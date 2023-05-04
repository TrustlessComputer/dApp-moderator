package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapPairSync(ctx context.Context, filter entity.SwapPairSyncFilter) (*entity.SwapPairSync, error) {
	var swapPairSync entity.SwapPairSync
	err := r.DB.Collection(utils.COLLECTION_SWAP_PAIR_SYNC).FindOne(ctx, r.parseSwapPairSyncFilter(filter)).Decode(&swapPairSync)
	if err != nil {
		return nil, err
	}
	return &swapPairSync, nil
}

func (r *Repository) parseSwapPairSyncFilter(filter entity.SwapPairSyncFilter) bson.M {
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

func (r *Repository) FindSwapPairSyncs(ctx context.Context, filter entity.SwapPairSyncFilter) ([]*entity.SwapPairSync, error) {
	var pairs []*entity.SwapPairSync

	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_PAIR_SYNC).Find(ctx, r.parseSwapPairSyncFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		pair := &entity.SwapPairSync{}
		err = cursor.Decode(pair)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

func (r *Repository) UpdateSwapPairSync(ctx context.Context, sync *entity.SwapPairSync) error {
	collectionName := sync.CollectionName()
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"tx_hash": sync.TxHash, "contract_address": sync.ContractAddress}, bson.M{"$set": sync})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
