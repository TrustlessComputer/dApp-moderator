package entity

import (
	"time"
)

type BaseEntity struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
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

type FilterString struct {
	Keyword           string
	ListCollectionIDs string
	ListPrices        string
	ListIDs           string
}
