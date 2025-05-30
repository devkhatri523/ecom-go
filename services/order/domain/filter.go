package domain

type Filters struct {
	Pagination *Pagination
}

type Pagination struct {
	TotalCount   int32
	AfterCursor  string
	BeforeCursor string
	PageSize     int
	SortOrder    string // ASC or DESC
	SortBy       string
}
