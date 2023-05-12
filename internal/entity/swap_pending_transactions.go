package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"
)

type SwapPendingTransactions struct {
	BaseEntity `bson:",inline"`
	TxHash     string    `json:"tx_hash" bson:"tx_hash,omitempty"`
	Timestamp  time.Time `json:"timestamp"  bson:"timestamp,omitempty"`
}

func (t *SwapPendingTransactions) CollectionName() string {
	return utils.COLLECTION_SWAP_PENDING_TRANSACTION
}

type SwapPendingTransactionsFilter struct {
	BaseFilters
	TxHash []string
}

func (t *SwapPendingTransactionsFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
