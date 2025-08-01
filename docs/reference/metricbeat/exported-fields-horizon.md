---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/exported-fields-horizon.html
---

% This file is generated! See scripts/generate_fields_docs.py

# Horizon module fields [exported-fields-horizon]

Horizon module collects metrics from VMware Horizon components.

## horizon [_horizon]



## health [_health]

health

## connection_server [_connection_server]

Horizon connection server information

**`horizon.health.connection_server.name`**
:   type: keyword


**`horizon.health.connection_server.status`**
:   type: keyword


**`horizon.health.connection_server.connection_count`**
:   type: long


**`horizon.health.connection_server.tunnel_connection_count`**
:   type: long


**`horizon.health.connection_server.default_certificate`**
:   type: boolean


## certificate [_certificate]

Certificate info

**`horizon.health.connection_server.certificate.valid`**
:   type: boolean


**`horizon.health.connection_server.certificate.valid_from`**
:   type: date


**`horizon.health.connection_server.certificate.valid_to`**
:   type: date


## details [_details]

Details about the connection server

**`horizon.health.connection_server.details.version`**
:   type: keyword


**`horizon.health.connection_server.details.build`**
:   type: keyword


## desktop_pool [_desktop_pool]

Horizon desktop pool information

**`horizon.health.desktop_pool.id`**
:   type: keyword


**`horizon.health.desktop_pool.name`**
:   type: keyword


**`horizon.health.desktop_pool.display_name`**
:   type: keyword


**`horizon.health.desktop_pool.description`**
:   type: keyword


**`horizon.health.desktop_pool.type`**
:   type: keyword


**`horizon.health.desktop_pool.source`**
:   type: keyword


**`horizon.health.desktop_pool.enabled`**
:   type: boolean


## settings [_settings]

Settings for the pool

**`horizon.health.desktop_pool.settings.delete_in_progress`**
:   type: boolean


**`horizon.health.desktop_pool.settings.enable_client_restrictions`**
:   type: boolean


**`horizon.health.desktop_pool.settings.allow_multiple_sessions_per_user`**
:   type: boolean


**`horizon.health.desktop_pool.settings.session_type`**
:   type: keyword


**`horizon.health.desktop_pool.settings.cloud_managed`**
:   type: boolean


**`horizon.health.desktop_pool.settings.cloud_assigned`**
:   type: boolean


## session_settings [_session_settings]

Session settings

**`horizon.health.desktop_pool.settings.session_settings.power_policy`**
:   type: keyword


**`horizon.health.desktop_pool.settings.session_settings.disconnected_session_timeout_policy`**
:   type: keyword


**`horizon.health.desktop_pool.settings.session_settings.disconnected_session_timeout_minutes`**
:   type: integer


**`horizon.health.desktop_pool.settings.session_settings.allow_users_to_reset_machines`**
:   type: boolean


**`horizon.health.desktop_pool.settings.session_settings.allow_multiple_sessions_per_user`**
:   type: boolean


**`horizon.health.desktop_pool.settings.session_settings.delete_or_refresh_machine_after_logoff`**
:   type: keyword


**`horizon.health.desktop_pool.settings.session_settings.refresh_os_disk_after_logoff`**
:   type: keyword


## display_protocol_settings [_display_protocol_settings]

Protocol settings

**`horizon.health.desktop_pool.settings.display_protocol_settings.display_protocols`**
:   type: keyword


**`horizon.health.desktop_pool.settings.display_protocol_settings.default_display_protocol`**
:   type: keyword


**`horizon.health.desktop_pool.settings.display_protocol_settings.allow_users_to_choose_protocol`**
:   type: boolean


**`horizon.health.desktop_pool.settings.display_protocol_settings.html_access_enabled`**
:   type: boolean


**`horizon.health.desktop_pool.settings.display_protocol_settings.session_collaboration_enabled`**
:   type: boolean


**`horizon.health.desktop_pool.settings.display_protocol_settings.renderer3d`**
:   type: keyword


**`horizon.health.desktop_pool.settings.display_protocol_settings.grid_vgpus_enabled`**
:   type: boolean


**`horizon.health.desktop_pool.settings.display_protocol_settings.max_number_of_monitors`**
:   type: integer


**`horizon.health.desktop_pool.settings.display_protocol_settings.max_resolution_of_any_one_monitor`**
:   type: keyword


## installed_application [_installed_application]

Application installed in a desktop pool

**`horizon.health.desktop_pool.installed_application.name`**
:   type: keyword


**`horizon.health.desktop_pool.installed_application.version`**
:   type: keyword


**`horizon.health.desktop_pool.installed_application.publisher`**
:   type: keyword


**`horizon.health.desktop_pool.installed_application.executable_path`**
:   type: keyword


**`horizon.health.desktop_pool.installed_application.file_types`**
:   type: keyword


## session [_session]

