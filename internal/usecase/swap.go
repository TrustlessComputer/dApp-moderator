package usecase

import (
	"context"
	"dapp-moderator/external/blockchain_api"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *Usecase) TcSwapScanEvents(ctx context.Context) error {
	configName := "swap_scan_current_block_number"
	startBlocks, err := u.Repo.ParseConfigByInt(ctx, configName)
	if err != nil {
		return err
	}

	contracts := []string{}
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "swap_factory_contract_address"))
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "swap_router_contract_address"))
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "gm_payment_contract_address"))

	eventResp, err := u.BlockChainApi.TcSwapEvents(contracts, 0, startBlocks, 0)
	if err != nil {
		return err
	}
	errs := u.TcSwapEventsByTransactionEventResp(
		ctx, eventResp,
	)
	if len(errs) > 0 {
		return errs[0]
	} else {
		u.TcSwapCreateOrUpdateCurrentScanBlock(ctx, eventResp.LastBlockNumber, configName)
	}

	u.TcSwapScanPairEvents(ctx, startBlocks)
	return nil
}

func (u *Usecase) TcSwapScanPairEvents(ctx context.Context, startBlocks int64) error {
	configName := "swap_scan_pair_current_block_number"
	currentBlocks, _ := u.Repo.ParseConfigByInt(ctx, configName)
	if currentBlocks == 0 {
		currentBlocks = startBlocks
	}
	contracts := []string{}
	pairQuery := entity.SwapPairFilter{}
	pairQuery.Limit = 10000
	pairQuery.Page = 1

	pairs, err := u.Repo.FindSwapPairs(ctx, pairQuery)
	if err != nil {
		logger.AtLog.Logger.Error("TcSwapScanPairEvents", zap.Error(err))
		return err
	}
	for _, item := range pairs {
		contracts = append(contracts, item.Pair)
	}
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "swap_factory_contract_address"))
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "swap_router_contract_address"))

	eventResp, err := u.BlockChainApi.TcSwapEvents(contracts, 0, currentBlocks, 0)
	if err != nil {
		return err
	}
	errs := u.TcSwapEventsByTransactionEventResp(
		ctx, eventResp,
	)
	if len(errs) > 0 {
		return errs[0]
	} else {
		u.TcSwapCreateOrUpdateCurrentScanBlock(ctx, eventResp.LastBlockNumber, configName)
	}
	return nil
}

