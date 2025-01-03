package syslogger

import (
	"fmt"
	"strings"

	"github.com/juju/rfc/v2/rfc5424"
)

// Message contains the set of information that is specific
// to a given message: its priority (in terms of facility and
// severity), the message type identifier (to group similar
// messages, the message text and optionally a set of
// parameters.
type Message struct {
	Facility rfc5424.Facility
	Severity rfc5424.Severity
	ID       string
	Content  interface{} // either a string or an object that will be marshalled to JSON
	Data     map[string][]string
}

// Element is an implementation of the StructuredDataElement interface.
type Element struct {
	id     rfc5424.StructuredDataName
	params []rfc5424.StructuredDataParam
}

func (s *SysLogger) newStructuredDataElement(id string, parameters ...string) *Element {
	params := make([]rfc5424.StructuredDataParam, len(parameters))
	for i, str := range parameters {
		parts := strings.SplitN(str, "=", 2)
		params[i].Name = rfc5424.StructuredDataName(parts[0])
		params[i].Value = rfc5424.StructuredDataParamValue(parts[1])
	}
	return &Element{
		id:     rfc5424.StructuredDataName(fmt.Sprintf("%s@%s", id, s.enterprise)),
		params: params,
	}
}

func (e *Element) ID() rfc5424.StructuredDataName {
	return e.id
}

func (e *Element) Params() []rfc5424.StructuredDataParam {
	return e.params
}

func (e *Element) Validate() error {
	return nil
}
