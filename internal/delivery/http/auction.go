package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"errors"
	"net/http"
	"strings"
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

func (h *httpDelivery) listBidByAuction(w http.Request, r *http.Request) {

}
