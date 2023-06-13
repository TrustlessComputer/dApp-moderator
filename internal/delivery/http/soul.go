package http

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/delivery/http/response"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

// UserCredits godoc
// @Summary Get Soul's Nfts
// @Description Soul's Nfts
// @Tags Soul
// @Accept  json
// @Produce  json
// @Param rarity query string false "min,max - separated by comma"
// @Param price query string false "min,max - separated by comma"
// @Param attributes query string false "key:value,key:value - separated by comma ex: Base colour:Red,Base colour:Orange"
// @Param token_id query string false "token id"
// @Param is_big_file query bool false "true|false, default: all"
// @Param buyable query bool false "true|false, default: all"
// @Param limit query int false "limit"
// @Param sort_by query string false "sort by field: default volume"
// @Param sort query int false "sort default: -1 desc"
// @Param page query int false "page"
// @Success 200 {object} response.JsonResponse{}
// @Router /soul/nfts [GET]
func (h *httpDelivery) soulNfts(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			iPagination := ctx.Value(utils.PAGINATION)
			p := iPagination.(request.PaginationReq)

			if p.SortBy == nil {
				sortBy := "token_id_int"
				p.SortBy = &sortBy
			}

			if p.Sort == nil {
				s := int(entity.SORT_DESC)
				p.Sort = &s
			}

			f := entity.FilterNfts{
				BaseFilters: entity.BaseFilters{
					Limit:  int64(*p.Limit),
					Offset: int64(*p.Offset),
					SortBy: *p.SortBy,
					Sort:   entity.SortType(*p.Sort),
				},
			}

			ca := os.Getenv("SOUL_CONTRACT")
			f.ContractAddress = &ca

			tokenID := strings.ToLower(r.URL.Query().Get("token_id"))
			if tokenID != "" {
				f.TokenID = &tokenID
			}

			rarity := strings.ToLower(r.URL.Query().Get("rarity"))
			if rarity != "" {
				rarity = strings.ReplaceAll(rarity, " ", "")
				rArray := strings.Split(rarity, ",")
				if len(rArray) == 2 {

					sort.SliceIsSorted(rArray, func(i, j int) bool {
						return rArray[i] > rArray[j]
					})

					min, _ := strconv.ParseFloat(rArray[0], 10)
					max, _ := strconv.ParseFloat(rArray[1], 10)

					f.Rarity = &entity.Rarity{
						Min: min,
						Max: max,
					}
				}
			}

			price := strings.ToLower(r.URL.Query().Get("price"))
			if price != "" {
				price = strings.ReplaceAll(price, " ", "")
				prArray := strings.Split(price, ",")
				if len(prArray) == 2 {

					sort.SliceIsSorted(prArray, func(i, j int) bool {
						return prArray[i] > prArray[j]
					})

					min, _ := strconv.ParseFloat(prArray[0], 10)
					max, _ := strconv.ParseFloat(prArray[1], 10)

					f.Price = &entity.Rarity{
						Min: min,
						Max: max,
					}
				}
			}

			attributes := r.URL.Query().Get("attributes")
			if attributes != "" {
				attributeArr := strings.Split(attributes, ",")
				val := []string{}
				key := []string{}
				for _, attr := range attributeArr {
					sep := strings.Split(attr, ":")
					if len(sep) == 2 {
						key = append(key, sep[0])
						val = append(val, sep[1])
					}
				}
				f.AttrKey = key
				f.AttrValue = val
			}

			isBigFile := r.URL.Query().Get("is_big_file")
			if isBigFile != "" {
				isBigFileBool, err := strconv.ParseBool(isBigFile)
				if err == nil {
					f.IsBigFile = &isBigFileBool
				}
			}

			isBuyable := r.URL.Query().Get("buyable")
			if isBuyable != "" {
				isBuyableBool, err := strconv.ParseBool(isBuyable)
				if err == nil {
					f.IsBuyable = &isBuyableBool
				}
			}

			data, err := h.Usecase.FilterSoulNfts(ctx, f)
			bnsAddress := strings.ToLower(os.Getenv("BNS_ADDRESS"))
			for _, i := range data {
				if i.Name == "" {
					if bnsAddress == ca {
						key := helpers.BnsTokenNameKey(i.TokenID)
						existed, _ := h.Usecase.Cache.Exists(key)
						if existed != nil && *existed == true {
							cached, _ := h.Usecase.Cache.GetData(key)
							if cached != nil {
								i.Name = *cached
							}
						} else {
							bnsName, _ := h.Usecase.BnsService.NameByToken(i.TokenID)
							if bnsName != nil {
								i.Name = bnsName.Name
								h.Usecase.Cache.SetStringData(key, i.Name)
							}

						}
					} else {
						tokenIdBigInt, _ := new(big.Int).SetString(i.TokenID, 10)
						g, _ := new(big.Int).SetString("1000000", 10)
						if tokenIdBigInt.Cmp(g) > 0 {
							// TODO maybe from generative
							i.Name = i.Collection.Name + " #" + tokenIdBigInt.Mod(tokenIdBigInt, g).String()
						} else {
							i.Name = i.Collection.Name + " #" + i.TokenID
						}
					}
				}
			}
			if err != nil {
				logger.AtLog.Logger.Error("Nfts", zap.Any("iPagination", iPagination), zap.Error(err))
				return nil, err
			}

			logger.AtLog.Logger.Info("Nfts", zap.Any("iPagination", iPagination), zap.Any("data", len(data)))
			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Get Soul's Nft
// @Description Soul's Nft
// @Tags Soul
// @Accept  json
// @Produce  json
// @Param token_id path string true "token_id"
// @Success 200 {object} response.JsonResponse{}
// @Router /soul/nfts/{token_id} [GET]
func (h *httpDelivery) soulNft(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			ca := strings.ToLower(os.Getenv("SOUL_CONTRACT"))
			tokID := vars["token_id"]

			data, err := h.Usecase.CollectionNftDetail(ctx, ca, tokID)
			if err != nil {
				logger.AtLog.Logger.Error("Nfts", zap.Error(err))
				return nil, err
			}

			return data, nil
		},
	).ServeHTTP(w, r)
}

// UserCredits godoc
// @Summary Create signature
// @Description Create signature
// @Tags Soul
// @Accept  json
// @Produce  json
// @Param requestdata body int true "request data"
// @Success 200 {object} response.JsonResponse{}
// @Router /soul/signature [POST]
func (h *httpDelivery) SoulCreateSignature(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			return nil, nil
		},
	).ServeHTTP(w, r)
}
