package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
)

type SwapBlackListToken struct {
	BaseEntity `bson:",inline"`
	Address    string `json:"address" bson:"address"`
}

func (t *SwapBlackListToken) CollectionName() string {
	return utils.COLLECTION_SWAP_BLACKLIST_TOKENS
}

type SwapBlackListokenFilter struct {
	BaseFilters
	Address string
}

func (t *SwapBlackListokenFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 500
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
