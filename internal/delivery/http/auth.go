package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/logger"
	req "dapp-moderator/utils/request"
	"encoding/json"
	"errors"
	"fmt"
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

// @Summary  Current user profile
// @Description Current user profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResponse{}
// @Router /profile/me [GET]
func (h *httpDelivery) currentUerProfile(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iwalletAdress := ctx.Value(utils.SIGNED_WALLET_ADDRESS)
			walletAdress, ok := iwalletAdress.(string)
			if !ok {
				err := errors.New("Token is incorect")
				logger.AtLog.Logger.Error("currentUerProfile", zap.String("walletAdress", walletAdress), zap.Error(err))
				return nil, err
			}

			profile, err := h.Usecase.GetUserProfileByWalletAddress(walletAdress)
			if err != nil {
				profile, err = h.Usecase.GetUserProfileByBtcAddressTaproot(walletAdress)
				if err != nil {
					logger.AtLog.Logger.Error("currentUerProfile failed", zap.String("walletAdress", walletAdress), zap.Error(err))
					profile = &entity.Users{}
				}
			}

			return profile, nil
		},
	).ServeHTTP(w, r)
}

// @Summary  Create profile's history
// @Description Create profile's history
// @Tags Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body structure.CreateHistoryMessage true "Generate message request"
// @Success 200 {object} response.JsonResponse{}
// @Router /profile/histories [POST]
func (h *httpDelivery) createProfileHistory(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iwalletAdress := ctx.Value(utils.SIGNED_WALLET_ADDRESS)
			walletAdress, ok := iwalletAdress.(string)
			if !ok {
				err := errors.New("Token is incorect")
				logger.AtLog.Logger.Error("createProfileHistory", zap.String("walletAdress", walletAdress), zap.Error(err))
				return nil, err
			}

			reqBody := &structure.CreateHistoryMessage{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				return nil, err
			}
			reqBody.WalletAddress = walletAdress
			resp, err := h.Usecase.CreateUserHistory(ctx, reqBody)
			if err != nil {
				logger.AtLog.Logger.Error("createProfileHistory", zap.String("walletAdress", walletAdress), zap.Error(err))
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}

// @Summary  confirm profile's history
// @Description confirm profile's history
// @Tags Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param req body request.ConfirmHistoriesReq true "request"
// @Success 200 {object} response.JsonResponse{}
// @Router /profile/histories [PUT]
func (h *httpDelivery) confirmProfileHistory(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iwalletAdress := ctx.Value(utils.SIGNED_WALLET_ADDRESS)
			walletAdress, ok := iwalletAdress.(string)
			if !ok {
				err := errors.New("Token is incorect")
				logger.AtLog.Logger.Error("confirmProfileHistory", zap.String("walletAdress", walletAdress), zap.Error(err))
				return nil, err
			}

			var reqBody request.ConfirmHistoriesReq
			err := req.BindJson(r, &reqBody)

			if err != nil {
				return nil, err
			}

			resp, err := h.Usecase.ConfirmUserHistory(ctx, walletAdress, &reqBody)
			if err != nil {
				logger.AtLog.Logger.Error("confirmProfileHistory", zap.String("walletAdress", walletAdress), zap.Error(err))
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}

// @Summary  Current user collections
// @Description Current user collections (created collections and collection has the owned nft)
// @Tags Profile
// @Accept json
// @Produce json
// @Param contract query string false "contract"
// @Param name query string false "name"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param sort_by query string false "default deployed_at_block"
// @Param sort query int false "default -1"
// @Success 200 {object} response.JsonResponse{}
// @Param walletAddress path string true "Wallet address"
// @Router /profile/wallet/{walletAddress}/collections [GET]
func (h *httpDelivery) currentUserProfileCollections(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAdress := vars["walletAddress"]

			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			var err error

			collectionAddress := r.URL.Query().Get("contract")
			name := r.URL.Query().Get("name")
			filter := request.CollectionsFilter{
				Owner:         &walletAdress,
				Address:       &collectionAddress,
				Name:          &name,
				PaginationReq: p,
			}

			nfts, err := h.Usecase.UserCollections(ctx, filter)
			if err != nil {
				logger.AtLog.Logger.Error("currentUerProfileCollections", zap.String("walletAdress", walletAdress), zap.Error(err), zap.Any("filter", filter))
				return nil, err
			}
			logger.AtLog.Logger.Error("currentUerProfileCollections", zap.String("walletAdress", walletAdress), zap.Error(err), zap.Int("resp", len(nfts)))
			return nfts, nil
		},
	).ServeHTTP(w, r)
}

// @Summary  Current user bouhght-tokens
// @Description Current user bouhght-tokens  (the tokens that the user has spent)
// @Tags Profile
// @Accept json
// @Produce json
// @Param walletAddress path string true "Wallet address"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param sort_by query string false "default token_id_int"
// @Param sort query int false "default -1"
// @Success 200 {object} response.JsonResponse{}
// @Router /profile/wallet/{walletAddress}/tokens/bought [GET]
func (h *httpDelivery) currentUserProfileBoughtTokens(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			pagination, ok := iPagination.(request.PaginationReq)
			if !ok {
				err := fmt.Errorf("invalid pagination params")
				logger.AtLog.Logger.Error("currentUerProfileBoughtTokens", zap.Error(err))
				return nil, err
			}

			walletAdress := vars["walletAddress"]
			data, err := h.Usecase.FindTokens(ctx, pagination, req.Query(r, "key", walletAdress))
			if err != nil {
				logger.AtLog.Logger.Error("currentUerProfileBoughtTokens", zap.Any("pagination", pagination), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("currentUerProfileBoughtTokens", zap.Any("pagination", pagination))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// @Summary  Current user histories
// @Description Current user histories
// @Tags Profile
// @Accept json
// @Produce json
// @Param tx_hash query string false "tx_hash"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param sort_by query string false "default deployed_at_block"
// @Param sort query int false "default -1"
// @Success 200 {object} response.JsonResponse{}
// @Param walletAddress path string true "Wallet address"
// @Router /profile/wallet/{walletAddress}/histories [GET]
func (h *httpDelivery) currentUerProfileHistories(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAdress := vars["walletAddress"]

			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			var err error
			txHash := r.URL.Query().Get("tx_hash")
			filter := request.HistoriesFilter{
				WalletAdress:  &walletAdress,
				PaginationReq: p,
				TxHash:        &txHash,
			}

			h, err := h.Usecase.GetUserHistories(ctx, filter)
			if err != nil {
				logger.AtLog.Logger.Error("currentUerProfileHistories", zap.Any("filter", filter), zap.Error(err))
				return nil, err

			}

			logger.AtLog.Logger.Info("currentUerProfileHistories", zap.Any("filter", filter), zap.Int("data", len(h)))
			return h, nil
		},
	).ServeHTTP(w, r)
}
