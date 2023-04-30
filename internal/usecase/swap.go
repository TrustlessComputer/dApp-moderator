package usecase

import (
	"context"
	"dapp-moderator/external/blockchain_api"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *Usecase) TcSwapScanEvents(ctx context.Context) error {
	startBlocks, err := u.Repo.ParseConfigByInt(ctx, "swap_scan_current_block_number")
	if err != nil {
		return err
	}

	contracts := []string{}
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "swap_factory_contract_address"))
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "swap_router_contract_address"))

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
		u.TcSwapCreateOrUpdateCurrentScanBlock(ctx, eventResp.LastBlockNumber)
	}
	return nil
}

func (u *Usecase) TcSwapCreateOrUpdateCurrentScanBlock(ctx context.Context, endBlock int64) error {
	configName := "swap_scan_current_block_number"
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

func (u *Usecase) TcSwapScanEventsByTransactionHash(ctx context.Context, txHash string) error {
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
		swapPair := &entity.SwapPairEvents{}
		swapPair.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPair.TxHash = strings.ToLower(eventResp.TxHash)
		swapPair.Amount0, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount0, 18).String())
		swapPair.Amount1, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Amount1, 18).String())
		swapPair.Sender = eventResp.Sender
		swapPair.To = eventResp.To
		swapPair.EventType = eventType
		swapPair.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)
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
		swapPair := &entity.SwapPairSync{}
		swapPair.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPair.TxHash = strings.ToLower(eventResp.TxHash)
		swapPair.Reserve0, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Reserve0, 18).String())
		swapPair.Reserve1, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Reserve1, 18).String())
		swapPair.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)
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

		_, err = u.Repo.InsertOne(swapPair)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}
