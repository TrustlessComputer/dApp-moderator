package http

import (
	"context"
	"net/http"

	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)

func (h *httpDelivery) swapJobUpdateIdoStatus(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			data, err := h.Usecase.SwapJobUpdateIdoStatus(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("swapJobUpdateIdoStatus", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("swapJobUpdateIdoStatus", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
