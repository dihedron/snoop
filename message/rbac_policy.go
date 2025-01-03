package message

import (
	"encoding/json"
	"log/slog"
)

// RBACPolicyNotification is sent when a new tenant is being created.
type RBACPolicyNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          RBACPolicyNotificationPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type RBACPolicyNotificationPayload struct {
	RBACPolicyID string `json:"rbac_policy_id,omitempty"` // ???
	RBACPolicy   struct {
		ID           string `json:"id,omitempty" yaml:"id,omitempty"`
		ProjectID    string `json:"project_id,omitempty" yaml:"project_id,omitempty"`
		TenantID     string `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
		Action       string `json:"action,omitempty" yaml:"action,omitempty"`
		ObjectID     string `json:"object_id,omitempty" yaml:"object_id,omitempty"`
		ObjectType   string `json:"object_type,omitempty" yaml:"object_type,omitempty"`
		TargetTenant string `json:"target_tenant,omitempty" yaml:"target_tenant,omitempty"`
	} `json:"rbac_policy,omitempty" yaml:"rbac_policy,omitempty"`
}

// type RBACPolicy struct {
// 	ID           string `json:"id,omitempty" yaml:"id,omitempty"`
// 	ProjectID    string `json:"project_id,omitempty" yaml:"project_id,omitempty"`
// 	TenantID     string `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
// 	Action       string `json:"action,omitempty" yaml:"action,omitempty"`
// 	ObjectID     string `json:"object_id,omitempty" yaml:"object_id,omitempty"`
// 	ObjectType   string `json:"object_type,omitempty" yaml:"object_type,omitempty"`
// 	TargetTenant string `json:"target_tenant,omitempty" yaml:"target_tenant,omitempty"`
// }

// FromJSON populates a RBACPolicyNotification using
// the data in the provided JSON.
func (msg *RBACPolicyNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing RBAC policy notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the RBACPolicyNotification
// in JSON format, pretty-printed or not.
func (msg *RBACPolicyNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling RBAC policy notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling RBAC policy notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("RBAC policy notification marshaled to JSON")
	return string(bytes), nil

}

// ToString converts the RBACPolicyNotification into a string containing its
// JSON one-liner representation.
func (msg *RBACPolicyNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
