package message

import (
	"encoding/json"
	"log/slog"
	"time"
)

// PortNotification is the container for all neutron's port related notifications.
type PortNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          PortNotificationPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type PortNotificationPayload struct {
	Port Port   `json:"port,omitempty" yaml:"port,omitempty"`
	ID   string `json:"id,omitempty" yaml:"id,omitempty"`
}

type Port struct {
	ID           string `json:"id,omitempty" yaml:"id,omitempty"`
	Name         string `json:"name,omitempty" yaml:"name,omitempty"`
	NetworkID    string `json:"network_id,omitempty" yaml:"network_id,omitempty"`
	TenantID     string `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
	MacAddress   string `json:"mac_address,omitempty" yaml:"mac_address,omitempty"`
	AdminStateUp bool   `json:"admin_state_up,omitempty" yaml:"admin_state_up,omitempty"`
	Status       string `json:"status,omitempty" yaml:"status,omitempty"`
	DeviceID     string `json:"device_id,omitempty" yaml:"device_id,omitempty"`
	DeviceOwner  string `json:"device_owner,omitempty" yaml:"device_owner,omitempty"`
	FixedIps     []struct {
		SubnetID  string `json:"subnet_id,omitempty" yaml:"subnet_id,omitempty"`
		IPAddress string `json:"ip_address,omitempty" yaml:"ip_address,omitempty"`
	} `json:"fixed_ips,omitempty" yaml:"fixed_ips,omitempty"`
	ProjectID           string   `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	QosPolicyID         string   `json:"qos_policy_id,omitempty" yaml:"qos_policy_id,omitempty"`
	PortSecurityEnabled bool     `json:"port_security_enabled,omitempty" yaml:"port_security_enabled,omitempty"`
	SecurityGroups      []string `json:"security_groups,omitempty" yaml:"security_groups,omitempty"`
	BindingVnicType     string   `json:"binding:vnic_type,omitempty" yaml:"binding:vnic_type,omitempty"`
	AllowedAddressPairs []string `json:"allowed_address_pairs,omitempty" yaml:"allowed_address_pairs,omitempty"`
	ExtraDhcpOpts       []string `json:"extra_dhcp_opts,omitempty" yaml:"extra_dhcp_opts,omitempty"`
	Description         string   `json:"description,omitempty" yaml:"description,omitempty"`
	BindingProfile      struct {
	} `json:"binding:profile,omitempty" yaml:"binding:profile,omitempty"`
	BindingHostID     string `json:"binding:host_id,omitempty" yaml:"binding:host_id,omitempty"`
	BindingVifType    string `json:"binding:vif_type,omitempty" yaml:"binding:vif_type,omitempty"`
	BindingVifDetails struct {
		PortFilter bool `json:"port_filter,omitempty" yaml:"port_filter,omitempty"`
	} `json:"binding:vif_details,omitempty" yaml:"binding:vif_details,omitempty"`
	DNSName       string `json:"dns_name,omitempty" yaml:"dns_name,omitempty"`
	DNSAssignment []struct {
		IPAddress string `json:"ip_address,omitempty" yaml:"ip_address,omitempty"`
		Hostname  string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
		Fqdn      string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	} `json:"dns_assignment,omitempty" yaml:"dns_assignment,omitempty"`
	ResourceRequest string    `json:"resource_request,omitempty" yaml:"resource_request,omitempty"`
	IPAllocation    string    `json:"ip_allocation,omitempty" yaml:"ip_allocation,omitempty"`
	Tags            []string  `json:"tags,omitempty" yaml:"tags,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	RevisionNumber  int       `json:"revision_number,omitempty" yaml:"revision_number,omitempty"`
}

// type PortFixedIP struct {
// 	SubnetID  string `json:"subnet_id,omitempty" yaml:"subnet_id,omitempty"`
// 	IPAddress string `json:"ip_address,omitempty" yaml:"ip_address,omitempty"`
// }

// type PortBindingProfile struct {
// }

// type PortBindingVifDetails struct {
// 	PortFilter bool `json:"port_filter,omitempty" yaml:"port_filter,omitempty"`
// }

// type PortDNSAssignment struct {
// 	IPAddress string `json:"ip_address,omitempty" yaml:"ip_address,omitempty"`
// 	Hostname  string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
// 	Fqdn      string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
// }

// FromJSON populates a PortNotification using
// the data in the provided JSON.
func (msg *PortNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing port notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the PortNotification
// in JSON format, pretty-printed or not.
func (msg *PortNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling port notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling port notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("port notification marshaled to JSON")
	return string(bytes), nil

}

// ToString converts the PortNotification into a string containing its
// JSON one-liner representation.
func (msg *PortNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
