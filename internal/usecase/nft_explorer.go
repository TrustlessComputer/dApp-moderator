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
	"errors"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
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

	f = append(f,
		bson.E{
			"$or",
			bson.A{
				bson.D{{"status", 0}},
				bson.D{{"status", primitive.Null{}}},
			},
		})

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

	if updateData.Name != nil && *updateData.Name != obj.Name {
		obj.Name = *updateData.Name
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
			if strings.Contains(strings.ToLower(item.Name), strings.ToLower("LP Token")) {
				logger.AtLog.Logger.Error("GetCollectionFromBlock LP Token", zap.Any("contract", item.Contract), zap.Int32("fromBlock", fromBlock), zap.Int32("toBlock", toBlock))
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

		logger.AtLog.Logger.Info(fmt.Sprintf("UpdateCollectionItems.page.%d.limit.%d", page, limit), zap.Any("filter", filter))
		nfts, err := u.CollectionsWithoutLogic(ctx, filter)
		if err != nil {
			logger.AtLog.Logger.Info(fmt.Sprintf("UpdateCollectionItems.page.%d.limit.%d", page, limit), zap.Any("filter", filter), zap.Error(err))
			break
		}

		if len(nfts) == 0 {
			err = errors.New("nfts is empty")
			logger.AtLog.Logger.Info(fmt.Sprintf("UpdateCollectionItems.page.%d.limit.%d", page, limit), zap.Any("filter", filter), zap.Error(err))
			break
		}

		var wg sync.WaitGroup
		for _, nft := range nfts {
			contract := strings.ToLower(nft.Contract)
			//logger.AtLog.Logger.Info(fmt.Sprintf("UpdateCollectionItems.%s", contract), zap.String("contract", contract))

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
	key := helpers.NftsOfContractPageKey(contract)
	page := 1

	existed := false
	cachedPage, err := u.Cache.GetData(key)
	if err == nil && cachedPage != nil {
		pageInt, err := strconv.Atoi(*cachedPage)
		if err == nil {
			page = pageInt
			existed = true
		}
	}

	items := []*nft_explorer.NftsResp{}
	itemsLimit := 100

	total := 0

	if page > 1 {
		page-- //Move to the last page that has items
	}

	channelItems := make(chan []*nft_explorer.NftsResp)
	for {

		go func(ctx context.Context, colectionPage int, itemsLimit int, channelItems chan []*nft_explorer.NftsResp) {
			offset := itemsLimit * (colectionPage - 1)

			tmpItems := []*nft_explorer.NftsResp{}
			defer func() {
				channelItems <- tmpItems
				//logger.AtLog.Logger.Info(fmt.Sprintf("GetNftsFromCollection.Routine.%s", contract), zap.String("contract", contract), zap.Any("page", colectionPage), zap.Any("itemsLimit", itemsLimit), zap.Any("Offset", offset), zap.Any("tmpItems", len(tmpItems)))

			}()

			tmpItems, err := u.CollectionNftsFrom3rdService(ctx, contract, request.PaginationReq{
				Limit:  &itemsLimit,
				Offset: &offset,
			})

			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("GetNftsFromCollection.%s", contract), zap.String("contract", contract), zap.Error(err))
				return
			}

		}(ctx, page, itemsLimit, channelItems)

		tmpItems := <-channelItems

		if len(tmpItems) == 0 {
			break
		}

		var itemWg sync.WaitGroup
		items = append(items, tmpItems...)
		total += len(tmpItems)
		for _, tmpItem := range tmpItems {
			itemWg.Add(1)
			go func(itemWg *sync.WaitGroup, ctx context.Context, nft *nft_explorer.NftsResp) {
				defer itemWg.Done()
				u.InsertOrUpdateNft(ctx, nft)
			}(&itemWg, ctx, tmpItem)
		}
		itemWg.Wait()

		page++
		//Update current page here.
		u.Cache.SetStringData(key, fmt.Sprintf("%d", page))

	}

	totalItems := len(items)
	if totalItems == 0 {
		return
	}

	if existed {
		offset := (page - 2) * itemsLimit //get offset = (page-1)*limit ( - 1), and the cached page that was not updated by the loop ( - 1): total - 2
		totalItems = offset + totalItems  // items of the previous offset and the current page (totalItems)
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
		logger.AtLog.Logger.Error(fmt.Sprintf("GetNftsFromCollection.%s", contract), zap.String("contract", contract), zap.Error(err))
		return
	}

	logger.AtLog.Logger.Info(fmt.Sprintf("GetNftsFromCollection.%s", contract), zap.String("contract", contract), zap.Int("items", totalItems), zap.Any("updated", updated))
}

