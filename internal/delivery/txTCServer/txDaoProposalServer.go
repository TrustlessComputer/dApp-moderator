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

type txDaoProposalServer struct {
	//Txconsumer txconsumer.HttpTxConsumer
	Usecase                   *usecase.Usecase
	Cache                     redis.IRedisCache
	Blockchain                *blockchain.TcNetwork
	DefaultLastProcessedBlock int64
	CronJobPeriod             int32
	BatchLogSize              int32
	Dao                       *DaoProposal
}

type DaoProposal struct {
	Contract string
	Events   map[string]string
}

func NewTxDaoProposalServer(global *global.Global, uc usecase.Usecase) (*txDaoProposalServer, error) {
	startBlock := os.Getenv("DAO_CONSUMER_START_BLOCK")
	startBlockInt, err := strconv.Atoi(startBlock)
	if err != nil {
		startBlockInt = 1
	}

	period := os.Getenv("DAO_CONSUMER_CRON_JOB_PERIOD")
	periodInt, err := strconv.Atoi(period)
	if err != nil {
		periodInt = 15
	}

	blockBatchSize := os.Getenv("DAO_CONSUMER_BATCH_LOG_SIZE")
	blockBatchSizeInt, err := strconv.Atoi(blockBatchSize)
	if err != nil {
		blockBatchSizeInt = 100
	}

	bc, err := blockchain.NewTcNetwork()
	if err != nil {
		return nil, err
	}
	daoEvents := make(map[string]string)
	daoEvents["DAO_PROPOSAL_CREATED"] = strings.ToLower(os.Getenv("DAO_PROPOSAL_CREATED"))
	daoEvents["DAO_PROPOSAL_CAST_VOTE"] = strings.ToLower(os.Getenv("DAO_PROPOSAL_CAST_VOTE"))

	m := &DaoProposal{
		Events: daoEvents,
	}

	t := &txDaoProposalServer{
		Usecase:                   &uc,
		Cache:                     global.Cache,
		DefaultLastProcessedBlock: int64(startBlockInt),
		CronJobPeriod:             int32(periodInt),
		BatchLogSize:              int32(blockBatchSizeInt),
		Blockchain:                bc,
		Dao:                       m,
	}

	return t, nil
}

func (c *txDaoProposalServer) getLastProcessedBlock(ctx context.Context) (int64, error) {

	lastProcessed := c.DefaultLastProcessedBlock
	redisKey := c.getRedisKey(nil)
	//c.Cache.Delete(redisKey)
	exists, err := c.Cache.Exists(redisKey)
	if err != nil {
		logger.AtLog.Logger.Error("txDaoProposalServer", zap.String("redisKey", redisKey), zap.Error(err))
		return 0, err
	}
	if *exists {
		processed, err := c.Cache.GetData(redisKey)
		if err != nil {
			logger.AtLog.Logger.Error("txDaoProposalServer", zap.Error(err))
			return 0, err
		}
		if processed == nil {
			return (c.DefaultLastProcessedBlock), nil
		}
		lastProcessedSavedOnRedis, err := strconv.ParseInt(*processed, 10, 64)
		if err != nil {
			logger.AtLog.Logger.Error("txDaoProposalServer", zap.Error(err))
			return 0, err
		}
		lastProcessed = int64(math.Max(float64(lastProcessed), float64(lastProcessedSavedOnRedis)))
	}
	return lastProcessed, nil
}

func (c *txDaoProposalServer) getRedisKey(postfix *string) string {
	key := fmt.Sprintf("dao-consumer:latest_processed")
	if postfix != nil {
		key = fmt.Sprintf("%s_%s", key, *postfix)
	}
	return key
}

func (c *txDaoProposalServer) Task(wg *sync.WaitGroup, taskName string, processFunc func(ctx context.Context) error) {
	defer wg.Done()
	logger.AtLog.Logger.Info(fmt.Sprintf("Task: %s is running \n", taskName), zap.String("taskName", taskName))
	err := processFunc(context.Background())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("Task: %s is running \n", taskName), zap.String("taskName", taskName), zap.Error(err))
	}
}

