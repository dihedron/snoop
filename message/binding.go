package message

import (
	"encoding/json"
	"log/slog"
)

type BindingNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          BindingNotificationPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type BindingNotificationPayload struct {
	BindingID string `json:"binding_id,omitempty" yaml:"binding_id,omitempty"`
	Binding   struct {
		Host     string `json:"host,omitempty" yaml:"host,omitempty"`
		VifType  string `json:"vif_type,omitempty" yaml:"vif_type,omitempty"`
		VnicType string `json:"vnic_type,omitempty" yaml:"vnic_type,omitempty"`
		Status   string `json:"status,omitempty" yaml:"status,omitempty"`
		Profile  struct {
			MigratingTo string `json:"migrating_to,omitempty" yaml:"migrating_to,omitempty"`
		} `json:"profile,omitempty" yaml:"profile,omitempty"`
		VifDetails struct {
			PortFilter bool `json:"port_filter,omitempty" yaml:"port_filter,omitempty"`
		} `json:"vif_details,omitempty" yaml:"vif_details,omitempty"`
	} `json:"binding,omitempty" yaml:"binding,omitempty"`
}

// FromJSON populates a BindingNotification using
// the data in the provided JSON.
func (msg *BindingNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing binding notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the BindingNotification
// in JSON format, pretty-printed or not.
func (msg *BindingNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling binding notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling binding notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("binding notification marshaled to JSON")
	return string(bytes), nil
}

// ToString converts the BindingNotification into its JSON one-liner representation.
func (msg *BindingNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
