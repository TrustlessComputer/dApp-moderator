package entity

import "dapp-moderator/utils"

type MarketplaceStatus int

const (
	MarketPlaceOpen   MarketplaceStatus = 0
	MarketPlaceCancel MarketplaceStatus = 1
	MarketPlaceDone   MarketplaceStatus = 2
)

type MarketplaceListings struct {
	BaseEntity         `bson:",inline"`
	OfferingId         string            `bson:"offering_id" json:"offering_id"`
	CollectionContract string            `bson:"collection_contract" json:"collection_contract"`
	TokenId            string            `bson:"token_id" json:"token_id"`
	Seller             string            `bson:"seller" json:"seller"`
	Erc20Token         string            `bson:"erc_20_token" json:"erc_20_token"`
	Price              string            `bson:"price" json:"price"`
	Status             MarketplaceStatus `bson:"status" json:"status"`
	DurationTime       string            `bson:"duration_time" json:"duration_time"`
	BlockNumber        uint64            `bson:"block_number" json:"block_number"`
	OwnerAddress       *string           `bson:"owner_address" json:"owner_address"`
}

func (u MarketplaceListings) CollectionName() string {
	return utils.COLLECTION_MARKETPLACE_LISTING
}

type FilterMarketplaceListings struct {
	BaseFilters
	CollectionContract *string
	TokenId            *string
	Erc20Token         *string
	SellerAddress      *string
	Status             *int
}

type FilterMarketplaceOffer struct {
	BaseFilters
	CollectionContract *string
	TokenId            *string
	Erc20Token         *string
	BuyerAddress       *string
	Status             *int
}

type FilterTokenActivities struct {
	BaseFilters
	ContractAddress *string
	TokenID         *string
}

type FilterNfts struct {
	BaseFilters
	ContractAddress *string
	TokenID         *string
}

type FilterMarketplaceAggregationData struct {
	BaseFilters
	CollectionContract *string
	Name               *string
}
