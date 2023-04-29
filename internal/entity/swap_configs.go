package entity

import (
	"dapp-moderator/utils"
)

type SwapConfigs struct {
	BaseEntity `bson:",inline"`
	Name       string `json:"name"  bson:"name,omitempty"`
	Value      string `json:"value"  bson:"value,omitempty"`
}

func (t *SwapConfigs) CollectionName() string {
	return utils.COLLECTION_SWAP_CONFIGS
}

type SwapConfigsFilter struct {
	BaseFilters
	Name  string
	Value string
}
