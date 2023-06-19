package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"net/http"
)

func (h *httpDelivery) auctionDetail(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			r.URL.Query().Get("auction_id")
			r.Context().Value()
			return nil, nil
		},
	).ServeHTTP(w, r)
}
