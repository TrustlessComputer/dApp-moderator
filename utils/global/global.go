package global

import (
	"dapp-moderator/external/nfts"
	"dapp-moderator/external/ord_service"
	"dapp-moderator/utils/blockchain"
	"dapp-moderator/utils/btc"
	"dapp-moderator/utils/config"
	_pConnection "dapp-moderator/utils/connections"
	"dapp-moderator/utils/delegate"
	discordclient "dapp-moderator/utils/discord"
	"dapp-moderator/utils/eth"
	"dapp-moderator/utils/googlecloud"
	_logger "dapp-moderator/utils/logger"
	"dapp-moderator/utils/oauth2service"
	"dapp-moderator/utils/redis"
	"dapp-moderator/utils/redisv9"
	"dapp-moderator/utils/slack"

	"github.com/gorilla/mux"
)

type Global struct {
	Conf                *config.Config
	Logger              _logger.Ilogger
	MuxRouter           *mux.Router
	DBConnection        _pConnection.IConnection
	Cache               redis.IRedisCache
	CacheAuthService    redis.IRedisCache
	RedisV9             redisv9.Client
	Pubsub              redis.IPubSubClient
	Auth2               oauth2service.Auth2
	GCS                 googlecloud.IGcstorage
	S3Adapter           googlecloud.S3Adapter
	MoralisNFT          nfts.MoralisNfts
	CovalentNFT         nfts.CovalentNfts
	OrdService          *ord_service.BtcOrd
	OrdServiceDeveloper *ord_service.BtcOrd
	Blockchain          blockchain.Blockchain
	TcNetwotkchain      blockchain.TcNetwork
	Slack               slack.Slack
	DiscordClient       *discordclient.Client
	DelegateService     *delegate.Service

	TcClient, EthClient *eth.Client
	BsClient            *btc.BlockcypherService
}
