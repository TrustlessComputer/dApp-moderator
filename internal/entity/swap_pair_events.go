package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapPairEventsType string

const SwapPairEventsTypeMint SwapPairEventsType = "mint"
const SwapPairEventsTypeBurn SwapPairEventsType = "burn"

type SwapPairEvents struct {
	BaseEntity      `bson:",inline"`
	TxHash          string               `json:"tx_hash" bson:"tx_hash,omitempty"`
	ContractAddress string               `json:"contract_address"  bson:"contract_address,omitempty"`
	Timestamp       time.Time            `json:"timestamp"  bson:"timestamp,omitempty"`
	Sender          string               `json:"sender"  bson:"sender,omitempty"`
	To              string               `json:"to"  bson:"to,omitempty"`
	EventType       SwapPairEventsType   `json:"event_type"  bson:"event_type,omitempty"`
	Amount0         primitive.Decimal128 `json:"amount0"  bson:"amount0,omitempty"`
	Amount1         primitive.Decimal128 `json:"amount1"  bson:"amount1,omitempty"`
	Index           uint                 `json:"log_index"  bson:"log_index,omitempty"`
	Pair            *SwapPair            `json:"pair" bson:"pair,omitempty"`
}

func (t *SwapPairEvents) CollectionName() string {
	return utils.COLLECTION_SWAP_PAIR_EVENTS
}

type SwapPairEventFilter struct {
	BaseFilters
	ContractAddress string
	TxHash          string
}

func (t *SwapPairEventFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
