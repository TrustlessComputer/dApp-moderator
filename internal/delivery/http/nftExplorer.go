package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
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
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections [GET]
func (h *httpDelivery) collections(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			data, err := h.Usecase.Collections(ctx)
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
// @Param collectionAddress path string true "collectionAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{collectionAddress} [GET]
func (h *httpDelivery) collectionDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			collectionAddress := vars["collectionAddress"]
			data, err := h.Usecase.CollectionDetail(ctx, collectionAddress)
			if err != nil {
				logger.AtLog.Logger.Error("collectionDetail", zap.String("collectionAddress", collectionAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionDetail", zap.String("collectionAddress", collectionAddress), zap.Any("data", data))
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
// @Param collectionAddress path string true "collectionAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{collectionAddress}/nfts [GET]
func (h *httpDelivery) collectionNfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			collectionAddress := vars["collectionAddress"]
			data, err := h.Usecase.CollectionNfts(ctx, collectionAddress)
			if err != nil {
				logger.AtLog.Logger.Error("collectionNfts", zap.String("collectionAddress", collectionAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionNfts", zap.String("collectionAddress", collectionAddress), zap.Any("data", data))
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
// @Param collectionAddress path string true "collectionAddress"
// @Param tokenID path string true "tokenID"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{collectionAddress}/nfts/{tokenID} [GET]
func (h *httpDelivery) collectionNftDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			collectionAddress := vars["collectionAddress"]
			tokenID := vars["tokenID"]
			data, err := h.Usecase.CollectionNftDetail(ctx, collectionAddress, tokenID)
			if err != nil {
				logger.AtLog.Logger.Error("collectionNftDetail", zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionNftDetail", zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
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
// @Param collectionAddress path string true "collectionAddress"
// @Param tokenID path string true "tokenID"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{collectionAddress}/nfts/{tokenID}/content [GET]
func (h *httpDelivery) collectionNftContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := context.Background()
	
	collectionAddress := vars["collectionAddress"]
	tokenID := vars["tokenID"]
	data, ctype, err := h.Usecase.CollectionNftContent(ctx, collectionAddress, tokenID)
	if err != nil {
		logger.AtLog.Logger.Error("collectionNftContent", zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Error(err))
		h.Response.RespondWithError(w, http.StatusBadRequest, response.Error, err)
		return 
	}

	logger.AtLog.Logger.Info("collectionNftContent", zap.String("collectionAddress", collectionAddress), zap.String("tokenID", tokenID), zap.Any("data", data))
	
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
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/nfts [GET]
func (h *httpDelivery) nfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			data, err := h.Usecase.Nfts(ctx)
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
// @Param ownerAddress path string true "ownerAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/owner-address/{ownerAddress}/nfts [GET]
func (h *httpDelivery) nftByWalletAddress(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			tokenID := vars["ownerAddress"]
			data, err := h.Usecase.NftByWalletAddress(ctx, tokenID)
			if err != nil {
				logger.AtLog.Logger.Error("Nfts", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Nfts", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}
