package usecase

import (
	"dapp-moderator/external/bfs_service"
	"dapp-moderator/external/block_stream"
	"dapp-moderator/external/blockchain_api"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/internal/repository"
	"dapp-moderator/utils/config"
	discordclient "dapp-moderator/utils/discord"
	"dapp-moderator/utils/global"
	"dapp-moderator/utils/googlecloud"
	"dapp-moderator/utils/oauth2service"
	"dapp-moderator/utils/redis"
)

type Usecase struct {
	Repo          repository.Repository
	Config        *config.Config
	QuickNode     *quicknode.QuickNode
	BlockStream   *block_stream.BlockStream
	NftExplorer   *nft_explorer.NftExplorer
	TokenExplorer *token_explorer.TokenExplorer
	BfsService    *bfs_service.BfsService
	BnsService    *bns_service.BNSService
	Cache         redis.IRedisCache
	Auth2         *oauth2service.Auth2
	Storage       googlecloud.IGcstorage
	DiscordClient *discordclient.Client
	BlockChainApi *blockchain_api.BlockChainApi
}

func NewUsecase(global *global.Global, r repository.Repository) (*Usecase, error) {
	u := new(Usecase)
	u.Repo = r
	u.QuickNode = global.QuickNode
	u.BlockStream = global.BlockStream
	u.NftExplorer = global.NftExplorer
	u.BfsService = global.BfsService
	u.TokenExplorer = global.TokenExplorer
	u.Cache = global.Cache
	u.Storage = global.GCS
	u.Auth2 = global.Auth2
	u.BnsService = global.BnsService
	u.DiscordClient = global.DiscordClient
	u.BlockChainApi = global.BlockChainApi
	return u, nil
}

func (u *Usecase) Version() string {
	return "dAPP-API Server - version 1"
}
