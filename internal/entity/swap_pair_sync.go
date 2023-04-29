package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"math/big"
	"time"
)

type SwapPairSync struct {
	BaseEntity      `bson:",inline"`
	TxHash          string    `json:"tx_hash" bson:"tx_hash"`
	ContractAddress string    `json:"contract_address" bson:"contract_address"`
	Timestamp       time.Time `json:"timestamp" bson:"timestamp"`
	Reserve0        *big.Int  `json:"reserve0" bson:"reserve0"`
	Reserve1        *big.Int  `json:"reserve1" bson:"reserve1"`
	Index           uint      `json:"log_index" bson:"log_index"`
}

func (t *SwapPairSync) CollectionName() string {
	return utils.COLLECTION_SWAP_PAIR_SYNC
}

type SwapPairSyncFilter struct {
	BaseFilters
	ContractAddress string
	TxHash          string
}

func (t *SwapPairSyncFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
