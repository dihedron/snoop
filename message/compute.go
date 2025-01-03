package message

import (
	"encoding/json"
	"log/slog"
)

// ComputeInstanceNotification is the notification for all instance-related events,
// such as those pertaining to existence, creation, update, deletion.
type ComputeInstanceNotification struct {
	BaseNotification `json:",inline" yaml:",inline"`
	Payload          ComputeInstanceNotificationPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type ComputeInstanceNotificationPayload struct {
	TenantID         string      `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty" diff:"tenant_id"`
	UserID           string      `json:"user_id,omitempty" yaml:"user_id,omitempty" diff:"user_id"`
	InstanceID       string      `json:"instance_id,omitempty" yaml:"instance_id,omitempty" diff:"instance_id"`
	DisplayName      string      `json:"display_name,omitempty" yaml:"display_name,omitempty" diff:"display_name"`
	ReservationID    string      `json:"reservation_id,omitempty" yaml:"reservation_id,omitempty" diff:"reservation_id"`
	Hostname         string      `json:"hostname,omitempty" yaml:"hostname,omitempty" diff:"hostname"`
	InstanceType     string      `json:"instance_type,omitempty" yaml:"instance_type,omitempty" diff:"instance_type"`
	InstanceTypeID   int         `json:"instance_type_id,omitempty" yaml:"instance_type_id,omitempty"`
	InstanceFlavorID string      `json:"instance_flavor_id,omitempty" yaml:"instance_flavor_id,omitempty"`
	Architecture     string      `json:"architecture,omitempty" yaml:"architecture,omitempty"`
	MemoryMb         int         `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
	DiskGb           int         `json:"disk_gb,omitempty" yaml:"disk_gb,omitempty"`
	VCPUs            int         `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
	RootGb           int         `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
	EphemeralGb      int         `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
	Host             string      `json:"host,omitempty" yaml:"host,omitempty"`
	Node             string      `json:"node,omitempty" yaml:"node,omitempty"`
	AvailabilityZone string      `json:"availability_zone,omitempty" yaml:"availability_zone,omitempty"`
	CellName         string      `json:"cell_name,omitempty" yaml:"cell_name,omitempty"`
	CreatedAt        string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	TerminatedAt     string      `json:"terminated_at,omitempty" yaml:"terminated_at,omitempty"`
	DeletedAt        string      `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
	LaunchedAt       string      `json:"launched_at,omitempty" yaml:"launched_at,omitempty"`
	ImageRefURL      string      `json:"image_ref_url,omitempty" yaml:"image_ref_url,omitempty"`
	OSType           string      `json:"os_type,omitempty" yaml:"os_type,omitempty"`
	KernelID         string      `json:"kernel_id,omitempty" yaml:"kernel_id,omitempty"`
	RamdiskID        string      `json:"ramdisk_id,omitempty" yaml:"ramdisk_id,omitempty"`
	State            string      `json:"state,omitempty" yaml:"state,omitempty"`
	StateDescription string      `json:"state_description,omitempty" yaml:"state_description,omitempty"`
	Progress         interface{} `json:"progress,omitempty" yaml:"progress,omitempty"`
	AccessIPV4       string      `json:"access_ip_v4,omitempty" yaml:"access_ip_v4,omitempty"`
	AccessIPV6       string      `json:"access_ip_v6,omitempty" yaml:"access_ip_v6,omitempty"`
	FixedIPs         []struct {
		Address string `json:"address,omitempty" yaml:"address,omitempty"`
		Type    string `json:"type,omitempty" yaml:"type,omitempty"`
		Version int    `json:"version,omitempty" yaml:"version,omitempty"`
		Meta    struct {
		} `json:"meta,omitempty" yaml:"meta,omitempty"`
		FloatingIPs []string `json:"floating_ips,omitempty" yaml:"floating_ips,omitempty"`
		Label       string   `json:"label,omitempty" yaml:"label,omitempty"`
		VifMac      string   `json:"vif_mac,omitempty" yaml:"vif_mac,omitempty"`
	} `json:"fixed_ips,omitempty" yaml:"fixed_ips,omitempty"`
	ImageMeta struct {
		Architecture                  string `json:"architecture,omitempty" yaml:"architecture,omitempty"`
		Description                   string `json:"description,omitempty" yaml:"description,omitempty"`
		CommitSha                     string `json:"commit_sha,omitempty" yaml:"commit_sha,omitempty"`
		HwDiskBus                     string `json:"hw_disk_bus,omitempty" yaml:"hw_disk_bus,omitempty"`
		HwQEMUGuestAgent              string `json:"hw_qemu_guest_agent,omitempty" yaml:"hw_qemu_guest_agent,omitempty"`
		HwRngModel                    string `json:"hw_rng_model,omitempty" yaml:"hw_rng_model,omitempty"`
		HwScsiModel                   string `json:"hw_scsi_model,omitempty" yaml:"hw_scsi_model,omitempty"`
		ImageType                     string `json:"image_type,omitempty" yaml:"image_type,omitempty"`
		OsDistro                      string `json:"os_distro,omitempty" yaml:"os_distro,omitempty"`
		OwnerSpecifiedOpenstackMd5    string `json:"owner_specified.openstack.md5,omitempty" yaml:"owner_specified.openstack.md5,omitempty"`
		OwnerSpecifiedOpenstackObject string `json:"owner_specified.openstack.object,omitempty" yaml:"owner_specified.openstack.object,omitempty"`
		OwnerSpecifiedOpenstackSha256 string `json:"owner_specified.openstack.sha256,omitempty" yaml:"owner_specified.openstack.sha256,omitempty"`
		MinRAM                        string `json:"min_ram,omitempty" yaml:"min_ram,omitempty"`
		MinDisk                       string `json:"min_disk,omitempty" yaml:"min_disk,omitempty"`
		DiskFormat                    string `json:"disk_format,omitempty" yaml:"disk_format,omitempty"`
		ContainerFormat               string `json:"container_format,omitempty" yaml:"container_format,omitempty"`
		BaseImageRef                  string `json:"base_image_ref,omitempty" yaml:"base_image_ref,omitempty"`
	} `json:"image_meta,omitempty" yaml:"image_meta,omitempty"`
	// Metadata struct {
	// 	Backup      string `json:"BACKUP,omitempty" yaml:"BACKUP,omitempty"`
	// 	ClusterName string `json:"cluster_name,omitempty" yaml:"cluster_name,omitempty"`
	// 	ImageName   string `json:"image_name,omitempty" yaml:"image_name,omitempty"`
	// 	HostID      string `json:"host_id,omitempty" yaml:"host_id,omitempty"`
	// 	VMRole      string `json:"vmrole,omitempty" yaml:"vmrole,omitempty"`
	// } `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Metadata             map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Message              string            `json:"message,omitempty" yaml:"message,omitempty"`
	Exception            string            `json:"exception,omitempty" yaml:"exception,omitempty"`
	Code                 int               `json:"code,omitempty" yaml:"code,omitempty"`
	ImageName            string            `json:"image_name,omitempty" yaml:"image_name,omitempty"`
	OldState             string            `json:"old_state,omitempty" yaml:"old_state,omitempty"`
	OldTaskState         string            `json:"old_task_state,omitempty" yaml:"old_task_state,omitempty"`
	NewTaskState         string            `json:"new_task_state,omitempty" yaml:"new_task_state,omitempty"`
	AuditPeriodBeginning string            `json:"audit_period_beginning,omitempty" yaml:"audit_period_beginning,omitempty"`
	AuditPeriodEnding    string            `json:"audit_period_ending,omitempty" yaml:"audit_period_ending,omitempty"`
	VolumeID             string            `json:"volume_id,omitempty" yaml:"volume_id,omitempty"`
	NewInstanceType      string            `json:"new_instance_typ,omitempty" yaml:"new_instance_typ,omitempty"`
	NewInstanceTypeID    int               `json:"new_instance_type_id,omitempty" yaml:"new_instance_type_id,omitempty"`
	Bandwidth            struct {
	} `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`
}

// type ComputeInstanceImageMeta struct {
// 	Architecture                  string `json:"architecture,omitempty" yaml:"architecture,omitempty"`
// 	Description                   string `json:"description,omitempty" yaml:"description,omitempty"`
// 	CommitSha                     string `json:"commit_sha,omitempty" yaml:"commit_sha,omitempty"`
// 	HwDiskBus                     string `json:"hw_disk_bus,omitempty" yaml:"hw_disk_bus,omitempty"`
// 	HwQEMUGuestAgent              string `json:"hw_qemu_guest_agent,omitempty" yaml:"hw_qemu_guest_agent,omitempty"`
// 	HwRngModel                    string `json:"hw_rng_model,omitempty" yaml:"hw_rng_model,omitempty"`
// 	HwScsiModel                   string `json:"hw_scsi_model,omitempty" yaml:"hw_scsi_model,omitempty"`
// 	ImageType                     string `json:"image_type,omitempty" yaml:"image_type,omitempty"`
// 	OsDistro                      string `json:"os_distro,omitempty" yaml:"os_distro,omitempty"`
// 	OwnerSpecifiedOpenstackMd5    string `json:"owner_specified.openstack.md5,omitempty" yaml:"owner_specified.openstack.md5,omitempty"`
// 	OwnerSpecifiedOpenstackObject string `json:"owner_specified.openstack.object,omitempty" yaml:"owner_specified.openstack.object,omitempty"`
// 	OwnerSpecifiedOpenstackSha256 string `json:"owner_specified.openstack.sha256,omitempty" yaml:"owner_specified.openstack.sha256,omitempty"`
// 	MinRAM                        string `json:"min_ram,omitempty" yaml:"min_ram,omitempty"`
// 	MinDisk                       string `json:"min_disk,omitempty" yaml:"min_disk,omitempty"`
// 	DiskFormat                    string `json:"disk_format,omitempty" yaml:"disk_format,omitempty"`
// 	ContainerFormat               string `json:"container_format,omitempty" yaml:"container_format,omitempty"`
// 	BaseImageRef                  string `json:"base_image_ref,omitempty" yaml:"base_image_ref,omitempty"`
// }

// type ComputeInstanceMetadata struct {
// 	Backup      string `json:"BACKUP,omitempty" yaml:"BACKUP,omitempty"`
// 	ClusterName string `json:"cluster_name,omitempty" yaml:"cluster_name,omitempty"`
// 	ImageName   string `json:"image_name,omitempty" yaml:"image_name,omitempty"`
// 	HostID      string `json:"host_id,omitempty" yaml:"host_id,omitempty"`
// 	VMRole      string `json:"vmrole,omitempty" yaml:"vmrole,omitempty"`
// }

// type ComputeInstanceFixedIP struct {
// 	Address     string             `json:"address,omitempty" yaml:"address,omitempty"`
// 	Type        string             `json:"type,omitempty" yaml:"type,omitempty"`
// 	Version     int                `json:"version,omitempty" yaml:"version,omitempty"`
// 	Meta        ComputeFixedIPMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
// 	FloatingIPs []string           `json:"floating_ips,omitempty" yaml:"floating_ips,omitempty"`
// 	Label       string             `json:"label,omitempty" yaml:"label,omitempty"`
// 	VifMac      string             `json:"vif_mac,omitempty" yaml:"vif_mac,omitempty"`
// }

// type ComputeFixedIPMeta struct {
// }

// FromJSON populates a ComputeInstanceNotification using
// the data in the provided JSON.
func (msg *ComputeInstanceNotification) FromJSON(data string) error {
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		slog.Error("failure parsing compute instance notification from JSON", "error", err)
		return err
	}
	return nil
}

// ToJSON returns a string representation of the ComputeInstanceNotification
// in JSON format, pretty-printed or not.
func (msg *ComputeInstanceNotification) ToJSON(pretty bool) (string, error) {
	var bytes []byte
	var err error
	if pretty {
		if bytes, err = json.MarshalIndent(msg, "", "  "); err != nil {
			slog.Error("failure marshaling compute instance notification to JSON", "error", err)
			return "", err
		}
	} else {
		if bytes, err = json.Marshal(msg); err != nil {
			slog.Error("failure marshaling compute instance notification to JSON", "error", err)
			return "", err
		}
	}
	slog.Debug("compute instance notification marshaled to JSON")
	return string(bytes), nil
}

// ToString converts the ComputeInstanceNotification into its JSON one-liner representation.
func (msg *ComputeInstanceNotification) ToString() string {
	value, _ := msg.ToJSON(false)
	return value
}
