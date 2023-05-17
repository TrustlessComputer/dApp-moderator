package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"errors"
	"math/big"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *Usecase) GmPaymentClaim(ctx context.Context, userAddress string) (interface{}, error) {
	var err error
	query := entity.SwapUserGmBalanceFilter{}
	query.Address = strings.ToLower(userAddress)

	startTime := u.Repo.ParseConfigByTime(ctx, "gm_payment_start_time")
	if time.Now().Before(*startTime) {
		err = errors.New("invalid start time")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
	adminWallet, err := u.SwapGetWalletAddress(ctx, config.GmPaymentAdminAddr)
	if err != nil {
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	userBalance, _ := u.Repo.FindUserGmBalance(ctx, query)
	if userBalance == nil {
		err = errors.New("GM Balance not found")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	resp := entity.SwapUserGmClaimSignature{}
	if userBalance != nil {
		mgAmount, _ := big.NewFloat(0).SetString(userBalance.Balance.String())
		chainId, _ := big.NewFloat(0).SetString(config.GmPaymentChainId)
		adminSign, err := u.BlockChainApi.GmPaymentSignMessage(
			config.GmPaymentContractAddr,
			config.GmPaymentAdminAddr,
			adminWallet.Prk,
			userAddress,
			config.GmTokenContractAddr,
			helpers.EtherToWei(chainId),
			helpers.EtherToWei(mgAmount),
		)
		if err != nil {
			logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
			return nil, err
		}
		if !strings.HasPrefix(adminSign, "0x") {
			adminSign = "0x" + adminSign
		}

		resp = entity.SwapUserGmClaimSignature{
			Signature: adminSign,
			Amount:    helpers.EtherToWei(mgAmount).String(),
		}

		return resp, nil
	}

	logger.AtLog.Logger.Info("GmPaymentClaim", zap.Any("data", resp))
	return resp, nil
}

func (u *Usecase) AddTestGmbalance(ctx context.Context, userAddress string) (interface{}, error) {
	swapPairSync := &entity.SwapUserGmBalance{}
	swapPairSync.UserAddress = strings.ToLower(userAddress)
	swapPairSync.Balance, _ = primitive.ParseDecimal128("1000")
	_, err := u.Repo.InsertOne(swapPairSync)
	if err != nil {
		logger.AtLog.Logger.Error("Insert mongo entity failed", zap.Error(err))
		return false, nil
	}
	return true, nil
}
