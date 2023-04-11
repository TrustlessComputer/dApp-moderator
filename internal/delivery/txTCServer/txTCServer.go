package txTCServer

import (
	"context"
	"dapp-moderator/internal/usecase"
	"fmt"
	"math"
	"os"
	"strconv"
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
		wg.Add(3)
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

		wg.Wait()
		processedTime := time.Now().Unix() - previousTime.Unix()
		if processedTime < int64(c.CronJobPeriod) {
			time.Sleep(time.Duration(c.CronJobPeriod-int32(processedTime)) * time.Second)
		}
		logger.AtLog.Logger.Info("StartServer", zap.Any("previousTime", previousTime), zap.Any("processedTime", processedTime))
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

	// logs, err := c.Blockchain.GetEventLogs(*big.NewInt(fromBlock), *big.NewInt(toBlock), c.Addresses)
	// if err != nil {
	// 	logger.AtLog.Logger.Error("err.GetEventLogs", zap.String("err", err.Error()))
	// 	return err
	// }

	c.processTxTransaction(ctx, int32(fromBlock), int32(toBlock))

	err = c.Cache.SetStringData(c.getRedisKey(nil), strconv.FormatInt(toBlock, 10))
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Error(err))
		return err
	}
	logger.AtLog.Logger.Info("resolveTransaction", zap.Int64("lastProcessedBlock", lastProcessedBlock), zap.Int64("blockNumber", blockNumber.Int64()), zap.Int64("fromBlock", fromBlock), zap.Int64("toBlock", toBlock))
	return nil
}

func (c *txTCServer) processTxTransaction(ctx context.Context, fromBlock int32, toBlock int32) {
	c.Usecase.GetCollectionFromBlock(ctx, fromBlock, toBlock)
}

func (c *txTCServer) fetchToken(ctx context.Context) {
	tokenPage := "token_page"
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

	fromPage = 1
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
