package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapBotConfig struct {
	BaseEntity    `bson:",inline"`
	Address       string `json:"address" bson:"address,omitempty"`
	SwapPair      `json:"pair" bson:"pair,omitempty"`
	Enabled       bool                 `json:"enabled" bson:"enabled,omitempty"`
	BeginPrice    primitive.Decimal128 `json:"begin_price" bson:"begin_price,omitempty"`
	BeginReserve0 primitive.Decimal128 `json:"begin_reserve0" bson:"begin_reserve0,omitempty"`
	BeginReserve1 primitive.Decimal128 `json:"begin_reserve1" bson:"begin_reserve1,omitempty"`
	CurrentDate   string               `json:"current_date" bson:"current_date,omitempty"`
	MinValue      primitive.Decimal128 `json:"min_value" bson:"min_value,omitempty"`
	MaxValue      primitive.Decimal128 `json:"max_value" bson:"max_value,omitempty"`
}

func (t *SwapBotConfig) CollectionName() string {
	return utils.COLLECTION_SWAP_BOT_CONFIG
}

type SwapBotConfigFilter struct {
	BaseFilters
	Enabled bool
	Address string
}

func (t *SwapBotConfigFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
