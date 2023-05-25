package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/contracts/generative_marketplace_lib"
	"dapp-moderator/utils/generative_nft_contract"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
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
		logger.AtLog.Logger.Error("TransferToken - InsertActivity", zap.Error(err), zap.String("tokenId", activity.InscriptionID), zap.String("txHash", activity.TxHash))
		return err
	}

	return nil
}
