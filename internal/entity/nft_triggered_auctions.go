package entity

import (
	"dapp-moderator/utils"
)

type NftTriggeredAuctions struct {
	BaseEntity `bson:",inline"`

	ContractAddress string `json:"collection_address" bson:"collection_address"`
	TokenID         string `json:"token_id" bson:"token_id"`
	TokenIDInt      int64  `json:"token_id_int" bson:"token_id_int"` //use it for sort
	IsAuction       bool   `bson:"is_auction" json:"is_auction"`
	TxHash          string `json:"tx_hash" bson:"tx_hash"`
}

func (u NftTriggeredAuctions) CollectionName() string {
	return utils.COLLECTION_NFT_TRIGGERED_AUCTIONS
}
