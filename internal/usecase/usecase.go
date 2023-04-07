package usecase

import (
	"dapp-moderator/external/bfs_service"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/internal/repository"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/global"
	"dapp-moderator/utils/oauth2service"
	"dapp-moderator/utils/redis"
)

type Usecase struct {
	Repo          repository.Repository
	Config        *config.Config
	QuickNode     *quicknode.QuickNode
	NftExplorer   *nft_explorer.NftExplorer
	TokenExplorer *token_explorer.TokenExplorer
	BfsService    *bfs_service.BfsService
	Cache         redis.IRedisCache
	Auth2         *oauth2service.Auth2
}

func NewUsecase(global *global.Global, r repository.Repository) (*Usecase, error) {
	u := new(Usecase)
	u.Repo = r
	u.QuickNode = global.QuickNode
	u.NftExplorer = global.NftExplorer
	u.BfsService = global.BfsService
	u.TokenExplorer = global.TokenExplorer
	u.Cache = global.Cache
	return u, nil
}

func (c *Usecase) Version() string {
	return "dAPP-API Server - version 1"
}
