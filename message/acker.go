package message

type Acker interface {
	Ack(multiple bool) error
}
