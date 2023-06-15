package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (u *Usecase) UpdateAuctionStatus(ctx context.Context) {
	chainLatestBlock, err := u.TCPublicNode.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
	}

	limit := int64(100)
	page := int64(1)
	success := 0
	for {
		var auctions []*entity.Auction
		if err := u.Repo.Find(utils.COLLECTION_AUCTION, bson.D{
			{"status", entity.AuctionStatusInProgress.Ordinal()},
			{"end_time_block", bson.M{"$lt": entity.AuctionStatusInProgress.Ordinal()}},
		}, limit, page, &auctions, bson.D{{"_id", 1}}); err != nil {
			logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
			return
		}

		if len(auctions) == 0 {
			break
		}

		updateResult, err := u.Repo.UpdateMany(utils.COLLECTION_AUCTION, bson.D{
			{"status", entity.AuctionStatusInProgress.Ordinal()},
			{"end_time_block", bson.M{"$lt": entity.AuctionStatusInProgress.Ordinal()}},
		}, bson.M{"$set": bson.M{"status": entity.AuctionStatusEnded.Ordinal()}})
		if err != nil {
			logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
			return
		}

		success += int(updateResult.ModifiedCount)
	}

	logger.AtLog.Logger.Info("Finish Usecase.UpdateAuctionStatus", zap.Int64("chainLatestBlock", chainLatestBlock.Int64()),
		zap.Int("success", success))
}
