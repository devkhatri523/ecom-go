package event

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafka2 "github.com/devkhatri523/ecom-go/go-service/kafka"
	"github.com/devkhatri523/ecom-go/go-service/logger"
)

type KafkaLooper struct {
	Client        kafka2.ConsumerClient
	PollTimeoutMs int
	Callback      func() bool
	stop          bool
}

func GetNewKafkaLooper(client kafka2.ConsumerClient, pollTimeoutMs int, callback func() bool) KafkaLooper {
	return KafkaLooper{
		Client:        client,
		PollTimeoutMs: pollTimeoutMs,
		Callback:      callback,
		stop:          true,
	}
}
func (k *KafkaLooper) Loop(handler Handler) {
	consumer := k.Client.Get().(*kafka.Consumer)
	if k.PollTimeoutMs == 0 {
		k.PollTimeoutMs = 100
	}
	var cb func() bool
	if k.Callback != nil {
		cb = func() bool {
			if k.Callback() {
				go func() {
					_, err := consumer.Commit()
					if err != nil {
						logger.Default().Errorf("Error while committing to kafka : ", err)
					}
				}()
				return true
			}
			return false
		}
	}
	for k.stop {
		ev := consumer.Poll(k.PollTimeoutMs)
		switch e := ev.(type) {
		case *kafka.Message:
			var err error
			if cb != nil {
				err = handler.HandleWithCb(e.Value, cb)
			} else {
				err = handler.Handle(e.Value)
			}
			if err != nil {
				logger.Default().Errorf("Error while consuming kafka payload : %s", err)
			}
			break
		case kafka.PartitionEOF:
			logger.Default().Debugf("kafka consumer reached %v", e)
			break
		case kafka.Error:
			logger.Default().Errorf("Error on kafka consumer %v", e)
			break
		}
	}
	logger.Default().Debug("Looper exited.")
}

func (k *KafkaLooper) Stop() {
	k.stop = true
}
