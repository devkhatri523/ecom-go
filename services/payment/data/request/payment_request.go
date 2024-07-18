package request

import "v01/enum"

type PaymentRequest struct {
	Id             string             `json:"id,omitempty"`
	Amount         float64            `json:"amount"`
	PaymentMethod  enum.PaymentMethod `json:"paymentMethod"`
	OrderId        int32              `json:"orderId"`
	OrderReference string             `json:"orderReference"`
	//	Customer       domain.Customer    `json:"customer"`
}
