package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapSlackReport(ctx context.Context) (*entity.SwapPairSlackReport, error) {
	var swapSlackReport entity.SwapPairSlackReport
	err := r.DB.Collection(utils.COLLECTION_SWAP_REPOR_SLACK).FindOne(ctx, entity.SwapPairFilter{}).Decode(&swapSlackReport)
	if err != nil {
		return nil, err
	}
	return &swapSlackReport, nil
}

func (r *Repository) FindSwapPair(ctx context.Context, filter entity.SwapPairFilter) (*entity.SwapPair, error) {
	var swapPair entity.SwapPair
	err := r.DB.Collection(utils.COLLECTION_SWAP_PAIR).FindOne(ctx, r.parseSwapPairFilter(filter)).Decode(&swapPair)
	if err != nil {
		return nil, err
	}
	return &swapPair, nil
}

func (r *Repository) parseSwapPairFilter(filter entity.SwapPairFilter) bson.M {
	andCond := make([]bson.M, 0)
	// Define your OR query
	if filter.Pair != "" {
		andCond = append(andCond, bson.M{"pair": filter.Pair})
	}
	if filter.TxHash != "" {
		andCond = append(andCond, bson.M{"tx_hash": filter.TxHash})
	}
	if filter.Token != "" {
		andCond = append(andCond, bson.M{"$or": []bson.M{
			{"token0": filter.Token},
			{"token1": filter.Token},
		}})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindSwapPairs(ctx context.Context, filter entity.SwapPairFilter) ([]entity.SwapPair, error) {
	var pairs []entity.SwapPair

	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_PAIR).Find(ctx, r.parseSwapPairFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pair entity.SwapPair
		err = cursor.Decode(&pair)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

func (r *Repository) UpdateSwapPair(ctx context.Context, pair *entity.SwapPair) error {
	collectionName := pair.CollectionName()
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"pair": pair.Pair}, bson.M{"$set": pair})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *Repository) FindTokensInPoolByContracts(ctx context.Context, contracts []string, filter entity.TokenFilter) ([]*entity.Token, error) {
	tokens := []*entity.Token{}
	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	// Set the options for the query
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)
	f := bson.D{{"address", bson.M{"$in": contracts}}}
	cursor, err := r.DB.Collection(utils.COLLECTION_TOKENS).Find(ctx, f, options)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		token := &entity.Token{}
		err = cursor.Decode(token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
