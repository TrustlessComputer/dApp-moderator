package bns_service

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

type BrowsedFileResp struct {
	Files   []string `json:"files"`
	Folders []string `json:"folders"`
	Name    string   `json:"name"`
}

type NameResp struct {
	Owner    string `json:"owner"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Resolver string `json:"resolver"`
	Pfp      string `json:"pfp"`
}

func (sr ServiceResp) ToNames() []*NameResp {
	resp := []*NameResp{}
	err := helpers.JsonTransform(sr.Result, &resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToName() *NameResp {
	resp := &NameResp{}
	err := helpers.JsonTransform(sr.Result, resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToAvailable() *bool {
	resp := false
	err := helpers.JsonTransform(sr.Result, &resp)
	if err == nil {
		return &resp
	}

	return &resp
}
