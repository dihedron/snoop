package message

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"regexp"
	"sync"

	"github.com/dihedron/snoop/format"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// Notification is the common interface to all OpenStack notifications.
type Notification interface {
	// Info returns the most relevant information about the notifications,
	// which allows to identify its type and correlate it with other other
	// adjacent notifications in the same event sequence.
	Info() *NotificationInfo
	// ToJSON returns a JSON representation of the Notification, as a single
	// or in pretty-print format.
	ToJSON(pretty bool) (string, error)
	// FromJSON populates a Notification using the data in the provided JSON.
	FromJSON(data string) error
	// ToString returns a one-liner string representation of the Notification.
	ToString() string
	// Ack allows to acknowledge the original amqp091.Delivery if set inside the
	// BaseNotification.
	Ack(multiple bool) error

	GetAMQPDeliveryReference() *amqp091.Delivery
	SetAMQPDeliveryReference(delivery *amqp091.Delivery)
}

// NotificationInfo is a compact representation of the most relevant information
// in an OpenStack Notification.
type NotificationInfo struct {
	EventType       string
	UserID          string
	UserName        string
	ProjectID       string
	ProjectName     string
	RequestID       string
	GlobalRequestID string
}

// BaseNotification is the base set of information contained in all
// OpenStack events, both from Nova and from Neutron; it does not implement
// the Notification interface.
type BaseNotification struct {
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
	// mutex protects the delivery reference from concurrent access
	mutex sync.RWMutex
	// delivery is the underlying RabbitMQ Delivery
	delivery *amqp091.Delivery
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

func (msg *BaseNotification) Info() *NotificationInfo {
	return &NotificationInfo{
		EventType:       msg.EventType,
		UserID:          msg.ContextUserID,
		UserName:        msg.ContextUserName,
		ProjectID:       msg.ContextProjectID,
		ProjectName:     msg.ContextProjectName,
		RequestID:       msg.ContextRequestID,
		GlobalRequestID: msg.ContextGlobalRequestID,
	}
}

func (msg *BaseNotification) SetAMQPDeliveryReference(delivery *amqp091.Delivery) {
	if delivery != nil {
		msg.mutex.Lock()
		defer msg.mutex.Unlock()
		msg.delivery = delivery
	}
}

func (msg *BaseNotification) GetAMQPDeliveryReference() *amqp091.Delivery {
	msg.mutex.RLock()
	defer msg.mutex.RUnlock()
	return msg.delivery
}

// FromJSON populates a KeyPairNotification using
// the data in the provided JSON.
func (msg *BaseNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing keypair notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the BaseNotification
// in JSON format, pretty-printed or not.
func (msg *BaseNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling base notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling base notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("keypair notification marshaled to JSON")
	return string(bytes), nil
}

// ToString converts the BaseNotification into its JSON one-liner representation.
func (msg *BaseNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}

// Ack allows to acknowledge the amqp091.Delivery from which this notification
// originated, if the underlying reference is valid.
func (msg *BaseNotification) Ack(multiple bool) error {
	slog.Debug("acknowledging Base notification", "notification", msg, "type", fmt.Sprintf("%T", msg))
	if msg != nil {
		delivery := msg.GetAMQPDeliveryReference()
		if delivery != nil {
			slog.Debug("acknowledging message", "tag", msg.delivery.DeliveryTag)
			err := msg.delivery.Ack(multiple)
			if err == nil {
				msg.delivery = nil
			}
			return err
		} else {
			slog.Warn("no reference to original delivery!")
		}
	}
	return nil
}

// NewNotificationFromOslo parses an Oslo message and extracts an
// OpenStack notification if it is one of the supported event types;
// it may include the original amqp091.Delivery if available in the Oslo
// message and the includeDelivery flag is set; this allows to acknowledge
// the AQMP delivery through the notification.
func NewNotificationFromOslo(oslo *Oslo, includeDelivery bool) (Notification, error) {
	if oslo == nil {
		slog.Error("input must not be nil", "error", ErrInvalidInput)
		return nil, ErrInvalidInput
	}
	notification, err := NewNotificationFromJSON(oslo.Payload)
	if err == nil && notification != nil && includeDelivery {
		slog.Debug("acquiring reference to original delivery from Oslo message...")
		oslo.mutex.RLock()
		defer oslo.mutex.RUnlock()
		delivery := oslo.delivery
		if delivery != nil {
			slog.Debug("valid delivery reference acquired, adding to notification...")
			notification.SetAMQPDeliveryReference(delivery)
			slog.Debug("added reference to original delivery to notification", "tag", delivery.DeliveryTag)
		}
	}
	return notification, err
}

var (
	// EventTypePattern is the regular expression that extracts the event
	// type from an Oslo message BEFORE unescaping quotes.
	EventTypePattern = regexp.MustCompile(`\"event_type\":\s*\"([a-zA-Z0-9\._-]+)\"`)
)

// NewNotificationFromJSON parses an Oslo message and extracts an
// OpenStack notification if it is one of the supported event types.
func NewNotificationFromJSON(input string) (Notification, error) {
	// parse body
	tokens := EventTypePattern.FindStringSubmatch(input)
	//slog.Debug("regular expression applied, value: %q, tokens: %v", input, tokens)
	slog.Debug("regular expression applied", "tokens", tokens)

	if len(tokens) == 0 {
		slog.Error("failure finding event type token in Oslo message body", "error", ErrInvalidPayload)
		return nil, ErrInvalidPayload
	}
	var notification Notification
	switch tokens[1] {
	case
		"compute.instance.exists",
		"compute.instance.update",
		"compute.instance.create.start",
		"compute.instance.create.end",
		"compute.instance.create.error",
		"compute.instance.delete.end",
		"compute.instance.delete.start",
		//"compute.instance.delete.error",
		"compute.instance.reboot.start",
		"compute.instance.reboot.end",
		//"compute.instance.reboot.error",
		"compute.instance.power_on.start",
		"compute.instance.power_on.end",
		//"compute.instance.power_on.error",
		"compute.instance.power_off.start",
		"compute.instance.power_off.end",
		//"compute.instance.power_off.error",
		"compute.instance.pause.start",
		"compute.instance.pause.end",
		//"compute.instance.pause.error",
		"compute.instance.unpause.start",
		"compute.instance.unpause.end",
		//"compute.instance.unpause.error",
		"compute.instance.resume.start",
		"compute.instance.resume.end",
		//"compute.instance.resume.error",
		"compute.instance.suspend.start",
		"compute.instance.suspend.end",
		//"compute.instance.suspend.error",
		"compute.instance.shutdown.start",
		"compute.instance.shutdown.end",
		//"compute.instance.shutdown.error",
		"compute.instance.resize.prep.start",
		"compute.instance.resize.prep.end",
		//"compute.instance.resize.prep.error",
		"compute.instance.resize.confirm.start",
		"compute.instance.resize.confirm.end",
		//"compute.instance.resize.confirm.error",
		"compute.instance.finish_resize.start",
		"compute.instance.finish_resize.end",
		//"compute.instance.finish_resize.error",
		"compute.instance.resize.start",
		"compute.instance.resize.end",
		//"compute.instance.resize.error",
		"compute.instance.volume.attach",
		"compute.instance.volume.detach",
		//"compute,instance.volume.detach.error",
		"compute.instance.evacuate",
		//"compute.instance.evacuate.error",
		"compute.instance.rebuild.scheduled",
		"compute.instance.rebuild.error",
		"compute.instance.shelve_offload.start",
		"compute.instance.shelve_offload.end",
		"compute.instance.unshelve.start",
		"compute.instance.unshelve.end",
		"compute.instance.live_migration.pre.start",
		"compute.instance.live_migration.pre.end",
		"compute.instance.live_migration.post.dest.start",
		"compute.instance.live_migration.post.dest.end",
		"compute.instance.live_migration._post.start",
		"compute.instance.live_migration._post.end",
		"compute.instance.live_migration.rollback.dest.start",
		"compute.instance.live_migration.rollback.dest.end",
		"compute.instance.live_migration._rollback.start",
		"compute.instance.live_migration._rollback.end":
		slog.Debug("parsing compute message", "event type", tokens[1])
		notification = &ComputeInstanceNotification{}
	case
		"rebuild_instance",
		"stop_instance",
		"resize_instance",
		"get_console_output",
		"get_instance_diagnostics",
		"create_key_pair",
		"attach_interface",
		"detach_interface",
		"attach_volume",
		"detach_volume",
		"pre_live_migration":
		slog.Debug("parsing exception message", "event type", tokens[1])
		notification = &ExceptionNotification{}
	// case "compute.libvirt.error":
	// 	slog.Debug("parsing libvirt message", "event type", tokens[1])
	// 	notification = &ComputeLibvirtNotification{}
	case
		"keypair.create.start",
		"keypair.create.end",
		//"keypair.create.error",
		"keypair.delete.start",
		"keypair.delete.end",
		//"keypair.delete.error",
		"keypair.import.start",
		"keypair.import.end":
		//"keypair.import.error"
		slog.Debug("parsing keypair message", "event type", tokens[1])
		notification = &KeyPairNotification{}
	case
		"port.create.start",
		"port.create.end",
		//"port.create.error",
		"port.update.start",
		"port.update.end",
		//"port.update.error",
		"port.delete.start",
		//"port.delete.error",
		"port.delete.end":
		slog.Debug("parsing port message", "event type", tokens[1])
		notification = &PortNotification{}
	case
		"rbac_policy.create.start",
		"rbac_policy.create.end",
		"rbac_policy.delete.start",
		"rbac_policy.delete.end":
		slog.Debug("parsing RBAC policy message", "event type", tokens[1])
		notification = &RBACPolicyNotification{}
	case
		"compute_task.rebuild_server",
		"compute_task.build_instances",
		"scheduler.select_destinations.start",
		"scheduler.select_destinations.end",
		"compute.libvirt.error":
		slog.Debug("parsing compute task rebuild server message", "event type", tokens[1])
		notification = &ComputeTaskNotification{}
	case
		"security_group.create.start",
		"security_group.create.end",
		"security_group.update.start",
		"security_group.update.end",
		"security_group.delete.start",
		"security_group.delete.end":
		slog.Debug("parsing security group message", "event type", tokens[1])
		notification = &SecurityGroupNotification{}
	case
		"security_group_rule.create.start",
		"security_group_rule.create.end",
		"security_group_rule.delete.start",
		"security_group_rule.delete.end":
		slog.Debug("parsing security group rule message", "event type", tokens[1])
		notification = &SecurityGroupRuleNotification{}
	case
		"tag.create.start",
		"tag.create.end",
		"tag.update.start",
		"tag.update.end",
		"tag.delete.start",
		"tag.delete.end",
		"tag.delete_all.start",
		"tag.delete_all.end":
		slog.Debug("parsing tag message", "event type", tokens[1])
		notification = &TagNotification{}
	case
		"binding.create.start",
		"binding.create.end",
		"binding.delete.start",
		"binding.delete.end":
		slog.Debug("parsing binding message", "event type", tokens[1])
		notification = &BindingNotification{}
	case
		"identity.authenticate",
		"identity.user.created",
		"identity.user.updated",
		"identity.user.deleted",
		"identity.project.created",
		"identity.project.updated",
		"identity.project.deleted",
		"identity.application_credential.created",
		"identity.application_credential.deleted",
		"identity.role_assignment.created",
		"identity.role_assignment.deleted",
		"identity.endpoint.updated":
		slog.Debug("parsing identity message", "event type", tokens[1])
		notification = &IdentityNotification{}
	default:
		slog.Debug("unsupported event type", "event type", tokens[1])
		format.WriteToFileAsJSON(".", tokens[1]+"-*.json", input)
		return nil, fmt.Errorf("unsupported event type: %T", tokens[1])
	}

	// unescape the payload to make it a valid JSON
	// payload := strings.ReplaceAll(input, `\"`, `"`)
	// this is to handle embedded Python stack traces
	// payload = strings.ReplaceAll(payload, `''`, `\"`)
	payload := input

	// debugging dumps out messages to file
	const debug bool = false
	if debug {
		switch notification.(type) {
		case *ComputeTaskNotification:
			format.WriteToFileAsJSON(".", "compute-task-*.json", input)
		case *ExceptionNotification:
			format.WriteToFileAsJSON(".", "exception-*.json", input)
		}
	}

	if err := notification.FromJSON(payload); err != nil {
		slog.Error("failure parsing notification of type", "notification type", tokens[1], "error", err)
		return nil, err
	}
	switch notification := notification.(type) {
	case *ComputeInstanceNotification:
		slog.Debug("compute instance notification parsed", "event id", notification.UniqueID)
	case *ExceptionNotification:
		slog.Debug("exception notification parsed", "event id", notification.UniqueID)
	case *IdentityNotification:
		slog.Debug("identity notification parsed", "event id", notification.UniqueID)
	case *PortNotification:
		slog.Debug("port notification parsed", "event id", notification.UniqueID)
	case *ComputeTaskNotification:
		slog.Debug("compute task notification parsed", "event id", notification.UniqueID)
	case *RBACPolicyNotification:
		slog.Debug("RBAC policy notification parsed", "event id", notification.UniqueID)
	case *KeyPairNotification:
		slog.Debug("keypair notification parsed", "event id", notification.UniqueID)
	case *SecurityGroupNotification:
		slog.Debug("security group notification parsed", "event id", notification.UniqueID)
	case *SecurityGroupRuleNotification:
		slog.Debug("security group rule notification parsed", "event id", notification.UniqueID)
	case *TagNotification:
		slog.Debug("tag notification parsed", "event id", notification.UniqueID)
	case *BindingNotification:
		slog.Debug("binding notification parsed", "event id", notification.UniqueID)
	default:
		slog.Error("unsupported notification type", "type", fmt.Sprintf("%T", notification))
		return nil, fmt.Errorf("unsupported notification type: %T", notification)
	}
	slog.Debug("notification parsed")
	return notification, nil
}
