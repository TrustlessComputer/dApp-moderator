package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"net/url"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (c *Usecase) Collections(ctx context.Context, filter request.PaginationReq) ([]entity.Nfts, error) {
	res := []entity.Nfts{}
	err := c.Repo.Find(utils.COLLECTION_NFTS, bson.D{}, int64(*filter.Limit), int64(*filter.Offset), &res)
	if err != nil {
		logger.AtLog.Logger.Error("Collections", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Collections", zap.Any("data", len(res)))
	return res, nil
}

func (c *Usecase) CollectionDetail(ctx context.Context, contractAddress string) (*entity.Nfts, error) {
	obj := &entity.Nfts{}
 	sr, err := c.Repo.FindOne(utils.COLLECTION_NFTS, bson.D{
		{"contract", primitive.Regex{Pattern: contractAddress, Options: "i"}},
	})
	
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.Error(err))
		return nil, err
	}

	err = sr.Decode(obj)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Any("obj", obj))
	return obj, nil
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

func (c *Usecase) GetCollectionFromBlock(ctx context.Context, fromBlock int32, toBlock int32) (interface{}, error) {

	params := url.Values{}
	params.Set("filter", fmt.Sprintf(`{"deployed_at_block":{"$gte":%d,"$lte":%d}}`, fromBlock, toBlock) ) 

	data, err := c.NftExplorer.Collections(params)
	if err != nil {
		logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock) ,  zap.Error(err))
		return nil, err
	}

	if len(data) == 0 {
		return data, nil
	} 

	for _, item := range data {
		tmp := &entity.Nfts{}
		err := helpers.JsonTransform(item, tmp)
		if err != nil {
			logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("contract", item.Contract) , zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock) ,  zap.Error(err))
			continue
		}

		tmp.Slug = helpers.GenerateSlug(tmp.Name)
		tmp.Contract = strings.ToLower(tmp.Contract)
		tmp.Creator = strings.ToLower(tmp.Creator)
		
		inserted, err := c.Repo.InsertOne(tmp)
		if err != nil {
			logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("contract", item.Contract) , zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock) ,  zap.Error(err))
			continue
		}

		_ = inserted
	}

	logger.AtLog.Logger.Info("GetCollectionFromBlock", zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock) , zap.Any("data", data))
	return data, nil
}