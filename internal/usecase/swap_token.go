package usecase

import (
	"context"
	"dapp-moderator/external/blockchain_api"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *Usecase) TcSwaTmTokenpScanEvents(ctx context.Context) error {
	configName := "swap_scan_tm_transfer_current_block_number"
	startBlocks, err := u.Repo.ParseConfigByInt(ctx, configName)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	contracts := []string{}
	contracts = append(contracts, u.Repo.ParseConfigByString(ctx, "swap_tm_token_contract_address"))

	eventResp, err := u.BlockChainApi.TcTmTokenEvents(contracts, 0, startBlocks, 0)
	if err != nil {
		return err
	}
	errs := u.TcTmTokenTransactionEventResp(
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

func (u *Usecase) TcTmTokenTransactionEventResp(ctx context.Context, eventResp *blockchain_api.TcTmTokenEventResp) []error {
	var err error
	var errs []error
	for _, event := range eventResp.Transfer {
		err = u.TcTmTokenCreatedTransfer(ctx, event)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (u *Usecase) TcTmTokenCreatedTransfer(ctx context.Context, eventResp *blockchain_api.TcTmTokenTransferEventResp) error {
	// check if token exist
	dbSwapPair, err := u.Repo.FindTmTransferHistory(ctx, entity.SwapTmTransferHistoriesFilter{
		TxHash: strings.ToLower(eventResp.TxHash),
		Index:  eventResp.Index,
	})
	if err != nil && err != mongo.ErrNoDocuments {
		logger.AtLog.Logger.Error("Find mongo entity failed", zap.Error(err))
		return err
	}

	if dbSwapPair != nil {
		return nil
	} else {
		swapPair := &entity.SwapTmTransferHistories{}
		swapPair.ContractAddress = strings.ToLower(eventResp.ContractAddress)
		swapPair.From = eventResp.From
		swapPair.To = eventResp.To
		swapPair.Value, _ = primitive.ParseDecimal128(helpers.ConvertWeiToBigFloat(eventResp.Value, 18).String())
		swapPair.TxHash = strings.ToLower(eventResp.TxHash)
		swapPair.Timestamp = time.Unix(int64(eventResp.Timestamp), 0)
		swapPair.Index = eventResp.Index
		_, err = u.Repo.InsertOne(swapPair)
		if err != nil {
			logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
			return err
		}
	}
	return nil
}
