package entity

import (
	"dapp-moderator/utils"
	"time"
)

type FilterCollectionChart struct {
	ContractAddress *string
	FromDate        *time.Time
	ToDate          *time.Time
}

type FilterCollectionNftOwners struct {
	ContractAddress *string
	BaseFilters
}

type CollectionChart struct {
	Contract            string     `bson:"contract" json:"contract"`
	OfferingID          string     `bson:"offering_id" json:"-"`
	Erc20Token          string     `bson:"erc_20_token" json:"-"`
	Price               string     `bson:"price" json:"-"`
	VolumeCreatedAT     *time.Time `bson:"volume_created_at" json:"volume_created_at"`
	VolumeType          string     `bson:"volume_type" json:"volume_type"`
	VolumeCreatedAtDate string     `bson:"volume_created_at_date" json:"volume_created_at_date"`
	USDT                float64    `bson:"-" json:"usdt"`
	BTC                 string     `bson:"-" json:"btc"`
	USDTRate            float64    `bson:"-" json:"-"`
}

func (u CollectionChart) CollectionName() string {
	return utils.VIEW_MARKETPLACE_COLLECTION_CHART
}
