package syslogger

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/juju/rfc/v2/rfc5424"
)

const (
	DefaultEnterprise  = "dihedron"
	DefaultSendTimeout = time.Duration(10 * time.Second)
	DefaultSendMaxSize = 0
)

// SysLogger wraps a syslog connection and stores all common
// configuration elements.
type SysLogger struct {
	application string
	hostname    string
	enterprise  string
	process     string
	client      *rfc5424.Client
}

// New creates a new SysLogger, initialising all relevant fields
// to the provided valuer or to the defaults; if application is
// not provided, it defaults to os.Args[0]; if enterprise is not
// provided, it defaults to DefaultEnterprise, which is set to
// "bancaditalia"; if process is not provided, it defaults to the
// current PID.
func New(application, enterprise, process string) (result *SysLogger, err error) {
	result = &SysLogger{
		application: application,
		enterprise:  enterprise,
		process:     process,
	}
	if result.application == "" {
		result.application = os.Args[0]
	}
	if result.enterprise == "" {
		result.enterprise = DefaultEnterprise
	}
	if result.process == "" {
		result.process = fmt.Sprintf("%d", os.Getpid())
	}
	if result.hostname, err = os.Hostname(); err != nil {
		return nil, err
	}
	result.client, err = rfc5424.Open("/dev/log", rfc5424.ClientConfig{
		MaxSize:     DefaultSendMaxSize,
		SendTimeout: DefaultSendTimeout,
	}, func(network, address string) (rfc5424.Conn, error) {
		return net.Dial("unixgram", address)
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// NewWithDefaults uses all defaults values to initialise a SysLogger.
func NewWithDefaults() (*SysLogger, error) {
	return New("", "", "")
}

// Send prepares a message in RFC5424-compliat format and
// sends it through the client.
func (s *SysLogger) Send(message *Message) error {
	msg := rfc5424.Message{
		Header: rfc5424.Header{
			Priority: rfc5424.Priority{
				Severity: message.Severity,
				Facility: message.Facility,
			},
			Timestamp: rfc5424.Timestamp{Time: time.Now()},
			Hostname:  rfc5424.Hostname{FQDN: s.hostname},
			AppName:   rfc5424.AppName(s.application),
			ProcID:    rfc5424.ProcID(s.process),
			MsgID:     rfc5424.MsgID(message.ID),
		},
		StructuredData: rfc5424.StructuredData{},
	}
	if s, ok := (message.Content).(string); ok {
		msg.Msg = s
	} else if message.Content != nil {
		if s, err := json.Marshal(message.Content); err != nil {
			return err
		} else {
			msg.Msg = string(s)
		}
	} else {
		return errors.New("no text in message")
	}
	for k, v := range message.Data {
		msg.StructuredData = append(msg.StructuredData, s.newStructuredDataElement(k, v...))
	}
	return s.client.Send(msg)
}
