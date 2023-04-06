package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)

func (c *Usecase) Collections(ctx context.Context, filter request.PaginationReq) (interface{}, error) {

	data, err := c.NftExplorer.Collections(filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Collections", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Collections", zap.Any("data", data))
	return data, nil
}

func (c *Usecase) CollectionDetail(ctx context.Context, contractAddress string) (interface{}, error) {
	data, err := c.NftExplorer.CollectionDetail(contractAddress)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Any("data", data))
	return data, nil
}

func (c *Usecase) CollectionNfts(ctx context.Context, contractAddress string, filter request.PaginationReq) (interface{}, error) {
	data, err := c.NftExplorer.CollectionNfts(contractAddress, filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Any("data", data))
	return data, nil
}

func (c *Usecase) CollectionNftDetail(ctx context.Context, contractAddress string, tokenID string) (interface{}, error) {
	data, err := c.NftExplorer.CollectionNftDetail(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionNfts", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	return data, nil
}

func (c *Usecase) CollectionNftContent(ctx context.Context, contractAddress string, tokenID string) ([]byte, string, error) {

	data, contentType, err := c.NftExplorer.CollectionNftContent(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, "", err
	}

	logger.AtLog.Logger.Info("CollectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	return data, contentType, nil
}

func (c *Usecase) Nfts(ctx context.Context, filter request.PaginationReq) (interface{}, error) {
	data, err := c.NftExplorer.Nfts(filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Nfts", zap.Any("data", data))
	return data, nil
}

func (c *Usecase) NftByWalletAddress(ctx context.Context, walletAddress string, filter request.PaginationReq) (interface{}, error) {
	data, err := c.NftExplorer.NftOfWalletAddress(walletAddress, filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.String("walletAddress", walletAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Nfts", zap.String("walletAddress", walletAddress), zap.Any("data", data))
	return data, nil
}
