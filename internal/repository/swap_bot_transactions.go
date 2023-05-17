package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapBotTransaction(ctx context.Context, filter entity.SwapBotTransactionFilter) (*entity.SwapBotTransaction, error) {
	var swapIdo entity.SwapBotTransaction
	err := r.DB.Collection(utils.COLLECTION_SWAP_BOT_TRANSACTION).FindOne(ctx, r.parseSwapBotTransactionFilter(filter)).Decode(&swapIdo)
	if err != nil {
		return nil, err
	}
	return &swapIdo, nil
}

func (r *Repository) parseSwapBotTransactionFilter(filter entity.SwapBotTransactionFilter) bson.M {
	andCond := make([]bson.M, 0)
	if filter.Address != "" {
		andCond = append(andCond, bson.M{"address": filter.Address})
	}

	if filter.Status >= 0 {
		andCond = append(andCond, bson.M{"status": filter.Status})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindSwapBotTransactions(ctx context.Context, filter entity.SwapBotTransactionFilter) ([]*entity.SwapBotTransaction, error) {
	idos := []*entity.SwapBotTransaction{}
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_BOT_TRANSACTION).Find(ctx, r.parseSwapBotTransactionFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pair entity.SwapBotTransaction
		err = cursor.Decode(&pair)
		if err != nil {
			return nil, err
		}
		idos = append(idos, &pair)
	}
	return idos, nil
}

func (r *Repository) UpdateSwapBotTransaction(ctx context.Context, pair *entity.SwapBotTransaction) error {
	collectionName := pair.CollectionName()
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"_id": pair.ID}, bson.M{"$set": pair})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
