package usecase

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (c *Usecase) Collections(ctx context.Context, filter request.CollectionsFilter) ([]entity.Nfts, error) {
	res := []entity.Nfts{}
	f := bson.D{
		// {"total_items", bson.M{"$gt": 0}},
	}

	if filter.AllowEmpty != nil && *filter.AllowEmpty == false {
		f = append(f, bson.E{"total_items", bson.M{"$gt": 0}})
	}

	if filter.Address != nil {
		f = append(f, bson.E{"contract", primitive.Regex{Pattern: *filter.Address, Options: "i"}})
	}

	if filter.Name != nil {
		f = append(f, bson.E{"name", primitive.Regex{Pattern: *filter.Name, Options: "i"}})
	}

	if filter.Owner != nil {
		f = append(f, bson.E{"creator", primitive.Regex{Pattern: *filter.Owner, Options: "i"}})
	}

	sortBy := "deployed_at_block"
	if filter.SortBy != nil && *filter.SortBy != "" {
		sortBy = *filter.SortBy
	}

	sort := 1
	if filter.Sort != nil {
		sort = *filter.Sort
	}

	s := bson.D{{sortBy, sort}}
	err := c.Repo.Find(utils.COLLECTION_NFTS, f, int64(*filter.Limit), int64(*filter.Offset), &res, s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Usecase) CollectionsWithoutLogic(ctx context.Context, filter request.PaginationReq) ([]entity.Nfts, error) {
	res := []entity.Nfts{}
	f := bson.D{}

	sort := bson.D{{"deployed_at_block", 1}}

	err := c.Repo.Find(utils.COLLECTION_NFTS, f, int64(*filter.Limit), int64(*filter.Offset), &res, sort)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Usecase) CollectionDetail(ctx context.Context, contractAddress string) (*entity.Nfts, error) {
	obj := &entity.Nfts{}
	sr, err := c.Repo.FindOne(utils.COLLECTION_NFTS, bson.D{
		{"contract", primitive.Regex{Pattern: contractAddress, Options: "i"}},
	})

	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	err = sr.Decode(obj)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Any("obj", obj))
	return obj, nil
}


func (c *Usecase) UpdateCollection(ctx context.Context, contractAddress string, walletAdress string, updateData *structure.UpdateCollection) (*entity.Nfts, error) {
	obj := &entity.Nfts{}

	f := bson.D{
		{"contract", primitive.Regex{Pattern: contractAddress, Options: "i"}},
		{"creator", primitive.Regex{Pattern: walletAdress, Options: "i"}},
	}
	sr, err := c.Repo.FindOne(utils.COLLECTION_NFTS, f)

	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	err = sr.Decode(obj)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	if updateData.Cover != nil && *updateData.Cover != obj.Cover {
		obj.Cover = *updateData.Cover
	}
	
	if updateData.Thumbnail != nil && *updateData.Thumbnail != obj.Thumbnail {
		obj.Thumbnail = *updateData.Thumbnail
	}
	
	if updateData.Description != nil && *updateData.Description != obj.Description {
		obj.Description = *updateData.Description
	}
	
	if updateData.Social.DisCord != nil && *updateData.Social.DisCord != obj.Social.DisCord {
		obj.Social.DisCord = *updateData.Social.DisCord
	}
	
	if updateData.Social.Instagram != nil && *updateData.Social.Instagram != obj.Social.Instagram {
		obj.Social.Instagram = *updateData.Social.Instagram
	}
	
	if updateData.Social.Medium != nil && *updateData.Social.Medium != obj.Social.Medium {
		obj.Social.Medium = *updateData.Social.Medium
	}

	if updateData.Social.Telegram != nil && *updateData.Social.Telegram != obj.Social.Telegram {
		obj.Social.Telegram = *updateData.Social.Telegram
	}
	
	if updateData.Social.Twitter != nil && *updateData.Social.Twitter != obj.Social.Twitter {
		obj.Social.Twitter = *updateData.Social.Twitter
	}
	
	if updateData.Social.Website != nil && *updateData.Social.Website != obj.Social.Website {
		obj.Social.Website = *updateData.Social.Website
	}

	_, err = c.Repo.ReplaceOne(f, obj)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Any("obj", obj))
	return obj, nil
}

func (c *Usecase) CollectionNfts(ctx context.Context, contractAddress string, filter request.PaginationReq) ([]*nft_explorer.NftsResp, error) {
	data, err := c.NftExplorer.CollectionNfts(contractAddress, filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Any("filter", filter), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Any("filter", filter), zap.Any("data", len(data)))
	return data, nil
}

