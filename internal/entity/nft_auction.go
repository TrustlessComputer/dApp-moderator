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

	Name         string      `json:"name" bson:"-"`
	Owner        string      `json:"owner" bson:"-"`
	TokenURI     string      `json:"token_uri" bson:"-"`
	Image        string      `json:"image" bson:"-"`
	MintedAt     float64     `json:"minted_at" bson:"-"`
	Attributes   []NftAttr   `json:"attributes" bson:"-"`
	Metadata     interface{} `json:"metadata" bson:"-"`
	MetadataType string      `json:"metadata_type"  bson:"-"`
	Size         int64       `bson:"-" json:"-"`

	AnimationFileUrl string `json:"animation_file_url,omitempty" bson:"animation_file_url,omitempty"`
	ImageCapture     string `json:"image_capture,omitempty" bson:"image_capture,omitempty"`

	ContractAddress string `json:"collection_address" bson:"collection_address"`
	TokenID         string `json:"token_id" bson:"token_id"`
	TokenIDInt      int64  `json:"token_id_int" bson:"token_id_int"` //use it for sort
	IsAuction       bool   `bson:"is_auction" json:"is_auction"`
}

func (u NftAuctions) CollectionName() string {
	return utils.COLLECTION_NFT_AUCTIONS
}
