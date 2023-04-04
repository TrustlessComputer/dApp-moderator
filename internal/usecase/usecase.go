package usecase

import (
	"dapp-moderator/external/quicknode"
	"dapp-moderator/internal/repository"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/global"
)

type Usecase struct {
	Repo      repository.Repository
	Config    *config.Config
	QuickNode *quicknode.QuickNode
}

func NewUsecase(global *global.Global, r repository.Repository) (*Usecase, error) {
	u := new(Usecase)
	u.Repo = r
	u.QuickNode = global.QuickNode
	return u, nil
}

func (uc *Usecase) Version() string {
	return "dAPP-API Server - version 1"
}
