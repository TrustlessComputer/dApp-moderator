package entity

import (
	"dapp-moderator/utils"
)

type Nfts struct {
	BaseEntity `bson:",inline"`

	//Collection      string      `json:"collection" bson:"collection"`
	ContractAddress string      `json:"collection_address" bson:"collection_address"`
	TokenID         string      `json:"token_id" bson:"token_id"`
	TokenIDInt      int64       `json:"token_id_int" bson:"token_id_int"` //use it for sort
	ContentType     string      `json:"content_type" bson:"content_type"`
	Name            string      `json:"name" bson:"name"`
	Owner           string      `json:"owner" bson:"owner"`
	TokenURI        string      `json:"token_uri" bson:"token_uri"`
	Image           string      `json:"image" bson:"image"`
	MintedAt        float64     `json:"minted_at" bson:"minted_at"`
	Attributes      []NftAttr   `json:"attributes" bson:"attributes"`
	Metadata        interface{} `json:"metadata" bson:"metadata"`
	MetadataType    string      `json:"metadata_type"  bson:"metadata_type"`
}

type NftAttr struct {
	TraitType string `json:"trait_type" bson:"trait_type"`
	Value     string `json:"value"  bson:"value"`
}

func (u Nfts) CollectionName() string {
	return utils.COLLECTION_NFTS
}
