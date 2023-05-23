package txTCServer

import (
	"context"
	"dapp-moderator/internal/usecase"
	"dapp-moderator/utils/generative_nft_contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	redis2 "github.com/go-redis/redis"

	"dapp-moderator/utils/logger"
	"dapp-moderator/utils/redis"

	"dapp-moderator/utils/blockchain"
	"dapp-moderator/utils/global"

	"go.uber.org/zap"
)

type txTCServer struct {
	//Txconsumer txconsumer.HttpTxConsumer
	Usecase                   *usecase.Usecase
	Cache                     redis.IRedisCache
	Blockchain                *blockchain.TcNetwork
	DefaultLastProcessedBlock int64
	CronJobPeriod             int32
	BatchLogSize              int32
}

type eventLog struct {
	Log   types.Log
	Data  []common.Hash
	Event string
}

func NewTxTCServer(global *global.Global, uc usecase.Usecase) (*txTCServer, error) {
	startBlock := os.Getenv("TX_CONSUMER_START_BLOCK")
	startBlockInt, err := strconv.Atoi(startBlock)
	if err != nil {
		startBlockInt = 1
	}

	period := os.Getenv("TX_CONSUMER_CRON_JOB_PERIOD")
	periodInt, err := strconv.Atoi(period)
	if err != nil {
		periodInt = 15
	}

	blockBatchSize := os.Getenv("TX_CONSUMER_BATCH_LOG_SIZE")
	blockBatchSizeInt, err := strconv.Atoi(blockBatchSize)
	if err != nil {
		blockBatchSizeInt = 100
	}

	bc, err := blockchain.NewTcNetwork()
	if err != nil {
		return nil, err
	}

	t := &txTCServer{
		Usecase:                   &uc,
		Cache:                     global.Cache,
		DefaultLastProcessedBlock: int64(startBlockInt),
		CronJobPeriod:             int32(periodInt),
		BatchLogSize:              int32(blockBatchSizeInt),
		Blockchain:                bc,
	}

	return t, nil
}

func (c *txTCServer) getLastProcessedBlock(ctx context.Context) (int64, error) {

	lastProcessed := c.DefaultLastProcessedBlock
	redisKey := c.getRedisKey(nil)
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

func (c *txTCServer) getRedisKey(postfix *string) string {
	key := fmt.Sprintf("tx-consumer:latest_processed")
	if postfix != nil {
		key = fmt.Sprintf("%s_%s", key, *postfix)
	}
	return key
}

func (c *txTCServer) StartServer() {
	ctx := context.Background()
	for {
		previousTime := time.Now()
		var wg sync.WaitGroup
		wg.Add(4)
		go func() {
			defer wg.Done()
			c.resolveTxTransaction(ctx)
		}()

		go func() {
			defer wg.Done()
			c.fetchToken(ctx)
		}()

		go func() {
			defer wg.Done()
			c.Usecase.UpdateCollectionItems(ctx)
		}()

		go func() {
			defer wg.Done()
			c.Usecase.UpdateCollectionThumbnails(ctx)
		}()

		wg.Wait()
		processedTime := time.Now().Unix() - previousTime.Unix()
		if processedTime < int64(c.CronJobPeriod) {
			time.Sleep(time.Duration(c.CronJobPeriod-int32(processedTime)) * time.Second)
		}
		logger.AtLog.Logger.Info("StartServer", zap.Any("previousTime", previousTime), zap.Any("processedTime", processedTime))
	}
}

func (c *txTCServer) Worker(inputDataChan chan types.Log, result chan *eventLog) {
	log := <-inputDataChan

	topics := log.Topics
	event := topics[0]

	result <- &eventLog{
		Log:   log,
		Data:  topics,
		Event: event.String(),
	}

}

func (c *txTCServer) ProcessLog(resultChan chan *eventLog) {
	dataFromChan := <-resultChan
	switch dataFromChan.Event {

	//Transfer event
	case strings.ToLower("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"):
		if len(dataFromChan.Data) == 4 {
			client := c.Blockchain

			contract := dataFromChan.Log.Address.Hex()

			erc21, err := generative_nft_contract.NewGenerativeNftContract(dataFromChan.Log.Address, client.GetClient())
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("UpdateNftOwner %s  ", contract), zap.Error(err))
				return
			}

			erc21Transfer, err := erc21.ParseTransfer(dataFromChan.Log)
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("UpdateNftOwner %s  ParseTransfer", contract), zap.Error(err))
				return
			}

			from := erc21Transfer.From.Hex()
			to := erc21Transfer.To.Hex()
			tokenIDStr := dataFromChan.Data[3].Big().String()

			if strings.ToLower(os.Getenv("ENV")) == strings.ToLower("production") {
				updated, err := c.Usecase.UpdateNftOwner(context.Background(), contract, tokenIDStr, to)
				if err != nil {
					logger.AtLog.Logger.Error(fmt.Sprintf("UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Uint64("blockNumber", dataFromChan.Log.BlockNumber), zap.Error(err))
					return
				}

				logger.AtLog.Logger.Info(fmt.Sprintf("UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Any("updated", updated), zap.Uint64("blockNumber", dataFromChan.Log.BlockNumber))
			} else {
				logger.AtLog.Logger.Info(fmt.Sprintf("[Testing] UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Uint64("blockNumber", dataFromChan.Log.BlockNumber))
			}

		}
		break
	}
}

func (c *txTCServer) resolveTxTransaction(ctx context.Context) error {
	lastProcessedBlock, err := c.getLastProcessedBlock(ctx)
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Any("err", err))
		return err
	}

	fromBlock := lastProcessedBlock + 1
	blockNumber, err := c.Blockchain.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Any("err", err))
		return err
	}

	toBlock := int64(math.Min(float64(blockNumber.Int64()), float64(fromBlock+int64(c.BatchLogSize))))
	if toBlock < fromBlock {
		fromBlock = toBlock
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go c.TransferTokenEvents(&wg, ctx, int64(fromBlock), int64(toBlock))

	go c.processTxTransaction(&wg, ctx, int32(fromBlock), int32(toBlock))

	wg.Wait()
	err = c.Cache.SetStringData(c.getRedisKey(nil), strconv.FormatInt(toBlock, 10))
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Error(err))
		return err
	}
	logger.AtLog.Logger.Info("resolveTransaction", zap.Int64("lastProcessedBlock", lastProcessedBlock), zap.Int64("blockNumber", blockNumber.Int64()), zap.Int64("fromBlock", fromBlock), zap.Int64("toBlock", toBlock))
	return nil
}