func (u *Usecase) TcSwapCreateOrUpdateCurrentScanBlock(ctx context.Context, endBlock int64, configName string) error {
	dbSwapConfig, err := u.Repo.FindSwapConfig(ctx, entity.SwapConfigsFilter{
		Name: configName,
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}
	isCreated := false
	if dbSwapConfig == nil {
		dbSwapConfig = &entity.SwapConfigs{}
		isCreated = true
	}
	dbSwapConfig.Name = configName
	dbSwapConfig.Value = strconv.FormatInt(endBlock, 10)
	if isCreated {
		_, err = u.Repo.InsertOne(dbSwapConfig)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	} else {
		err = u.Repo.UpdateSwapConfig(ctx, dbSwapConfig)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) TcSwapScanEventsByTransactionHash(txHash string) error {
	ctx := context.Background()
	pendingTx, _ := u.Repo.FindSwapPendingTransaction(ctx, entity.SwapPendingTransactionsFilter{TxHash: []string{txHash}})
	if pendingTx == nil {
		pendingTx := &entity.SwapPendingTransactions{}
		pendingTx.Timestamp = time.Now()
		pendingTx.TxHash = txHash
		_, err := u.Repo.InsertOne(pendingTx)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}

	eventResp, err := u.BlockChainApi.TcSwapEventsByTransaction(txHash)
	if err != nil {
		return err
	}
	errs := u.TcSwapEventsByTransactionEventResp(
		ctx, eventResp,
	)
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

func (u *Usecase) TcSwapEventsByTransactionEventResp(ctx context.Context, eventResp *blockchain_api.TcSwapEventResp) []error {
	var err error
	var errs []error
	for _, event := range eventResp.PairCreated {
		err = u.TcSwapCreatedPair(ctx, event)
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, event := range eventResp.PairMint {
		err = u.TcSwapPairCreateEvent(ctx, event, entity.SwapPairEventsTypeMint)
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, event := range eventResp.PairBurn {
		err = u.TcSwapPairCreateEvent(ctx, event, entity.SwapPairEventsTypeBurn)
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, event := range eventResp.PairSync {
		err = u.TcSwapPairCreateSyncEvent(ctx, event)
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, event := range eventResp.Swap {
		err = u.TcSwapPairCreateSwapEvent(ctx, event)
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, event := range eventResp.GmPaymentPaid {
		err = u.TcGmPaymentPaidEvent(ctx, event)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (u *Usecase) TcSwapCreatedPair(ctx context.Context, eventResp *blockchain_api.TcSwapPairCreatedEventResp) error {
	// check if token exist
	dbSwapPair, err := u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
		Pair:   strings.ToLower(eventResp.Pair),
		TxHash: strings.ToLower(eventResp.TxHash),
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	if dbSwapPair != nil {
		return nil
	} else {
		swapPair := &entity.SwapPair{}
		swapPair.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPair.Pair = strings.ToLower(eventResp.Pair)
		swapPair.TxHash = strings.ToLower(eventResp.TxHash)
		swapPair.Arg3 = eventResp.Arg3.Int64()
		swapPair.Token0 = eventResp.Token0
		swapPair.Token1 = eventResp.Token1
		swapPair.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)

		token0, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: eventResp.Token0})
		if token0 != nil {
			swapPair.Token0Obj = token0
		}

		token1, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: eventResp.Token1})
		if token1 != nil {
			swapPair.Token1Obj = token1
		}
		_, err = u.Repo.InsertOne(swapPair)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) TcSwapPairCreateEvent(ctx context.Context, eventResp *blockchain_api.TcSwapMintBurnEventResp, eventType entity.SwapPairEventsType) error {
	// check if token exist
	dbSwapPair, err := u.Repo.FindSwapPairEvents(ctx, entity.SwapPairEventFilter{
		ContractAddress: strings.ToLower(eventResp.ContractAddress),
		TxHash:          strings.ToLower(eventResp.TxHash),
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	if dbSwapPair != nil {
		return nil
	} else {
		pair, _ := u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
			Pair: strings.ToLower(eventResp.ContractAddress),
		})

		swapPair := &entity.SwapPairEvents{}
		swapPair.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPair.TxHash = strings.ToLower(eventResp.TxHash)
		swapPair.Amount0, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount0, 18).String())
		swapPair.Amount1, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount1, 18).String())
		swapPair.Sender = eventResp.Sender
		swapPair.To = eventResp.To
		swapPair.EventType = eventType
		swapPair.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)
		if pair != nil {
			swapPair.Pair = pair
		}
		_, err = u.Repo.InsertOne(swapPair)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) TcSwapPairCreateSyncEvent(ctx context.Context, eventResp *blockchain_api.TcSwapSyncEventResp) error {
	// check if token exist
	dbSwapPair, err := u.Repo.FindSwapPairSync(ctx, entity.SwapPairSyncFilter{
		ContractAddress: strings.ToLower(eventResp.ContractAddress),
		TxHash:          strings.ToLower(eventResp.TxHash),
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	if dbSwapPair != nil {
		return nil
	} else {
		pair, _ := u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
			Pair: strings.ToLower(eventResp.ContractAddress),
		})

		token, baseToken, baseIndex, err := u.TcSwapGetBaseTokenOnPair(ctx, pair)
		if err != nil {
			logger.AtLog.Logger.Error("TcSwapPairCreateSwapEvent", zap.Error(err))
			return err
		}

		swapPairSync := &entity.SwapPairSync{}
		swapPairSync.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPairSync.TxHash = strings.ToLower(eventResp.TxHash)
		swapPairSync.Reserve0, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Reserve0, 18).String())
		swapPairSync.Reserve1, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Reserve1, 18).String())
		swapPairSync.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)
		if token != nil && pair != nil && baseToken != nil {
			swapPairSync.Token = token.Address
			tmpPrice := big.NewFloat(0).Quo(helpers.ConvertWeiToBigFloat(eventResp.Reserve0, 18), helpers.ConvertWeiToBigFloat(eventResp.Reserve1, 18))
			if baseIndex == 1 {
				tmpPrice = big.NewFloat(0).Quo(helpers.ConvertWeiToBigFloat(eventResp.Reserve1, 18), helpers.ConvertWeiToBigFloat(eventResp.Reserve0, 18))
			}
			swapPairSync.Price, _ = primitive.ParseDecimal128(tmpPrice.String())
			swapPairSync.BaseTokenSymbol = baseToken.Symbol
		}
		if pair != nil {
			swapPairSync.Pair = pair
		}
		_, err = u.Repo.InsertOne(swapPairSync)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) TcGmPaymentPaidEvent(ctx context.Context, eventResp *blockchain_api.TcGmPaymentPaidEventResp) error {
	// check if token exist
	dbSwapPair, err := u.Repo.FindUserGmPaid(ctx, entity.SwapUserGmPaidFilter{
		Address: strings.ToLower(eventResp.User),
		TxHash:  strings.ToLower(eventResp.TxHash),
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	if dbSwapPair != nil {
		return nil
	} else {

		swapPair := &entity.SwapUserGmPaid{}
		swapPair.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPair.TxHash = strings.ToLower(eventResp.TxHash)
		swapPair.UserAddress = strings.ToLower(eventResp.User)
		swapPair.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)
		swapPair.Amount, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.AmountGM, 18).String())

		_, err = u.Repo.InsertOne(swapPair)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) TcSwapPairCreateSwapEvent(ctx context.Context, eventResp *blockchain_api.TcSwapSwapEventResp) error {
	// check if token exist
	dbSwapPair, err := u.Repo.FindSwapPairSwapHistory(ctx, entity.SwapPairSwapHistoriesFilter{
		ContractAddress: strings.ToLower(eventResp.ContractAddress),
		TxHash:          strings.ToLower(eventResp.TxHash),
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	if dbSwapPair != nil {
		return nil
	} else {
		pair, _ := u.Repo.FindSwapPair(ctx, entity.SwapPairFilter{
			Pair: strings.ToLower(eventResp.ContractAddress),
		})

		token, baseToken, baseIndex, err := u.TcSwapGetBaseTokenOnPair(ctx, pair)
		if err != nil {
			logger.AtLog.Logger.Error("TcSwapPairCreateSwapEvent", zap.Error(err))
			return err
		}

		swapPair := &entity.SwapPairSwapHistories{}
		swapPair.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPair.TxHash = strings.ToLower(eventResp.TxHash)
		swapPair.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)
		swapPair.Amount0In, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount0In, 18).String())
		swapPair.Amount0Out, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount0Out, 18).String())
		swapPair.Amount1In, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount1In, 18).String())
		swapPair.Amount1Out, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount1Out, 18).String())
		swapPair.Sender = eventResp.Sender
		swapPair.To = eventResp.To
		swapPair.Index = eventResp.Index
		if token != nil && pair != nil && baseToken != nil {
			swapPair.Token = token.Address
			tmpAmount0 := big.NewFloat(0).Add(helpers.ConvertWeiToBigFloat(eventResp.Amount0In, 18), helpers.ConvertWeiToBigFloat(eventResp.Amount0Out, 18))
			tmpAmount1 := big.NewFloat(0).Add(helpers.ConvertWeiToBigFloat(eventResp.Amount1In, 18), helpers.ConvertWeiToBigFloat(eventResp.Amount1Out, 18))

			tmpVolume := tmpAmount0
			tmpPrice := big.NewFloat(0).Quo(tmpAmount0, tmpAmount1)
			if baseIndex == 1 {
				tmpVolume = tmpAmount1
				tmpPrice = big.NewFloat(0).Quo(tmpAmount1, tmpAmount0)
			}

			swapPair.Volume, _ = primitive.ParseDecimal128(tmpVolume.String())
			swapPair.Price, _ = primitive.ParseDecimal128(tmpPrice.String())
			swapPair.BaseTokenSymbol = baseToken.Symbol

		}
		if pair != nil {
			swapPair.Pair = pair
		}

		_, err = u.Repo.InsertOne(swapPair)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) TcSwapAddFronEndLog(ctx context.Context, logBody map[string]interface{}) error {
	swapFeLog := &entity.SwapFrontEndLog{}
	swapFeLog.Log = logBody

	_, err := u.Repo.InsertOne(swapFeLog)
	if err != nil {
		logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
		return err
	}
	return nil
}

