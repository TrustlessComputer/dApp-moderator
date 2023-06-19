package nft_explorer

import (
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/helpers"
)

type Erc721 struct {
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

type RequestData struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type ServiceResp struct {
	Code   string      `json:"code"`
	Error  error       `json:"error"`
	Result interface{} `json:"result"`
}

type CollectionsResp struct {
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Contract        string `json:"contract"`
	ContractType    string `json:"contract_type"`
	Creator         string `json:"creator"`
	Description     string `json:"description"`
	TotalItems      int    `json:"total_items"`
	TotalOwners     int    `json:"total_owners"`
	Cover           string `json:"cover"`
	Thumbnail       string `json:"thumbnail"`
	DeployedAtBlock int64  `json:"deployed_at_block" json:"deployed_at_block"`
}

type MkpNftsResp struct {
	ContractAddress  string                            `bson:"collection_address" json:"collection_address"`
	TokenID          string                            `bson:"token_id" json:"token_id"`
	ContentType      string                            `bson:"content_type" json:"content_type"`
	Name             string                            `bson:"name" json:"name"`
	Owner            string                            `bson:"owner" json:"owner"`
	TokenURI         string                            `bson:"token_uri" json:"token_uri"`
	Image            string                            `bson:"image" json:"image"`
	MintedAt         float64                           `bson:"minted_at" json:"minted_at"`
	Attributes       []MkpNftAttr                      `json:"attributes" bson:"attributes"`
	Metadata         interface{}                       `json:"metadata" bson:"metadata"`
	MetadataType     string                            `json:"metadata_type" bson:"metadata_type"`
	Activities       []entity.MarketplaceTokenActivity `json:"activities" bson:"activities"`
	BlockNumber      string                            `json:"block_number" bson:"block_number"`
	ListingForSales  []entity.MarketplaceListings      `json:"listing_for_sales" bson:"listing_for_sales"`
	MakeOffers       []entity.MarketplaceOffers        `json:"make_offers" bson:"make_offers"`
	Buyable          bool                              `bson:"buyable" json:"buyable"`
	PriceERC20       *MkpPriceERC20                    `bson:"price_erc20" json:"price_erc20"`
	Collection       entity.Collections                `json:"collection" bson:"collection"`
	Size             int64                             `json:"size" bson:"size"`
	BnsData          []*entity.FilteredBNS             `json:"bns_data,omitempty"`
	ImageCapture     string                            `bson:"image_capture" json:"image_capture"`           // capture thumbnail from html - animation_file_url
	AnimationFileUrl string                            `bson:"animation_file_url" json:"animation_file_url"` // capture thumbnail from html
}

type MkpNftsPagination struct {
	Items     []*MkpNftsResp `bson:"items" json:"items"`
	TotalItem int            `json:"total_item" bson:"total_item"`
}

type SoulNft struct {
	MkpNftsResp `bson:",inline"`
	IsAuction   bool `json:"is_auction" bson:"is_auction"`
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

type NftsResp struct {
	//Collection      string                            `json:"collection"`
	ContractAddress string                             `json:"collection_address"`
	TokenID         string                             `json:"token_id"`
	ContentType     string                             `json:"content_type"`
	Name            string                             `json:"name"`
	Owner           string                             `json:"owner"`
	TokenURI        string                             `json:"token_uri"`
	Image           string                             `json:"image"`
	MintedAt        float64                            `json:"minted_at"`
	Attributes      []NftAttr                          `json:"attributes"`
	Metadata        interface{}                        `json:"metadata"`
	MetadataType    string                             `json:"metadata_type"`
	Activities      []*entity.MarketplaceTokenActivity `json:"activities"`
	BlockNumber     string                             `json:"block_number"`
	ListingForSales []entity.MarketplaceListings       `json:"listing_for_sales"`
	MakeOffers      []entity.MarketplaceOffers         `json:"make_offers"`
	Collection      entity.Collections                 `json:"collection"`
}

type NftAttr struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

func (sr ServiceResp) ToCollections() []CollectionsResp {
	resp := []CollectionsResp{}
	err := helpers.JsonTransform(sr.Result, &resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToCollection() *CollectionsResp {
	resp := &CollectionsResp{}
	err := helpers.JsonTransform(sr.Result, resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToNfts() []*NftsResp {
	resp := []*NftsResp{}
	err := helpers.JsonTransform(sr.Result, &resp)
	if err == nil {
		return resp
	}

	return resp
}

func (sr ServiceResp) ToNft() *NftsResp {
	resp := &NftsResp{}
	err := helpers.JsonTransform(sr.Result, resp)
	if err == nil {
		return resp
	}

	return resp
}
