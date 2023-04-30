package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"

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
	var data interface{}
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

	var tokens []*entity.Token
	if len(contracts) > 0 {
		tokens, err = u.Repo.FindTokensInPoolByContracts(ctx, contracts, query)
		if err != nil {
			logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
			return nil, err
		}
	}

	if len(tokens) == 0 && isTest != "" {
		token0 := &entity.Token{}
		token0.Address = "0x435bdab1bcB2fcf80e5cF47dba209E28c340c3Bf"
		token0.Name = "WBTC"
		token0.Symbol = "WBTC"
		token0.Decimal = 18
		tokens = append(tokens, token0)

		token1 := &entity.Token{}
		token1.Address = "0xA9CBb5F80445ff60faED26bFa37128F91Ac7E0E5"
		token1.Name = "DUNGT"
		token1.Symbol = "DUNGT"
		token1.Decimal = 18
		tokens = append(tokens, token1)
	}
	data = tokens

	logger.AtLog.Logger.Info("FindTokensInPool", zap.Any("data", data))
	return data, nil
}
