package entity

type IEntity interface {
	CollectionName() string
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
