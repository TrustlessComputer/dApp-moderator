package txTCServer

import (
	"dapp-moderator/internal/usecase"
	"time"

	"dapp-moderator/utils/global"
)

type JobDisCord struct {
	UseCase *usecase.Usecase
}

func NewJobDisCord(global *global.Global, uc usecase.Usecase) *JobDisCord {

	t := &JobDisCord{
		UseCase: &uc,
	}
	return t
}

func (c *JobDisCord) StartServer() {
	for {
		c.UseCase.JobSendDiscord()
		time.Sleep(30 * time.Second)
	}
}
