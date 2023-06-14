package txTCServer

import (
	"context"
	"dapp-moderator/internal/usecase"
	"dapp-moderator/utils/logger"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
	"sync"
	"time"

	"dapp-moderator/utils/redis"

	"dapp-moderator/utils/global"
)

type trendingServer struct {
	Usecase       *usecase.Usecase
	Cache         redis.IRedisCache
	CronJobPeriod int32
}

func NewTrendingServer(global *global.Global, uc usecase.Usecase) (*trendingServer, error) {
	defaultPeriod := int32(600) //second - 10 min
	if os.Getenv("TRENDING_SERVER_PERIOD") != "" {
		p, err := strconv.Atoi(os.Getenv("TRENDING_SERVER_PERIOD"))
		if err == nil {
			defaultPeriod = int32(p)
		}
	}

	t := &trendingServer{
		Usecase:       &uc,
		Cache:         global.Cache,
		CronJobPeriod: defaultPeriod,
	}
	return t, nil
}

func (c *trendingServer) Task(wg *sync.WaitGroup, taskName string, processFunc func(ctx context.Context) error) {
	defer wg.Done()
	logger.AtLog.Logger.Info(fmt.Sprintf("trendingServer - Task: %s is running \n", taskName), zap.String("taskName", taskName))
	err := processFunc(context.Background())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("trendingServer - Task: %s is running \n", taskName), zap.String("taskName", taskName), zap.Error(err))
	}
}

func (c *trendingServer) StartServer() {

	tasks := make(map[string]func(ctx context.Context) error)
	//market-place
	tasks["trending_collections"] = c.UpdateCollectionMkplaceStat

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

func (c *trendingServer) UpdateCollectionMkplaceStat(ctx context.Context) error {
	return c.Usecase.AggregateCollectionStats()
}
