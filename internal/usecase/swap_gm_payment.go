package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *Usecase) TestGG(ctx context.Context) (interface{}, error) {
	encryptedText, err := helpers.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_WALLET_PRIVATE_KEY_ENCRYPTED"))
	if err != nil {
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	walletCipherKey, err := helpers.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_ENCRYPTED_SAT"))
	if err != nil {
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	m := make(map[string]string)
	m["GSM_KEY_NAME__DAPP_TOKEN_WALLET_PRIVATE_KEY_ENCRYPTED"] = encryptedText
	m["GSM_KEY_NAME__DAPP_TOKEN_ENCRYPTED_SAT"] = walletCipherKey
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
	encryptedText, _ := u.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_WALLET_PRIVATE_KEY_ENCRYPTED"))
	walletCipherKey, _ := u.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_ENCRYPTED_SAT"))

	if encryptedText == "" || walletCipherKey == "" {
		err = errors.New("Cannot get encrypted key")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	decryptedPrk, err := helpers.DecryptToString(encryptedText, walletCipherKey)
	if err != nil {
		err = errors.New("Cannot decrypted prk")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

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

func (u *Usecase) GetGoogleSecretKey(keyName string) (string, error) {
	redisKey := fmt.Sprintf("tc-swap:google-secret-key-%s", keyName)
	exists, err := u.Cache.Exists(redisKey)
	if err != nil {
		logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
		return "", err
	}
	if *exists {
		encryptedText, err := u.Cache.GetData(redisKey)
		if err != nil {
			logger.AtLog.Logger.Error("c.Cache.Exists", zap.String("redisKey", redisKey), zap.Error(err))
			return "", err
		}
		return *encryptedText, nil
	} else {
		encryptedText, err := helpers.GetGoogleSecretKey(keyName)
		if err != nil {
			err = errors.New("Cannot get encryptedText")
			logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
			return "", err
		}
		err = u.Cache.SetStringDataWithExpTime(redisKey, encryptedText, 30*60)
		if err != nil {
			logger.AtLog.Logger.Error("Save the last fetched page to redis failed", zap.Error(err))
			return encryptedText, nil
		}

		return encryptedText, nil
	}
}

func (u *Usecase) GmPaymentClaimTestnet(ctx context.Context, userAddress string) (interface{}, error) {
	var err error
	query := entity.SwapUserGmBalanceFilter{}
	query.Address = strings.ToLower(userAddress)

	config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
	userBalance, _ := u.Repo.FindUserGmBalance(ctx, query)
	if userBalance == nil {
		err = errors.New("GM Balance not found")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	decryptedPrk := "1c373998059152166f8d4c7fcfb42c5403360668d45b6acc922ef4c2c1a67f7d"
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

func (u *Usecase) GmPaymentClaimTestMainnet(ctx context.Context, userAddress string) (interface{}, error) {
	var err error
	query := entity.SwapUserGmBalanceFilter{}
	query.Address = strings.ToLower(userAddress)

	dbPaidGm, _ := u.Repo.FindUserGmPaid(ctx, entity.SwapUserGmPaidFilter{
		Address: strings.ToLower(userAddress),
	})

	if dbPaidGm != nil {
		err = errors.New("User is already claimed")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	config, _ := u.TcSwapGetWrapTokenContractAddr(ctx)
	encryptedText, _ := u.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_WALLET_PRIVATE_KEY_ENCRYPTED"))
	walletCipherKey, _ := u.GetGoogleSecretKey(os.Getenv("GSM_KEY_NAME__DAPP_TOKEN_ENCRYPTED_SAT"))

	if encryptedText == "" || walletCipherKey == "" {
		err = errors.New("Cannot get encrypted key")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	decryptedPrk, err := helpers.DecryptToString(encryptedText, walletCipherKey)
	if err != nil {
		err = errors.New("Cannot decrypted prk")
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

func (u *Usecase) AddGmbalanceFromFile(ctx context.Context) error {
	f, err := os.Open("/Users/autonomous/Desktop/gm_results.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	listData := []entity.IEntity{}
	for i, row := range data {
		if i != 0 {
			swapPairSync := &entity.SwapUserGmBalance{}
			swapPairSync.UserAddress = strings.ToLower(row[0])
			swapPairSync.Balance, _ = primitive.ParseDecimal128(row[1])
			swapPairSync.IsContract = false
			if row[2] == "1" {
				swapPairSync.IsContract = true
			}
			listData = append(listData, swapPairSync)
			// _, err := u.Repo.InsertOne(swapPairSync)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		}
	}
	u.Repo.InsertMany(listData)
	return nil
}
