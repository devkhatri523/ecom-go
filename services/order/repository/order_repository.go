package repository

import "github.com/devkhatri523/order-service/domain"

type OrderRepository interface {
	CreateOrder(order domain.Order) (int32, error)
	FindAllOrder(filter *domain.Filters) ([]*domain.Order, *domain.Pagination, error)
	FindByOrderId(orderId int32) (domain.Order, error)
}
