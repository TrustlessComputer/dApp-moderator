package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"strings"
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

// uploadFile godoc
// @Summary Upload File multipart fake
// @Description  Upload File multipart fake
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Param file formData file true "file"
// @Success 200 {object} response.UploadResponse{}
// @Security ApiKeyAuth
// @Router /upload/file/multipart-fake [POST]
func (h *httpDelivery) uploadFileMultiPartFake(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			_, fileHeader, err := r.FormFile("file")
			if err != nil || fileHeader == nil {
				return nil, fmt.Errorf("invalid file data")
			}

			path, err := h.Usecase.UploadFileMultipart(fileHeader)
			if err != nil {
				return nil, err
			}

			return path, nil
		},
	).ServeHTTP(w, r)
}

// uploadFile godoc
// @Summary Get chunks of the uploaded file
// @Description Get chunks of the uploaded file
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Param file_id path string  true "fileID"
// @Param chunk_id path string  true "chunk_id"
// @Param tx_hash path string  true "tx_hash"
// @Param status query int false "0: new, 1: processing, 2: done - default: all"
// @Success 200 {object} response.UploadResponse{}
// @Security ApiKeyAuth
// @Router /upload/file/{file_id}/chunks [GET]
func (h *httpDelivery) fileChunks(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			fileID, err := primitive.ObjectIDFromHex(vars["file_id"])
			if err != nil {
				return nil, err
			}

			status := r.URL.Query().Get("status")

			f := &entity.FilterChunks{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				FileID: &fileID,
			}

			if status != "" {
				statusInt, err := strconv.Atoi(status)
				if err == nil {
					cs := entity.ChunkStatus(statusInt)
					f.Status = &cs
				}

			}

			chunks, err := h.Usecase.FilterChunks(f)
			if err != nil {
				return nil, err
			}

			return chunks, nil
		},
	).ServeHTTP(w, r)
}

// uploadFile godoc
// @Summary Get chunk by ID
// @Description Get chunk by ID
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Param file_id path string  true "fileID"
// @Param chunk_id path string  true "chunk_id"
// @Param status query int false "0: new, 1: processing, 2: done - default: all"
// @Success 200 {object} response.UploadResponse{}
// @Security ApiKeyAuth
// @Router /upload/file/{file_id}/chunks/{chunk_id} [GET]
func (h *httpDelivery) getChunkByID(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			chunkID := vars["chunk_id"]
			fileID := vars["file_id"]
			chunk, err := h.Usecase.GetChunkByID(fileID, chunkID)
			if err != nil {
				return nil, err
			}

			return chunk, nil
		},
	).ServeHTTP(w, r)
}

// uploadFile godoc
// @Summary Get uploaded Files
// @Description Get uploaded Files
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Param contract_address query string false "contract_address"
// @Param token_id query string false "token_id"
// @Param wallet_address query string false "wallet_address"
// @Param tx_hash query string false "tx_hash"
// @Param status query string false "0: new, 1: has tx_hash and not fully uploaded to blockchain, 2: done. Statuses are separated by comma"
// @Success 200 {object} response.UploadResponse{}
// @Router /upload/file [GET]
func (h *httpDelivery) filterUploadedFile(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)
			contractAddress := strings.ToLower(r.URL.Query().Get("contract_address"))
			tokenID := r.URL.Query().Get("token_id")
			walletAddress := strings.ToLower(r.URL.Query().Get("wallet_address"))
			txHash := strings.ToLower(r.URL.Query().Get("tx_hash"))

			status := []int{}
			statsStr := strings.ToLower(r.URL.Query().Get("status"))
			if statsStr != "" {
				statsStrs := strings.Split(statsStr, ",")
				for _, i := range statsStrs {
					iInt, err := strconv.Atoi(i)
					if err != nil {
						continue
					}
					status = append(status, iInt)
				}
			}

			f := &entity.FilterUploadedFile{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					//SortBy: *p.SortBy,
					//Sort:   entity.SortType(*p.Sort),
				},
				ContractAddress: &contractAddress,
				TokenID:         &tokenID,
				WalletAddress:   &walletAddress,
				TxHash:          &txHash,
				Status:          status,
			}

			files, err := h.Usecase.GetUploadedFiles(f)
			if err != nil {
				return nil, err
			}

			return files, nil
		},
	).ServeHTTP(w, r)
}

