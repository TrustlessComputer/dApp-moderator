package entity

import (
	"dapp-moderator/utils"
)

type SoulImages struct {
	BaseEntity `bson:",inline"`

	ContractAddress    string             `json:"collection_address" bson:"collection_address"`
	TokenID            string             `json:"token_id" bson:"token_id"`
	TokenIDInt         int64              `json:"token_id_int" bson:"token_id_int"` //use it for sort
	Image              *string            `json:"image" bson:"image"`
	AnimationURL       *string            `json:"animation_url" bson:"animation_url"`
	ReplacedAttributes *map[string]string `json:"replaced_attributes" bson:"replaced_attributes"`
}

func (u SoulImages) CollectionName() string {
	return utils.COLLECTION_SOUL_IMAGES
}
