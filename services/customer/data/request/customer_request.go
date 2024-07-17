package request

type CustomerRequest struct {
	Id        string          `json:"id"`
	FirstName string          `json:"firstName" validate:"required"`
	LastName  string          `json:"lastName" validate:"required"`
	Email     string          `json:"email" validate:"required,email"`
	Address   CustomerAddress `json:"address" validate:"required"`
}

type CustomerAddress struct {
	Street      string `json:"street" validate:"required"`
	HouseNumber string `json:"houseNumber" validate:"required"`
	ZipCode     string `json:"zipCode" validate:"required"`
}
type CustomerUpdateRequest struct {
	Id        string                `json:"id"`
	FirstName string                `json:"firstName"`
	LastName  string                `json:"lastName"`
	Email     string                `json:"email"`
	Address   CustomerUpdateAddress `json:"address"`
}
type CustomerUpdateAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
}
