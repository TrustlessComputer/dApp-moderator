package repository

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) CollectionsByNfts(ownerAddress string) ([]entity.GroupedCollection, error) {
	f2 := bson.A{
		bson.D{{"$match", bson.D{{"owner", primitive.Regex{Pattern: ownerAddress, Options: "i"}}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", bson.D{{"collection_address", "$collection_address"}}},
					{"tokens", bson.D{{"$sum", 1}}},
				},
			},
		},
		bson.D{{"$project", bson.D{{"tokens", 1}}}},
	}

	groupedNfts := []entity.GroupedCollection{}
	cursor, err := r.DB.Collection(entity.Nfts{}.CollectionName()).Aggregate(context.TODO(), f2)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cursor.All((context.TODO()), &groupedNfts); err != nil {
		return nil, errors.WithStack(err)
	}

	return groupedNfts, nil
}

func (r *Repository) CreateNftHistories(histories *entity.NftHistories) (*mongo.InsertOneResult, error) {
	inserted, err := r.InsertOne(histories)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) GetNft(contract string, tokenID string) (*entity.Nfts, error) {
	nftResp, err := r.FindOne(entity.Nfts{}.CollectionName(), bson.D{
		{"collection_address", contract},
		{"token_id", tokenID},
	})

	if err != nil {
		return nil, err
	}

	nft := &entity.Nfts{}
	err = nftResp.Decode(nft)

	if err != nil {
		return nil, err
	}

	return nft, nil

}

func (r *Repository) UpdateNftOwner(contract string, tokenID string, owner string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"collection_address", contract},
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"owner": strings.ToLower(owner)}}

	updated, err := r.UpdateOne(entity.Nfts{}.CollectionName(), f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil

}

func (r *Repository) GetNfts(collectionAddress string, skip int, limit int) ([]entity.Nfts, error) {
	f2 := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"collection_address", strings.ToLower(collectionAddress)},
					{"image", bson.D{{"$ne", ""}}},
				},
			},
		},

		bson.D{{"$skip", skip}},
		bson.D{{"$limit", limit}},
	}

	groupedNfts := []entity.Nfts{}
	cursor, err := r.DB.Collection(entity.Nfts{}.CollectionName()).Aggregate(context.TODO(), f2)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cursor.All((context.TODO()), &groupedNfts); err != nil {
		return nil, errors.WithStack(err)
	}

	return groupedNfts, nil
}

func (r *Repository) RefreshNft(contract string, tokenID string, metadataType string, contentType string, attributes []nft_explorer.NftAttr, mintedAt float64, metadata interface{}) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"collection_address", contract},
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{
		"content_type":  strings.ToLower(contentType),
		"metadata_type": strings.ToLower(metadataType),
		"minted_at":     mintedAt,
		"attributes":    attributes,
		"metadata":      metadata,
	}}

	updated, err := r.UpdateOne(entity.Nfts{}.CollectionName(), f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil

}

func (r *Repository) GetMarketplaceListing(offeringID string) (*entity.MarketplaceListings, error) {
	nftResp, err := r.FindOne(entity.MarketplaceListings{}.CollectionName(), bson.D{
		{"offering_id", strings.ToLower(offeringID)},
	})

	if err != nil {
		return nil, err
	}

	ml := &entity.MarketplaceListings{}
	err = nftResp.Decode(ml)

	if err != nil {
		return nil, err
	}

	return ml, nil

}

func (r *Repository) SoulNfts(contract string) ([]entity.Nfts, error) {
	result := []entity.Nfts{}

	f := bson.A{
		bson.D{
			{"$match", bson.D{
				{"collection_address", strings.ToLower(contract)},
			}},
		},
	}

	cursor, err := r.DB.Collection(entity.Nfts{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) UpdateNftSize(contract string, tokenID string, size int64) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"collection_address", contract},
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"size": size}}

	updated, err := r.UpdateOne(entity.Nfts{}.CollectionName(), f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil

}

func (r *Repository) GetNftsWithoutSize(collectionAddress string, skip int, limit int) ([]entity.Nfts, error) {
	f2 := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"collection_address", strings.ToLower(collectionAddress)},
					{"image", bson.D{{"$ne", ""}}},
					{"size", 0},
				},
			},
		},

		bson.D{{"$skip", skip}},
		bson.D{{"$limit", limit}},
	}

	groupedNfts := []entity.Nfts{}
	cursor, err := r.DB.Collection(utils.VIEW_NFTS_WITH_SIZE).Aggregate(context.TODO(), f2)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cursor.All((context.TODO()), &groupedNfts); err != nil {
		return nil, errors.WithStack(err)
	}

	return groupedNfts, nil
}

