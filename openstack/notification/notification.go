package notification

import (
	"log/slog"

	"github.com/goccy/go-json"

	"github.com/dihedron/snoop/format"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// Notification is the common interface to all OpenStack notifications.
type Notification interface {
	// Summary returns the most relevant information about the notification,
	// which allows to identify its type and correlate it with adjacent
	// notifications in the same event sequence.
	Summary() *Summary
}

// String returns a string representation of Base notification as a JSON one-liner.
func ToString[N Notification](n N) string {
	return format.ToJSON(n)
}

// FromJSON populates a Notification using the data in the provided JSON.
func FromJSON[N Notification](notification N, data string) error {
	if err := json.Unmarshal([]byte(data), notification); err != nil {
		slog.Error("failure parsing exception notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the ExceptionNotification
// in JSON format, pretty-printed or not.
func ToJSON[N Notification](notification N, pretty bool) string {
	if pretty {
		return format.ToPrettyJSON(notification)
	} else {
		return format.ToJSON(notification)
	}
}

// Summary is a compact representation of the most relevant information
// in an OpenStack Notification.
type Summary struct {
	EventType       string
	UserID          string
	UserName        string
	ProjectID       string
	ProjectName     string
	RequestID       string
	GlobalRequestID string
}

// Commons is the base set of information contained in all
// OpenStack events, both from Nova and from Neutron; it does not
// implement the Notification interface.
type Base struct {
	MessageID              string           `json:"message_id,omitempty" yaml:"message_id,omitempty"`
	PublisherID            string           `json:"publisher_id,omitempty" yaml:"publisher_id,omitempty"`
	EventType              string           `json:"event_type,omitempty" yaml:"event_type,omitempty"`
	Priority               string           `json:"priority,omitempty" yaml:"priority,omitempty"`
	Timestamp              string           `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
	UniqueID               string           `json:"_unique_id,omitempty" yaml:"_unique_id,omitempty"`
	ContextUser            string           `json:"_context_user,omitempty" yaml:"_context_user,omitempty"`
	ContextTenant          string           `json:"_context_tenant,omitempty" yaml:"_context_tenant,omitempty"`
	ContextSystemScope     string           `json:"_context_system_scope,omitempty" yaml:"_context_system_scope,omitempty"`
	ContextProject         string           `json:"_context_project,omitempty" yaml:"_context_project,omitempty"`
	ContextDomain          string           `json:"_context_domain,omitempty" yaml:"_context_domain,omitempty"`
	ContextUserDomain      string           `json:"_context_user_domain,omitempty" yaml:"_context_user_domain,omitempty"`
	ContextProjectDomain   string           `json:"_context_project_domain,omitempty" yaml:"_context_project_domain,omitempty"`
	ContextIsAdmin         bool             `json:"_context_is_admin,omitempty" yaml:"_context_is_admin,omitempty"`
	ContextReadOnly        bool             `json:"_context_read_only,omitempty" yaml:"_context_read_only,omitempty"`
	ContextShowDeleted     bool             `json:"_context_show_deleted,omitempty" yaml:"_context_show_deleted,omitempty"`
	ContextAuthToken       string           `json:"_context_auth_token,omitempty" yaml:"_context_auth_token,omitempty"`
	ContextRequestID       string           `json:"_context_request_id,omitempty" yaml:"_context_request_id,omitempty"` // this
	ContextGlobalRequestID string           `json:"_context_global_request_id,omitempty" yaml:"_context_global_request_id,omitempty"`
	ContextResourceUUID    string           `json:"_context_resource_uuid,omitempty" yaml:"_context_resource_uuid,omitempty"`
	ContextRoles           []string         `json:"_context_roles,omitempty" yaml:"_context_roles,omitempty"`
	ContextUserIdentity    string           `json:"_context_user_identity,omitempty" yaml:"_context_user_identity,omitempty"`
	ContextIsAdminProject  bool             `json:"_context_is_admin_project,omitempty" yaml:"_context_is_admin_project,omitempty"`
	ContextUserID          string           `json:"_context_user_id,omitempty" yaml:"_context_user_id,omitempty"`
	ContextReadDeleted     string           `json:"_context_read_deleted" yaml:"_context_read_deleted"`
	ContextRemoteAddress   string           `json:"_context_remote_address" yaml:"_context_remote_address"`
	ContextQuotaClass      string           `json:"_context_quota_class" yaml:"_context_quota_class"`
	ContextTenantID        string           `json:"_context_tenant_id,omitempty" yaml:"_context_tenant_id,omitempty"`
	ContextProjectID       string           `json:"_context_project_id,omitempty" yaml:"_context_project_id,omitempty"`
	ContextTimestamp       string           `json:"_context_timestamp,omitempty" yaml:"_context_timestamp,omitempty"`
	ContextTenantName      string           `json:"_context_tenant_name,omitempty" yaml:"_context_tenant_name,omitempty"`
	ContextProjectName     string           `json:"_context_project_name,omitempty" yaml:"_context_project_name,omitempty"`
	ContextUserName        string           `json:"_context_user_name,omitempty" yaml:"_context_user_name,omitempty"`
	ContextServiceCatalog  []ServiceCatalog `json:"_context_service_catalog,omitempty" yaml:"_context_service_catalog,omitempty"`
	// backref is a reference to the underlying RabbitMQ delivery
	backref *amqp091.Delivery
}

type ServiceCatalog struct {
	Type      string     `json:"type,omitempty" yaml:"type,omitempty"`
	Name      string     `json:"name,omitempty" yaml:"name,omitempty"`
	Endpoints []Endpoint `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
}

type Endpoint struct {
	Region      string `json:"region,omitempty" yaml:"region,omitempty"`
	InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
	PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
	AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
}

// Summary returns a small subset of an OpenStack notification/event
// related information.
func (b *Base) Summary() *Summary {
	return &Summary{
		EventType:       b.EventType,
		UserID:          b.ContextUserID,
		UserName:        b.ContextUserName,
		ProjectID:       b.ContextProjectID,
		ProjectName:     b.ContextProjectName,
		RequestID:       b.ContextRequestID,
		GlobalRequestID: b.ContextGlobalRequestID,
	}
}

func (b *Base) SetBackRef(delivery *amqp091.Delivery) {
	b.backref = delivery
}

func (b *Base) BackRef() *amqp091.Delivery {
	return b.backref
}

// Ack allows to acknowledge the Notification's underlying amqp091.Delivery, if set.
func (b *Base) Ack(multiple bool) error {
	slog.Debug("acknowledging Notification message...", "type", format.TypeAsString(b))
	if b.backref != nil {
		slog.Debug("acknowledging message", "correlation id", b.backref.CorrelationId)
		if err := b.backref.Ack(multiple); err != nil {
			slog.Error("error acnowledging message", "correlation id", b.backref.CorrelationId, "error", err)
			return err
		}
		b.backref = nil
	}
	return nil
}
