package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"v01/data/request"
	"v01/data/response"
	"v01/domain"
	"v01/repository"
)

type CustomerServiceImpl struct {
	repository repository.CustomerRepository
	Validate   *validator.Validate
}

func NewCustomerServiceImpl(customerRepository repository.CustomerRepository, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		repository: customerRepository,
		Validate:   validate,
	}
}

func (c CustomerServiceImpl) CreateCustomer(customerRequest request.CustomerRequest) (string, error) {
	err := c.Validate.Struct(customerRequest)
	if err != nil {
		return "", err
	}

	customer := domain.Customer{
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Email:     customerRequest.Email,
		Address: domain.Address{
			Street:      customerRequest.Address.Street,
			ZipCode:     customerRequest.Address.ZipCode,
			HouseNumber: customerRequest.Address.HouseNumber,
		},
	}
	res, err := c.repository.CreateCustomer(customer)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (c CustomerServiceImpl) FindAllCustomer() ([]response.CustomerResponse, error) {
	var customerRes []response.CustomerResponse
	result, err := c.repository.FindAllCustomer()
	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	for _, value := range result {
		customer := response.CustomerResponse{
			Id:        value.Id.Hex(),
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Email:     value.Email,
			Address: response.CustomerAddress{
				Street:      value.Address.Street,
				HouseNumber: value.Address.HouseNumber,
				ZipCode:     value.Address.ZipCode,
			},
		}
		customerRes = append(customerRes, customer)
	}
	return customerRes, nil
}

func (c CustomerServiceImpl) UpdateCustomer(customerId string, customerUpdateRequest request.CustomerUpdateRequest) (string, error) {
	value := domain.Customer{
		FirstName: customerUpdateRequest.FirstName,
		LastName:  customerUpdateRequest.LastName,
		Email:     customerUpdateRequest.Email,
		Address: domain.Address{
			Street:      customerUpdateRequest.Address.Street,
			ZipCode:     customerUpdateRequest.Address.ZipCode,
			HouseNumber: customerUpdateRequest.Address.HouseNumber,
		},
	}
	customerId, err := c.repository.UpdateCustomer(customerId, value)
	if err != nil {
		return "", err
	}
	return customerId, nil
}

func (c CustomerServiceImpl) IsCustomerExists(customerId string) (bool, error) {
	data, err := c.FindById(customerId)
	if err != nil {
		return false, err
	}

	res := response.CustomerResponse{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Address: response.CustomerAddress{
			Street:      data.Address.Street,
			HouseNumber: data.Address.HouseNumber,
			ZipCode:     data.Address.ZipCode,
		},
	}
	return res.Id != "", nil
}

func (c CustomerServiceImpl) FindById(customerId string) (response.CustomerResponse, error) {
	data, err := c.repository.FindById(customerId)
	if err != nil {
		return response.CustomerResponse{}, err
	}

	res := response.CustomerResponse{
		Id:        data.Id.Hex(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Address: response.CustomerAddress{
			Street:      data.Address.Street,
			HouseNumber: data.Address.HouseNumber,
			ZipCode:     data.Address.ZipCode,
		},
	}
	return res, nil
}

func (c CustomerServiceImpl) DeleteCustomer(customerId string) (string, error) {
	customerId, err := c.repository.DeleteCustomer(customerId)

	if err != nil {
		return "", err
	}

	return customerId, nil
}