Horizon session information

**`horizon.health.session.id`**
:   type: keyword


**`horizon.health.session.user_id`**
:   type: keyword


**`horizon.health.session.broker_user_id`**
:   type: keyword


**`horizon.health.session.access_group_id`**
:   type: keyword


**`horizon.health.session.machine_id`**
:   type: keyword


**`horizon.health.session.desktop_pool_id`**
:   type: keyword


**`horizon.health.session.agent_version`**
:   type: keyword


## client_data [_client_data]

Client data

**`horizon.health.session.client_data.location_id`**
:   type: keyword


**`horizon.health.session.client_data.type`**
:   type: keyword


**`horizon.health.session.client_data.address`**
:   type: keyword


**`horizon.health.session.client_data.name`**
:   type: keyword


**`horizon.health.session.client_data.version`**
:   type: keyword


## security_gateway_data [_security_gateway_data]

Gateway info

**`horizon.health.session.security_gateway_data.domain_name`**
:   type: keyword


**`horizon.health.session.security_gateway_data.address`**
:   type: keyword


**`horizon.health.session.security_gateway_data.location`**
:   type: keyword


**`horizon.health.session.session_type`**
:   type: keyword


**`horizon.health.session.session_protocol`**
:   type: keyword


**`horizon.health.session.session_state`**
:   type: keyword


**`horizon.health.session.start_time`**
:   type: date


**`horizon.health.session.disconnected_time`**
:   type: date


**`horizon.health.session.last_session_duration`**
:   type: long


**`horizon.health.session.resourced_remotely`**
:   type: boolean


**`horizon.health.session.unauthenticated`**
:   type: boolean


**`horizon.health.session.idle_duration`**
:   type: long


## gateway [_gateway]

Horizon gateway info

**`horizon.health.gateway.id`**
:   type: keyword


**`horizon.health.gateway.name`**
:   type: keyword


**`horizon.health.gateway.status`**
:   type: keyword


**`horizon.health.gateway.active_connection_count`**
:   type: long


**`horizon.health.gateway.pcoip_connection_count`**
:   type: long


**`horizon.health.gateway.blast_connection_count`**
:   type: long


## details [_details]

Gateway details

**`horizon.health.gateway.details.type`**
:   type: keyword


**`horizon.health.gateway.details.address`**
:   type: keyword


**`horizon.health.gateway.details.internal`**
:   type: boolean


**`horizon.health.gateway.details.version`**
:   type: keyword


## virtual_center [_virtual_center]

vCenter information

**`horizon.health.virtual_center.id`**
:   type: keyword


**`horizon.health.virtual_center.version`**
:   type: keyword


**`horizon.health.virtual_center.description`**
:   type: keyword


**`horizon.health.virtual_center.instance_uuid`**
:   type: keyword


**`horizon.health.virtual_center.server_name`**
:   type: keyword


**`horizon.health.virtual_center.port`**
:   type: integer


**`horizon.health.virtual_center.use_ssl`**
:   type: boolean


**`horizon.health.virtual_center.user_name`**
:   type: keyword


**`horizon.health.virtual_center.se_sparse_reclamation_enabled`**
:   type: boolean


**`horizon.health.virtual_center.enabled`**
:   type: boolean


**`horizon.health.virtual_center.vmc_deployment`**
:   type: boolean


## limits [_limits]

Resource limits

**`horizon.health.virtual_center.limits.provisioning_limit`**
:   type: integer


**`horizon.health.virtual_center.limits.power_operations_limit`**
:   type: integer


**`horizon.health.virtual_center.limits.instant_clone_engine_provisioning_limit`**
:   type: integer


## storage_accelerator_data [_storage_accelerator_data]

Storage cache

**`horizon.health.virtual_center.storage_accelerator_data.enabled`**
:   type: boolean


**`horizon.health.virtual_center.storage_accelerator_data.default_cache_size_mb`**
:   type: integer


## certificate_override [_certificate_override]

Certificate override

**`horizon.health.virtual_center.certificate_override.certificate`**
:   type: keyword


**`horizon.health.virtual_center.certificate_override.type`**
:   type: keyword


## machine [_machine]

Machine info

**`horizon.health.machine.id`**
:   type: keyword


**`horizon.health.machine.name`**
:   type: keyword


**`horizon.health.machine.dns_name`**
:   type: keyword


**`horizon.health.machine.desktop_pool_id`**
:   type: keyword


**`horizon.health.machine.state`**
:   type: keyword


**`horizon.health.machine.type`**
:   type: keyword


**`horizon.health.machine.operating_system`**
:   type: keyword


**`horizon.health.machine.operating_system_architecture`**
:   type: keyword


**`horizon.health.machine.agent_version`**
:   type: keyword


**`horizon.health.machine.agent_build_number`**
:   type: keyword


