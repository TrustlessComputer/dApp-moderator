package bns_service

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"fmt"
	"net/url"
)

type BNSService struct {
	conf      *config.Config
	serverURL string
	cache     redis.IRedisCache
}

func NewBNSService(conf *config.Config, cache redis.IRedisCache) *BNSService {
	return &BNSService{
		conf:      conf,
		serverURL: conf.BNSService,
		cache:     cache,
	}
}

func (q BNSService) Names(params url.Values) ([]*NameResp, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/names?%s", q.serverURL, params.Encode())

	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp.ToNames(), nil
}

func (q BNSService) Name(name string) (*NameResp, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/name/%s", q.serverURL, name)

	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp.ToName(), nil
}

func (q BNSService) NameByToken(tokenID string) (*NameResp, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/token/%s", q.serverURL, tokenID)

	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp.ToName(), nil
}

func (q BNSService) NameAvailable(name string) (*bool, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/available/%s", q.serverURL, name)

	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp.ToAvailable(), nil
}

func (q BNSService) NameOnwedByWalletAddress(walletAddress string, params url.Values) ([]*NameResp, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/names/%s?%s", q.serverURL, walletAddress, params.Encode())

	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp.ToNames(), nil
}

func (q BNSService) ParseData(data []byte) (*ServiceResp, error) {
	resp := &ServiceResp{}
	err := helpers.ParseData(data, resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp, nil
}
