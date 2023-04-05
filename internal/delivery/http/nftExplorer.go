package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get Collections
// @Description Get Collections
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections [GET]
func (h *httpDelivery) collections(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			
			data, err := h.Usecase.Collections(ctx, iPagination.(request.PaginationReq))
			if err != nil {
				logger.AtLog.Logger.Error("Collections", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Collections", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Collections
// @Description Get Collections
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param contractAddress path string true "contractAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{contractAddress} [GET]
func (h *httpDelivery) collectionDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := vars["contractAddress"]
			data, err := h.Usecase.CollectionDetail(ctx, contractAddress)
			if err != nil {
				logger.AtLog.Logger.Error("collectionDetail", zap.String("contractAddress", contractAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionDetail", zap.String("contractAddress", contractAddress), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get nfts of a Collectionc
// @Description Get nfts of a Collectionc
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param contractAddress path string true "contractAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{contractAddress}/nfts [GET]
func (h *httpDelivery) collectionNfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := vars["contractAddress"]
			iPagination := ctx.Value(utils.PAGINATION)

			data, err := h.Usecase.CollectionNfts(ctx, contractAddress, iPagination.(request.PaginationReq))
			if err != nil {
				logger.AtLog.Logger.Error("collectionNfts", zap.String("contractAddress", contractAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionNfts", zap.String("contractAddress", contractAddress), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get nft detail of a Collection
// @Description Get nft detail of a Collection
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param contractAddress path string true "contractAddress"
// @Param tokenID path string true "tokenID"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{contractAddress}/nfts/{tokenID} [GET]
func (h *httpDelivery) collectionNftDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			contractAddress := vars["contractAddress"]
			tokenID := vars["tokenID"]
			data, err := h.Usecase.CollectionNftDetail(ctx, contractAddress, tokenID)
			if err != nil {
				logger.AtLog.Logger.Error("collectionNftDetail", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionNftDetail", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get nft content of a Collection
// @Description Get nft content of a Collection
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param contractAddress path string true "contractAddress"
// @Param tokenID path string true "tokenID"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{contractAddress}/nfts/{tokenID}/content [GET]
func (h *httpDelivery) collectionNftContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := context.Background()
	
	contractAddress := vars["contractAddress"]
	tokenID := vars["tokenID"]
	data, ctype, err := h.Usecase.CollectionNftContent(ctx, contractAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("collectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Error(err))
		h.Response.RespondWithError(w, http.StatusBadRequest, response.Error, err)
		return 
	}

	logger.AtLog.Logger.Info("collectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	
	w.Header().Set("Content-Type", ctype)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)
	return 
}

// UserCredits godoc
// @Summary Get Nfts
// @Description Get Nfts
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/nfts [GET]
func (h *httpDelivery) nfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)

			data, err := h.Usecase.Nfts(ctx, iPagination.(request.PaginationReq))
			if err != nil {
				logger.AtLog.Logger.Error("Nfts", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Nfts", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get nfts of a wallet address
// @Description Get nfts of a wallet address
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param ownerAddress path string true "ownerAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/owner-address/{ownerAddress}/nfts [GET]
func (h *httpDelivery) nftByWalletAddress(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			tokenID := vars["ownerAddress"]
			iPagination := ctx.Value(utils.PAGINATION)
			data, err := h.Usecase.NftByWalletAddress(ctx, tokenID, iPagination.(request.PaginationReq))
			if err != nil {
				logger.AtLog.Logger.Error("Nfts", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Nfts", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
