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
	Types           []int
}

type FilterNfts struct {
	BaseFilters
	ContractAddress *string
	TokenID         *string
	Owner           *string
	//Rarity          *string
	AttrKey   []string
	AttrValue []string
	Rarity    *Rarity
	Price     *Rarity
	IsBigFile *bool
	IsBuyable *bool
}

type Rarity struct {
	Min float64
	Max float64
}

type FilterMarketplaceAggregationData struct {
	BaseFilters
	CollectionContract *string
	Name               *string
}

type MkpNftAttr struct {
	TraitType string  `json:"trait_type" bson:"trait_type"`
	Value     string  `json:"value" bson:"value"`
	Count     int64   `json:"count" bson:"count"`
	Total     int64   `json:"total" bson:"total"`
	Percent   float64 `json:"percent" bson:"percent"`
}

type MkpPriceERC20 struct {
	OfferingID string `bson:"offering_id" json:"offering_id"`
	TokenID    string `bson:"token_id" json:"token_id"`
	Erc20Token string `bson:"erc_20_token" json:"erc_20_token"`
	Price      string `bson:"price" json:"price"`
}

type MkpNftsResp struct {
	ContractAddress  string                     `bson:"collection_address" json:"collection_address"`
	TokenID          string                     `bson:"token_id" json:"token_id"`
	ContentType      string                     `bson:"content_type" json:"content_type"`
	Name             string                     `bson:"name" json:"name"`
	Owner            string                     `bson:"owner" json:"owner"`
	TokenURI         string                     `bson:"token_uri" json:"token_uri"`
	Image            string                     `bson:"image" json:"image"`
	ImageCapture     string                     `bson:"image_capture" json:"image_capture"`           // capture thumbnail from html - animation_file_url
	AnimationFileUrl string                     `bson:"animation_file_url" json:"animation_file_url"` // capture thumbnail from html
	MintedAt         float64                    `bson:"minted_at" json:"minted_at"`
	Attributes       []MkpNftAttr               `json:"attributes" bson:"attributes"`
	Metadata         interface{}                `json:"metadata" bson:"metadata"`
	MetadataType     string                     `json:"metadata_type" bson:"metadata_type"`
	Activities       []MarketplaceTokenActivity `json:"activities" bson:"activities"`
	BlockNumber      string                     `json:"block_number" bson:"block_number"`
	ListingForSales  []MarketplaceListings      `json:"listing_for_sales" bson:"listing_for_sales"`
	MakeOffers       []MarketplaceOffers        `json:"make_offers" bson:"make_offers"`
	Buyable          bool                       `bson:"buyable" json:"buyable"`
	PriceERC20       *MkpPriceERC20             `bson:"price_erc20" json:"price_erc20"`
	Collection       Collections                `json:"collection" bson:"collection"`
	Size             int64                      `json:"size" bson:"size"`
	BnsData          []*Bns                     `json:"bns_data,omitempty" bson:"bns_data"`
	BnsDefault       []*BNSDefault              `json:"bns_default,omitempty" bson:"bns_default"`
}

type MkpNftsPagination struct {
	Items     []*MkpNftsResp `bson:"items" json:"items"`
	TotalItem int64          `json:"total_item" bson:"total_item"`
}
