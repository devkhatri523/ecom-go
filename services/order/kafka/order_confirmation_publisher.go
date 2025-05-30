package kafka

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-service/event"
	"github.com/devkhatri523/ecom-go/go-service/logger"
	"github.com/devkhatri523/order-service/data/response"
)

type OrderConfirmationPublisher interface {
	PublishOrderCreated(ctx context.Context, confirmation *response.OrderConfirmation)
}

type OrderConfirmationPublisherImpl struct {
	publisher event.Publisher[kafka.Event]
}

func NewOrderConfirmationPublisher(kafkapublisher event.Publisher[kafka.Event]) OrderConfirmationPublisher {
	return &OrderConfirmationPublisherImpl{
		publisher: kafkapublisher,
	}
}

func (o OrderConfirmationPublisherImpl) PublishOrderCreated(ctx context.Context, confirmation *response.OrderConfirmation) {
	logger.Default().InfoWithCtxf(ctx, "Publishing order created event for order id  : %s", confirmation.OrderReference)
	topic := config.Default().GetString("event.order.created.topic")
	b, err := buildOrderEventPayload(ctx, confirmation)
	if err != nil {
		logger.Default().ErrorWithCtxf(ctx, "Error while producing order created event for id  : %s Error : %s", confirmation.OrderReference, err)
		return
	}
	deliveryReport := make(chan kafka.Event, 100000)
	err = o.publisher.Publish(topic, b, deliveryReport)
	if err != nil {
		logger.Default().ErrorWithCtxf(ctx, "Error in kafka while sending order created event for id  : %s error : %s topic : %s", confirmation.OrderReference, err, topic)
		return
	}
	c := <-deliveryReport
	m := c.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		logger.Default().ErrorWithCtxf(ctx, "Error in kafka broker while sending order created event for id  : %s error : %s topic : %s", confirmation.OrderReference, m.TopicPartition.Error, topic)
	}

	defer close(deliveryReport)
}
func buildOrderEventPayload(ctx context.Context, order *response.OrderConfirmation) ([]byte, error) {

	b, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	return b, nil
}
