package fibonacci

import (
	"context"
	"log/slog"
	"math/big"
	"testing"
)

func TestFibonacci(t *testing.T) {
	aVeryLargeNumber := big.NewInt(1_000_000_000_000_000_000)
	aVeryLargeNumber.Mul(aVeryLargeNumber, aVeryLargeNumber)
	aVeryLargeNumber.Mul(aVeryLargeNumber, aVeryLargeNumber)
	aVeryLargeNumber.Mul(aVeryLargeNumber, aVeryLargeNumber)

	source, err := New()
	if err != nil {
		slog.Error("error creating fibonacci source", "error", err)
	}
	messages, _ := source.Emit(context.Background())
loop:
	for message := range messages {
		if message, ok := message.(*Message); ok {
			if message.Value.Cmp(aVeryLargeNumber) >= 0 {
				slog.Info("upper bound reached, cancelling...")
				break loop
			}
			slog.Info("message received", "value", message.Value)
			message.Ack(false)
		}
	}
	slog.Info("no more messages")
}

func TestFibonacciGenerator(t *testing.T) {
	for n := range Fibonacci(1_000_000_000_000_000_000) {
		slog.Info("received item", "value", n)
	}
}
