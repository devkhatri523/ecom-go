package domain

import (
	"encoding/json"
	"v01/enum"
)

type OrderConfirmation struct {
	OrderReference   string             `json:"orderReference"`
	TotalAmount      float64            `json:"totalAmount"`
	PaymentMethod    enum.PaymentMethod `json:"paymentMethod"`
	CustomerResponse Customer           `json:"customer"`
	Products         []Purchase         `json:"products"`
}

func (r OrderConfirmation) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
