package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapPairSync struct {
	BaseEntity      `bson:",inline"`
	TxHash          string               `json:"tx_hash" bson:"tx_hash,omitempty"`
	ContractAddress string               `json:"contract_address" bson:"contract_address,omitempty"`
	Timestamp       time.Time            `json:"timestamp" bson:"timestamp,omitempty"`
	Reserve0        primitive.Decimal128 `json:"reserve0" bson:"reserve0,omitempty"`
	Reserve1        primitive.Decimal128 `json:"reserve1" bson:"reserve1,omitempty"`
	Index           uint                 `json:"log_index" bson:"log_index,omitempty"`
	Token           string               `json:"token" bson:"token,omitempty"`
	Price           primitive.Decimal128 `json:"price" bson:"price,omitempty"`
	Pair            *SwapPair            `json:"pair" bson:"pair,omitempty"`
	BaseTokenSymbol string               `json:"base_token_symbol"  bson:"base_token_symbol,omitempty"`
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
