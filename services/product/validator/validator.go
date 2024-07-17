package validator

import (
	"errors"
	"v01/data/request"
)

func ValidateProductCreateRequest(request request.ProductRequest) error {
	if request.Name == "" {
		return errors.New("product  name cannot be null")
	}
	if request.Description == "" {
		return errors.New("product description cannot be null")
	}
	if request.AvailableQuantity == 0 {
		return errors.New("product quantity   should  be greater than zero")
	}
	if request.Price == 0 {
		return errors.New("product quantity   should  be greater than zero")
	}
	return nil
}
