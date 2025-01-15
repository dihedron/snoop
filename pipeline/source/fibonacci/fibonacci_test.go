package fibonacci

import (
	"log/slog"
	"testing"
)

func TestFibonacciGenerator(t *testing.T) {
	for n := range Series(1_000_000_000_000_000_000) {
		slog.Info("received item", "value", n)
	}
}
