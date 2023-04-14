package entity

import (
	"dapp-moderator/utils"
)

type NftHistories struct {
	BaseEntity        `bson:",inline"`
	Collection        string `json:"collection" bson:"collection"`
	ContractAddress   string `json:"collection_address" bson:"collection_address"`
	TokenID           string `json:"token_id" bson:"token_id"`
	TokenIDInt        int64  `json:"token_id_int" bson:"token_id_int"` //use it for sort
	FromWalletAddress string `json:"from_wallet_address" bson:"from_wallet_address"`
	ToWalletAddress   string `json:"to_wallet_address" bson:"to_wallet_address"`
	Action            string `json:"action" bson:"action"`
}

func (u NftHistories) CollectionName() string {
	return utils.COLLECTION_NFT_HISTORIES
}
