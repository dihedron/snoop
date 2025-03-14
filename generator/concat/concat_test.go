package concat

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/generator/integer"
	"github.com/dihedron/snoop/generator/textfile"
	"github.com/dihedron/snoop/test"
)

func TestConcatGenerator(t *testing.T) {
	test.Setup(t)
	files := textfile.New()
	for n := range Concat(files.AllLines("../file/a2m.txt"), files.AllLines("../file/n2z.txt")) {
		slog.Info("received item", "value", n)
	}
}

func TestConcatGeneratorContext(t *testing.T) {
	test.Setup(t)
	ctx1, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	ctx2, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	for n := range Concat(integer.SequenceContext(ctx1, 0, 0, 0), integer.SequenceContext(ctx2, 1, 0, 0)) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
}
