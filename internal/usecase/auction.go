package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"math/big"

	"go.mongodb.org/mongo-driver/bson"
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
