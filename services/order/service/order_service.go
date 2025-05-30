package service

import (
	"github.com/devkhatri523/order-service/data/request"
	response2 "github.com/devkhatri523/order-service/data/response"
	"github.com/devkhatri523/order-service/domain"
)

type OrderService interface {
	CreateOrder(req request.OrderRequest) (orderId int32, err error)
	FindAllOrder(filter *domain.Filters) ([]*response2.OrderResponse, *domain.Pagination, error)
	FindByOrderId(orderId int32) (response2.OrderResponse, error)
}
