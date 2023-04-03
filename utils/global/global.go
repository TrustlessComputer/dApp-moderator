package global

import (
	"dapp-moderator/utils/config"
	_pConnection "dapp-moderator/utils/connections"
	"dapp-moderator/utils/googlecloud"
	_logger "dapp-moderator/utils/logger"

	"github.com/gorilla/mux"
)

type Global struct {
	Conf                *config.Config
	Logger              _logger.Ilogger
	MuxRouter           *mux.Router
	DBConnection        _pConnection.IConnection
	GCS                 googlecloud.IGcstorage
	S3Adapter           googlecloud.S3Adapter
}
