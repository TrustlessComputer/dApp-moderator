package nft_explorer

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type NftExplorer struct {
	conf      *config.Config
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

func (q NftExplorer) Collections(params url.Values) ([]CollectionsResp, error) {
	headers := make(map[string]string)
	data, _, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s?%s", q.serverURL, "collections", params.Encode()), "GET", headers, nil)
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
	data, _, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s/%s", q.serverURL, "collection", contractAddress), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp.ToCollection(), nil
}

func (q NftExplorer) CollectionNfts(contractAddress string, params url.Values) ([]*NftsResp, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/%s/%s/nfts?%s", q.serverURL, "collection", contractAddress, params.Encode())
	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	res := resp.ToNfts()
	q.FillDataMultiple(res)
	return res, nil
}

func (q NftExplorer) CollectionNftDetail(contractAddress string, tokenID string) (*NftsResp, error) {
	headers := make(map[string]string)
	fullURL := fmt.Sprintf("%s/%s/%s/nft/%s", q.serverURL, "collection", contractAddress, tokenID)

	data, _, _, err := helpers.JsonRequest(fullURL, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	rs := resp.ToNft()
	q.FillData(rs)
	return rs, nil
}

func (q NftExplorer) CollectionNftContent(contractAddress string, tokenID string) ([]byte, string, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/%s/%s/nft/%s/content", q.serverURL, "collection", contractAddress, tokenID)

	data, resHeader, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, "", err
	}

	contentType := resHeader.Get("content-type")
	txt := string(data)
	if strings.Contains(txt, `<svg xmlns="http://www.w3.org/2000/svg" `) {
		contentType = "image/svg+xml"
	}

	return data, contentType, nil
}

func (q NftExplorer) Nfts(params url.Values) ([]*NftsResp, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/nfts?%s", q.serverURL, params.Encode())

	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	res := resp.ToNfts()
	q.FillDataMultiple(res)
	return res, nil
}

func (q NftExplorer) NftOfWalletAddress(walletAddress string, params url.Values) ([]*NftsResp, error) {
	headers := make(map[string]string)
	url := fmt.Sprintf("%s/nfts/%s?%s", q.serverURL, walletAddress, params.Encode())
	data, _, _, err := helpers.JsonRequest(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	res := resp.ToNfts()
	q.FillDataMultiple(res)
	return res, nil
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

func (q *NftExplorer) FillData(nft *NftsResp) {
	nft.Image = fmt.Sprintf("%s/dapp/api/nft-explorer/collections/%s/nfts/%s/content", os.Getenv("URL"), nft.ContractAddress, nft.TokenID)
	nft.TokenURI = fmt.Sprintf("%s/dapp/api/nft-explorer/collections/%s/nfts/%s", os.Getenv("URL"), nft.ContractAddress, nft.TokenID)

	if strings.Index(nft.ContentType, "image") == -1 {
		if strings.Index(nft.ContentType, "json") != -1 {
			bytes, _, err := q.CollectionNftContent(nft.ContractAddress, nft.TokenID)
			if err != nil {
				return
			}

			nft.Metadata = string(bytes)

			erc721 := &Erc721{}
			err = helpers.ParseData(bytes, erc721)
			if err != nil {
				return
			}
			nft.Image = erc721.Image
		}
	}

}

func (q *NftExplorer) FillDataMultiple(nfts []*NftsResp) {
	for _, nft := range nfts {
		q.FillData(nft)
	}
}

func (q NftExplorer) RefreshNft(contractAddress string, tokenID string) (*ServiceResp, error) {
	headers := make(map[string]string)
	data, _, _, err := helpers.JsonRequest(fmt.Sprintf("%s/%s/%s/%s", q.serverURL, "refresh-nft", contractAddress, tokenID), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
