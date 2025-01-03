package file

import (
	"bufio"
	"context"
	"iter"
	"log/slog"
	"os"

	"github.com/dihedron/snoop/pipeline"
)

// Source is the concrete source capable of reading data one
// line at a time from a text file, and producing it as items in
// a channel.
type Source struct {
	file *os.File
}

// New creates a new Source.
func New(path string) (*Source, error) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("failure opening input file", "path", path, "error", err)
		return nil, err
	}
	slog.Debug("input file open", "path", path)
	source := &Source{
		file: file,
	}
	return source, nil
}

// Emit opens the channel on which text lines will be emitted one at a time.
func (s *Source) Emit(ctx context.Context) (<-chan pipeline.Message, error) {
	messages := make(chan pipeline.Message, 1)
	go func(ctx context.Context) {
		defer func() {
			slog.Debug("closing output message channel")
			close(messages)
		}()
		scanner := bufio.NewScanner(s.file)
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				slog.Debug("context cancelled")
				return
			default:
				text := scanner.Text()
				slog.Debug("sending text as message")
				messages <- &Message{Value: text}
			}
		}
		slog.Info("done reading file line by line", "path", s.file.Name())
	}(ctx)
	return messages, nil
}

// Close closes the output file descriptor.
func (s *Source) Close() error {
	slog.Info("closing the source")
	if s != nil && s.file != nil {
		slog.Info("closing the underlying file")
		s.file.Close()
	}
	return nil
}

// File uses the new Go 1.23 style generator.
func File(paths ...string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		for _, path := range paths {
			file, err := os.Open(path)
			if err != nil {
				slog.Error("failure opening input file", "path", path, "error", err)
				yield("", err)
				return
			}
			defer file.Close()
			for {
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {

					text := scanner.Text()
					if !yield(text, nil) {
						return
					}
				}
				if err = scanner.Err(); err != nil {
					yield("", err)
					return
				} else {
					break
				}
			}
		}
	}
}

// File uses the new Go 1.23 style generator.
func FileContext(ctx context.Context, paths ...string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
	outer:
		for _, path := range paths {
			file, err := os.Open(path)
			if err != nil {
				slog.Error("failure opening input file", "path", path, "error", err)
				yield("", err)
				return
			}
			defer file.Close()
			for {
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					select {
					case <-ctx.Done():
						break outer
					default:
						text := scanner.Text()
						if !yield(text, nil) {
							return
						}
					}
				}
				if err = scanner.Err(); err != nil {
					yield("", err)
					return
				} else {
					break
				}
			}
		}
	}
}
