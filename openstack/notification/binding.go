package notification

type Binding struct {
	Base    `json:",inline" yaml:",inline"`
	Payload BindingPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type BindingPayload struct {
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
