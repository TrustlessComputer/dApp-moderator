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

func (u Usecase) CollectionDetail(ctx context.Context, collectionAddress string) (interface{}, error) {
	data, err := u.NftExplorer.CollectionDetail(collectionAddress)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("collectionAddress", collectionAddress), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("CollectionDetail",zap.String("collectionAddress", collectionAddress), zap.Any("data", data))
	return data, nil
}

func (u Usecase) CollectionNfts(ctx context.Context, collectionAddress string) (interface{}, error) {
	data, err := u.NftExplorer.CollectionNfts(collectionAddress)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts",zap.String("collectionAddress", collectionAddress), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("CollectionNfts",zap.String("collectionAddress", collectionAddress), zap.Any("data", data))
	return data, nil
}

func (u Usecase) CollectionNftDetail(ctx context.Context, collectionAddress string, tokenID string) (interface{}, error) {
	data, err := u.NftExplorer.CollectionNftDetail(collectionAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts",zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, err
	}
	
	logger.AtLog.Logger.Info("CollectionNfts",zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	return data, nil
}

func (u Usecase) CollectionNftContent(ctx context.Context,collectionAddress string, tokenID string) ([]byte, string, error) {
	
	data, contentType, err := u.NftExplorer.CollectionNftContent(collectionAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNftContent",zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, "", err
	}
	
	logger.AtLog.Logger.Info("CollectionNftContent",zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
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
