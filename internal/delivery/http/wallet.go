package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"
	"net/http"

	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get Wallet's info
// @Description Get Wallet's info
// @Tags Wallet
// @Accept  json
// @Produce  json
// @Param walletAddress path string true "walletAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /wallets/{walletAddress} [GET]
func (h *httpDelivery) walletInfo(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["walletAddress"]
			data, err := h.Usecase.GetBTCWalletInfo(ctx, walletAddress)
			if err != nil {
				logger.AtLog.Logger.Error("walletAddress", zap.String("walletAddress", walletAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("walletAddress", zap.String("walletAddress", walletAddress), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Wallet's txs
// @Description Get Wallet's txs
// @Tags Wallet
// @Accept  json
// @Produce  json
// @Param walletAddress path string true "walletAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /wallets/{walletAddress}/txs [GET]
func (h *httpDelivery) walletTxs(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["walletAddress"]
			data, err := h.Usecase.GetBTCWalletTXS(ctx, walletAddress)
			if err != nil {
				logger.AtLog.Logger.Error("walletAddress", zap.String("walletAddress", walletAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("walletAddress", zap.String("walletAddress", walletAddress), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
