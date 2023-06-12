package txTCServer

import (
	"context"
	"dapp-moderator/internal/usecase"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

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
	MarketPlace               *MarketPlace
}

type MarketPlace struct {
	Contract string
	Events   map[string]string
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
	mkpEvents := make(map[string]string)
	mkpEvents["MARKETPLACE_TRANSFER_EVENT"] = strings.ToLower(os.Getenv("MARKETPLACE_TRANSFER_EVENT"))
	mkpEvents["MARKETPLACE_LIST_TOKEN"] = strings.ToLower(os.Getenv("MARKETPLACE_LIST_TOKEN"))
	mkpEvents["MARKETPLACE_PURCHASE_TOKEN"] = strings.ToLower(os.Getenv("MARKETPLACE_PURCHASE_TOKEN"))
	mkpEvents["MARKETPLACE_MAKE_OFFER"] = strings.ToLower(os.Getenv("MARKETPLACE_MAKE_OFFER"))
	mkpEvents["MARKETPLACE_ACCEPT_MAKE_OFFER"] = strings.ToLower(os.Getenv("MARKETPLACE_ACCEPT_MAKE_OFFER"))
	mkpEvents["MARKETPLACE_CANCEL_MAKE_OFFER"] = strings.ToLower(os.Getenv("MARKETPLACE_CANCEL_MAKE_OFFER"))
	mkpEvents["MARKETPLACE_CANCEL_LISTING"] = strings.ToLower(os.Getenv("MARKETPLACE_CANCEL_LISTING"))
	mkpEvents["MARKETPLACE_BNS_RESOLVER_UPDATED"] = strings.ToLower(os.Getenv("MARKETPLACE_BNS_RESOLVER_UPDATED"))
	mkpEvents["MARKETPLACE_BNS_REGISTERED"] = strings.ToLower(os.Getenv("MARKETPLACE_BNS_REGISTERED"))
	mkpEvents["MARKETPLACE_BNS_SET_FPF"] = strings.ToLower(os.Getenv("MARKETPLACE_BNS_SET_FPF"))

	m := &MarketPlace{
		Contract: os.Getenv("MARKETPLACE_CONTRACT"),
		Events:   mkpEvents,
	}

	t := &txTCServer{
		Usecase:                   &uc,
		Cache:                     global.Cache,
		DefaultLastProcessedBlock: int64(startBlockInt),
		CronJobPeriod:             int32(periodInt),
		BatchLogSize:              int32(blockBatchSizeInt),
		Blockchain:                bc,
		MarketPlace:               m,
	}

	return t, nil
}

func (c *txTCServer) getLastProcessedBlock(ctx context.Context) (int64, error) {

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

func (c *txTCServer) getRedisKey(postfix *string) string {
	key := fmt.Sprintf("tx-consumer:latest_processed")
	if postfix != nil {
		key = fmt.Sprintf("%s_%s", key, *postfix)
	}
	return key
}

func (c *txTCServer) Task(wg *sync.WaitGroup, taskName string, processFunc func(ctx context.Context) error) {
	defer wg.Done()
	logger.AtLog.Logger.Info(fmt.Sprintf("Task: %s is running \n", taskName), zap.String("taskName", taskName))
	err := processFunc(context.Background())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("Task: %s is running \n", taskName), zap.String("taskName", taskName), zap.Error(err))
	}
}

func (c *txTCServer) StartServer() {

	tasks := make(map[string]func(ctx context.Context) error)
	//function is being developed
	tasks["checkSoulOwnerCrontab"] = c.checkSoulOwnerCrontab

	//function have been done in develop
	if os.Getenv("ENV") == "production" {
		tasks["checkTxHashChunks"] = c.checkTxHashChunks
		tasks["resolveTxTransaction"] = c.resolveTxTransaction
		tasks["fetchToken"] = c.fetchToken
		tasks["UpdateCollectionItems"] = c.Usecase.UpdateCollectionItems
		tasks["UpdateCollectionThumbnails"] = c.Usecase.UpdateCollectionThumbnails
	}

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

func (c *txTCServer) resolveTxTransaction(ctx context.Context) error {
	lastProcessedBlock, err := c.getLastProcessedBlock(ctx)
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Error(err))
		return err
	}

	// get new block from db
	chainBlock, err := c.Blockchain.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Error(err))
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

	logger.AtLog.Logger.Info("resolveTransaction", zap.Int64("fromBlock", fromBlock), zap.Int64("toBlock", toBlock), zap.Int64("chainBlock", chainBlock.Int64()))

	var wg sync.WaitGroup
	wg.Add(2)

	go c.TokenEvents(&wg, ctx, int64(fromBlock), int64(toBlock))

	go c.processTxTransaction(&wg, ctx, int32(fromBlock), int32(toBlock))

	wg.Wait()

	err = c.Cache.SetStringData(c.getRedisKey(nil), strconv.FormatInt(toBlock, 10))
	if err != nil {
		logger.AtLog.Logger.Error("resolveTransaction", zap.Error(err))
		return err
	}

	return nil
}

func (c *txTCServer) processTxTransaction(wg *sync.WaitGroup, ctx context.Context, fromBlock int32, toBlock int32) {
	defer wg.Done()

	c.Usecase.GetCollectionFromBlock(ctx, fromBlock, toBlock)
}

func (c *txTCServer) fetchToken(ctx context.Context) error {
	tokenPage := "tokens_page"
	key := c.getRedisKey(&tokenPage)

	lastFetchedPage, err := c.Cache.GetData(key)
	if err != nil && err != redis2.Nil {
		logger.AtLog.Logger.Error("fetchToken", zap.Error(err))
		return err
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
		return err
	}

	err = c.Cache.SetStringData(key, strconv.Itoa(toPage))
	if err != nil {
		logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
		return err
	}

	return nil
}

func (c *txTCServer) checkTxHashChunks(ctx context.Context) error {
	return c.Usecase.ListenedChunks()
}

func (c *txTCServer) checkSoulOwnerCrontab(ctx context.Context) error {
	c.Usecase.SoulCrontab()
	return nil
}

// the contracts, that will be listened (collection address erc721) + marketplace contract
func (c *txTCServer) ListenedContracts() []common.Address {
	addresses := []common.Address{}
	mkpContract := common.HexToAddress(c.MarketPlace.Contract)
	addresses = append(addresses, mkpContract)

	allCollections, err := c.Usecase.Repo.AllCollections()
	if err != nil {
		logger.AtLog.Logger.Error("err.GetEventLogs", zap.String("err", err.Error()))
		return addresses
	}

	for _, collection := range allCollections {
		hexAddress := common.HexToAddress(collection.Contract)
		addresses = append(addresses, hexAddress)
	}

	return addresses
}
