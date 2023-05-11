package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *Usecase) TcSwapFindSwapPairs(ctx context.Context, filter request.PaginationReq, fromToken string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapPairFilter{}
	query.FromPagination(filter)
	if fromToken != "" {
		query.FromToken = fromToken
	}

	data, err = u.Repo.FindSwapPairs(ctx, query)

	if err != nil {
		logger.AtLog.Logger.Error("FindSwapPairs", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("FindSwapPairs", zap.Any("data", data))
	return data, nil
}

func (u *Usecase) TcSwapFindSwapHistories(ctx context.Context, filter request.PaginationReq,
	tokenAddress, pariAddress, userAddress string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapPairSwapHistoriesFilter{}
	query.FromPagination(filter)
	query.Token = tokenAddress
	query.ContractAddress = pariAddress
	query.UserAddress = userAddress

	data, err = u.Repo.FindSwapPairHistories(ctx, query)

	if err != nil {
		logger.AtLog.Logger.Error("TcSwapFindSwapHistories", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("TcSwapFindSwapHistories", zap.Any("data", data))
	return data, nil
}

func (u *Usecase) FindTokensInPool(ctx context.Context, filter request.PaginationReq, fromToken string) (interface{}, error) {
	var err error
	query := entity.TokenFilter{}
	query.FromPagination(filter)

	contracts := []string{}
	pairQuery := entity.SwapPairFilter{}
	pairQuery.Limit = 10000
	pairQuery.Page = 1

	pairs, err := u.Repo.FindSwapPairs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
		return nil, err
	}

	isWbtcInArray := false
	wbtcContractAddr := u.Repo.ParseConfigByString(ctx, "wbtc_contract_address")
	for _, pair := range pairs {
		if fromToken == "" || (fromToken != "" && strings.EqualFold(fromToken, pair.Token1)) {
			contracts = append(contracts, pair.Token0)
			if strings.EqualFold(wbtcContractAddr, pair.Token0) {
				isWbtcInArray = true
			}
		}

		if fromToken == "" || (fromToken != "" && strings.EqualFold(fromToken, pair.Token0)) {
			contracts = append(contracts, pair.Token1)
			if strings.EqualFold(wbtcContractAddr, pair.Token1) {
				isWbtcInArray = true
			}
		}
	}

	if isWbtcInArray && fromToken != "" {
		for _, pair := range pairs {
			if strings.EqualFold(wbtcContractAddr, pair.Token0) &&
				!strings.EqualFold(fromToken, pair.Token1) {
				contracts = append(contracts, pair.Token1)
			} else if strings.EqualFold(wbtcContractAddr, pair.Token1) &&
				!strings.EqualFold(fromToken, pair.Token0) {
				contracts = append(contracts, pair.Token0)
			}
		}
	}

	tokens := []*entity.Token{}
	if len(contracts) > 0 {
		tokens, err = u.Repo.FindTokensInPoolByContracts(ctx, contracts, query)
		if err != nil {
			logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
			return nil, err
		}
	}

	logger.AtLog.Logger.Info("FindTokensInPool", zap.Any("data", tokens))
	return tokens, nil
}

func (u *Usecase) ClearCache() error {
	redisKey := fmt.Sprintf("tc-swap:token-reports-%!s(int64=1)-%!s(int64=500)")
	u.Cache.Delete(redisKey)
	return nil

}
func (u *Usecase) FindTokensPrice(ctx context.Context, contractAddress string, chartType string) (interface{}, error) {
	reports, err := u.Repo.FindTokePrice(ctx, contractAddress, chartType)
	if err != nil {
		//logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
		return reports, nil
	}
	btcPrice := u.Repo.ParseConfigByFloat64(ctx, "swap_btc_price")

	for _, item := range reports {
		if s, err := strconv.ParseFloat(item.Close.String(), 64); err == nil {
			item.BtcPrice = s

			item.UsdPrice = fmt.Sprint(s * btcPrice)
			item.CloseUsd = fmt.Sprint(s * btcPrice)
		}
		if s, err := strconv.ParseFloat(item.Open.String(), 64); err == nil {
			item.OpenUsd = fmt.Sprint(s * btcPrice)
		}
		if s, err := strconv.ParseFloat(item.High.String(), 64); err == nil {
			item.HighUsd = fmt.Sprint(s * btcPrice)
		}
		if s, err := strconv.ParseFloat(item.Low.String(), 64); err == nil {
			item.LowUsd = fmt.Sprint(s * btcPrice)
		}
		if s, err := strconv.ParseFloat(item.VolumeTo.String(), 64); err == nil {
			item.VolumeToUsd = fmt.Sprint(s * btcPrice)
		}
		if s, err := strconv.ParseFloat(item.VolumeFrom.String(), 64); err == nil {
			item.VolumeFromUsd = fmt.Sprint(s * btcPrice)
		}
		item.TotalVolumeUsd = fmt.Sprint(item.TotalVolume * btcPrice)
	}
	return reports, nil
}

func (u *Usecase) FindTokensReport(ctx context.Context, filter request.PaginationReq, address, sortBy string, sortType int) (interface{}, error) {
	query := entity.TokenReportFilter{}
	query.FromPagination(filter)
	query.Address = address
	query.SortBy = sortBy
	query.SortType = sortType

	redisKey := fmt.Sprintf("tc-swap:token-reports-%d-%d-%s-%s-%d", query.Page, query.Limit, address, sortBy, sortType)
	exists, err := u.Cache.Exists(redisKey)
	if err != nil {
		logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
		return 0, err
	}

	if *exists {
		dataInCache, err := u.Cache.GetData(redisKey)
		if err != nil {
			logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
			return nil, err
		}

		b := []byte(*dataInCache)
		reports := []entity.SwapPairReport{}
		err = json.Unmarshal(b, &reports)
		if err != nil {
			return nil, err
		}
		return reports, nil
	} else {
		reports, err := u.Repo.FindTokenReport(ctx, query)
		if err != nil {
			logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
			return nil, err
		}

		btcPrice := u.Repo.ParseConfigByFloat64(ctx, "swap_btc_price")

		for _, item := range reports {
			if s, err := strconv.ParseFloat(item.Price.String(), 64); err == nil {
				item.BtcPrice = s
				item.UsdPrice = s * btcPrice
			}

			if s, err := strconv.ParseFloat(item.Volume.String(), 64); err == nil {
				item.BtcVolume = s
				item.UsdVolume = s * btcPrice
			}

			if s, err := strconv.ParseFloat(item.TotalVolume.String(), 64); err == nil {
				item.BtcTotalVolume = s
				item.UsdTotalVolume = s * btcPrice
			}

			if item.Address == "0xfB83c18569fB43f1ABCbae09Baf7090bFFc8CBBD" {
				item.UsdPrice = btcPrice
			}
		}

		logger.AtLog.Logger.Info("FindTokensReport", zap.Any("data", reports))
		reportsStr, err := json.Marshal(&reports)
		if err != nil {
			logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
			return reports, nil
		}
		err = u.Cache.SetStringDataWithExpTime(redisKey, string(reportsStr), 5*60)
		if err != nil {
			logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
			return reports, nil
		}
		return reports, nil
	}
}

func (u *Usecase) UpdateDataSwapSync(ctx context.Context) error {
	pairQuery := entity.SwapPairSyncFilter{}
	pairQuery.Limit = 1000
	pairQuery.Page = 1

	pairSyncs, err := u.Repo.FindSwapPairSyncs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
		return err
	}

	mapPair := map[string]*entity.SwapPair{}
	// mapToken := map[string]*entity.Token{}

	// wbtcContractAddr := u.Repo.ParseConfigByString(ctx, "wbtc_contract_address")
	for _, pairSync := range pairSyncs {
		// if pairSync != nil && pairSync.Pair == nil {
		// var token *entity.Token
		var pair *entity.SwapPair

		if p, ok := mapPair[strings.ToLower(pairSync.ContractAddress)]; ok {
			pair = p
		} else {
			pair, _ = u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
				Pair: strings.ToLower(pairSync.ContractAddress),
			})
			mapPair[strings.ToLower(pairSync.ContractAddress)] = pair
		}

		// if pair != nil {
		// 	tokenAddress := ""
		// 	if strings.EqualFold(pair.Token0, wbtcContractAddr) {
		// 		tokenAddress = pair.Token1
		// 	} else if strings.EqualFold(pair.Token1, wbtcContractAddr) {
		// 		tokenAddress = pair.Token0
		// 	}

		// 	if tokenAddress != "" {
		// 		if p, ok := mapToken[tokenAddress]; ok {
		// 			token = p
		// 		} else {
		// 			token, _ = u.Repo.FindToken(ctx, entity.TokenFilter{
		// 				Address: tokenAddress,
		// 			})
		// 			mapToken[tokenAddress] = token
		// 		}
		// 	}
		// }

		if pair != nil {
			// pairSync.Token = token.Address
			// tmpReserce0, _ := new(big.Float).SetString(pairSync.Reserve0.String())
			// tmpReserce1, _ := new(big.Float).SetString(pairSync.Reserve1.String())
			// tmpPrice := big.NewFloat(0).Quo(tmpReserce0, tmpReserce1)
			// if strings.EqualFold(pair.Token1, wbtcContractAddr) {
			// 	tmpPrice = big.NewFloat(0).Quo(tmpReserce1, tmpReserce0)
			// }
			// pairSync.Price, _ = primitive.ParseDecimal128(tmpPrice.String())
			pairSync.Pair = pair

			err := u.Repo.UpdateSwapPairSync(ctx, pairSync)
			if err != nil {
				logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
				return err
			}
		}
		// }
	}
	return nil
}