func (c *txTCServer) processTxTransaction(wg *sync.WaitGroup, ctx context.Context, fromBlock int32, toBlock int32) {
	defer wg.Done()

	c.Usecase.GetCollectionFromBlock(ctx, fromBlock, toBlock)
}

func (c *txTCServer) fetchToken(ctx context.Context) {
	tokenPage := "tokens_page"
	key := c.getRedisKey(&tokenPage)

	lastFetchedPage, err := c.Cache.GetData(key)
	if err != nil && err != redis2.Nil {
		logger.AtLog.Logger.Error("fetchToken", zap.Error(err))
		return
	}
	fromPage := 1
	if lastFetchedPage != nil {
		fromPage, err = strconv.Atoi(*lastFetchedPage)
		if err != nil {
			fromPage = 1
		}
	}

	toPage, err := c.Usecase.CrawToken(ctx, fromPage)
	if err != nil {
		logger.AtLog.Logger.Error("failed to fetch token from token-explorer", zap.Error(err))
		return
	}

	err = c.Cache.SetStringData(key, strconv.Itoa(toPage))
	if err != nil {
		logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
		return
	}
}

func (c *txTCServer) TransferTokenEvents(wg *sync.WaitGroup, ctx context.Context, fromBlock int64, toBlock int64) {
	defer wg.Done()

	allCollections, err := c.Usecase.Repo.AllCollections()
	if err != nil {
		logger.AtLog.Logger.Error("err.GetEventLogs", zap.String("err", err.Error()))
		return
	}

	addresses := []common.Address{}
	for _, collection := range allCollections {
		hexAddress := common.HexToAddress(collection.Contract)
		addresses = append(addresses, hexAddress)
	}

	//logs are only heard by our collection addresses
	logs, err := c.Blockchain.GetEventLogs(*big.NewInt(fromBlock), *big.NewInt(toBlock), addresses)
	if err != nil {
		logger.AtLog.Logger.Error("err.GetEventLogs", zap.String("err", err.Error()))
		return
	}

	//start channel
	inputDataChan := make(chan types.Log, len(logs))
	resultChan := make(chan *eventLog, len(logs))

	//init workers
	for _, _ = range logs {
		go c.Worker(inputDataChan, resultChan)
	}

	// push data to worker
	for i, log := range logs {
		inputDataChan <- log
		if i > 0 && i%100 == 0 {
			time.Sleep(time.Millisecond * 500)
		}
	}

	//processing
	for _, _ = range logs {
		c.ProcessLog(resultChan)
	}
}
