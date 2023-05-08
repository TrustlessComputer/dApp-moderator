package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapIdo(ctx context.Context, filter entity.SwapIdoFilter) (*entity.SwapIdo, error) {
	var swapIdo entity.SwapIdo
	err := r.DB.Collection(utils.COLLECTION_SWAP_IDO).FindOne(ctx, r.parseSwapIdoFilter(filter)).Decode(&swapIdo)
	if err != nil {
		return nil, err
	}
	return &swapIdo, nil
}

func (r *Repository) parseSwapIdoFilter(filter entity.SwapIdoFilter) bson.M {
	andCond := make([]bson.M, 0)
	if filter.ID != "" {
		objectId, err := primitive.ObjectIDFromHex(filter.ID)
		if err != nil {
			log.Println("Invalid id")
		}
		andCond = append(andCond, bson.M{"_id": objectId})
	}

	if filter.Address != "" {
		andCond = append(andCond, bson.M{"token.address": filter.Address})
	}

	if filter.WalletAddress != "" {
		andCond = append(andCond, bson.M{"user_wallet_address": filter.WalletAddress})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindSwapIdos(ctx context.Context, filter entity.SwapIdoFilter) ([]*entity.SwapIdo, error) {
	idos := []*entity.SwapIdo{}
	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_IDO).Find(ctx, r.parseSwapIdoFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pair entity.SwapIdo
		err = cursor.Decode(&pair)
		if err != nil {
			return nil, err
		}
		idos = append(idos, &pair)
	}
	return idos, nil
}

func (r *Repository) UpdateSwapIdo(ctx context.Context, pair *entity.SwapIdo) error {
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

func (r *Repository) DetelteSwapIdo(ctx context.Context, filter entity.SwapIdoFilter) error {
	result, err := r.DB.Collection(utils.COLLECTION_SWAP_IDO).DeleteOne(ctx, r.parseSwapIdoFilter(filter))
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
