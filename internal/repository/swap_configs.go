package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) FindSwapConfig(ctx context.Context, filter entity.SwapConfigsFilter) (*entity.SwapConfigs, error) {
	var swapConfigs entity.SwapConfigs
	err := r.DB.Collection(utils.COLLECTION_SWAP_CONFIGS).FindOne(ctx, r.parseSwapConfigFilter(filter)).Decode(&swapConfigs)
	if err != nil {
		return nil, err
	}
	return &swapConfigs, nil
}

func (r *Repository) parseSwapConfigFilter(filter entity.SwapConfigsFilter) bson.M {
	andCond := make([]bson.M, 0)
	// Define your OR query
	if filter.Name != "" {
		andCond = append(andCond, bson.M{"name": filter.Name})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) UpdateSwapConfig(ctx context.Context, cf *entity.SwapConfigs) error {
	collectionName := cf.CollectionName()
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"name": cf.Name}, bson.M{"$set": cf})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
