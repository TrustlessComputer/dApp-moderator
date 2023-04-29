package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"math/big"
	"time"
)

type SwapPairEventsType string

const SwapPairEventsTypeMint SwapPairEventsType = "mint"
const SwapPairEventsTypeBurn SwapPairEventsType = "burn"

type SwapPairEvents struct {
	BaseEntity      `bson:",inline"`
	TxHash          string             `json:"tx_hash" bson:"tx_hash"`
	ContractAddress string             `json:"contract_address"  bson:"contract_address"`
	Timestamp       time.Time          `json:"timestamp"  bson:"timestamp"`
	Sender          string             `json:"sender"  bson:"sender"`
	To              string             `json:"to"  bson:"to"`
	EventType       SwapPairEventsType `json:"event_type"  bson:"event_type"`
	Amount0         *big.Int           `json:"amount0"  bson:"amount0"`
	Amount1         *big.Int           `json:"amount1"  bson:"amount1"`
	Index           uint               `json:"log_index"  bson:"tx_hash"`
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
