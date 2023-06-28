package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	soul_contract "dapp-moderator/utils/contracts/soul"
	"dapp-moderator/utils/logger"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (u *Usecase) UpdateAuctionStatus(ctx context.Context) {
	chainLatestBlock, err := u.TCPublicNode.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
	}

	limit := int64(100)
	offset := int64(0)
	success := 0
	for {
		var auctions []*entity.Auction
		if err := u.Repo.Find(utils.COLLECTION_AUCTION, bson.D{
			{"status", entity.AuctionStatusInProgress.Ordinal()},
		}, limit, offset, &auctions, bson.D{{"_id", 1}}); err != nil {
			logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
			return
		}

		if len(auctions) == 0 {
			break
		}
		for _, item := range auctions {
			endTime, ok := new(big.Int).SetString(item.EndTimeBlock, 10)
			if !ok {
				logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus-parseEndTimeToBigInt", zap.Error(err))
			}

			if chainLatestBlock.Cmp(endTime) > 0 {
				if updateResult, err := u.Repo.UpdateOne(utils.COLLECTION_AUCTION, bson.D{{"_id", item.ID}}, bson.M{
					"$set": bson.M{"status": entity.AuctionStatusEnded.Ordinal()},
				}); err != nil {
					logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
					return
				} else {
					success += int(updateResult.ModifiedCount)
				}
			}
		}
		offset = offset + limit
	}

	logger.AtLog.Logger.Info("Finish Usecase.UpdateAuctionStatus", zap.Int64("chainLatestBlock", chainLatestBlock.Int64()),
		zap.Int("success", success))
}

func (u *Usecase) AuctionDetail(contractAddr, tokenID string) (*response.AuctionDetailResponse, error) {
	soulContract, err := soul_contract.NewSoul(common.HexToAddress(contractAddr), u.TCPublicNode.GetClient())
	if err != nil {
		return nil, err
	}
	tokenIDBigInt, ok := new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, errors.New("invalid token id")
	}
	resp, err := soulContract.Auctions(&bind.CallOpts{
		Context: context.Background(),
	}, tokenIDBigInt)
	if err != nil {
		logger.AtLog.Logger.Error("httpDelivery.auctionDetail", zap.Error(err))
		return nil, err
	}

	available, err := soulContract.Available(&bind.CallOpts{Context: context.Background()}, tokenIDBigInt)
	if err != nil {
		return nil, err
	}
	var auctionEntity = &entity.Auction{}
	err = u.Repo.FindOneWithResult(utils.COLLECTION_AUCTION, bson.M{
		"collection_address": contractAddr,
		"token_id":           tokenID,
	}, auctionEntity, &options.FindOneOptions{
		Sort: bson.M{"_id": -1},
	})
	if err != nil || auctionEntity.ID.IsZero() {
		return &response.AuctionDetailResponse{
			Available: available,
		}, nil
	}

	chainLatestBlock, err := u.TCPublicNode.GetBlockNumber()
	if err != nil {
		logger.AtLog.Logger.Error("Usecase.UpdateAuctionStatus", zap.Error(err))
	}

	status := entity.AuctionStatusInProgress
	if chainLatestBlock.Cmp(resp.EndTime) > 0 {
		status = entity.AuctionStatusEnded
	}
	if resp.Settled {
		status = entity.AuctionStatusSettled
	}

	if auctionEntity.Status.Ordinal() != status.Ordinal() {
		if _, err := u.Repo.UpdateOne(utils.COLLECTION_AUCTION, bson.D{{"_id", auctionEntity.ID}}, bson.M{
			"$set": bson.M{"status": status.Ordinal()},
		}); err != nil {
			logger.AtLog.Logger.Error("Usecase.AuctionDetail", zap.Error(err))
			return nil, err
		}
	}

	availableToFE := available
	if status == entity.AuctionStatusInProgress || status == entity.AuctionStatusEnded {
		availableToFE = true
	}

	return &response.AuctionDetailResponse{
		Available:      availableToFE,
		AuctionStatus:  status,
		HighestBid:     resp.Amount.String(),
		EndTime:        resp.EndTime.String(),
		DBAuctionID:    auctionEntity.ID.Hex(),
		ChainAuctionID: new(big.Int).SetBytes(resp.AuctionId[:]).String(),
	}, nil
}

