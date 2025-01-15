package file

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/pipeline/source/concat"
)

func TestFileContextGenerator(t *testing.T) {
	t.Log("test with cancellation after 10 items")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range FileContext(ctx, "../../flow/test.txt") {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
}

func TestFilesGenerator(t *testing.T) {
	t.Log("test with one file")
	for n := range File("../../flow/test.txt") {
		slog.Info("received item", "value", n)
	}
	t.Log("test with two files")
	for n := range concat.Concat(File("../../flow/test.txt"), File("../../flow/test.txt")) {
		slog.Info("received item", "value", n)
	}
	t.Log("test with non-existing file")
	for n := range File("../../flow/non_existing.txt") {
		slog.Info("received item", "value", n)
	}
}
