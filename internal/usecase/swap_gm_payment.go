package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"errors"
	"math/big"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *Usecase) TestGG(ctx context.Context) (interface{}, error) {
	// prk := "1c373998059152166f8d4c7fcfb42c5403360668d45b6acc922ef4c2c1a67f7d"
	// prkEncrypted, err := helpers.EncryptToString(prk, os.Getenv("GM_PAYMENT_SALT"))
	// if err != nil {
	// 	err = errors.New("Cannot get encryptedText")
	// 	logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
	// 	return nil, err
	// }

	// decryptedPrk, err := helpers.DecryptToString("AAAAAAAAAADBKOfbj4B7wdqlEc/JYV/nBrm4kcpE23mEU3vG/8ir3aVCYL+ttlCtA6hzWpwgr/ZFTdNPaMPA7+daCmTliWbD", "au3Cao8NSguLZAgIpZkquvrVyjutEzct")
	// if err != nil {
	// 	err = errors.New("Cannot decrypted prk")
	// 	logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
	// 	return nil, err
	// }

	encryptedText, err := helpers.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_WALLET_PRIVATE_KEY_ENCRYPTED"))
	if err != nil {
		err = errors.New("Cannot get encryptedText")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	walletCipherKey, err := helpers.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_ENCRYPTED_SAT"))
	if err != nil {
		err = errors.New("Cannot get encryptedText")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	m := make(map[string]string)
	m["GM_PAYMENT_PRIVATE_KEY"] = encryptedText
	m["GM_PAYMENT_SALT"] = walletCipherKey
	// m["PRK"] = prkEncrypted
	// m["decryptedPrk"] = decryptedPrk
	return m, nil
}

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

	dbPaidGm, _ := u.Repo.FindUserGmPaid(ctx, entity.SwapUserGmPaidFilter{
		Address: strings.ToLower(userAddress),
	})

	if dbPaidGm != nil {
		err = errors.New("User is already claimed")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
	encryptedText, err := helpers.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_WALLET_PRIVATE_KEY_ENCRYPTED"))
	if err != nil {
		err = errors.New("Cannot get encryptedText")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	walletCipherKey, err := helpers.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_ENCRYPTED_SAT"))
	if err != nil {
		err = errors.New("Cannot get encryptedText")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	decryptedPrk, err := helpers.DecryptToString(encryptedText, walletCipherKey)
	if err != nil {
		err = errors.New("Cannot decrypted prk")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	// adminWallet, err := u.SwapGetWalletAddress(ctx, config.GmPaymentAdminAddr)
	// if err != nil {
	// 	logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
	// 	return nil, err
	// }

	userBalance, _ := u.Repo.FindUserGmBalance(ctx, query)
	if userBalance == nil {
		err = errors.New("GM Balance not found")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	if userBalance.IsContract {
		timeIn := startTime.Add(time.Minute * time.Duration(15))
		if time.Now().Before(timeIn) {
			time.Sleep(time.Minute * time.Duration(5))
			err = errors.New("Bad request")
			logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
			return nil, err
		}
	}

	resp := entity.SwapUserGmClaimSignature{}
	if userBalance != nil {
		mgAmount, _ := big.NewFloat(0).SetString(userBalance.Balance.String())
		chainId, _ := new(big.Int).SetString(config.GmPaymentChainId, 10)
		adminSign, err := u.BlockChainApi.GmPaymentSignMessage(
			config.GmPaymentContractAddr,
			config.GmPaymentAdminAddr,
			decryptedPrk,
			userAddress,
			config.GmTokenContractAddr,
			chainId,
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
