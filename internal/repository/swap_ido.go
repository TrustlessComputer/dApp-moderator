package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"log"
	"strings"
	"time"

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

func (r *Repository) FindSwapIdoView(ctx context.Context, filter entity.SwapIdoFilter) (*entity.SwapIdo, error) {
	var swapIdo entity.SwapIdo
	err := r.DB.Collection(utils.COLLECTION_SWAP_IDO_LIST_VIEW).FindOne(ctx, r.parseSwapIdoFilter(filter)).Decode(&swapIdo)
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

	if filter.Status != "" {
		andCond = append(andCond, bson.M{"status": filter.Status})
	}

	if filter.WalletAddress != "" {
		andCond = append(andCond, bson.M{"user_wallet_address": strings.ToLower(filter.WalletAddress)})
	}

	if filter.CheckStartTime > 0 {
		andCond = append(andCond, bson.M{"start_at": bson.M{"$gte": time.Now()}})
	} else if filter.CheckStartTime < 0 {
		andCond = append(andCond, bson.M{"start_at": bson.M{"$lte": time.Now()}})
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
	options.SetSort(bson.D{{"start_at", 1}})

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

func (r *Repository) FindSwapIdosView(ctx context.Context, filter entity.SwapIdoFilter) ([]*entity.SwapIdo, error) {
	idos := []*entity.SwapIdo{}
	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)
	options.SetSort(bson.D{{"start_at", 1}})

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_IDO_LIST_VIEW).Find(ctx, r.parseSwapIdoFilter(filter), options)
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

func (r *Repository) FindIdoTokens(ctx context.Context, filter entity.IdoTokenFilter) ([]*entity.Token, error) {
	tokens := []*entity.Token{}
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.VIEW_SWAP_IDO_TOKEN).Find(ctx, r.parseIdoTokenFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var token entity.Token
		err = cursor.Decode(&token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, &token)
	}
	return tokens, nil
}

func (r *Repository) parseIdoTokenFilter(filter entity.IdoTokenFilter) bson.M {

	andCond := make([]bson.M, 0)

	if filter.CreatedBy != "" {
		andCond = append(andCond, bson.M{"owner": filter.CreatedBy})
	}

	if len(filter.Address) > 0 {
		andCond = append(andCond, bson.M{"token.address": bson.M{"$nin": filter.Address}})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}
