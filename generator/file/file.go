package file

import (
	"bufio"
	"context"
	"iter"
	"log/slog"
	"os"
)

// Lines uses the new Go 1.23 style generator to read the given
// files line by line.
func Lines(paths ...string) iter.Seq[string] {
	return LinesContext(context.Background(), paths...)
}

// LinesContext uses the new Go 1.23 style generator to read the given
// files line by line; if aborts when the given context is cancelled.
func LinesContext(ctx context.Context, paths ...string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, path := range paths {
			file, err := os.Open(path)
			if err != nil {
				slog.Error("failure opening input file", "path", path, "error", err)
				return
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				select {
				case <-ctx.Done():
					return
				default:
					text := scanner.Text()
					if !yield(text) {
						return
					}
				}
			}
			if err = scanner.Err(); err != nil {
				slog.Error("error reading file line by line", "path", path, "error", err)
			}
		}
	}
}
