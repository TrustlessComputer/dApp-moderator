package entity

import "dapp-moderator/utils"

type MarketplaceCollectionAttribute struct {
	Count     int64   `bson:"count" json:"count"`
	Contract  string  `bson:"contract" json:"contract"`
	TraitType string  `bson:"trait_type" json:"trait_type"`
	Value     string  `bson:"value" json:"value"`
	Total     int64   `bson:"total" json:"total"`
	Percent   float64 `bson:"percent" json:"percent"`
}

func (u MarketplaceCollectionAttribute) CollectionName() string {
	return utils.VIEW_MARKETPLACE_COLLECTION_ATTRIBUTES_PERCENT
}

type FilterMarketplaceCollectionAttribute struct {
	BaseFilters
	ContractAddress *string
	TraitType       *string
	Value           *string
	Percent         *float64
}
