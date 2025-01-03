package flow

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"sync"

	"github.com/dihedron/snoop/pipeline"
)

// Flow represents a data flow pipeline; it allows to specify a single
// Source (which must comply with the pipeline.Source interface) and an
// ordered list of Filters (which must comply with the pipeline.Filter
// interface). The pipeline also accepts a set of callbacks which are
// called whenever an error occurs and have a chance to abort the whole
// pipeline.
type Flow struct {
	source    pipeline.Source
	filters   []pipeline.Filter
	sink      pipeline.Sink
	buffering int
	wg        sync.WaitGroup
	cancel    context.CancelFunc
	lock      sync.RWMutex
}

// Option is the type for functional options.
type Option func(*Flow)

// New creates a new Flow, applying all the provided functional options.
func New(options ...Option) *Flow {
	f := &Flow{
		filters: []pipeline.Filter{},
	}
	for _, option := range options {
		option(f)
	}
	return f
}

// From allows to specify the Flow source.
func From(source pipeline.Source) Option {
	return func(f *Flow) {
		if source != nil {
			f.source = source
		}
	}
}

// Through allows to add one or more Filters to the Flow.
func Through(filters ...pipeline.Filter) Option {
	return func(f *Flow) {
		for _, filter := range filters {
			if filter != nil {
				f.filters = append(f.filters, filter)
			}
		}
	}
}

// Into allows to specify the Flow sink.
func Into(sink pipeline.Sink) Option {
	return func(f *Flow) {
		if sink != nil {
			f.sink = sink
		}
	}
}

// DefaultBuffering is the default value for the internal buffering.
const DefaultBuffering = 1

// WithBuffering sets the number of messages that might pile up in
// the output channel; this is useful when the pipeline is cancelled
// or terminates because it allows up to "buffering" messages to be
// retrieved from the receiving goroutine; note: a value of 0 means
// the pipeline is unbuffered.
func WithBuffering(buffering int) Option {
	return func(p *Flow) {
		buffering = DefaultBuffering
		if buffering >= 0 {
			p.buffering = buffering
		}
	}
}

// // Start starts the pipeline by retrieving messages from the underlying Source and
// // piping the Messages into the filters and eventually the output message channel.
// // Source, filters and error callbacks are also provided a post-only quit channel
// // on which they can signal if the pipeline has to abort immediately.
// func (f *Flow) Start(ctx context.Context) (<-chan pipeline.Message, error) {
// 	logger := logging.GetLogger()
// 	if p == nil || p.source == nil {
// 		return nil, errors.New("invalid pipeline")
// 	}
// 	inputs, err := p.source.Open(ctx)
// 	if err != nil {
// 		logger.Errorf("error opening source: %v", err)
// 		return nil, err
// 	}
// 	outputs := make(chan pipeline.Message, p.buffering)
// 	// the whole pipeline is cancellable
// 	cancellable, cancel := context.WithCancel(ctx)
// 	p.setCancelFunc(cancel)
// 	p.wg.Add(1)
// 	go func(ctx context.Context) {
// 		var err error
// 		defer func() {
// 			logger.Debug("closing the output channel")
// 			close(outputs)
// 			logger.Debug("pipeline message pump goroutine terminated")
// 			p.wg.Done()

// 		}()
// 		logger.Debug("entering the pipeline message pump goroutine...")
// 	messages:
// 		for {
// 			select {
// 			case message, ok := <-inputs:
// 				if ok {
// 					// loop over the filters and let them process the Message; when
// 					// the chain is complete, output the transformed message.
// 				filtering:
// 					for _, filter := range p.filters {
// 						ctx, message, err = filter.Process(ctx, message)
// 						switch err {
// 						case nil:
// 							// logger.Debugf("filter '%s' processing successful", filter.Name())
// 							continue filtering
// 						case pipeline.ErrSkip:
// 							logger.Warnf("filter '%s' requested skipping of message %v", filter.Name(), message)
// 							continue messages
// 						case pipeline.ErrDone:
// 							logger.Warnf("filter '%s' requested skipping of further processing for message %v", filter.Name(), message)
// 							break filtering
// 						case pipeline.ErrAbort:
// 							logger.Errorf("filter '%s' requested aborting the pipeline", filter.Name())
// 							break messages
// 						default:
// 							logger.Warnf("filter '%s' returned a (recoverable?) error:%v", filter.Name(), err)
// 						}
// 					}
// 					// we emit the message; there must be a sink that
// 					// acknowledges it; if the message has been transformed,
// 					// the filter must make sure that a reference to the
// 					// original message's Ack() method is available and
// 					// invoked when acknowledging the final incarnation of
// 					// the message, or there will be no ack at all.
// 					outputs <- message
// 				} else {
// 					logger.Debug("no more input message, quitting immediately!")
// 					break messages
// 				}
// 			case <-ctx.Done():
// 				logger.Debug("context cancelled, quitting immediately!")
// 				break messages
// 			}
// 		}
// 		logger.Debugf("terminating pipeline message pump...")
// 	}(cancellable)
// 	return outputs, nil
// }