func (u *Usecase) AuctionListBid(filterReq *request.FilterAuctionBid) (*response.AuctionListBidResponse, error) {
	filterStage := bson.M{}
	if filterReq.DBAuctionID != nil && *filterReq.DBAuctionID != "" {
		objectID, err := primitive.ObjectIDFromHex(*filterReq.DBAuctionID)
		if err != nil {
			logger.AtLog.Logger.Error("Usecase.AuctionListBid", zap.Error(err))
			return nil, err
		}

		var auctionEntity = &entity.Auction{}
		if err = u.Repo.FindOneWithResult(utils.COLLECTION_AUCTION, bson.M{"_id": objectID}, auctionEntity); err != nil {
			logger.AtLog.Logger.Error("Usecase.AuctionListBid", zap.Error(err))
			return nil, err
		}

		filterStage["chain_auction_id"] = auctionEntity.AuctionID
	}

	if filterReq.Sender != nil && *filterReq.Sender != "" {
		filterStage["sender"] = strings.ToLower(*filterReq.Sender)
	}

	type AuctionBidSummary struct {
		entity.AuctionBidSummary `bson:",inline"`
		BnsData                  []*entity.Bns      `bson:"bns_data,omitempty"`
		BnsDefault               *entity.BNSDefault `bson:"bns_default"`
		Auction                  *entity.Auction    `bson:"auction"`
		Ranking                  *int               `bson:"ranking"`
		TxHash                   string             `json:"tx_hash"`
		BlockNumber              uint               `json:"block_number_int"`
	}

	pipelines := bson.A{}
	if len(filterStage) > 0 {
		pipelines = append(pipelines, bson.M{
			"$match": filterStage,
		})
	}

	pipelines = append(pipelines,
		bson.M{"$lookup": bson.M{
			"from":         "bns",
			"localField":   "sender",
			"foreignField": "resolver",
			"as":           "bns_data",
		}},
		bson.M{"$lookup": bson.M{
			"from":         "bns_default",
			"localField":   "sender",
			"foreignField": "resolver",
			"pipeline": bson.A{
				bson.M{"$skip": 0},
				bson.M{"$limit": 1},
				bson.M{
					"$lookup": bson.M{
						"from":         "bns",
						"localField":   "bns_default_id",
						"foreignField": "_id",
						"as":           "bns_default_data",
					},
				},
				bson.M{
					"$unwind": bson.M{
						"path":                       "$bns_default_data",
						"preserveNullAndEmptyArrays": true,
					},
				},
			},
			"as": "bns_default",
		}},
		bson.M{
			"$unwind": bson.M{
				"path":                       "$bns_default",
				"preserveNullAndEmptyArrays": true,
			},
		},
	)

	if filterReq.Sender != nil && *filterReq.Sender != "" {
		pipelines = append(pipelines,
			bson.M{"$lookup": bson.M{
				"from":         "auction",
				"localField":   "db_auction_id",
				"foreignField": "_id",
				"as":           "auction",
			}},
			bson.M{
				"$unwind": bson.M{
					"path":                       "$auction",
					"preserveNullAndEmptyArrays": true,
				},
			},
			bson.M{"$lookup": bson.M{
				"from":         "auction_bid_summary",
				"localField":   "db_auction_id",
				"foreignField": "db_auction_id",
				"pipeline": bson.A{
					bson.M{"$addFields": bson.M{"amount_number": bson.M{"$toDouble": "$total_amount"}}},
					bson.M{"$sort": bson.M{"amount_number": -1}},
				},
				"as": "user_auctions",
			}},
			bson.M{"$addFields": bson.M{"ranking": bson.M{"$add": bson.A{
				bson.M{"$indexOfArray": bson.A{"$user_auctions.sender", strings.ToLower(*filterReq.Sender)}},
				1,
			}}}},
			bson.M{"$project": bson.M{"user_auctions": 0}},
		)
	}

	total, err := u.Repo.CountTotalFromPipeline(utils.COLLECTION_AUCTION_BID_SUMMARY, pipelines)
	if err != nil {
		return nil, err
	}

	limit, offset := filterReq.PaginationReq.GetOffsetAndLimit()
	//sortBy := "updated_at"
	//sort := -1
	//if filterReq.SortBy != nil && *filterReq.SortBy != "" {
	//	sortBy = *filterReq.SortBy
	//}
	//if filterReq.Sort != nil && *filterReq.Sort != 0 {
	//	sort = *filterReq.Sort
	//}
	pipelines = append(pipelines,
		bson.M{"$sort": bson.D{
			{"block_number_int", entity.SORT_DESC},
			{"log_index", entity.SORT_ASC},
		}})
	pipelines = append(pipelines, bson.D{{"$skip", offset}})
	pipelines = append(pipelines, bson.D{{"$limit", limit}})
	cursor, err := u.Repo.DB.Collection(utils.COLLECTION_AUCTION_BID_SUMMARY).Aggregate(context.TODO(), pipelines)
	if err != nil {
		return nil, err
	}

	resp := make([]*AuctionBidSummary, 0)
	err = cursor.All(context.TODO(), &resp)
	if err != nil {
		return nil, err
	}

	result := &response.AuctionListBidResponse{
		Items: make([]*response.AuctionListBidResponseItem, 0, len(resp)),
		Total: int64(total),
	}

	for _, item := range resp {
		name, avatar := "", ""
		if item.BnsDefault != nil && item.BnsDefault.BNSDefaultData != nil {
			avatar = item.BnsDefault.BNSDefaultData.PfpData.GCSUrl
			name = item.BnsDefault.BNSDefaultData.Name
		} else {
			if len(item.BnsData) > 0 {
				avatar = item.BnsData[0].PfpData.GCSUrl
				name = item.BnsData[0].Name
			}
		}
		updatedAt := item.UpdatedAt
		if updatedAt == nil {
			updatedAt = utils.ToPtr(time.Now())
		}
		responseItem := &response.AuctionListBidResponseItem{
			Amount:       item.TotalAmount,
			Sender:       item.Sender,
			BidderAvatar: avatar,
			BidderName:   name,
			Time:         *updatedAt,
			Auction:      item.Auction,
			Ranking:      item.Ranking,
			TxHash:       item.AuctionBidSummary.TxHash,
			BlockNumber:  item.AuctionBidSummary.BlockNumberInt,
		}

		if filterReq.Sender != nil && *filterReq.Sender != "" {
			if nftResp, err := u.GetMkplaceNft(context.TODO(), item.CollectionAddress, item.TokenID); err == nil {
				responseItem.MkpNftsResp = nftResp
			}
		}

		result.Items = append(result.Items, responseItem)
	}
	return result, nil
}
