package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/response"
	"fmt"
	"net/http"
	"time"
)

// uploadFile godoc
// @Summary Upload File
// @Description Upload File
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Param file formData file true "file"
// @Success 200 {object} response.UploadResponse{}
// @Security ApiKeyAuth
// @Router /upload/file [POST]
func (h *httpDelivery) uploadFile(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			_, fileHeader, err := r.FormFile("file")
			if err != nil || fileHeader == nil {
				return nil, fmt.Errorf("invalid file data")
			}

			path, err := h.Usecase.UploadFile(fileHeader)
			if err != nil {
				return nil, err
			}

			return response.UploadResponse{
				FileName:  fileHeader.Filename,
				URL:       path,
				CreatedAt: time.Now(),
			}, nil
		},
	).ServeHTTP(w, r)
}
