package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"net/http"
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
			
			return nil, nil
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
// @Router /nft-explorer/{collectionAddress} [GET]
func (h *httpDelivery) collectionDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			
			return nil, nil
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
// @Router /nft-explorer/{collectionAddress}/nfts [GET]
func (h *httpDelivery) collectionNfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			
			return nil, nil
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
// @Param nftID path string true "nftID"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/{collectionAddress}/nfts/{nftID} [GET]
func (h *httpDelivery) collectionNftDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			
			return nil, nil
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
// @Param nftID path string true "nftID"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/{collectionAddress}/nfts/{nftID}/content [GET]
func (h *httpDelivery) collectionNftContent(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			
			return nil, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Collections
// @Description Get Collections
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/nfts [GET]
func (h *httpDelivery) nfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			
			return nil, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Collections
// @Description Get Collections
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/nfts [GET]
func (h *httpDelivery) collectionnfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			
			return nil, nil
		},
	).ServeHTTP(w, r)
}