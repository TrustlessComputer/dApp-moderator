package usecase

import (
	"context"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func (u *Usecase) AddressBalance(ctx context.Context, walletAddress string) ([]structure.WalletAddressBalanceResp, error) {
	resp := []structure.WalletAddressBalanceResp{}

	outputs, err := u.QuickNode.AddressBalance(walletAddress)
	if err != nil {
		logger.AtLog.Logger.Error("AddressBalance", zap.String("walletAddress", walletAddress), zap.Error(err))
	}

	for _, output := range outputs {
		tmp := &structure.WalletAddressBalanceResp{}
		err = helpers.JsonTransform(output, tmp)
		if err != nil {
			continue
		}

		/*out := fmt.Sprintf("%s:%d", output.Hash, output.Index)
		data, err := u.GetInscriptionByOutput(out)
		if err != nil {
			continue
		}*/

		/*if len(data.Inscriptions) > 0 {
			tmp.IsOrdinal = true
		}*/

		resp = append(resp, *tmp)

	}

	logger.AtLog.Logger.Info("AddressBalance", zap.String("walletAddress", walletAddress), zap.Any("data", outputs))
	return resp, err
}

func (u *Usecase) GetInscriptionByOutput(ouput string) (*structure.InscriptionByOutput, error) {
	ordServer := os.Getenv("CUSTOM_ORD_SERVER")
	if ordServer == "" {
		ordServer = "https://dev-v5.generativeexplorer.com"
	}

	url := fmt.Sprintf("%s/api/output/%s", ordServer, ouput)
	headers := make(map[string]string)

	resp, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		logger.AtLog.Logger.Error("getInscriptionByOutput", zap.String("url", url), zap.Error(err))
		return nil, err
	}

	data := &structure.InscriptionByOutput{}
	err = helpers.ParseData(resp, data)
	if err != nil {
		logger.AtLog.Logger.Error("getInscriptionByOutput", zap.String("url", url), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("getInscriptionByOutput", zap.String("url", url), zap.Any("data", data))

	return data, nil
}
