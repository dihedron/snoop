package notification

import (
	"errors"
	"fmt"
	"log/slog"
	"regexp"

	"github.com/goccy/go-json"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/openstack/oslo"
)

// NewNotificationFromOslo parses an Oslo message and extracts an
// OpenStack notification if it is one of the supported event types;
// it may include the original amqp091.Delivery if available in the Oslo
// message and the includeDelivery flag is set; this allows to acknowledge
// the AQMP delivery through the notification.
func OsloToNotification(includeBackRef bool) func(*oslo.Oslo) (Notification, error) {
	return func(oslo *oslo.Oslo) (Notification, error) {
		notification, err := JSONToNotification()(oslo.Payload)
		if err == nil && oslo != nil && includeBackRef && oslo.BackRef() != nil {
			slog.Debug("adding back-reference to original AMQP delivery", "reference", oslo.BackRef().DeliveryTag)
			if base, ok := notification.(*Base); ok {
				slog.Debug("setting back reference")
				base.backref = oslo.BackRef()
			}
		}
		return notification, err
	}
}

// JSONToNotification parses an Oslo message's JSON payload and
// extracts an OpenStack notification if it is one of the supported
// types.
func JSONToNotification() func(string) (Notification, error) {
	return func(input string) (Notification, error) {
		// in order to detect the kind of structure into which we will
		// parse the input JSON, it is necessary to look up a pattern
		// matching the type in the JSON string< once we have a match,
		// we can switch on the string and use the proper Notification
		// concrtete type; the pattees extracts the event type from an
		// Oslo message BEFORE unescaping quotes.
		pattern := regexp.MustCompile(`\"event_type\":\s*\"([a-zA-Z0-9\._-]+)\"`)

		tokens := pattern.FindStringSubmatch(input)
		//slog.Debug("regular expression applied, value: %q, tokens: %v", input, tokens)
		slog.Debug("regular expression applied", "tokens", tokens)

		if len(tokens) == 0 {
			slog.Error("failure finding event type token in Oslo message body")
			return nil, errors.New("invalid payload")
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
			notification = &ComputeInstance{}
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
			notification = &Exception{}
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
			notification = &KeyPair{}
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
			notification = &Port{}
		case
			"rbac_policy.create.start",
			"rbac_policy.create.end",
			"rbac_policy.delete.start",
			"rbac_policy.delete.end":
			slog.Debug("parsing RBAC policy message", "event type", tokens[1])
			notification = &RBACPolicy{}
		case
			"compute_task.rebuild_server",
			"compute_task.build_instances",
			"scheduler.select_destinations.start",
			"scheduler.select_destinations.end",
			"compute.libvirt.error":
			slog.Debug("parsing compute task rebuild server message", "event type", tokens[1])
			notification = &ComputeTask{}
		case
			"security_group.create.start",
			"security_group.create.end",
			"security_group.update.start",
			"security_group.update.end",
			"security_group.delete.start",
			"security_group.delete.end":
			slog.Debug("parsing security group message", "event type", tokens[1])
			notification = &SecurityGroup{}
		case
			"security_group_rule.create.start",
			"security_group_rule.create.end",
			"security_group_rule.delete.start",
			"security_group_rule.delete.end":
			slog.Debug("parsing security group rule message", "event type", tokens[1])
			notification = &SecurityGroupRule{}
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
			notification = &Tag{}
		case
			"binding.create.start",
			"binding.create.end",
			"binding.delete.start",
			"binding.delete.end":
			slog.Debug("parsing binding message", "event type", tokens[1])
			notification = &Binding{}
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
			notification = &Identity{}
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
			case *ComputeTask:
				format.WriteToFileAsJSON(".", "compute-task-*.json", input)
			case *Exception:
				format.WriteToFileAsJSON(".", "exception-*.json", input)
			}
		}

		if err := json.Unmarshal([]byte(payload), notification); err != nil {
			slog.Error("failure parsing exception notification from JSON", "notification type", tokens[1], "error", err)
			return nil, err
		}

		switch notification := notification.(type) {
		case *ComputeInstance:
			slog.Debug("compute instance notification parsed", "event id", notification.UniqueID)
		case *Exception:
			slog.Debug("exception notification parsed", "event id", notification.UniqueID)
		case *Identity:
			slog.Debug("identity notification parsed", "event id", notification.UniqueID)
		case *Port:
			slog.Debug("port notification parsed", "event id", notification.UniqueID)
		case *ComputeTask:
			slog.Debug("compute task notification parsed", "event id", notification.UniqueID)
		case *RBACPolicy:
			slog.Debug("RBAC policy notification parsed", "event id", notification.UniqueID)
		case *KeyPair:
			slog.Debug("keypair notification parsed", "event id", notification.UniqueID)
		case *SecurityGroup:
			slog.Debug("security group notification parsed", "event id", notification.UniqueID)
		case *SecurityGroupRule:
			slog.Debug("security group rule notification parsed", "event id", notification.UniqueID)
		case *Tag:
			slog.Debug("tag notification parsed", "event id", notification.UniqueID)
		case *Binding:
			slog.Debug("binding notification parsed", "event id", notification.UniqueID)
		default:
			slog.Error("unsupported notification type", "type", fmt.Sprintf("%T", notification))
			return nil, fmt.Errorf("unsupported notification type: %T", notification)
		}
		slog.Debug("notification parsed")
		return notification, nil
	}
}
