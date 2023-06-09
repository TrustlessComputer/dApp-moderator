package entity

import (
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NftAuctions struct {
	ID        primitive.ObjectID `json:"-" bson:"-"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt *time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at"`

	ContractAddress string `json:"collection_address" bson:"collection_address"`
	TokenID         string `json:"token_id" bson:"token_id"`
	TokenIDInt      int64  `json:"token_id_int" bson:"token_id_int"` //use it for sort
	IsAuction       bool   `bson:"is_auction" json:"is_auction"`
}

func (u NftAuctions) CollectionName() string {
	return utils.COLLECTION_NFT_AUCTIONS
}
