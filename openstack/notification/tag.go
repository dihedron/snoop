package notification

type Tag struct {
	Base    `json:",inline" yaml:",inline"`
	Payload TagPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type TagPayload struct {
	ParentResource   string   `json:"parent_resource,omitempty" yaml:"parent_resource,omitempty"`
	ParentResourceID string   `json:"parent_resource_id,omitempty" yaml:"parent_resource_id,omitempty"`
	Tags             []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
