package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindUserGmBalance(ctx context.Context, filter entity.SwapUserGmBalanceFilter) (*entity.SwapUserGmBalance, error) {
	var swapIdo entity.SwapUserGmBalance
	err := r.DB.Collection(utils.COLLECTION_SWAP_USER_GM_BALANCE).FindOne(ctx, r.parseSwapUserGmBalanceFilter(filter)).Decode(&swapIdo)
	if err != nil {
		return nil, err
	}
	return &swapIdo, nil
}

func (r *Repository) parseSwapUserGmBalanceFilter(filter entity.SwapUserGmBalanceFilter) bson.M {
	andCond := make([]bson.M, 0)

	if filter.Address != "" {
		andCond = append(andCond, bson.M{"user_address": strings.ToLower(filter.Address)})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}
