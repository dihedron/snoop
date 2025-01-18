package throttler

import (
	"log/slog"
	"time"
)

// Delay is a filter that delays the value processing
// by inserting a configurable delay.
type Delay[T any] struct {
	delay time.Duration
}

func New[T any](delay time.Duration) *Delay[T] {
	return &Delay[T]{delay: delay}
}

func (d *Delay[T]) Apply(value T) (T, error) {
	slog.Debug("throttling", "delay", d.delay.String(), "value", value)
	time.Sleep(d.delay)
	return value, nil
}
