package entity

import (
	"dapp-moderator/utils"
)

type SoulAnimationHtml struct {
	BaseEntity `bson:",inline"`

	ContractAddress string  `json:"collection_address" bson:"collection_address"`
	TokenID         string  `json:"token_id" bson:"token_id"`
	TokenIDInt      int64   `json:"token_id_int" bson:"token_id_int"` //use it for sort
	AnimationURL    *string `json:"animation_url" bson:"animation_url"`
}

func (u SoulAnimationHtml) CollectionName() string {
	return utils.COLLECTION_SOUL_ANIMATION_HTML
}
