package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

// TokenExplorer godoc
// @Summary Get tokens
// @Description Get tokens
// @Tags token-explorer
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /token-explorer/tokens [GET]
func (h *httpDelivery) tokens(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}

			data, err := h.Usecase.Tokens(ctx, pagination)
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
