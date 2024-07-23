package event

type Looper interface {
	Loop(handler Handler)
	Stop()
}
