package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get market place listing
// @Description Get market place listing
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract address"
// @Param token_id path string true "token_id"
// @Param status query bool false "0: open, 1: cancel, 2: done, default all"
// @Param sort_by query string false "sort by field"
// @Param sort query int false "1: ASC, -1: DESC"
// @Param limit query int false "limit default 10"
// @Param page query int false "page start with 1"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/listing/{contract_address}/token/{token_id} [GET]
func (h *httpDelivery) getListingViaGenAddressTokenID(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			_ = p

			f := entity.FilterMarketplaceListings{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
			}

			resp, err := h.Usecase.FilterMKListing(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get market place offers
// @Description Get market place offers
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract address"
// @Param token_id path string true "token_id"
// @Param status query bool false "0: open, 1: cancel, 2: done, default all"
// @Param sort_by query string false "sort by field"
// @Param sort query int false "1: ASC, -1: DESC"
// @Param limit query int false "limit default 10"
// @Param page query int false "page start with 1"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/offers/{contract_address}/token/{token_id} [GET]
func (h *httpDelivery) getOfferViaGenAddressTokenID(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			_ = p

			f := entity.FilterMarketplaceOffer{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
			}

			resp, err := h.Usecase.FilterMKOffers(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get listing of a profile
// @Description listing of a profile
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param wallet_address path string true "wallet_address"
// @Param status query bool false "0: open, 1: cancel, 2: done, default all"
// @Param sort_by query string false "sort by field"
// @Param sort query int false "1: ASC, -1: DESC"
// @Param limit query int false "limit default 10"
// @Param page query int false "page start with 1"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/wallet/{wallet_address}/listing [GET]
func (h *httpDelivery) getListingOfAProfile(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			wa := vars["wallet_address"]

			f := entity.FilterMarketplaceListings{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				SellerAddress: &wa,
			}

			resp, err := h.Usecase.FilterMKListing(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get offers of a profile
// @Description Offers of a profile
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param wallet_address path string true "wallet_address"
// @Param status query bool false "0: open, 1: cancel, 2: done, default all"
// @Param sort_by query string false "sort by field"
// @Param sort query int false "1: ASC, -1: DESC"
// @Param limit query int false "limit default 10"
// @Param page query int false "page start with 1"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/wallet/{wallet_address}/offer [GET]
func (h *httpDelivery) getOffersOfAProfile(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			wa := vars["wallet_address"]

			f := entity.FilterMarketplaceOffer{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				BuyerAddress: &wa,
			}

			resp, err := h.Usecase.FilterMKOffers(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get token's activities
// @Description Get token's activities
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract_address"
// @Param token_id path string true "token_id"
// @Param status query bool false "0: open, 1: cancel, 2: done, default all"
// @Param sort_by query string false "sort by field"
// @Param sort query int false "1: ASC, -1: DESC"
// @Param limit query int false "limit default 10"
// @Param page query int false "page start with 1"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/contract/{contract_address}/token/{token_id}/activities [GET]
func (h *httpDelivery) getTokenActivities(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			tokenID := vars["token_id"]
			contractAddresss := vars["contract_address"]

			f := entity.FilterTokenActivities{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				TokenID:         &tokenID,
				ContractAddress: &contractAddresss,
			}

			resp, err := h.Usecase.FilterTokenActivities(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Collections
// @Description Get Collections
// @Tags MarketPlace
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
// @Router /marketplace/collections [GET]
func (h *httpDelivery) mkpCollections(w http.ResponseWriter, r *http.Request) {
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
			data, err := h.Usecase.MarketplaceCollections(ctx, filter)
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
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract address"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/collections/{contract_address} [GET]
func (h *httpDelivery) mkpCollectionDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := strings.ToLower(vars["contract_address"])
			data, err := h.Usecase.MarketplaceCollectionDetail(ctx, contractAddress)
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
// @Summary Get marketplace Nfts
// @Description Get marketplace Nfts
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address query string false "contract_address"
// @Param token_id query string false "token_id"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/nfts [GET]
func (h *httpDelivery) mkplaceNfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)

			if p.SortBy == nil {
				sortBy := "token_id_int"
				p.SortBy = &sortBy
			}

			if p.Sort == nil {
				sort := int(entity.SORT_DESC)
				p.Sort = &sort
			}

			f := entity.FilterNfts{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					SortBy: *p.SortBy,
					Sort:   entity.SortType(*p.Sort),
				},
			}

			ca := r.URL.Query().Get("contract_address")
			tokID := r.URL.Query().Get("token_id")

			if ca != "" {
				f.ContractAddress = &ca
			}

			if tokID != "" {
				f.TokenID = &tokID
			}

			data, err := h.Usecase.FilterMkplaceNfts(ctx, f)
			if err != nil {
				logger.AtLog.Logger.Error("Nfts", zap.Any("iPagination", iPagination), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Nfts", zap.Any("iPagination", iPagination), zap.Any("data", len(data.Items)))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get marketplace Nfts of a collection
// @Description Get marketplace Nfts of a collection
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param rarity query string false "min,max - separated by comma"
// @Param price query string false "min,max - separated by comma"
// @Param attributes query string false "key:value,key:value - separated by comma ex: Base colour:Red,Base colour:Orange"
// @Param token_id query string false "token id"
// @Param owner query string false "owner"
// @Param contract_address path string true "contract_address"
// @Param is_big_file query bool false "true|false, default: all"
// @Param buyable query bool false "true|false, default: all"
// @Param limit query int false "limit"
// @Param sort_by query string false "sort by field: default volume"
// @Param sort query int false "sort default: -1 desc"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/collections/{contract_address}/nfts [GET]
func (h *httpDelivery) mkplaceNftsOfACollection(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)

			if p.SortBy == nil {
				sortBy := "token_id_int"
				p.SortBy = &sortBy
			}

			if p.Sort == nil {
				s := int(entity.SORT_DESC)
				p.Sort = &s
			}

			f := entity.FilterNfts{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					SortBy: *p.SortBy,
					Sort:   entity.SortType(*p.Sort),
				},
			}
			ca := strings.ToLower(vars["contract_address"])
			if ca != "" {
				f.ContractAddress = &ca
			}

			tokenID := strings.ToLower(r.URL.Query().Get("token_id"))
			if tokenID != "" {
				f.TokenID = &tokenID
			}

			owner := strings.ToLower(r.URL.Query().Get("owner"))
			if owner != "" {
				f.Owner = &owner
			}
			isOrPhan := strings.ToLower(r.URL.Query().Get("is_orphan"))
			if isOrPhan != "" {
				if isOrhanVal, err := strconv.Atoi(isOrPhan); err == nil {
					f.IsOrphan = &isOrhanVal
				}
			}

			rarity := strings.ToLower(r.URL.Query().Get("rarity"))
			if rarity != "" {
				rarity = strings.ReplaceAll(rarity, " ", "")
				rArray := strings.Split(rarity, ",")
				if len(rArray) == 2 {

					sort.SliceIsSorted(rArray, func(i, j int) bool {
						return rArray[i] > rArray[j]
					})

					min, _ := strconv.ParseFloat(rArray[0], 10)
					max, _ := strconv.ParseFloat(rArray[1], 10)

					f.Rarity = &entity.Rarity{
						Min: min,
						Max: max,
					}
				}
			}

			price := strings.ToLower(r.URL.Query().Get("price"))
			if price != "" {
				price = strings.ReplaceAll(price, " ", "")
				prArray := strings.Split(price, ",")
				if len(prArray) == 2 {

					min, _ := strconv.ParseFloat(prArray[0], 10)
					if min == -1 {
						min = 0
					}

					max, _ := strconv.ParseFloat(prArray[1], 10)
					if max == -1 {
						max = 999999999
					}

					f.Price = &entity.Rarity{
						Min: min,
						Max: max,
					}
				}
			}

			attributes := r.URL.Query().Get("attributes")
			if attributes != "" {
				attributeArr := strings.Split(attributes, ",")
				val := []string{}
				key := []string{}
				for _, attr := range attributeArr {
					sep := strings.Split(attr, ":")
					if len(sep) == 2 {
						key = append(key, sep[0])
						val = append(val, sep[1])
					}
				}
				f.AttrKey = key
				f.AttrValue = val
			}

			isBigFile := r.URL.Query().Get("is_big_file")
			if isBigFile != "" {
				isBigFileBool, err := strconv.ParseBool(isBigFile)
				if err == nil {
					f.IsBigFile = &isBigFileBool
				}
			}

			isBuyable := r.URL.Query().Get("buyable")
			if isBuyable != "" {
				isBuyableBool, err := strconv.ParseBool(isBuyable)
				if err == nil {
					f.IsBuyable = &isBuyableBool
				}
			}

			iWalletAddress := ctx.Value(utils.SIGNED_WALLET_ADDRESS)
			walletAddress, ok := iWalletAddress.(string)
			if ok {
				f.CurrentUser = &walletAddress
			}

			data, err := h.Usecase.FilterMkplaceNftNew(ctx, f)
			if err != nil {
				logger.AtLog.Logger.Error("can not get nfts", zap.Any("iPagination", iPagination), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Nfts", zap.Any("iPagination", iPagination),
				zap.Any("filter", f),
				zap.Any("data", len(data.Items)),
			)

			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get marketplace Nft's detail
// @Description Get marketplace Nft's detail
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract_address"
// @Param token_id path string true "token_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/collections/{contract_address}/nfts/{token_id} [GET]
func (h *httpDelivery) mkplaceNftDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			ca := vars["contract_address"]
			tokID := vars["token_id"]

			data, err := h.Usecase.GetMkplaceNft(ctx, ca, tokID)
			if err != nil {
				logger.AtLog.Logger.Error("Nfts", zap.Error(err))
				return nil, err
			}

			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Collection's attributes
// @Description  Get Collection's attributes
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param trait_type query string false "trait_type"
// @Param value query string false "value"
// @Param contract_address path string true "contract address"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/collections/{contract_address}/attributes [GET]
func (h *httpDelivery) mkpCollectionAttributes(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := vars["contract_address"]

			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			_ = p

			f := entity.FilterMarketplaceCollectionAttribute{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				ContractAddress: &contractAddress,
			}

			traitType := r.URL.Query().Get("trait_type")
			if traitType != "" {
				f.TraitType = &traitType
			}

			value := r.URL.Query().Get("value")
			if value != "" {
				f.Value = &value
			}

			data, err := h.Usecase.MarketplaceCollectionAttributes(ctx, f)
			if err != nil {
				logger.AtLog.Logger.Error("mkpCollectionAttributes", zap.String("contractAddress", contractAddress), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("mkpCollectionAttributes", zap.String("contractAddress", contractAddress), zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get collection's activities
// @Description Get collection's activities
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract_address"
// @Param types query string false "0: mint, 1: listing, 2: cancel listing, 3: token matched, default all"
// @Param limit query int false "limit default 10"
// @Param page query int false "page start with 1"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/collections/{contract_address}/activities [GET]
func (h *httpDelivery) getCollectionActivities(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)

			contractAddresss := vars["contract_address"]
			types := r.URL.Query().Get("types")

			f := entity.FilterTokenActivities{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				ContractAddress: &contractAddresss,
			}

			if types != "" {
				types = strings.ReplaceAll(types, " ", "")
				t := strings.Split(types, ",")
				for _, i := range t {
					iInt, err := strconv.Atoi(i)
					if err != nil {
						return nil, err
					}
					f.Types = append(f.Types, iInt)
				}

			}

			resp, err := h.Usecase.FilterTokenActivities(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get collection's chart
// @Description Get collection's chart
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract_address"
// @Param date_range query string false "date range: 7D, month - default 7D"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/collections/{contract_address}/chart [GET]
func (h *httpDelivery) getCollectionChart(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddresss := vars["contract_address"]
			dateRange := r.URL.Query().Get("date_range")

			to := time.Now().UTC()

			day := 7 * 24
			if dateRange == strings.ToLower("month") {
				day = 30 * 24
			}

			from := to.Add(time.Duration(day*-1) * time.Hour)

			f := entity.FilterCollectionChart{
				ContractAddress: &contractAddresss,
				FromDate:        &from,
				ToDate:          &to,
			}

			resp, err := h.Usecase.FilterCollectionChart(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get marketplace Nft owners of a collection
// @Description Get marketplace Nft owners of a collection
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract_address"
// @Param limit query int false "limit"
// @Param sort_by query string false "sort by field: default volume"
// @Param sort query int false "sort default: -1 desc"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/collections/{contract_address}/nft-owners [GET]
func (h *httpDelivery) mkplaceNftOwnerCollection(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)

			if p.SortBy == nil {
				sortBy := "count"
				p.SortBy = &sortBy
			}

			if p.Sort == nil {
				s := int(entity.SORT_DESC)
				p.Sort = &s
			}

			contractAddress := vars["contract_address"]
			f := entity.FilterCollectionNftOwners{
				ContractAddress: &contractAddress,
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Page:   int64(*p.Page),
					Offset: int64(*p.Offset),
					SortBy: *p.SortBy,
					Sort:   entity.SortType(*p.Sort),
				},
			}

			data, err := h.Usecase.FilterNftOwners(ctx, f)
			if err != nil {
				return nil, err
			}
			logger.AtLog.Logger.Info("Nfts", zap.Any("iPagination", iPagination), zap.Any("data", len(data.Items)))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get token's histories
// @Description Get token's histories
// @Tags MarketPlace
// @Accept  json
// @Produce  json
// @Param contract_address path string true "contract_address"
// @Param token_id path string true "token_id"
// @Param status query bool false "0: open, 1: cancel, 2: done, default all"
// @Param sort_by query string false "sort by field"
// @Param sort query int false "1: ASC, -1: DESC"
// @Param limit query int false "limit default 10"
// @Param page query int false "page start with 1"
// @Success 200 {object} response.JsonResponse{}
// @Router /marketplace/contract/{contract_address}/token/{token_id}/soul_histories [GET]
func (h *httpDelivery) getSoulHistories(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			tokenID := vars["token_id"]
			contractAddresss := vars["contract_address"]

			f := entity.FilterTokenActivities{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				TokenID:         &tokenID,
				ContractAddress: &contractAddresss,
			}

			resp, err := h.Usecase.FilterTokenSoulHistories(ctx, f)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	).ServeHTTP(w, r)
}
