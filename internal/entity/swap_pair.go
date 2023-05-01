package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapPairReport struct {
	Address         string               `json:"address" bson:"address"`
	TotalSupply     string               `json:"total_supply" bson:"total_supply"`
	Owner           string               `json:"owner" bson:"owner"` // Owner of a contract (contract address)
	Decimal         int                  `json:"decimal" bson:"decimal"`
	DeployedAtBlock int                  `json:"deployed_at_block" bson:"deployed_at_block"`
	Slug            string               `json:"slug" bson:"slug"`
	Symbol          string               `json:"symbol" bson:"symbol"`
	Name            string               `json:"name" bson:"name"`
	Thumbnail       string               `json:"thumbnail" bson:"thumbnail"`
	Description     string               `json:"description" bson:"description"`
	Social          Social               `json:"social" bson:"social"`
	Index           int64                `json:"index" bson:"index"`
	Volume          primitive.Decimal128 `json:"volume" bson:"volume"`
	Price           primitive.Decimal128 `json:"price" bson:"price"`
	Percent         primitive.Decimal128 `json:"percent" bson:"percent"`
}

type SwapPair struct {
	BaseEntity      `bson:",inline"`
	TxHash          string    `json:"tx_hash"  bson:"tx_hash,omitempty"`
	ContractAddress string    `json:"contract_address"  bson:"contract_address,omitempty"`
	Timestamp       time.Time `json:"timestamp"  bson:"timestamp,omitempty"`
	Token0          string    `json:"token0"  bson:"token0,omitempty"`
	Token1          string    `json:"token1"  bson:"token1,omitempty"`
	Pair            string    `json:"pair"  bson:"pair,omitempty"`
	Arg3            int64     `json:"arg3"  bson:"arg3,omitempty"`
	Index           uint      `json:"log_index"  bson:"log_index,omitempty"`
}

func (t *SwapPair) CollectionName() string {
	return utils.COLLECTION_SWAP_PAIR
}

type SwapPairFilter struct {
	BaseFilters
	Pair   string
	TxHash string
	Token  string
}

func (t *SwapPairFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
