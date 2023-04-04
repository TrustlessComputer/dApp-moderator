package nft_explorer

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"fmt"
)

type NftExplorer struct {
	conf *config.Config
	serverURL string
	cache     redis.IRedisCache
}

func NewNftExplorer(conf *config.Config, cache redis.IRedisCache) *NftExplorer {
	return &NftExplorer{
		conf:      conf,
		serverURL: conf.NftExplorer,
		cache:     cache,
	}
}

func (q NftExplorer) Collections() ([]CollectionsResp, error) {
	headers := make(map[string]string)	
	data, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s",q.serverURL, "collections"), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	
	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToCollections(), nil
}

func (q NftExplorer) CollectionDetail(contractAddress string) (*CollectionsResp, error) {
	headers := make(map[string]string)	
	data, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s/%s",q.serverURL, "collection", contractAddress), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	
	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToCollection(), nil
}

func (q NftExplorer) CollectionNfts(contractAddress string) ([]NftsResp, error) {
	headers := make(map[string]string)	
	data, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s/%s/nfts",q.serverURL, "collection", contractAddress), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	
	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToNfts(), nil
}

func (q NftExplorer) CollectionNftDetail(contractAddress string, tokenID string) (*NftsResp, error) {
	headers := make(map[string]string)	
	fullURL := fmt.Sprintf("%s/%s/%s/nft/%s",q.serverURL, "collection", contractAddress, tokenID)
	
	data, _, err := helpers.JsonRequest(fullURL, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	
	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToNft(), nil
}

func (q NftExplorer) CollectionNftContent(contractAddress string, tokenID string) ([]byte, string, error) {
	headers := make(map[string]string)	
	data, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s/%s/nft/%s/content",q.serverURL, "collection", contractAddress, tokenID), "GET", headers, nil)
	if err != nil {
		return nil, "", err
	}

	nft, err := q.CollectionNftDetail(contractAddress, tokenID)
	if err != nil {
		return nil, "", err
	}

	return data, nft.ContentType,  nil
}

func (q NftExplorer) Nfts() ([]NftsResp, error) {
	headers := make(map[string]string)	
	data, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s",q.serverURL, "nfts"), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	
	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToNfts(), nil
}

func (q NftExplorer) NftOfWalletAddress(walletAddress string) ([]NftsResp, error) {
	headers := make(map[string]string)	
	data, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s/%s",q.serverURL, "nfts",walletAddress), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToNfts(), nil
}

func (q NftExplorer) ParseData(data []byte) (*ServiceResp, error) {
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