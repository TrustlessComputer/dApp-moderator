package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils/logger"
	"go.uber.org/zap"
)

func (c *Usecase) Tokens(ctx context.Context, filter request.PaginationReq) (interface{}, error) {

	data, err := c.TokenExplorer.Tokens(filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Tokens", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
	return data, nil
}
