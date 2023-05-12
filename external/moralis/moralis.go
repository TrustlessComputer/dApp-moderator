package moralis

import (
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"dapp-moderator/utils/redis"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"os"
)

type MoralisService struct {
	serverURL string
	key       string
	chain     string
	cache     redis.IRedisCache
}

func NewMoralisService(cache redis.IRedisCache) *MoralisService {
	return &MoralisService{
		cache:     cache,
		serverURL: os.Getenv("MORALIS_URL"),
		key:       os.Getenv("MORALIS_KEY"),
		chain:     os.Getenv("MORALIS_CHAIN"),
	}
}

func (m *MoralisService) Erc20TokenBalanceByWallet(walletAddr string, erc20Addresses []string) ([]Erc20BalanceResp, error) {
	headers := make(map[string]string)
	headers["X-API-Key"] = m.key

	fullUrl := fmt.Sprintf(Erc20TokenBalance, m.serverURL, walletAddr)

	params := url.Values{}
	params.Set("chain", "eth")

	for key, erc20Address := range erc20Addresses {
		params.Set(fmt.Sprintf("token_addresses[%d]", key), erc20Address)
	}

	fullUrl = fullUrl + "?" + params.Encode()
	data, _, _, err := helpers.JsonRequest(fullUrl, "GET", headers, nil)
	if err != nil {
		logger.AtLog.Logger.Error("Erc20TokenBalanceByWallet", zap.String("url", fullUrl), zap.Error(err))
		return nil, err
	}

	resp := []Erc20BalanceResp{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		logger.AtLog.Logger.Error("Erc20TokenBalanceByWallet", zap.String("url", fullUrl), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Logger.Info("Erc20TokenBalanceByWallet", zap.String("url", fullUrl), zap.Any("resp", resp))
	return resp, nil
}
