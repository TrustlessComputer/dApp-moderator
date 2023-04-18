package http

import (
	"context"
	"encoding/json"
	"net/http"

	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
)

// UserCredits godoc
// @Summary Get Redis
// @Description Get Redis
// @Tags Admin
// @Accept  json
// @Produce  json
// @Success 200 {object} response.JsonResponse{data=response.RedisResponse}
// @Router /admin/redis [GET]
func (h *httpDelivery) getRedisKeys(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			res, err := h.Usecase.GetAllRedis()
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Redis
// @Description Get Redis
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param key path string true "Redis key"
// @Success 200 {object} response.JsonResponse{data=response.RedisResponse}
// @Router /admin/redis/{key} [GET]
func (h *httpDelivery) getRedis(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			redisKey := vars["key"]
			res, err := h.Usecase.GetRedis(redisKey)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Upsert Redis
// @Description Upsert Redis
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param request body request.UpsertRedisRequest true "Upsert redis key"
// @Success 200 {object} response.JsonResponse{data=response.RedisResponse}
// @Router /admin/redis [POST]
func (h *httpDelivery) upsertRedis(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			var reqBody request.UpsertRedisRequest
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&reqBody)
			if err != nil {
				return nil, err
			}
			res, err := h.Usecase.UpsertRedis(reqBody.Key, reqBody.Value)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Delete Redis
// @Description Delete Redis
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param key path string true "Redis key"
// @Success 200 {object} response.JsonResponse{data=string}
// @Router /admin/redis/{key} [DELETE]
func (h *httpDelivery) deleteRedis(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			redisKey := vars["key"]
			err := h.Usecase.DeleteRedis(redisKey)
			if err != nil {
				return "error", err
			}

			return "success", nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Delete Redis
// @Description Delete Redis
// @Tags Admin
// @Accept  json
// @Produce  json
// @Success 200 {object} response.JsonResponse{data=string}
// @Router /admin/redis [DELETE]
func (h *httpDelivery) deleteAllRedis(w http.ResponseWriter, r *http.Request) {

	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			_, err := h.Usecase.DeleteAllRedis()
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
	).ServeHTTP(w, r)
}
