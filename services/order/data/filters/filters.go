package filters

import "github.com/devkhatri523/order-service/data/pagination"

type Filters struct {
	SortBy     string
	SortOrder  string
	Query      string
	Pagination *pagination.CursorPagination `json:"pagination"`
}
