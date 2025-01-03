package message

import (
	"encoding/json"
	"log/slog"
)

type TagNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          TagNotificationPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type TagNotificationPayload struct {
	ParentResource   string   `json:"parent_resource,omitempty" yaml:"parent_resource,omitempty"`
	ParentResourceID string   `json:"parent_resource_id,omitempty" yaml:"parent_resource_id,omitempty"`
	Tags             []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// FromJSON populates a TagNotification using
// the data in the provided JSON.
func (msg *TagNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing tag notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the TagNotification
// in JSON format, pretty-printed or not.
func (msg *TagNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling tag notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling tag notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("tag notification marshaled to JSON")
	return string(bytes), nil
}

// ToString converts the TagNotification into its JSON one-liner representation.
func (msg *TagNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