func (u *Usecase) UpdateDataSwapHistory(ctx context.Context) error {
	pairQuery := entity.SwapPairSwapHistoriesFilter{}
	pairQuery.Limit = 1000
	pairQuery.Page = 1

	pairSyncs, err := u.Repo.FindSwapPairHistories(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
		return err
	}

	mapPair := map[string]*entity.SwapPair{}
	// mapToken := map[string]*entity.Token{}

	// wbtcContractAddr := u.Repo.ParseConfigByString(ctx, "wbtc_contract_address")
	for _, pairSync := range pairSyncs {
		// if pairSync != nil && pairSync.Pair == nil {
		// var token *entity.Token
		var pair *entity.SwapPair

		if p, ok := mapPair[strings.ToLower(pairSync.ContractAddress)]; ok {
			pair = p
		} else {
			pair, _ = u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
				Pair: strings.ToLower(pairSync.ContractAddress),
			})
			mapPair[strings.ToLower(pairSync.ContractAddress)] = pair
		}

		// if pair != nil {
		// 	tokenAddress := ""
		// 	if strings.EqualFold(pair.Token0, wbtcContractAddr) {
		// 		tokenAddress = pair.Token1
		// 	} else if strings.EqualFold(pair.Token1, wbtcContractAddr) {
		// 		tokenAddress = pair.Token0
		// 	}

		// 	if tokenAddress != "" {
		// 		if p, ok := mapToken[tokenAddress]; ok {
		// 			token = p
		// 		} else {
		// 			token, _ = u.Repo.FindToken(ctx, entity.TokenFilter{
		// 				Address: tokenAddress,
		// 			})
		// 			mapToken[tokenAddress] = token
		// 		}
		// 	}
		// }

		if pair != nil {
			// pairSync.Token = token.Address
			// tmpAmount0In, _ := new(big.Float).SetString(pairSync.Amount0In.String())
			// tmpAmount0Out, _ := new(big.Float).SetString(pairSync.Amount0Out.String())
			// tmpAmount1In, _ := new(big.Float).SetString(pairSync.Amount1In.String())
			// tmpAmount1Out, _ := new(big.Float).SetString(pairSync.Amount1Out.String())

			// tmpAmount0 := big.NewFloat(0).Add(tmpAmount0In, tmpAmount0Out)
			// tmpAmount1 := big.NewFloat(0).Add(tmpAmount1In, tmpAmount1Out)

			// tmpVolume := tmpAmount0
			// tmpPrice := big.NewFloat(0).Quo(tmpAmount0, tmpAmount1)
			// if strings.EqualFold(pair.Token1, wbtcContractAddr) {
			// 	tmpVolume = tmpAmount1
			// 	tmpPrice = big.NewFloat(0).Quo(tmpAmount1, tmpAmount0)
			// }

			// pairSync.Volume, _ = primitive.ParseDecimal128(tmpVolume.String())
			// pairSync.Price, _ = primitive.ParseDecimal128(tmpPrice.String())
			pairSync.Pair = pair

			err := u.Repo.UpdateSwapPairHistory(ctx, pairSync)
			if err != nil {
				logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
				return err
			}
		}

		// }
	}
	return nil
}

