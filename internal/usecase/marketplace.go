package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/contracts/bns"
	"dapp-moderator/utils/contracts/generative_marketplace_lib"
	soul_contract "dapp-moderator/utils/contracts/soul"
	"dapp-moderator/utils/generative_nft_contract"
	"dapp-moderator/utils/logger"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
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

		activity.UserAAddress = strings.ToLower(event.Buyer.Hex())
		activity.UserBAddress = strings.ToLower(event.Data.Seller.Hex())
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
	case entity.BNSResolverUpdated:
		bnsContract, err := bns.NewBns(chainLog.Address, u.TCPublicNode.GetClient())
		event, err := bnsContract.ParseResolverUpdated(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("marketplaceContract.ParseResolverUpdated", zap.Error(err))
			return nil, nil, err
		}

		activity.UserBAddress = strings.ToLower(event.Addr.Hex())
		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.Id.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		//activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.BNSResolverCreated:
		bnsContract, err := bns.NewBns(chainLog.Address, u.TCPublicNode.GetClient())
		event, err := bnsContract.ParseNameRegistered(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("marketplaceContract.ParseResolverUpdated", zap.Error(err))
			return nil, nil, err
		}

		//activity.UserBAddress = strings.ToLower(event.)
		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.Id.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		//activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.BNSPfpUpdated:
		bnsContract, err := bns.NewBns(chainLog.Address, u.TCPublicNode.GetClient())
		event, err := bnsContract.ParsePfpUpdated(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("marketplaceContract.ParseResolverUpdated", zap.Error(err))
			return nil, nil, err
		}
		//activity.UserBAddress = strings.ToLower(event.Filename)
		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.Id.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		//activity.OfferingID = strings.ToLower(fmt.Sprintf("%x", event.OfferingId))
		return activity, event, nil
	case entity.AuctionCreatedActivity:
		soulContract, err := soul_contract.NewSoul(chainLog.Address, u.TCPublicNode.GetClient())
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.NewSoulContract", zap.Error(err))
			return nil, nil, err
		}

		event, err := soulContract.ParseAuctionCreated(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.ParseAuctionCreated", zap.Error(err))
			return nil, nil, err
		}

		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.TokenId.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		activity.UserAAddress = strings.ToLower(event.Sender.Hex())
		activity.AuctionID = utils.ToPtr(new(big.Int).SetBytes(event.AuctionId[:]).String())

		return activity, event, nil
	case entity.AuctionBidActivity:
		soulContract, err := soul_contract.NewSoul(chainLog.Address, u.TCPublicNode.GetClient())
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.NewSoulContract", zap.Error(err))
			return nil, nil, err
		}

		event, err := soulContract.ParseAuctionBid(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.ParseAuctionCreated", zap.Error(err))
			return nil, nil, err
		}

		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.TokenId.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		activity.UserAAddress = strings.ToLower(event.Sender.Hex())
		activity.Amount = event.Value.Int64()
		activity.AmountStr = event.Value.String()
		activity.AuctionID = utils.ToPtr(new(big.Int).SetBytes(event.Auction.AuctionId[:]).String())

		return activity, event, nil
	case entity.AuctionSettledActivity:
		soulContract, err := soul_contract.NewSoul(chainLog.Address, u.TCPublicNode.GetClient())
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.NewSoulContract", zap.Error(err))
			return nil, nil, err
		}

		event, err := soulContract.ParseAuctionSettled(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.ParseAuctionSettled", zap.Error(err))
			return nil, nil, err
		}

		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.TokenId.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		activity.AuctionID = utils.ToPtr(new(big.Int).SetBytes(event.Auction.AuctionId[:]).String())

		return activity, event, nil
	case entity.AuctionClaimActivity:
		soulContract, err := soul_contract.NewSoul(chainLog.Address, u.TCPublicNode.GetClient())
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.NewSoulContract", zap.Error(err))
			return nil, nil, err
		}

		event, err := soulContract.ParseAuctionClaimBid(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.ParseAuctionClaimBid", zap.Error(err))
			return nil, nil, err
		}

		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.TokenId.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		activity.AmountStr = event.Value.String()
		activity.UserAAddress = strings.ToLower(event.Sender.String())
		activity.AuctionID = utils.ToPtr(new(big.Int).SetBytes(event.AuctionId[:]).String())

		return activity, event, nil
	case entity.SoulUnlockFeature:
		soulContract, err := soul_contract.NewSoul(chainLog.Address, u.TCPublicNode.GetClient())
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.NewSoulContract", zap.Error(err))
			return nil, nil, err
		}

		event, err := soulContract.ParseUnlockFeature(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error("soul_contract.ParseAuctionClaimBid", zap.Error(err))
			return nil, nil, err
		}

		activity.Time = &tm
		activity.InscriptionID = strings.ToLower(event.TokenId.String())
		activity.CollectionContract = strings.ToLower(chainLog.Address.Hex())
		activity.BlockNumber = event.BlockNumber.Uint64()
		activity.UserAAddress = strings.ToLower(event.User.String())
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

	//Send message to discord
	_, err = u.NewListForSaleNotify(listing)
	if err != nil {
		logger.AtLog.Logger.Error("MarketplaceCreateListing - ListForSaleNotify", zap.Error(err), zap.String("offeringId", listing.OfferingId), zap.String("tokenId", listing.TokenId))
		//return err
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

	//Send message to discord
	_, err = u.NewPurchaseTokenNotify(offeringID)
	if err != nil {
		logger.AtLog.Logger.Error("MarketplaceCreateListing - NewPurchaseTokenNotify", zap.Error(err), zap.String("offeringId", offeringID))
		//return err
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

	if strings.ToLower(os.Getenv("ENV")) == strings.ToLower("production") ||
		strings.ToLower(os.Getenv("ENV")) == strings.ToLower("develop") {

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

func (u *Usecase) MarketplaceCollectionAttributes(ctx context.Context, f entity.FilterMarketplaceCollectionAttribute) ([]structure.MarketplaceCollectionAttributeResp, error) {
	obj := []entity.MarketplaceCollectionAttribute{}

	offset := f.Offset
	limit := int64(10000)

	filter := bson.D{}
	if f.ContractAddress != nil && *f.ContractAddress != "" {
		filter = append(filter, bson.E{"contract", *f.ContractAddress})
	}

	if f.TraitType != nil && *f.TraitType != "" {
		filter = append(filter, bson.E{"trait_type", *f.TraitType})
	}

	if f.Value != nil && *f.Value != "" {
		filter = append(filter, bson.E{"value", *f.Value})
	}

	if f.Percent != nil && *f.Percent != 0 {
		filter = append(filter, bson.E{"percent", *f.Percent})
	}

	sort := bson.D{
		{"trait_type", entity.SORT_ASC},
	}

	err := u.Repo.Find(utils.VIEW_MARKETPLACE_COLLECTION_ATTRIBUTES_PERCENT, filter, limit, offset, &obj, sort)
	if err != nil {
		logger.AtLog.Logger.Error("MarketplaceCollectionAttributes", zap.Error(err), zap.Any("filter", f))
		return nil, err
	}

	resp := []structure.MarketplaceCollectionAttributeResp{}
	group := make(map[string][]entity.MarketplaceCollectionAttribute)

	for _, item := range obj {
		gd, ok := group[item.TraitType]
		if ok {
			group[item.TraitType] = append(gd, item)
		} else {
			gd = []entity.MarketplaceCollectionAttribute{}
			group[item.TraitType] = append(gd, item)
		}

	}

	for key, item := range group {
		respItem := structure.MarketplaceCollectionAttributeResp{}
		respItem.TraitName = key

		itemValues := []structure.MarketplaceCollectionAttributeValue{}
		for _, iv := range item {
			itemValue := structure.MarketplaceCollectionAttributeValue{
				Value:  iv.Value,
				Rarity: iv.Percent * 100,
			}

			itemValues = append(itemValues, itemValue)
		}

		respItem.TraitValuesStat = itemValues

		resp = append(resp, respItem)
	}

	logger.AtLog.Logger.Info("MarketplaceCollectionAttributes", zap.Any("obj", obj), zap.Any("filter", f))
	return resp, nil
}

func (u *Usecase) MarketplaceBNSResolverUpdated(eventData interface{}, chainLog types.Log) error {
	event := eventData.(*bns.BnsResolverUpdated)
	tokenID := event.Id.String()
	resolver := event.Addr.Hex()

	logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceBNSResolverUpdated - bns: %s", tokenID), zap.String("tokenID", tokenID), zap.String("resolver", resolver))

	updated, err := u.Repo.UpdateBnsResolver(tokenID, resolver)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceBNSResolverUpdated -  %s", tokenID), zap.String("resolver", resolver), zap.Error(err))
		return err
	}

	logger.AtLog.Logger.Info(fmt.Sprintf("MarketplaceBNSResolverUpdated -  %s", tokenID), zap.String("resolver", resolver), zap.Any("updated", updated))

	// nếu bns này đang sử dụng cho 1 resolver nào đó -> update lại default cho resolver đó
	bnsEntity := &entity.Bns{}
	result, err := u.Repo.FindOne(utils.COLLECTION_BNS, bson.D{{"token_id", tokenID}})
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceBNSResolverUpdated -  %s", tokenID),
			zap.String("resolver", resolver), zap.Error(err))
		return err
	}
	if err := result.Decode(bnsEntity); err != nil {
		return err
	}

	result, err = u.Repo.FindOne(utils.COLLECTION_BNS_DEFAULT, bson.D{{"bns_default_id", bnsEntity.ID}})
	if err == nil {
		bnsDefault := &entity.BNSDefault{}
		if err := result.Decode(bnsDefault); err == nil && bnsDefault.Resolver != resolver {
			if _, err := u.Repo.DeleteOne(utils.COLLECTION_BNS_DEFAULT, bson.D{{"_id", bnsDefault.ID}}); err == nil {
				_, _ = u.BnsDefault(context.TODO(), bnsDefault.Resolver)
			}
		}
	}

	return nil
}

