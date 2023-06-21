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
	Owner               *string
	Name                *string
	Address             *string
	AllowEmpty          *bool
	ContentTypeNotEmpty *bool
	IsBigFile           *bool
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
	PFP       *string
	Resolver  *string
	Owner     *string
	Name      *string
	TokenID   *string
}

func (pagination *PaginationReq) GetOffsetAndLimit() (int, int) {
	limit := 32
	offset := 0

	if pagination != nil {
		if pagination.Offset != nil {
			offset = *pagination.Offset
		}
		if pagination.Page != nil {
			offset = (int(*pagination.Page) - 1) * limit
		}
		if pagination.Limit != nil {
			limit = *pagination.Limit
		}
	}
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 || limit > 1000 {
		limit = 32
	}

	return limit, offset
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
