package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *Repository) FindListUserGmBalance(ctx context.Context, filter entity.SwapUserGmBalanceFilter) ([]*entity.SwapUserGmBalance, error) {
	tokens := []*entity.SwapUserGmBalance{}
	options := options.Find()
	options.SetSkip(0)
	options.SetLimit(50000)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_USER_GM_BALANCE).Find(ctx, r.parseSwapUserGmBalanceFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var token entity.SwapUserGmBalance
		err = cursor.Decode(&token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, &token)
	}
	return tokens, nil
}

func (r *Repository) UpdateSwapUserGmBalance(ctx context.Context, pair *entity.SwapUserGmBalance) error {
	collectionName := pair.CollectionName()
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"user_address": pair.UserAddress}, bson.M{"$set": pair})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
