package entity

import (
	"dapp-moderator/external/token_explorer"
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
}

type TokenFilter struct {
	BaseFilters
	Address   string
	Name      string
	Symbol    string
	CreatedBy string
	Decimal   int
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
