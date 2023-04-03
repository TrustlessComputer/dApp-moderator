package usecase

import (
	"dapp-moderator/internal/repository"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/global"
)


type Usecase struct {
	Repo                repository.Repository
	Config              *config.Config
}

func NewUsecase(global *global.Global, r repository.Repository) (*Usecase, error) {
	u := new(Usecase)
	u.Repo = r
	return u, nil
}

func (uc *Usecase) Version() string {
	return "Generateve-API Server - version 1"
}
