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
// @Param token_id query string false "token_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /auctions [GET]
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

func (h *httpDelivery) listBid(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
		iPagination := ctx.Value(utils.PAGINATION)
		pagination, ok := iPagination.(request.PaginationReq)
		if !ok {
			err := fmt.Errorf("invalid pagination params")
			logger.AtLog.Logger.Error("invalid pagination params", zap.Error(err))
			return nil, err
		}

		dbAuction := vars["dbAuctionID"]
		if dbAuction == "" {
			return nil, errors.New("missing required info")
		}
		return h.Usecase.AuctionListBid(dbAuction, &pagination)
	}).ServeHTTP(w, r)
}
