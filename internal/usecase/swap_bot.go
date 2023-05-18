package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
	"strings"

	"go.uber.org/zap"
)

func (u *Usecase) AddSwapBotConfig(ctx context.Context, req *request.SwapBotConfigRequest) error {
	pair, _ := u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
		Pair: strings.ToLower(req.Pair),
	})
	botConfig := &entity.SwapBotConfig{}
	botConfig.SwapPair = *pair
	botConfig.Address = req.Address
	botConfig.MinValue = req.MinValue
	botConfig.MaxValue = req.MaxValue
	_, err := u.Repo.InsertOne(botConfig)
	if err != nil {
		logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
		return err
	}
	return nil
}

// func (u *Usecase) DoJobSwapBot(ctx context.Context) error {
// 	contractConfig, _ := u.TcSwapGetWrapTokenContractAddr(ctx)

// 	configs, err := u.Repo.FindSwapBotConfigs(ctx, entity.SwapBotConfigFilter{})
// 	if err != nil {
// 		logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
// 		return err
// 	}

// 	strDate := time.Now().Format("2006-01-02")
// 	for _, item := range configs {
// 		reserve0, reserve1, err := u.BlockChainApi.TcSwapGetReserves(item.Pair)
// 		if err != nil {
// 			logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
// 			return err
// 		}
// 		token, baseToken, baseIndex, err := u.TcSwapGetBaseTokenOnPair(ctx, &item.SwapPair)
// 		if err != nil {
// 			logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
// 			return err
// 		}

// 		tmpReserve0 := helpers.ConvertWeiToBigFloat(reserve0, 18)
// 		tmpReserve1 := helpers.ConvertWeiToBigFloat(reserve1, 18)

// 		currentPrice := big.NewFloat(0).Quo(tmpReserve0, tmpReserve1)
// 		if baseIndex == 1 {
// 			currentPrice = big.NewFloat(0).Quo(tmpReserve1, tmpReserve0)
// 		}

// 		if !strings.EqualFold(item.CurrentDate, strDate) {
// 			item.CurrentDate = strDate
// 			item.BeginReserve0, _ = primitive.ParseDecimal128(tmpReserve0.String())
// 			item.BeginReserve1, _ = primitive.ParseDecimal128(tmpReserve1.String())

// 			if token != nil && baseToken != nil {
// 				expectPercent := item.MinValue + rand.Float64()*(item.MaxValue-item.MinValue)
// 				item.ExpectValue = expectPercent
// 				item.BeginPrice, _ = primitive.ParseDecimal128(currentPrice.String())
// 				if baseIndex == 1 {
// 					expectReseve0 := big.NewFloat(0).Sub(tmpReserve0, big.NewFloat(0).Mul(tmpReserve0, big.NewFloat(expectPercent/2/100)))
// 					item.ExpectReserve0, _ = primitive.ParseDecimal128(expectReseve0.String())

// 					expectReseve1 := big.NewFloat(0).Add(tmpReserve1, big.NewFloat(0).Mul(tmpReserve1, big.NewFloat(expectPercent/2/100)))
// 					item.ExpectReserve1, _ = primitive.ParseDecimal128(expectReseve1.String())
// 				} else {
// 					expectReseve0 := big.NewFloat(0).Add(tmpReserve0, big.NewFloat(0).Mul(tmpReserve0, big.NewFloat(expectPercent/2/100)))
// 					item.ExpectReserve0, _ = primitive.ParseDecimal128(expectReseve0.String())

// 					expectReseve1 := big.NewFloat(0).Sub(tmpReserve1, big.NewFloat(0).Mul(tmpReserve1, big.NewFloat(expectPercent/2/100)))
// 					item.ExpectReserve1, _ = primitive.ParseDecimal128(expectReseve1.String())
// 				}
// 			}

// 			err = u.Repo.UpdateSwapBotConfig(ctx, item)
// 			if err != nil {
// 				logger.AtLog.Logger.Error("Update mongo entity failed", zap.Error(err))
// 				return err
// 			}
// 		}
// 		tx, err := u.Repo.FindSwapBotTransaction(ctx, entity.SwapBotTransactionFilter{
// 			Status: 0,
// 		})

// 		if err != nil && err != mongo.ErrNoDocuments {
// 			logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
// 			return err
// 		}

// 		if tx != nil {
// 			return nil
// 		}

// 		beginPrice, _ := new(big.Float).SetString(item.BeginPrice.String())
// 		fmt.Println(fmt.Sprintf(`beginPrice=%s`, beginPrice.String()))

// 		fmt.Println(fmt.Sprintf(`currentPrice=%s`, currentPrice.String()))

// 		gapPrice := big.NewFloat(0).Sub(currentPrice, beginPrice)
// 		fmt.Println(fmt.Sprintf(`gapPrice=%s`, gapPrice.String()))

