package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindAuctionByChainAuctionID(ctx context.Context, auctionID string) (*entity.Auction, error) {
	filter := bson.D{
		{"auction_id", auctionID},
	}

	result := &entity.Auction{}
	resp, err := r.FindOne(utils.COLLECTION_AUCTION, filter)
	if err != nil {
		return nil, err
	}

	if err := resp.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
