package notification

// resize_instance, create_key_pair
type Exception struct {
	Base    `json:",inline" yaml:",inline"`
	Payload struct {
		Exception string `json:"exception"`
		Args      struct {
			UserID            string      `json:"user_id,omitempty" yaml:"user_id,omitempty"`
			KeyName           string      `json:"key_name,omitempty" yaml:"key_name,omitempty"`
			KeyType           string      `json:"key_type,omitempty" yaml:"key_type,omitempty"`
			InjectedFiles     interface{} `json:"injected_files,omitempty" yaml:"injected_files,omitempty"`
			ImageRef          interface{} `json:"image_ref,omitempty" yaml:"image_ref,omitempty"`
			OrigImageRef      interface{} `json:"orig_image_ref,omitempty" yaml:"orig_image_ref,omitempty"`
			OrigSysMetadata   interface{} `json:"orig_sys_metadata,omitempty" yaml:"orig_sys_metadata,omitempty"`
			Bdms              interface{} `json:"bdms,omitempty" yaml:"bdms,omitempty"`
			Recreate          bool        `json:"recreate,omitempty" yaml:"recreate,omitempty"`
			OnSharedStorage   bool        `json:"on_shared_storage,omitempty" yaml:"on_shared_storage,omitempty"`
			PreserveEphemeral bool        `json:"preserve_ephemeral,omitempty" yaml:"preserve_ephemeral,omitempty"`
			TailLength        int         `json:"tail_length,omitempty" yaml:"tail_length,omitempty"`
			ScheduledNode     string      `json:"scheduled_node,omitempty" yaml:"scheduled_node,omitempty"`
			Limits            struct {
			} `json:"limits,omitempty" yaml:"limits,omitempty"`
			Instance struct {
				ID          int    `json:"id,omitempty" yaml:"id,omitempty"`
				UserID      string `json:"user_id,omitempty" yaml:"user_id,omitempty"`
				ProjectID   string `json:"project_id,omitempty" yaml:"project_id,omitempty"`
				ImageRef    string `json:"image_ref,omitempty" yaml:"image_ref,omitempty"`
				KernelID    string `json:"kernel_id,omitempty" yaml:"kernel_id,omitempty"`
				RamdiskID   string `json:"ramdisk_id,omitempty" yaml:"ramdisk_id,omitempty"`
				Hostname    string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
				LaunchIndex int    `json:"launch_index,omitempty" yaml:"launch_index,omitempty"`
				KeyName     string `json:"key_name,omitempty" yaml:"key_name,omitempty"`
				KeyData     string `json:"key_data,omitempty" yaml:"key_data,omitempty"`
				PowerState  int    `json:"power_state,omitempty" yaml:"power_state,omitempty"`
				VMState     string `json:"vm_state,omitempty" yaml:"vm_state,omitempty"`
				TaskState   string `json:"task_state,omitempty" yaml:"task_state,omitempty"`
				Services    []struct {
					ID             int         `json:"id,omitempty" yaml:"id,omitempty"`
					UUID           string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
					Host           string      `json:"host,omitempty" yaml:"host,omitempty"`
					Binary         string      `json:"binary,omitempty" yaml:"binary,omitempty"`
					Topic          string      `json:"topic,omitempty" yaml:"topic,omitempty"`
					ReportCount    int         `json:"report_count,omitempty" yaml:"report_count,omitempty"`
					Disabled       bool        `json:"disabled,omitempty" yaml:"disabled,omitempty"`
					DisabledReason interface{} `json:"disabled_reason,omitempty" yaml:"disabled_reason,omitempty"`
					LastSeenUp     string      `json:"last_seen_up,omitempty" yaml:"last_seen_up,omitempty"`
					ForcedDown     bool        `json:"forced_down,omitempty" yaml:"forced_down,omitempty"`
					Version        int         `json:"version,omitempty" yaml:"version,omitempty"`
					CreatedAt      string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
					UpdatedAt      string      `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
					DeletedAt      interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
					Deleted        bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
				} `json:"services,omitempty" yaml:"services,omitempty"`
				MemoryMb               int               `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
				Vcpus                  int               `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
				RootGb                 int               `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
				EphemeralGb            int               `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
				EphemeralKeyUUID       interface{}       `json:"ephemeral_key_uuid,omitempty" yaml:"ephemeral_key_uuid,omitempty"`
				Host                   string            `json:"host,omitempty" yaml:"host,omitempty"`
				Node                   string            `json:"node,omitempty" yaml:"node,omitempty"`
				InstanceTypeID         int               `json:"instance_type_id,omitempty" yaml:"instance_type_id,omitempty"`
				UserData               string            `json:"user_data,omitempty" yaml:"user_data,omitempty"`
				ReservationID          string            `json:"reservation_id,omitempty" yaml:"reservation_id,omitempty"`
				LaunchedAt             string            `json:"launched_at,omitempty" yaml:"launched_at,omitempty"`
				TerminatedAt           interface{}       `json:"terminated_at,omitempty" yaml:"terminated_at,omitempty"`
				AvailabilityZone       string            `json:"availability_zone,omitempty" yaml:"availability_zone,omitempty"`
				DisplayName            string            `json:"display_name,omitempty" yaml:"display_name,omitempty"`
				DisplayDescription     string            `json:"display_description,omitempty" yaml:"display_description,omitempty"`
				LaunchedOn             string            `json:"launched_on,omitempty" yaml:"launched_on,omitempty"`
				Locked                 bool              `json:"locked,omitempty" yaml:"locked,omitempty"`
				LockedBy               interface{}       `json:"locked_by,omitempty" yaml:"locked_by,omitempty"`
				OsType                 interface{}       `json:"os_type,omitempty" yaml:"os_type,omitempty"`
				Architecture           string            `json:"architecture,omitempty" yaml:"architecture,omitempty"`
				VMMode                 interface{}       `json:"vm_mode,omitempty" yaml:"vm_mode,omitempty"`
				UUID                   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
				RootDeviceName         string            `json:"root_device_name,omitempty" yaml:"root_device_name,omitempty"`
				DefaultEphemeralDevice interface{}       `json:"default_ephemeral_device,omitempty" yaml:"default_ephemeral_device,omitempty"`
				DefaultSwapDevice      interface{}       `json:"default_swap_device,omitempty" yaml:"default_swap_device,omitempty"`
				ConfigDrive            string            `json:"config_drive,omitempty" yaml:"config_drive,omitempty"`
				AccessIPV4             interface{}       `json:"access_ip_v4,omitempty" yaml:"access_ip_v4,omitempty"`
				AccessIPV6             interface{}       `json:"access_ip_v6,omitempty" yaml:"access_ip_v6,omitempty"`
				AutoDiskConfig         bool              `json:"auto_disk_config,omitempty" yaml:"auto_disk_config,omitempty"`
				Progress               int               `json:"progress,omitempty" yaml:"progress,omitempty"`
				ShutdownTerminate      bool              `json:"shutdown_terminate,omitempty" yaml:"shutdown_terminate,omitempty"`
				DisableTerminate       bool              `json:"disable_terminate,omitempty" yaml:"disable_terminate,omitempty"`
				CellName               interface{}       `json:"cell_name,omitempty" yaml:"cell_name,omitempty"`
				Metadata               map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
				SystemMetadata         struct {
					ImageArchitecture                  string `json:"image_architecture,omitempty" yaml:"image_architecture,omitempty"`
					ImageCommitSha                     string `json:"image_commit_sha,omitempty" yaml:"image_commit_sha,omitempty"`
					ImageDescription                   string `json:"image_description,omitempty" yaml:"image_description,omitempty"`
					ImageHwDiskBus                     string `json:"image_hw_disk_bus,omitempty" yaml:"image_hw_disk_bus,omitempty"`
					ImageHwQemuGuestAgent              string `json:"image_hw_qemu_guest_agent,omitempty" yaml:"image_hw_qemu_guest_agent,omitempty"`
					ImageHwRngModel                    string `json:"image_hw_rng_model,omitempty" yaml:"image_hw_rng_model,omitempty"`
					ImageHwScsiModel                   string `json:"image_hw_scsi_model,omitempty" yaml:"image_hw_scsi_model,omitempty"`
					ImageImageType                     string `json:"image_image_type,omitempty" yaml:"image_image_type,omitempty"`
					ImageOsDistro                      string `json:"image_os_distro" yaml:"image_os_distro"`
					ImageOwnerSpecifiedOpenstackMd5    string `json:"image_owner_specified.openstack.md5" yaml:"image_owner_specified.openstack.md5"`
					ImageOwnerSpecifiedOpenstackObject string `json:"image_owner_specified.openstack.object" yaml:"image_owner_specified.openstack.object"`
					ImageOwnerSpecifiedOpenstackSha256 string `json:"image_owner_specified.openstack.sha256" yaml:"image_owner_specified.openstack.sha256"`
					ImageMinRAM                        string `json:"image_min_ram,omitempty" yaml:"image_min_ram,omitempty"`
					ImageMinDisk                       string `json:"image_min_disk,omitempty" yaml:"image_min_disk,omitempty"`
					ImageDiskFormat                    string `json:"image_disk_format,omitempty" yaml:"image_disk_format,omitempty"`
					ImageContainerFormat               string `json:"image_container_format,omitempty" yaml:"image_container_format,omitempty"`
					ImageBaseImageRef                  string `json:"image_base_image_ref,omitempty" yaml:"image_base_image_ref,omitempty"`
					OwnerUserName                      string `json:"owner_user_name,omitempty" yaml:"owner_user_name,omitempty"`
					OwnerProjectName                   string `json:"owner_project_name,omitempty" yaml:"owner_project_name,omitempty"`
					BootRoles                          string `json:"boot_roles,omitempty" yaml:"boot_roles,omitempty"`
					CleanAttempts                      string `json:"clean_attempts,omitempty" yaml:"clean_attempts,omitempty"`
					OldVMState                         string `json:"old_vm_state,omitempty" yaml:"old_vm_state,omitempty"`
				} `json:"system_metadata,omitempty" yaml:"system_metadata,omitempty"`
				InfoCache struct {
					ChangedFields []interface{} `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
					Context       struct {
						UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
						ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
						DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
						UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
						ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
						AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
						UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
						ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
						DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
						SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
						UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
						ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
						IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
						IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
						ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
						ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
						ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
						Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
						ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
						ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
						ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
						ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
						ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
						ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
						RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
						GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
						ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
						RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
						Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
						ServiceCatalog           []struct {
							Type      string `json:"type,omitempty" yaml:"type,omitempty"`
							Name      string `json:"name,omitempty" yaml:"name,omitempty"`
							Endpoints []struct {
								Region      string `json:"region,omitempty" yaml:"region,omitempty"`
								InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
								PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
								AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
							} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
						} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
						QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
						DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
						MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
						CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
						UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
					} `json:"_context,omitempty" yaml:"_context,omitempty"`
					Version         string `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
					ObjInstanceUUID string `json:"_obj_instance_uuid,omitempty" yaml:"_obj_instance_uuid,omitempty"`
					ObjNetworkInfo  []struct {
						ID      string `json:"id,omitempty" yaml:"id,omitempty"`
						Address string `json:"address,omitempty" yaml:"address,omitempty"`
						Network struct {
							ID      string `json:"id,omitempty" yaml:"id,omitempty"`
							Bridge  string `json:"bridge,omitempty" yaml:"bridge,omitempty"`
							Label   string `json:"label,omitempty" yaml:"label,omitempty"`
							Subnets []struct {
								Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`
								DNS  []struct {
									Address string `json:"address,omitempty" yaml:"address,omitempty"`
									Type    string `json:"type,omitempty" yaml:"type,omitempty"`
									Version int    `json:"version,omitempty" yaml:"version,omitempty"`
									Meta    struct {
									} `json:"meta,omitempty" yaml:"meta,omitempty"`
								} `json:"dns,omitempty" yaml:"dns,omitempty"`
								Gateway struct {
									Address string `json:"address,omitempty" yaml:"address,omitempty"`
									Type    string `json:"type,omitempty" yaml:"type,omitempty"`
									Version int    `json:"version,omitempty" yaml:"version,omitempty"`
									Meta    struct {
									} `json:"meta,omitempty" yaml:"meta,omitempty"`
								} `json:"gateway,omitempty" yaml:"gateway,omitempty"`
								Ips []struct {
									Address string `json:"address,omitempty" yaml:"address,omitempty"`
									Type    string `json:"type,omitempty" yaml:"type,omitempty"`
									Version int    `json:"version,omitempty" yaml:"version,omitempty"`
									Meta    struct {
									} `json:"meta,omitempty" yaml:"meta,omitempty"`
									FloatingIps []interface{} `json:"floating_ips,omitempty" yaml:"floating_ips,omitempty"`
									Label       string        `json:"label" yaml:"label"`
									VifMac      string        `json:"vif_mac" yaml:"vif_mac"`
								} `json:"ips,omitempty" yaml:"ips,omitempty"`
								Routes  []interface{} `json:"routes,omitempty" yaml:"routes,omitempty"`
								Version int           `json:"version,omitempty" yaml:"version,omitempty"`
								Meta    struct {
									DhcpServer string `json:"dhcp_server,omitempty" yaml:"dhcp_server,omitempty"`
								} `json:"meta,omitempty" yaml:"meta,omitempty"`
							} `json:"subnets,omitempty" yaml:"subnets,omitempty"`
							Meta struct {
								Injected        bool   `json:"injected,omitempty" yaml:"injected,omitempty"`
								TenantID        string `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
								Mtu             int    `json:"mtu,omitempty" yaml:"mtu,omitempty"`
								PhysicalNetwork string `json:"physical_network,omitempty" yaml:"physical_network,omitempty"`
								Tunneled        bool   `json:"tunneled,omitempty" yaml:"tunneled,omitempty"`
							} `json:"meta,omitempty" yaml:"meta,omitempty"`
						} `json:"network,omitempty" yaml:"network,omitempty"`
						Type    string `json:"type,omitempty" yaml:"type,omitempty"`
						Details struct {
							PortFilter bool `json:"port_filter,omitempty" yaml:"port_filter,omitempty"`
						} `json:"details,omitempty" yaml:"details,omitempty"`
						Devname        string      `json:"devname,omitempty" yaml:"devname,omitempty"`
						OvsInterfaceid string      `json:"ovs_interfaceid,omitempty" yaml:"ovs_interfaceid,omitempty"`
						QbhParams      interface{} `json:"qbh_params,omitempty" yaml:"qbh_params,omitempty"`
						QbgParams      interface{} `json:"qbg_params,omitempty" yaml:"qbg_params,omitempty"`
						Active         bool        `json:"active,omitempty" yaml:"active,omitempty"`
						VnicType       string      `json:"vnic_type,omitempty" yaml:"vnic_type,omitempty"`
						Profile        struct {
						} `json:"profile,omitempty" yaml:"profile,omitempty"`
						PreserveOnDelete bool `json:"preserve_on_delete,omitempty" yaml:"preserve_on_delete,omitempty"`
						Meta             struct {
						} `json:"meta,omitempty" yaml:"meta,omitempty"`
					} `json:"_obj_network_info,omitempty" yaml:"_obj_network_info,omitempty"`
					ObjCreatedAt string      `json:"_obj_created_at,omitempty" yaml:"_obj_created_at,omitempty"`
					ObjUpdatedAt string      `json:"_obj_updated_at,omitempty" yaml:"_obj_updated_at,omitempty"`
					ObjDeletedAt interface{} `json:"_obj_deleted_at,omitempty" yaml:"_obj_deleted_at,omitempty"`
					ObjDeleted   bool        `json:"_obj_deleted,omitempty" yaml:"_obj_deleted,omitempty"`
				} `json:"info_cache,omitempty" yaml:"info_cache,omitempty"`
				SecurityGroups []interface{} `json:"security_groups,omitempty" yaml:"security_groups,omitempty"`
				Cleaned        bool          `json:"cleaned,omitempty" yaml:"cleaned,omitempty"`
				DeviceMetadata interface{}   `json:"device_metadata,omitempty" yaml:"device_metadata,omitempty"`
				PciDevices     []interface{} `json:"pci_devices,omitempty" yaml:"pci_devices,omitempty"`
				NumaTopology   interface{}   `json:"numa_topology,omitempty" yaml:"numa_topology,omitempty"`
				Hidden         bool          `json:"hidden,omitempty" yaml:"hidden,omitempty"`
				Resources      interface{}   `json:"resources,omitempty" yaml:"resources,omitempty"`
				CreatedAt      string        `json:"created_at,omitempty" yaml:"created_at,omitempty"`
				UpdatedAt      string        `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
				DeletedAt      string        `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
				Deleted        bool          `json:"deleted,omitempty" yaml:"deleted,omitempty"`
				Name           string        `json:"name,omitempty" yaml:"name,omitempty"`
				PciRequests    struct {
					InstanceUUID string        `json:"instance_uuid,omitempty" yaml:"instance_uuid,omitempty"`
					Requests     []interface{} `json:"requests,omitempty" yaml:"requests,omitempty"`
				} `json:"pci_requests,omitempty" yaml:"pci_requests,omitempty"`
				Flavor struct {
					ID          int     `json:"id,omitempty" yaml:"id,omitempty"`
					Name        string  `json:"name,omitempty" yaml:"name,omitempty"`
					MemoryMb    int     `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
					Vcpus       int     `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
					RootGb      int     `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
					EphemeralGb int     `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
					Flavorid    string  `json:"flavorid,omitempty" yaml:"flavorid,omitempty"`
					Swap        int     `json:"swap,omitempty" yaml:"swap,omitempty"`
					RxtxFactor  float64 `json:"rxtx_factor,omitempty" yaml:"rxtx_factor,omitempty"`
					VcpuWeight  int     `json:"vcpu_weight,omitempty" yaml:"vcpu_weight,omitempty"`
					Disabled    bool    `json:"disabled,omitempty" yaml:"disabled,omitempty"`
					IsPublic    bool    `json:"is_public,omitempty" yaml:"is_public,omitempty"`
					ExtraSpecs  struct {
						HwCPUCores   string `json:"hw:cpu_cores,omitempty" yaml:"hw:cpu_cores,omitempty"`
						HwCPUSockets string `json:"hw:cpu_sockets,omitempty" yaml:"hw:cpu_sockets,omitempty"`
						HwRngAllowed string `json:"hw_rng:allowed,omitempty" yaml:"hw_rng:allowed,omitempty"`
					} `json:"extra_specs,omitempty" yaml:"extra_specs,omitempty"`
					Description interface{} `json:"description,omitempty" yaml:"description,omitempty"`
					CreatedAt   string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
					UpdatedAt   interface{} `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
					DeletedAt   interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
					Deleted     bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
				} `json:"flavor,omitempty" yaml:"flavor,omitempty"`
				OldFlavor interface{} `json:"old_flavor,omitempty" yaml:"old_flavor,omitempty"`
				NewFlavor struct {
					ID          int     `json:"id,omitempty" yaml:"id,omitempty"`
					Name        string  `json:"name,omitempty" yaml:"name,omitempty"`
					MemoryMb    int     `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
					Vcpus       int     `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
					RootGb      int     `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
					EphemeralGb int     `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
					Flavorid    string  `json:"flavorid,omitempty" yaml:"flavorid,omitempty"`
					Swap        int     `json:"swap,omitempty" yaml:"swap,omitempty"`
					RxtxFactor  float64 `json:"rxtx_factor,omitempty" yaml:"rxtx_factor,omitempty"`
					VcpuWeight  int     `json:"vcpu_weight,omitempty" yaml:"vcpu_weight,omitempty"`
					Disabled    bool    `json:"disabled,omitempty" yaml:"disabled,omitempty"`
					IsPublic    bool    `json:"is_public,omitempty" yaml:"is_public,omitempty"`
					ExtraSpecs  struct {
						HwCPUCores   string `json:"hw:cpu_cores,omitempty" yaml:"hw:cpu_cores,omitempty"`
						HwCPUSockets string `json:"hw:cpu_sockets,omitempty" yaml:"hw:cpu_sockets,omitempty"`
						HwRngAllowed string `json:"hw_rng:allowed,omitempty" yaml:"hw_rng:allowed,omitempty"`
					} `json:"extra_specs,omitempty" yaml:"extra_specs,omitempty"`
					Description interface{} `json:"description,omitempty" yaml:"description,omitempty"`
					CreatedAt   string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
					UpdatedAt   interface{} `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
					DeletedAt   interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
					Deleted     bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
				} `json:"new_flavor,omitempty" yaml:"new_flavor,omitempty"`
				MigrationContext struct {
					ChangedFields []string `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
					Context       struct {
						UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
						ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
						DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
						UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
						ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
						AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
						UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
						ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
						DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
						SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
						UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
						ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
						IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
						IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
						ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
						ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
						ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
						Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
						ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
						ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
						ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
						ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
						ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
						ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
						RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
						GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
						ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
						RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
						Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
						ServiceCatalog           []struct {
							Type      string `json:"type,omitempty" yaml:"type,omitempty"`
							Name      string `json:"name,omitempty" yaml:"name,omitempty"`
							Endpoints []struct {
								Region      string `json:"region,omitempty" yaml:"region,omitempty"`
								InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
								PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
								AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
							} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
						} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
						QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
						DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
						MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
						CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
						UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
					} `json:"_context,omitempty" yaml:"_context,omitempty"`
					Version            string        `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
					ObjInstanceUUID    string        `json:"_obj_instance_uuid,omitempty" yaml:"_obj_instance_uuid,omitempty"`
					ObjMigrationID     int           `json:"_obj_migration_id,omitempty" yaml:"_obj_migration_id,omitempty"`
					ObjNewNumaTopology interface{}   `json:"_obj_new_numa_topology,omitempty" yaml:"_obj_new_numa_topology,omitempty"`
					ObjOldNumaTopology interface{}   `json:"_obj_old_numa_topology,omitempty" yaml:"_obj_old_numa_topology,omitempty"`
					ObjNewPciDevices   []interface{} `json:"_obj_new_pci_devices,omitempty" yaml:"_obj_new_pci_devices,omitempty"`
					ObjOldPciDevices   []interface{} `json:"_obj_old_pci_devices,omitempty" yaml:"_obj_old_pci_devices,omitempty"`
					ObjNewPciRequests  struct {
						InstanceUUID string        `json:"instance_uuid,omitempty" yaml:"instance_uuid,omitempty"`
						Requests     []interface{} `json:"requests,omitempty" yaml:"requests,omitempty"`
					} `json:"_obj_new_pci_requests,omitempty" yaml:"_obj_new_pci_requests,omitempty"`
					ObjOldPciRequests struct {
						InstanceUUID string        `json:"instance_uuid,omitempty" yaml:"instance_uuid,omitempty"`
						Requests     []interface{} `json:"requests,omitempty" yaml:"requests,omitempty"`
					} `json:"_obj_old_pci_requests,omitempty" yaml:"_obj_old_pci_requests,omitempty"`
					ObjNewResources interface{} `json:"_obj_new_resources,omitempty" yaml:"_obj_new_resources,omitempty"`
					ObjOldResources interface{} `json:"_obj_old_resources,omitempty" yaml:"_obj_old_resources,omitempty"`
				} `json:"migration_context,omitempty" yaml:"migration_context,omitempty"`
			} `json:"instance,omitempty" yaml:"instance,omitempty"`
			VolumeID       string      `json:"volume_id,omitempty" yaml:"volume_id,omitempty"`
			AttachmentID   string      `json:"attachment_id,omitempty" yaml:"attachment_id,omitempty"`
			BlockMigration interface{} `json:"block_migration" yaml:"block_migration"`
			Disk           interface{} `json:"disk" yaml:"disk"`
			PortID         string      `json:"port_id,omitempty" yaml:"port_id,omitempty"`
			MigrateData    struct {
				ChangedFields []string `json:"_changed_fields" yaml:"_changed_fields"`
				Context       struct {
					UserID                   string        `json:"_user_id" yaml:"_user_id"`
					ProjectID                string        `json:"_project_id" yaml:"_project_id"`
					DomainID                 interface{}   `json:"_domain_id" yaml:"_domain_id"`
					UserDomainID             string        `json:"_user_domain_id" yaml:"_user_domain_id"`
					ProjectDomainID          string        `json:"_project_domain_id" yaml:"_project_domain_id"`
					AuthToken                string        `json:"auth_token" yaml:"auth_token"`
					UserName                 string        `json:"user_name" yaml:"user_name"`
					ProjectName              string        `json:"project_name" yaml:"project_name"`
					DomainName               interface{}   `json:"domain_name" yaml:"domain_name"`
					SystemScope              interface{}   `json:"system_scope" yaml:"system_scope"`
					UserDomainName           interface{}   `json:"user_domain_name" yaml:"user_domain_name"`
					ProjectDomainName        interface{}   `json:"project_domain_name" yaml:"project_domain_name"`
					IsAdmin                  bool          `json:"is_admin" yaml:"is_admin"`
					IsAdminProject           bool          `json:"is_admin_project" yaml:"is_admin_project"`
					ReadOnly                 bool          `json:"read_only" yaml:"read_only"`
					ShowDeleted              bool          `json:"show_deleted" yaml:"show_deleted"`
					ResourceUUID             interface{}   `json:"resource_uuid" yaml:"resource_uuid"`
					Roles                    []string      `json:"roles" yaml:"roles"`
					ServiceToken             interface{}   `json:"service_token" yaml:"service_token"`
					ServiceUserID            interface{}   `json:"service_user_id" yaml:"service_user_id"`
					ServiceUserName          interface{}   `json:"service_user_name" yaml:"service_user_name"`
					ServiceUserDomainID      interface{}   `json:"service_user_domain_id" yaml:"service_user_domain_id"`
					ServiceUserDomainName    interface{}   `json:"service_user_domain_name" yaml:"service_user_domain_name"`
					ServiceProjectID         interface{}   `json:"service_project_id" yaml:"service_project_id"`
					ServiceProjectName       interface{}   `json:"service_project_name" yaml:"service_project_name"`
					ServiceProjectDomainID   interface{}   `json:"service_project_domain_id" yaml:"service_project_domain_id"`
					ServiceProjectDomainName interface{}   `json:"service_project_domain_name" yaml:"service_project_domain_name"`
					ServiceRoles             []interface{} `json:"service_roles" yaml:"service_roles"`
					RequestID                string        `json:"request_id" yaml:"request_id"`
					GlobalRequestID          interface{}   `json:"global_request_id" yaml:"global_request_id"`
					ReadDeleted              string        `json:"_read_deleted" yaml:"_read_deleted"`
					RemoteAddress            string        `json:"remote_address" yaml:"remote_address"`
					Timestamp                string        `json:"timestamp" yaml:"timestamp"`
					ServiceCatalog           []struct {
						Type      string `json:"type" yaml:"type"`
						Name      string `json:"name" yaml:"name"`
						Endpoints []struct {
							Region      string `json:"region" yaml:"region"`
							InternalURL string `json:"internalURL" yaml:"internalURL"`
							PublicURL   string `json:"publicURL" yaml:"publicURL"`
							AdminURL    string `json:"adminURL" yaml:"adminURL"`
						} `json:"endpoints" yaml:"endpoints"`
					} `json:"service_catalog" yaml:"service_catalog"`
					QuotaClass     interface{} `json:"quota_class" yaml:"quota_class"`
					DbConnection   interface{} `json:"db_connection" yaml:"db_connection"`
					MqConnection   interface{} `json:"mq_connection" yaml:"mq_connection"`
					CellUUID       interface{} `json:"cell_uuid" yaml:"cell_uuid"`
					UserAuthPlugin interface{} `json:"user_auth_plugin" yaml:"user_auth_plugin"`
				} `json:"_context" yaml:"_context"`
				Version                     string      `json:"VERSION" yaml:"VERSION"`
				ObjFilename                 string      `json:"_obj_filename" yaml:"_obj_filename"`
				ObjImageType                string      `json:"_obj_image_type" yaml:"_obj_image_type"`
				ObjBlockMigration           bool        `json:"_obj_block_migration" yaml:"_obj_block_migration"`
				ObjDiskAvailableMb          int         `json:"_obj_disk_available_mb" yaml:"_obj_disk_available_mb"`
				ObjIsSharedInstancePath     bool        `json:"_obj_is_shared_instance_path" yaml:"_obj_is_shared_instance_path"`
				ObjIsSharedBlockStorage     bool        `json:"_obj_is_shared_block_storage" yaml:"_obj_is_shared_block_storage"`
				ObjInstanceRelativePath     string      `json:"_obj_instance_relative_path" yaml:"_obj_instance_relative_path"`
				ObjGraphicsListenAddrVnc    string      `json:"_obj_graphics_listen_addr_vnc" yaml:"_obj_graphics_listen_addr_vnc"`
				ObjGraphicsListenAddrSpice  string      `json:"_obj_graphics_listen_addr_spice" yaml:"_obj_graphics_listen_addr_spice"`
				ObjSerialListenAddr         interface{} `json:"_obj_serial_listen_addr" yaml:"_obj_serial_listen_addr"`
				ObjDstWantsFileBackedMemory bool        `json:"_obj_dst_wants_file_backed_memory" yaml:"_obj_dst_wants_file_backed_memory"`
				ObjFileBackedMemoryDiscard  bool        `json:"_obj_file_backed_memory_discard" yaml:"_obj_file_backed_memory_discard"`
				ObjIsVolumeBacked           bool        `json:"_obj_is_volume_backed" yaml:"_obj_is_volume_backed"`
				ObjVifs                     []struct {
					ChangedFields []string `json:"_changed_fields" yaml:"_changed_fields"`
					Context       struct {
						UserID                   string        `json:"_user_id" yaml:"_user_id"`
						ProjectID                string        `json:"_project_id" yaml:"_project_id"`
						DomainID                 interface{}   `json:"_domain_id" yaml:"_domain_id"`
						UserDomainID             string        `json:"_user_domain_id" yaml:"_user_domain_id"`
						ProjectDomainID          string        `json:"_project_domain_id" yaml:"_project_domain_id"`
						AuthToken                string        `json:"auth_token" yaml:"auth_token"`
						UserName                 string        `json:"user_name" yaml:"user_name"`
						ProjectName              string        `json:"project_name" yaml:"project_name"`
						DomainName               interface{}   `json:"domain_name" yaml:"domain_name"`
						SystemScope              interface{}   `json:"system_scope" yaml:"system_scope"`
						UserDomainName           interface{}   `json:"user_domain_name" yaml:"user_domain_name"`
						ProjectDomainName        interface{}   `json:"project_domain_name" yaml:"project_domain_name"`
						IsAdmin                  bool          `json:"is_admin" yaml:"is_admin"`
						IsAdminProject           bool          `json:"is_admin_project" yaml:"is_admin_project"`
						ReadOnly                 bool          `json:"read_only" yaml:"read_only"`
						ShowDeleted              bool          `json:"show_deleted" yaml:"show_deleted"`
						ResourceUUID             interface{}   `json:"resource_uuid" yaml:"resource_uuid"`
						Roles                    []string      `json:"roles" yaml:"roles"`
						ServiceToken             interface{}   `json:"service_token" yaml:"service_token"`
						ServiceUserID            interface{}   `json:"service_user_id" yaml:"service_user_id"`
						ServiceUserName          interface{}   `json:"service_user_name" yaml:"service_user_name"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id" yaml:"service_user_domain_id"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name" yaml:"service_user_domain_name"`
						ServiceProjectID         interface{}   `json:"service_project_id" yaml:"service_project_id"`
						ServiceProjectName       interface{}   `json:"service_project_name" yaml:"service_project_name"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id" yaml:"service_project_domain_id"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name" yaml:"service_project_domain_name"`
						ServiceRoles             []interface{} `json:"service_roles" yaml:"service_roles"`
						RequestID                string        `json:"request_id" yaml:"request_id"`
						GlobalRequestID          interface{}   `json:"global_request_id" yaml:"global_request_id"`
						ReadDeleted              string        `json:"_read_deleted" yaml:"_read_deleted"`
						RemoteAddress            string        `json:"remote_address" yaml:"remote_address"`
						Timestamp                string        `json:"timestamp" yaml:"timestamp"`
						ServiceCatalog           []struct {
							Type      string `json:"type" yaml:"type"`
							Name      string `json:"name" yaml:"name"`
							Endpoints []struct {
								Region      string `json:"region" yaml:"region"`
								InternalURL string `json:"internalURL" yaml:"internalURL"`
								PublicURL   string `json:"publicURL" yaml:"publicURL"`
								AdminURL    string `json:"adminURL" yaml:"adminURL"`
							} `json:"endpoints" yaml:"endpoints"`
						} `json:"service_catalog" yaml:"service_catalog"`
						QuotaClass     interface{} `json:"quota_class" yaml:"quota_class"`
						DbConnection   interface{} `json:"db_connection" yaml:"db_connection"`
						MqConnection   interface{} `json:"mq_connection" yaml:"mq_connection"`
						CellUUID       interface{} `json:"cell_uuid" yaml:"cell_uuid"`
						UserAuthPlugin interface{} `json:"user_auth_plugin" yaml:"user_auth_plugin"`
					} `json:"_context" yaml:"_context"`
					Version           string `json:"VERSION" yaml:"VERSION"`
					ObjPortID         string `json:"_obj_port_id" yaml:"_obj_port_id"`
					ObjVnicType       string `json:"_obj_vnic_type" yaml:"_obj_vnic_type"`
					ObjVifType        string `json:"_obj_vif_type" yaml:"_obj_vif_type"`
					ObjVifDetailsJSON string `json:"_obj_vif_details_json" yaml:"_obj_vif_details_json"`
					ObjProfileJSON    string `json:"_obj_profile_json" yaml:"_obj_profile_json"`
					ObjHost           string `json:"_obj_host" yaml:"_obj_host"`
					ObjSourceVif      struct {
						ID      string `json:"id" yaml:"id"`
						Address string `json:"address" yaml:"address"`
						Network struct {
							ID      string `json:"id" yaml:"id"`
							Bridge  string `json:"bridge" yaml:"bridge"`
							Label   string `json:"label" yaml:"label"`
							Subnets []struct {
								Cidr string `json:"cidr" yaml:"cidr"`
								DNS  []struct {
									Address string `json:"address" yaml:"address"`
									Type    string `json:"type" yaml:"type"`
									Version int    `json:"version" yaml:"version"`
									Meta    struct {
									} `json:"meta" yaml:"meta"`
								} `json:"dns" yaml:"dns"`
								Gateway struct {
									Address string `json:"address" yaml:"address"`
									Type    string `json:"type" yaml:"type"`
									Version int    `json:"version" yaml:"version"`
									Meta    struct {
									} `json:"meta" yaml:"meta"`
								} `json:"gateway" yaml:"gateway"`
								Ips []struct {
									Address string `json:"address" yaml:"address"`
									Type    string `json:"type" yaml:"type"`
									Version int    `json:"version" yaml:"version"`
									Meta    struct {
									} `json:"meta" yaml:"meta"`
									FloatingIps []interface{} `json:"floating_ips" yaml:"floating_ips"`
								} `json:"ips" yaml:"ips"`
								Routes  []interface{} `json:"routes" yaml:"routes"`
								Version int           `json:"version" yaml:"version"`
								Meta    struct {
									DhcpServer string `json:"dhcp_server" yaml:"dhcp_server"`
								} `json:"meta" yaml:"meta"`
							} `json:"subnets" yaml:"subnets"`
							Meta struct {
								Injected        bool   `json:"injected" yaml:"injected"`
								TenantID        string `json:"tenant_id" yaml:"tenant_id"`
								Mtu             int    `json:"mtu" yaml:"mtu"`
								PhysicalNetwork string `json:"physical_network" yaml:"physical_network"`
								Tunneled        bool   `json:"tunneled" yaml:"tunneled"`
							} `json:"meta" yaml:"meta"`
						} `json:"network" yaml:"network"`
						Type    string `json:"type" yaml:"type"`
						Details struct {
							PortFilter bool `json:"port_filter" yaml:"port_filter"`
						} `json:"details" yaml:"details"`
						Devname        string      `json:"devname" yaml:"devname"`
						OvsInterfaceid string      `json:"ovs_interfaceid" yaml:"ovs_interfaceid"`
						QbhParams      interface{} `json:"qbh_params" yaml:"qbh_params"`
						QbgParams      interface{} `json:"qbg_params" yaml:"qbg_params"`
						Active         bool        `json:"active" yaml:"active"`
						VnicType       string      `json:"vnic_type" yaml:"vnic_type"`
						Profile        struct {
						} `json:"profile" yaml:"profile"`
						PreserveOnDelete bool `json:"preserve_on_delete" yaml:"preserve_on_delete"`
						Meta             struct {
						} `json:"meta" yaml:"meta"`
					} `json:"_obj_source_vif" yaml:"_obj_source_vif"`
				} `json:"_obj_vifs" yaml:"_obj_vifs"`
				ObjOldVolAttachmentIds struct {
				} `json:"_obj_old_vol_attachment_ids" yaml:"_obj_old_vol_attachment_ids"`
			} `json:"migrate_data" yaml:"migrate_data"`
			Migration struct {
				ID                int         `json:"id,omitempty" yaml:"id,omitempty"`
				UUID              string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
				SourceCompute     string      `json:"source_compute,omitempty" yaml:"source_compute,omitempty"`
				DestCompute       string      `json:"dest_compute,omitempty" yaml:"dest_compute,omitempty"`
				SourceNode        string      `json:"source_node,omitempty" yaml:"source_node,omitempty"`
				DestNode          string      `json:"dest_node,omitempty" yaml:"dest_node,omitempty"`
				DestHost          string      `json:"dest_host,omitempty" yaml:"dest_host,omitempty"`
				OldInstanceTypeID int         `json:"old_instance_type_id,omitempty" yaml:"old_instance_type_id,omitempty"`
				NewInstanceTypeID int         `json:"new_instance_type_id,omitempty" yaml:"new_instance_type_id,omitempty"`
				InstanceUUID      string      `json:"instance_uuid,omitempty" yaml:"instance_uuid,omitempty"`
				Status            string      `json:"status,omitempty" yaml:"status,omitempty"`
				MigrationType     string      `json:"migration_type,omitempty" yaml:"migration_type,omitempty"`
				Hidden            bool        `json:"hidden,omitempty" yaml:"hidden,omitempty"`
				MemoryTotal       interface{} `json:"memory_total,omitempty" yaml:"memory_total,omitempty"`
				MemoryProcessed   interface{} `json:"memory_processed,omitempty" yaml:"memory_processed,omitempty"`
				MemoryRemaining   interface{} `json:"memory_remaining,omitempty" yaml:"memory_remaining,omitempty"`
				DiskTotal         interface{} `json:"disk_total,omitempty" yaml:"disk_total,omitempty"`
				DiskProcessed     interface{} `json:"disk_processed,omitempty" yaml:"disk_processed,omitempty"`
				DiskRemaining     interface{} `json:"disk_remaining,omitempty" yaml:"disk_remaining,omitempty"`
				CrossCellMove     bool        `json:"cross_cell_move,omitempty" yaml:"cross_cell_move,omitempty"`
				UserID            string      `json:"user_id,omitempty" yaml:"user_id,omitempty"`
				ProjectID         string      `json:"project_id,omitempty" yaml:"project_id,omitempty"`
				CreatedAt         string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
				UpdatedAt         string      `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
				DeletedAt         interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
				Deleted           bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
			} `json:"migration,omitempty" yaml:"migration,omitempty"`
			Image struct {
				ID              string `json:"id,omitempty" yaml:"id,omitempty"`
				Name            string `json:"name,omitempty" yaml:"name,omitempty"`
				Status          string `json:"status,omitempty" yaml:"status,omitempty"`
				Checksum        string `json:"checksum,omitempty" yaml:"checksum,omitempty"`
				Owner           string `json:"owner,omitempty" yaml:"owner,omitempty"`
				Size            int    `json:"size,omitempty" yaml:"size,omitempty"`
				ContainerFormat string `json:"container_format,omitempty" yaml:"container_format,omitempty"`
				DiskFormat      string `json:"disk_format,omitempty" yaml:"disk_format,omitempty"`
				CreatedAt       string `json:"created_at,omitempty" yaml:"created_at,omitempty"`
				UpdatedAt       string `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
				MinRAM          int    `json:"min_ram,omitempty" yaml:"min_ram,omitempty"`
				MinDisk         int    `json:"min_disk,omitempty" yaml:"min_disk,omitempty"`
				Properties      struct {
					HwArchitecture   string `json:"hw_architecture,omitempty" yaml:"hw_architecture,omitempty"`
					HwDiskBus        string `json:"hw_disk_bus,omitempty" yaml:"hw_disk_bus,omitempty"`
					HwQemuGuestAgent bool   `json:"hw_qemu_guest_agent,omitempty" yaml:"hw_qemu_guest_agent,omitempty"`
					HwRngModel       string `json:"hw_rng_model,omitempty" yaml:"hw_rng_model,omitempty"`
					HwScsiModel      string `json:"hw_scsi_model,omitempty" yaml:"hw_scsi_model,omitempty"`
				} `json:"properties,omitempty" yaml:"properties,omitempty"`
			} `json:"image,omitempty" yaml:"image,omitempty"`
			InstanceType struct {
				ID          int     `json:"id,omitempty" yaml:"id,omitempty"`
				Name        string  `json:"name,omitempty" yaml:"name,omitempty"`
				MemoryMb    int     `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
				Vcpus       int     `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
				RootGb      int     `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
				EphemeralGb int     `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
				Flavorid    string  `json:"flavorid,omitempty" yaml:"flavorid,omitempty"`
				Swap        int     `json:"swap,omitempty" yaml:"swap,omitempty"`
				RxtxFactor  float64 `json:"rxtx_factor,omitempty" yaml:"rxtx_factor,omitempty"`
				VcpuWeight  int     `json:"vcpu_weight,omitempty" yaml:"vcpu_weight,omitempty"`
				Disabled    bool    `json:"disabled,omitempty" yaml:"disabled,omitempty"`
				IsPublic    bool    `json:"is_public,omitempty" yaml:"is_public,omitempty"`
				ExtraSpecs  struct {
					HwCPUCores   string `json:"hw:cpu_cores,omitempty" yaml:"hw:cpu_cores,omitempty"`
					HwCPUSockets string `json:"hw:cpu_sockets,omitempty" yaml:"hw:cpu_sockets,omitempty"`
					HwRngAllowed string `json:"hw_rng:allowed,omitempty" yaml:"hw_rng:allowed,omitempty"`
				} `json:"extra_specs,omitempty" yaml:"extra_specs,omitempty"`
				Description interface{} `json:"description,omitempty" yaml:"description,omitempty"`
				CreatedAt   string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
				UpdatedAt   interface{} `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
				DeletedAt   interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
				Deleted     bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
			} `json:"instance_type,omitempty" yaml:"instance_type,omitempty"`
			CleanShutdown bool `json:"clean_shutdown,omitempty" yaml:"clean_shutdown,omitempty"`
			RequestSpec   struct {
				ChangedFields []string `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
				Context       struct {
					UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
					ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
					DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
					UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
					ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
					AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
					UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
					ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
					DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
					SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
					UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
					ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
					IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
					IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
					ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
					ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
					ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
					Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
					ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
					ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
					ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
					ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
					ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
					ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
					ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
					ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
					ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
					ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
					RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
					GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
					ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
					RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
					Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
					ServiceCatalog           []struct {
						Type      string `json:"type,omitempty" yaml:"type,omitempty"`
						Name      string `json:"name,omitempty" yaml:"name,omitempty"`
						Endpoints []struct {
							Region      string `json:"region,omitempty" yaml:"region,omitempty"`
							InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
							PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
							AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
						} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
					} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
					QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
					DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
					MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
					CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
					UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
				} `json:"_context,omitempty" yaml:"_context,omitempty"`
				Version  string `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
				ObjID    int    `json:"_obj_id,omitempty" yaml:"_obj_id,omitempty"`
				ObjImage struct {
					ChangedFields []string `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
					Context       struct {
						UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
						ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
						DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
						UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
						ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
						AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
						UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
						ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
						DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
						SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
						UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
						ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
						IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
						IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
						ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
						ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
						ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
						Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
						ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
						ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
						ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
						ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
						ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
						ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
						RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
						GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
						ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
						RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
						Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
						ServiceCatalog           []struct {
							Type      string `json:"type,omitempty" yaml:"type,omitempty"`
							Name      string `json:"name,omitempty" yaml:"name,omitempty"`
							Endpoints []struct {
								Region      string `json:"region,omitempty" yaml:"region,omitempty"`
								InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
								PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
								AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
							} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
						} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
						QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
						DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
						MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
						CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
						UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
					} `json:"_context,omitempty" yaml:"_context,omitempty"`
					Version            string `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
					ObjID              string `json:"_obj_id,omitempty" yaml:"_obj_id,omitempty"`
					ObjName            string `json:"_obj_name,omitempty" yaml:"_obj_name,omitempty"`
					ObjStatus          string `json:"_obj_status,omitempty" yaml:"_obj_status,omitempty"`
					ObjChecksum        string `json:"_obj_checksum,omitempty" yaml:"_obj_checksum,omitempty"`
					ObjOwner           string `json:"_obj_owner,omitempty" yaml:"_obj_owner,omitempty"`
					ObjSize            int    `json:"_obj_size,omitempty" yaml:"_obj_size,omitempty"`
					ObjContainerFormat string `json:"_obj_container_format,omitempty" yaml:"_obj_container_format,omitempty"`
					ObjDiskFormat      string `json:"_obj_disk_format,omitempty" yaml:"_obj_disk_format,omitempty"`
					ObjCreatedAt       string `json:"_obj_created_at,omitempty" yaml:"_obj_created_at,omitempty"`
					ObjUpdatedAt       string `json:"_obj_updated_at,omitempty" yaml:"_obj_updated_at,omitempty"`
					ObjMinRAM          int    `json:"_obj_min_ram,omitempty" yaml:"_obj_min_ram,omitempty"`
					ObjMinDisk         int    `json:"_obj_min_disk,omitempty" yaml:"_obj_min_disk,omitempty"`
					ObjProperties      struct {
						ChangedFields       []string    `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
						Context             interface{} `json:"_context,omitempty" yaml:"_context,omitempty"`
						Version             string      `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
						ObjHwArchitecture   string      `json:"_obj_hw_architecture,omitempty" yaml:"_obj_hw_architecture,omitempty"`
						ObjHwDiskBus        string      `json:"_obj_hw_disk_bus,omitempty" yaml:"_obj_hw_disk_bus,omitempty"`
						ObjHwQemuGuestAgent bool        `json:"_obj_hw_qemu_guest_agent,omitempty" yaml:"_obj_hw_qemu_guest_agent,omitempty"`
						ObjHwRngModel       string      `json:"_obj_hw_rng_model,omitempty" yaml:"_obj_hw_rng_model,omitempty"`
						ObjHwScsiModel      string      `json:"_obj_hw_scsi_model,omitempty" yaml:"_obj_hw_scsi_model,omitempty"`
					} `json:"_obj_properties,omitempty" yaml:"_obj_properties,omitempty"`
				} `json:"_obj_image,omitempty" yaml:"_obj_image,omitempty"`
				ObjNumaTopology interface{} `json:"_obj_numa_topology,omitempty" yaml:"_obj_numa_topology,omitempty"`
				ObjPciRequests  struct {
					Requests []interface{} `json:"requests,omitempty" yaml:"requests,omitempty"`
				} `json:"_obj_pci_requests,omitempty" yaml:"_obj_pci_requests,omitempty"`
				ObjProjectID        string `json:"_obj_project_id,omitempty" yaml:"_obj_project_id,omitempty"`
				ObjUserID           string `json:"_obj_user_id,omitempty" yaml:"_obj_user_id,omitempty"`
				ObjAvailabilityZone string `json:"_obj_availability_zone,omitempty" yaml:"_obj_availability_zone,omitempty"`
				ObjFlavor           struct {
					ID          int     `json:"id,omitempty" yaml:"id,omitempty"`
					Name        string  `json:"name,omitempty" yaml:"name,omitempty"`
					MemoryMb    int     `json:"memory_mb,omitempty" yaml:"memory_mb,omitempty"`
					Vcpus       int     `json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
					RootGb      int     `json:"root_gb,omitempty" yaml:"root_gb,omitempty"`
					EphemeralGb int     `json:"ephemeral_gb,omitempty" yaml:"ephemeral_gb,omitempty"`
					Flavorid    string  `json:"flavorid,omitempty" yaml:"flavorid,omitempty"`
					Swap        int     `json:"swap,omitempty" yaml:"swap,omitempty"`
					RxtxFactor  float64 `json:"rxtx_factor,omitempty" yaml:"rxtx_factor,omitempty"`
					VcpuWeight  int     `json:"vcpu_weight,omitempty" yaml:"vcpu_weight,omitempty"`
					Disabled    bool    `json:"disabled,omitempty" yaml:"disabled,omitempty"`
					IsPublic    bool    `json:"is_public,omitempty" yaml:"is_public,omitempty"`
					ExtraSpecs  struct {
						HwCPUCores   string `json:"hw:cpu_cores,omitempty" yaml:"hw:cpu_cores,omitempty"`
						HwCPUSockets string `json:"hw:cpu_sockets,omitempty" yaml:"hw:cpu_sockets,omitempty"`
						HwRngAllowed string `json:"hw_rng:allowed,omitempty" yaml:"hw_rng:allowed,omitempty"`
					} `json:"extra_specs,omitempty" yaml:"extra_specs,omitempty"`
					Description interface{} `json:"description,omitempty" yaml:"description,omitempty"`
					CreatedAt   string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
					UpdatedAt   interface{} `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
					DeletedAt   interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
					Deleted     bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
				} `json:"_obj_flavor,omitempty" yaml:"_obj_flavor,omitempty"`
				ObjNumInstances         int         `json:"_obj_num_instances,omitempty" yaml:"_obj_num_instances,omitempty"`
				ObjIgnoreHosts          []string    `json:"_obj_ignore_hosts,omitempty" yaml:"_obj_ignore_hosts,omitempty"`
				ObjForceHosts           interface{} `json:"_obj_force_hosts,omitempty" yaml:"_obj_force_hosts,omitempty"`
				ObjForceNodes           interface{} `json:"_obj_force_nodes,omitempty" yaml:"_obj_force_nodes,omitempty"`
				ObjRequestedDestination struct {
					ChangedFields []string `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
					Context       struct {
						UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
						ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
						DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
						UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
						ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
						AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
						UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
						ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
						DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
						SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
						UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
						ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
						IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
						IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
						ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
						ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
						ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
						Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
						ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
						ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
						ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
						ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
						ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
						ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
						RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
						GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
						ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
						RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
						Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
						ServiceCatalog           []struct {
							Type      string `json:"type,omitempty" yaml:"type,omitempty"`
							Name      string `json:"name,omitempty" yaml:"name,omitempty"`
							Endpoints []struct {
								Region      string `json:"region,omitempty" yaml:"region,omitempty"`
								InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
								PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
								AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
							} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
						} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
						QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
						DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
						MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
						CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
						UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
					} `json:"_context,omitempty" yaml:"_context,omitempty"`
					Version string `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
					ObjCell struct {
						ChangedFields         []interface{} `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
						Context               interface{}   `json:"_context,omitempty" yaml:"_context,omitempty"`
						Version               string        `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
						ObjID                 int           `json:"_obj_id,omitempty" yaml:"_obj_id,omitempty"`
						ObjUUID               string        `json:"_obj_uuid,omitempty" yaml:"_obj_uuid,omitempty"`
						ObjName               string        `json:"_obj_name,omitempty" yaml:"_obj_name,omitempty"`
						ObjTransportURL       string        `json:"_obj_transport_url,omitempty" yaml:"_obj_transport_url,omitempty"`
						ObjDatabaseConnection string        `json:"_obj_database_connection,omitempty" yaml:"_obj_database_connection,omitempty"`
						ObjDisabled           bool          `json:"_obj_disabled,omitempty" yaml:"_obj_disabled,omitempty"`
						ObjCreatedAt          string        `json:"_obj_created_at,omitempty" yaml:"_obj_created_at,omitempty"`
						ObjUpdatedAt          interface{}   `json:"_obj_updated_at,omitempty" yaml:"_obj_updated_at,omitempty"`
					} `json:"_obj_cell,omitempty" yaml:"_obj_cell,omitempty"`
				} `json:"_obj_requested_destination,omitempty" yaml:"_obj_requested_destination,omitempty"`
				ObjRetry  interface{} `json:"_obj_retry,omitempty" yaml:"_obj_retry,omitempty"`
				ObjLimits struct {
					ChangedFields []string `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
					Context       struct {
						UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
						ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
						DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
						UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
						ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
						AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
						UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
						ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
						DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
						SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
						UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
						ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
						IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
						IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
						ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
						ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
						ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
						Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
						ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
						ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
						ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
						ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
						ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
						ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
						RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
						GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
						ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
						RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
						Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
						ServiceCatalog           []struct {
							Type      string `json:"type,omitempty" yaml:"type,omitempty"`
							Name      string `json:"name,omitempty" yaml:"name,omitempty"`
							Endpoints []struct {
								Region      string `json:"region,omitempty" yaml:"region,omitempty"`
								InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
								PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
								AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
							} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
						} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
						QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
						DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
						MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
						CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
						UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
					} `json:"_context,omitempty" yaml:"_context,omitempty"`
					Version         string      `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
					ObjNumaTopology interface{} `json:"_obj_numa_topology,omitempty" yaml:"_obj_numa_topology,omitempty"`
					ObjVcpu         interface{} `json:"_obj_vcpu,omitempty" yaml:"_obj_vcpu,omitempty"`
					ObjDiskGb       interface{} `json:"_obj_disk_gb,omitempty" yaml:"_obj_disk_gb,omitempty"`
					ObjMemoryMb     interface{} `json:"_obj_memory_mb,omitempty" yaml:"_obj_memory_mb,omitempty"`
				} `json:"_obj_limits,omitempty" yaml:"_obj_limits,omitempty"`
				ObjInstanceGroup  interface{} `json:"_obj_instance_group,omitempty" yaml:"_obj_instance_group,omitempty"`
				ObjSchedulerHints struct {
				} `json:"_obj_scheduler_hints,omitempty" yaml:"_obj_scheduler_hints,omitempty"`
				ObjInstanceUUID   string `json:"_obj_instance_uuid,omitempty" yaml:"_obj_instance_uuid,omitempty"`
				ObjSecurityGroups []struct {
					ChangedFields []string `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
					Context       struct {
						UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
						ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
						DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
						UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
						ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
						AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
						UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
						ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
						DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
						SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
						UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
						ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
						IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
						IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
						ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
						ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
						ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
						Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
						ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
						ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
						ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
						ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
						ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
						ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
						RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
						GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
						ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
						RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
						Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
						ServiceCatalog           []struct {
							Type      string `json:"type,omitempty" yaml:"type,omitempty"`
							Name      string `json:"name,omitempty" yaml:"name,omitempty"`
							Endpoints []struct {
								Region      string `json:"region,omitempty" yaml:"region,omitempty"`
								InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
								PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
								AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
							} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
						} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
						QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
						DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
						MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
						CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
						UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
					} `json:"_context,omitempty" yaml:"_context,omitempty"`
					Version string `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
					ObjUUID string `json:"_obj_uuid,omitempty" yaml:"_obj_uuid,omitempty"`
					ObjName string `json:"_obj_name,omitempty" yaml:"_obj_name,omitempty"`
				} `json:"_obj_security_groups,omitempty" yaml:"_obj_security_groups,omitempty"`
				ObjNetworkMetadata struct {
					ChangedFields []string `json:"_changed_fields,omitempty" yaml:"_changed_fields,omitempty"`
					Context       struct {
						UserID                   string        `json:"_user_id,omitempty" yaml:"_user_id,omitempty"`
						ProjectID                string        `json:"_project_id,omitempty" yaml:"_project_id,omitempty"`
						DomainID                 interface{}   `json:"_domain_id,omitempty" yaml:"_domain_id,omitempty"`
						UserDomainID             string        `json:"_user_domain_id,omitempty" yaml:"_user_domain_id,omitempty"`
						ProjectDomainID          string        `json:"_project_domain_id,omitempty" yaml:"_project_domain_id,omitempty"`
						AuthToken                string        `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
						UserName                 string        `json:"user_name,omitempty" yaml:"user_name,omitempty"`
						ProjectName              string        `json:"project_name,omitempty" yaml:"project_name,omitempty"`
						DomainName               interface{}   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
						SystemScope              interface{}   `json:"system_scope,omitempty" yaml:"system_scope,omitempty"`
						UserDomainName           interface{}   `json:"user_domain_name,omitempty" yaml:"user_domain_name,omitempty"`
						ProjectDomainName        interface{}   `json:"project_domain_name,omitempty" yaml:"project_domain_name,omitempty"`
						IsAdmin                  bool          `json:"is_admin,omitempty" yaml:"is_admin,omitempty"`
						IsAdminProject           bool          `json:"is_admin_project,omitempty" yaml:"is_admin_project,omitempty"`
						ReadOnly                 bool          `json:"read_only,omitempty" yaml:"read_only,omitempty"`
						ShowDeleted              bool          `json:"show_deleted,omitempty" yaml:"show_deleted,omitempty"`
						ResourceUUID             interface{}   `json:"resource_uuid,omitempty" yaml:"resource_uuid,omitempty"`
						Roles                    []string      `json:"roles,omitempty" yaml:"roles,omitempty"`
						ServiceToken             interface{}   `json:"service_token,omitempty" yaml:"service_token,omitempty"`
						ServiceUserID            interface{}   `json:"service_user_id,omitempty" yaml:"service_user_id,omitempty"`
						ServiceUserName          interface{}   `json:"service_user_name,omitempty" yaml:"service_user_name,omitempty"`
						ServiceUserDomainID      interface{}   `json:"service_user_domain_id,omitempty" yaml:"service_user_domain_id,omitempty"`
						ServiceUserDomainName    interface{}   `json:"service_user_domain_name,omitempty" yaml:"service_user_domain_name,omitempty"`
						ServiceProjectID         interface{}   `json:"service_project_id,omitempty" yaml:"service_project_id,omitempty"`
						ServiceProjectName       interface{}   `json:"service_project_name,omitempty" yaml:"service_project_name,omitempty"`
						ServiceProjectDomainID   interface{}   `json:"service_project_domain_id,omitempty" yaml:"service_project_domain_id,omitempty"`
						ServiceProjectDomainName interface{}   `json:"service_project_domain_name,omitempty" yaml:"service_project_domain_name,omitempty"`
						ServiceRoles             []interface{} `json:"service_roles,omitempty" yaml:"service_roles,omitempty"`
						RequestID                string        `json:"request_id,omitempty" yaml:"request_id,omitempty"`
						GlobalRequestID          interface{}   `json:"global_request_id,omitempty" yaml:"global_request_id,omitempty"`
						ReadDeleted              string        `json:"_read_deleted,omitempty" yaml:"_read_deleted,omitempty"`
						RemoteAddress            string        `json:"remote_address,omitempty" yaml:"remote_address,omitempty"`
						Timestamp                string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
						ServiceCatalog           []struct {
							Type      string `json:"type,omitempty" yaml:"type,omitempty"`
							Name      string `json:"name,omitempty" yaml:"name,omitempty"`
							Endpoints []struct {
								Region      string `json:"region,omitempty" yaml:"region,omitempty"`
								InternalURL string `json:"internalURL,omitempty" yaml:"internalURL,omitempty"`
								PublicURL   string `json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
								AdminURL    string `json:"adminURL,omitempty" yaml:"adminURL,omitempty"`
							} `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
						} `json:"service_catalog,omitempty" yaml:"service_catalog,omitempty"`
						QuotaClass     interface{} `json:"quota_class,omitempty" yaml:"quota_class,omitempty"`
						DbConnection   interface{} `json:"db_connection,omitempty" yaml:"db_connection,omitempty"`
						MqConnection   interface{} `json:"mq_connection,omitempty" yaml:"mq_connection,omitempty"`
						CellUUID       interface{} `json:"cell_uuid,omitempty" yaml:"cell_uuid,omitempty"`
						UserAuthPlugin interface{} `json:"user_auth_plugin,omitempty" yaml:"user_auth_plugin,omitempty"`
					} `json:"_context,omitempty" yaml:"_context,omitempty"`
					Version     string   `json:"VERSION,omitempty" yaml:"VERSION,omitempty"`
					ObjPhysnets []string `json:"_obj_physnets,omitempty" yaml:"_obj_physnets,omitempty"`
					ObjTunneled bool     `json:"_obj_tunneled,omitempty" yaml:"_obj_tunneled,omitempty"`
				} `json:"_obj_network_metadata,omitempty" yaml:"_obj_network_metadata,omitempty"`
				ObjIsBfv              bool          `json:"_obj_is_bfv,omitempty" yaml:"_obj_is_bfv,omitempty"`
				ObjRequestedResources []interface{} `json:"_obj_requested_resources,omitempty" yaml:"_obj_requested_resources,omitempty"`
			} `json:"request_spec,omitempty" yaml:"request_spec,omitempty"`
			Bdm struct {
				ID                  int         `json:"id,omitempty" yaml:"id,omitempty"`
				UUID                string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
				InstanceUUID        string      `json:"instance_uuid,omitempty" yaml:"instance_uuid,omitempty"`
				SourceType          string      `json:"source_type,omitempty" yaml:"source_type,omitempty"`
				DestinationType     string      `json:"destination_type,omitempty" yaml:"destination_type,omitempty"`
				GuestFormat         interface{} `json:"guest_format,omitempty" yaml:"guest_format,omitempty"`
				DeviceType          interface{} `json:"device_type,omitempty" yaml:"device_type,omitempty"`
				DiskBus             interface{} `json:"disk_bus,omitempty" yaml:"disk_bus,omitempty"`
				BootIndex           interface{} `json:"boot_index,omitempty" yaml:"boot_index,omitempty"`
				DeviceName          string      `json:"device_name,omitempty" yaml:"device_name,omitempty"`
				DeleteOnTermination bool        `json:"delete_on_termination,omitempty" yaml:"delete_on_termination,omitempty"`
				SnapshotID          interface{} `json:"snapshot_id,omitempty" yaml:"snapshot_id,omitempty"`
				VolumeID            string      `json:"volume_id,omitempty" yaml:"volume_id,omitempty"`
				VolumeSize          int         `json:"volume_size,omitempty" yaml:"volume_size,omitempty"`
				ImageID             interface{} `json:"image_id,omitempty" yaml:"image_id,omitempty"`
				NoDevice            bool        `json:"no_device,omitempty" yaml:"no_device,omitempty"`
				ConnectionInfo      string      `json:"connection_info,omitempty" yaml:"connection_info,omitempty"`
				Tag                 interface{} `json:"tag,omitempty" yaml:"tag,omitempty"`
				AttachmentID        string      `json:"attachment_id,omitempty" yaml:"attachment_id,omitempty"`
				VolumeType          interface{} `json:"volume_type,omitempty" yaml:"volume_type,omitempty"`
				CreatedAt           string      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
				UpdatedAt           string      `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
				DeletedAt           interface{} `json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
				Deleted             bool        `json:"deleted,omitempty" yaml:"deleted,omitempty"`
			} `json:"bdm,omitempty" yaml:"bdm,omitempty"`
		} `json:"args,omitempty" yaml:"args,omitempty"`
	} `json:"payload,omitempty" yaml:"payload,omitempty"`
}
