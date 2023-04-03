package http

import (
	"net/http"

	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get Redis
// @Description Get Redis
// @Tags Admin
// @Accept  json
// @Produce  json
// @Success 200 {object} response.JsonResponse{data=[]response.RedisResponse}
// @Router /admin/redis [GET]
func (h *httpDelivery) getRedisKeys(w http.ResponseWriter, r *http.Request) {
	res, err := h.Usecase.GetAllRedis()

	if err != nil {
		logger.AtLog.Logger.Error("err", zap.Error(err))
		h.Response.RespondWithError(w, http.StatusBadRequest, response.Error, err)
		return
	}

	h.Response.RespondSuccess(w, http.StatusOK, response.Success, res, "")
}