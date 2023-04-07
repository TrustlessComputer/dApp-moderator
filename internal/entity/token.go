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
	Symbol          string `json:"symbol" bson:"symbol"`
	Slug            string `json:"slug" bson:"slug"`
	Decimal         int    `json:"decimal" bson:"decimal"`
	Name            string `json:"name" bson:"name"`
	TotalSupply     string `json:"total_supply" bson:"total_supply"`
	Owner           string `json:"owner" bson:"owner"`
	DeployedAtBlock int    `json:"deployed_at_block" bson:"deployed_at_block"`
	Thumbnail       string `json:"thumbnail" bson:"thumbnail"`
	Description     string `json:"description" bson:"description"`
	Social          Social `json:"social" bson:"social"`
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
	Address   string
	Key       string
	CreatedBy string
}

func (t *TokenFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 10
	if pag.Limit != nil {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil {
		t.Page = int64(*pag.Page)
	}
}
