package entity

import (
	"go.mongodb.org/mongo-driver/bson"
)

type IEntity interface {
	CollectionName() string
	ToBson() (*bson.D, error)
	SetCreatedAt()
	SetUpdatedAt()
	SetDeletedAt()
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
