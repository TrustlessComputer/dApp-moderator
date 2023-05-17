package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapBotConfig(ctx context.Context, filter entity.SwapBotConfigFilter) (*entity.SwapBotConfig, error) {
	var swapIdo entity.SwapBotConfig
	err := r.DB.Collection(utils.COLLECTION_SWAP_BOT_CONFIG).FindOne(ctx, r.parseSwapBotFilter(filter)).Decode(&swapIdo)
	if err != nil {
		return nil, err
	}
	return &swapIdo, nil
}

func (r *Repository) parseSwapBotFilter(filter entity.SwapBotConfigFilter) bson.M {
	andCond := make([]bson.M, 0)
	if filter.Address != "" {
		andCond = append(andCond, bson.M{"address": filter.Address})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindSwapBotConfigs(ctx context.Context, filter entity.SwapBotConfigFilter) ([]*entity.SwapBotConfig, error) {
	idos := []*entity.SwapBotConfig{}
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_BOT_CONFIG).Find(ctx, r.parseSwapBotFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pair entity.SwapBotConfig
		err = cursor.Decode(&pair)
		if err != nil {
			return nil, err
		}
		idos = append(idos, &pair)
	}
	return idos, nil
}

func (r *Repository) UpdateSwapBotConfig(ctx context.Context, pair *entity.SwapBotConfig) error {
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
