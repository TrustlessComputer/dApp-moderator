package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	mapBlackListContract := make(map[string]string)
	listBlackList, _ := u.Repo.FindBlackListTokens(ctx, entity.SwapBlackListokenFilter{BaseFilters: entity.BaseFilters{Limit: 10000, Page: 1}})
	for _, item := range listBlackList {
		mapBlackListContract[item.Address] = "1"
	}

	isWbtcInArray := false
	wbtcContractAddr := u.Repo.ParseConfigByString(ctx, "wbtc_contract_address")
	for _, pair := range pairs {
		if fromToken == "" || (fromToken != "" && strings.EqualFold(fromToken, pair.Token1)) {
			if _, ok := mapBlackListContract[pair.Token0]; !ok {
				contracts = append(contracts, pair.Token0)
				if strings.EqualFold(wbtcContractAddr, pair.Token0) {
					isWbtcInArray = true
				}
			}
		}

		if fromToken == "" || (fromToken != "" && strings.EqualFold(fromToken, pair.Token0)) {
			if _, ok := mapBlackListContract[pair.Token1]; !ok {
				contracts = append(contracts, pair.Token1)
				if strings.EqualFold(wbtcContractAddr, pair.Token1) {
					isWbtcInArray = true
				}
			}
		}
	}

	if isWbtcInArray && fromToken != "" {
		for _, pair := range pairs {
			if strings.EqualFold(wbtcContractAddr, pair.Token0) &&
				!strings.EqualFold(fromToken, pair.Token1) {
				if _, ok := mapBlackListContract[pair.Token1]; !ok {
					contracts = append(contracts, pair.Token1)
				}
			} else if strings.EqualFold(wbtcContractAddr, pair.Token1) &&
				!strings.EqualFold(fromToken, pair.Token0) {
				if _, ok := mapBlackListContract[pair.Token0]; !ok {
					contracts = append(contracts, pair.Token0)
				}
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
	redisKey := "tc-swap:wrap-token-config"
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

func (u *Usecase) GetWrapTokenPriceBySymbol(ctx context.Context) (float64, float64) {
	btcPrice := u.Repo.ParseConfigByFloat64(ctx, "swap_btc_price")
	ethPrice := u.Repo.ParseConfigByFloat64(ctx, "swap_eth_price")
	return btcPrice, ethPrice
}

func (u *Usecase) FindTokensReport(ctx context.Context, filter request.PaginationReq, address, search, sortBy string, sortType int) (interface{}, error) {
	query := entity.TokenReportFilter{}
	query.FromPagination(filter)
	query.Address = address
	query.SortBy = sortBy
	query.SortType = sortType
	query.Search = search

	redisKey := fmt.Sprintf("tc-swap:token-reports-%d-%d-%s-%s-%s-%d", query.Page, query.Limit, address, search, sortBy, sortType)
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
		wtokenConfig, _ := u.TcSwapGetWrapTokenContractAddr(ctx)

		reports, err := u.Repo.FindTokenReport(ctx, query)
		if err != nil {
			logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
			return nil, err
		}
		btcPrice, ethPrice := u.GetWrapTokenPriceBySymbol(ctx)

		for _, item := range reports {
			if item.BaseTokenSymbol == "" {
				item.BaseTokenSymbol = string(entity.SwapBaseTokenSymbolWBTC)
			}

			tmUsdPrice := float64(0)
			if item.BaseTokenSymbol == string(entity.SwapBaseTokenSymbolWETH) {
				tmUsdPrice = ethPrice
			} else if item.BaseTokenSymbol == string(entity.SwapBaseTokenSymbolWBTC) {
				tmUsdPrice = btcPrice
			}

			if s, err := strconv.ParseFloat(item.Price.String(), 64); err == nil {
				item.BtcPrice = s
				item.UsdPrice = s * tmUsdPrice
			}

			if s, err := strconv.ParseFloat(item.Volume.String(), 64); err == nil {
				item.BtcVolume = s
				item.UsdVolume = s * tmUsdPrice
			}

			if s, err := strconv.ParseFloat(item.TotalVolume.String(), 64); err == nil {
				item.BtcTotalVolume = s
				item.UsdTotalVolume = s * tmUsdPrice
			}

			if s, err := strconv.ParseFloat(item.MarketCap.String(), 64); err == nil {
				item.UsdMarketCap = s * tmUsdPrice
			}

			if item.Address == wtokenConfig.WbtcContractAddr {
				item.UsdPrice = btcPrice
			}
			// else if item.Address == wtokenConfig.WethContractAddr {
			// 	item.UsdPrice = ethPrice
			// }
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
	pairQuery.Limit = 2000
	pairQuery.Page = 1
	pairQuery.Symbol = "WBTC"

	pairSyncs, err := u.Repo.FindSwapPairSyncs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
		return err
	}

	mapPair := map[string]*entity.SwapPair{}
	for _, pairSync := range pairSyncs {
		var pair *entity.SwapPair

		if p, ok := mapPair[strings.ToLower(pairSync.ContractAddress)]; ok {
			pair = p
		} else {
			pair, _ = u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
				Pair: strings.ToLower(pairSync.ContractAddress),
			})
			mapPair[strings.ToLower(pairSync.ContractAddress)] = pair
		}

		if pair != nil {
			token, baseToken, baseIndex, err := u.TcSwapGetBaseTokenOnPair(ctx, pair)
			if err != nil {
				logger.AtLog.Logger.Error("TcSwapPairCreateSwapEvent", zap.Error(err))
				return err
			}
			if token != nil && baseToken != nil {
				pairSync.Token = token.Address
				tmpReserce0, _ := new(big.Float).SetString(pairSync.Reserve0.String())
				tmpReserce1, _ := new(big.Float).SetString(pairSync.Reserve1.String())
				tmpPrice := big.NewFloat(0).Quo(tmpReserce0, tmpReserce1)
				if baseIndex == 1 {
					tmpPrice = big.NewFloat(0).Quo(tmpReserce1, tmpReserce0)
				}
				pairSync.Price, _ = primitive.ParseDecimal128(tmpPrice.String())
				pairSync.BaseTokenSymbol = baseToken.Symbol
			}
			pairSync.Pair = pair

			err = u.Repo.UpdateSwapPairSync(ctx, pairSync)
			if err != nil {
				fmt.Printf(pairSync.Id())
				logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
				// return err
			}
		}
		// }
	}
	return nil
}

func (u *Usecase) UpdateDataSwapHistory(ctx context.Context) error {
	pairQuery := entity.SwapPairSwapHistoriesFilter{}
	pairQuery.Limit = 2000
	pairQuery.Page = 1
	pairQuery.Symbol = "WBTC"

	pairSyncs, err := u.Repo.FindSwapPairHistories(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
		return err
	}

	mapPair := map[string]*entity.SwapPair{}
	for _, pairSync := range pairSyncs {
		var pair *entity.SwapPair

		if p, ok := mapPair[strings.ToLower(pairSync.ContractAddress)]; ok {
			pair = p
		} else {
			pair, _ = u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
				Pair: strings.ToLower(pairSync.ContractAddress),
			})
			mapPair[strings.ToLower(pairSync.ContractAddress)] = pair
		}

		if pair != nil {
			token, baseToken, baseIndex, err := u.TcSwapGetBaseTokenOnPair(ctx, pair)
			if err != nil {
				logger.AtLog.Logger.Error("TcSwapPairCreateSwapEvent", zap.Error(err))
				return err
			}
			if token != nil && baseToken != nil {
				pairSync.Token = token.Address
				tmpAmount0In, _ := new(big.Float).SetString(pairSync.Amount0In.String())
				tmpAmount0Out, _ := new(big.Float).SetString(pairSync.Amount0Out.String())
				tmpAmount1In, _ := new(big.Float).SetString(pairSync.Amount1In.String())
				tmpAmount1Out, _ := new(big.Float).SetString(pairSync.Amount1Out.String())

				tmpAmount0 := big.NewFloat(0).Add(tmpAmount0In, tmpAmount0Out)
				tmpAmount1 := big.NewFloat(0).Add(tmpAmount1In, tmpAmount1Out)

				tmpVolume := tmpAmount0
				tmpPrice := big.NewFloat(0).Quo(tmpAmount0, tmpAmount1)
				if baseIndex == 1 {
					tmpVolume = tmpAmount1
					tmpPrice = big.NewFloat(0).Quo(tmpAmount1, tmpAmount0)
				}

				pairSync.Volume, _ = primitive.ParseDecimal128(tmpVolume.String())
				pairSync.Price, _ = primitive.ParseDecimal128(tmpPrice.String())
				pairSync.BaseTokenSymbol = baseToken.Symbol
			}
			pairSync.Pair = pair

			err = u.Repo.UpdateSwapPairHistory(ctx, pairSync)
			if err != nil {
				fmt.Println(pairSync.Id())
				// logger.AtLog.Logger.Error("FindTokensInPool", zap.Error(err))
				// return err
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
		config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
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
				if strings.EqualFold(pairObj.Token0, config.WbtcContractAddr) {
					poolLiquidity, _ = new(big.Float).SetString(pairLiquidity.Reserve0.String())
				} else if strings.EqualFold(pairObj.Token1, config.WbtcContractAddr) {
					poolLiquidity, _ = new(big.Float).SetString(pairLiquidity.Reserve1.String())
				} else if strings.EqualFold(pairObj.Token0, config.WethContractAddr) {
					poolLiquidity, _ = new(big.Float).SetString(pairLiquidity.Reserve0.String())
				} else if strings.EqualFold(pairObj.Token1, config.WethContractAddr) {
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

func (u *Usecase) SwapGetPairAprListReport(ctx context.Context, filter request.PaginationReq, search string) (interface{}, error) {
	query := entity.TokenReportFilter{}
	query.FromPagination(filter)
	query.Search = search

	redisKey := fmt.Sprintf("tc-swap:pair-apr-reports-%d-%d-%s", query.Page, query.Limit, search)
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
		reports := []entity.SwapPairAprReport{}
		err = json.Unmarshal(b, &reports)
		if err != nil {
			return nil, err
		}
		return reports, nil
	} else {
		btcPrice, ethPrice := u.GetWrapTokenPriceBySymbol(ctx)
		reports, err := u.Repo.FindPairAprReport(ctx, query)
		if err != nil {
			logger.AtLog.Logger.Error("SwapGetPairAprListReport", zap.Error(err))
			return nil, err
		}
		for _, item := range reports {
			tmUsdPrice := float64(0)
			if item.BaseTokenSymbol == string(entity.SwapBaseTokenSymbolWETH) {
				tmUsdPrice = ethPrice
			} else if item.BaseTokenSymbol == string(entity.SwapBaseTokenSymbolWBTC) {
				tmUsdPrice = btcPrice
			}
			if s, err := strconv.ParseFloat(item.Volume.String(), 64); err == nil {
				item.UsdVolume = s * tmUsdPrice
			}

			if s, err := strconv.ParseFloat(item.TotalVolume.String(), 64); err == nil {
				item.UsdTotalVolume = s * tmUsdPrice
			}
		}

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
		reserve0, reserve1, err := u.BlockChainApi.TcSwapGetReserves(pair.Pair)
		if err != nil {
			logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
			return err
		}

		tmpReserve0 := helpers.ConvertWeiToBigFloat(reserve0, 18)
		tmpReserve1 := helpers.ConvertWeiToBigFloat(reserve1, 18)

		if pair.Token0Obj == nil {
			token0, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: pair.Token0})
			if token0 != nil {
				pair.Token0Obj = token0
			}
		}

		if pair.Token1Obj == nil {
			token1, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: pair.Token1})
			if token1 != nil {
				pair.Token1Obj = token1
			}
		}
		pair.Reserve0, _ = primitive.ParseDecimal128(tmpReserve0.String())
		pair.Reserve1, _ = primitive.ParseDecimal128(tmpReserve1.String())
		pair.SetUpdatedAt()

		err = u.Repo.UpdateSwapPair(ctx, pair)
		if err != nil {
			logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) UpdateBaseSymbolToken(ctx context.Context) error {
	config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
	pairQuery := entity.SwapPairFilter{}
	pairQuery.Limit = 10000
	pairQuery.Page = 1
	pairQuery.FromToken = config.WbtcContractAddr

	//base WBTC
	pairs, err := u.Repo.FindSwapPairs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
		return err
	}

	for _, pair := range pairs {
		tmpTokenAddr := pair.Token0
		baseToken := pair.Token1Obj
		if strings.EqualFold(pair.Token0, config.WbtcContractAddr) {
			tmpTokenAddr = pair.Token1
			baseToken = pair.Token0Obj
		}
		token, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: tmpTokenAddr})
		if token != nil && token.BaseTokenSymbol == "" {
			token.BaseTokenSymbol = baseToken.Symbol
			err = u.Repo.UpdateBaseSymbolToken(ctx, token)
			if err != nil {
				logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
				return err
			}
		}
	}

	pairQuery.FromToken = config.WethContractAddr
	ethPairs, err := u.Repo.FindSwapPairs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
		return err
	}

	for _, pair := range ethPairs {
		tmpTokenAddr := pair.Token0
		baseToken := pair.Token1Obj
		if strings.EqualFold(pair.Token0, config.WbtcContractAddr) {
			tmpTokenAddr = pair.Token1
			baseToken = pair.Token0Obj
		}
		token, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: tmpTokenAddr})
		if token != nil && token.BaseTokenSymbol == "" {
			token.BaseTokenSymbol = baseToken.Symbol
			token.SetUpdatedAt()
			err = u.Repo.UpdateBaseSymbolToken(ctx, token)
			if err != nil {
				logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
				return err
			}
		}
	}
	return nil
}

// func (u *Usecase) UpdateDataSwapToken(ctx context.Context) error {
// 	pairQuery := entity.TokenFilter{}
// 	pairQuery.Limit = 10000
// 	pairQuery.Page = 1

// 	tokens, err := u.Repo.FindListTokens(ctx, pairQuery)
// 	if err != nil {
// 		logger.AtLog.Logger.Error("UpdateDataSwapPair", zap.Error(err))
// 		return err
// 	}

// 	for _, token := range tokens {
// 		if token.Network == "" {
// 			token.Network = "Bitcoin - TC"
// 			token.Priority = 0
// 			err := u.Repo.UpdateToken(ctx, token)
// 			if err != nil {
// 				logger.AtLog.Logger.Error("UpdateDataSwapToken", zap.Error(err))
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }
