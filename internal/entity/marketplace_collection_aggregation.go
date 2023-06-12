package entity

import (
	"dapp-moderator/utils"
)

type MarketplaceCollectionAggregation struct {
	BaseEntity                   `bson:",inline" json:"base_entity"`
	Contract                     string               `bson:"contract" json:"contract"`
	MarketPlaceVolumes           []*MarketPlaceVolume `bson:"market_place_volumes" json:"market_place_volumes"`
	FloorPriceMarketPlaceVolumes []*MarketPlaceVolume `bson:"floor_price_market_place_volumes" json:"floor_price_market_place_volumes"`
	UniqueOwners                 int64                `json:"unique_owners" bson:"unique_owners"`
	TotalNfts                    int64                `json:"total_nfts" bson:"total_nfts"`
	TotalSales                   int64                `json:"total_sales" bson:"total_sales"`

	//USDT
	FloorPrice float64 `json:"floor_price" bson:"floor_price"` //USDT
	Volume     float64 `json:"volume" bson:"volume"`           //USDT

	BtcFloorPrice float64 `json:"btc_floor_price" bson:"btc_floor_price"` //USDT
	EthFloorPrice float64 `json:"eth_floor_price" bson:"eth_floor_price"` //USDT
	BtcVolume     float64 `json:"btc_volume" bson:"btc_volume"`           //USDT
	EthVolume     float64 `json:"eth_volume" bson:"eth_volume"`           //USDT
}

type MarketPlaceVolume struct {
	TotalVolume     float64 `bson:"total_volume" json:"total_volume"`
	TotalSales      int64   `bson:"total_sales" json:"total_sales"`
	Erc20Token      string  `bson:"erc_20_token" json:"erc_20_token"`
	Contract        string  `bson:"contract" json:"contract"`
	MarketplaceType string  `bson:"marketplace_type" json:"marketplace_type"`
	Erc20Rate       float64 `bson:"erc_20_rate" json:"erc_20_rate"`
	Erc20Decimal    int     `bson:"erc_20_decimal" json:"erc_20_decimal"`
	USDTValue       float64 `bson:"usdt_value" json:"usdt_value"`
	WBTCRate        float64 `bson:"wbtc_rate" json:"wbtc_rate"`
	WEthRate        float64 `bson:"weth_rate" json:"weth_rate"`
	BTCValue        float64 `bson:"btc_value" json:"btc_value"`
	EthValue        float64 `bson:"eth_value" json:"eth_value"`
}

type MarketplaceCollections struct {
	BaseEntity  `bson:",inline"`
	Collections `bson:",inline"`
	TotalNfts   int64 `json:"-" bson:"total_nfts"`
	TotalSales  int64 `json:"total_sales" bson:"total_sales"`
	UniqueOwner int64 `json:"unique_owners" bson:"unique_owners"`

	//USDT
	FloorPrice float64 `json:"floor_price" bson:"floor_price"` //USDT
	Volume     float64 `json:"volume" bson:"volume"`           //USDT
}

func (u MarketplaceCollectionAggregation) CollectionName() string {
	return utils.COLLECTION_MARKETPLACE_AGGREGATED_COLLECTIONS
}
