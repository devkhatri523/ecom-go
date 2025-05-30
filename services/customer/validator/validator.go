package validator

import (
	"errors"
	"github.com/devkhatri523/ecom-go/go-utils/utils"
	"v01/data/request"
)

func ValidateCustomerCreateRequest(request request.CustomerRequest) error {
	if utils.IsBlank(request.FirstName) {
		return errors.New("customer first name cannot be null")
	}
	if request.LastName == "" {
		return errors.New("customer last name cannot be null")
	}
	if request.Email == "" {
		return errors.New("customer email  cannot be null")
	}
	if request.Address.Street == "" {
		return errors.New("customer street name cannot be null")
	}
	if request.Address.HouseNumber == "" {
		return errors.New("customer house number cannot be null")
	}
	if request.Address.ZipCode == "" {
		return errors.New("customer zipcode  cannot be null")
	}
	return nil
}