func (u *Usecase) MarketplaceBNSCreated(eventData interface{}, chainLog types.Log) error {
	event := eventData.(*bns.BnsNameRegistered)
	tokenID := event.Id.String()
	contractAddress := chainLog.Address.Hex()
	logger.AtLog.Logger.Info(fmt.Sprintf("MarketplaceBNSCreated - bns: %s", tokenID), zap.String("tokenID", tokenID), zap.String("contract_address", contractAddress))

	inputChan := make(chan entity.Nfts, 1)
	outputChan := make(chan structure.BnsRespChan, 1)

	bnsS, err := bns.NewBns(common.HexToAddress(chainLog.Address.Hex()), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceBNSCreated - bns: %s", tokenID), zap.Error(err))
		return err
	}

	go u.BnsItemWorker(context.Background(), bnsS, inputChan, outputChan)

	inputChan <- entity.Nfts{
		TokenID:         tokenID,
		ContractAddress: contractAddress,
	}

	dataFChan := <-outputChan
	if dataFChan.Err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceBNSCreated - bns: %s", tokenID), zap.Error(dataFChan.Err))
		return dataFChan.Err
	}

	bnsFChan := dataFChan.Bns
	_, err = u.Repo.InsertOne(bnsFChan)
	if dataFChan.Err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplaceBNSCreated - InsertOne bns: %s", tokenID), zap.Error(dataFChan.Err), zap.Any("bnsFChan", bnsFChan))
		return dataFChan.Err
	}

	return nil
}