func (c *txDaoProposalServer) StartServer() {

	tasks := make(map[string]func(ctx context.Context) error)
	tasks["UpdateCollectionThumbnails"] = c.resolveTxTransaction

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

func (c *txDaoProposalServer) resolveTxTransaction(ctx context.Context) error {
	lastProcessedBlock, err := c.getLastProcessedBlock(ctx)
	if err != nil {
		logger.AtLog.Logger.Error("txDaoProposalServer", zap.Error(err))
		return err
	}

	// get new block from db
	chainBlock, err := c.Blockchain.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("txDaoProposalServer", zap.Error(err))
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

	logs, err := c.Blockchain.GetEventLogs(*big.NewInt(fromBlock), *big.NewInt(toBlock), c.ListenedContracts())
	if err != nil {
		logger.AtLog.Logger.Error("txDaoProposalServer", zap.Int64("fromBlock", fromBlock), zap.Int64("toBlock", toBlock), zap.Int64("chainBlock", chainBlock.Int64()), zap.String("err", err.Error()))
		return err
	}

	logger.AtLog.Logger.Info("txDaoProposalServer", zap.Int64("fromBlock", fromBlock), zap.Int64("toBlock", toBlock), zap.Int64("chainBlock", chainBlock.Int64()), zap.Int("logs", len(logs)))
	//loop

	//start channel
	inputDataChan := make(chan types.Log, len(logs))
	resultChan := make(chan *daoEventLog, len(logs))
	var wg sync.WaitGroup
	//init workers (pool)
	for _, _ = range logs {
		go c.Worker(&wg, inputDataChan, resultChan)
	}

	// push data to worker, process each 100 logs
	for i, log := range logs {
		wg.Add(1)
		inputDataChan <- log
		if (i > 0 && i%100 == 0) || (i == len(logs)-1) {
			wg.Wait()
		}
	}

	//processing
	for _, _ = range logs {
		c.ProcessLog(resultChan)
	}

	err = c.Cache.SetStringData(c.getRedisKey(nil), strconv.FormatInt(toBlock, 10))
	if err != nil {
		logger.AtLog.Logger.Error("txDaoProposalServer", zap.Error(err))
		return err
	}

	return nil
}

// the contracts, that will be listened (collection address erc721) + marketplace contract
func (c *txDaoProposalServer) ListenedContracts() []common.Address {
	addresses := []common.Address{}
	contracts, err := c.Usecase.Repo.ListenedDaoContract()
	if err != nil {
		logger.AtLog.Logger.Error("txDaoProposalServer", zap.String("err", err.Error()))
		return addresses
	}

	logger.AtLog.Logger.Error("txDaoProposalServer", zap.Int("contracts", len(contracts)))
	for _, contract := range contracts {
		hexAddress := common.HexToAddress(contract.ContractAddress)
		addresses = append(addresses, hexAddress)
	}
	return addresses
}

type daoEventLog struct {
	Log       types.Log
	EventData interface{}
	Err       error
	PFunc     proccessFunc
	BlockInfo *types.Block
}

func (c *txDaoProposalServer) Worker(wg *sync.WaitGroup, inputDataChan chan types.Log, result chan *daoEventLog) {
	defer wg.Done()
	log := <-inputDataChan

	topics := log.Topics
	event := topics[0]
	daoEventName := strings.ToLower(event.String())

	var pFunction proccessFunc
	var eventType entity.TokenActivityType
	var err error
	var eventData interface{}

	//parse dao data
	pFunction, eventData, blockInfo, err := c.Usecase.ParseDao(log, eventType, daoEventName)
	result <- &daoEventLog{
		Log:       log,
		EventData: eventData,
		Err:       err,
		PFunc:     pFunction,
		BlockInfo: blockInfo,
	}

}

func (c *txDaoProposalServer) ProcessLog(resultChan chan *daoEventLog) {
	dataFromChan := <-resultChan
	err := dataFromChan.Err
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("txDaoProposalServer.ProcessLog - %s", dataFromChan.Log.TxHash), zap.Error(err))
		return
	}

	_func := dataFromChan.PFunc
	err = _func(dataFromChan.EventData, dataFromChan.Log)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("txDaoProposalServer.Proccess funnc - %s", dataFromChan.Log.TxHash), zap.Error(err))
		return
	}
	logger.AtLog.Logger.Info(fmt.Sprintf("txDaoProposalServer.Proccess funnc - %s", dataFromChan.Log.TxHash), zap.Any("dataFromChan", dataFromChan))
}
