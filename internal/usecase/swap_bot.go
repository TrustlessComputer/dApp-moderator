package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"math/big"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *Usecase) DoJobSwapBot(ctx context.Context, idoReq *request.IdoRequest) error {
	configs, err := u.Repo.FindSwapBotConfigs(ctx, entity.SwapBotConfigFilter{})
	if err != nil {
		logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
		return err
	}

	strDate := time.Now().Format("2006-01-02")
	for _, item := range configs {
		if !strings.EqualFold(item.CurrentDate, strDate) {
			item.CurrentDate = strDate
			reserve0, reserve1, err := u.BlockChainApi.TcSwapGetReserves(item.Pair)
			if err != nil {
				logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
				return err
			}
			item.BeginReserve0, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(reserve0, 18).String())
			item.BeginReserve1, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(reserve1, 18).String())

			token, baseToken, baseIndex, err := u.TcSwapGetBaseTokenOnPair(ctx, &item.SwapPair)
			if err != nil {
				logger.AtLog.Logger.Error("TcSwapPairCreateSwapEvent", zap.Error(err))
				return err
			}
			if token != nil && baseToken != nil {
				tmpPrice := big.NewFloat(0).Quo(helpers.ConvertWeiToBigFloat(reserve0, 18), helpers.ConvertWeiToBigFloat(reserve1, 18))
				if baseIndex == 1 {
					tmpPrice = big.NewFloat(0).Quo(helpers.ConvertWeiToBigFloat(reserve1, 18), helpers.ConvertWeiToBigFloat(reserve0, 18))
				}
				item.BeginPrice, _ = primitive.ParseDecimal128(tmpPrice.String())
			}
			err = u.Repo.UpdateSwapBotConfig(ctx, item)
			if err != nil {
				logger.AtLog.Logger.Error("Update mongo entity failed", zap.Error(err))
				return err
			}
		}
		tx, err := u.Repo.FindSwapBotTransaction(ctx, entity.SwapBotTransactionFilter{
			Status: 0,
		})

		if err != nil {
			logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
			return err
		}

		if tx != nil {
			return nil
		}

	}
	return nil
}
