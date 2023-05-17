package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindSwapPairSwapHistory(ctx context.Context, filter entity.SwapPairSwapHistoriesFilter) (*entity.SwapPairSwapHistories, error) {
	var swapPairHistories entity.SwapPairSwapHistories
	err := r.DB.Collection(utils.COLLECTION_SWAP_HISTORIES).FindOne(ctx, r.parseSwapPairSwapHistories(filter)).Decode(&swapPairHistories)
	if err != nil {
		return nil, err
	}
	return &swapPairHistories, nil
}

func (r *Repository) parseSwapPairSwapHistories(filter entity.SwapPairSwapHistoriesFilter) bson.M {
	andCond := make([]bson.M, 0)
	if filter.ContractAddress != "" {
		andCond = append(andCond, bson.M{"contract_address": filter.ContractAddress})
	}

	if filter.TxHash != "" {
		andCond = append(andCond, bson.M{"tx_hash": filter.TxHash})
	}

	if filter.UserAddress != "" {
		andCond = append(andCond, bson.M{"to": filter.UserAddress})
	}

	if filter.Token != "" {
		andCond = append(andCond, bson.M{"token": filter.Token})
	}

	if filter.Symbol != "" {
		andCond = append(andCond, bson.M{"base_token_symbol": bson.M{"$ne": filter.Symbol}})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindTokenReport(ctx context.Context, filter entity.TokenReportFilter) ([]*entity.SwapPairReport, error) {
	var tokens []*entity.SwapPairReport
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)
	if filter.SortBy != "" {
		options.SetSort(bson.D{{filter.SortBy, filter.SortType}})
	} else {
		options.SetSort(bson.D{{"priority", -1}, {"total_volume", -1}, {"percent_7day", -1}})
	}

	cursor, err := r.DB.Collection(utils.VIEW_SWAP_REPORT_FINAL).Find(ctx, r.parseTokenReportFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var token *entity.SwapPairReport
		err = cursor.Decode(&token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (r *Repository) parseTokenReportFilter(filter entity.TokenReportFilter) bson.M {
	andCond := make([]bson.M, 0)
	if filter.Address != "" {
		andCond = append(andCond, bson.M{"address": filter.Address})
	}

	if filter.CreatedBy != "" {
		andCond = append(andCond, bson.M{"owner": filter.CreatedBy})
	}

	if filter.Search != "" {
		andCond = append(andCond, bson.M{"$or": []bson.M{
			{"symbol": primitive.Regex{Pattern: filter.Search, Options: "i"}},
			{"name": primitive.Regex{Pattern: filter.Search, Options: "i"}},
			{"address": primitive.Regex{Pattern: filter.Search, Options: "i"}},
			{"owner": primitive.Regex{Pattern: filter.Search, Options: "i"}},
		}})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindTokePrice(ctx context.Context, contract string, chartType string) ([]*entity.ChartDataResp, error) {
	var tokens []*entity.ChartDataResp

	// pagination
	// Set the options for the query

	options := options.Find()
	options.SetSort(bson.D{{"created_at", 1}})
	var swapPair entity.SwapPair

	err := r.DB.Collection(utils.COLLECTION_SWAP_PAIR).FindOne(ctx, bson.D{
		{"token0", contract},
		{"token1", "0xfB83c18569fB43f1ABCbae09Baf7090bFFc8CBBD"},
	}).Decode(&swapPair)
	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_HISTORIES).Find(ctx, bson.D{
		{"contract_address", swapPair.Pair},
	}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		//var token *entity.ChartDataResp
		var h *entity.SwapPairSwapHistories
		err = cursor.Decode(&h)

		if err != nil {
			return nil, err
		}
		tokens = updateChartData(h, tokens, chartType)
		//tokens = append(tokens, token)
	}
	return tokens, nil
}

func CompareDecimal128(d1, d2 primitive.Decimal128) int {
	b1, exp1, err := d1.BigInt()
	if err != nil {
		return 0
	}
	b2, exp2, err := d2.BigInt()
	if err != nil {
		return 0
	}

	sign := b1.Sign()
	if sign != b2.Sign() {
		if b1.Sign() > 0 {
			return 1
		} else {
			return -1
		}
	}

	if exp1 == exp2 {
		return b1.Cmp(b2)
	}

	if sign < 0 {
		if exp1 < exp2 {
			return 1
		}
		return -1
	} else {
		if exp1 < exp2 {
			return -1
		}

		return 1
	}
}

func updateChartData(h *entity.SwapPairSwapHistories, res []*entity.ChartDataResp, chartType string) []*entity.ChartDataResp {
	t := FindTime(*h.CreatedAt, chartType)
	isExit := false
	for i, s := range res {
		if s.Time == t {
			isExit = true
			res[i].Close = h.Price
			volmeString := h.Volume.String()
			s, _ := strconv.ParseFloat(volmeString, 32)

			res[i].TotalVolume = res[i].TotalVolume + s
			if CompareDecimal128(res[i].VolumeFrom, h.Volume) > 0 {
				res[i].VolumeFrom = h.Volume
			}
			if CompareDecimal128(res[i].VolumeTo, h.Volume) < 0 {
				res[i].VolumeTo = h.Volume
			}
			if CompareDecimal128(res[i].High, h.Price) < 0 {
				res[i].High = h.Price
			}
			if CompareDecimal128(res[i].Low, h.Price) > 0 {
				res[i].Low = h.Price
			}
			break
		}

	}
	if !isExit {
		var token *entity.ChartDataResp
		token = new(entity.ChartDataResp)
		token.Time = t
		token.Timestamp = t.Unix()
		token.Close = h.Price
		token.Open = h.Price
		volmeString := h.Volume.String()
		s, _ := strconv.ParseFloat(volmeString, 32)
		token.TotalVolume = s
		token.VolumeFrom = h.Volume
		token.VolumeTo = h.Volume
		token.High = h.Price
		token.Low = h.Price
		token.ConversionSymbol = ""
		token.ConversionType = ""
		res = append(res, token)
	}
	return res
}

func FindTime(t time.Time, chartType string) time.Time {
	year, month, day := t.Date()
	hr, min, _ := t.Clock()
	if chartType == "minute" {
		return time.Date(year, month, day, hr, min, 0, 0, t.Location())
	}
	if chartType == "hour" {
		return time.Date(year, month, day, hr, 0, 0, 0, t.Location())
	}
	if chartType == "day" {
		return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	}
	if chartType == "month" {
		return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	}
	return t
}

func (r *Repository) FindSwapPairHistories(ctx context.Context, filter entity.SwapPairSwapHistoriesFilter) ([]*entity.SwapPairSwapHistories, error) {
	var pairs []*entity.SwapPairSwapHistories

	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)
	options.SetSort(bson.D{{"timestamp", -1}})

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_HISTORIES).Find(ctx, r.parseSwapPairSwapHistories(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		pair := &entity.SwapPairSwapHistories{}
		err = cursor.Decode(pair)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

func (r *Repository) UpdateSwapPairHistory(ctx context.Context, sync *entity.SwapPairSwapHistories) error {
	collectionName := sync.CollectionName()
	r.DB.Collection(collectionName).FindOneAndUpdate(ctx, bson.M{"tx_hash": sync.TxHash, "contract_address": sync.ContractAddress}, bson.M{"$set": sync}, nil)
	// if err != nil {
	// 	return err
	// }
	// if err != nil {
	// 	return err
	// }
	return nil
}
