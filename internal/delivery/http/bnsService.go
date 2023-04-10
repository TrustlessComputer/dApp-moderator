package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"net/http"

	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get bns names
// @Description Get bns names
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/names [GET]
func (h *httpDelivery) bnsNames(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			
			filter := request.FilterBNSNames{
				PaginationReq: iPagination.(request.PaginationReq),
			}

			data, err := h.Usecase.BnsNames(ctx, filter)
			if err != nil {
				logger.AtLog.Logger.Error("bnsNames",zap.Any("filter",filter), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsNames",zap.Any("filter",filter), zap.Any("data", len(data)))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get detail of bns name
// @Description detail of bns name
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Param name path string false "name"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/names/{name} [GET]
func (h *httpDelivery) bnsName(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			name := vars["name"]

			data, err := h.Usecase.BnsName(ctx, name)
			if err != nil {
				logger.AtLog.Logger.Error("bnsName",zap.Any("filter",name), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsName",zap.Any("filter",name), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Check bns name available for register
// @Description Check bns name available for register
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Param name path string false "name"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/names/{name}/available [GET]
func (h *httpDelivery) bnsNameAvailable(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			name := vars["name"]

			data, err := h.Usecase.BnsNameAvailable(ctx, name)
			if err != nil {
				logger.AtLog.Logger.Error("bnsName",zap.Any("filter",name), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsName",zap.Any("filter",name), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get name of a wallet-address
// @Description Get name of a wallet-address
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Param wallet_address path string false "wallet_address"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/names/owned/{wallet_address} [GET]
func (h *httpDelivery) bnsNameOwnedByWalletAddress(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["wallet_address"]
			iPagination := ctx.Value(utils.PAGINATION)
			filter := request.FilterBNSNames{
				PaginationReq: iPagination.(request.PaginationReq),
			}

			data, err := h.Usecase.BnsNamesOnwedByWalletAddress(ctx, walletAddress, filter)
			if err != nil {
				logger.AtLog.Logger.Error("bnsName",zap.Any("filter",walletAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsName",zap.Any("filter",walletAddress), zap.Any("data", len(data)))
			return data, nil
		},
	).ServeHTTP(w, r)
}
