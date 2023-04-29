package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapPairSwapHistories struct {
	BaseEntity      `bson:",inline"`
	TxHash          string               `json:"tx_hash" bson:"tx_hash"`
	ContractAddress string               `json:"contract_address"  bson:"contract_address"`
	Timestamp       time.Time            `json:"timestamp"  bson:"timestamp"`
	Sender          string               `json:"sender"  bson:"sender"`
	To              string               `json:"to"  bson:"to"`
	Amount0In       primitive.Decimal128 `json:"amount0_in"  bson:"amount0_in"`
	Amount1In       primitive.Decimal128 `json:"amount1_in"  bson:"amount1_in"`
	Amount0Out      primitive.Decimal128 `json:"amount0_out"  bson:"amount0_out"`
	Amount1Out      primitive.Decimal128 `json:"amount1_out"  bson:"amount1_out"`
	Index           uint                 `json:"log_index"  bson:"log_index"`
}

func (t *SwapPairSwapHistories) CollectionName() string {
	return utils.COLLECTION_SWAP_HISTORIES
}

type SwapPairSwapHistoriesFilter struct {
	BaseFilters
	ContractAddress string
	TxHash          string
}

func (t *SwapPairSwapHistoriesFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