// uploadFile godoc
// @Summary update tx_hash for the uploaded file
// @Description update tx_hash for the uploaded file
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Param file_id path string true "file_id"
// @Param tx_hash path string true "tx_hash"
// @Param request_body body structure.UpdateUploadedFileTxHash true "request body"
// @Success 200 {object} response.UploadResponse{}
// @Security ApiKeyAuth
// @Router /upload/file/{file_id}/tx_hash/{tx_hash} [PUT]
func (h *httpDelivery) updateTxHashUploadedFile(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			fileID := vars["file_id"]
			txHash := vars["tx_hash"]

			reqBody := &structure.UpdateUploadedFileTxHash{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				return nil, err
			}

			reqBody.TxHash = txHash
			reqBody.FileID = fileID

			resp, err := h.Usecase.UpdateTxHashForUploadedFile(reqBody)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}

// CreateMultipartUpload godoc
// @Summary Create multipart upload
// @Description Create multipart upload.
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Security Authorization
// @Param request body request.CreateMultipartUploadRequest true "Create multipart upload request"
// @Success 200 {object} response.JsonResponse{}
// @Router /upload/file/multipart [POST]
func (h *httpDelivery) CreateMultipartUpload(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			var reqBody request.CreateMultipartUploadRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqBody)
			if err != nil {
				return nil, err
			}

			resp, err := h.Usecase.CreateMultipartUpload(ctx, reqBody.Group, reqBody.FileName)
			if err != nil {
				return nil, err
			}

			return response.FileResponse{UploadID: *resp}, nil
		},
	).ServeHTTP(w, r)
}

// UploadPart godoc
// @Summary Upload multipart file
// @Description Upload multipart file
// @Tags Uploader
// @Content-Type: multipart/form-data
// @Security Authorization
// @Produce  multipart/form-data
// @Param file formData file true "file"
// @Param uploadID path string true "upload ID"
// @Param partNumber query string  false  "part number"
// @Success 200 {object} response.JsonResponse{}
// @Router /upload/file/multipart/{uploadID} [PUT]
func (h *httpDelivery) UploadPart(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			_, handler, err := r.FormFile("file")
			if err != nil {
				return nil, err
			}

			uploadID := vars["uploadID"]
			var partNumber int
			partNumberStr := r.URL.Query().Get("partNumber")
			if partNumberStr == "" {
				err = errors.New("missing part number")
			} else {
				partNumber, err = strconv.Atoi(partNumberStr)
			}

			if err != nil {
				return nil, err
			}

			data, err := handler.Open()
			if err != nil {
				return nil, err
			}

			err = h.Usecase.UploadPart(ctx, uploadID, data, handler.Size, partNumber)
			if err != nil {
				return nil, err
			}

			return "OK", nil
		},
	).ServeHTTP(w, r)
}

// CompleteMultipartUpload godoc
// @Summary Finish multipart upload
// @Description Finish multipart upload
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Security Authorization
// @Param uploadID path string true "upload ID"
// @Success 200 {object} response.JsonResponse{}
// @Router /upload/file/multipart/{uploadID} [POST]
func (h *httpDelivery) CompleteMultipartUpload(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			uploadID := vars["uploadID"]
			uploaded, err := h.Usecase.CompleteMultipartUpload(ctx, uploadID)

			if err != nil {
				return nil, err
			}

			return response.MultipartUploadResponse{
				FileURL: uploaded.FullPath,
				FileID:  uploaded.ID.Hex(),
			}, nil
		},
	).ServeHTTP(w, r)
}

// uploadFile godoc
// @Summary update tx_hash for a chunk
// @Description update tx_hash for a chunk
// @Tags Uploader
// @Accept  json
// @Produce  json
// @Param file_id path string true "file_id"
// @Param chunk_id path string true "chunk_id"
// @Param tx_hash path string false "tx_hash"
// @Success 200 {object} response.UploadResponse{}
// @Security ApiKeyAuth
// @Router /upload/file/{file_id}/chunks/{chunk_id}/tx_hash/{tx_hash} [PUT]
func (h *httpDelivery) updateTxHashForAChunk(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			fileID := vars["file_id"]
			chunkID := vars["chunk_id"]
			txHash := vars["tx_hash"]

			resp, err := h.Usecase.UpdateTxHashForAChunk(fileID, chunkID, txHash)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}
