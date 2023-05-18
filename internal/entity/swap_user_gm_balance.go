package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapUserGmClaimSignature struct {
	Amount    string `json:"amount" bson:"amount"`
	Signature string `json:"signature" bson:"signature"`
}

type SwapUserGmBalance struct {
	BaseEntity  `bson:",inline"`
	UserAddress string               `json:"user_address" bson:"user_address,omitempty"`
	Balance     primitive.Decimal128 `json:"balance" bson:"balance"`
	IsContract  bool                 `json:"is_contract" bson:"is_contract"`
	BalanceSign string               `json:"balance_sign" bson:"balance_sign"`
	Signature   string               `json:"signature" bson:"signature"`
}

func (t *SwapUserGmBalance) CollectionName() string {
	return utils.COLLECTION_SWAP_USER_GM_BALANCE
}

type SwapUserGmBalanceFilter struct {
	BaseFilters
	Address string
}

func (t *SwapUserGmBalanceFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
