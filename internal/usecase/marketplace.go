package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/contracts/generative_marketplace_lib"
	"dapp-moderator/utils/generative_nft_contract"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"math/big"
	"os"
	"strings"
	"time"
)

func (u *Usecase) ParseMkplaceData(chainLog types.Log, eventType entity.TokenActivityType) (*entity.MarketplaceTokenActivity, interface{}, error) {
	marketplaceContract, err := generative_marketplace_lib.NewGenerativeMarketplaceLib(chainLog.Address, u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error("parseMkplaceData - init marketplace", zap.Uint64("chainLog", chainLog.BlockNumber), zap.Error(err))
		return nil, nil, err
	}

	activity := &entity.MarketplaceTokenActivity{
		Type:         eventType,
		Title:        entity.TokenActivityName[eventType],
		UserBAddress: "",
		BlockNumber:  chainLog.BlockNumber,
		TxHash:       strings.ToLower(chainLog.TxHash.Hex()),
		LogIndex:     chainLog.Index,
	}

	bn := big.NewInt(int64(chainLog.BlockNumber))
	blockInfo, err := u.TCPublicNode.GetBlockByNumber(*bn)
	if err != nil {
		logger.AtLog.Logger.Error("parseMkplaceData - init marketplace", zap.Uint64("chainLog", chainLog.BlockNumber), zap.Error(err))
		return nil, nil, err
	}

	blockTime := blockInfo.Header().Time
	tm := time.Unix(int64(blockTime), 0).UTC()

	switch eventType {
	case entity.TokenListing:
		event, err := marketplaceContract.ParseListingToken(chainLog)
		if err != nil {
			//logger.AtLog.Logger.Error(" marketplaceContract.ParseListingToken", zap.Error(err))
			return nil, nil, err
		}

		activity.UserAAddress = strings.ToLower(event.Data.Seller.Hex())
		activity.Amount = event.Data.Price.Int64()
		activity.Erc20Address = strings.ToLower(event.Data.Erc20Token.Hex())
		activity.Time = &tm
		activity.InscriptionID = event.Data.TokenId.String()
		activity.CollectionContract = strings.ToLower(event.Data.CollectionContract.String())
		activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.TokenPurchase:
		event, err := marketplaceContract.ParsePurchaseToken(chainLog)
		if err != nil {
			//logger.AtLog.Logger.Error("marketplaceContract.ParsePurchaseToken", zap.Error(err))
			return nil, nil, err
		}

		activity.UserAAddress = strings.ToLower(event.Data.Seller.Hex())
		activity.Amount = event.Data.Price.Int64()
		activity.Erc20Address = strings.ToLower(event.Data.Erc20Token.Hex())
		activity.Time = &tm
		activity.InscriptionID = event.Data.TokenId.String()
		activity.CollectionContract = strings.ToLower(event.Data.CollectionContract.String())
		activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.TokenCancelListing:
		event, err := marketplaceContract.ParseCancelListing(chainLog)
		if err != nil {
			//logger.AtLog.Logger.Error(" marketplaceContract.ParseCancelListing", zap.Error(err))
			return nil, nil, err
		}

		activity.UserAAddress = strings.ToLower(event.Data.Seller.Hex())
		activity.Amount = event.Data.Price.Int64()
		activity.Erc20Address = strings.ToLower(event.Data.Erc20Token.Hex())
		activity.Time = &tm
		activity.InscriptionID = event.Data.TokenId.String()
		activity.CollectionContract = strings.ToLower(event.Data.CollectionContract.String())
		activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.TokenMakeOffer:
		event, err := marketplaceContract.ParseMakeOffer(chainLog)
		if err != nil {
			//logger.AtLog.Logger.Error("marketplaceContract.ParseMakeOffer", zap.Error(err))
			return nil, nil, err
		}

		activity.UserAAddress = strings.ToLower(event.Data.Buyer.Hex())
		activity.Amount = event.Data.Price.Int64()
		activity.Erc20Address = strings.ToLower(event.Data.Erc20Token.Hex())
		activity.Time = &tm
		activity.InscriptionID = event.Data.TokenId.String()
		activity.CollectionContract = strings.ToLower(event.Data.CollectionContract.String())
		activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.TokenAcceptOffer:
		event, err := marketplaceContract.ParseAcceptMakeOffer(chainLog)
		if err != nil {
			//logger.AtLog.Logger.Error("marketplaceContract.ParseAcceptMakeOffer", zap.Error(err))
			return nil, nil, err
		}

		activity.UserBAddress = strings.ToLower(event.Buyer.Hex())
		activity.Amount = event.Data.Price.Int64()
		activity.Erc20Address = strings.ToLower(event.Data.Erc20Token.Hex())
		activity.Time = &tm
		activity.InscriptionID = event.Data.TokenId.String()
		activity.CollectionContract = strings.ToLower(event.Data.CollectionContract.String())
		activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.TokenCancelOffer:
		event, err := marketplaceContract.ParseCancelMakeOffer(chainLog)
		if err != nil {
			//logger.AtLog.Logger.Error("marketplaceContract.ParseAcceptMakeOffer", zap.Error(err))
			return nil, nil, err
		}

		activity.UserAAddress = strings.ToLower(event.Data.Buyer.Hex())
		activity.Amount = event.Data.Price.Int64()
		activity.Erc20Address = strings.ToLower(event.Data.Erc20Token.Hex())
		activity.Time = &tm
		activity.InscriptionID = event.Data.TokenId.String()
		activity.CollectionContract = strings.ToLower(event.Data.CollectionContract.String())
		activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.TokenTransfer:
		erc721Contract, err := generative_nft_contract.NewGenerativeNftContract(chainLog.Address, u.TCPublicNode.GetClient())
		event, err := erc721Contract.ParseTransfer(chainLog)
		if err != nil {
			//logger.AtLog.Logger.Error("marketplaceContract.ParseAcceptMakeOffer", zap.Error(err))
			return nil, nil, err
		}

		activity.UserAAddress = strings.ToLower(event.From.Hex())
		activity.UserBAddress = strings.ToLower(event.To.Hex())
		//activity.Amount =
		//activity.Erc20Address =
		activity.Time = &tm
		activity.InscriptionID = event.TokenId.String()
		activity.CollectionContract = strings.ToLower(event.Raw.Address.Hex())
		//activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	}

	return nil, nil, errors.New(fmt.Sprintf("Cannot detect event log - %d - txHash: %s, topics %s ", eventType, chainLog.TxHash, chainLog.Topics[0].String()))
}

func (u *Usecase) MarketplaceCreateListing(eventData interface{}, chainLog types.Log) error {

	event := eventData.(*generative_marketplace_lib.GenerativeMarketplaceLibListingToken)
	listing := &entity.MarketplaceListings{
		OfferingId:         strings.ToLower(fmt.Sprintf("%x", event.OfferingId)),
		CollectionContract: strings.ToLower(event.Data.CollectionContract.String()),
		TokenId:            event.Data.TokenId.String(),
		Seller:             strings.ToLower(event.Data.Seller.String()),
		Erc20Token:         strings.ToLower(event.Data.Erc20Token.String()),
		Price:              event.Data.Price.String(),
		BlockNumber:        chainLog.BlockNumber,
		Status:             entity.MarketPlaceOpen,
		DurationTime:       event.Data.DurationTime.String(),
	}

	err := u.Repo.InsertListing(listing)
	if err != nil {
		logger.AtLog.Logger.Error("MarketplaceCreateListing - InsertListing", zap.Error(err), zap.String("offeringId", listing.OfferingId), zap.String("tokenId", listing.TokenId))

		return err
	}

	return nil
}

func (u *Usecase) MarketplacePurchaseListing(eventData interface{}, chainLog types.Log) error {
	event := eventData.(*generative_marketplace_lib.GenerativeMarketplaceLibPurchaseToken)
	offeringID := strings.ToLower(fmt.Sprintf("%x", event.OfferingId))

	err := u.Repo.PurchaseListing(context.Background(), offeringID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplacePurchaseListing - %s", offeringID), zap.Error(err))
		return err
	}
	return nil
}

func (u *Usecase) MarketplaceCancelListing(eventData interface{}, chainLog types.Log) error {
	event := eventData.(*generative_marketplace_lib.GenerativeMarketplaceLibCancelListing)
	offeringID := strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
	err := u.Repo.CancelListing(context.Background(), offeringID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceCancelListing - %s", offeringID), zap.Error(err))
		return err
	}
	return nil
}

func (u *Usecase) MarketplaceMakeOffer(eventData interface{}, chainLog types.Log) error {

	event := eventData.(*generative_marketplace_lib.GenerativeMarketplaceLibMakeOffer)
	offer := &entity.MarketplaceOffers{
		OfferingId:         strings.ToLower(fmt.Sprintf("%x", event.OfferingId)),
		CollectionContract: strings.ToLower(event.Data.CollectionContract.String()),
		TokenId:            event.Data.TokenId.String(),
		Buyer:              strings.ToLower(event.Data.Buyer.String()),
		Erc20Token:         strings.ToLower(event.Data.Erc20Token.String()),
		Price:              event.Data.Price.String(),
		Status:             entity.MarketPlaceOpen,
		BlockNumber:        chainLog.BlockNumber,
		DurationTime:       event.Data.DurationTime.String(),
	}

	err := u.Repo.InsertOffer(offer)
	if err != nil {
		logger.AtLog.Logger.Error("MarketplaceCreateListing - InsertListing", zap.Error(err), zap.String("offeringId", offer.OfferingId), zap.String("tokenId", offer.TokenId))

		return err
	}

	return nil
}

func (u *Usecase) MarketplaceAcceptOffer(eventData interface{}, chainLog types.Log) error {

	event := eventData.(*generative_marketplace_lib.GenerativeMarketplaceLibAcceptMakeOffer)
	offeringID := strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
	err := u.Repo.AcceptOffer(context.Background(), offeringID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceAcceptOffer - %s", offeringID), zap.Error(err))
		return err
	}
	return nil
}

func (u *Usecase) MarketplaceCancelOffer(eventData interface{}, chainLog types.Log) error {

	event := eventData.(*generative_marketplace_lib.GenerativeMarketplaceLibCancelMakeOffer)
	offeringID := strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
	err := u.Repo.CancelOffer(context.Background(), offeringID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceCancelOffer - %s", offeringID), zap.Error(err))
		return err
	}

	return nil
}

func (u *Usecase) TransferToken(eventData interface{}, chainLog types.Log) error {
	event := eventData.(*generative_nft_contract.GenerativeNftContractTransfer)
	contract := chainLog.Address.Hex()
	tokenIDStr := event.TokenId.String()
	to := event.To.Hex()
	from := event.From.Hex()

	go u.UpdateUploadedFile(eventData, chainLog)

	if strings.ToLower(os.Getenv("ENV")) == strings.ToLower("production") {

		updated, err := u.UpdateNftOwner(context.Background(), contract, tokenIDStr, to)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Uint64("blockNumber", chainLog.BlockNumber), zap.Error(err))
			return err
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Any("updated", updated), zap.Uint64("blockNumber", chainLog.BlockNumber))
	} else {
		logger.AtLog.Logger.Info(fmt.Sprintf("[Testing] UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Uint64("blockNumber", chainLog.BlockNumber))
	}

	return nil
}

func (u *Usecase) InsertActivity(activity *entity.MarketplaceTokenActivity) error {
	err := u.Repo.InsertActivity(activity)
	if err != nil {
		logger.AtLog.Logger.Error("TransferToken - InsertActivity", zap.Error(err), zap.String("tokenId", activity.InscriptionID), zap.String("txHash", activity.TxHash), zap.Uint("log_index", activity.LogIndex), zap.Uint64("block_number", activity.BlockNumber))
		return err
	}

	return nil
}

func (u *Usecase) UpdateUploadedFile(eventData interface{}, chainLog types.Log) error {
	event := eventData.(*generative_nft_contract.GenerativeNftContractTransfer)
	contract := chainLog.Address.Hex()
	tokenIDStr := event.TokenId.String()
	to := event.To.Hex()
	from := event.From.Hex()
	txHash := chainLog.TxHash.Hex()

	if true {
		updated, err := u.Repo.FindUploadedFileByTxHash(txHash)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Uint64("blockNumber", chainLog.BlockNumber), zap.Error(err))
			return err
		}

		err = u.Repo.UpdateUploadedFileTokenID(txHash, tokenIDStr, to, contract)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateNftOwner - UpdateUploadedFileTokenID %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Uint64("blockNumber", chainLog.BlockNumber), zap.Error(err))
			return err
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Any("updated", updated), zap.Uint64("blockNumber", chainLog.BlockNumber))
	} else {
		logger.AtLog.Logger.Info(fmt.Sprintf("[Testing] UpdateNftOwner %s - %s ", contract, tokenIDStr), zap.String("from", from), zap.String("to", to), zap.Uint64("blockNumber", chainLog.BlockNumber))
	}

	return nil
}

func (u *Usecase) MarketplaceCollections(ctx context.Context, filter request.CollectionsFilter) ([]entity.MarketplaceCollections, error) {
	res := []entity.MarketplaceCollections{}
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

	s := bson.D{{sortBy, sort}, {"index", -1}}
	err := u.Repo.Find(utils.VIEW_MARKETPLACE_AGGREGATED_COLLECTIONS, f, int64(*filter.Limit), int64(*filter.Offset), &res, s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *Usecase) MarketplaceCollectionDetail(ctx context.Context, contractAddress string) (*entity.MarketplaceCollections, error) {
	obj := &entity.MarketplaceCollections{}
	sr, err := u.Repo.FindOne(utils.VIEW_MARKETPLACE_AGGREGATED_COLLECTIONS, bson.D{
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