func (u *Usecase) SwapGetPairApr(ctx context.Context, pair string) (interface{}, error) {
	var err error
	aprPercent := float64(0)
	query := entity.SwapPairFilter{}
	query.Id = strings.ToLower(pair)

	pairObj, err := u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
		Pair: strings.ToLower(pair),
	})
	if err != nil {
		logger.AtLog.Logger.Error("SwapGetPairApr", zap.Error(err))
		return nil, err
	}

	if pairObj != nil {
		wbtcContractAddr := u.Repo.ParseConfigByString(ctx, "wbtc_contract_address")
		pairVolume, err := u.Repo.FindSwapPairVolume(ctx, query)
		if err != nil && err != mongo.ErrNoDocuments {
			logger.AtLog.Logger.Error("SwapGetPairApr", zap.Error(err))
			return nil, err
		}
		if pairVolume != nil {
			volume24H, _ := new(big.Float).SetString(pairVolume.Volume24H.String())
			tradingFee24H := big.NewFloat(0).Mul(volume24H, big.NewFloat(0.02))
			tradingFeeYear := big.NewFloat(0).Mul(tradingFee24H, big.NewFloat(365))

			pairLiquidity, err := u.Repo.FindSwapPairCurrentReserve(ctx, query)
			if err != nil && err != mongo.ErrNoDocuments {
				logger.AtLog.Logger.Error("SwapGetPairApr", zap.Error(err))
				return nil, err
			}

			poolLiquidity := big.NewFloat(0)
			if pairLiquidity != nil {
				if strings.EqualFold(pairObj.Token0, wbtcContractAddr) {
					poolLiquidity, _ = new(big.Float).SetString(pairLiquidity.Reserve0.String())
				} else if strings.EqualFold(pairObj.Token1, wbtcContractAddr) {
					poolLiquidity, _ = new(big.Float).SetString(pairLiquidity.Reserve1.String())
				}
				poolLiquidity = big.NewFloat(0).Mul(poolLiquidity, big.NewFloat(2))
			}

			if poolLiquidity.Cmp(big.NewFloat(0)) != 0 {
				fmt.Println(poolLiquidity.String())
				poolApr := big.NewFloat(0).Quo(tradingFeeYear, poolLiquidity)
				aprPercent, _ = (big.NewFloat(0).Mul(poolApr, big.NewFloat(100))).Float64()
			}
		}
	}

	logger.AtLog.Logger.Info("SwapGetPairApr", zap.Any("data", aprPercent))
	return aprPercent, nil
}

