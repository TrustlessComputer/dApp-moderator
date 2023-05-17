package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindUserGmPaid(ctx context.Context, filter entity.SwapUserGmPaidFilter) (*entity.SwapUserGmPaid, error) {
	var swapIdo entity.SwapUserGmPaid
	err := r.DB.Collection(utils.COLLECTION_SWAP_USER_GM_PAID).FindOne(ctx, r.parseSwapUserGmPaidFilter(filter)).Decode(&swapIdo)
	if err != nil {
		return nil, err
	}
	return &swapIdo, nil
}

func (r *Repository) parseSwapUserGmPaidFilter(filter entity.SwapUserGmPaidFilter) bson.M {
	andCond := make([]bson.M, 0)

	if filter.Address != "" {
		andCond = append(andCond, bson.M{"user_address": strings.ToLower(filter.Address)})
	}

	if filter.TxHash != "" {
		andCond = append(andCond, bson.M{"tx_hash": strings.ToLower(filter.TxHash)})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}
