package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

func (u *Usecase) TcSwapFindSwapPairs(ctx context.Context, filter request.PaginationReq, key string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapPairFilter{}
	query.FromPagination(filter)

	data, err = u.Repo.FindSwapPairs(ctx, query)

	if err != nil {
		logger.AtLog.Logger.Error("FindSwapPairs", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("FindSwapPairs", zap.Any("data", data))
	return data, nil
}

func (u *Usecase) TcSwapFindSwapHistories(ctx context.Context, filter request.PaginationReq, key string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapPairSwapHistoriesFilter{}
	query.FromPagination(filter)

	data, err = u.Repo.FindSwapPairSwapHistory(ctx, query)

	if err != nil {
		logger.AtLog.Logger.Error("TcSwapFindSwapHistories", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("TcSwapFindSwapHistories", zap.Any("data", data))
	return data, nil
}

func (u *Usecase) FindTokensInPool(ctx context.Context, filter request.PaginationReq, fromToken, isTest string) (interface{}, error) {
	var err error
	query := entity.TokenFilter{}
	query.FromPagination(filter)

	contracts := []string{}
	pairQuery := entity.SwapPairFilter{}
	pairQuery.Limit = 10000
	pairQuery.Page = 1
	if fromToken != "" {
		pairQuery.Token = fromToken
	}

	pairs, err := u.Repo.FindSwapPairs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
		return nil, err
	}

	for _, pair := range pairs {
		if fromToken == "" || (fromToken != "" && fromToken != pair.Token0) {
			contracts = append(contracts, pair.Token0)
		}

		if fromToken == "" || (fromToken != "" && fromToken != pair.Token1) {
			contracts = append(contracts, pair.Token1)
		}
	}

	tokens := []*entity.Token{}
	if len(contracts) > 0 {
		tokens, err = u.Repo.FindTokensInPoolByContracts(ctx, contracts, query)
		if err != nil {
			logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
			return nil, err
		}
	}

	logger.AtLog.Logger.Info("FindTokensInPool", zap.Any("data", tokens))
	return tokens, nil
}

func (u *Usecase) ClearCache() error {
	redisKey := fmt.Sprintf("tc-swap:token-reports-*")
	u.Cache.Delete(redisKey)
	return nil

}

func (u *Usecase) FindTokensReport(ctx context.Context, filter request.PaginationReq, isTest string) (interface{}, error) {
	query := entity.SwapPairFilter{}
	query.FromPagination(filter)

	redisKey := fmt.Sprintf("tc-swap:token-reports-%s-%s", query.Page, query.Limit)
	exists, err := u.Cache.Exists(redisKey)
	if err != nil {
		logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
		return 0, err
	}

	if *exists {
		dataInCache, err := u.Cache.GetData(redisKey)
		if err != nil {
			logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
			return nil, err
		}

		b := []byte(*dataInCache)
		reports := []entity.SwapPairReport{}
		err = json.Unmarshal(b, &reports)
		if err != nil {
			return nil, err
		}
		return reports, nil

	} else {
		reports, err := u.Repo.FindTokenReport(ctx, query)
		if err != nil {
			logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
			return nil, err
		}

		// btcPrice, _ := u.BlockChainApi.GetBitcoinPrice()
		btcPrice := u.Repo.ParseConfigByFloat64(ctx, "swap_btc_price")

		for _, item := range reports {
			if s, err := strconv.ParseFloat(item.Price.String(), 64); err == nil {
				item.BtcPrice = s
				item.UsdPrice = s * btcPrice
			}

			if s, err := strconv.ParseFloat(item.Volume.String(), 64); err == nil {
				item.BtcVolume = s
				item.UsdVolume = s * btcPrice
			}

			if item.Address == "0xfB83c18569fB43f1ABCbae09Baf7090bFFc8CBBD" {
				item.UsdPrice = btcPrice
			}
		}

		logger.AtLog.Logger.Info("FindTokensReport", zap.Any("data", reports))
		reportsStr, err := json.Marshal(&reports)
		if err != nil {
			logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
			return reports, nil
		}
		err = u.Cache.SetStringDataWithExpTime(redisKey, string(reportsStr), 5*60)
		if err != nil {
			logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
			return reports, nil
		}
		return reports, nil
	}
}
