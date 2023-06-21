package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	soul_contract "dapp-moderator/utils/contracts/soul"
	"dapp-moderator/utils/logger"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (u *Usecase) UpdateAuctionStatus(ctx context.Context) {
	chainLatestBlock, err := u.TCPublicNode.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
	}

	limit := int64(100)
	offset := int64(0)
	success := 0
	for {
		var auctions []*entity.Auction
		if err := u.Repo.Find(utils.COLLECTION_AUCTION, bson.D{
			{"status", entity.AuctionStatusInProgress.Ordinal()},
		}, limit, offset, &auctions, bson.D{{"_id", 1}}); err != nil {
			logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
			return
		}

		if len(auctions) == 0 {
			break
		}
		for _, item := range auctions {
			endTime, ok := new(big.Int).SetString(item.EndTimeBlock, 10)
			if !ok {
				logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus-parseEndTimeToBigInt", zap.Error(err))
			}

			if chainLatestBlock.Cmp(endTime) > 0 {
				if updateResult, err := u.Repo.UpdateOne(utils.COLLECTION_AUCTION, bson.D{{"_id", item.ID}}, bson.M{
					"$set": bson.M{"status": entity.AuctionStatusEnded.Ordinal()},
				}); err != nil {
					logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
					return
				} else {
					success += int(updateResult.ModifiedCount)
				}
			}
		}
		offset = offset + limit
	}

	logger.AtLog.Logger.Info("Finish Usecase.UpdateAuctionStatus", zap.Int64("chainLatestBlock", chainLatestBlock.Int64()),
		zap.Int("success", success))
}

func (u *Usecase) AuctionDetail(contractAddr, tokenID string) (*response.AuctionDetailResponse, error) {
	soulContract, err := soul_contract.NewSoul(common.HexToAddress(contractAddr), u.TCPublicNode.GetClient())
	if err != nil {
		return nil, err
	}
	tokenIDBigInt, ok := new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, errors.New("invalid token id")
	}
	resp, err := soulContract.Auctions(&bind.CallOpts{
		Context: context.Background(),
	}, tokenIDBigInt)
	if err != nil {
		logger.AtLog.Logger.Error("httpDelivery.auctionDetail", zap.Error(err))
		return nil, err
	}

	available, err := soulContract.Available(&bind.CallOpts{Context: context.Background()}, tokenIDBigInt)
	if err != nil {
		return nil, err
	}
	var auctionEntity = &entity.Auction{}
	err = u.Repo.FindOneWithResult(utils.COLLECTION_AUCTION, bson.M{
		"collection_address": contractAddr,
		"token_id":           tokenID,
	}, auctionEntity, &options.FindOneOptions{
		Sort: bson.M{"_id": -1},
	})
	if err != nil || auctionEntity.ID.IsZero() {
		return &response.AuctionDetailResponse{
			Available: available,
		}, nil
	}

	chainLatestBlock, err := u.TCPublicNode.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
	}

	status := entity.AuctionStatusInProgress
	if chainLatestBlock.Cmp(resp.EndTime) > 0 {
		status = entity.AuctionStatusEnded
	}
	if resp.Settled {
		status = entity.AuctionStatusSettled
	}

	if auctionEntity.Status.Ordinal() != status.Ordinal() {
		if _, err := u.Repo.UpdateOne(utils.COLLECTION_AUCTION, bson.D{{"_id", auctionEntity.ID}}, bson.M{
			"$set": bson.M{"status": status.Ordinal()},
		}); err != nil {
			logger.AtLog.Logger.Error("Usecase.AuctionDetail", zap.Error(err))
			return nil, err
		}
	}

	availableToFE := available
	if status == entity.AuctionStatusInProgress || status == entity.AuctionStatusEnded {
		availableToFE = true
	}

	return &response.AuctionDetailResponse{
		Available:      availableToFE,
		AuctionStatus:  status,
		HighestBid:     resp.Amount.String(),
		EndTime:        resp.EndTime.String(),
		DBAuctionID:    auctionEntity.ID.Hex(),
		ChainAuctionID: new(big.Int).SetBytes(resp.AuctionId[:]).String(),
	}, nil
}
