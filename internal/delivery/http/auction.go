package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"errors"
	"net/http"
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
			auctionID := r.URL.Query().Get("auction_id")
			if auctionID == "" {
				return nil, errors.New("auction_id is required")
			}

			return nil, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) listBidByAuction(w http.Request, r *http.Request) {

}
