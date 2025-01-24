package syslog

import (
	"log"
	"log/slog"
	"testing"

	"github.com/dihedron/snoop/test"
	"github.com/juju/rfc/v2/rfc5424"
)

const (
	ApplicationName = "RHOSP-Keystone"
)

type TestData struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

func TestSendPlainTextToSyslog(t *testing.T) {

	test.Setup(t)
	message := "a message sent to syslog"

	syslog, err := New(WithApplication(ApplicationName))
	if err != nil {
		log.Fatal(err)
	}

	if err := syslog.Send(&Message{
		Facility: rfc5424.FacilityAuthpriv,
		Severity: rfc5424.SeverityEmergency,
		ID:       "Login",
		Content:  message,
		Data: map[string][]string{
			"user":   {"name=John", "surname=Doe", "id=a123456"},
			"tenant": {"region=regionOne", "name=event-broker", "id=1234567890"},
		},
	}); err != nil {
		t.Fatal(err)
	}
}

func TestSendObjectAsJSONToSyslog(t *testing.T) {
	test.Setup(t)
	syslog, err := New(WithApplication(ApplicationName))
	if err != nil {
		slog.Error("error opeinnig client", "error", err)
		t.FailNow()
	}

	if err := syslog.Send(&Message{
		Facility: rfc5424.FacilityAuthpriv,
		Severity: rfc5424.SeverityEmergency,
		ID:       "Login",
		Content: &TestData{
			Name:    "Andrea",
			Surname: "Funtò",
			Phone:   "555-123456",
			Address: "Main st., 12 - 00123 (NJ)",
		},
		Data: map[string][]string{
			"user":   {"name=Andrea", "surname=Funtò", "id=a123456"},
			"tenant": {"region=regionOne", "name=event-broker", "id=1234567890"},
		},
	}); err != nil {
		t.Fatal(err)
	}
}
