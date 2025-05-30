package pagination

type CursorPagination struct {
	PageSize     int
	AfterCursor  string
	BeforeCursor string
}
