package quicknode

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"encoding/json"
	"strconv"
)

type QuickNode struct {
	conf      *config.Config
	serverURL string
	cache     redis.IRedisCache
}

func NewQuickNode(conf *config.Config, cache redis.IRedisCache) *QuickNode {
	return &QuickNode{
		conf:      conf,
		serverURL: conf.QuickNode,
		cache:     cache,
	}
}

type QuickNodeUTXO_NEW struct {
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	Value         string `json:"value"`
	Height        int    `json:"height"`
	Confirmations int    `json:"confirmations"`
}

func (q QuickNode) AddressBalance(walletAddress string) ([]WalletAddressBalanceResp, error) {
	headers := make(map[string]string)
	reqBody := RequestData{
		Method: "bb_getutxos",
		Params: []string{
			walletAddress,
		},
	}

	data, _, _, err := helpers.HttpRequest(q.serverURL, "POST", headers, reqBody)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Result []QuickNodeUTXO_NEW `json:"result"`
	}{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	result := []WalletAddressBalanceResp{}
	for _, v := range resp.Result {
		value, err := strconv.ParseUint(v.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, WalletAddressBalanceResp{
			Version:  0,
			Height:   int64(v.Height),
			Script:   "",
			Address:  walletAddress,
			Coinbase: false,
			Hash:     v.Txid,
			Index:    v.Vout,
			Value:    value,
		})
	}

	return result, nil
}
