package request

import (
	"encoding/json"
	"github.com/devkhatri523/order-service/data/filters"
	"github.com/devkhatri523/order-service/enum"
)

type OrderRequest struct {
	Id            int32              `json:"id"`
	Reference     string             `json:"reference"`
	Amount        float64            `json:"amount"`
	PaymentMethod enum.PaymentMethod `json:"paymentMethod"`
	CustomerId    string             `json:"customerId"`
	Products      []PurchaseRequest  `json:"products"`
}

type GetAllOrderRequest struct {
	Filters *filters.Filters `json:"filters"`
}

func (r OrderRequest) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
