package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"time"
	"v01/data/request"
	"v01/domain"
	"v01/repository"
)

type PaymentServiceImpl struct {
	paymentRepository repository.PaymentRepository
	Validate          *validator.Validate
}

func NewPaymentServiceImpl(respository repository.PaymentRepository, validate *validator.Validate) (service PaymentService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}

	return &PaymentServiceImpl{
		paymentRepository: respository,
		Validate:          validate,
	}, err
}

func (p PaymentServiceImpl) CreatePayment(request request.PaymentRequest) (int32, error) {
	err := p.Validate.Struct(request)
	if err != nil {
		return 0, err
	}
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 100
	randomNumber := rand.Int31n(100)
	m := domain.Payment{
		Id:            randomNumber,
		Amount:        request.Amount,
		PaymentMethod: request.PaymentMethod.String(),
		OrderId:       request.OrderId,
		CreatedAt:     time.Now(),
	}
	paymentId, err := p.paymentRepository.CreatePayment(m)
	if err != nil {
		return 0, err
	}
	return paymentId, nil
}
