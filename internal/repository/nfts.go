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
	if filter.IsBuyable != nil {
		match = append(match, bson.E{"buyable", *filter.IsBuyable})
	}

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

	if filter.Price != nil {
		//btcRate := u.GetExternalRate(os.Getenv("WBTC_ADDRESS"))
		//ethRate := u.GetExternalRate(os.Getenv("WETH_ADDRESS"))
		//rate := btcRate / ethRate

		minPrice := filter.Price.Min
		maxPrice := filter.Price.Max

		//minPriceEth := minPrice * rate
		//maxPriceEth := maxPrice * rate

		fPrice := bson.A{
			bson.D{
				{"$and",
					bson.A{
						//bson.D{{"erc20", strings.ToLower(os.Getenv("WBTC_ADDRESS"))}},
						bson.D{{"price", bson.D{{"$gte", minPrice}}}},
						bson.D{{"price", bson.D{{"$lte", maxPrice}}}},
					},
				},
			},
			//bson.D{
			//	{"$and",
			//		bson.A{
			//			//bson.D{{"erc20", strings.ToLower(os.Getenv("WETH_ADDRESS"))}},
			//			bson.D{{"price", bson.D{{"$gte", minPriceEth}}}},
			//			bson.D{{"price", bson.D{{"$lte", maxPriceEth}}}},
			//		},
			//	},
			//},
		}

		match = append(match, bson.E{"$or", fPrice})

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
		bson.D{
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
		},
		bson.D{
			{"$sort",
				bson.D{
					{"buyable", -1},
					{"price", 1},
					{filter.SortBy, filter.Sort},
				},
			},
		},
	}

	fP := append(f, f1...)
	fP = append(fP, bson.D{{"$skip", filter.Offset}})
	fP = append(fP, bson.D{{"$limit", filter.Limit}})

	resp := []*entity.MkpNftsResp{}
	cursor, err := r.DB.Collection(utils.COLLECTION_NFTS).Aggregate(context.TODO(), fP)
	if err != nil {
		return nil, err
	}

	if err = cursor.All((context.TODO()), &resp); err != nil {
		return nil, err
	}

	pResp := &entity.MkpNftsPagination{}

	count, err := r.DB.Collection(utils.COLLECTION_NFTS).CountDocuments(context.TODO(), match)
	if err != nil {
		return nil, err
	}

	for index, item := range resp {
		if len(item.BnsDefault) > 0 && item.BnsDefault[0].Resolver != "" {
			for j, bnsItem := range resp[index].BnsData {
				if bnsItem.ID.Hex() == item.BnsDefault[0].BNSDefaultID.Hex() {
					resp[index].BnsData[0], resp[index].BnsData[j] = resp[index].BnsData[j], resp[index].BnsData[0]
					break
				}
			}
		}
	}
	pResp.Items = resp
	pResp.TotalItem = count
	return pResp, nil
}
