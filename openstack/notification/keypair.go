package notification

// KeyPair notifications are received when key-pairs are
// being created, added, removed.
type KeyPair struct {
	Base    `json:",inline" yaml:",inline"`
	Payload KeyPairPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type KeyPairPayload struct {
	TenantID string `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
	UserID   string `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	KeyName  string `json:"key_name,omitempty" yaml:"key_name,omitempty"`
}
