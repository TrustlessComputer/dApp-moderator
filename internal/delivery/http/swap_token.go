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

func (h *httpDelivery) findTmTokenHistories(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			data, err := h.Usecase.TcTmTokenTransferHistories(ctx, pagination, req.Query(r, "address", ""))
			if err != nil {
				logger.AtLog.Logger.Error("TcSwapFindSwapHistories", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("TcSwapFindSwapHistories", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