**`horizon.health.machine.remote_experience_agent_build_number`**
:   type: keyword


**`horizon.health.machine.message_security_mode`**
:   type: keyword


**`horizon.health.machine.message_security_enhanced_mode_supported`**
:   type: boolean


**`horizon.health.machine.pairing_state`**
:   type: keyword


**`horizon.health.machine.configured_by_connection_server`**
:   type: keyword


**`horizon.health.machine.user_ids`**
:   type: keyword


## managed_machine_data [_managed_machine_data]

Managed machine data info

**`horizon.health.machine.managed_machine_data.virtual_center_id`**
:   type: keyword


**`horizon.health.machine.managed_machine_data.host_name`**
:   type: keyword


**`horizon.health.machine.managed_machine_data.path`**
:   type: keyword


**`horizon.health.machine.managed_machine_data.virtual_machine_power_state`**
:   type: keyword


**`horizon.health.machine.managed_machine_data.storage_accelerator_state`**
:   type: keyword


**`horizon.health.machine.managed_machine_data.memory_mb`**
:   type: integer


**`horizon.health.machine.managed_machine_data.missing_in_vcenter`**
:   type: boolean


**`horizon.health.machine.managed_machine_data.in_hold_customization`**
:   type: boolean


**`horizon.health.machine.managed_machine_data.create_time`**
:   type: date


**`horizon.health.machine.managed_machine_data.in_maintenance_mode`**
:   type: boolean


## virtual_disk [_virtual_disk]

Virtual disks info

**`horizon.health.machine.virtual_disk.path`**
:   type: keyword


**`horizon.health.machine.virtual_disk.datastore_path`**
:   type: keyword


**`horizon.health.machine.virtual_disk.capacity_mb`**
:   type: integer


## rds_server [_rds_server]

RDS server information

**`horizon.health.rds_server.id`**
:   type: keyword


**`horizon.health.rds_server.name`**
:   type: keyword


**`horizon.health.rds_server.dns_name`**
:   type: keyword


**`horizon.health.rds_server.desktop_pool_id`**
:   type: keyword


**`horizon.health.rds_server.state`**
:   type: keyword


**`horizon.health.rds_server.type`**
:   type: keyword


**`horizon.health.rds_server.operating_system`**
:   type: keyword


**`horizon.health.rds_server.operating_system_architecture`**
:   type: keyword


**`horizon.health.rds_server.agent_version`**
:   type: keyword


**`horizon.health.rds_server.agent_build_number`**
:   type: keyword


**`horizon.health.rds_server.remote_experience_agent_build_number`**
:   type: keyword


**`horizon.health.rds_server.message_security_mode`**
:   type: keyword


**`horizon.health.rds_server.message_security_enhanced_mode_supported`**
:   type: boolean


**`horizon.health.rds_server.pairing_state`**
:   type: keyword


**`horizon.health.rds_server.configured_by_connection_server`**
:   type: keyword


**`horizon.health.rds_server.user_ids`**
:   type: keyword


## managed_machine_data [_managed_machine_data]

Managed machine data info

**`horizon.health.rds_server.managed_machine_data.virtual_center_id`**
:   type: keyword


**`horizon.health.rds_server.managed_machine_data.host_name`**
:   type: keyword


**`horizon.health.rds_server.managed_machine_data.path`**
:   type: keyword


**`horizon.health.rds_server.managed_machine_data.virtual_machine_power_state`**
:   type: keyword


**`horizon.health.rds_server.managed_machine_data.storage_accelerator_state`**
:   type: keyword


**`horizon.health.rds_server.managed_machine_data.memory_mb`**
:   type: integer


**`horizon.health.rds_server.managed_machine_data.missing_in_vcenter`**
:   type: boolean


**`horizon.health.rds_server.managed_machine_data.in_hold_customization`**
:   type: boolean


**`horizon.health.rds_server.managed_machine_data.create_time`**
:   type: date


**`horizon.health.rds_server.managed_machine_data.in_maintenance_mode`**
:   type: boolean


## virtual_disks [_virtual_disks]

Virtual disks info

**`horizon.health.rds_server.managed_machine_data.virtual_disks.path`**
:   type: keyword


**`horizon.health.rds_server.managed_machine_data.virtual_disks.datastore_path`**
:   type: keyword


**`horizon.health.rds_server.managed_machine_data.virtual_disks.capacity_mb`**
:   type: integer


## farm [_farm]

Horizon farm information

**`horizon.health.farm.id`**
:   type: keyword


**`horizon.health.farm.name`**
:   type: keyword


**`horizon.health.farm.display_name`**
:   type: keyword


**`horizon.health.farm.description`**
:   type: keyword


**`horizon.health.farm.enabled`**
:   type: boolean


**`horizon.health.farm.source`**
:   type: keyword


**`horizon.health.farm.type`**
:   type: keyword


## settings [_settings]

