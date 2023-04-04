package global

import (
	"dapp-moderator/external/bfs_service"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/utils/config"
	_pConnection "dapp-moderator/utils/connections"
	"dapp-moderator/utils/googlecloud"
	_logger "dapp-moderator/utils/logger"
	"dapp-moderator/utils/redis"

	"github.com/gorilla/mux"
)

type Global struct {
	Conf                *config.Config
	Logger              _logger.Ilogger
	MuxRouter           *mux.Router
	DBConnection        _pConnection.IConnection
	GCS                 googlecloud.IGcstorage
	S3Adapter           googlecloud.S3Adapter
	Cache               redis.IRedisCache
	CacheAuthService    redis.IRedisCache
	QuickNode			*quicknode.QuickNode
	NftExplorer			*nft_explorer.NftExplorer
	BfsService			*bfs_service.BfsService
}
