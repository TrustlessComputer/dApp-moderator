package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) InsertActivity(obj *entity.MarketplaceTokenActivity) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) FilterTokenActivites(filter entity.FilterTokenActivities) ([]entity.MarketplaceTokenActivity, error) {
	match := bson.D{}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		match = append(match, bson.E{"collection_contract", *filter.ContractAddress})
	}
	if filter.TokenID != nil && *filter.TokenID != "" {
		match = append(match, bson.E{"inscription_id", *filter.TokenID})
	}

	mkpListing := []entity.MarketplaceTokenActivity{}
	f := bson.A{
		bson.D{
			{"$match", match},
		},
		bson.D{{"$sort", bson.D{{"block_number", -1}, {"log_index", 1}}}},
		bson.D{{"$skip", filter.Offset}},
		bson.D{{"$limit", filter.Limit}},
	}

	cursor, err := r.DB.Collection(entity.MarketplaceTokenActivity{}.CollectionName()).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &mkpListing)
	if err != nil {
		return nil, err
	}

	return mkpListing, nil
}
