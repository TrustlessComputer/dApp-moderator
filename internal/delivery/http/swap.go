package http

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"
	req "dapp-moderator/utils/request"
)

func (h *httpDelivery) swapScanEvents(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.TcSwapScanEvents(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) swapScanHash(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.TcSwapScanEventsByTransactionHash(ctx, req.Query(r, "tx_hash", ""))
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}