func (u *Usecase) MarketplacePFPUpdated(eventData interface{}, chainLog types.Log) error {
	event := eventData.(*bns.BnsPfpUpdated)
	tokenID := event.Id.String()
	pfp := event.Filename

	logger.AtLog.Logger.Error(fmt.Sprintf("MarketplacePFPUpdated - bns: %s", tokenID), zap.String("tokenID", tokenID), zap.String("pfp", pfp))

	updated, err := u.Repo.UpdateBnsPfp(tokenID, pfp)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplacePFPUpdated -  %s", tokenID), zap.String("pfp", pfp), zap.Error(err))
		return err
	}

	if err := u.UploadBnsPFPToGCS(chainLog.Address.Hex(), tokenID); err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MarketplacePFPUpdated - UploadBnsPFPToGCS -  %s", tokenID), zap.String("pfp", pfp), zap.Error(err))
	}

	logger.AtLog.Logger.Info(fmt.Sprintf("MarketplacePFPUpdated -  %s", tokenID), zap.String("pfp", pfp), zap.Any("updated", updated))
	return nil
}

func (u *Usecase) UploadBnsPFPToGCS(contractAddress string, tokenID string) error {
	bnsRow, err := u.Repo.FindOne(utils.COLLECTION_BNS, bson.D{{"token_id", tokenID}})
	if err != nil {
		logger.AtLog.Logger.Error("UploadBnsPFPToGCS.FindOne got error", zap.String("tokenID", tokenID), zap.Error(err))
		return err
	}
	var bnsEntity = &entity.Bns{}
	if err := bnsRow.Decode(bnsEntity); err != nil {
		return err
	}

	if bnsEntity.Pfp == "" {
		return errors.New("pfp is empty")
	}

	bnsS, err := bns.NewBns(common.HexToAddress(contractAddress), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UploadBnsPFPToGCS - bns: %s", tokenID), zap.Error(err))
		return err
	}

	tokenIdInt, err := strconv.Atoi(bnsEntity.TokenID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UploadBnsPFPToGCS - bns: %s", tokenID), zap.Error(err))
		return err
	}
	tokenId := big.NewInt(int64(tokenIdInt))
	bytes, err := bnsS.GetPfp(&bind.CallOpts{Context: context.Background()}, tokenId)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UploadBnsPFPToGCS - getPfp from contract with token_id: %s", tokenID), zap.Error(err))
		return err
	}

	arr := strings.Split(bnsEntity.Pfp, "/")
	fileName := fmt.Sprintf("%s_%s", tokenID, arr[len(arr)-1])
	base64Str := base64.StdEncoding.EncodeToString(bytes)

	object, err := u.Storage.UploadBaseToBucket(base64Str, fileName)
	if err != nil {
		logger.AtLog.Logger.Error("UploadBnsPFPToGCS - UploadBaseToBucket", zap.String("tokenID", tokenID), zap.Any("bns", bnsEntity), zap.Error(err))
		return err
	}

	logger.AtLog.Logger.Info("upload pfp to gcs success", zap.Any("response", object))
	_, err = u.Repo.UpdateBnsPfpData(tokenID, &entity.BnsPfpData{
		GCSUrl:   fmt.Sprintf("%v/%v/%v", u.Config.Gcs.Endpoint, u.Config.Gcs.Bucket, fileName),
		Filename: fileName,
	})
	if err != nil {
		logger.AtLog.Logger.Error("UploadBnsPFPToGCSRepo - UpdateBnsPfpData", zap.String("tokenID", tokenID), zap.Any("bns", bnsEntity), zap.Error(err))
		return err
	}

	return nil
}

