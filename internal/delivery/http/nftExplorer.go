package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get Collections
// @Description Get Collections
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param owner query string false "owner"
// @Param contract query string false "contract"
// @Param allow_empty query bool false "allow_empty, default: false"
// @Param name query string false "name"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param sort_by query string false "default deployed_at_block"
// @Param sort query int false "default -1"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections [GET]
func (h *httpDelivery) collections(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			isAllowEmptyBool := false
			var err error

			owner := r.URL.Query().Get("owner")
			collectionAddress := r.URL.Query().Get("contract")
			name := r.URL.Query().Get("name")

			filter := request.CollectionsFilter{
				Owner:         &owner,
				Address:       &collectionAddress,
				Name:          &name,
				PaginationReq: p,
			}

			isAllowEmpty := r.URL.Query().Get("allow_empty")
			if isAllowEmpty != "" {
				isAllowEmptyBool, err = strconv.ParseBool(isAllowEmpty)
				if err != nil {
					isAllowEmptyBool = false
				}

			}

			filter.AllowEmpty = &isAllowEmptyBool
			data, err := h.Usecase.Collections(ctx, filter)
			if err != nil {
				logger.AtLog.Logger.Error("collections", zap.Any("filter", filter), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collections", zap.Any("filter", filter), zap.Int("data", len(data)))
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
// @Summary Update Collection
// @Description Update Collection
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param request body structure.UpdateCollection true "UpdateCollection"
// @Param contractAddress path string true "contractAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{contractAddress} [PUT]
func (h *httpDelivery) updateCollectionDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := vars["contractAddress"]
			reqBody := &structure.UpdateCollection{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				logger.AtLog.Logger.Error("collectionDetail", zap.String("contractAddress", contractAddress), zap.Any("reqBody", reqBody), zap.Error(err))
				return nil, err
			}

			iwalletAdress := ctx.Value(utils.SIGNED_WALLET_ADDRESS)
			walletAdress, ok := iwalletAdress.(string)
			if !ok {
				err := errors.New("Token is incorect")
				logger.AtLog.Logger.Error("collectionDetail", zap.String("contractAddress", contractAddress), zap.Any("reqBody", reqBody), zap.Error(err))
				return nil, err
			}

			data, err := h.Usecase.UpdateCollection(ctx, contractAddress, walletAdress, reqBody)
			if err != nil {
				logger.AtLog.Logger.Error("collectionDetail", zap.String("contractAddress", contractAddress), zap.Any("reqBody", reqBody), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionDetail", zap.String("contractAddress", contractAddress), zap.Any("reqBody", reqBody), zap.Any("data", data))
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
// @Param is_big_file query bool false "is_big_file"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param name query string false "name"
// @Param owner query string false "owner"
// @Param contractAddress path string true "contractAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/collections/{contractAddress}/nfts [GET]
func (h *httpDelivery) collectionNfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := strings.ToLower(vars["contractAddress"])
			iPagination := ctx.Value(utils.PAGINATION)

			p := iPagination.(request.PaginationReq)
			var err error

			owner := r.URL.Query().Get("owner")
			name := r.URL.Query().Get("name")

			filter := request.CollectionsFilter{
				Owner:         &owner,
				Address:       &contractAddress,
				Name:          &name,
				PaginationReq: p,
			}

			isBigFile := r.URL.Query().Get("is_big_file")
			if isBigFile != "" {
				isBigFileBool, err := strconv.ParseBool(isBigFile)
				if err == nil {
					filter.IsBigFile = &isBigFileBool
				}
			}

			if strings.ToLower(contractAddress) == strings.ToLower("0x16EfDc6D3F977E39DAc0Eb0E123FefFeD4320Bc0") {
				if r.URL.Query().Get("allow_empty") == "false" {
					// artifact
					tmp := true
					filter.ContentTypeNotEmpty = &tmp
				}
			}

			coll, err := h.Usecase.CollectionDetail(ctx, contractAddress)
			if err != nil {
				logger.AtLog.Logger.Error("collectionNfts", zap.Any("iPagination", iPagination), zap.String("contractAddress", contractAddress), zap.Error(err))
			}
			data, err := h.Usecase.CollectionNfts(ctx, contractAddress, filter)
			bnsAddress := strings.ToLower(os.Getenv("BNS_ADDRESS"))
			for _, i := range data {
				if i.Name == "" {
					if bnsAddress == contractAddress {
						key := helpers.BnsTokenNameKey(i.TokenID)
						existed, _ := h.Usecase.Cache.Exists(key)
						if existed != nil && *existed == true {
							cached, _ := h.Usecase.Cache.GetData(key)
							if cached != nil {
								i.Name = *cached
							}
						} else {
							bnsName, _ := h.Usecase.BnsService.NameByToken(i.TokenID)
							if bnsName != nil {
								i.Name = bnsName.Name
								h.Usecase.Cache.SetStringData(key, i.Name)
							}

						}
					} else {
						tokenIdBigInt, _ := new(big.Int).SetString(i.TokenID, 10)
						g, _ := new(big.Int).SetString("1000000", 10)
						if tokenIdBigInt.Cmp(g) > 0 {
							// TODO maybe from generative
							i.Name = coll.Name + " #" + tokenIdBigInt.Mod(tokenIdBigInt, g).String()
						} else {
							i.Name = coll.Name + " #" + i.TokenID
						}
					}
				}
			}
			if err != nil {
				logger.AtLog.Logger.Error("collectionNfts", zap.Any("iPagination", iPagination), zap.String("contractAddress", contractAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("collectionNfts", zap.Any("iPagination", iPagination), zap.String("contractAddress", contractAddress), zap.Any("data", len(data)))
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

			contractAddress := strings.ToLower(vars["contractAddress"])
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

	logger.AtLog.Logger.Info("collectionNftContent", zap.String("contractAddress", contractAddress), zap.String("tokenID", tokenID), zap.Any("data", len(data)))

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
				logger.AtLog.Logger.Error("Nfts", zap.Any("iPagination", iPagination), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Nfts", zap.Any("iPagination", iPagination), zap.Any("data", len(data)))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get tokens of a wallet address
// @Description Get tokens of a wallet address
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
				logger.AtLog.Logger.Error("nftByWalletAddress", zap.Any("pagination", iPagination), zap.String("tokenID", tokenID), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("nftByWalletAddress", zap.Any("pagination", iPagination), zap.String("tokenID", tokenID), zap.Int("data", len(data)))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary refresh-nft
// @Description refresh-nft
// @Tags nft-explorer
// @Accept  json
// @Produce  json
// @Param contractAddress path string true "contractAddress"
// @Param tokenID path string true "tokenID"
// @Success 200 {object} response.JsonResponse{}
// @Router /nft-explorer/refresh-nft/contracts/{contractAddress}/token/{tokenID} [GET]
func (h *httpDelivery) refreshNft(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := vars["contractAddress"]
			tokenID := vars["tokenID"]

			resp, err := h.Usecase.RefreshNft(context.Background(), contractAddress, tokenID)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}
