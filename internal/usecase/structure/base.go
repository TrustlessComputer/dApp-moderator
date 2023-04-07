package structure

type BaseFilters struct {
	Limit  *string
	Page   *string
	SortBy *string
	Cursor *string
	Sort   *int
}

type UpdateCollection struct {
	Cover       *string `json:"cover"`
	Thumbnail   *string `json:"thumbnail"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Social      Social `json:"social"`
}

type Social struct {
	Website   *string `json:"website"`
	DisCord   *string `json:"discord"`
	Twitter   *string `json:"twitter"`
	Telegram  *string `json:"telegram"`
	Medium    *string `json:"medium"`
	Instagram *string `json:"instagram"`
}
