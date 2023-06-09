package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func (r *Repository) UpdateBnsResolver(tokenID string, resolver string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"resolver": strings.ToLower(resolver)}}
	updated, err := r.UpdateOne(utils.COLLECTION_BNS, f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r *Repository) UpdateBnsPfp(tokenID string, pfp string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"pfp": strings.ToLower(pfp)}}
	updated, err := r.UpdateOne(utils.COLLECTION_BNS, f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r *Repository) FilterBNS(filter entity.FilterBns) ([]*entity.FilteredBNS, error) {
	resp := []*entity.FilteredBNS{}
	f := bson.A{}
	match := bson.D{}
	if filter.Resolver != nil && *filter.Resolver != "" {
		match = append(match, bson.E{"resolver", strings.ToLower(*filter.Resolver)})
	}

	if filter.PFP != nil && *filter.PFP != "" {
		match = append(match, bson.E{"pfp", strings.ToLower(*filter.PFP)})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		match = append(match, bson.E{"owner", strings.ToLower(*filter.Owner)})
	}

	if filter.Name != nil && *filter.Name != "" {
		match = append(match, bson.E{"name", primitive.Regex{Pattern: *filter.Name, Options: "i"}})
	}

	if filter.TokenID != nil && *filter.TokenID != "" {
		match = append(match, bson.E{"token_id", *filter.TokenID})
	}

	if len(match) > 0 {
		f = append(f, bson.D{{"$match", match}})
	}

	f = append(f, bson.D{{"$sort", bson.D{{"token_id_int", entity.SORT_DESC}}}})
	f = append(f, bson.D{{"$skip", filter.Offset}})
	f = append(f, bson.D{{"$limit", filter.Limit}})

	cursor, err := r.DB.Collection(utils.VIEW_BNS).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) FindBNS(tokenID string) (*entity.FilteredBNS, error) {
	resp := &entity.FilteredBNS{}
	f := bson.D{{"token_id", tokenID}}

	cursor := r.DB.Collection(utils.VIEW_BNS).FindOne(context.TODO(), f)

	err := cursor.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
