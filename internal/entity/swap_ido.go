package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"
)

type SwapIdo struct {
	BaseEntity        `bson:",inline"`
	Token             `json:"token" bson:"token,omitempty"`
	UserWalletAddress string    `json:"user_wallet_address" bson:"user_wallet_address,omitempty"`
	StartAt           time.Time `json:"start_at"  bson:"start_at,omitempty"`
	Price             string    `json:"price" bson:"price,omitempty"`
	Link              string    `json:"link" bson:"link,omitempty"`
	Website           string    `json:"website" bson:"website,omitempty"`
	Twitter           string    `json:"twitter" bson:"twitter,omitempty"`
	WhitePaper        string    `json:"white_papper" bson:"white_papper,omitempty"`
	Discord           string    `json:"discord" bson:"discord,omitempty"`
}

func (t *SwapIdo) CollectionName() string {
	return utils.COLLECTION_SWAP_IDO
}

type SwapIdoFilter struct {
	BaseFilters
	ID             string
	Address        string
	WalletAddress  string
	CheckStartTime bool
}

func (t *SwapIdoFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}

type IdoTokenFilter struct {
	BaseFilters
	CreatedBy string
	Address   []string
}

func (t *IdoTokenFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
