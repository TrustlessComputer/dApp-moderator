package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *Usecase) SwapJobUpdateIdoStatus(ctx context.Context) (bool, error) {
	var err error
	query := entity.SwapIdoFilter{}
	// query.FromPagination(filter)
	query.CheckStartTime = -1
	query.Status = entity.SwapIdoStatusUpcoming

	idos, err := u.Repo.FindSwapIdos(ctx, query)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("SwapJobUpdateIdoStatus", zap.Error(err))
		return false, err
	}

	for _, item := range idos {
		pair, _ := u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
			FromToken: item.Address,
		})
		if pair != nil {
			item.Status = entity.SwapIdoStatusStated
			item.Pair = pair
			u.Repo.UpdateSwapIdo(ctx, item)
		}
	}

	logger.AtLog.Logger.Info("SwapJobUpdateIdoStatus", zap.Any("data", idos))
	return true, nil
}
