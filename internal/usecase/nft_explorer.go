package usecase

import (
	"context"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)


func (u Usecase) Collections(ctx context.Context) (interface{}, error) {
	data, err := u.NftExplorer.Collections()
	if err != nil {
		logger.AtLog.Logger.Error("Collections", zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("Collections", zap.Any("data", data))
	return data, nil
}

func (u Usecase) CollectionDetail(ctx context.Context, contractAddress string) (interface{}, error) {
	data, err := u.NftExplorer.CollectionDetail(contractAddress)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("CollectionDetail",zap.String("contractAddress", contractAddress), zap.Any("data", data))
	return data, nil
}

func (u Usecase) CollectionNfts(ctx context.Context, contractAddress string) (interface{}, error) {
	data, err := u.NftExplorer.CollectionNfts(contractAddress)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts",zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("CollectionNfts",zap.String("contractAddress", contractAddress), zap.Any("data", data))
	return data, nil
}

func (u Usecase) CollectionNftDetail(ctx context.Context, contractAddress string, tokenID string) (interface{}, error) {
	data, err := u.NftExplorer.CollectionNftDetail(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts",zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("CollectionNfts",zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	return data, nil
}

func (u Usecase) CollectionNftContent(ctx context.Context,contractAddress string, tokenID string) ([]byte, string, error) {
	
	data, contentType, err := u.NftExplorer.CollectionNftContent(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNftContent",zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, "", err
	}
	
	logger.AtLog.Logger.Info("CollectionNftContent",zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	return data, contentType, nil
}

func (u Usecase) Nfts(ctx context.Context) (interface{}, error) {
	data, err := u.NftExplorer.Nfts()
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("Nfts", zap.Any("data", data))
	return data, nil
}

func (u Usecase) NftByWalletAddress(ctx context.Context, walletAddress string) (interface{}, error) {
	data, err := u.NftExplorer.NftOfWalletAddress(walletAddress)
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.String("walletAddress", walletAddress), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("Nfts", zap.String("walletAddress", walletAddress), zap.Any("data", data))
	return data, nil
}
