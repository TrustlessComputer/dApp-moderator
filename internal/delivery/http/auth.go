package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Generate a message
// @Description Generate a message for user's wallet
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body structure.GenerateMessage true "Generate message request"
// @Success 200 {object} response.JsonResponse{}
// @Router /auth/nonce [POST]
func (h *httpDelivery) generateMessage(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			reqBody := &structure.GenerateMessage{}

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				return nil, err
			}

			resp, err := h.Usecase.GenerateMessage(ctx, reqBody)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}


// UserCredits godoc
// @Summary Verified the generated message
// @Description Verified the generated message
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body structure.VerifyMessage true "Verify message request"
// @Success 200 {object} response.JsonResponse{}
// @Router /auth/nonce/verify [POST]
func (h *httpDelivery) verifyMessage(w http.ResponseWriter, r *http.Request) {

	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			reqBody := &structure.VerifyMessage{}

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				return nil, err
			}

			resp, err := h.Usecase.VerifyMessage(ctx, reqBody)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}


// @Summary User profile via wallet address
// @Description User profile via wallet address
// @Tags Profile
// @Accept json
// @Produce json
// @Param walletAddress path string true "Wallet address"
// @Success 200 {object} response.JsonResponse{}
// @Router /profile/wallet/{walletAddress} [GET]
func (h *httpDelivery) profileByWallet(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["walletAddress"]
			profile, err := h.Usecase.GetUserProfileByWalletAddress(walletAddress)
			if err != nil {
				profile, err = h.Usecase.GetUserProfileByBtcAddressTaproot(walletAddress)
				if err != nil {
					logger.AtLog.Logger.Error("GetUserProfileByWalletAddress failed", zap.Error(err))
					profile = &entity.Users{}
				}
			}
			
			return profile, nil
		},
	).ServeHTTP(w, r)
}