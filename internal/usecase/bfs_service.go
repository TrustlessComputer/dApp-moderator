package usecase

import (
	"context"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)

func (c *Usecase) BfsFiles(ctx context.Context, walletAddress string) (interface{}, error) {
	data, err := c.BfsService.Files(walletAddress)
	if err != nil {
		logger.AtLog.Logger.Error("BfsFiles", zap.String("walletAddress", walletAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("BfsFiles", zap.String("walletAddress", walletAddress), zap.Any("data", data))
	return data, nil
}

func (c *Usecase) BfsBrowsedFile(ctx context.Context, walletAddress string, path string) (interface{}, error) {
	data, err := c.BfsService.BrowseFiles(walletAddress, path)
	if err != nil {
		logger.AtLog.Logger.Error("BrowseFiles", zap.String("walletAddress", walletAddress), zap.String("path", path), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("BrowseFiles", zap.String("walletAddress", walletAddress), zap.String("path", path), zap.Any("data", data))
	return data, nil
}

func (c *Usecase) BfsFileInfo(ctx context.Context, walletAddress string, path string) (interface{}, error) {
	data, err := c.BfsService.FileInfo(walletAddress, path)
	if err != nil {
		logger.AtLog.Logger.Error("FileInfo", zap.String("walletAddress", walletAddress), zap.String("path", path), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("FileInfo", zap.String("walletAddress", walletAddress), zap.String("path", path), zap.Any("data", data))
	return data, nil
}

func (c *Usecase) BfsFileContent(ctx context.Context, walletAddress string, path string) ([]byte, string, error) {
	data, contentType, err := c.BfsService.FileContent(walletAddress, path)
	if err != nil {
		logger.AtLog.Logger.Error("FileInfo", zap.String("BfsFileContent", walletAddress), zap.String("path", path), zap.Error(err))
		return nil, contentType, err
	}

	logger.AtLog.Logger.Info("FileInfo", zap.String("BfsFileContent", walletAddress), zap.String("path", path), zap.Any("data", data))
	return data, contentType, nil
}
