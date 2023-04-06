package token_explorer

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"fmt"
	"net/url"
)

type TokenExplorer struct {
	conf      *config.Config
	serverURL string
	cache     redis.IRedisCache
}

func NewTokenExplorer(conf *config.Config, cache redis.IRedisCache) *TokenExplorer {
	return &TokenExplorer{
		conf:      conf,
		serverURL: conf.TokenExplorer,
		cache:     cache,
	}
}

func (q *TokenExplorer) Tokens(params url.Values) ([]Token, error) {
	headers := make(map[string]string)
	data, _, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s?%s", q.serverURL, "tokens", params.Encode()), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp.ToTokens()
}

func (q *TokenExplorer) ParseData(data []byte) (*Response, error) {
	resp := &Response{}
	err := helpers.ParseData(data, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != "OK" {
		return nil, resp.Error
	}

	return resp, nil
}
