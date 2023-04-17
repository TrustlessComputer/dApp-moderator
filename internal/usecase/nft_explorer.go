package usecase

import (
	"context"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

func (u *Usecase) Collections(ctx context.Context, filter request.CollectionsFilter) ([]entity.Collections, error) {
	res := []entity.Collections{}
	f := bson.D{
		// {"total_items", bson.M{"$gt": 0}},
	}

	if filter.AllowEmpty != nil && *filter.AllowEmpty == false {
		f = append(f, bson.E{"total_items", bson.M{"$gt": 0}})
	}

	if filter.Address != nil && *filter.Address != "" {
		f = append(f, bson.E{"contract", primitive.Regex{Pattern: *filter.Address, Options: "i"}})
	}

	if filter.Name != nil && *filter.Name != "" {
		f = append(f, bson.E{"name", primitive.Regex{Pattern: *filter.Name, Options: "i"}})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		f = append(f, bson.E{"creator", primitive.Regex{Pattern: *filter.Owner, Options: "i"}})
	}

	sortBy := "deployed_at_block"
	if filter.SortBy != nil && *filter.SortBy != "" {
		sortBy = *filter.SortBy
	}

	sort := -1
	if filter.Sort != nil {
		sort = *filter.Sort
	}

	s := bson.D{{sortBy, sort}, {"index", 1}}
	err := u.Repo.Find(utils.COLLECTION_COLLECTIONS, f, int64(*filter.Limit), int64(*filter.Offset), &res, s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Usecase) CollectionNftsFrom3rdService(ctx context.Context, contractAddress string, filter request.PaginationReq) ([]*nft_explorer.NftsResp, error) {
	data, err := u.NftExplorer.CollectionNfts(contractAddress, filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Any("filter", filter), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Any("filter", filter), zap.Any("data", len(data)))
	return data, nil
}

func (u *Usecase) CollectionsWithoutLogic(ctx context.Context, filter request.PaginationReq) ([]entity.Collections, error) {
	res := []entity.Collections{}
	f := bson.D{}

	sort := bson.D{{"deployed_at_block", 1}}

	err := u.Repo.Find(utils.COLLECTION_COLLECTIONS, f, int64(*filter.Limit), int64(*filter.Offset), &res, sort)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Usecase) CollectionDetail(ctx context.Context, contractAddress string) (*entity.Collections, error) {
	obj := &entity.Collections{}
	sr, err := u.Repo.FindOne(utils.COLLECTION_COLLECTIONS, bson.D{
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

func (u *Usecase) UpdateCollection(ctx context.Context, contractAddress string, walletAdress string, updateData *structure.UpdateCollection) (*entity.Collections, error) {
	obj := &entity.Collections{}

	f := bson.D{
		{"contract", primitive.Regex{Pattern: contractAddress, Options: "i"}},
		{"creator", primitive.Regex{Pattern: walletAdress, Options: "i"}},
	}
	sr, err := u.Repo.FindOne(utils.COLLECTION_COLLECTIONS, f)

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

	_, err = u.Repo.ReplaceOne(f, obj)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionDetail", zap.String("contractAddress", contractAddress), zap.Any("obj", obj))
	return obj, nil
}

func (u *Usecase) CollectionNfts(ctx context.Context, contractAddress string, filter request.CollectionsFilter) ([]entity.Nfts, error) {
	// data, err := u.NftExplorer.CollectionNfts(contractAddress, filter.ToNFTServiceUrlQuery())
	// if err != nil {
	// 	logger.AtLog.Logger.Error("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Any("filter", filter), zap.Error(err))
	// 	return nil, err
	// }

	// logger.AtLog.Logger.Info("CollectionNfts", zap.String("contractAddress", contractAddress), zap.Any("filter", filter), zap.Any("data", len(data)))
	// return data, nil

	res := []entity.Nfts{}
	f := bson.D{}

	if filter.Address != nil && *filter.Address != "" {
		f = append(f, bson.E{"collection_address", primitive.Regex{Pattern: *filter.Address, Options: "i"}})
	}

	if filter.Name != nil && *filter.Name != "" {
		f = append(f, bson.E{"collection", primitive.Regex{Pattern: *filter.Name, Options: "i"}})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		f = append(f, bson.E{"owner", primitive.Regex{Pattern: *filter.Owner, Options: "i"}})
	}

	sortBy := "token_id_int"
	if filter.SortBy != nil && *filter.SortBy != "" {
		sortBy = *filter.SortBy
	}

	sort := -1
	if filter.Sort != nil {
		sort = *filter.Sort
	}

	s := bson.D{{sortBy, sort}}
	err := u.Repo.Find(utils.COLLECTION_NFTS, f, int64(*filter.Limit), int64(*filter.Offset), &res, s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Usecase) CollectionNftDetail(ctx context.Context, contractAddress string, tokenID string) (*nft_explorer.NftsResp, error) {
	data, err := u.NftExplorer.CollectionNftDetail(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNfts", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("CollectionNfts", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	return data, nil
}

func (u *Usecase) CollectionNftContent(ctx context.Context, contractAddress string, tokenID string) ([]byte, string, error) {

	data, contentType, err := u.NftExplorer.CollectionNftContent(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("CollectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, "", err
	}

	logger.AtLog.Logger.Info("CollectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", len(data)))
	return data, contentType, nil
}

func (u *Usecase) Nfts(ctx context.Context, filter request.PaginationReq) ([]*nft_explorer.NftsResp, error) {
	data, err := u.NftExplorer.Nfts(filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Nfts", zap.Any("data", data))
	return data, nil
}

func (u *Usecase) NftByWalletAddress(ctx context.Context, walletAddress string, filter request.PaginationReq) ([]*nft_explorer.NftsResp, error) {
	data, err := u.NftExplorer.NftOfWalletAddress(walletAddress, filter.ToNFTServiceUrlQuery())
	if err != nil {
		logger.AtLog.Logger.Error("Nfts", zap.String("walletAddress", walletAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Nfts", zap.String("walletAddress", walletAddress), zap.Any("data", len(data)))
	return data, nil
}

func (u *Usecase) GetCollectionFromBlock(ctx context.Context, fromBlock int32, toBlock int32) error {
	params := url.Values{}
	page := 1
	limit := 100
	for {

		offset := limit * (page - 1)
		params.Set("filter", fmt.Sprintf(`{"deployed_at_block":{"$gte":%d,"$lte":%d}}`, fromBlock, toBlock))
		params.Set("limit", fmt.Sprintf("%d", limit))
		params.Set("offset", fmt.Sprintf("%d", offset))

		data, err := u.NftExplorer.Collections(params)
		if err != nil {
			logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("params", params), zap.Error(err))
			break
		}

		if len(data) == 0 {
			break
		}

		countInt := int64(0)
		count, _, err := u.Repo.CountDocuments(utils.COLLECTION_COLLECTIONS, bson.D{})
		if err != nil || count == nil {
			countInt = 0
		} else {
			countInt = *count
		}
		countInt++

		//revert the array to index
		for i := len(data) - 1; i >= 0; i = i - 1 {
			item := data[i]

			tmp := &entity.Collections{}
			err := helpers.JsonTransform(item, tmp)
			if err != nil {
				logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("contract", item.Contract), zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Error(err))
				continue
			}

			_, err = u.CollectionDetail(ctx, item.Contract)
			if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
				tmp.Index = countInt
				tmp.Slug = helpers.GenerateSlug(tmp.Name)
				tmp.Contract = strings.ToLower(tmp.Contract)
				tmp.Creator = strings.ToLower(tmp.Creator)

				_, err := u.Repo.InsertOne(tmp)
				if err != nil {
					logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("contract", item.Contract), zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Error(err))
					continue
				}
				u.NewCollectionNotify(tmp)
			}

			// else {
			// 	updatedData := bson.M{
			// 		"$set": bson.M{"index": countInt},
			// 	}
			// 	_, err := u.Repo.UpdateOne(utils.COLLECTION_COLLECTIONS, bson.D{{"contract", nft.Contract}}, updatedData)
			// 	if err != nil {
			// 		logger.AtLog.Logger.Error("GetCollectionFromBlock", zap.Any("contract", item.Contract), zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Error(err))
			// 		continue
			// 	}
			// }

			countInt++
		}

		logger.AtLog.Logger.Info("GetCollectionFromBlock", zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock), zap.Any("data", len(data)))

		page++
	}

	return nil
}

func (u *Usecase) UpdateCollectionItems(ctx context.Context) error {
	filter := request.PaginationReq{}
	page := 1
	limit := 10

	for {

		//filter again
		offset := limit * (page - 1)
		filter.Page = &page
		filter.Limit = &limit
		filter.Offset = &offset
		nfts, err := u.CollectionsWithoutLogic(ctx, filter)
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
			go u.GetNftsFromCollection(ctx, &wg, contract, nft)

		}
		wg.Wait()
		page++
	}

	return nil
}

func (u *Usecase) GetNftsFromCollection(ctx context.Context, wg *sync.WaitGroup, contract string, nft entity.Collections) {
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
			defer func() {
				channelItems <- tmpItems
			}()

			tmpItems, err := u.CollectionNftsFrom3rdService(ctx, contract, request.PaginationReq{
				Limit:  &itemsLimit,
				Offset: &offset,
			})

			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("UpdateCollection.%s", contract), zap.String("contract", contract), zap.Error(err))
				return
			}

		}(ctx, page, itemsLimit, channelItems)

		tmpItems := <-channelItems
		if len(tmpItems) == 0 {
			break
		}

		for _, tmpItem := range tmpItems {
			items = append(items, tmpItem)
			err := u.InsertOrUpdateNft(tmpItem)
			if err == nil {
				total += len(tmpItems)
			}
		}

		page++
	}

	totalItems := len(items)
	if totalItems == 0 {
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

	updated, err := u.Repo.UpdateOne(nft.CollectionName(), f, updateData)
	if err != nil {
		return
	}

	logger.AtLog.Logger.Info(fmt.Sprintf("UpdateCollection.%s", contract), zap.String("contract", contract), zap.Int("items", totalItems), zap.Any("updated", updated))
}

func (u *Usecase) UserCollections(ctx context.Context, filter request.CollectionsFilter) ([]entity.Collections, error) {
	return u.Repo.UserCollections(filter)
}

func (u *Usecase) InsertOrUpdateNft(item *nft_explorer.NftsResp) error {
	tmp := &entity.Nfts{}

	err := helpers.JsonTransform(item, tmp)
	if err != nil {
		return err
	}

	contract := item.ContractAddress
	tokenIDInt, err := strconv.Atoi(tmp.TokenID)
	if err == nil {
		tmp.TokenIDInt = int64(tokenIDInt)
	}

	nft, err := u.Repo.GetNft(tmp.ContractAddress, tmp.TokenID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, err = u.Repo.CreateNftHistories(&entity.NftHistories{
				Collection:        strings.ToLower(tmp.Collection),
				ContractAddress:   strings.ToLower(tmp.ContractAddress),
				TokenID:           tmp.TokenID,
				TokenIDInt:        tmp.TokenIDInt,
				FromWalletAddress: strings.ToLower(tmp.Owner),
				ToWalletAddress:   strings.ToLower(tmp.Owner),
				Action:            "mint",
			})
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("UpdateCollection.%s.CreateNftHistories", contract), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
				return err
			}

			_, err = u.Repo.InsertOne(tmp)
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("UpdateCollection.%s.InsertOne", contract), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
				return err
			}
			if tmp.Collection == strings.ToLower("0x16EfDc6D3F977E39DAc0Eb0E123FefFeD4320Bc0") {
				u.NewArtifactNotify(tmp)
			}
			if tmp.Collection == strings.ToLower("0x8b46F89BBA2B1c1f9eE196F43939476E79579798") {
				u.NewNameNotify(&bns_service.NameResp{
					Owner: tmp.Owner,
					Name:  tmp.Name,
					ID:    tmp.TokenID,
				})
			}

		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateCollection.%s", contract), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
			return err
		}
	} else {
		//the current owner != owner from chain
		if strings.ToLower(nft.Owner) != strings.ToLower(tmp.Owner) {
			_, err := u.Repo.CreateNftHistories(&entity.NftHistories{
				Collection:        strings.ToLower(tmp.Collection),
				ContractAddress:   strings.ToLower(tmp.ContractAddress),
				TokenID:           tmp.TokenID,
				TokenIDInt:        tmp.TokenIDInt,
				FromWalletAddress: strings.ToLower(nft.Owner),
				ToWalletAddress:   strings.ToLower(tmp.Owner),
				Action:            "transfer",
			})

			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("UpdateCollection.%s.%s.history", tmp.ContractAddress, tmp.TokenID), zap.String("owner", tmp.Owner), zap.Error(err))
			}

			_, err = u.Repo.UpdateNftOwner(tmp.ContractAddress, tmp.TokenID, tmp.Owner)
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("UpdateCollection.%s.%s.onwer", tmp.ContractAddress, tmp.TokenID), zap.String("owner", tmp.Owner), zap.Error(err))
			}

		}
	}

	return nil
}
