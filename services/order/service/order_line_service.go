package service

import (
	response2 "github.com/devkhatri523/order-service/data/response"
	"github.com/devkhatri523/order-service/domain"
)

type OrderLineService interface {
	SaveOrderLine(line domain.OrderLine) (int32, error)
	FindAllByOrderId(orderLineId int32) (response2.OrderLineResponse, error)
}
