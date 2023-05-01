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

func (u *Usecase) FindTokensReport(ctx context.Context, filter request.PaginationReq, isTest string) (interface{}, error) {
	var data interface{}
	query := entity.SwapPairFilter{}
	query.FromPagination(filter)

	reports, err := u.Repo.FindTokenReport(ctx, query)
	if err != nil {
		logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("FindTokensReport", zap.Any("data", data))
	return reports, nil
}
