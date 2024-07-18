package repository

import (
	"gorm.io/gorm"
	"v01/domain"
)

type PaymentRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentRepositoryImpl(Db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		Db: Db,
	}

}

func (p PaymentRepositoryImpl) CreatePayment(payment domain.Payment) (int32, error) {
	result := p.Db.Create(&payment)
	if result.Error != nil {
		return 0, result.Error
	}

	return payment.Id, nil
}
