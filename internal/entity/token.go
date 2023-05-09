package entity

import (
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"encoding/json"
	"strings"
)

type OwnedToken struct {
	Token
	Balance string `json:"balance" bson:"balance"`
	Decimal int    `json:"decimal" bson:"decimal"`
}

type Token struct {
	BaseEntity      `bson:",inline"`
	Address         string `json:"address" bson:"address"`
	TotalSupply     string `json:"total_supply" bson:"total_supply"`
	Owner           string `json:"owner" bson:"owner"` // Owner of a contract (contract address)
	Decimal         int    `json:"decimal" bson:"decimal"`
	DeployedAtBlock int    `json:"deployed_at_block" bson:"deployed_at_block"`
	Slug            string `json:"slug" bson:"slug"`

	// edit able
	Symbol      string `json:"symbol" bson:"symbol"`
	Name        string `json:"name" bson:"name"`
	Thumbnail   string `json:"thumbnail" bson:"thumbnail"`
	Description string `json:"description" bson:"description"`
	Social      Social `json:"social" bson:"social"`
	Index       int64  `json:"index" bson:"index"`
	Network     string `json:"network" bson:"network"`
	Priority    int64  `json:"priority" bson:"priority"`
}

func (t *Token) CollectionName() string {
	return utils.COLLECTION_TOKENS
}

func (t *Token) OwnedToken() *OwnedToken {
	resp := &OwnedToken{
		Token:   *t,
		Balance: "0", Decimal: 1}
	return resp
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
	t.Limit = 500
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