Settings for the farm

**`horizon.health.farm.settings.delete_in_progress`**
:   type: boolean


**`horizon.health.farm.settings.desktop_id`**
:   type: keyword


## display_protocol_settings [_display_protocol_settings]

Display protocol settings

**`horizon.health.farm.settings.display_protocol_settings.allow_display_protocol_override`**
:   type: boolean


**`horizon.health.farm.settings.display_protocol_settings.default_display_protocol`**
:   type: keyword


**`horizon.health.farm.settings.display_protocol_settings.grid_vgpus_enabled`**
:   type: boolean


**`horizon.health.farm.settings.display_protocol_settings.html_access_enabled`**
:   type: boolean


**`horizon.health.farm.settings.display_protocol_settings.session_collaboration_enabled`**
:   type: boolean


**`horizon.health.farm.settings.display_protocol_settings.vgpu_grid_profile`**
:   type: keyword


## load_balancer_settings [_load_balancer_settings]

Load balancer settings

**`horizon.health.farm.settings.load_balancer_settings.custom_script_in_use`**
:   type: boolean


## lb_metric_settings [_lb_metric_settings]

Load balancer metric settings

**`horizon.health.farm.settings.load_balancer_settings.lb_metric_settings.cpu_threshold`**
:   type: integer


**`horizon.health.farm.settings.load_balancer_settings.lb_metric_settings.disk_queue_length_threshold`**
:   type: integer


**`horizon.health.farm.settings.load_balancer_settings.lb_metric_settings.disk_read_latency_threshold`**
:   type: integer


**`horizon.health.farm.settings.load_balancer_settings.lb_metric_settings.disk_write_latency_threshold`**
:   type: integer


**`horizon.health.farm.settings.load_balancer_settings.lb_metric_settings.include_session_count`**
:   type: boolean


**`horizon.health.farm.settings.load_balancer_settings.lb_metric_settings.memory_threshold`**
:   type: integer


**`horizon.health.farm.settings.server_error_threshold`**
:   type: integer


## session_settings [_session_settings]

Session settings

**`horizon.health.farm.settings.session_settings.disconnected_session_timeout_minutes`**
:   type: integer


**`horizon.health.farm.settings.session_settings.disconnected_session_timeout_policy`**
:   type: keyword


**`horizon.health.farm.settings.session_settings.empty_session_timeout_minutes`**
:   type: integer


**`horizon.health.farm.settings.session_settings.empty_session_timeout_policy`**
:   type: keyword


**`horizon.health.farm.settings.session_settings.logoff_after_timeout`**
:   type: boolean


**`horizon.health.farm.settings.session_settings.pre_launch_session_timeout_minutes`**
:   type: integer


**`horizon.health.farm.settings.session_settings.pre_launch_session_timeout_policy`**
:   type: keyword


## certificate_data [_certificate_data]

Certificate data

**`horizon.health.certificate_data.certificate_usage`**
:   type: keyword


**`horizon.health.certificate_data.dns_subject_alternative_names`**
:   type: keyword


**`horizon.health.certificate_data.in_use`**
:   type: boolean


**`horizon.health.certificate_data.invalid_reasons`**
:   type: keyword


**`horizon.health.certificate_data.is_valid`**
:   type: boolean


**`horizon.health.certificate_data.issuer_name`**
:   type: keyword


**`horizon.health.certificate_data.serial_number`**
:   type: keyword


**`horizon.health.certificate_data.sha1_thumbprint`**
:   type: keyword


**`horizon.health.certificate_data.signature_algorithm`**
:   type: keyword


**`horizon.health.certificate_data.subject_name`**
:   type: keyword


**`horizon.health.certificate_data.valid_from`**
:   type: date


**`horizon.health.certificate_data.valid_until`**
:   type: date


## license_data [_license_data]

License data

**`horizon.health.license_data.application_pool_launch_enabled`**
:   type: boolean


**`horizon.health.license_data.desktop_pool_launch_enabled`**
:   type: boolean


**`horizon.health.license_data.expiration_time`**
:   type: long


**`horizon.health.license_data.grace_period_days`**
:   type: integer


**`horizon.health.license_data.help_desk_enabled`**
:   type: boolean


**`horizon.health.license_data.instant_clone_enabled`**
:   type: boolean


**`horizon.health.license_data.license_edition`**
:   type: keyword


**`horizon.health.license_data.license_health`**
:   type: keyword


**`horizon.health.license_data.license_key`**
:   type: keyword


**`horizon.health.license_data.license_mode`**
:   type: keyword


**`horizon.health.license_data.licensed`**
:   type: boolean


**`horizon.health.license_data.session_collaboration_enabled`**
:   type: boolean


**`horizon.health.license_data.subscription_slice_expiry`**
:   type: long


**`horizon.health.license_data.usage_model`**
:   type: keyword


