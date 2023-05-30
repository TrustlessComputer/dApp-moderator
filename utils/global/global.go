package global

import (
	"dapp-moderator/external/bfs_service"
	"dapp-moderator/external/block_stream"
	"dapp-moderator/external/blockchain_api"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/external/moralis"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/utils/config"
	_pConnection "dapp-moderator/utils/connections"
	discordclient "dapp-moderator/utils/discord"
	"dapp-moderator/utils/googlecloud"
	"dapp-moderator/utils/oauth2service"
	"dapp-moderator/utils/redis"

	"github.com/gorilla/mux"
)

type Global struct {
	Conf                   *config.Config
	MuxRouter              *mux.Router
	DBConnection           _pConnection.IConnection
	GenerativeDBConnection _pConnection.IConnection
	GCS                    googlecloud.IGcstorage
	S3Adapter              *googlecloud.S3Adapter
	Cache                  redis.IRedisCache
	CacheAuthService       redis.IRedisCache
	QuickNode              *quicknode.QuickNode
	BlockStream            *block_stream.BlockStream
	NftExplorer            *nft_explorer.NftExplorer
	BfsService             *bfs_service.BfsService
	BnsService             *bns_service.BNSService
	TokenExplorer          *token_explorer.TokenExplorer
	Auth2                  *oauth2service.Auth2
	DiscordClient          *discordclient.Client
	BlockChainApi          *blockchain_api.BlockChainApi
	Moralis                *moralis.MoralisService
}
