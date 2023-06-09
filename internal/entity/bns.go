package entity

import "dapp-moderator/utils"

type Bns struct {
	BaseEntity        `bson:",inline"`
	Name              string `json:"name" bson:"name"`
	TokenID           string `json:"token_id" bson:"token_id"`
	CollectionAddress string `json:"collection_address" bson:"collection_address"`
	Owner             string `json:"owner" bson:"owner"`
	Resolver          string `json:"resolver" bson:"resolver"`
	Pfp               string `json:"pfp" bson:"pfp"`
}

func (u Bns) CollectionName() string {
	return utils.COLLECTION_BNS
}
