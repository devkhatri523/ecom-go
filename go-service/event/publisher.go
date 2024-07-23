package event

type Publisher[T any] interface {
	Publish(topic string, payload []byte, deliveryReport chan T) error
}
