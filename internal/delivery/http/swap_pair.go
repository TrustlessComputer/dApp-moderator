package http

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"dapp-moderator/internal/delivery/http/response"
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
