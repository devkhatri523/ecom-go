package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/devkhatri523/ecom-go/go-service/event"
	"github.com/devkhatri523/ecom-go/go-service/logger"
	"v01/domain"
	service "v01/email_service"
)

func NewOrderEventHandler(service service.EmailService) event.Handler {
	return &OrderEventHandler{
		emailService: service,
	}

}

type OrderEventHandler struct {
	emailService service.EmailService
}

func (o OrderEventHandler) Handle(payload []byte) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderEventHandler) HandleWithCb(payload []byte, cb func() bool) error {
	ctx := context.Background()
	if payload == nil {
		logger.Default().ErrorWithCtx(ctx, "Payload null ")
		cb()
		return errors.New("payload null on event")
	}
	defer func() {
		if er := recover(); er != nil {
			logger.Default().ErrorWithCtxf(ctx, "Panic on processing transaction event. Panic : %s", er)
		}
	}()
	e := domain.OrderConfirmation{}
	err := json.Unmarshal(payload, &e)
	if err != nil {
		logger.Default().ErrorWithCtxf(ctx, "Could not unmarshall. %s", err)
		cb()
		return errors.New("could not unmarshall")
	}

	if err != nil {
		logger.Default().ErrorWithCtxf(ctx, "Could not unmarshall error : %s", err)
		cb()
		return errors.New("could not unmarshall  payload")
	}
	err = o.emailService.SendEmail(e)
	if err != nil {
		return err
	}
	cb()
	return nil
}
