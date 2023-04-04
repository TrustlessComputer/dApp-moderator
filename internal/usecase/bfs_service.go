package usecase

import (
	"context"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)


func (u Usecase) BfsFiles(ctx context.Context, walletAddress string) (interface{}, error) {
	data, err := u.BfsService.Files(walletAddress)
	if err != nil {
		logger.AtLog.Logger.Error("BfsFiles", zap.String("walletAddress",walletAddress), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("BfsFiles",zap.String("walletAddress",walletAddress), zap.Any("data", data))
	return data, nil
}

func (u Usecase) BfsBrowsedFile(ctx context.Context, walletAddress string, path string) (interface{}, error) {
	data, err := u.BfsService.BrowseFiles(walletAddress, path)
	if err != nil {
		logger.AtLog.Logger.Error("BrowseFiles", zap.String("walletAddress",walletAddress), zap.String("path",path), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("BrowseFiles",zap.String("walletAddress",walletAddress), zap.String("path",path), zap.Any("data", data))
	return data, nil
}

func (u Usecase) BfsFileInfo(ctx context.Context, walletAddress string, path string) (interface{}, error) {
	data, err := u.BfsService.FileInfo(walletAddress, path)
	if err != nil {
		logger.AtLog.Logger.Error("FileInfo", zap.String("walletAddress",walletAddress), zap.String("path",path), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("FileInfo",zap.String("walletAddress",walletAddress), zap.String("path",path), zap.Any("data", data))
	return data, nil
}
