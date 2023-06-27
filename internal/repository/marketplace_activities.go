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

	if len(filter.Types) > 0 {
		match = append(match, bson.E{"type", bson.D{{"$in", filter.Types}}})
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

func (r Repository) FilterSoulHistories(filter entity.FilterTokenActivities) ([]*entity.SoulTokenHistoriesFiltered, error) {
	match := bson.D{}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		match = append(match, bson.E{"collection_contract", strings.ToLower(*filter.ContractAddress)})
	}
	if filter.TokenID != nil && *filter.TokenID != "" {
		match = append(match, bson.E{"inscription_id", strings.ToLower(*filter.TokenID)})
	}

	match = append(match, bson.E{"type", bson.D{{"$in", []entity.TokenActivityType{entity.SoulUnlockFeature}}}})

	mkpListing := []*entity.SoulTokenHistoriesFiltered{}
	f := bson.A{
		bson.D{
			{"$match", match},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "soul_image_histories"},
					{"localField", "tx_hash"},
					{"foreignField", "tx_hash"},
					{"let", bson.D{{"log_index", "$log_index"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$eq",
													bson.A{
														"$log_index",
														"$$log_index",
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{"as", "soul_image_histories"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$soul_image_histories"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"image_capture", "$soul_image_histories.image_capture"},
					{"thumbnail", "$soul_image_histories.image_capture"},
					{"feature_name", "$soul_image_histories.feature_name"},
					{"balance", "$soul_image_histories.erc_20_amount"},
					{"erc_20_address", "$soul_image_histories.erc_20_address"},
					{"token_id", "$inscription_id"},
					{"amount",
						bson.D{
							{"$divide",
								bson.A{
									bson.D{{"$toDouble", "$soul_image_histories.erc_20_amount"}},
									bson.D{
										{"$pow",
											bson.A{
												10,
												18,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		bson.D{{"$project", bson.D{{"soul_image_histories", 0}}}},
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
