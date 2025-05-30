package request

import (
	"encoding/json"
	"github.com/devkhatri523/order-service/enum"
)

type PaymentRequest struct {
	Id             string             `json:"id"`
	Amount         float64            `json:"amount"`
	PaymentMethod  enum.PaymentMethod `json:"paymentMethod"`
	OrderId        int32              `json:"orderId"`
	OrderReference string             `json:"orderReference"`
	Customer       Customer           `json:"customer"`
}

func (r PaymentRequest) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

type Customer struct {
	Id        string `json:"customerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `validate:"required,email"`
}
