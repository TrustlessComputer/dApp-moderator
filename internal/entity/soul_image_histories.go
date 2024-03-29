package entity

import (
	"dapp-moderator/utils"
	"time"
)

type SoulTokenHistoriesFiltered struct {
	BaseEntity   `bson:",inline" json:"-"`
	Type         TokenActivityType `bson:"type" json:"type"`
	Title        string            `bson:"title" json:"-"`
	Owner        string            `bson:"owner" json:"owner"`
	AmountStr    float64           `bson:"amount" json:"amount"`
	Erc20Address string            `bson:"erc_20_address" json:"erc_20_address"`
	Time         *time.Time        `bson:"time" json:"time"`
	//InscriptionID      string            `bson:"inscription_id" json:"inscription_id"`
	TokenID string `bson:"token_id" json:"token_id"`
	//Thumbnail          string `bson:"thumbnail" json:"thumbnail"`
	CollectionContract string `bson:"collection_contract" json:"collection_contract"`
	BlockNumber        uint64 `bson:"block_number" json:"block_number"`
	MintedAt           uint64 `bson:"minted_at" json:"minted_at"`
	HoldTime           uint64 `bson:"hold_time" json:"hold_time"`
	TxHash             string `json:"tx_hash" bson:"tx_hash"`
	LogIndex           uint   `json:"log_index" bson:"log_index"`
	ImageCapture       string `json:"image_capture" bson:"image_capture"`
	FeatureName        string `json:"feature_name" bson:"feature_name"`
	Balance            string `json:"balance" bson:"balance"`
}

type SoulImageHistories struct {
	BaseEntity `bson:",inline"`

	ContractAddress string `json:"collection_address" bson:"collection_address"`
	TokenID         string `json:"token_id" bson:"token_id"`
	TokenIDInt      int64  `json:"token_id_int" bson:"token_id_int"` //use it for sort

	ImageCaptureAt      *time.Time        `json:"image_capture_at" bson:"image_capture_at"`
	ImageCaptureDate    string            `json:"image_capture_date" bson:"image_capture_date"`
	ImageCapture        string            `json:"image_capture,omitempty" bson:"image_capture,omitempty"`
	Owner               string            `json:"owner" bson:"owner"`
	TxHash              string            `bson:"tx_hash" json:"tx_hash"`
	LogIndex            int               `bson:"log_index" json:"log_index"`
	BlockNumber         uint64            `json:"block_number" bson:"block_number"`
	Erc20Amount         string            `json:"erc_20_amount" bson:"erc_20_amount"`
	Erc20Address        string            `json:"erc_20_address" bson:"erc_20_address"`
	BitcoinDexWETHPrice string            `json:"bitcoin_dex_weth_price" bson:"bitcoin_dex_weth_price"`
	BitcoinDexWBTCPrice string            `json:"bitcoin_dex_wbtc_price" bson:"bitcoin_dex_wbtc_price"`
	BitcoinDexUSDTPrice string            `json:"bitcoin_dex_usdt_price" bson:"bitcoin_dex_usdt_price"`
	Event               TokenActivityType `bson:"event" bson:"event"`
	FeatureName         string            `bson:"feature_name" bson:"feature_name"`
}

func (u SoulImageHistories) CollectionName() string {
	return utils.COLLECTION_SOUL_IMAGE_HISTORIES
}
