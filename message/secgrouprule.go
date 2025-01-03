package message

import (
	"encoding/json"
	"log/slog"
	"time"
)

type SecurityGroupRuleNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          SecurityGroupRuleNotificationPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type SecurityGroupRuleNotificationPayload struct {
	SecurityGroupRuleID string `json:"security_group_rule_id,omitempty" yaml:"security_group_rule_id,omitempty"`
	SecurityGroupRules  []struct {
		ID              string    `json:"id,omitempty" yaml:"id,omitempty"`
		TenantID        string    `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
		SecurityGroupID string    `json:"security_group_id,omitempty" yaml:"security_group_id,omitempty"`
		Ethertype       string    `json:"ethertype,omitempty" yaml:"ethertype,omitempty"`
		Direction       string    `json:"direction,omitempty" yaml:"direction,omitempty"`
		Protocol        string    `json:"protocol,omitempty" yaml:"protocol,omitempty"`
		PortRangeMin    int       `json:"port_range_min,omitempty" yaml:"port_range_min,omitempty"`
		PortRangeMax    int       `json:"port_range_max,omitempty" yaml:"port_range_max,omitempty"`
		RemoteIPPrefix  string    `json:"remote_ip_prefix,omitempty" yaml:"remote_ip_prefix,omitempty"`
		RemoteGroupID   string    `json:"remote_group_id,omitempty" yaml:"remote_group_id,omitempty"`
		Description     string    `json:"description,omitempty" yaml:"description,omitempty"`
		CreatedAt       time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
		UpdatedAt       time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
		RevisionNumber  int       `json:"revision_number,omitempty" yaml:"revision_number,omitempty"`
		ProjectID       string    `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	} `json:"security_group_rules,omitempty" yaml:"security_group_rules,omitempty"`
}

// type SecurityGroupRule struct {
// 	ID              string    `json:"id,omitempty" yaml:"id,omitempty"`
// 	TenantID        string    `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
// 	SecurityGroupID string    `json:"security_group_id,omitempty" yaml:"security_group_id,omitempty"`
// 	Ethertype       string    `json:"ethertype,omitempty" yaml:"ethertype,omitempty"`
// 	Direction       string    `json:"direction,omitempty" yaml:"direction,omitempty"`
// 	Protocol        string    `json:"protocol,omitempty" yaml:"protocol,omitempty"`
// 	PortRangeMin    int       `json:"port_range_min,omitempty" yaml:"port_range_min,omitempty"`
// 	PortRangeMax    int       `json:"port_range_max,omitempty" yaml:"port_range_max,omitempty"`
// 	RemoteIPPrefix  string    `json:"remote_ip_prefix,omitempty" yaml:"remote_ip_prefix,omitempty"`
// 	RemoteGroupID   string    `json:"remote_group_id,omitempty" yaml:"remote_group_id,omitempty"`
// 	Description     string    `json:"description,omitempty" yaml:"description,omitempty"`
// 	CreatedAt       time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
// 	UpdatedAt       time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
// 	RevisionNumber  int       `json:"revision_number,omitempty" yaml:"revision_number,omitempty"`
// 	ProjectID       string    `json:"project_id,omitempty" yaml:"project_id,omitempty"`
// }

// FromJSON populates a SecurityGroupRuleNotification using
// the data in the provided JSON.
func (msg *SecurityGroupRuleNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing security group rule notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the SecurityGroupRuleNotification
// in JSON format, pretty-printed or not.
func (msg *SecurityGroupRuleNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling security group rule notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling security group rule notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("security group rule notification marshaled to JSON")
	return string(bytes), nil
}

// ToString converts the SecurityGroupRuleNotification into its JSON one-liner representation.
func (msg *SecurityGroupRuleNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
