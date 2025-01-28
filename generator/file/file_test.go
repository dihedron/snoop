package file

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/generator/concat"
	"github.com/dihedron/snoop/test"
)

func TestFileContextGenerator(t *testing.T) {
	test.Setup(t)
	slog.Info("test with cancellation after 10 items")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range LinesContext(ctx, "test.txt") {
		slog.Debug("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
}

func TestFilesGenerator(t *testing.T) {
	test.Setup(t)
	slog.Info("test with one file")
	for n := range Lines("test.txt") {
		slog.Debug("received item", "value", n)
	}
	slog.Info("test with non-existing file")
	for n := range Lines("non_existing.txt") {
		slog.Debug("received item", "value", n)
	}
}

func TestMultipleFilesGenerator(t *testing.T) {
	test.Setup(t)
	slog.Info("test with two files")
	for n := range Lines("a2m.txt", "n2z.txt") {
		slog.Debug("received item", "value", n)
	}
}

func TestConcatFilesGenerator(t *testing.T) {
	test.Setup(t)
	slog.Info("test with two files")
	for n := range concat.Concat(Lines("a2m.txt"), Lines("n2z.txt")) {
		slog.Debug("received item", "value", n)
	}
}
