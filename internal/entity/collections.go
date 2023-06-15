package entity

import (
	"dapp-moderator/utils"
)

type GroupedCollection struct {
	ID     GroupedCollectionID `bson:"_id"`
	Tokens int64               `bson:"tokens"`
}

type GroupedCollectionID struct {
	CollectionAddress string `bson:"collection_address"`
}

type CollectionNftOwner struct {
	Address           string `json:"address" bson:"address"`
	CollectionAddress string `json:"-" bson:"collection_address"`
	Name              string `json:"name" bson:"name"`
	Avatar            string `json:"avatar" bson:"avatar"`
	Count             int64  `bson:"count" json:"count"`
}

type CollectionNftOwnerFiltered struct {
	Items      []*CollectionNftOwner `json:"items" bson:"items"`
	TotalItems int                   `json:"total_items" bson:"total_items"`
}

type FilteredCollections struct {
	Contract        string `json:"contract" bson:"contract"`
	ContractType    string `json:"contract_type" bson:"contract_type"`
	DeployedAtBlock int64  `bson:"deployed_at_block" json:"deployed_at_block"`
}

type Collections struct {
	BaseEntity `bson:",inline"`

	Slug            string `bson:"slug" json:"slug"`
	ContractType    string `bson:"contract_type" json:"contract_type"`
	Contract        string `bson:"contract" json:"contract"`
	Creator         string `bson:"creator" json:"creator"`
	TotalItems      int    `bson:"total_items" json:"total_items"`
	TotalOwners     int    `bson:"total_owners" json:"total_owners"`
	Index           int64  `bson:"index" json:"index"` //autoinscreament
	DeployedAtBlock int64  `bson:"deployed_at_block" json:"deployed_at_block"`

	//TODO - Updateable
	Cover       string `bson:"cover" json:"cover"`
	Thumbnail   string `bson:"thumbnail" json:"thumbnail"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Social      Social `json:"social" bson:"social"`
	Status      int    `json:"status" bson:"status"` // -1: disable, 0: enable
}

type CollectionNftThumbnail struct {
	Contract      string `bson:"contract"`
	Thumbnail     string `bson:"thumbnail"`
	NftImage      string `bson:"nft_image"`
	NftTokenUri   string `bson:"nft_token_uri"`
	NftTokenIDInt int64  `bson:"nft_token_id_int"`
}

type FilterFiles struct {
	BaseFilters
	Name       *string
	UploadedBy *string
}

func (u Collections) CollectionName() string {
	return utils.COLLECTION_COLLECTIONS
}
