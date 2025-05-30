package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/devkhatri523/order-service/client"
	"github.com/devkhatri523/order-service/data/request"
	response2 "github.com/devkhatri523/order-service/data/response"
	"github.com/devkhatri523/order-service/domain"
	"github.com/devkhatri523/order-service/kafka"
	"github.com/devkhatri523/order-service/repository"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

type OrderServiceImpl struct {
	orderRepository            repository.OrderRepository
	validate                   *validator.Validate
	orderLineRepository        repository.OrderLineRepository
	orderConfirmationPublisher kafka.OrderConfirmationPublisher
}

func NewOrderServiceImpl(orderRepository repository.OrderRepository, validate *validator.Validate, orderLineRepository repository.OrderLineRepository, orderConfirmationPublisher kafka.OrderConfirmationPublisher) (service OrderService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}
	return &OrderServiceImpl{
		orderRepository:            orderRepository,
		validate:                   validate,
		orderLineRepository:        orderLineRepository,
		orderConfirmationPublisher: orderConfirmationPublisher,
	}, err
}

func (o OrderServiceImpl) CreateOrder(req request.OrderRequest) (orderId int32, err error) {
	customer, err := client.GetCustomerDetails(req.CustomerId)
	if err != nil {
		log.Fatal(err)
	}
	purchaseProducts, err := client.GetPurchaseProducts(req)
	if err != nil {
		log.Fatal(err)
	}
	order := domain.Order{
		Id:            req.Id,
		CustomerId:    req.CustomerId,
		Reference:     req.Reference,
		TotalAmount:   req.Amount,
		PaymentMethod: req.PaymentMethod.String(),
		CreatedAt:     time.Now(),
	}
	result, err := o.orderRepository.CreateOrder(order)
	fmt.Println(result)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range purchaseProducts {
		purchaseProduct := domain.OrderLine{
			Id:        req.Id,
			Quantity:  value.Quantity,
			ProductId: value.ProductId,
			OrderId:   order.Id,
		}
		_, err = o.orderLineRepository.SaveOrderLine(purchaseProduct)
		if err != nil {
			return 0, err
		}
	}
	paymentRequest := request.PaymentRequest{
		Id:             req.CustomerId,
		Amount:         req.Amount,
		PaymentMethod:  req.PaymentMethod,
		OrderId:        order.Id,
		OrderReference: order.Reference,
		Customer: request.Customer{
			Id:        customer.Id,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
		},
	}
	err = client.RequestPayment(paymentRequest)
	if err != nil {
		log.Fatal(err)
	}
	orderConfirmation := &response2.OrderConfirmation{
		OrderReference: order.Reference,
		TotalAmount:    req.Amount,
		PaymentMethod:  req.PaymentMethod,
		CustomerResponse: response2.CustomerResponse{
			Id:        customer.Id,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
		},
		Products: purchaseProducts,
	}
	o.orderConfirmationPublisher.PublishOrderCreated(context.Background(), orderConfirmation)
	return result, nil
}

func (o OrderServiceImpl) FindAllOrder(filter *domain.Filters) (orderResponse []*response2.OrderResponse, pagination *domain.Pagination, err error) {
	result, pagination, err := o.orderRepository.FindAllOrder(filter)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(result)

	for _, value := range result {
		order := &response2.OrderResponse{
			Id:         value.Id,
			CustomerId: value.CustomerId,
			Amount:     value.TotalAmount,
			//	PaymentMethod: value.PaymentMethod,
			Reference: value.Reference,
		}
		orderResponse = append(orderResponse, order)
	}
	return orderResponse, pagination, nil
}

func (o OrderServiceImpl) FindByOrderId(orderId int32) (response2.OrderResponse, error) {
	data, err := o.orderRepository.FindByOrderId(orderId)
	if err != nil {
		return response2.OrderResponse{}, err
	}

	res := response2.OrderResponse{
		Id:         data.Id,
		CustomerId: data.CustomerId,
		Amount:     data.TotalAmount,
		//	PaymentMethod: value.PaymentMethod,
		Reference: data.Reference,
	}
	return res, nil
}
