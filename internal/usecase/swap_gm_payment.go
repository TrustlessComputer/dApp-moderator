package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"math/big"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *Usecase) GmPaymentClaim(ctx context.Context, userAddress string) (interface{}, error) {
	var data interface{}
	var err error
	query := entity.SwapUserGmBalanceFilter{}
	query.Address = strings.ToLower(userAddress)

	config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
	adminWallet, err := u.SwapGetWalletAddress(ctx, config.GmPaymentAdminAddr)
	if err != nil {
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	userBalance, err := u.Repo.FindUserGmBalance(ctx, query)
	if err != nil {
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

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
		return adminSign, nil
	}

	logger.AtLog.Logger.Info("SwapFindSwapIdoDetail", zap.Any("data", data))
	return "", nil
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
