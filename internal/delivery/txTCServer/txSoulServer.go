package txTCServer

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"dapp-moderator/utils/logger"
	"dapp-moderator/utils/redis"

	"dapp-moderator/utils/blockchain"
	"dapp-moderator/utils/global"

	"go.uber.org/zap"
)

type txSoulServer struct {
	//Txconsumer txconsumer.HttpTxConsumer
	Usecase                   *usecase.Usecase
	Cache                     redis.IRedisCache
	Blockchain                *blockchain.TcNetwork
	DefaultLastProcessedBlock int64
	CronJobPeriod             int32
	BatchLogSize              int32
	Soul                      *Soul
}

type Soul struct {
	Contract string
	Events   map[string]string
}

func NewTxSoulServer(global *global.Global, uc usecase.Usecase) (*txSoulServer, error) {
	startBlock := os.Getenv("TX_SOUL_CONSUMER_START_BLOCK")
	if startBlock == "" {
		startBlock = os.Getenv("TX_CONSUMER_CRON_JOB_PERIOD")
	}

	startBlockInt, err := strconv.Atoi(startBlock)
	if err != nil {
		startBlockInt = 1
	}

	period := os.Getenv("TX_SOUL_CONSUMER_CRON_JOB_PERIOD")
	if period == "" {
		period = os.Getenv("TX_CONSUMER_CRON_JOB_PERIOD")
	}
	periodInt, err := strconv.Atoi(period)
	if err != nil {
		periodInt = 15
	}

	blockBatchSize := os.Getenv("TX_SOUL_CONSUMER_BATCH_LOG_SIZE")
	if blockBatchSize == "" {
		blockBatchSize = os.Getenv("TX_CONSUMER_BATCH_LOG_SIZE")
	}

	blockBatchSizeInt, err := strconv.Atoi(blockBatchSize)
	if err != nil {
		blockBatchSizeInt = 100
	}

	bc, err := blockchain.NewTcNetwork()
	if err != nil {
		return nil, err
	}
	mkpEvents := make(map[string]string)

	mkpEvents["AUCTION_CREATED_EVENT"] = strings.ToLower(os.Getenv("AUCTION_CREATED_EVENT"))
	mkpEvents["AUCTION_BID_EVENT"] = strings.ToLower(os.Getenv("AUCTION_BID_EVENT"))
	mkpEvents["AUCTION_SETTLE_EVENT"] = strings.ToLower(os.Getenv("AUCTION_SETTLE_EVENT"))
	mkpEvents["AUCTION_CLAIM_EVENT"] = strings.ToLower(os.Getenv("AUCTION_CLAIM_EVENT"))

	//Move to mkplace event - to prevent block
	//mkpEvents["SOUL_UNLOCK_FEATURE_EVENT"] = strings.ToLower(os.Getenv("SOUL_UNLOCK_FEATURE_EVENT"))

	m := &Soul{
		Contract: os.Getenv("SOUL_CONTRACT"),
		Events:   mkpEvents,
	}

	t := &txSoulServer{
		Usecase:                   &uc,
		Cache:                     global.Cache,
		DefaultLastProcessedBlock: int64(startBlockInt),
		CronJobPeriod:             int32(periodInt),
		BatchLogSize:              int32(blockBatchSizeInt),
		Blockchain:                bc,
		Soul:                      m,
	}

	return t, nil
}

func (c *txSoulServer) getLastProcessedBlock(ctx context.Context) (int64, error) {

	lastProcessed := c.DefaultLastProcessedBlock
	redisKey := c.getRedisKey(nil)
	//c.Cache.Delete(redisKey)
	exists, err := c.Cache.Exists(redisKey)
	if err != nil {
		logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
		return 0, err
	}
	if *exists {
		processed, err := c.Cache.GetData(redisKey)
		if err != nil {
			logger.AtLog.Logger.Error("error get from redis", zap.Error(err))
			return 0, err
		}
		if processed == nil {
			return (c.DefaultLastProcessedBlock), nil
		}
		lastProcessedSavedOnRedis, err := strconv.ParseInt(*processed, 10, 64)
		if err != nil {
			logger.AtLog.Logger.Error("err.getLastProcessedBlock", zap.Error(err))
			return 0, err
		}
		lastProcessed = int64(math.Max(float64(lastProcessed), float64(lastProcessedSavedOnRedis)))
	}
	return lastProcessed, nil
}

func (c *txSoulServer) getRedisKey(postfix *string) string {
	key := fmt.Sprintf("tx-consumer:soul:latest_processed")
	if postfix != nil {
		key = fmt.Sprintf("%s_%s", key, *postfix)
	}
	return key
}

func (c *txSoulServer) Task(wg *sync.WaitGroup, taskName string, processFunc func(ctx context.Context) error) {
	defer wg.Done()

	err := processFunc(context.Background())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("Task: %s is running", taskName), zap.String("taskName", taskName), zap.Error(err))
	}
}

func (c *txSoulServer) StartServer() {

	tasks := make(map[string]func(ctx context.Context) error)
	tasks["soul-tx-transaction"] = c.resolveTxTransaction

	var wg sync.WaitGroup
	for {
		// tasks ==> start
		for key, task := range tasks {
			wg.Add(1)
			go c.Task(&wg, key, task)
		}
		wg.Wait()

		time.Sleep(time.Duration(c.CronJobPeriod) * time.Second)
	}
}

