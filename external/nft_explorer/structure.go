package nft_explorer

import "dapp-moderator/utils/helpers"

type RequestData struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type ServiceResp struct {
	Code   string      `json:"code"`
	Error  error       `json:"error"`
	Result interface{} `json:"result"`
}

type CollectionsResp struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Contract     string `json:"contract"`
	ContractType string `json:"contract_type"`
	Creator      string `json:"creator"`
	Description  string `json:"description"`
	TotalItems   int    `json:"total_items"`
	TotalOwners  int    `json:"total_owners"`
	Cover        string `json:"cover"`
	Thumbnail    string `json:"thumbnail"`
}

type NftsResp struct {
	Collection        string  `json:"collection"`
	ContractAddress string  `json:"collection_address"`
	TokenID           string  `json:"token_id"`
	ContentType       string  `json:"content_type"`
	Name              string  `json:"name"`
	Owner             string  `json:"owner"`
	MintedAt          float64 `json:"mintedAt"`
	Attributes        []NftAttr `json:"attributes"`
}

type NftAttr struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

func (sr ServiceResp) ToCollections() []CollectionsResp {
	resp := []CollectionsResp{}
	err := helpers.JsonTransform(sr.Result, &resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToCollection() *CollectionsResp {
	resp := &CollectionsResp{}
	err := helpers.JsonTransform(sr.Result, resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToNfts() []NftsResp {
	resp := []NftsResp{}
	err := helpers.JsonTransform(sr.Result, &resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToNft() *NftsResp {
	resp := &NftsResp{}
	err := helpers.JsonTransform(sr.Result, resp)
	if err == nil {
		return resp
	}

	return resp
}
