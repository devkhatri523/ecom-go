package repository

import "v01/domain"

type CustomerRepository interface {
	CreateCustomer(customer domain.Customer) (string, error)
	FindAllCustomer() ([]domain.Customer, error)
	UpdateCustomer(customerId string, customer domain.Customer) (string, error)
	IsCustomerExists(customerId string) (bool, error)
	FindById(customerId string) (domain.Customer, error)
	DeleteCustomer(customerId string) (string, error)
}
