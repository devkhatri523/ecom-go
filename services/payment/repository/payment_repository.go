package repository

import "v01/domain"

type PaymentRepository interface {
	CreatePayment(customer domain.Payment) (int32, error)
}
