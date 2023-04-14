package http

import (
	"net/http"

	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"

	"go.uber.org/zap"
)

func (h *httpDelivery) getRedisKeys(w http.ResponseWriter, r *http.Request) {
	res, err := h.Usecase.GetAllRedis()

	if err != nil {
		logger.AtLog.Logger.Error("err", zap.Error(err))
		h.Response.RespondWithError(w, http.StatusBadRequest, response.Error, err)
		return
	}

	h.Response.RespondSuccess(w, http.StatusOK, response.Success, res, "")
}
