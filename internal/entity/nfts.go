package entity

import (
	"dapp-moderator/utils"
)

type Files struct {
	BaseEntity      `bson:",inline"`
	Name            string `bson:"name" json:"name"`
	Slug            string `bson:"slug" json:"slug"`
	Contract        string `bson:"contract" json:"contract"`
	Creator         string `bson:"creator" json:"creator"`
	Description     string `bson:"description" json:"description"`
	TotalItems      int    `bson:"total_items" json:"total_items"`
	TotalOwners     int    `bson:"total_owners" json:"total_owners"`
	Cover           string `bson:"cover" json:"cover"`
	Thumbnail       string `bson:"thumbnail" json:"thumbnail"`
	DeployedAtBlock int64  `bson:"deployed_at_block" json:"deployed_at_block"`
}

type FilterFiles struct {
	BaseFilters
	Name       *string
	UploadedBy *string
}

func (u Files) CollectionName() string {
	return utils.COLLECTION_NFTS
}
