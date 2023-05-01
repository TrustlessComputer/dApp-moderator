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

func (h *httpDelivery) swapScanPairEvents(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.TcSwapScanPairEvents(ctx, 0)
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
			go h.Usecase.TcSwapScanEventsByTransactionHash(ctx, req.Query(r, "tx_hash", ""))
			// if err != nil {
			// 	logger.AtLog.Logger.Error("Tokens", zap.Error(err))
			// 	return false, err
			// }

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) findSwapPairs(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			data, err := h.Usecase.TcSwapFindSwapPairs(ctx, pagination, req.Query(r, "key", ""))
			if err != nil {
				logger.AtLog.Logger.Error("TcSwapFindSwapPairs", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("TcSwapFindSwapPairs", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) findSwapHistories(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			data, err := h.Usecase.TcSwapFindSwapHistories(ctx, pagination, req.Query(r, "key", ""))
			if err != nil {
				logger.AtLog.Logger.Error("TcSwapFindSwapHistories", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("TcSwapFindSwapHistories", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getTokensInPool(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			isTest := req.Query(r, "is_test", "")
			fromToken := req.Query(r, "from_token", "")
			data, err := h.Usecase.FindTokensInPool(ctx, pagination, fromToken, isTest)
			if err != nil {
				logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("FindTokensInPool", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getTokensReport(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			isTest := req.Query(r, "is_test", "")
			data, err := h.Usecase.FindTokensReport(ctx, pagination, isTest)
			if err != nil {
				logger.AtLog.Logger.Error("FindTokensReport", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("FindTokensReport", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
