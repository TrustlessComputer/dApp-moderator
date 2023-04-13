package request

import (
	"fmt"
	"net/url"
)

type PaginationReq struct {
	Limit  *int
	Page   *int
	Offset *int
	SortBy *string
	Sort   *int
}

type CollectionsFilter struct {
	PaginationReq
	Owner      *string
	Name       *string
	Address    *string
	AllowEmpty *bool
}

type HistoriesFilter struct {
	PaginationReq
	WalletAdress *string
	TxHash       *string
}

type ConfirmHistoriesReq struct {
	Data []struct {
		TxHash  []string `json:"tx_hash"`
		BTCHash string   `json:"btc_hash"`
		Status  string   `json:"status"`
	} `json:"data"`
}

type NftItemsFilter struct {
	PaginationReq
	Owner      *string
	Name       *string
	Address    *string
	AllowEmpty *bool
}

type FilterBNSNames struct {
	PaginationReq
	FromBlock *int
	ToBlock   *int
}

func (pq PaginationReq) ToNFTServiceUrlQuery() url.Values {
	q := url.Values{}

	if pq.Limit != nil && *pq.Limit != 0 {
		q.Set("limit", fmt.Sprintf("%d", *pq.Limit))
	}

	if pq.Offset != nil {
		q.Set("offset", fmt.Sprintf("%d", *pq.Offset))
	}

	return q
}
