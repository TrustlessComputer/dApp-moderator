package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapBotTransaction struct {
	BaseEntity   `bson:",inline"`
	Address      string `json:"address" bson:"address,omitempty"`
	SwapPair     `json:"pair" bson:"pair,omitempty"`
	TxHash       string               `json:"tx_hash" bson:"tx_hash,omitempty"`
	AmountIn     primitive.Decimal128 `json:"amount_in" bson:"amount_in,omitempty"`
	AmountOutMin primitive.Decimal128 `json:"amount_out_min" bson:"amount_out_min,omitempty"`
	Status       int                  `json:"status" bson:"status,omitempty"`
}

func (t *SwapBotTransaction) CollectionName() string {
	return utils.COLLECTION_SWAP_BOT_TRANSACTION
}

type SwapBotTransactionFilter struct {
	BaseFilters
	Address string
	Status  int
}

func (t *SwapBotTransactionFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
