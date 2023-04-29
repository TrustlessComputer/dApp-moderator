package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"math/big"
	"time"
)

type SwapPair struct {
	BaseEntity      `bson:",inline"`
	TxHash          string    `json:"tx_hash"  bson:"tx_hash"`
	ContractAddress string    `json:"contract_address"  bson:"contract_address"`
	Timestamp       time.Time `json:"timestamp"  bson:"timestamp"`
	Token0          string    `json:"token0"  bson:"token0"`
	Token1          string    `json:"token1"  bson:"token1"`
	Pair            string    `json:"pair"  bson:"pair"`
	Arg3            *big.Int  `json:"arg3"  bson:"arg3"`
	Index           uint      `json:"log_index"  bson:"log_index"`
}

func (t *SwapPair) CollectionName() string {
	return utils.COLLECTION_SWAP_PAIR
}

type SwapPairFilter struct {
	BaseFilters
	Pair   string
	TxHash string
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
