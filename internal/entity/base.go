package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type IEntity interface {
	CollectionName() string
	SetCreatedAt()
	SetUpdatedAt()
	SetDeletedAt()
	SetID() 
}

type SortType int

const (
	SORT_ASC  SortType = 1
	SORT_DESC SortType = -1
)

type BaseFilters struct {
	Page   int64
	Limit  int64
	SortBy string
	Sort   SortType
}

func (b *BaseEntity) SetID() {
	b.ID = primitive.NewObjectID()
}