func (u *Usecase) GetRoutePair(ctx context.Context, fromToken, toToken string) (interface{}, error) {
	var err error

	listPairs := []*entity.SwapPair{}
	pair, err := u.Repo.FindSwapPairByTokens(ctx, fromToken, toToken)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("GetRoutePair", zap.Error(err))
		return nil, err
	}
	if pair != nil {
		listPairs = append(listPairs, pair)
	} else {
		wbtcContractAddr := u.Repo.ParseConfigByString(ctx, "wbtc_contract_address")
		pair1, err := u.Repo.FindSwapPairByTokens(ctx, fromToken, wbtcContractAddr)
		if err != nil {
			err := errors.New("Pair is not exist")
			logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
			return nil, err
		}
		if pair1 != nil {
			listPairs = append(listPairs, pair1)
		}

		pair2, err := u.Repo.FindSwapPairByTokens(ctx, wbtcContractAddr, toToken)
		if err != nil {
			err := errors.New("Pair is not exist")
			logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
			return nil, err
		}
		if pair2 != nil {
			listPairs = append(listPairs, pair2)
		}
	}

	if len(listPairs) == 0 {
		err := errors.New("Pair is not exist")
		logger.AtLog.Logger.Error("SwapAddOrUpdateIdo", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("GetRoutePair", zap.Any("data", listPairs))
	return listPairs, nil
}

func (u *Usecase) UpdateDataSwapPair(ctx context.Context) error {
	pairQuery := entity.SwapPairFilter{}
	pairQuery.Limit = 10000
	pairQuery.Page = 1

	pairs, err := u.Repo.FindSwapPairs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
		return err
	}

	for _, pair := range pairs {
		token0, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: pair.Token0})
		if token0 != nil {
			pair.Token0Obj = *token0
		}

		token1, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: pair.Token1})
		if token1 != nil {
			pair.Token1Obj = *token1
		}

		err := u.Repo.UpdateSwapPair(ctx, pair)
		if err != nil {
			logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) UpdateDataSwapToken(ctx context.Context) error {
	pairQuery := entity.TokenFilter{}
	pairQuery.Limit = 10000
	pairQuery.Page = 1

	tokens, err := u.Repo.FindListTokens(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
		return err
	}

	for _, token := range tokens {
		if token.Network == "" {
			token.Network = "TC"
			token.Priority = 0
			err := u.Repo.UpdateToken(ctx, token)
			if err != nil {
				logger.AtLog.Logger.Error("UpdateDataSwapToken", zap.Error(err))
				return err
			}
		}
	}
	return nil
}
