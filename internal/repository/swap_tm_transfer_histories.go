package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindTmTransferHistory(ctx context.Context, filter entity.SwapTmTransferHistoriesFilter) (*entity.SwapTmTransferHistories, error) {
	var swapPairHistories entity.SwapTmTransferHistories
	err := r.DB.Collection(utils.COLLECTION_SWAP_TOKEN_TRANSFER_HISTORY).FindOne(ctx, r.parseTmTransferHistories(filter)).Decode(&swapPairHistories)
	if err != nil {
		return nil, err
	}
	return &swapPairHistories, nil
}

func (r *Repository) parseTmTransferHistories(filter entity.SwapTmTransferHistoriesFilter) bson.M {
	andCond := make([]bson.M, 0)
	if filter.TxHash != "" {
		andCond = append(andCond, bson.M{"tx_hash": filter.TxHash})
	}

	if filter.Index > 0 {
		andCond = append(andCond, bson.M{"index": filter.Index})
	}

	if filter.UserAddress != "" {
		orCond := make([]bson.M, 0)
		orCond = append(orCond, bson.M{"from": filter.UserAddress})
		orCond = append(orCond, bson.M{"to": filter.UserAddress})
		filterOr := bson.M{"$or": orCond}

		andCond = append(andCond, filterOr)
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindTmTransferHistories(ctx context.Context, filter entity.SwapTmTransferHistoriesFilter) ([]*entity.SwapTmTransferHistories, error) {
	var pairs []*entity.SwapTmTransferHistories

	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)
	options.SetSort(bson.D{{"timestamp", -1}})

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_TOKEN_TRANSFER_HISTORY).Find(ctx, r.parseTmTransferHistories(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		pair := &entity.SwapTmTransferHistories{}
		err = cursor.Decode(pair)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}