// 		gapPricePercent := big.NewFloat(0).Mul(big.NewFloat(0).Quo(gapPrice, beginPrice), big.NewFloat(100))
// 		fmt.Println(fmt.Sprintf(`gapPricePercent=%s`, gapPricePercent.String()))

// 		if gapPricePercent.Cmp(big.NewFloat(item.MinValue-1)) < 0 ||
// 			gapPricePercent.Cmp(big.NewFloat(item.MaxValue+1)) > 0 {
// 			// wallet, err := u.SwapGetWalletAddress(ctx, strings.ToLower(item.Address))
// 			// if err != nil && err != mongo.ErrNoDocuments {
// 			// 	logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
// 			// 	return err
// 			// }

// 			expectReseve0, _ := big.NewFloat(0).SetString(item.ExpectReserve0.String())
// 			expectReseve1, _ := big.NewFloat(0).SetString(item.ExpectReserve1.String())
// 			// beginReseve0, _ := big.NewFloat(0).SetString(item.BeginReserve0.String())
// 			// beginReseve1, _ := big.NewFloat(0).SetString(item.BeginReserve1.String())

// 			isBuy := true
// 			buySellValue := big.NewFloat(0).Sub(expectReseve0, tmpReserve0)
// 			amountOutMin := big.NewFloat(0).Sub(tmpReserve1, expectReseve1)
// 			if expectReseve0.Cmp(tmpReserve0) < 0 {
// 				isBuy = false
// 				buySellValue = big.NewFloat(0).Sub(tmpReserve0, expectReseve0)
// 				amountOutMin = big.NewFloat(0).Sub(expectReseve1, tmpReserve1)
// 			}

// 			buySellValue = big.NewFloat(0).Add(buySellValue, big.NewFloat(0).Mul(buySellValue, big.NewFloat(0.02)))
// 			amountOutMin = big.NewFloat(0).Sub(amountOutMin, big.NewFloat(0).Mul(amountOutMin, big.NewFloat(0.1)))
// 			fmt.Println(fmt.Sprintf(`buySellValue=%s`, buySellValue.String()))
// 			fmt.Println(fmt.Sprintf(`amountOutMin=%s`, amountOutMin.String()))

// 			// if wallet != nil {
// 			// 	txHash := ""
// 			// 	if isBuy {
// 			// 		coinBlanceInt, _ := u.BlockChainApi.Erc20GetCoinBalance(item.Token0, item.Address)
// 			// 		coinBlance := helpers.ConvertWeiToBigFloat(coinBlanceInt, 18)
// 			// 		fmt.Println(fmt.Sprintf(`coinBlance = %s`, coinBlance.String()))

// 			// 		if coinBlance.Cmp(buySellValue) > 0 {
// 			// 			// estAmountIn, err := u.BlockChainApi.TcSwapGetAmountOut(
// 			// 			// 	contractConfig.RouterContractAddr,
// 			// 			// 	helpers.EtherToWei(buySellValue),
// 			// 			// )

// 			// 			txHash, err = u.BlockChainApi.TcSwapExactTokensForTokens(
// 			// 				contractConfig.RouterContractAddr,
// 			// 				helpers.EtherToWei(buySellValue),
// 			// 				helpers.EtherToWei(amountOutMin),
// 			// 				item.Address,
// 			// 				wallet.Prk,
// 			// 				item.Token0,
// 			// 				item.Token1,
// 			// 			)
// 			// 			if err != nil {
// 			// 				logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
// 			// 				return err
// 			// 			}
// 			// 		}
// 			// 	} else {
// 			// 		coinBlanceInt, _ := u.BlockChainApi.Erc20GetCoinBalance(item.Token1, item.Address)
// 			// 		coinBlance := helpers.ConvertWeiToBigFloat(coinBlanceInt, 18)
// 			// 		fmt.Println(fmt.Sprintf(`coinBlance = %s`, coinBlance))

// 			// 		if coinBlance.Cmp(buySellValue) > 0 {
// 			// 			txHash, err = u.BlockChainApi.TcSwapExactTokensForTokens(
// 			// 				contractConfig.RouterContractAddr,
// 			// 				helpers.EtherToWei(buySellValue),
// 			// 				helpers.EtherToWei(amountOutMin),
// 			// 				item.Address,
// 			// 				wallet.Prk,
// 			// 				item.Token1,
// 			// 				item.Token0,
// 			// 			)
// 			// 			if err != nil {
// 			// 				logger.AtLog.Logger.Error("DoJobSwapBot", zap.Error(err))
// 			// 				return err
// 			// 			}
// 			// 		}
// 			// 	}

// 			// 	if txHash != "" {

// 			// 	}
// 			// }
// 		}
// 	}
// 	return nil
// }
