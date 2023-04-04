package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"
	"net/http"

	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary qn_addressBalance RPC Method
// @Description getaddress balance
// @Tags QuickNode
// @Accept  json
// @Produce  json
// @Param walletAddress path string true "BTC walletAddress"
// @Success 200 {object} response.JsonResponse{data=[]response.RedisResponse}
// @Router /quicknode/address/{walletAddress}/balance [GET]
func (h *httpDelivery) addressBalance(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["walletAddress"]
			resp, err := h.Usecase.AddressBalance(ctx, walletAddress)
			if err != nil {
				logger.AtLog.Logger.Error("addressBalance", zap.Error(err))
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}