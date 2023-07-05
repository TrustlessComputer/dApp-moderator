package repository

import (
	"context"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"os"
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

func (r *Repository) getPipelineForAuctionRequest(filter *entity.FilterNfts) bson.A {
	pipeline := bson.A{bson.M{"$lookup": bson.D{
		{"from", "nft_auction_available"},
		{"localField", "token_id"},
		{"foreignField", "token_id"},
		{"let", bson.D{{"contract_address", "$collection_address"}}},
		{"pipeline",
			bson.A{
				bson.D{
					{"$match",
						bson.D{
							{"$expr",
								bson.D{
									{"$eq",
										bson.A{
											"$contract",
											"$$contract_address",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{"as", "nft_auction_available"},
	}},
		bson.M{"$unwind": bson.D{
			{"path", "$nft_auction_available"},
			{"preserveNullAndEmptyArrays", true},
		}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "auction"},
					{"localField", "token_id"},
					{"foreignField", "token_id"},
					{"let", bson.D{{"contract_address", "$collection_address"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$eq",
													bson.A{
														"$collection_address",
														"$$contract_address",
													},
												},
											},
										},
										{"status", bson.M{"$in": bson.A{1, 2}}}, // 1 token ở 1 contract tại 1 thời điểm chỉ có 1 auction diễn ra
									},
								},
							},
							bson.M{"$sort": bson.M{"created_at": -1}},
							bson.M{"$skip": 0},
							bson.M{"$limit": 1},
						},
					},
					{"as", "auction"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$auction"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"db_auction_id",
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
																		"$auction",
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
									"",
									"$auction._id",
								},
							},
						},
					},
					{"start_time_block",
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
																		"$auction",
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
									"",
									"$auction.start_time_block",
								},
							},
						},
					},
					{"end_time_block",
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
																		"$auction",
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
									"",
									"$auction.end_time_block",
								},
							},
						},
					},
					{"auction_id",
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
																		"$auction",
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
									"",
									"$auction.auction_id",
								},
							},
						},
					},
					{"auction_status",
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
																		"$auction",
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
									-1,
									"$auction.status",
								},
							},
						},
					},
					{"is_available_for_auction", "$nft_auction_available.is_auction"},
					{"is_live_auction",
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
																		"$auction",
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
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"auction", 0},
					{"nft_auction_available", 0},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "nfts_attributes_percent_view"},
					{"localField", "token_id"},
					{"foreignField", "token_id"},
					{"let", bson.D{{"contract", "$collection_address"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$eq",
													bson.A{
														"$collection_address",
														"$$contract",
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{"as", "percent_attributes"},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"attributes", "$percent_attributes"},
					{"rarity", bson.D{{"$avg", "$percent_attributes.percent"}}},
				},
			},
		},
		bson.D{{"$project", bson.D{{"percent_attributes", 0}}}},
	}

	if filter.IsOrphan != nil && *filter.IsOrphan > 0 {
		pipeline = append(pipeline, bson.D{
			{"$match", bson.D{{"$or", bson.A{
				bson.M{
					"is_available_for_auction": true,
				}, bson.M{
					"is_live_auction": true,
				}}}}},
		})
	}

	return pipeline
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

	f1 := bson.A{}
	fsort := bson.D{}
	fBNS := bson.D{}
	fBNSDefault := bson.D{}
	fMarketPlaceOffer := bson.D{}
	fName := bson.A{}

	collection := utils.COLLECTION_NFTS
	// Only for soul
	if filter.ContractAddress != nil {
		contractAddress := strings.ToLower(*filter.ContractAddress)
		if contractAddress == strings.ToLower(os.Getenv("SOUL_CONTRACT")) {
			sortDoc := bson.D{}
			if filter.SortBy == "orphanage" {
				sortDoc = append(sortDoc,
					bson.E{"is_available_for_auction", filter.Sort},
					bson.E{"is_live_auction", filter.Sort},
					bson.E{"auction_status", 1},
				)
			} else {
				sortDoc = append(sortDoc, bson.E{filter.SortBy, filter.Sort})
			}
			if filter.SortBy != "token_id_int" {
				sortDoc = append(sortDoc, bson.E{"token_id_int", entity.SORT_DESC})
			}
			fsort = bson.D{{"$sort", sortDoc}}
			f = append(f, r.getPipelineForAuctionRequest(&filter)...)
		} else {
			addFields := bson.D{
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
			}

			//used for marketplace
			fMarketPlaceOffer = append(fMarketPlaceOffer, bson.E{"$lookup",
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
			})
			f1 = append(f1, bson.D{
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
			})

			f1 = append(f1, bson.D{{"$addFields", bson.D{{"price_erc20", "$listing_for_sales"}}}})
			f1 = append(f1, bson.D{
				{"$unwind",
					bson.D{
						{"path", "$price_erc20"},
						{"preserveNullAndEmptyArrays", true},
					},
				},
			})
			f1 = append(f1, bson.D{
				{"$addFields", addFields},
			})
			f1 = append(f1, bson.D{
				{"$addFields",
					bson.D{
						{"price",
							bson.D{
								{"$divide",
									bson.A{
										"$price",
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
			})
			fsort = bson.D{
				{"$sort",
					bson.D{
						{"buyable", -1},
						{"price", 1},
						{filter.SortBy, filter.Sort},
					},
				},
			}

			switch contractAddress {
			case strings.ToLower(os.Getenv("BNS_ADDRESS")):
				fName = append(fName, bson.D{
					{"$lookup",
						bson.D{
							{"from", "bns"},
							{"localField", "token_id"},
							{"foreignField", "token_id"},
							{"pipeline",
								bson.A{
									bson.D{
										{"$match",
											bson.D{
												{"collection_address", contractAddress},
											},
										},
									},
									bson.D{{"$skip", 0}},
									bson.D{{"$limit", 1}},
									bson.D{{"$project", bson.D{{"name", 1}}}},
								},
							},
							{"as", "bns"},
						},
					}},
					bson.D{
						{"$unwind",
							bson.D{
								{"path", "$bns"},
								{"preserveNullAndEmptyArrays", true},
							},
						},
					},
				)

				fieldName := bson.M{"name": "$bns.name"}
				fName = append(fName, bson.D{{
					"$addFields", fieldName,
				}})
				break
			default:
				fieldName := bson.M{"name": bson.D{
					{"$cond",
						bson.D{
							{"if",
								bson.D{
									{"$eq",
										bson.A{
											"$name",
											"",
										},
									},
								},
							},
							{"then",
								bson.D{
									{"$cond",
										bson.D{
											{"if",
												bson.D{
													{"$gt",
														bson.A{
															"$token_id_int",
															1000000,
														},
													},
												},
											},
											{"then",
												bson.D{
													{"$concat",
														bson.A{
															"$collection.name",
															" #",
															bson.D{
																{"$toString",
																	bson.D{
																		{"$mod",
																			bson.A{
																				"$token_id_int",
																				1000000,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
											{"else",
												bson.D{
													{"$concat",
														bson.A{
															"$collection.name",
															" #",
															"$token_id",
														},
													},
												},
											},
										},
									},
								},
							},
							{"else", "$name"},
						},
					},
				}}
				fName = append(fName, bson.D{{
					"$addFields", fieldName,
				}})
				break
			}
		}
	}

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

	fPagination = append(fPagination, fsort)
	fPagination = append(fPagination, bson.D{{"$skip", filter.Offset}})
	fPagination = append(fPagination, bson.D{{"$limit", filter.Limit}})

	//move them after limit and skip for performance
	fPagination = append(fPagination, bson.D{
		{"$lookup",
			bson.D{
				{"from", "collections"},
				{"localField", "collection_address"},
				{"foreignField", "contract"},
				{"as", "collection"},
			},
		},
	})
	fPagination = append(fPagination, bson.D{
		{"$unwind",
			bson.D{
				{"path", "$collection"},
				{"preserveNullAndEmptyArrays", true},
			},
		},
	})

	//lookup BNS
	fBNS = bson.D{
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
	}

	//lookup bns default
	fBNSDefault = bson.D{{"$lookup",
		bson.D{
			{"from", "bns_default"},
			{"localField", "owner"},
			{"foreignField", "resolver"},
			{"pipeline",
				bson.A{
					bson.D{{"$skip", 0}},
					bson.D{{"$limit", 1}},
					bson.M{"$lookup": bson.M{
						"from":         "bns",
						"localField":   "bns_default_id",
						"foreignField": "_id",
						"as":           "bns_default_data",
					}},
					bson.M{"$unwind": bson.M{
						"path":                       "$bns_default_data",
						"preserveNullAndEmptyArrays": true,
					}},
				},
			},
			{"as", "bns_default"},
		},
	}}

	if len(fMarketPlaceOffer) > 0 {
		fPagination = append(fPagination, fMarketPlaceOffer)
	}

	if len(fBNS) > 0 {
		fPagination = append(fPagination, fBNS)
	}

	if len(fBNSDefault) > 0 {
		fPagination = append(fPagination, fBNSDefault)
		fPagination = append(fPagination, bson.M{"$addFields": bson.M{"bns_data": bson.M{"$concatArrays": bson.A{"$bns_default.bns_default_data", "$bns_data"}}}})
	}

	if len(fName) > 0 {
		fPagination = append(fPagination, fName...)
	}
	//end

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

	cursor, err := r.DB.Collection(collection).Aggregate(context.TODO(), fAll)
	if err != nil {
		return nil, err
	}

	pResp := []*entity.MkpNftsPagination{}
	err = cursor.All(context.TODO(), &pResp)
	if err != nil {
		return nil, err
	}

	if len(pResp) == 0 {
		return nil, errors.New("Cannot get nfts")
	}

	return pResp[0], nil
}
