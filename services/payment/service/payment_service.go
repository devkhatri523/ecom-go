package service

import (
	"v01/data/request"
)

type PaymentService interface {
	CreatePayment(request request.PaymentRequest) (int32, error)
}
