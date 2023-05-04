package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"net/http"
)

// UserCredits godoc
// @Summary Get bns names
// @Description Get bns names
// @Tags BNS-service
// @Accept  json
// @Produce  json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /bns-service/names [GET]
func (h *httpDelivery) compileContract(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			data, err := h.Usecase.CompileContract(r)
			if err != nil {
				return nil, err
			}

			return data, nil
		},
	).ServeHTTP(w, r)
}
