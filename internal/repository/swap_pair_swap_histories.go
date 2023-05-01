package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapPairSwapHistory(ctx context.Context, filter entity.SwapPairSwapHistoriesFilter) (*entity.SwapPairSwapHistories, error) {
	var swapPairHistories entity.SwapPairSwapHistories
	err := r.DB.Collection(utils.COLLECTION_SWAP_HISTORIES).FindOne(ctx, r.parseSwapPairSwapHistories(filter)).Decode(&swapPairHistories)
	if err != nil {
		return nil, err
	}
	return &swapPairHistories, nil
}

func (r *Repository) parseSwapPairSwapHistories(filter entity.SwapPairSwapHistoriesFilter) bson.M {
	andCond := make([]bson.M, 0)
	// Define your OR query
	if filter.ContractAddress != "" {
		andCond = append(andCond, bson.M{"contract_address": filter.ContractAddress})
	}
	if filter.TxHash != "" {
		andCond = append(andCond, bson.M{"tx_hash": filter.TxHash})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindTokenReport(ctx context.Context, filter entity.SwapPairFilter) ([]*entity.SwapPairReport, error) {
	var tokens []*entity.SwapPairReport

	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	// Set the options for the query
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_REPORT_VOLUME).Find(ctx, r.parseSwapPairFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var token *entity.SwapPairReport
		err = cursor.Decode(&token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
