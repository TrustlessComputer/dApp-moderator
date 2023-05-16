package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	req "dapp-moderator/utils/request"
)

func (h *httpDelivery) swapTmTokenScanEvents(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.TcSwaTmTokenpScanEvents(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

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
			go h.Usecase.TcSwapScanEventsByTransactionHash(req.Query(r, "tx_hash", ""))
			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) clearCache(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			h.Usecase.ClearCache()
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
			data, err := h.Usecase.TcSwapFindSwapPairs(ctx, pagination, req.Query(r, "from_token", ""))
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
			tokenAddress := req.Query(r, "contract_address", "")
			data, err := h.Usecase.TcSwapFindSwapHistories(ctx, pagination, tokenAddress, "", "")
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
			fromToken := req.Query(r, "from_token", "")
			data, err := h.Usecase.FindTokensInPool(ctx, pagination, fromToken)
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
			address := req.Query(r, "address", "")
			sortCollum := req.Query(r, "sort", "")
			sortTypePrams := req.Query(r, "sort_type", "-1")
			sortType, _ := strconv.Atoi(sortTypePrams)
			data, err := h.Usecase.FindTokensReport(ctx, pagination, address, sortCollum, sortType)
			if err != nil {
				logger.AtLog.Logger.Error("FindTokensReport", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("FindTokensReport", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getTokensPrice(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := req.Query(r, "contract_address", "")
			chartType := req.Query(r, "chart_type", "")
			data, err := h.Usecase.FindTokensPrice(ctx, contractAddress, chartType)
			if err != nil {
				logger.AtLog.Logger.Error("FindTokensReport", zap.Error(err))
				return nil, err
			}

			//logger.AtLog.Logger.Info("FindTokensReport", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) jobGetBtcPrice(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.TcSwapUpdateWrapTokenPriceJob(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) addFrontEndLog(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			buf, bodyErr := ioutil.ReadAll(r.Body)
			var bodyRequest string
			if bodyErr == nil {
				bodyRequest = string(buf)
			}

			result := make(map[string]interface{})
			json.Unmarshal([]byte(bodyRequest), &result)

			err := h.Usecase.TcSwapAddFronEndLog(ctx, result)
			if err != nil {
				logger.AtLog.Logger.Error("addFrontEndLog", zap.Error(err))
				return nil, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) jobUpdateDataSwapSync(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.UpdateDataSwapSync(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("UpdateDataSwap", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) jobUpdateDataSwapHistory(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.UpdateDataSwapHistory(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("UpdateDataSwapHistory", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getSlackReport(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			channel := req.Query(r, "channel", "")
			err := h.Usecase.TcSwapSlackReport(ctx, channel)
			if err != nil {
				logger.AtLog.Logger.Error("getSlackReport", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) jobUpdateDataSwapPair(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.UpdateDataSwapPair(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("jobUpdateDataSwapPair", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) jobUpdateDataSwapToken(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.UpdateDataSwapToken(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("jobUpdateDataSwapToken", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) jobUpdateTotalSupply(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			err := h.Usecase.TcSwapUpdateTotalSupplyJob(ctx)
			if err != nil {
				logger.AtLog.Logger.Error("Tokens", zap.Error(err))
				return false, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) findPendingTransactionHistories(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
				return nil, err
			}
			txs := req.Query(r, "txs", "")

			data, err := h.Usecase.PendingTransactionHistories(ctx, pagination, txs)
			if err != nil {
				logger.AtLog.Logger.Error("findPendingTransactionHistories", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("findPendingTransactionHistories", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) addOrUpdateSwapWallet(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			var reqBody request.SwapWalletAddressRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqBody)
			if err != nil {
				return nil, err
			}
			res, err := h.Usecase.SwapAddOrUpdateWalletAddress(ctx, &reqBody)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) getSwapWallet(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			var reqBody request.SwapWalletAddressRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqBody)
			if err != nil {
				return nil, err
			}
			address := req.Query(r, "address", "")
			res, err := h.Usecase.SwapGetWalletAddress(ctx, address)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) addSwapBotConfig(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			var reqBody request.SwapBotConfigRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqBody)
			if err != nil {
				return nil, err
			}
			err = h.Usecase.AddSwapBotConfig(ctx, &reqBody)
			if err != nil {
				return nil, err
			}

			return true, nil
		},
	).ServeHTTP(w, r)
}
