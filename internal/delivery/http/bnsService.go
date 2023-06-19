package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get bns names
// @Description Get bns names
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Param name query string false "name"
// @Param token_id query string false "token id"
// @Param resolver query string false "resolver"
// @Param owner query string false "owner"
// @Param pfp query string false "pfp"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/names [GET]
func (h *httpDelivery) bnsNames(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			filter, err := h.createFilterBns(ctx, r, vars)
			if err != nil {
				logger.AtLog.Logger.Error("bnsNames", zap.Error(err))
				return nil, err
			}

			data, err := h.Usecase.BnsNames(ctx, *filter)
			if err != nil {
				logger.AtLog.Logger.Error("bnsNames", zap.Any("filter", filter), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsNames", zap.Any("filter", filter), zap.Any("data", len(data)))
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
// @Param token_id path string true "token_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/names/{token_id} [GET]
func (h *httpDelivery) bnsName(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			tokenID := vars["token_id"]

			data, err := h.Usecase.BnsName(ctx, tokenID)
			if err != nil {
				logger.AtLog.Logger.Error("bnsName", zap.Any("tokenID", tokenID), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsName", zap.Any("filter", tokenID), zap.Any("data", data))
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
				logger.AtLog.Logger.Error("bnsName", zap.Any("filter", name), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsName", zap.Any("filter", name), zap.Any("data", data))
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
			filter, err := h.createFilterBns(ctx, r, vars)
			if err != nil {
				logger.AtLog.Logger.Error("bnsNameOwnedByWalletAddress", zap.Error(err))
				return nil, err
			}

			data, err := h.Usecase.BnsNames(ctx, *filter)
			if err != nil {
				logger.AtLog.Logger.Error("bnsNameOwnedByWalletAddress", zap.Any("filter", filter), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bnsNameOwnedByWalletAddress", zap.Any("filter", filter), zap.Any("data", len(data)))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// @Summary bnsDefault
// @Description bnsDefault
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Param wallet_address path string true "wallet_address"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/default/{wallet_address} [GET]
func (h *httpDelivery) bnsDefault(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
		resolver := vars["wallet_address"]
		if resolver == "" {
			return nil, errors.New("resolver is required")
		}

		data, err := h.Usecase.BnsDefault(ctx, resolver)
		if err != nil {
			return nil, err
		}

		return data, nil
	}).ServeHTTP(w, r)
}

// @Summary updateBnsDefault
// @Description updateBnsDefault
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param request body request.UpdateBNSDefaultRequest true "body"
// @Param wallet_address path string true "user wallet address
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/default/{wallet_address} [PUT]
func (h *httpDelivery) updateBnsDefault(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
		resolver := vars["wallet_address"]
		if resolver == "" {
			return nil, errors.New("resolver is required")
		}
		var reqBody = &request.UpdateBNSDefaultRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(reqBody)
		if err != nil {
			return nil, err
		}
		if os.Getenv("ENV") != "local" {
			if resolver != ctx.Value(utils.SIGNED_WALLET_ADDRESS) {
				return nil, errors.New("permission denied")
			}
		}
		reqBody.Resolver = strings.ToLower(resolver)
		reqBody.TokenID = strings.ToLower(reqBody.TokenID)
		bns, err := h.Usecase.UpdateBnsDefault(ctx, reqBody)
		if err != nil {
			return nil, err
		}
		return bns, nil
	}).ServeHTTP(w, r)
}

func (h *httpDelivery) createFilterBns(ctx context.Context, r *http.Request, vars map[string]string) (*request.FilterBNSNames, error) {
	iPagination := ctx.Value(utils.PAGINATION)

	filter := request.FilterBNSNames{
		PaginationReq: iPagination.(request.PaginationReq),
	}

	name := r.URL.Query().Get("name")
	if name != "" {
		filter.Name = &name
	}

	owner := r.URL.Query().Get("owner")
	if owner != "" {
		filter.Owner = &owner
	}

	ownerVar := vars["wallet_address"]
	if ownerVar != "" {
		filter.Owner = &ownerVar
	}

	tokenID := r.URL.Query().Get("token_id")
	if tokenID != "" {
		filter.TokenID = &tokenID
	}

	tokenIDVar := vars["token_id"]
	if tokenIDVar != "" {
		filter.TokenID = &tokenIDVar
	}

	resolver := r.URL.Query().Get("resolver")
	if resolver != "" {
		filter.Resolver = &resolver
	}

	pfp := r.URL.Query().Get("pfp")
	if resolver != "" {
		filter.PFP = &pfp
	}

	return &filter, nil
}
