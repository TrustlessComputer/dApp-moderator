package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func (r *Repository) InsertActivity(obj *entity.MarketplaceTokenActivity) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) FilterTokenActivites(filter entity.FilterTokenActivities) ([]*entity.MarketplaceTokenActivity, error) {
	match := bson.D{}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		match = append(match, bson.E{"collection_contract", strings.ToLower(*filter.ContractAddress)})
	}
	if filter.TokenID != nil && *filter.TokenID != "" {
		match = append(match, bson.E{"inscription_id", strings.ToLower(*filter.TokenID)})
	}

	mkpListing := []*entity.MarketplaceTokenActivity{}
	f := bson.A{
		bson.D{
			{"$match", match},
		},
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

	for _, ac := range mkpListing {
		ac.AmountStr = fmt.Sprintf("%d", ac.Amount)
		ac.TokenID = ac.InscriptionID
		ac.Thumbnail = fmt.Sprintf("https://dapp.trustless.computer/dapp/api/nft-explorer/collections/%s/nfts/%s/content", ac.CollectionContract, ac.InscriptionID)
	}

	return mkpListing, nil
}

func (r Repository) PurchaseMKPActivity(offeringID string) (*entity.MarketplaceTokenActivity, error) {
	match := bson.D{}
	match = append(match, bson.E{"offering_id", strings.ToLower(offeringID)})
	match = append(match, bson.E{"type", entity.TokenPurchase})

	mkpListing := &entity.MarketplaceTokenActivity{}

	cursor := r.DB.Collection(entity.MarketplaceTokenActivity{}.CollectionName()).FindOne(context.TODO(), match, nil)

	err := cursor.Decode(&mkpListing)
	if err != nil {
		return nil, err
	}

	return mkpListing, nil
}
