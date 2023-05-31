package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

func (r *Repository) InsertMarketPlaceAggregation(obj *entity.MarketplaceCollectionAggregation) error {
	f := bson.D{{
		"contract", strings.ToLower(obj.Contract),
	}}

	data := bson.M{}
	err := helpers.Transform(obj, &data)
	if err != nil {
		return err
	}

	updateOpts := options.Update().SetUpsert(true)
	_, err = r.UpdateOneWithOptions(obj.CollectionName(), f, bson.M{"$set": data}, updateOpts)
	if err != nil {
		return err
	}

	return nil
}

// Aggregate data for view
func (r *Repository) AggregatetMarketPlaceData(filter entity.FilterMarketplaceAggregationData) ([]entity.MarketplaceCollectionAggregation, error) {
	match := bson.D{}
	if filter.CollectionContract != nil && *filter.CollectionContract != "" {
		match = append(match, bson.E{"contract", strings.ToLower(*filter.CollectionContract)})
	}

	if filter.Limit == 0 {
		filter.Limit = 100
	}

	f := bson.A{

		bson.D{
			{"$lookup",
				bson.D{
					{"from", "nfts"},
					{"localField", "contract"},
					{"foreignField", "collection_address"},
					{"let", bson.D{{"owner", "$onwer"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$group",
									bson.D{
										{"_id", bson.D{{"owner", "$owner"}}},
										{"items", bson.D{{"$sum", 1}}},
									},
								},
							},
							bson.D{
								{"$project",
									bson.D{
										{"items", 1},
										{"owner_address", "$_id.owner"},
									},
								},
							},
						},
					},
					{"as", "nft_owners"},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "marketplace_listings"},
					{"localField", "contract"},
					{"foreignField", "collection_contract"},
					{"pipeline",
						bson.A{
							bson.D{{"$match", bson.D{{"status", 2}}}},
							bson.D{
								{"$group",
									bson.D{
										{"_id",
											bson.D{
												{"contract", "$collection_contract"},
												{"erc_20_token", "$erc_20_token"},
											},
										},
										{"total_volume", bson.D{{"$sum", bson.D{{"$toDouble", "$price"}}}}},
										{"total_sales", bson.D{{"$sum", 1}}},
									},
								},
							},
							bson.D{
								{"$addFields",
									bson.D{
										{"erc_20_token", "$_id.erc_20_token"},
										{"contract", "$_id.contract"},
										{"marketplace_type", "marketplace_listings"},
									},
								},
							},
							bson.D{{"$project", bson.D{{"_id", 0}}}},
						},
					},
					{"as", "marketplace_listings"},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "marketplace_offers"},
					{"localField", "contract"},
					{"foreignField", "collection_contract"},
					{"pipeline",
						bson.A{
							bson.D{{"$match", bson.D{{"status", 2}}}},
							bson.D{
								{"$group",
									bson.D{
										{"_id",
											bson.D{
												{"contract", "$collection_contract"},
												{"erc_20_token", "$erc_20_token"},
											},
										},
										{"total_volume", bson.D{{"$sum", bson.D{{"$toDouble", "$price"}}}}},
										{"total_sales", bson.D{{"$sum", 1}}},
									},
								},
							},
							bson.D{
								{"$addFields",
									bson.D{
										{"erc_20_token", "$_id.erc_20_token"},
										{"contract", "$_id.contract"},
										{"marketplace_type", "marketplace_offers"},
									},
								},
							},
							bson.D{{"$project", bson.D{{"_id", 0}}}},
						},
					},
					{"as", "marketplace_offers"},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"market_place_volumes",
						bson.D{
							{"$concatArrays",
								bson.A{
									"$marketplace_listings",
									"$marketplace_offers",
								},
							},
						},
					},
					{"unique_owners", bson.D{{"$size", "$nft_owners"}}},
					{"total_owners", "$unique_onwers"},
					{"total_nfts", "$total_items"},
					{"total_sales",
						bson.D{
							{"$sum",
								bson.A{
									bson.D{{"$size", "$marketplace_listings"}},
									bson.D{{"$size", "$marketplace_offers"}},
								},
							},
						},
					},
					{"floor_price", 0},
					{"volume", 0},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"nft_owners", 0},
					{"marketplace_listings", 0},
					{"marketplace_offers", 0},
				},
			},
		},
	}

	f = append(f, bson.D{{"$skip", filter.Offset}})
	f = append(f, bson.D{{"$limit", filter.Limit}})
	f = append(f, bson.D{{"$sort", bson.D{
		{"total_sales", -1},
	}}})

	resp := []entity.MarketplaceCollectionAggregation{}

	cursor, err := r.DB.Collection(entity.Collections{}.CollectionName()).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) FilterMarketPlaceCollectionAggregation(filter entity.FilterMarketplaceAggregationData) ([]entity.MarketplaceCollectionAggregation, error) {
	match := bson.D{}
	if filter.CollectionContract != nil && *filter.CollectionContract != "" {
		match = append(match, bson.E{"contract", strings.ToLower(*filter.CollectionContract)})
	}

	if filter.Limit == 0 {
		filter.Limit = 100
	}

	f :=
		bson.A{
			bson.D{
				{"$match", match},
			},
			bson.D{{"$skip", filter.Offset}},
			bson.D{{"$limit", filter.Limit}},
			bson.D{{"$sort", bson.D{
				{"total_sales", -1},
			}}},
		}

	resp := []entity.MarketplaceCollectionAggregation{}

	cursor, err := r.DB.Collection(entity.Collections{}.CollectionName()).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
