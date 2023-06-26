package txTCServer

import (
	"context"
	"dapp-moderator/internal/usecase"
	"dapp-moderator/utils/global"
	"sync"
	"time"
)

type JobCaptureThumbnailHistories struct {
	UseCase *usecase.Usecase
}

func NewJobCaptureThumbnailHistories(global *global.Global, uc usecase.Usecase) *JobCaptureThumbnailHistories {
	t := &JobCaptureThumbnailHistories{
		UseCase: &uc,
	}
	return t
}

func (c *JobCaptureThumbnailHistories) StartServer() {
	var wg sync.WaitGroup
	for {

		wg.Add(1)
		//go func(wg *sync.WaitGroup) {
		//	defer wg.Done()
		//	//c.UseCase.SoulNftImageHistoriesCrontab()
		//}(&wg)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			c.UseCase.CrontabUpdateImageSize(context.Background())
		}(&wg)

		wg.Wait()
		time.Sleep(time.Minute * 20) //20 min
	}
}
