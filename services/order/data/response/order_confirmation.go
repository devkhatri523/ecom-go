package response

import (
	"encoding/json"
	"github.com/devkhatri523/order-service/enum"
)

type OrderConfirmation struct {
	OrderReference   string             `json:"orderReference"`
	TotalAmount      float64            `json:"totalAmount"`
	PaymentMethod    enum.PaymentMethod `json:"paymentMethod"`
	CustomerResponse CustomerResponse   `json:"customer"`
	Products         []PurchaseResponse `json:"purchase"`
}

func (r OrderConfirmation) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
