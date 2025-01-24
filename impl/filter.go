package syslog2

import (
	"log/slog"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/message"
	"github.com/juju/rfc/v2/rfc5424"
)

// Writer is a filter that logs messages to the local system syslogd.
type Writer[T any] struct {
	name       string
	enterprise string
	process    string
	acceptor   MessageAcceptor[T]
	logger     *SysLogger
}

type Option[T any] func(*Writer[T])

func NewWriter[T any](options ...Option[T]) (*Writer[T], error) {
	var err error
	writer := &Writer[T]{}
	for _, option := range options {
		option(writer)
	}
	writer.logger, err = New(writer.name, writer.enterprise, writer.process)
	if err != nil {
		slog.Error("error creating syslogger", "error", err)
		return nil, err
	}
	return writer, nil
}

func WithApplicationName[T any](value string) Option[T] {
	return func(s *Writer[T]) {
		s.name = value
	}
}

func WithEnterpriseId[T any](value string) Option[T] {
	return func(s *Writer[T]) {
		s.enterprise = value
	}
}

func WithProcessId[T any](value string) Option[T] {
	return func(s *Writer[T]) {
		s.process = value
	}
}

// MessageAcceptor is used to mark messages that should be logged
// to syslog.
type MessageAcceptor[T any] func(message T) bool

func WithAcceptor[T any](acceptor MessageAcceptor[T]) Option[T] {
	return func(s *Writer[T]) {
		s.acceptor = acceptor
	}
}

func (w *Writer[T]) Apply(value T) (T, error) {
	if w.acceptor == nil || w.acceptor(value) {
		slog.Debug("logging message for inclusion into syslog...", "type", format.TypeAsString(value))
		if m, ok := value.(*message.IdentityNotification); ok {
			slog.Info("sending message to syslog", "type", m.EventType)
			// we cannot stop the pipeline even if there's an error
			// writing to syslog, but we should at least log it
			var (
				json string
				err  error
			)
			if json, err = m.ToJSON(false); err != nil {
				slog.Warn("error marshalling OpenStack notification to JSON", "error", err)
			}
			if err := w.logger.Send(&Message{
				Facility: rfc5424.FacilityAuthpriv,
				Severity: rfc5424.SeverityEmergency,
				ID:       "OpenStack",
				Content:  json,
				Data: map[string][]string{
					"user":   {"name=Andrea", "surname=Funtò", "id=a123456"},
					"tenant": {"region=regionOne", "name=event-broker", "id=1234567890"},
				},
			}); err != nil {
				slog.Warn("error sending OpenStack notification to syslog", "error", err)
			} else {
				slog.Debug("message sent to syslog", "type", m.EventType)
			}
		}
	}
	return value, nil
}
