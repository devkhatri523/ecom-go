package repository

import "github.com/devkhatri523/order-service/domain"

type OrderLineRepository interface {
	SaveOrderLine(line domain.OrderLine) (int32, error)
	FindAllByOrderId(orderLineId int32) (domain.OrderLine, error)
}