func (r *Repository) FilterMKPNfts(filter entity.FilterNfts) (*entity.MkpNftsPagination, error) {

	f := bson.A{}
	match := bson.D{}

	if filter.TokenID != nil && *filter.TokenID != "" {
		match = append(match, bson.E{"token_id", *filter.TokenID})
	}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		match = append(match, bson.E{"collection_address", *filter.ContractAddress})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		match = append(match, bson.E{"owner", strings.ToLower(*filter.Owner)})
	}

	if filter.Rarity != nil {
		filter.Rarity.Min = filter.Rarity.Min / 100
		filter.Rarity.Max = filter.Rarity.Max / 100
		//f = append(f, bson.E{"$and", bson.A{
		//	bson.E{"attributes.percent", bson.M{"$lte": filter.Rarity.Max / 100}},
		//	bson.E{"attributes.percent", bson.M{"$gte": filter.Rarity.Min / 100}},
		//}})

		attrs, err := r.FilterCollectionAttributeByPercent(entity.FilterMarketplaceCollectionAttribute{
			ContractAddress: filter.ContractAddress,
			MaxPercent:      &filter.Rarity.Max,
			MinPercent:      &filter.Rarity.Min,
		})

		if err != nil {
			return nil, err
		}

		key := []string{}
		value := []string{}
		for _, attr := range attrs {
			key = append(key, attr.TraitType)
			value = append(value, attr.Value)
		}

		filter.AttrKey = key
		filter.AttrValue = value
	}

	if len(filter.AttrKey) > 0 {
		match = append(match, bson.E{"attributes.trait_type", bson.M{"$in": filter.AttrKey}})
	}

	if len(filter.AttrValue) > 0 {
		match = append(match, bson.E{"attributes.value", bson.M{"$in": filter.AttrValue}})
	}

	if len(match) > 0 {
		f = append(f, bson.D{{"$match", match}})
	}

	f1 := bson.A{
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "collections"},
					{"localField", "collection_address"},
					{"foreignField", "contract"},
					{"as", "collection"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$collection"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "marketplace_listings"},
					{"localField", "token_id"},
					{"foreignField", "token_id"},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"collection_contract", strings.ToLower(*filter.ContractAddress)},
										{"status", 0},
									},
								},
							},
							bson.D{{"$skip", 0}},
							bson.D{{"$limit", 1}},
						},
					},
					{"as", "listing_for_sales"},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "marketplace_offers"},
					{"localField", "token_id"},
					{"foreignField", "token_id"},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"collection_contract", strings.ToLower(*filter.ContractAddress)},
										{"status", 0},
									},
								},
							},
							bson.D{{"$skip", 0}},
							bson.D{{"$limit", 100}},
						},
					},
					{"as", "make_offers"},
				},
			},
		},
		bson.D{{"$addFields", bson.D{{"price_erc20", "$listing_for_sales"}}}},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$price_erc20"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"buyable",
						bson.D{
							{"$cond",
								bson.A{
									bson.D{
										{"$or",
											bson.A{
												bson.D{
													{"$eq",
														bson.A{
															bson.D{
																{"$ifNull",
																	bson.A{
																		"$price_erc20",
																		0,
																	},
																},
															},
															0,
														},
													},
												},
											},
										},
									},
									false,
									true,
								},
							},
						},
					},
					{"price",
						bson.D{
							{"$cond",
								bson.A{
									bson.D{
										{"$or",
											bson.A{
												bson.D{
													{"$eq",
														bson.A{
															bson.D{
																{"$ifNull",
																	bson.A{
																		"$price_erc20",
																		0,
																	},
																},
															},
															0,
														},
													},
												},
											},
										},
									},
									0,
									bson.D{{"$toDouble", "$price_erc20.price"}},
								},
							},
						},
					},
					{"erc20", "$price_erc20.erc_20_token"},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "bns"},
					{"localField", "owner"},
					{"foreignField", "resolver"},
					{"pipeline",
						bson.A{
							bson.D{{"$skip", 0}},
							bson.D{{"$limit", 1}},
						},
					},
					{"as", "bns_data"},
				},
			},
		},
	}

	f1 = append(f1, bson.D{
		{"$lookup",
			bson.D{
				{"from", "bns_default"},
				{"localField", "owner"},
				{"foreignField", "resolver"},
				{"pipeline",
					bson.A{
						bson.D{{"$skip", 0}},
						bson.D{{"$limit", 1}},
					},
				},
				{"as", "bns_default"},
			},
		},
	})
	f1 = append(f1, bson.D{
		{"$sort",
			bson.D{
				{"buyable", -1},
				{"price", 1},
				{filter.SortBy, filter.Sort},
			},
		},
	})

	matchPrice := bson.D{}
	if filter.IsBuyable != nil {
		matchPrice = append(matchPrice, bson.E{"buyable", *filter.IsBuyable})
	}

	if filter.Price != nil {
		//matchPrice = append(matchPrice, bson.E{"buyable", true})
		matchPrice = append(matchPrice, bson.E{"$and", bson.A{
			bson.D{{"price", bson.D{{"$lte", filter.Price.Max}}}},
			bson.D{{"price", bson.D{{"$gte", filter.Price.Min}}}},
		}})
	}

	if len(matchPrice) > 0 {
		f1 = append(f1, bson.D{{"$match", matchPrice}})
	}

	fPagination := append(f, f1...)

	//count all items
	fCount := fPagination
	fCount = append(fCount, bson.D{
		{"$group",
			bson.D{
				{"_id", bson.D{{"collection_address", "$collection_address"}}},
				{"all", bson.D{{"$sum", 1}}},
			},
		},
	})

	fPagination = append(fPagination, bson.D{{"$skip", filter.Offset}})
	fPagination = append(fPagination, bson.D{{"$limit", filter.Limit}})

	fAll := bson.A{
		bson.D{
			{"$facet",
				bson.D{
					{"items",
						fPagination, //filter with pagination
					},
					{"count",
						fCount, // count all items
					},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$count"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{{"$addFields", bson.D{{"total_item", "$count.all"}}}},
		bson.D{{"$project", bson.D{{"count", 0}}}},
	}

	pResp := []entity.MkpNftsPagination{}
	cursor, err := r.DB.Collection(utils.COLLECTION_NFTS).Aggregate(context.TODO(), fAll)
	if err != nil {
		return nil, err
	}
	err = cursor.All((context.TODO()), &pResp)
	if err != nil {
		return nil, err
	}

	if len(pResp) == 0 {
		return nil, errors.New("Cannot get nfts")
	}

	return &pResp[0], nil
}
