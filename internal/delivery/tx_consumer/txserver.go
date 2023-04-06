package txserver

import (
	"dapp-moderator/internal/usecase"
	"fmt"
	"math"
	"os"
	"strconv"

	"dapp-moderator/utils/config"
	"dapp-moderator/utils/logger"
	"dapp-moderator/utils/redis"

	"dapp-moderator/utils/global"

	"go.uber.org/zap"
)

type TxServer struct {
	//Txconsumer txconsumer.HttpTxConsumer
	Usecase    usecase.Usecase
	Config     *config.Config
	Cache      redis.IRedisCache

	DefaultLastProcessedBlock int64
	CronJobPeriod             int32
	BatchLogSize              int32
}

func NewTxServer(global *global.Global, uc usecase.Usecase, cfg config.Config) (*TxServer, error) {
	startBlock :=  os.Getenv("TX_CONSUMER_START_BLOCK")
	startBlockInt, err := strconv.Atoi(startBlock)
	if err != nil {
		startBlockInt = 1
	}
	
	period :=  os.Getenv("TX_CONSUMER_CRON_JOB_PERIOD")
	periodInt, err := strconv.Atoi(period)
	if err != nil {
		periodInt = 15
	}
	
	blockBatchSize :=  os.Getenv("TX_CONSUMER_BATCH_LOG_SIZE")
	blockBatchSizeInt, err := strconv.Atoi(blockBatchSize)
	if err != nil {
		blockBatchSizeInt = 100
	}

	t := &TxServer{
		Usecase: uc,
		Config:  &cfg,
		Cache:   global.Cache,
		DefaultLastProcessedBlock: int64(startBlockInt),
		CronJobPeriod: int32(periodInt),
		BatchLogSize: int32(blockBatchSizeInt),
	}

	return t, nil
}

func (c *TxServer) getLastProcessedBlock() (int64, error) {

	lastProcessed := c.DefaultLastProcessedBlock
	redisKey := c.getRedisKey()
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

func (c *TxServer) getRedisKey() string {
	return fmt.Sprintf("tx-consumer:latest_processed")
}

func (tx TxServer) StartServer() {
	for {


	}
}
