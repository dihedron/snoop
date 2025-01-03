package accumulator

/*
type Accumulator[T any] struct {
	_t     T
	Values []T
}

func (a *Accumulator[T]) Name() string {
	return fmt.Sprintf("github.com/bancaditalia/brokerd/pipeline/filter/counter/Counter[%T]", a._t)
}

func (a *Accumulator[T]) Process(ctx context.Context, message pipeline.Message) (context.Context, pipeline.Message, error) {
	logger := logging.GetLogger()
	select {
	case <-ctx.Done():
		logger.Debug("context cancelled")
		return ctx, nil, pipeline.ErrAbort
	default:
		logger.Debug("processing message")
		if message, ok := message.(*integer.Message); ok {
			value := T(*message)
			if value%2 == 0 {
				logger.Info("even value %d, forwarding...", value)
				return ctx, message, nil
			} else {
				logger.Info("odd value %d: adding to skipped values in accumulator...", value)
				a.Values = append(a.Values, value)
				message.Ack(true)
				return ctx, nil, pipeline.ErrSkip
			}
		}
	}
	return ctx, nil, fmt.Errorf("invalud message type: expected *int64, git %T", message)
}
*/
