package notification

// RBACPolicy is sent when a new tenant is being created.
type RBACPolicy struct {
	Base    `json:",inline" yaml:",inline"`
	Payload RBACPolicyPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type RBACPolicyPayload struct {
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
