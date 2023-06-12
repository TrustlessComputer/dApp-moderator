package usecase

import (
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/contracts/erc20"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"math/big"
	"strconv"
	"strings"
	"sync"
)

type CheckGMBalanceOutputChan struct {
	Nft     entity.Nfts
	Err     error
	Balance *big.Int
}

func (u *Usecase) SoulCrontab() error {
	maxProcess := 10
	minBalance := float64(1)
	erc20Addr := "0x2fe8d5A64afFc1d703aECa8a566f5e9FaeE0C003"
	instance, err := erc20.NewErc20(common.HexToAddress(erc20Addr), u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	collection, err := u.Repo.GetSoulCollection()
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	nfts, err := u.Repo.SoulNfts(collection.Contract)
	if err != nil {
		logger.AtLog.Logger.Error("SoulCrontab", zap.Error(err))
		return err
	}

	inputChan := make(chan entity.Nfts, len(nfts))
	outputChan := make(chan CheckGMBalanceOutputChan, len(nfts))
	wg := sync.WaitGroup{}
	logger.AtLog.Logger.Info("SoulCrontab", zap.String("contract_address", collection.Contract), zap.Int("nfts", len(nfts)))

	for i := 0; i < len(nfts); i++ {
		go u.CheckGMBalanceWorker(&wg, instance, inputChan, outputChan)
	}

	for i, nft := range nfts {
		wg.Add(1)
		inputChan <- nft
		if i%maxProcess == 0 && i > 0 {
			wg.Wait()
		}
	}

	for i := 0; i < len(nfts); i++ {
		out := <-outputChan
		if out.Err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.Error(out.Err))
			continue
		}

		tokenIDInt, err := strconv.Atoi(out.Nft.TokenID)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.Error(err))
			continue
		}

		isAuction := false

		value := helpers.GetValue(fmt.Sprintf("%d", out.Balance.Int64()), 18)
		if value < minBalance {
			isAuction = true
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("SoulCrontab - %s - %s", out.Nft.ContractAddress, out.Nft.TokenID), zap.String("contract_address", collection.Contract), zap.String("token_id", out.Nft.TokenID), zap.String("owner", out.Nft.Owner), zap.String("balance", fmt.Sprintf("%d", out.Balance.Int64())))

		insertData := &entity.NftAuctions{
			TokenID:         out.Nft.TokenID,
			TokenIDInt:      int64(tokenIDInt),
			ContractAddress: strings.ToLower(out.Nft.ContractAddress),
			IsAuction:       isAuction,
		}

		err = u.Repo.InsertAuction(insertData)
	}
	return nil
}

func (u *Usecase) CheckGMBalanceWorker(wg *sync.WaitGroup, erc20Instance *erc20.Erc20, input chan entity.Nfts, output chan CheckGMBalanceOutputChan) {
	defer wg.Done()
	nft := <-input

	owner := nft.Owner
	balanceOf, err := erc20Instance.BalanceOf(nil, common.HexToAddress(owner))

	output <- CheckGMBalanceOutputChan{
		Nft:     nft,
		Balance: balanceOf,
		Err:     err,
	}
}
