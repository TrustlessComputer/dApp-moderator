package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils/logger"
	"go.uber.org/zap"
)

func (c *Usecase) Tokens(ctx context.Context, filter request.PaginationReq, key string) (interface{}, error) {
	var data interface{}
	var err error

	params := filter.ToNFTServiceUrlQuery()

	if key == "" {
		data, err = c.TokenExplorer.Tokens(params)
	} else {
		params["query"] = []string{key}
		data, err = c.TokenExplorer.Search(params)
	}

	if err != nil {
		logger.AtLog.Logger.Error("Tokens", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
	return data, nil
}

func (c *Usecase) Token(ctx context.Context, address string) (interface{}, error) {

	data, err := c.TokenExplorer.Token(address)
	if err != nil {
		logger.AtLog.Logger.Error("Token", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Token", zap.Any("data", data))
	return data, nil
}
