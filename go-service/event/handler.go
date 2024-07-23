package event

type Handler interface {
	Handle(payload []byte) error
	HandleWithCb(payload []byte, cb func() bool) error
}
