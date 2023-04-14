package entity

import (
	"dapp-moderator/utils/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt *time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at"`
}

func (b *BaseEntity) ToBson() (*bson.D, error) {
	now := time.Now().UTC()
	b.CreatedAt = &now
	return helpers.ToDoc(b)
}

func (b *BaseEntity) SetCreatedAt() {
	now := time.Now().UTC()
	b.CreatedAt = &now
}

func (b *BaseEntity) SetUpdatedAt() {
	now := time.Now().UTC()
	b.UpdatedAt = &now

}

func (b *BaseEntity) SetDeletedAt() {
	now := time.Now().UTC()
	b.DeletedAt = &now
}

func (b *BaseEntity) Id() string {
	return b.ID.Hex()
}

type FilterString struct {
	Keyword           string
	ListCollectionIDs string
	ListPrices        string
	ListIDs           string
}
