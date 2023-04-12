package block_stream

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"fmt"
)

type BlockStream struct {
	conf *config.Config
	serverURL string
	cache     redis.IRedisCache
}

func NewBlockStream(conf *config.Config, cache redis.IRedisCache) *BlockStream {
	return &BlockStream{
		conf:      conf,
		serverURL: conf.BlockStream,
		cache:     cache,
	}
}


func (q BlockStream) Txs(walletAddress string) (interface{}, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/address/%s/txs", q.serverURL, walletAddress)
	data,_, _, err := helpers.HttpRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}
