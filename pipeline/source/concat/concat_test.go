package concat

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/pipeline/source/file"
	"github.com/dihedron/snoop/pipeline/source/integer"
)

func TestConcatGenerator(t *testing.T) {
	for n := range Concat(file.File("../file/a2m.txt"), file.File("../file/n2z.txt")) {
		slog.Info("received item", "value", n)
	}
}

func TestConcatGeneratorContext(t *testing.T) {
	ctx1, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	ctx2, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	for n := range Concat(integer.SequenceContext(ctx1, 0, 0, 0), integer.SequenceContext(ctx2, 1, 0, 0)) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
}
