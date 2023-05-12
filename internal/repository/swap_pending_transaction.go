package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapPendingTransaction(ctx context.Context, filter entity.SwapPendingTransactionsFilter) (*entity.SwapPendingTransactions, error) {
	var swapPairHistories entity.SwapPendingTransactions
	err := r.DB.Collection(utils.COLLECTION_SWAP_PENDING_TRANSACTION).FindOne(ctx, r.parseSwapPendingTransaction(filter)).Decode(&swapPairHistories)
	if err != nil {
		return nil, err
	}
	return &swapPairHistories, nil
}

func (r *Repository) parseSwapPendingTransaction(filter entity.SwapPendingTransactionsFilter) bson.M {
	andCond := make([]bson.M, 0)
	if len(filter.TxHash) > 0 {
		andCond = append(andCond, bson.M{"tx_hash": bson.M{"$in": filter.TxHash}})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindSwapPendingTransactionList(ctx context.Context, filter entity.SwapPendingTransactionsFilter) ([]*entity.SwapPendingTransactions, error) {
	pairs := []*entity.SwapPendingTransactions{}

	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)
	options.SetSort(bson.D{{"timestamp", -1}})

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_PENDING_TRANSACTION).Find(ctx, r.parseSwapPendingTransaction(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		pair := &entity.SwapPendingTransactions{}
		err = cursor.Decode(pair)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}
