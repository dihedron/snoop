package message

import (
	"encoding/json"
	"log/slog"
)

// KeyPairNotifications are received when key-pairs are
// being created, added, removed.
type KeyPairNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          KeyPairNotificationPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type KeyPairNotificationPayload struct {
	TenantID string `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
	UserID   string `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	KeyName  string `json:"key_name,omitempty" yaml:"key_name,omitempty"`
}

// FromJSON populates a KeyPairNotification using
// the data in the provided JSON.
func (msg *KeyPairNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing keypair notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the KeyPairNotification
// in JSON format, pretty-printed or not.
func (msg *KeyPairNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling keypair notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling keypair notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("keypair notification marshaled to JSON")
	return string(bytes), nil
}

// ToString converts the KeyPairNotification into its JSON one-liner representation.
func (msg *KeyPairNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
