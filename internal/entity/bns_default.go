package entity

import (
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BNSDefault struct {
	BaseEntity `bson:",inline"`

	Resolver     string             `json:"resolver" bson:"resolver"`
	BNSDefaultID primitive.ObjectID `json:"bns_default_id" bson:"bns_default_id"`

	BNSDefaultData *Bns `json:"bns_default_data,omitempty" bson:"bns_default_data,omitempty"`
}

func (b *BNSDefault) CollectionName() string {
	return utils.COLLECTION_BNS_DEFAULT
}
