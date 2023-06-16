package repository

import (
	"context"
	"dapp-moderator/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindAuctionByChainAuctionID(ctx context.Context, auctionID uint64) (*entity.Auction, error) {
	filter := bson.D{
		{"auction_id", auctionID},
	}

	result := &entity.Auction{}
	resp, err := r.FindOne(entity.Auction{}.CollectionName(), filter)
	if err != nil {
		return nil, err
	}

	if err := resp.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
