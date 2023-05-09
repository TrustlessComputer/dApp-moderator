package http

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"
)

func (h *httpDelivery) swapTransactions(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.ScanTransactions(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}
