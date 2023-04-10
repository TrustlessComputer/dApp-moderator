package entity

import (
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"encoding/json"
	"strings"
)

type Token struct {
	BaseEntity      `bson:",inline"`
	Address         string `json:"address" bson:"address"`
	TotalSupply     string `json:"total_supply" bson:"total_supply"`
	Owner           string `json:"owner" bson:"owner"`
	Decimal         int    `json:"decimal" bson:"decimal"`
	DeployedAtBlock int    `json:"deployed_at_block" bson:"deployed_at_block"`
	Slug            string `json:"slug" bson:"slug"`

	// edit able
	Symbol      string `json:"symbol" bson:"symbol"`
	Name        string `json:"name" bson:"name"`
	Thumbnail   string `json:"thumbnail" bson:"thumbnail"`
	Description string `json:"description" bson:"description"`
	Social      Social `json:"social" bson:"social"`
}

func (t *Token) CollectionName() string {
	return utils.COLLECTION_TOKENS
}

func (t *Token) FromTokenExplorer(te token_explorer.Token) error {
	data, err := json.Marshal(te)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	t.Slug = strings.ToLower(t.Symbol)

	return err
}

type TokenFilter struct {
	BaseFilters
	Key       string
	Address   string
	CreatedBy string
}

func (t *TokenFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
