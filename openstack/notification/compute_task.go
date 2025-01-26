package notification

// ComputeTask is a family of events emitted by Nova scheduler when
// placing a new workload upon one of the hypervisors.
type ComputeTask struct {
	Base    `json:",inline" yaml:",inline"`
	Payload ComputeTaskPayload `json:"payload,omitempty" yaml:"payload,omitempty"`
}

type ComputeTaskPayload struct {
	RequestSpec struct {
		NumInstances int `json:"num_instances,omitempty" yaml:"num_instances,omitempty"`
		Image        struct {
			ID              string `json:"id,omitempty" yaml:"id,omitempty"`
			Name            string `json:"name,omitempty" yaml:"name,omitempty"`
			Status          string `json:"status,omitempty" yaml:"status,omitempty"`
			Checksum        string `json:"checksum,omitempty" yaml:"checksum,omitempty"`
			Owner           string `json:"owner,omitempty" yaml:"owner,omitempty"`
			Size            int64  `json:"size,omitempty" yaml:"size,omitempty"`
			ContainerFormat string `json:"container_format,omitempty" yaml:"container_format,omitempty"`
			DiskFormat      string `json:"disk_format,omitempty" yaml:"disk_format,omitempty"`
			CreatedAt       string `json:"created_at,omitempty" yaml:"created_at,omitempty"`
			UpdatedAt       string `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
			MinRAM          int    `json:"min_ram,omitempty" yaml:"min_ram,omitempty"`
			MinDisk         int    `json:"min_disk,omitempty" yaml:"min_ram,omitempty"`
			Properties      struct {
				HwArchitecture   string `json:"hw_architecture"`
				HwDiskBus        string `json:"hw_disk_bus,omitempty" yaml:"hw_disk_bus,omitempty"`
				HwQemuGuestAgent bool   `json:"hw_qemu_guest_agent,omitempty" yaml:"hw_qemu_guest_agent,omitempty"`
				HwRngModel       string `json:"hw_rng_model,omitempty" yaml:"hw_rng_model,omitempty"`
				HwScsiModel      string `json:"hw_scsi_model,omitempty" yaml:"hw_scsi_model,omitempty"`
			} `json:"properties,omitempty" yaml:"properties,omitempty"`
		} `json:"image,omitempty" yaml:"image,omitempty"`
		InstanceProperties struct {
			NUMATopology interface{} `json:"numa_topology,omitempty" yaml:"numa_topology,omitempty"`
			PCIRequests  struct {
				Requests []interface{} `json:"requests,omitempty" yaml:"requests,omitempty"`
			} `json:"pci_requests,omitempty" yaml:"pci_requests,omitempty"`
			ProjectID        string `json:"project_id,omitempty" yaml:"project_id,omitempty"`
			UserID           string `json:"user_id,omitempty" yaml:"user_id,omitempty"`
			AvailabilityZone string `json:"availability_zone,omitempty" yaml:"availability_zone,omitempty"`
			UUID             string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
			RootGb           int    `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
			EphemeralGb      int    `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
			MemoryMb         int    `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
			VCPUs            int    `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
		} `json:"instance_properties,omitempty" yaml:"instance_properties,omitempty"`
		InstanceType struct {
			ID          int     `json:"id,omitempty" yaml:"id,omitempty"`
			Name        string  `json:"name,omitempty" yaml:"name,omitempty"`
			MemoryMb    int     `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
			VCPUs       int     `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
			RootGb      int     `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
			EphemeralGb int     `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
			FlavorID    string  `json:"flavorid,omitempty" yaml:"flavorid,omitempty"`
			Swap        int     `json:"swap,omitempty" yaml:"swap,omitempty"`
			RXTXFactor  float64 `json:"rxtx_factor,omitempty" yaml:"rxtx_factor,omitempty"`
			VCPUWeight  int     `json:"vcpu_weight,omitempty" yaml:"vcpu_weight,omitempty"`
			Disabled    bool    `json:"disabled,omitempty" yaml:"disabled,omitempty"`
			IsPublic    bool    `json:"is_public,omitempty" yaml:"is_public,omitempty"`
			ExtraSpecs  struct {
				HwCPUCores                 string `json:"hw:cpu_cores,omitempty" yaml:"hw:cpu_cores,omitempty"`
				HwCPUSockets               string `json:"hw:cpu_sockets,omitempty" yaml:"hw:cpu_sockets,omitempty"`
				HwRngAllowed               string `json:"hw_rng:allowed,omitempty" yaml:"hw_rng:allowed,omitempty"`
				TraitCOMPUTESTATUSDISABLED string `json:"trait:COMPUTE_STATUS_DISABLED,omitempty" yaml:"trait:COMPUTE_STATUS_DISABLED,omitempty"`
			} `json:"extra_specs,omitempty" yaml:"extra_specs,omitempty"`
			Description string      `json:"description,omitempty" yaml:"description,omitempty"`
			CreatedAt   string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
			UpdatedAt   interface{} `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
			DeletedAt   interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
			Deleted     bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
		} `json:"instance_type,omitempty" yaml:"instance_type,omitempty"`
	} `json:"request_spec,omitempty" yaml:"request_spec,omitempty"`
	InstanceProperties struct {
		NUMATopology interface{} `json:"numa_topology,omitempty" yaml:"numa_topology,omitempty"`
		PCIRequests  struct {
			Requests []interface{} `json:"requests,omitempty" yaml:"requests,omitempty"`
		} `json:"pci_requests,omitempty" yaml:"pci_requests,omitempty"`
		ProjectID        string `json:"project_id,omitempty" yaml:"project_id,omitempty"`
		UserID           string `json:"user_id,omitempty" yaml:"user_id,omitempty"`
		AvailabilityZone string `json:"availability_zone,omitempty" yaml:"availability_zone,omitempty"`
		UUID             string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
		RootGb           int    `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
		EphemeralGb      int    `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
		MemoryMb         int    `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
		VCPUs            int    `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
	} `json:"instance_properties,omitempty" yaml:"instance_properties,omitempty"`
	InstanceID string `json:"instance_id,omitempty" yaml:"instance_id,omitempty"`
	State      string `json:"state,omitempty" yaml:"state,omitempty"`
	Method     string `json:"method,omitempty" yaml:"method,omitempty"`
	Reason     string `json:"reason,omitempty" yaml:"reason,omitempty"`
	IP         string `json:"ip,omitempty" yaml:"ip,omitempty"`
}
