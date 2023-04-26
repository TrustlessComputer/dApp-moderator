package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"encoding/json"
	"net/http"
)

// UserCredits godoc
// @Summary post dapp info
// @Description update load dapp info
// @Tags dapp-service
// @Accept  json
// @Produce  json
// @Success 200 {object} response.JsonResponse{}
// @Router /dapp-info/create [GET]
func (h *httpDelivery) createDAppInfo(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			reqBody := &request.CreateDappInfoReq{}

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				return nil, err
			}

			resp, err := h.Usecase.ApiCreateDappInfo(ctx, reqBody)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) listDAppInfo(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			resp, err := h.Usecase.ApiListDappInfo()
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}
