package entity

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapPairReport struct {
	Address         string               `json:"address" bson:"address"`
	TotalSupply     string               `json:"total_supply" bson:"total_supply"`
	Owner           string               `json:"owner" bson:"owner"` // Owner of a contract (contract address)
	Decimal         int                  `json:"decimal" bson:"decimal"`
	DeployedAtBlock int                  `json:"deployed_at_block" bson:"deployed_at_block"`
	Slug            string               `json:"slug" bson:"slug"`
	Symbol          string               `json:"symbol" bson:"symbol"`
	Name            string               `json:"name" bson:"name"`
	Thumbnail       string               `json:"thumbnail" bson:"thumbnail"`
	Description     string               `json:"description" bson:"description"`
	Social          Social               `json:"social" bson:"social"`
	Index           int64                `json:"index" bson:"index"`
	Volume          primitive.Decimal128 `json:"volume" bson:"volume"`
	TotalVolume     primitive.Decimal128 `json:"total_volume" bson:"total_volume"`
	BtcVolume       float64              `json:"btc_volume" bson:"btc_volume"`
	UsdVolume       float64              `json:"usd_volume" bson:"usd_volume"`
	BtcTotalVolume  float64              `json:"btc_total_volume" bson:"btc_total_volume"`
	UsdTotalVolume  float64              `json:"usd_total_volume" bson:"usd_total_volume"`
	Price           primitive.Decimal128 `json:"price" bson:"price"`
	BtcPrice        float64              `json:"btc_price" bson:"btc_price"`
	UsdPrice        float64              `json:"usd_price" bson:"usd_price"`
	Percent         primitive.Decimal128 `json:"percent" bson:"percent"`
	Percent7Day     primitive.Decimal128 `json:"percent_7day" bson:"percent_7day"`
}

type SwapPairSlackReport struct {
	Volume24h   primitive.Decimal128 `json:"volume_24h" bson:"volume_24h"`
	VolumeTotal primitive.Decimal128 `json:"volume_total" bson:"volume_total"`
	Tx24h       int32                `json:"tx_24h" bson:"tx_24h"`
	TxTotal     int32                `json:"tx_total" bson:"tx_total"`
	Users24h    int32                `json:"users_24h" bson:"users_24h"`
	UsersTotal  int32                `json:"users_total" bson:"users_total"`
}

type SwapPair struct {
	BaseEntity      `bson:",inline"`
	TxHash          string    `json:"tx_hash"  bson:"tx_hash,omitempty"`
	ContractAddress string    `json:"contract_address"  bson:"contract_address,omitempty"`
	Timestamp       time.Time `json:"timestamp"  bson:"timestamp,omitempty"`
	Token0          string    `json:"token0"  bson:"token0,omitempty"`
	Token1          string    `json:"token1"  bson:"token1,omitempty"`
	Pair            string    `json:"pair"  bson:"pair,omitempty"`
	Arg3            int64     `json:"arg3"  bson:"arg3,omitempty"`
	Index           uint      `json:"log_index"  bson:"log_index,omitempty"`
}

type ChartDataResp struct {
	Time             time.Time            `json:"time"`
	Timestamp        int64                `json:"timestamp"`
	VolumeFrom       primitive.Decimal128 `json:"volume_from"`
	VolumeTo         primitive.Decimal128 `json:"volume_to"`
	
	BtcPrice         float64              `json:"btc_price"`
	UsdPrice         float64              `json:"usd_price"`

	Low              primitive.Decimal128 `json:"low"`
	Open             primitive.Decimal128 `json:"open"`
	Close            primitive.Decimal128 `json:"close"`
	High             primitive.Decimal128 `json:"high"`

	VolumeFromUsd      float64 `json:"volume_from_usd"`
	VolumeToUsd        float64 `json:"volume_to_usd"`
	TotalVolumeUsd      float64              `json:"total_volume_usd"`
	TotalVolume     float64              `json:"total_volume"`
	LowUsd            float64 `json:"low_usd"`
	OpenUsd         float64 `json:"open_usd"`
	CloseUsd         float64 `json:"close_usd"`
	HighUsd            float64 `json:"high_usd"`


	ConversionType   string               `json:"conversion_type"`
	ConversionSymbol string               `json:"conversion_symbol"`
}

func (t *SwapPair) CollectionName() string {
	return utils.COLLECTION_SWAP_PAIR
}

type SwapPairFilter struct {
	BaseFilters
	Pair   string
	TxHash string
	Token  string
}
type ChartrFilter struct {
	BaseFilters
	contract_address string
}

func (t *SwapPairFilter) FromPagination(pag request.PaginationReq) {
	t.Limit = 100
	if pag.Limit != nil && *pag.Limit > 0 {
		t.Limit = int64(*pag.Limit)
	}

	t.Page = 1
	if pag.Page != nil && *pag.Page > 0 {
		t.Page = int64(*pag.Page)
	}
}
