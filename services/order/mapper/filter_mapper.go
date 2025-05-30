package mapper

import (
	"github.com/devkhatri523/order-service/data/filters"
	"github.com/devkhatri523/order-service/domain"
)

func MapToDomainFilter(filters *filters.Filters) *domain.Filters {
	var filter = new(domain.Filters)
	if filters != nil {
		var pg *domain.Pagination
		if filters.Pagination != nil {
			pg = &domain.Pagination{
				PageSize:     filters.Pagination.PageSize,
				AfterCursor:  filters.Pagination.AfterCursor,
				BeforeCursor: filters.Pagination.BeforeCursor,
				SortBy:       filters.SortBy,
				SortOrder:    filters.SortOrder,
			}

		} else {
			pg = &domain.Pagination{
				PageSize:     100,
				AfterCursor:  "",
				BeforeCursor: "",
				SortBy:       filters.SortBy,
				SortOrder:    filters.SortOrder,
			}

		}

		filter = &domain.Filters{
			Pagination: pg,
		}
	}
	return filter
}
