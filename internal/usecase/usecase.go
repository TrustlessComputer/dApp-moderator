package usecase

import (
	"dapp-moderator/external/bfs_service"
	"dapp-moderator/external/block_stream"
	"dapp-moderator/external/blockchain_api"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/external/moralis"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/internal/generative_respository"
	"dapp-moderator/internal/repository"
	"dapp-moderator/utils/blockchain"
	"dapp-moderator/utils/config"
	discordclient "dapp-moderator/utils/discord"
	"dapp-moderator/utils/global"
	"dapp-moderator/utils/googlecloud"
	"dapp-moderator/utils/oauth2service"
	"dapp-moderator/utils/redis"
)

type Usecase struct {
	Repo           *repository.Repository
	GenerativeRepo *generative_respository.GenerativeRepository
	Config         *config.Config
	QuickNode      *quicknode.QuickNode
	BlockStream    *block_stream.BlockStream
	NftExplorer    *nft_explorer.NftExplorer
	TokenExplorer  *token_explorer.TokenExplorer
	BfsService     *bfs_service.BfsService
	BnsService     *bns_service.BNSService
	Cache          redis.IRedisCache
	Auth2          *oauth2service.Auth2
	Storage        googlecloud.IGcstorage
	DiscordClient  *discordclient.Client
	BlockChainApi  *blockchain_api.BlockChainApi
	Moralis        *moralis.MoralisService
	TCPublicNode   *blockchain.TcNetwork
}

func NewUsecase(global *global.Global, r *repository.Repository, generativeRepository *generative_respository.GenerativeRepository) (*Usecase, error) {
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
	u.Config = global.Conf
	u.BlockChainApi = global.BlockChainApi
	u.Moralis = global.Moralis
	u.GenerativeRepo = generativeRepository

	tcPublicNode, err := blockchain.NewTcNetwork()
	if err != nil {
		return nil, err
	}

	u.TCPublicNode = tcPublicNode
	return u, nil
}

func (u *Usecase) Version() string {
	return "dAPP-API Server - version 1"
}
