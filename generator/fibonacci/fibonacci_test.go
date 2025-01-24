package fibonacci

import (
	"log/slog"
	"testing"

	"github.com/dihedron/snoop/test"
)

func TestFibonacciGenerator(t *testing.T) {
	test.Setup(t)
	for n := range Series(1_000_000_000_000_000_000) {
		slog.Info("received item", "value", n)
	}
}
