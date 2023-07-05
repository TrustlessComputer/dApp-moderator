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
	sleepMinute, _ := u.Repo.ParseConfigByInt(ctx, "gm_payment_sleep_time_minute")
	if time.Now().Before(*startTime) {
		err = errors.New("invalid start time")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	userBalance, _ := u.Repo.FindUserGmBalance(ctx, query)
	if userBalance == nil {
		err = errors.New("GM Balance not found")
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}

	userGmBalance, _ := big.NewFloat(0).SetString(userBalance.Balance.String())
	if userGmBalance.Cmp(big.NewFloat(10)) >= 0 {
		timeIn := startTime.Add(time.Minute * time.Duration(sleepMinute))
		if time.Now().Before(timeIn) {
			time.Sleep(time.Minute * time.Duration(3))
			err = errors.New("Bad request")
			logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
			return nil, err
		}
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
	user1Addr := "0x670543A06bf013138C2AB47605C339e87181B65d"
	mgAmount1, _ := big.NewFloat(0).SetString("0.1")

	user2Addr := "0x94d2172638014C7D0C3C4eA0535391AdED049Fd1"
	mgAmount2, _ := big.NewFloat(0).SetString("0.2")

	contractAddr := "0xBAd802Afa594e6bbFECf7248ED031617F86171D9"
	// adminAddr := "0xBD91528e1B91AdbddF9f049e4CF5A5D9A45F1B8B"
	// decryptedPrk := "1c373998059152166f8d4c7fcfb42c5403360668d45b6acc922ef4c2c1a67f7d"
	adminAddr := "0x825794e9cca48352ED02599400170E990CAE3A04"
	decryptedPrk := "aa613f99e94701131fa99a864f7ffa3ea99674e23cadf9f4738929bff0eec775"
	tokenAddr := "0x13f86cbF0476e1D867342adE6d60164F8E26c14F"

	// mgAmount, _ := big.NewFloat(0).SetString("0.01")
	chainId, _ := new(big.Int).SetString("22213", 10)
	adminSign1, err := u.BlockChainApi.GmPaymentSignMessage(
		contractAddr,
		adminAddr,
		decryptedPrk,
		user1Addr,
		tokenAddr,
		chainId,
		helpers.EtherToWei(mgAmount1),
	)
	if err != nil {
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}
	if !strings.HasPrefix(adminSign1, "0x") {
		adminSign1 = "0x" + adminSign1
	}

	fmt.Println(adminSign1)
	fmt.Println(helpers.EtherToWei(mgAmount1).String())
	resp1 := entity.SwapUserGmClaimSignature{
		Signature: adminSign1,
		Amount:    helpers.EtherToWei(mgAmount1).String(),
	}

	///////////////////////////////////////////
	adminSign2, err := u.BlockChainApi.GmPaymentSignMessage(
		contractAddr,
		adminAddr,
		decryptedPrk,
		user2Addr,
		tokenAddr,
		chainId,
		helpers.EtherToWei(mgAmount2),
	)
	if err != nil {
		logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
		return nil, err
	}
	if !strings.HasPrefix(adminSign2, "0x") {
		adminSign2 = "0x" + adminSign2
	}

	fmt.Println(adminSign2)
	fmt.Println(helpers.EtherToWei(mgAmount2).String())
	resp2 := entity.SwapUserGmClaimSignature{
		Signature: adminSign2,
		Amount:    helpers.EtherToWei(mgAmount2).String(),
	}

	listResp := []entity.SwapUserGmClaimSignature{}
	listResp = append(listResp, resp1)
	listResp = append(listResp, resp2)

	return listResp, nil
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

func (u *Usecase) GmPaymentGenerateSignature(ctx context.Context) (interface{}, error) {
	var err error
	query := entity.SwapUserGmBalanceFilter{}
	query.ListAddress = []string{
		"0x8a08f5f944e7097add58a0d79176ca56a341d4a3",
		"0x6503516d9d61e0dcb8aa52006c540da9d52b8065",
		"0x80a6278131de681def1859e03e20b5a1e589d30a",
		"0x1f65dea6b70afbb601d263ef9985aa427ab7bd88",
		"0x0bbeb88093e8cc6b84a52613241436f1e9a9d305",
		"0x9a6f3122a1463e457049d85ef189783cd5010b86",
		"0xd42556049a4db5befe62b351f73ade0b8f3e83f8",
		"0x9b3f6db2f794de2165f6708edb666f37bdef8f29",
		"0xfe9d66ccf9baa6bab9d937c827a0e8e827e1f21d",
		"0x54ab58e6404bfe1f236d27b294c925bf767c8292",
		"0x49a395e91f23dac263980078b470be68eb67dac1",
		"0xef6afe42e79609c31a71b8e4babf6d7215c171ae",
		"0xbae736a772f2db61119e5941aa86320627a00aa9",
		"0xab261acc43152b392fb99f936de52aae73b1542a",
		"0xf8133515878ad78d5076555c77d3bd57e56fc0f3",
		"0x7145536b9e2973ab824bde9d056008e1a59afe0b",
		"0x4706d76b21bd7667c24d8457384957c72b671af0",
		"0xa8cf72de7682525dbd73c851f9e5578822013148",
		"0x60e2e4cdf845ceae2508d90233db8f3833adb9e9",
		"0x87641b0cb2f91a01077a17765fa759a687a86f12",
		"0xfe687f8baa64dbc19c7aa4092a7fd5d65486a25e",
		"0x7a01217786a6c8d64016847c2b51ef796a5fe067",
		"0x851b10a2571eec8efce68fb75a72d1c77fb276c5",
		"0xd4984286554a91e66506732e3f470ce45bcea4d8",
		"0xecb84eb73cf4ed352dea6d1f1672ed384552e9af",
		"0x46e52013acdf018a9b46e0cfe4c41c4c0f90c9da",
		"0xa9e91ded664de6c8bc0bc43be43d2959879543aa",
		"0xb5beeb865f1fb98142471baba6dbdffe81bc7cf6",
		"0xa5717bf1c67c7c64bd441935a4cac897fe75ce8d",
		"0x65b0170fb3d4cc2646d093d0d094d2344d36e39d",
		"0xa0b1b2c371e4d5c3b40330dc74f359bd4af63726",
		"0x461b1e5724ab97bf73931c5760be9c5b20c3964c",
		"0x94d2172638014c7d0c3c4ea0535391aded049fd1",
		"0x5ff692088ce9496eb254dc6a7d79e2e9465700b0",
		"0xed7edf853e37f9da27c7da8a4f012237f4d06d8b",
		"0x670543a06bf013138c2ab47605c339e87181b65d",
		"0xc29b6bad9abaa4f7585526ac703fbf2b320b6f94",
		"0x7004a3df94eeb372042d4cd82386d9745d6cb1b8",
	}

	mapAddress := map[string]string{
		"0x8a08f5f944e7097add58a0d79176ca56a341d4a3": "0x0c76140c49e7a85c0d37783ea258722e89102a1e",
		"0x6503516d9d61e0dcb8aa52006c540da9d52b8065": "0x7a919d232823e5fecc9bb89a9205715064033d66",
		"0x80a6278131de681def1859e03e20b5a1e589d30a": "0x555c42320b6334253e2fc7bc6888305d6a5d988d",
		"0x1f65dea6b70afbb601d263ef9985aa427ab7bd88": "0x3cd863cf1d3e88316333245597a5b88fb357c102",
		"0x0bbeb88093e8cc6b84a52613241436f1e9a9d305": "0x6e9310b70d0440da9aea59f613ea484b57161ae9",
		"0x9a6f3122a1463e457049d85ef189783cd5010b86": "0x5edf0dfbd5e8a023c4dd55f0725cbe84c1ab2f69",
		"0xd42556049a4db5befe62b351f73ade0b8f3e83f8": "0x52f54f1bf61dbeee854c20e27e6346d59f91ead1",
		"0x9b3f6db2f794de2165f6708edb666f37bdef8f29": "0xb0d6c10715d6a85ae403403548f1d9a26e6adf02",
		"0xfe9d66ccf9baa6bab9d937c827a0e8e827e1f21d": "0x1f056fa0d63afa27f0899af9d9ab54c67a25f01f",
		"0x54ab58e6404bfe1f236d27b294c925bf767c8292": "0x80394c57eb082d35c7ea73239992454ab768807b",
		"0x49a395e91f23dac263980078b470be68eb67dac1": "0xe49f423670dd32bd6b24d941fc5d353f4c902dd9",
		"0xef6afe42e79609c31a71b8e4babf6d7215c171ae": "0x15063f48160e9ef8554416d32ce4cbb26fee462e",
		"0xbae736a772f2db61119e5941aa86320627a00aa9": "0x874d9c5f7cce9e06cd9742d1ef4a204ce9e5b175",
		"0xab261acc43152b392fb99f936de52aae73b1542a": "0x3981768a3b0e36be3746c0e88c913765c2a1411a",
		"0xf8133515878ad78d5076555c77d3bd57e56fc0f3": "0x9b7b50480df5c5cb221a1991c50e8d9625680ee7",
		"0x7145536b9e2973ab824bde9d056008e1a59afe0b": "0xa43b84f2be90ebadb4bdfcd38fef7422e41ec425",
		"0x4706d76b21bd7667c24d8457384957c72b671af0": "0x7a25a709503f1ba5ffc4ef78c9779a2813582e8a",
		"0xa8cf72de7682525dbd73c851f9e5578822013148": "0x28d0dc6e29b46bc8d21cd1c1a9622c7b384f1878",
		"0x60e2e4cdf845ceae2508d90233db8f3833adb9e9": "0xd4549c2a55dc95746e488b8976dfd5adf9c7441e",
		"0x87641b0cb2f91a01077a17765fa759a687a86f12": "0x568c229a40a03fdba9a23c854c24388f3b68aa6c",
		"0xfe687f8baa64dbc19c7aa4092a7fd5d65486a25e": "0xc99a2ee4cef26473daf9ef553f5673e6b1f5aeaf",
		"0x7a01217786a6c8d64016847c2b51ef796a5fe067": "0x960252ae3c22636ad721792c1b3d06f1df9d2b53",
		"0x851b10a2571eec8efce68fb75a72d1c77fb276c5": "0x81522aa51c3f98af67cd3d49735b09d805932f96",
		"0xd4984286554a91e66506732e3f470ce45bcea4d8": "0x0297452097f55ed93e1d06695f9b6fb0294aca76",
		"0xecb84eb73cf4ed352dea6d1f1672ed384552e9af": "0x2fcc16bb6c6cc7528a5ca32121ce661d55c0a5fb",
		"0x46e52013acdf018a9b46e0cfe4c41c4c0f90c9da": "0x0eaf92059fdbdf86a9bcdf1cc99658b8a70995f1",
		"0xa9e91ded664de6c8bc0bc43be43d2959879543aa": "0xc12a205be940a7bc1b604e770ed2d9aacd0e1ada",
		"0xb5beeb865f1fb98142471baba6dbdffe81bc7cf6": "0x91c9e5279cc51cec5789dda21a2df59cd26ec43b",
		"0xa5717bf1c67c7c64bd441935a4cac897fe75ce8d": "0xc5dd0224f10fed0a173a1ef13fad37b0cf44a27b",
		"0x65b0170fb3d4cc2646d093d0d094d2344d36e39d": "0x76b7e810f7fc39ddcbffcc8ac8122c5c2f6daa1a",
		"0xa0b1b2c371e4d5c3b40330dc74f359bd4af63726": "0x1ce7d875753ff327e411799714b16ad82c0aaad9",
		"0x461b1e5724ab97bf73931c5760be9c5b20c3964c": "0x8137c193d0c99fd3a49db9a88495577ceb158a7a",
		"0x94d2172638014c7d0c3c4ea0535391aded049fd1": "0x391c4eb280b2a3c3ab8b666b41cb88d96d249d50",
		"0x5ff692088ce9496eb254dc6a7d79e2e9465700b0": "0xa1b43cb8514d25e720523d8f79606a4e837c9ddd",
		"0xed7edf853e37f9da27c7da8a4f012237f4d06d8b": "0x3903e9195355bca1e0d2a2834dc226bfe19f87d0",
		"0x670543a06bf013138c2ab47605c339e87181b65d": "0x1df49c9073ab4f560748f4e8a7dd8a66ae8d1167",
		"0xc29b6bad9abaa4f7585526ac703fbf2b320b6f94": "0x2e8cebca515381b0ea47e34c1d79a817679061f5",
		"0x7004a3df94eeb372042d4cd82386d9745d6cb1b8": "0x0ddc5e07c370e5f1fba733a08eac48c5386dbd89",
	}

	userBalances, _ := u.Repo.FindListUserGmBalance(ctx, query)

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

	resp := entity.SwapUserGmClaimSignature{}
	for _, userBalance := range userBalances {
		oldAddress := userBalance.UserAddress
		newUserAddress := mapAddress[userBalance.UserAddress]
		mgAmount, _ := big.NewFloat(0).SetString(userBalance.Balance.String())
		chainId, _ := new(big.Int).SetString(config.GmPaymentChainId, 10)
		adminSign, err := u.BlockChainApi.GmPaymentSignMessage(
			config.GmPaymentContractAddr,
			config.GmPaymentAdminAddr,
			decryptedPrk,
			newUserAddress,
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

		userBalance.UserAddress = newUserAddress
		userBalance.BalanceSign = helpers.EtherToWei(mgAmount).String()
		userBalance.Signature = adminSign
		err = u.Repo.UpdateSwapUserGmBalanceWithAddress(ctx, oldAddress, userBalance)
		if err != nil {
			logger.AtLog.Logger.Error("GmPaymentClaim", zap.Error(err))
			return nil, err
		}
	}

	logger.AtLog.Logger.Info("GmPaymentClaim", zap.Any("data", resp))
	return resp, nil
}

func (u *Usecase) AddGmbalanceFromFile(ctx context.Context) error {
	f, err := os.Open("/Users/autonomous/Desktop/gm_results_batch_2.csv")
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
		}
	}
	u.Repo.InsertMany(listData)
	return nil
}
