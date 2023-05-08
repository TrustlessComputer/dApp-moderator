package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	req "dapp-moderator/utils/request"

	"go.uber.org/zap"
)

func (h *httpDelivery) addOrUpdateSwapIdo(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			var reqBody request.IdoRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqBody)
			if err != nil {
				return nil, err
			}
			res, err := h.Usecase.SwapAddOrUpdateIdo(ctx, &reqBody)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) findSwapIdoHistories(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			data, err := h.Usecase.SwapFindSwapIdoHistories(ctx, pagination)
			if err != nil {
				logger.AtLog.Logger.Error("TcSwapFindSwapHistories", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("TcSwapFindSwapHistories", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) findSwapIdoDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			id := req.Query(r, "id", "")
			data, err := h.Usecase.SwapFindSwapIdoDetail(ctx, id)
			if err != nil {
				logger.AtLog.Logger.Error("findSwapIdoDetail", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("findSwapIdoDetail", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) deleteSwapIdo(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			id := req.Query(r, "id", "")
			data, err := h.Usecase.SwapDeleteSwapIdo(ctx, id)
			if err != nil {
				logger.AtLog.Logger.Error("findSwapIdoDetail", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("findSwapIdoDetail", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getSwapIdoTokens(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			owner := req.Query(r, "owner", "")
			data, err := h.Usecase.SwapFindTokens(ctx, pagination, owner)
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Tokens", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
