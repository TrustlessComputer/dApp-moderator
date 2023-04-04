package usecase

import (
	"dapp-moderator/external/bfs_service"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/internal/repository"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/global"
)

type Usecase struct {
	Repo        repository.Repository
	Config      *config.Config
	QuickNode   *quicknode.QuickNode
	NftExplorer *nft_explorer.NftExplorer
	BfsService  *bfs_service.BfsService
}

func NewUsecase(global *global.Global, r repository.Repository) (*Usecase, error) {
	u := new(Usecase)
	u.Repo = r
	u.QuickNode = global.QuickNode
	u.NftExplorer = global.NftExplorer
	u.BfsService = global.BfsService
	return u, nil
}

func (uc *Usecase) Version() string {
	return "dAPP-API Server - version 1"
}
