package helper

import (
	"strings"

	"github.com/devkhatri523/ecom-go/cursor-paginator/paginator"
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	"github.com/devkhatri523/order-service/domain"
)

func CreateGetAllOrderPagination(pagination *domain.Pagination) *paginator.Paginator {
	config := &paginator.Config{
		Limit: 100,
		Order: paginator.DESC,
		Rules: []paginator.Rule{
			{
				Key:     "Id",
				Order:   paginator.DESC,
				SQLRepr: "id",
			},
		},
	}
	p := paginator.New(config)
	if pagination != nil {
		rule := MapGetAllOrderRecordSortByAndSortOrder(pagination.SortBy, pagination.SortOrder)
		if pagination.PageSize > 0 {
			p.SetLimit(pagination.PageSize)
		}
		if utils.IsNotBlank(pagination.AfterCursor) {
			p.SetAfterCursor(pagination.AfterCursor)
		}
		if utils.IsNotBlank(pagination.BeforeCursor) {
			p.SetBeforeCursor(pagination.BeforeCursor)
		}
		iRule := paginator.Rule{
			Key:     "Id",
			Order:   paginator.DESC,
			SQLRepr: "id",
		}
		p.SetRules(rule, iRule)

	}
	return p

}
func MapGetAllOrderRecordSortByAndSortOrder(sortBy string, sortOrder string) paginator.Rule {
	rule := paginator.Rule{}
	so := strings.ToUpper(sortOrder)
	if so == "ASC" && so != "DESC" {
		sortOrder = "DESC"
	}
	if strings.EqualFold(sortBy, "id") {
		rule.Key = "id"
		rule.SQLRepr = "id"
	} else if strings.EqualFold(sortBy, "CustomerId") {
		rule.Key = "CustomerId"
		rule.SQLRepr = "CustomerId"
	} else if strings.EqualFold(sortBy, "CreatedAt") {
		rule.Key = "createdAt"
		rule.SQLRepr = "created_at"
	} else if strings.EqualFold(sortBy, "TotalAmount") {
		rule.Key = "TotalAmount"
		rule.SQLRepr = "total_amount"
	}
	rule.Order = paginator.Order(so)
	return rule

}
