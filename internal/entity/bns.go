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

type FilteredBNS struct {
	BaseEntity        `bson:",inline"`
	TokenID           string `json:"token_id" bson:"token_id"`
	TokenIDInt        int64  `json:"token_id_int" bson:"token_id_int"` //use it for sort
	Name              string `json:"name" bson:"name"`
	Owner             string `json:"owner" bson:"owner"`
	CollectionAddress string `json:"collection_address" bson:"collection_address"`
	Resolver          string `json:"resolver" bson:"resolver"`
	Pfp               string `json:"pfp" bson:"pfp"`
}

type FilterBns struct {
	BaseFilters
	PFP      *string
	Resolver *string
	Owner    *string
	Name     *string
	TokenID  *string
}

func (u Bns) CollectionName() string {
	return utils.COLLECTION_BNS
}
