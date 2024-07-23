package event

import (
	"errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafka2 "github.com/devkhatri523/ecom-go/go-service/kafka"
	"github.com/devkhatri523/ecom-go/go-utils/utils"
)

type KafkaPublisher struct {
	Client kafka2.ProducerClient
}

func NewKafkaPublisher(client kafka2.ProducerClient) Publisher[kafka.Event] {
	return KafkaPublisher{
		Client: client,
	}
}

func (k KafkaPublisher) Publish(topic string, payload []byte, deliveryReport chan kafka.Event) error {
	if utils.IsBlank(topic) {
		return errors.New("kafka topic cannot be empty")
	}
	if payload == nil || len(payload) <= 0 {
		return errors.New("payload cannot be nil or empty")
	}
	producer := k.Client.Get().(*kafka.Producer)
	err := producer.Produce(&kafka.Message{
		Value: payload,
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
	}, deliveryReport)
	return err
}
