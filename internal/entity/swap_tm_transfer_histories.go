package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapTmTransferHistories struct {
	BaseEntity      `bson:",inline"`
	TxHash          string               `json:"tx_hash" bson:"tx_hash,omitempty"`
	ContractAddress string               `json:"contract_address" bson:"contract_address,omitempty"`
	Timestamp       time.Time            `json:"timestamp" bson:"timestamp,omitempty"`
	Index           uint                 `json:"log_index" bson:"log_index,omitempty"`
	From            string               `json:"from" bson:"from,omitempty"`
	To              string               `json:"to" bson:"to,omitempty"`
	Value           primitive.Decimal128 `json:"value" bson:"value,omitempty"`
}

func (t *SwapTmTransferHistories) CollectionName() string {
	return utils.COLLECTION_SWAP_TOKEN_TRANSFER_HISTORY
}

type SwapTmTransferHistoriesFilter struct {
	BaseFilters
	TxHash      string
	UserAddress string
	Index       uint
}

func (t *SwapTmTransferHistoriesFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
