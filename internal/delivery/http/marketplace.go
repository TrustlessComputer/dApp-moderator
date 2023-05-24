package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"net/http"
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
