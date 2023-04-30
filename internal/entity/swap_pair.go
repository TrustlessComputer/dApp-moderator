package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"
)

type SwapPair struct {
	BaseEntity      `bson:",inline"`
	TxHash          string    `json:"tx_hash"  bson:"tx_hash,omitempty"`
	ContractAddress string    `json:"contract_address"  bson:"contract_address,omitempty"`
	Timestamp       time.Time `json:"timestamp"  bson:"timestamp,omitempty"`
	Token0          string    `json:"token0"  bson:"token0,omitempty"`
	Token1          string    `json:"token1"  bson:"token1,omitempty"`
	Pair            string    `json:"pair"  bson:"pair,omitempty"`
	Arg3            int64     `json:"arg3"  bson:"arg3,omitempty"`
	Index           uint      `json:"log_index"  bson:"log_index,omitempty"`
}

func (t *SwapPair) CollectionName() string {
	return utils.COLLECTION_SWAP_PAIR
}

type SwapPairFilter struct {
	BaseFilters
	Pair   string
	TxHash string
	Token  string
}

func (t *SwapPairFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
