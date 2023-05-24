package entity

import (
	"dapp-moderator/utils/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IEntity interface {
	CollectionName() string
	SetCreatedAt()
	SetUpdatedAt()
	SetDeletedAt()
	SetID()
	Decode(from *primitive.D) error
}

type SortType int

const (
	SORT_ASC  SortType = 1
	SORT_DESC SortType = -1
)

type BaseFilters struct {
	Page   int64
	Offset int64
	Limit  int64
	SortBy string
	Sort   SortType
}

func (b *BaseEntity) SetID() {
	b.ID = primitive.NewObjectID()
}

func (b *BaseEntity) Decode(from *primitive.D) error {
	err := helpers.Transform(from, b)
	if err != nil {
		return err
	}
	return nil
}
