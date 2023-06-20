package txTCServer

import (
	"dapp-moderator/internal/usecase"
	"dapp-moderator/utils/global"
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
	for {
		c.UseCase.SoulNftImageHistoriesCrontab()
		time.Sleep(time.Minute * 20) //20 min
	}
}
