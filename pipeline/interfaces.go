package pipeline

// Acknowledgeable is the interface that delivered messages must comply with
// if they must be acknowledged.
type Acknowledgeable interface {
	// Ack is used to notify the source that the Message has been processed;
	// this is relevant in those cases where Messages have to be explicitly
	// removed from the source, depending on whether they have been acquired
	// by the client (see RabbitMQ deliveries).
	Ack(multiple bool) error
}

// MessageWrapper is a generic mechanism to convey the original message
// alongside the modified, or wrapping message; this allows to keep track
// of the acknowledging logic, so that if the original message must
type MessageWrapper struct {
	wrapped Acknowledgeable
}

func (w *MessageWrapper) Ack(multiple bool) error {
	return w.wrapped.Ack(multiple)
}
