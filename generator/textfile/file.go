package textfile

import (
	"bufio"
	"context"
	"errors"
	"iter"
	"log/slog"
	"os"
)

type TextFiles struct {
	err error
}

func New() *TextFiles {
	return &TextFiles{}
}

func (f *TextFiles) Err() error {
	return f.err
}

func (f *TextFiles) Reset() {
	f.err = nil
}

// Lines uses the new Go 1.23 style generator to read the given
// files line by line.
func (f *TextFiles) AllLines(paths ...string) iter.Seq[string] {
	return f.AllLinesContext(context.Background(), paths...)
}

// LinesContext uses the new Go 1.23 style generator to read the given
// files line by line; if aborts when the given context is cancelled.
func (f *TextFiles) AllLinesContext(ctx context.Context, paths ...string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, path := range paths {
			file, err := os.Open(path)
			if err != nil {
				f.err = errors.Join(f.err, err)
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
				f.err = errors.Join(f.err, err)
			}
		}
	}
}
