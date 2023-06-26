package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// @Summary auctionDetail
// @Description auctionDetail
// @Tags Auction
// @Accept  json
// @Produce  json
// @Param contractAddress path string true "contract address"
// @Param tokenID path string true "token_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /auction/detail/{contractAddress}/{tokenID} [GET]
func (h *httpDelivery) auctionDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			contractAddress := strings.ToLower(vars["contractAddress"])
			tokenID := strings.ToLower(vars["tokenID"])
			if contractAddress == "" || tokenID == "" {
				return nil, errors.New("missing required info")
			}

			return h.Usecase.AuctionDetail(contractAddress, tokenID)
		},
	).ServeHTTP(w, r)
}

// @Summary listBid
// @Description listBid
// @Tags Auction
// @Accept  json
// @Produce  json
// @Param dbAuctionID query string false "DB Auction ID"
// @Param owner query string false "Owner"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /auction/list-bid [GET]
func (h *httpDelivery) listBid(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
		iPagination := ctx.Value(utils.PAGINATION)
		pagination, ok := iPagination.(request.PaginationReq)
		if !ok {
			err := fmt.Errorf("invalid pagination params")
			logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
			return nil, err
		}

		filterRequest := &request.FilterAuctionBid{
			PaginationReq: pagination,
		}
		if dbAuctionID := r.URL.Query().Get("dbAuctionID"); dbAuctionID != "" {
			filterRequest.DBAuctionID = &dbAuctionID
		}
		if owner := r.URL.Query().Get("owner"); owner != "" {
			filterRequest.Sender = &owner
		}

		return h.Usecase.AuctionListBid(filterRequest)
	}).ServeHTTP(w, r)
}
