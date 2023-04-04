package quicknode

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"encoding/json"
)

type QuickNode struct {
	conf *config.Config
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


func (q QuickNode) AddressBalance(walletAddress string) ([]WalletAddressBalanceResp, error) {
	headers := make(map[string]string)
	reqBody := RequestData{
		Method: "qn_addressBalance",
		Params: []string{
			walletAddress,
		},
	}
	
	data, _, err := helpers.HttpRequest(q.serverURL, "POST", headers, reqBody)
	if err != nil {
		return nil, err
	}

	resp := []WalletAddressBalanceResp{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	
	return resp, nil
}