func (u *Usecase) HandleAuctionCreated(data interface{}, chainLog types.Log) error {
	eventData, ok := data.(*soul_contract.SoulAuctionCreated)
	if !ok {
		logger.AtLog.Logger.Error("HandleAuctionCreated - assert eventData failed", zap.String("tokenID", eventData.TokenId.String()))
		return errors.New("event data is not correct")
	}
	logger.AtLog.Logger.Info("HandleAuctionCreated", zap.String("tokenID", eventData.TokenId.String()),
		zap.String("contract", chainLog.Address.Hex()), zap.Uint64("startTime", eventData.StartTime.Uint64()),
		zap.Uint64("endTime", eventData.EndTime.Uint64()))

	tokenIDInt, _ := strconv.Atoi(eventData.TokenId.String())
	auctionEntity := &entity.Auction{
		CollectionAddress: strings.ToLower(chainLog.Address.Hex()),
		TokenID:           strings.ToLower(eventData.TokenId.String()),
		TokenIDInt:        uint64(tokenIDInt),
		AuctionID:         new(big.Int).SetBytes(eventData.AuctionId[:]).String(),
		StartTimeBlock:    eventData.StartTime.String(),
		EndTimeBlock:      eventData.EndTime.String(),
		Status:            entity.AuctionStatusInProgress,
	}

	_, err := u.Repo.InsertOne(auctionEntity)
	if err != nil {
		logger.AtLog.Logger.Error("useCase.HandleAuctionCreated-InsertOne", zap.Error(err))
		return err
	}

	return nil
}