func (c *Usecase) CollectionNftDetail(ctx context.Context, contractAddress string, tokenID string) (*nft_explorer.NftsResp, error) {
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

	logger.AtLog.Logger.Info("CollectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", len(data)))
	return data, contentType, nil
}

func (c *Usecase) Nfts(ctx context.Context, filter request.PaginationReq) ([]*nft_explorer.NftsResp, error) {
	data, err := c.NftExplorer.Nfts(filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Nfts", zap.Any("data", data))
	return data, nil
}

func (c *Usecase) NftByWalletAddress(ctx context.Context, walletAddress string, filter request.PaginationReq) ([]*nft_explorer.NftsResp, error) {
	data, err := c.NftExplorer.NftOfWalletAddress(walletAddress, filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.String("walletAddress", walletAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Nfts", zap.String("walletAddress", walletAddress), zap.Any("data", len(data)))
	return data, nil
}

func (c *Usecase) GetCollectionFromBlock(ctx context.Context, fromBlock int32, toBlock int32) (interface{}, error) {

	params := url.Values{}
	params.Set("filter", fmt.Sprintf(`{"deployed_at_block":{"$gte":%d,"$lte":%d}}`, fromBlock, toBlock))

	data, err := c.NftExplorer.Collections(params)
	if err != nil {
		logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Error(err))
		return nil, err
	}

	if len(data) == 0 {
		return data, nil
	}

	for _, item := range data {
		tmp := &entity.Nfts{}
		err := helpers.JsonTransform(item, tmp)
		if err != nil {
			logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("contract", item.Contract), zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Error(err))
			continue
		}

		tmp.Slug = helpers.GenerateSlug(tmp.Name)
		tmp.Contract = strings.ToLower(tmp.Contract)
		tmp.Creator = strings.ToLower(tmp.Creator)

		inserted, err := c.Repo.InsertOne(tmp)
		if err != nil {
			logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("contract", item.Contract), zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Error(err))
			continue
		}

		_ = inserted
	}

	logger.AtLog.Logger.Info("GetCollectionFromBlock", zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Any("data", data))
	return data, nil
}

func (c *Usecase) UpdateCollectionItems(ctx context.Context) error {
	filter := request.PaginationReq{}
	page := 1
	limit := 10
	for {

		//filter again
		offset := limit * (page - 1)
		filter.Page = &page
		filter.Limit = &limit
		filter.Offset = &offset
		nfts, err := c.CollectionsWithoutLogic(ctx, filter)
		if err != nil {
			break
		}

		if len(nfts) == 0 {
			break
		}

		var wg sync.WaitGroup
		for _, nft := range nfts {
			contract := strings.ToLower(nft.Contract)

			wg.Add(1)
			go func(wg *sync.WaitGroup, nft entity.Nfts) {
				defer wg.Done()

				items := []*nft_explorer.NftsResp{}
				itemsLimit := 100
				page := 1
				total := 0

				channelItems := make(chan []*nft_explorer.NftsResp)
				for {

					go func(ctx context.Context, page int, itemsLimit int, channelItems chan []*nft_explorer.NftsResp) {

						offset := itemsLimit * (page - 1)

						tmpItems := []*nft_explorer.NftsResp{}
						defer func  ()  {
							channelItems <- tmpItems
						}()

						//TODO - Paging the request data
						tmpItems, err = c.CollectionNfts(ctx, contract, request.PaginationReq{
							Limit:  &itemsLimit,
							Offset: &offset,
						})

						if err != nil {
							logger.AtLog.Logger.Error(fmt.Sprintf("UpdateCollection.%s", contract), zap.String("contract", contract), zap.Error(err))
							return
						}


					}(ctx, page, itemsLimit, channelItems)


					tmpItems := <- channelItems
					if len(tmpItems) == 0 {
						break
					}

					for _, tmpItem := range tmpItems {
						items = append(items, tmpItem)
					}


					total += len(tmpItems)
					page++
				}

				totalItems := len(items)
				if totalItems == 0 {
					return
				}

				if totalItems == nft.TotalItems {
					return
				}

				f := bson.D{
					{"contract", contract},
				}

				updateData := bson.M{
					"$set": bson.M{
						"total_items": totalItems,
					},
				}

				updated, err := c.Repo.UpdateOne(nft.CollectionName(), f, updateData)
				if err != nil {
					return
				}

				logger.AtLog.Logger.Info(fmt.Sprintf("UpdateCollection.%s", contract), zap.String("contract", contract), zap.Int("items", totalItems), zap.Any("updated", updated))
			}(&wg, nft)

			wg.Wait()

		}

		page++
	}

	return nil
}
