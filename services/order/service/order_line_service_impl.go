package service

import (
	response2 "github.com/devkhatri523/order-service/data/response"
	"github.com/devkhatri523/order-service/domain"
	"github.com/devkhatri523/order-service/repository"
	"github.com/go-playground/validator/v10"
)

type OrderLineServiceImpl struct {
	orderLineRepository repository.OrderLineRepository
	validate            *validator.Validate
}

func NewOrderLineServiceImpl(orderLineRepository repository.OrderLineRepository, validate *validator.Validate) (service OrderLineService, err error) {
	return &OrderLineServiceImpl{
		orderLineRepository: orderLineRepository,
		validate:            validate,
	}, err
}

func (o OrderLineServiceImpl) SaveOrderLine(line domain.OrderLine) (int32, error) {
	_, err := o.orderLineRepository.SaveOrderLine(line)
	if err != nil {
		return 0, err
	}
	return line.Id, nil
}

func (o OrderLineServiceImpl) FindAllByOrderId(orderLineId int32) (response2.OrderLineResponse, error) {
	data, err := o.orderLineRepository.FindAllByOrderId(orderLineId)
	if err != nil {
		return response2.OrderLineResponse{}, err
	}

	res := response2.OrderLineResponse{
		Id:       data.Id,
		Quantity: data.Quantity,
	}
	return res, nil
}