// // Stop stops the pipeline by cancelling the context; if Source,
// // Filters and error callbacks are well-behaved, the pipeline should
// // stop in a short while.
// func (p *Flow) Stop() {
// 	logger := logging.GetLogger()
// 	logger.Info("stopping the pipeline")
// 	p.lock.RLock()
// 	defer p.lock.RUnlock()
// 	logger.Info("cancelling the inner context")
// 	p.cancel()
// 	logger.Info("waiting for the pipeline message pump goroutine to terminate")
// 	p.wg.Wait()
// 	logger.Info("pipeline message pump goroutine terminated")
// }

// Execute starts the pipeline by retrieving messages from the underlying Source and
// piping the Messages into the filters and eventually the pipeline sink.
func (f *Flow) Execute(ctx context.Context) error {
	if f == nil || f.source == nil || f.sink == nil {
		return errors.New("invalid flow")
	}
	inputs, err := f.source.Emit(ctx)
	if err != nil {
		slog.Error("error opening source", "error", err)
		return err
	}
	slog.Debug("entering the pipeline message pump goroutine...")
messages:
	for {
		select {
		case message, ok := <-inputs:
			if ok {
				// loop over the filters and let them process the Message; when
				// the chain is complete, output the transformed message.
				slog.Debug("processing new incoming message...", "type", fmt.Sprintf("%T", message))
			filtering:
				for _, filter := range f.filters {
					backup := message
					ctx, message, err = filter.Process(ctx, message)
					switch err {
					case nil:
						slog.Debug("filter processing successful", "filter", filter.Name())
						continue filtering
					case pipeline.ErrSkip:
						slog.Warn("filter requested skipping of message", "filter", filter.Name(), "message", message)
						continue messages
					case pipeline.ErrDone:
						slog.Warn("filter requested skipping of further processing for message", "filter", filter.Name(), "message", message)
						break filtering
					case pipeline.ErrAbort:
						slog.Error("filter requested aborting the pipeline", "filter", filter.Name())
						break messages
					default:
						// TODO: check if sending the original message to the sink for collection is ok
						slog.Warn("filter returned an error on message ", "filter", filter.Name(), "message", backup, "type", fmt.Sprintf("%T", backup), "error", err)
						message = backup
						break filtering
					}
				}
				// we emit the message; there must be a sink that acknowledges it;
				// if the message has been transformed, the filter must make sure
				// that a reference to the original message's Ack() method is
				// available and invoked when acknowledging the final incarnation
				// of the message, or there will be no ack at all.
				err = f.sink.Collect(ctx, message)
				if err != nil {
					slog.Error("error collecting message into sink", "error", err)
					return err
				}
			} else {
				slog.Debug("no more input message, quitting immediately!")
				break messages
			}
		case <-ctx.Done():
			slog.Debug("context cancelled, quitting immediately!")
			break messages
		}
	}
	slog.Debug("terminating pipeline message pump...")
	return nil
}

// Close stops the flow and releases associated resources (e.g.
// by closing the Source).
func (f *Flow) Close() error {
	if f == nil || f.source == nil || f.sink == nil {
		return errors.New("invalid flow")
	}
	for _, filter := range f.filters {
		if closeable, ok := filter.(io.Closer); ok {
			slog.Info("closing filter", "filter", filter.Name())
			closeable.Close()
		}
	}
	if closeable, ok := f.sink.(io.Closer); ok {
		slog.Info("closing the sink")
		closeable.Close()
	}

	slog.Info("closing the pipeline")
	if closeable, ok := f.source.(io.Closer); ok {
		slog.Info("closing the source")
		return closeable.Close()
	}
	return nil
}

// setCancelFunc is an internal function to avoid race conditions on
// the cancel struct member, used to cancel the pipeline.
func (p *Flow) setCancelFunc(cancel context.CancelFunc) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.cancel = cancel
}