func (u *Usecase) TcSwapUpdateWrapTokenPriceJob(ctx context.Context) error {
	err := u.TcSwapUpdateBTCPriceJob(ctx, "swap_btc_price")
	if err != nil {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	err = u.TcSwapUpdateBTCPriceJob(ctx, "swap_eth_price")
	if err != nil {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}
	return nil
}

func (u *Usecase) TcSwapUpdateBTCPriceJob(ctx context.Context, configName string) error {
	dbSwapConfig, err := u.Repo.FindSwapConfig(ctx, entity.SwapConfigsFilter{
		Name: configName,
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}
	dbSwapConfig.Name = configName
	var btcPrice float64
	if configName == "swap_btc_price" {
		btcPrice, _ = u.BlockChainApi.GetBitcoinPrice()
	} else if configName == "swap_eth_price" {
		btcPrice, _ = u.BlockChainApi.GetEthereumPrice()
	}

	dbSwapConfig.Value = fmt.Sprintf("%f", btcPrice)
	err = u.Repo.UpdateSwapConfig(ctx, dbSwapConfig)
	if err != nil {
		logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
		return err
	}
	return nil
}

func (u *Usecase) TcSwapSlackReport(ctx context.Context, channel string) error {
	resp, err := u.Repo.FindSwapSlackReport(ctx)
	if err != nil {
		logger.AtLog.Logger.Error("TcSwapSlackReport", zap.Error(err))
		return err
	}
	respLiq, err := u.Repo.FindSwapSlackLiquidityReport(ctx)
	if err != nil {
		logger.AtLog.Logger.Error("TcSwapSlackReport", zap.Error(err))
		return err
	}

	pairReserves, err := u.Repo.FindSwapPairCurrentReserveList(ctx, entity.SwapPairFilter{})
	if err != nil {
		logger.AtLog.Logger.Error("FindSwapPairs", zap.Error(err))
		return err
	}
	wbtcContractAddr := u.Repo.ParseConfigByString(ctx, "wbtc_contract_address")

	poolBTCLiquidity := big.NewFloat(0)
	for _, item := range pairReserves {
		tmpPoolBTCLiquidity := big.NewFloat(0)
		if strings.EqualFold(item.Token0, wbtcContractAddr) {
			tmpPoolBTCLiquidity, _ = new(big.Float).SetString(item.Reserve0.String())
		} else if strings.EqualFold(item.Token1, wbtcContractAddr) {
			tmpPoolBTCLiquidity, _ = new(big.Float).SetString(item.Reserve1.String())
		}
		poolBTCLiquidity = big.NewFloat(0).Add(poolBTCLiquidity, tmpPoolBTCLiquidity)
	}

	if resp != nil && respLiq != nil {
		btcPrice := u.Repo.ParseConfigByFloat64(ctx, "swap_btc_price")

		totalVolumeBtc := float64(0)
		volume24hBtc := float64(0)
		totalVolumeUsd := float64(0)
		volume24hUsd := float64(0)
		if s, err := strconv.ParseFloat(resp.VolumeTotal.String(), 64); err == nil {
			totalVolumeUsd = s * btcPrice
			totalVolumeBtc = s
		}

		if s, err := strconv.ParseFloat(resp.Volume24h.String(), 64); err == nil {
			volume24hUsd = s * btcPrice
			volume24hBtc = s
		}

		slackString := "*TC SWAP Report*\n"
		slackString += fmt.Sprintf("*Total Volume:* %.2f BTC | $%.2f\n", totalVolumeBtc, totalVolumeUsd)
		slackString += fmt.Sprintf("*Total Txs:* %d\n", resp.TxTotal)
		slackString += fmt.Sprintf("*Total Users:* %d\n", resp.UsersTotal)
		slackString += fmt.Sprintf("*Last 24h Volume:* %.2f BTC | $%.2f\n", volume24hBtc, volume24hUsd)
		slackString += fmt.Sprintf("*Last 24h Txs:* %d\n", resp.Tx24h)
		slackString += fmt.Sprintf("*Last 24h Users:* %d\n", resp.Users24h)

		slackString += "\n*TC Liquidity Report*\n"
		slackString += fmt.Sprintf("*Total BTC In Pool:* %.2f BTC\n", poolBTCLiquidity)
		slackString += fmt.Sprintf("*Total Pair:* %d\n", respLiq.PairTotal)
		slackString += fmt.Sprintf("*Total Txs:* %d\n", respLiq.TxTotal)
		slackString += fmt.Sprintf("*Last 24h Pair:* %d\n", respLiq.Pair24h)
		slackString += fmt.Sprintf("*Last 24h Txs:* %d\n", respLiq.Tx24h)

		listName := []string{
			"wbtc_contract_address",
			"weth_contract_address",
			"wpepe_contract_address",
			"wusdc_contract_address",
			"wordi_contract_address",
		}
		dbSwapConfigs, err := u.Repo.FindSwapConfigByListName(ctx, listName)
		if err != nil && err != mongo.ErrNoDocuments {
			logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
			return err
		}

		slackString += "\n*TC Bridge Locked Report*\n"
		for _, item := range dbSwapConfigs {
			tmpAmount, _ := new(big.Float).SetString(item.TotalSupply.String())
			tmpAmountFloat, _ := tmpAmount.Float64()
			slackString += fmt.Sprintf("*Total %s:* %.2f\n", item.Symbol, tmpAmountFloat)
		}

		helpers.SlackHook(channel, slackString)
	}

	return nil
}

func (u *Usecase) TcSwapUpdateTotalSupplyJob(ctx context.Context) error {
	listName := []string{
		"wbtc_contract_address",
		"weth_contract_address",
		"wpepe_contract_address",
		"wusdc_contract_address",
		"wordi_contract_address",
	}
	dbSwapConfigs, err := u.Repo.FindSwapConfigByListName(ctx, listName)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	// listContractAddress := []string{}
	// for _, item := range dbSwapConfigs {
	// 	listContractAddress = append(listContractAddress, item.Value)
	// }

	// listTokens, _ := u.Repo.FindTokensByContracts(ctx, listContractAddress)
	// for _, item := range listTokens {
	// 	totalSupply, _ := u.BlockChainApi.Erc20TotalSupply(item.Address)
	// 	if totalSupply != nil {
	// 		// item.TotalSupply = helpers.ConvertWeiToBigFloat(totalSupply, 18).String()
	// 		item.TotalSupply = totalSupply.String()
	// 	}
	// 	err = u.Repo.UpdateToken(ctx, item)
	// 	if err != nil {
	// 		logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
	// 		return err
	// 	}
	// }

	for _, item := range dbSwapConfigs {
		totalSupply, _ := u.BlockChainApi.Erc20TotalSupply(item.Value)
		if totalSupply != nil {
			item.TotalSupply, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(totalSupply, 18).String())
		}
		err = u.Repo.UpdateSwapConfig(ctx, item)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *Usecase) PendingTransactionHistories(ctx context.Context, filter request.PaginationReq, txs string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapPendingTransactionsFilter{}
	query.FromPagination(filter)
	if txs != "" {
		query.TxHash = strings.Split(txs, ",")
	}
	data, err = u.Repo.FindSwapPendingTransactionList(ctx, query)
	if err != nil {
		logger.AtLog.Logger.Error("PendingTransactionHistories", zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("PendingTransactionHistories", zap.Any("data", data))
	return data, nil
}

// func (u *Usecase) SwapAddOrUpdateWalletAddress(ctx context.Context, walletReq *request.SwapWalletAddressRequest) (interface{}, error) {
// 	var err error
// 	wallet, err := u.Repo.FindSwapWalletByAddress(ctx, walletReq.WalletAddress)
// 	if err != nil && err != mongo.ErrNoDocuments {
// 		logger.AtLog.Logger.Error("SwapAddOrUpdateWalletAddress", zap.Error(err))
// 		return false, err
// 	}

// 	isCreated := false
// 	if wallet == nil {
// 		isCreated = true
// 		wallet = &entity.SwapWalletAddress{}
// 	}
// 	wallet.Address = strings.ToLower(walletReq.WalletAddress)

// 	ciphertext, err := helpers.GetAESEncrypted(u.Config.Swap.SecretKey, u.Config.Swap.IvKey, walletReq.WalletAddressPrivateKey)
// 	if err != nil {
// 		logger.AtLog.Logger.Error("SwapAddOrUpdateWalletAddress", zap.Error(err))
// 		return nil, err
// 	}

// 	wallet.Prk = ciphertext
// 	if isCreated {
// 		_, err = u.Repo.InsertOne(wallet)
// 		if err != nil {
// 			logger.AtLog.Logger.Error("SwapAddOrUpdateWalletAddress", zap.Error(err))
// 			return nil, err
// 		}
// 	} else {
// 		err = u.Repo.UpdateWallet(ctx, wallet)
// 		if err != nil {
// 			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
// 			return nil, err
// 		}
// 	}

// 	logger.AtLog.Logger.Info("SwapAddOrUpdateWalletAddress", zap.Any("data", true))
// 	return true, nil
// }

// func (u *Usecase) SwapGetWalletAddress(ctx context.Context, walletAddress string) (*entity.SwapWalletAddress, error) {
// 	var err error
// 	wallet, err := u.Repo.FindSwapWalletByAddress(ctx, strings.ToLower(walletAddress))
// 	if err != nil {
// 		logger.AtLog.Logger.Error("SwapGetWalletAddress", zap.Error(err))
// 		return nil, err
// 	}

// 	plaintext, err := helpers.GetAESDecrypted(u.Config.Swap.SecretKey, u.Config.Swap.IvKey, wallet.Prk)
// 	if err != nil {
// 		logger.AtLog.Logger.Error("SwapGetWalletAddress", zap.Error(err))
// 		return nil, err
// 	}

// 	fmt.Printf("plaintext: %s\n", plaintext)
// 	wallet.Prk = string(plaintext)

// 	return wallet, nil
// }

func (u *Usecase) TcSwapGetWrapTokenContractAddr(ctx context.Context) (*entity.SwapWrapTOkenContractAddrConfig, error) {
	redisKey := "tc-swap:wrap-token-config"
	exists, err := u.Cache.Exists(redisKey)
	if err != nil {
		logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
		return nil, err
	}
	if *exists {
		dataInCache, err := u.Cache.GetData(redisKey)
		if err != nil {
			logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
			return nil, err
		}

		b := []byte(*dataInCache)
		config := &entity.SwapWrapTOkenContractAddrConfig{}
		err = json.Unmarshal(b, config)
		if err != nil {
			return nil, err
		}
		return config, nil
	} else {
		config := &entity.SwapWrapTOkenContractAddrConfig{}
		configs, err := u.Repo.FindSwapConfigByListName(ctx, []string{
			"wbtc_contract_address",
			"weth_contract_address",
			"wpepe_contract_address",
			"wusdc_contract_address",
			"wordi_contract_address",
			"swap_router_contract_address",
			"swap_factory_contract_address",
			"gm_payment_contract_address",
			"gm_token_contract_address",
			"gm_payment_admin_address",
			"gm_payment_chain_id",
		})
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return nil, err
		}
		for _, item := range configs {
			token, _ := u.Repo.FindToken(ctx, entity.TokenFilter{Address: item.Value})
			switch item.Name {
			case "wbtc_contract_address":
				config.WbtcContractAddr = item.Value
				config.WbtcToken = token
			case "weth_contract_address":
				config.WethContractAddr = item.Value
				config.WethToken = token
			case "wpepe_contract_address":
				config.WpepeContractAddr = item.Value
				config.WpepeToken = token
			case "wusdc_contract_address":
				config.WusdcContractAddr = item.Value
				config.WusdcToken = token
			case "wordi_contract_address":
				config.WordiContractAddr = item.Value
				config.WordiToken = token
			case "swap_router_contract_address":
				config.RouterContractAddr = item.Value
			case "swap_factory_contract_address":
				config.FactoryContractAddr = item.Value
			case "gm_payment_contract_address":
				config.GmPaymentContractAddr = item.Value
			case "gm_token_contract_address":
				config.GmTokenContractAddr = item.Value
			case "gm_payment_chain_id":
				config.GmPaymentChainId = item.Value
			case "gm_payment_admin_address":
				config.GmPaymentAdminAddr = item.Value
			}
		}

		reportsStr, err := json.Marshal(&config)
		if err != nil {
			logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
			return config, nil
		}
		err = u.Cache.SetStringDataWithExpTime(redisKey, string(reportsStr), 30*60)
		if err != nil {
			logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
			return config, nil
		}

		return config, nil
	}
}

func (u *Usecase) TcSwapGetBaseTokenOnPair(ctx context.Context, pair *entity.SwapPair) (*entity.Token, *entity.Token, int, error) {
	config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
	var token *entity.Token
	var baseToken *entity.Token
	baseIndex := int(0)
	if pair != nil {
		tokenAddress := ""
		if strings.EqualFold(pair.Token0, config.WbtcContractAddr) {
			tokenAddress = pair.Token1
			baseToken = config.WbtcToken
		} else if strings.EqualFold(pair.Token1, config.WbtcContractAddr) {
			baseIndex = 1
			tokenAddress = pair.Token0
			baseToken = config.WbtcToken
		} else if strings.EqualFold(pair.Token0, config.WethContractAddr) {
			tokenAddress = pair.Token1
			baseToken = config.WethToken
		} else if strings.EqualFold(pair.Token1, config.WethContractAddr) {
			baseIndex = 1
			tokenAddress = pair.Token0
			baseToken = config.WethToken
		}
		// else if strings.EqualFold(pair.Token0, config.WusdcContractAddr) {
		// 	tokenAddress = pair.Token1
		// 	baseToken = config.WusdcToken
		// } else if strings.EqualFold(pair.Token1, config.WusdcContractAddr) {
		// 	baseIndex = 1
		// 	tokenAddress = pair.Token0
		// 	baseToken = config.WusdcToken
		// }

		if tokenAddress != "" {
			token, _ = u.Repo.FindToken(ctx, entity.TokenFilter{
				Address: tokenAddress,
			})
		}
	}
	return token, baseToken, baseIndex, nil
}
