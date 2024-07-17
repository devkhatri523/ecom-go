package response

type CustomerResponse struct {
	Id        string          `json:"id"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Email     string          `validate:"required,email"`
	Address   CustomerAddress `json:"address"`
}

type CustomerAddress struct {
	Street      string `json:"street" `
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode" `
}
