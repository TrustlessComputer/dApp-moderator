package usecase

import (
	"context"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)

func (u Usecase) AddressBalance(ctx context.Context, walletAddress string) (interface{}, error) {
	data, err := u.QuickNode.AddressBalance(walletAddress)
	if err != nil {
		logger.AtLog.Logger.Error("AddressBalance", zap.String("walletAddress", walletAddress), zap.Error(err))
	}

	logger.AtLog.Logger.Info("AddressBalance", zap.String("walletAddress", walletAddress), zap.Any("data", data))
	return data, err
}
