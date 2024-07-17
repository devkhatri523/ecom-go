package service

import (
	"v01/data/request"
	"v01/data/response"
)

type CustomerService interface {
	CreateCustomer(customerRequest request.CustomerRequest) (string, error)
	FindAllCustomer() ([]response.CustomerResponse, error)
	UpdateCustomer(customerId string, customerUpdateRequest request.CustomerUpdateRequest) (string, error)
	IsCustomerExists(customerId string) (bool, error)
	FindById(customerId string) (response.CustomerResponse, error)
	DeleteCustomer(customerId string) (string, error)
}
