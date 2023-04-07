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

// TokenExplorer godoc
// @Summary Get tokens
// @Description Get tokens
// @Tags token-explorer
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param key query string false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /token-explorer/tokens [GET]
func (h *httpDelivery) getTokens(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			data, err := h.Usecase.FindTokens(ctx, pagination, req.Query(r, "key", ""))
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// TokenExplorer godoc
// @Summary Get token detail
// @Description Get token detail
// @Tags token-explorer
// @Accept  json
// @Produce  json
// @Param address path string true "contractAddress"
// @Param payload body request.UpdateTokenReq true "contractAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /token-explorer/token/{address} [PUT]
func (h *httpDelivery) updateToken(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			address := vars["address"]
			var reqPayload request.UpdateTokenReq
			err := req.BindJson(r, &reqPayload)
			if err != nil {
				logger.AtLog.Logger.Error("invalid payload", zap.Error(err))
				return nil, err
			}

			err = h.Usecase.UpdateToken(ctx, address, reqPayload)
			if err != nil {
				logger.AtLog.Logger.Error("token", zap.String("address", address), zap.Error(err))
				return nil, err
			}

			return nil, nil
		},
	).ServeHTTP(w, r)
}

// TokenExplorer godoc
// @Summary Update token
// @Description Update token
// @Tags token-explorer
// @Accept  json
// @Produce  json
// @Param address path string true "contractAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /token-explorer/token/{address} [GET]
func (h *httpDelivery) getToken(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			address := vars["address"]
			data, err := h.Usecase.FindToken(ctx, address)
			if err != nil {
				logger.AtLog.Logger.Error("token", zap.String("address", address), zap.Error(err))
				return nil, err
			}
			return data, nil
		},
	).ServeHTTP(w, r)
}