func (u *Usecase) UserCollections(ctx context.Context, filter request.CollectionsFilter) ([]entity.Collections, error) {
	return u.Repo.UserCollections(filter)
}

func (u *Usecase) InsertOrUpdateNft(ctx context.Context, item *nft_explorer.NftsResp) error {
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

	artfactAddress := strings.ToLower(os.Getenv("ARTIFACT_ADDRESS"))
	bnsAddress := strings.ToLower(os.Getenv("BNS_ADDRESS"))

	//logger.AtLog.Logger.Info(fmt.Sprintf("InsertOrUpdateNft.%s", contract), zap.String("contract", tmp.ContractAddress), zap.String("tokenID", tmp.TokenID))

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
				logger.AtLog.Logger.Error(fmt.Sprintf("InsertOrUpdateNft.%s", contract), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
			}

			_, err = u.Repo.InsertOne(tmp)
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("InsertOrUpdateNft.%s", contract), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
				return err
			}
			if tmp.ContractAddress == artfactAddress {
				u.NewArtifactNotify(tmp)
			}
			if tmp.ContractAddress == bnsAddress {
				go func() {
					name, err := u.BnsService.NameByToken(tmp.TokenID)
					if err == nil {
						u.NewNameNotify(name)
					} else {
						logger.AtLog.Logger.Error("InsertOrUpdateNft.NewNameNotify", zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
					}
				}()
			}

		} else {
			logger.AtLog.Logger.Error(fmt.Sprintf("InsertOrUpdateNft.%s", contract), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
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
				logger.AtLog.Logger.Error(fmt.Sprintf("InsertOrUpdateNft.%s", contract), zap.String("owner", tmp.Owner), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
			}

			_, err = u.Repo.UpdateNftOwner(tmp.ContractAddress, tmp.TokenID, tmp.Owner)
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("InsertOrUpdateNft.%s", contract), zap.String("owner", tmp.Owner), zap.String("contract", contract), zap.Int("tokenID", int(tmp.TokenIDInt)), zap.Error(err))
			}

		}
	}

	return nil
}

func (u *Usecase) UpdateCollectionThumbnails(ctx context.Context) error {
	collections, err := u.Repo.CollectionThumbnailByNfts()
	if err != nil {
		logger.AtLog.Logger.Error("UpdateCollectionThumbnails", zap.Error(err))
		return err
	}

	for _, collection := range collections {
		err = u.Repo.UpdateCollectionThumbnail(ctx, collection.Contract, collection.NftImage)
		if err != nil {
			logger.AtLog.Logger.Error("UpdateCollectionThumbnails", zap.String("contract", collection.Contract), zap.String("nftImage", collection.NftImage), zap.Error(err))
			return err
		}

		logger.AtLog.Logger.Info("UpdateCollectionThumbnails", zap.String("contract", collection.Contract), zap.String("nftImage", collection.NftImage))

	}

	return nil
}

func (u *Usecase) UpdateNftOwner(ctx context.Context, contractAddress string, tokenID string, newOwner string) (*entity.Nfts, error) {
	contractAddress = strings.ToLower(contractAddress)
	tokenID = strings.ToLower(tokenID)
	newOwner = strings.ToLower(newOwner)
	nft, err := u.Repo.GetNft(contractAddress, tokenID)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(nft.Owner) == strings.ToLower(newOwner) {
		return nil, errors.New(fmt.Sprintf("Token is belong to %s", newOwner))
	}

	_, err = u.Repo.UpdateNftOwner(contractAddress, tokenID, newOwner)
	if err != nil {
		return nil, err
	}
	nft.Owner = newOwner
	return nft, nil
}

func (u *Usecase) RefreshNft(ctx context.Context, contractAddress string, tokenID string) (interface{}, error) {
	data, err := u.NftExplorer.RefreshNft(contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("RefreshNft", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (u *Usecase) UpdateCollectionIndex(ctx context.Context, contractAddress string, index int) error {
	err := u.Repo.UpdateCollectionIndex(ctx, contractAddress, index)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateCollectionIndex", zap.String("contractAddress", contractAddress), zap.Int("index", index), zap.Error(err))
		return err
	}
	return nil
}

func (u *Usecase) UpdateAllCollectionIndex(ctx context.Context) error {
	collections, err := u.Repo.AllCollections()
	if err != nil {
		logger.AtLog.Logger.Error("UpdateAllCollectionIndex", zap.Error(err))
		return err
	}
	sort.Slice(collections, func(i, j int) bool {
		return collections[i].DeployedAtBlock <= collections[j].DeployedAtBlock
	})

	for i, coll := range collections {
		u.UpdateCollectionIndex(ctx, coll.Contract, i+1)
	}
	return nil
}
