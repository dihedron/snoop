package syslogger

import (
	"context"
	"log/slog"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/message"
	"github.com/dihedron/snoop/pipeline"
	"github.com/juju/rfc/v2/rfc5424"
)

// SysLogWriter is a filter that logs messages to the
// local system syslogd.
type SysLogWriter struct {
	name       string
	enterprise string
	process    string
	acceptor   MessageAcceptor
	logger     *SysLogger
}

type SysLogWriterOption func(*SysLogWriter)

func NewSysLogWriter(options ...SysLogWriterOption) (*SysLogWriter, error) {
	var err error
	writer := &SysLogWriter{}
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

func WithApplicationName(value string) SysLogWriterOption {
	return func(s *SysLogWriter) {
		s.name = value
	}
}

func WithEnterpriseId(value string) SysLogWriterOption {
	return func(s *SysLogWriter) {
		s.enterprise = value
	}
}

func WithProcessId(value string) SysLogWriterOption {
	return func(s *SysLogWriter) {
		s.process = value
	}
}

// MessageAcceptor is used to mark messages that should be logged
// to syslog.
type MessageAcceptor func(message pipeline.Message) bool

func WithAcceptor(acceptor MessageAcceptor) SysLogWriterOption {
	return func(s *SysLogWriter) {
		s.acceptor = acceptor
	}
}

func (w *SysLogWriter) Name() string {
	return "github.com/dihedron/snoop/syslogger/SysLogWriter"
}

func (w *SysLogWriter) Process(ctx context.Context, msg pipeline.Message) (context.Context, pipeline.Message, error) {
	if w.acceptor == nil || w.acceptor(msg) {
		slog.Debug("logging message for inclusion into syslog...", "type", format.TypeAsString(msg))
		if m, ok := msg.(*message.IdentityNotification); ok {
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
	return ctx, msg, nil
}
