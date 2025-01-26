package transformers

import (
	"errors"

	"github.com/dihedron/snoop/syslog"
	"github.com/dihedron/snoop/transform"
)

// Accept accepts the value if the condition is true. This
// filter does not affect the value flowing through.
func WriteToSyslog[T any](syslog *syslog.Syslog, accept func(value T) bool) transform.F[T] {
	return func(value T) (T, error) {
		// if accept(value) {
		// 	slog.Debug("logging message for inclusion into syslog...", "type", format.TypeAsString(value))
		// 	if m, ok := any(value).(*notification.Identity); ok {
		// 		slog.Info("sending message to syslog", "type", m.EventType)
		// 		// we cannot stop the pipeline even if there's an error
		// 		// writing to syslog, but we should at least log it
		// 		json := format.ToJSON(m)
		// 		if err := syslog.Send(&syslog.Message{
		// 			Facility: rfc5424.FacilityAuthpriv,
		// 			Severity: rfc5424.SeverityEmergency,
		// 			ID:       "OpenStack",
		// 			Content:  json,
		// 			Data: map[string][]string{
		// 				"user":   {"name=Andrea", "surname=Funtò", "id=a123456"},
		// 				"tenant": {"region=regionOne", "name=event-broker", "id=1234567890"},
		// 			},
		// 		}); err != nil {
		// 			slog.Warn("error sending OpenStack notification to syslog", "error", err)
		// 		} else {
		// 			slog.Debug("message sent to syslog", "type", m.EventType)
		// 		}
		// 	}
		// }
		// return value, nil
		return value, errors.New("not implemented")
	}
}

// func (w *Writer[T]) Apply(value T) (T, error) {
// 	if w.acceptor == nil || w.acceptor(value) {
// 		slog.Debug("logging message for inclusion into syslog...", "type", format.TypeAsString(value))
// 		if m, ok := value.(*message.IdentityNotification); ok {
// 			slog.Info("sending message to syslog", "type", m.EventType)
// 			// we cannot stop the pipeline even if there's an error
// 			// writing to syslog, but we should at least log it
// 			var (
// 				json string
// 				err  error
// 			)
// 			if json, err = m.ToJSON(false); err != nil {
// 				slog.Warn("error marshalling OpenStack notification to JSON", "error", err)
// 			}
// 			if err := w.logger.Send(&Message{
// 				Facility: rfc5424.FacilityAuthpriv,
// 				Severity: rfc5424.SeverityEmergency,
// 				ID:       "OpenStack",
// 				Content:  json,
// 				Data: map[string][]string{
// 					"user":   {"name=Andrea", "surname=Funtò", "id=a123456"},
// 					"tenant": {"region=regionOne", "name=event-broker", "id=1234567890"},
// 				},
// 			}); err != nil {
// 				slog.Warn("error sending OpenStack notification to syslog", "error", err)
// 			} else {
// 				slog.Debug("message sent to syslog", "type", m.EventType)
// 			}
// 		}
// 	}
// 	return value, nil
// }
