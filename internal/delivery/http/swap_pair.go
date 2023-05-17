package http

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	req "dapp-moderator/utils/request"
)

func (h *httpDelivery) getLiquidityApr(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			pairAddress := req.Query(r, "pair", "")
			data, err := h.Usecase.SwapGetPairApr(ctx, pairAddress)
			if err != nil {
				logger.AtLog.Logger.Error("getLiquidityApr", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("getLiquidityApr", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getListLiquidityAprReport(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}

			search := req.Query(r, "search", "")
			data, err := h.Usecase.SwapGetPairAprListReport(ctx, pagination, search)
			if err != nil {
				logger.AtLog.Logger.Error("getListLiquidityAprReport", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("getListLiquidityAprReport", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getRoutePair(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			fromToken := req.Query(r, "from_token", "")
			toToken := req.Query(r, "to_token", "")
			data, err := h.Usecase.GetRoutePair(ctx, fromToken, toToken)
			if err != nil {
				logger.AtLog.Logger.Error("getRoutePair", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("getRoutePair", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
