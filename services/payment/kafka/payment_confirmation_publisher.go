package kafka

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-service/event"
	"github.com/devkhatri523/ecom-go/go-service/logger"
	"v01/domain"
)

type PaymentConfirmationPublisher interface {
	PublishPaymentConfirmation(ctx context.Context, confirmation *domain.PaymentNotification)
}

func NewPaymentConfirmationPublisher(kafkapublisher event.Publisher[kafka.Event]) PaymentConfirmationPublisher {
	return &PaymentConfirmationPublisherImpl{
		publisher: kafkapublisher,
	}
}

type PaymentConfirmationPublisherImpl struct {
	publisher event.Publisher[kafka.Event]
}

func (o PaymentConfirmationPublisherImpl) PublishPaymentConfirmation(ctx context.Context, confirmation *domain.PaymentNotification) {
	logger.Default().InfoWithCtxf(ctx, "Publishing payment created event for order id  : %s", confirmation.OrderReference)
	topic := config.Default().GetString("event.payment.created.topic")
	b, err := buildPaymnetConfirmationEventPayload(ctx, confirmation)
	if err != nil {
		logger.Default().ErrorWithCtxf(ctx, "Error while producing payment created event for id  : %s Error : %s", confirmation.OrderReference, err)
		return
	}
	deliveryReport := make(chan kafka.Event, 100000)
	err = o.publisher.Publish(topic, b, deliveryReport)
	if err != nil {
		logger.Default().ErrorWithCtxf(ctx, "Error in kafka while sending payment created event for id  : %s error : %s topic : %s", confirmation.OrderReference, err, topic)
		return
	}
	c := <-deliveryReport
	m := c.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		logger.Default().ErrorWithCtxf(ctx, "Error in kafka broker while sending payment created event for id  : %s error : %s topic : %s", confirmation.OrderReference, m.TopicPartition.Error, topic)
	}

	defer close(deliveryReport)
}

func buildPaymnetConfirmationEventPayload(ctx context.Context, confirmation *domain.PaymentNotification) ([]byte, error) {

	b, err := json.Marshal(confirmation)
	if err != nil {
		return nil, err
	}
	return b, nil
}
