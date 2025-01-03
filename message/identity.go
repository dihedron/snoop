package message

import (
	"encoding/json"
	"log/slog"

	"github.com/dihedron/snoop/message/osptime"
)

type IdentityNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          struct {
		TypeURI   string                `json:"typeURI,omitempty" yaml:"typeURI,omitempty"`
		EventType string                `json:"eventType,omitempty" yaml:"eventType,omitempty"`
		ID        string                `json:"id,omitempty" yaml:"id,omitempty"`
		EventTime osptime.OpenStackTime `json:"eventTime,omitempty" yaml:"eventTime,omitempty"`
		Action    string                `json:"action,omitempty" yaml:"action,omitempty"`
		Outcome   string                `json:"outcome,omitempty" yaml:"outcome,omitempty"`
		Observer  struct {
			ID      string `json:"id,omitempty" yaml:"id,omitempty"`
			TypeURI string `json:"typeURI,omitempty" yaml:"typeURI,omitempty"`
		} `json:"observer,omitempty" yaml:"observer,omitempty"`
		Initiator struct {
			ID      string `json:"id,omitempty" yaml:"id,omitempty"`
			TypeURI string `json:"typeURI,omitempty" yaml:"typeURI,omitempty"`
			Host    struct {
				Address string `json:"address,omitempty" yaml:"address,omitempty"`
				Agent   string `json:"agent,omitempty" yaml:"agent,omitempty"`
			} `json:"host,omitempty" yaml:"host,omitempty"`
			UserID     string `json:"user_id,omitempty" yaml:"user_id,omitempty"`
			ProjectID  string `json:"project_id,omitempty" yaml:"project_id,omitempty"`
			RequestID  string `json:"request_id,omitempty" yaml:"request_id,omitempty"`
			Username   string `json:"username,omitempty" yaml:"username,omitempty"`
			Credential struct {
				Type             string   `json:"type,omitempty" yaml:"type,omitempty"`
				Token            string   `json:"token,omitempty" yaml:"token,omitempty"`
				IdentityProvider string   `json:"identity_provider,omitempty" yaml:"identity_provider,omitempty"`
				User             string   `json:"user,omitempty" yaml:"user,omitempty"`
				Groups           []string `json:"groups,omitempty" yaml:"groups,omitempty"`
			} `json:"credential,omitempty" yaml:"credential,omitempty"`
		} `json:"initiator,omitempty" yaml:"initiator,omitempty"`
		Target struct {
			ID      string `json:"id,omitempty" yaml:"id,omitempty"`
			TypeURI string `json:"typeURI,omitempty" yaml:"typeURI,omitempty"`
		} `json:"target,omitempty" yaml:"target,omitempty"`
		ResourceInfo        string `json:"resource_info,omitempty" yaml:"resource_info,omitempty"`
		Role                string `json:"role,omitempty" yaml:"role,omitempty"`
		Project             string `json:"project,omitempty" yaml:"project,omitempty"`
		User                string `json:"user,omitempty" yaml:"user,omitempty"`
		InheritedToProjects bool   `json:"inherited_to_projects,omitempty" yaml:"inherited_to_projects,omitempty"`
		Group               string `json:"group,omitempty" yaml:"group,omitempty"`
		Reason              struct {
			ReasonCode string `json:"reasonCode,omitempty" yaml:"reasonCode,omitempty"`
			ReasonType string `json:"reasonType,omitempty" yaml:"reasonType,omitempty"`
		} `json:"reason,omitempty" yaml:"reason,omitempty"`
	} `json:"payload,omitempty" yaml:"payload,omitempty"`
}

// FromJSON populates a IdentityNotification using
// the data in the provided JSON.
func (msg *IdentityNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing identity notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the IdentityNotification
// in JSON format, pretty-printed or not.
func (msg *IdentityNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling identity notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling identity notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("identity notification marshaled to JSON")
	return string(bytes), nil
}

// ToString converts the IdentityNotification into a string containing
// its JSON one-liner representation.
func (msg *IdentityNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
