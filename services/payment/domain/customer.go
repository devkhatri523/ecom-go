package domain

type Customer struct {
	CustomerId string `json:"customerId" validate:"required"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"  validate:"required"`
	Email      string `json:"email"  validate:"required"`
}
