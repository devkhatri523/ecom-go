package response

import (
	"encoding/json"
	"github.com/devkhatri523/order-service/enum"
)

type OrderResponse struct {
	Id            int32              `json:"id"`
	Reference     string             `json:"reference"`
	Amount        float64            `json:"amount"`
	PaymentMethod enum.PaymentMethod `json:"paymentMethod"`
	CustomerId    string             `json:"customerId"`
}

func (r OrderResponse) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