func (c *txSoulServer) resolveTxTransaction(ctx context.Context) error {
	lastProcessedBlock, err := c.getLastProcessedBlock(ctx)
	if err != nil {
		logger.AtLog.Logger.Error("txSoulServer - resolveTransaction", zap.Error(err))
		return err
	}

	// get new block from db
	chainBlock, err := c.Blockchain.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("txSoulServer - resolveTransaction", zap.Error(err))
		return err
	}

	fromBlock := lastProcessedBlock + 1
	if fromBlock > chainBlock.Int64() {
		return err
	}

	toBlock := int64(math.Min(float64(chainBlock.Int64()), float64(fromBlock+int64(c.BatchLogSize))))
	if toBlock < fromBlock {
		fromBlock = toBlock
	}

	//Logs - start
	//logs are only heard by our collection addresses
	logs, err := c.Blockchain.GetEventLogs(*big.NewInt(fromBlock), *big.NewInt(toBlock), c.ListenedContracts())
	if err != nil {
		logger.AtLog.Logger.Error("txSoulServer - resolveTransaction",
			zap.String("err", err.Error()),
			zap.Int64("fromBlock", fromBlock),
			zap.Int64("toBlock", toBlock),
			zap.Int64("chainBlock",
				chainBlock.Int64()),
		)
		return err
	}

	logger.AtLog.Logger.Info("txSoulServer - resolveTransaction",
		zap.Int64("fromBlock", fromBlock),
		zap.Int64("toBlock", toBlock),
		zap.Int("logs", len(logs)))

	//start channel
	inputDataChan := make(chan types.Log, len(logs))
	resultChan := make(chan *eventLog, len(logs))
	var wg sync.WaitGroup

	//init workers (pool)
	for _, _ = range logs {
		go c.Worker(&wg, inputDataChan, resultChan)
	}

	for i, log := range logs {
		wg.Add(1)
		inputDataChan <- log
		if i > 0 && i%100 == 0 || i == len(logs) {
			wg.Wait()
		}
	}

	for _, _ = range logs {
		c.ProcessLog(resultChan)
	}

	err = c.Cache.SetStringData(c.getRedisKey(nil), strconv.FormatInt(toBlock, 10))
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Error(err))
		return err
	}

	return nil
}

// the contract, that will be listened (SOUL contract address)
func (c *txSoulServer) ListenedContracts() []common.Address {
	addresses := []common.Address{}
	soulContract := common.HexToAddress(c.Soul.Contract)
	addresses = append(addresses, soulContract)
	return addresses
}

func (c *txSoulServer) Worker(wg *sync.WaitGroup, inputDataChan chan types.Log, result chan *eventLog) {
	defer wg.Done()

	log := <-inputDataChan

	topics := log.Topics
	event := topics[0]
	mkpEventName := strings.ToLower(event.String())

	var pFunction proccessFunc
	var eventType entity.TokenActivityType
	var err error
	activity := &entity.MarketplaceTokenActivity{}
	var eventData interface{}

	switch mkpEventName {

	case c.Soul.Events["AUCTION_CREATED_EVENT"]:
		pFunction = c.Usecase.HandleAuctionCreated
		eventType = entity.AuctionCreatedActivity
		break
	case c.Soul.Events["AUCTION_BID_EVENT"]:
		pFunction = c.Usecase.HandleAuctionBid
		eventType = entity.AuctionBidActivity
		break
	case c.Soul.Events["AUCTION_SETTLE_EVENT"]:
		pFunction = c.Usecase.HandleAuctionSettle
		eventType = entity.AuctionSettledActivity
		break
	case c.Soul.Events["AUCTION_CLAIM_EVENT"]:
		pFunction = c.Usecase.HandleAuctionClaim
		eventType = entity.AuctionClaimActivity
		break

		//move to mkplace to prevent blocking
		//case c.Soul.Events["SOUL_UNLOCK_FEATURE_EVENT"]:
		//	pFunction = c.Usecase.HandleUnlockFeature
		//	eventType = entity.SoulUnlockFeature
		//	break
	}

	activity, eventData, err = c.Usecase.ParseMkplaceData(log, eventType)
	result <- &eventLog{
		Log:       log,
		EventData: eventData,
		Event:     eventType,
		Activity:  activity,
		Err:       err,
		PFunc:     pFunction,
	}

}

func (c *txSoulServer) ProcessLog(resultChan chan *eventLog) {
	dataFromChan := <-resultChan
	err := dataFromChan.Err
	if err == nil {
		_func := dataFromChan.PFunc
		err := _func(dataFromChan.EventData, dataFromChan.Log)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("txSoulServer - Proccess funnc - %s", dataFromChan.Log.TxHash), zap.Error(err))
		}

		activity := dataFromChan.Activity
		err = c.Usecase.InsertActivity(activity)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("txSoulServer - InsertActivity - %s", dataFromChan.Log.TxHash), zap.Error(err))
		}
	}
}
