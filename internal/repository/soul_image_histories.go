package repository

import (
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strings"
	"sync"
)

func (r *Repository) InsertSoulImageHistory(obj *entity.SoulImageHistories) error {
	_, err := r.InsertOne(obj)
	if err != nil {
		return err
	}
	return nil
}

// aggregate data to view_soul_nfts
func (r *Repository) PrepareSoulData(wg *sync.WaitGroup) error {
	defer wg.Done()

	fAll := bson.A{
		bson.D{{"$match", bson.D{{"collection_address", strings.ToLower(os.Getenv("SOUL_CONTRACT"))}}}},
		bson.D{
			{"$lookup",
				bson.D{
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
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$nft_auction_available"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
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
										{"status", 1},
									},
								},
							},
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
					{"from", "soul_attributes_percent_view"},
					{"localField", "token_id"},
					{"foreignField", "token_id"},
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
		bson.D{
			{"$merge",
				bson.D{
					{"into", utils.COLLECTION_SOUL_NFTS},
					{"whenMatched", "replace"},
				},
			},
		},
	}
	_, err := r.DB.Collection(utils.COLLECTION_NFTS).Aggregate(context.TODO(), fAll)
	if err != nil {
		return err
	}
	return nil
}