func (u *Usecase) HandleAuctionBid(data interface{}, chainLog types.Log) error {
	eventData, ok := data.(*soul_contract.SoulAuctionBid)
	if !ok {
		logger.AtLog.Logger.Error("HandleAuctionBid - assert eventData failed", zap.String("tokenID", eventData.TokenId.String()))
		return errors.New("event data is not correct")
	}
	logger.AtLog.Logger.Info("HandleAuctionBid", zap.String("tokenID", eventData.TokenId.String()),
		zap.Any("eventData", eventData), zap.String("contract", chainLog.Address.Hex()))

	chainAuctionID := new(big.Int).SetBytes(eventData.Auction.AuctionId[:]).String()
	// async here => so retry to have data auctionEntity
	auctionEntity := &entity.Auction{}
	for {
		err := u.Repo.FindOneWithResult(utils.COLLECTION_AUCTION, bson.M{
			"auction_id": chainAuctionID,
		}, auctionEntity)
		if err == nil {
			break
		}

		time.Sleep(3 * time.Second)
	}

	auctionBid, err := u.validateAuctionBid(eventData, auctionEntity, chainLog)
	if err != nil {
		logger.AtLog.Logger.Error("HandleAuctionBid - validateAuctionBid", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		return err
	}

	if _, err = u.Repo.InsertOne(auctionBid); err != nil {
		logger.AtLog.Logger.Error("HandleAuctionBid - InsertOne", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		return err
	}

	// Handle auction_bid_summary
	now := time.Now()
	auctionBidSummary := &entity.AuctionBidSummary{}
	if err := u.Repo.FindOneWithResult(utils.COLLECTION_AUCTION_BID_SUMMARY, bson.M{
		"chain_auction_id": chainAuctionID,
		"sender":           auctionBid.Sender,
	}, auctionBidSummary); err != nil {
		if _, err := u.Repo.InsertOne(&entity.AuctionBidSummary{
			BaseEntity: entity.BaseEntity{
				UpdatedAt: &now,
			},
			DBAuctionID:       auctionEntity.ID,
			ChainAuctionID:    chainAuctionID,
			TokenID:           auctionBid.TokenID,
			CollectionAddress: auctionBid.CollectionAddress,
			TotalAmount:       eventData.Value.String(),
			Sender:            auctionBid.Sender,
		}); err != nil {
			logger.AtLog.Logger.Error("HandleAuctionBid - InsertOne - AuctionBidSummary", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		}
	} else {
		currentAmount, _ := new(big.Int).SetString(auctionBidSummary.TotalAmount, 10)
		newAmount := new(big.Int).Add(currentAmount, eventData.Value)
		if _, err := u.Repo.UpdateOne(utils.COLLECTION_AUCTION_BID_SUMMARY, bson.D{
			{"chain_auction_id", chainAuctionID},
			{"sender", auctionBid.Sender},
		}, bson.M{"$set": bson.M{
			"total_amount": newAmount.String(),
			"updated_at":   now,
		}}); err != nil {
			logger.AtLog.Logger.Error("HandleAuctionBid - UpdateOne - AuctionBidSummary", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		}
	}

	// Update total amount of auction
	if auctionEntity.TotalAmount == "" {
		auctionEntity.TotalAmount = "0"
	}
	totalAmount, ok := new(big.Int).SetString(auctionEntity.TotalAmount, 10)
	if ok {
		totalAmount = new(big.Int).Add(totalAmount, eventData.Value)
	} else {
		totalAmount, _ = new(big.Int).SetString("0", 10)
	}
	if _, err := u.Repo.UpdateOne(utils.COLLECTION_AUCTION, bson.D{
		{"auction_id", chainAuctionID},
	}, bson.M{"$set": bson.M{
		"total_amount": totalAmount.String(),
	}}); err != nil {
		logger.AtLog.Logger.Error("HandleAuctionBid - UpdateOne", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		return err
	}

	return nil
}

func (u *Usecase) validateAuctionBid(auctionBidEvent *soul_contract.SoulAuctionBid, auction *entity.Auction, chainLog types.Log) (*entity.AuctionBid, error) {
	if os.Getenv("ENV") != "local" {
		chainLatestBlock, err := u.TCPublicNode.GetBlockNumber()
		if err != nil {
			logger.AtLog.Logger.Error("validateAuctionBid - GetBlockNumber", zap.Error(err))
			return nil, err
		}
		startTime, ok := new(big.Int).SetString(auction.StartTimeBlock, 10)
		if !ok {
			return nil, errors.New("invalid parse auction start time")
		}
		endTime, ok := new(big.Int).SetString(auction.EndTimeBlock, 10)
		if !ok {
			return nil, errors.New("invalid parse auction end time")
		}

		if !(startTime.Cmp(chainLatestBlock) <= 1 && chainLatestBlock.Cmp(endTime) <= 1) {
			logger.AtLog.Logger.Error("validateAuctionBid - auction is not in progress", zap.Any("auction", auctionBidEvent.Auction))
			return nil, errors.New("auction is not in progress")
		}
	}

	return &entity.AuctionBid{
		DBAuctionID:       auction.ID,
		ChainAuctionID:    auction.AuctionID,
		TokenID:           strings.ToLower(auctionBidEvent.TokenId.String()),
		CollectionAddress: strings.ToLower(chainLog.Address.Hex()),
		Amount:            auctionBidEvent.Value.String(),
		Sender:            strings.ToLower(auctionBidEvent.Sender.Hex()),
	}, nil
}

func (u *Usecase) HandleAuctionSettle(data interface{}, chainLog types.Log) error {
	eventData, ok := data.(*soul_contract.SoulAuctionSettled)
	if !ok {
		logger.AtLog.Logger.Error("HandleAuctionBid - assert eventData failed", zap.String("tokenID", eventData.TokenId.String()))
		return errors.New("event data is not correct")
	}

	logger.AtLog.Logger.Info("HandleAuctionSettle", zap.String("tokenID", eventData.TokenId.String()),
		zap.Any("eventData", eventData), zap.String("contract", chainLog.Address.Hex()))

	chainAuctionID := new(big.Int).SetBytes(eventData.Auction.AuctionId[:]).String()
	auction, err := u.Repo.FindAuctionByChainAuctionID(context.TODO(), chainAuctionID)
	if err != nil {
		logger.AtLog.Logger.Error("HandleAuctionSettle - FindAuctionByChainAuctionID", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		return err
	}

	updateAuction := &bson.M{
		"status":       entity.AuctionStatusSettled.Ordinal(),
		"winner":       utils.ToPtr(strings.ToLower(eventData.Winner.Hex())),
		"total_amount": eventData.Amount.String(),
	}

	_, err = u.Repo.UpdateOne(utils.COLLECTION_AUCTION, bson.D{{"_id", auction.ID}}, bson.M{
		"$set": updateAuction,
	})
	if err != nil {
		logger.AtLog.Logger.Error("HandleAuctionSettle - UpdateOne", zap.String("auctionObjectID", auction.ID.Hex()), zap.Error(err))
		return err
	}

	return nil
}

func (u *Usecase) HandleAuctionClaim(data interface{}, chainLog types.Log) error {
	eventData, ok := data.(*soul_contract.SoulAuctionClaimBid)
	if !ok {
		logger.AtLog.Logger.Error("HandleAuctionClaim - assert eventData failed", zap.String("tokenID", eventData.TokenId.String()))
		return errors.New("event data is not correct")
	}

	logger.AtLog.Logger.Info("HandleAuctionClaim", zap.String("tokenID", eventData.TokenId.String()),
		zap.Any("eventData", eventData), zap.String("contract", chainLog.Address.Hex()))

	chainAuctionID := new(big.Int).SetBytes(eventData.AuctionId[:]).String()
	auction, err := u.Repo.FindAuctionByChainAuctionID(context.TODO(), chainAuctionID)
	if err != nil {
		logger.AtLog.Logger.Error("HandleAuctionSettle - FindAuctionByChainAuctionID", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		return err
	}
	if auction.Status == entity.AuctionStatusInProgress {
		logger.AtLog.Logger.Error("HandleAuctionSettle - auctionInProgress - Cannot claim", zap.String("tokenID", eventData.TokenId.String()))
		return errors.New("auctionInProgress - Cannot claim")
	}

	auctionClaim := &entity.AuctionClaim{
		DBAuctionID:       auction.ID,
		ChainAuctionID:    chainAuctionID,
		TokenID:           strings.ToLower(eventData.TokenId.String()),
		CollectionAddress: strings.ToLower(chainLog.Address.Hex()),
		Claimer:           strings.ToLower(eventData.Sender.Hex()),
		Amount:            eventData.Value.String(),
	}

	if _, err := u.Repo.InsertOne(auctionClaim); err != nil {
		logger.AtLog.Logger.Error("HandleAuctionClaim - InsertOne", zap.String("tokenID", eventData.TokenId.String()), zap.Error(err))
		return err
	}

	return nil
}

func (u *Usecase) HandleUnlockFeature(data interface{}, chainLog types.Log) error {
	eventData, ok := data.(*soul_contract.SoulUnlockFeature)
	if !ok {
		logger.AtLog.Logger.Error("HandleUnlockFeature - assert eventData failed", zap.String("tokenID", eventData.TokenId.String()))
		return errors.New("event data is not correct")
	}
	logger.AtLog.Logger.Info("HandleUnlockFeature", zap.Any("eventData", eventData), zap.Any("chainLog", chainLog))
	err := u.SoulNftImageHistoriesCrontab([]string{strings.ToLower(eventData.TokenId.String())})
	if err != nil {
		logger.AtLog.Logger.Error("HandleUnlockFeature - assert eventData failed", zap.Error(err), zap.String("tokenID", eventData.TokenId.String()))
		return errors.New("event data is not correct")
	}
	return nil
}
