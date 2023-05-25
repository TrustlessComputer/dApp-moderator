package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) InsertListing(obj *entity.MarketplaceListings) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CancelListing(ctx context.Context, offeringID string) error {
	filter := bson.M{
		"offering_id": offeringID,
	}

	update := bson.M{
		"status": entity.MarketPlaceCancel,
	}

	result, err := r.DB.Collection(entity.MarketplaceListings{}.CollectionName()).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("No document")
	}

	return nil
}

func (r *Repository) PurchaseListing(ctx context.Context, offeringID string) error {
	filter := bson.M{
		"offering_id": offeringID,
	}

	update := bson.M{
		"status": entity.MarketPlaceDone,
	}

	result, err := r.DB.Collection(entity.MarketplaceListings{}.CollectionName()).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("No document")
	}

	return nil
}

func (r Repository) FilterMarketplaceListings(filter entity.FilterMarketplaceListings) ([]entity.MarketplaceListings, error) {
	match := bson.D{}

	if filter.CollectionContract != nil && *filter.CollectionContract != "" {
		match = append(match, bson.E{"collection_contract", *filter.CollectionContract})
	}
	if filter.TokenId != nil && *filter.TokenId != "" {
		match = append(match, bson.E{"token_id", *filter.TokenId})
	}
	if filter.Status != nil {
		match = append(match, bson.E{"status", *filter.Status})
	}
	if filter.SellerAddress != nil && *filter.SellerAddress != "" {
		match = append(match, bson.E{"seller", *filter.SellerAddress})
	}
	if filter.Erc20Token != nil && *filter.Erc20Token != "" {
		match = append(match, bson.E{"erc_20_token", *filter.Erc20Token})
	}

	mkpListing := []entity.MarketplaceListings{}
	f := bson.A{
		bson.D{
			{"$match", match},
		},
		bson.D{{"$sort", bson.D{{"block_number", -1}}}},
		bson.D{{"$skip", filter.Offset}},
		bson.D{{"$limit", filter.Limit}},
	}

	cursor, err := r.DB.Collection(entity.MarketplaceListings{}.CollectionName()).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &mkpListing)
	if err != nil {
		return nil, err
	}

	return mkpListing, nil
}
