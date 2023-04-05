package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/utils/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// UserCredits godoc
// @Summary Get files of a wallet
// @Description Get files of a wallet (uploader's wallet address)
// @Tags BFS-service
// @Accept  json
// @Produce  json
// @Param walletAddress path string true "walletAddress"
// @Success 200 {object} response.JsonResponse{}
// @Router /bfs-service/files/{walletAddress} [GET]
func (h *httpDelivery) bfsFiles(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["walletAddress"]
			data, err := h.Usecase.BfsFiles(ctx, walletAddress)
			if err != nil {
				logger.AtLog.Logger.Error("BfsFiles", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("BfsFiles", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Browse files of a wallet
// @Description Browse files of a wallet (uploader's wallet address)
// @Tags BFS-service
// @Accept  json
// @Produce  json
// @Param walletAddress path string true "walletAddress"
// @Param path query string false "path"
// @Success 200 {object} response.JsonResponse{}
// @Router /bfs-service/browse/{walletAddress} [GET]
func (h *httpDelivery) bfsBrowseFile(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["walletAddress"]
			path := r.URL.Query().Get("path")
			data, err := h.Usecase.BfsBrowsedFile(ctx, walletAddress, path)
			if err != nil {
				logger.AtLog.Logger.Error("bfsBrowseFile", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bfsBrowseFile", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get file info
// @Description Get file info of a wallet address (uploader's wallet address)
// @Tags BFS-service
// @Accept  json
// @Produce  json
// @Param walletAddress path string true "walletAddress"
// @Param path query string false "path"
// @Success 200 {object} response.JsonResponse{}
// @Router /bfs-service/info/{walletAddress} [GET]
func (h *httpDelivery) bfsFileInfo(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			walletAddress := vars["walletAddress"]
			path := r.URL.Query().Get("path")
			data, err := h.Usecase.BfsFileInfo(ctx, walletAddress, path)
			if err != nil {
				logger.AtLog.Logger.Error("bfsFileInfo", zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("bfsFileInfo", zap.Any("data", data))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get content file
// @Description Get file content of a wallet address (uploader's wallet address)
// @Tags BFS-service
// @Accept  json
// @Produce  json
// @Param walletAddress path string true "walletAddress"
// @Param path query string false "path"
// @Success 200 {object} response.JsonResponse{}
// @Router /bfs-service/content/{walletAddress} [GET]
func (h *httpDelivery) bfsFileContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctx := context.Background()
	
	walletAddress := vars["walletAddress"]
	path := r.URL.Query().Get("path")
	data, ctype, err := h.Usecase.BfsFileContent(ctx, walletAddress, path)
	if err != nil {
		logger.AtLog.Logger.Error("bfsFileInfo", zap.Error(err))
		return 
	}

	if err != nil {
		logger.AtLog.Logger.Error("collectionNftContent",  zap.Error(err))
		h.Response.RespondWithError(w, http.StatusBadRequest, response.Error, err)
		return 
	}

	logger.AtLog.Logger.Info("collectionNftContent", zap.Any("data", data))
	
	w.Header().Set("Content-Type", ctype)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)
	return 
}
