package http

import (
	"context"
	"dapp-moderator/external/cyberscope"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"encoding/json"
	"net/http"
)

func (h *httpDelivery) checkEvmBytescode(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			var reqBody request.BytescodeRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqBody)
			if err != nil {
				return nil, err
			}
			res, err := cyberscope.CheckBytescode(reqBody.Bytescode)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	).ServeHTTP(w, r)
}
