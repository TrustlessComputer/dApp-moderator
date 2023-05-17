package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapUserGmPaid struct {
	BaseEntity      `bson:",inline"`
	TxHash          string               `json:"tx_hash" bson:"tx_hash,omitempty"`
	ContractAddress string               `json:"contract_address"  bson:"contract_address,omitempty"`
	Timestamp       time.Time            `json:"timestamp"  bson:"timestamp,omitempty"`
	UserAddress     string               `json:"user_address" bson:"user_address,omitempty"`
	Amount          primitive.Decimal128 `json:"amount" bson:"amount"`
}

func (t *SwapUserGmPaid) CollectionName() string {
	return utils.COLLECTION_SWAP_USER_GM_PAID
}

type SwapUserGmPaidFilter struct {
	BaseFilters
	Address string
	TxHash  string
}

func (t *SwapUserGmPaidFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
