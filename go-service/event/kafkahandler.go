package event

import (
	"errors"
	"github.com/devkhatri523/ecom-go/go-service/kafka"
)

type KafkaHandler struct {
	Client kafka.ConsumerClient
}

func (h *KafkaHandler) Handle(payload []byte) error {
	return errors.New("method not implemented")
}

func (h *KafkaHandler) HandleWithCb(payload []byte, cb func() bool) error {
	return nil
}
