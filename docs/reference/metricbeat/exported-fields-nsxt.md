---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/exported-fields-nsxt.html
---

% This file is generated! See scripts/generate_fields_docs.py

# nsxt fields [exported-fields-nsxt]

nsxt module

## nsxt [_nsxt]



## health [_health]

Metrics retrieve via the NSX-T REST API

## backup_status [_backup_status]

Latest backup statuses.

## cluster [_cluster]

Latest cluster backup status.

**`nsxt.health.backup_status.cluster.backup_id`**
:   Cluster.backup id field.

type: keyword


**`nsxt.health.backup_status.cluster.start_time`**
:   Cluster.start time field.

type: keyword


**`nsxt.health.backup_status.cluster.end_time`**
:   Cluster.end time field.

type: keyword


**`nsxt.health.backup_status.cluster.success`**
:   Cluster.success field.

type: boolean


**`nsxt.health.backup_status.cluster.error_code`**
:   Cluster.error code field.

type: keyword


**`nsxt.health.backup_status.cluster.error_message`**
:   Cluster.error message field.

type: keyword


## node [_node]

Latest node backup status.

**`nsxt.health.backup_status.node.backup_id`**
:   Cluster.backup id field.

type: keyword


**`nsxt.health.backup_status.node.start_time`**
:   Cluster.start time field.

type: keyword


**`nsxt.health.backup_status.node.end_time`**
:   Cluster.end time field.

type: keyword


**`nsxt.health.backup_status.node.success`**
:   Cluster.success field.

type: boolean


**`nsxt.health.backup_status.node.error_code`**
:   Cluster.error code field.

type: keyword


**`nsxt.health.backup_status.node.error_message`**
:   Cluster.error message field.

type: keyword


## inventory [_inventory]

Latest inventory backup status.

**`nsxt.health.backup_status.inventory.backup_id`**
:   Cluster.backup id field.

type: keyword


**`nsxt.health.backup_status.inventory.start_time`**
:   Cluster.start time field.

type: keyword


**`nsxt.health.backup_status.inventory.end_time`**
:   Cluster.end time field.

type: keyword


**`nsxt.health.backup_status.inventory.success`**
:   Cluster.success field.

type: boolean


**`nsxt.health.backup_status.inventory.error_code`**
:   Cluster.error code field.

type: keyword


**`nsxt.health.backup_status.inventory.error_message`**
:   Cluster.error message field.

type: keyword


## cluster_node [_cluster_node]

Cluster node field.

**`nsxt.health.cluster_node.id`**
:   Id field.

type: keyword


**`nsxt.health.cluster_node.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.cluster_node.external_id`**
:   External id field.

type: keyword


**`nsxt.health.cluster_node.appliance_mgmt_listen_addr`**
:   Appliance mgmt listen addr field.

type: keyword


**`nsxt.health.cluster_node.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.cluster_node.manager_role`**
:   Manager role field.

type: keyword


**`nsxt.health.cluster_node.controller_role`**
:   Controller role field.

type: keyword


**`nsxt.health.cluster_node.create_time`**
:   Create time field.

type: date


**`nsxt.health.cluster_node.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.cluster_node.last_modified_time`**
:   Last modified time field.

type: date


**`nsxt.health.cluster_node.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.cluster_node.protection`**
:   Protection field.

type: keyword


**`nsxt.health.cluster_node.revision`**
:   Revision field.

type: keyword


**`nsxt.health.cluster_node.system_owned`**
:   System owned field.

type: keyword


## cluster_status [_cluster_status]

Cluster status field.

**`nsxt.health.cluster_status.cluster_id`**
:   Cluster id field.

type: keyword


**`nsxt.health.cluster_status.control_cluster_status`**
:   Control cluster status field.

type: keyword


**`nsxt.health.cluster_status.mgmt_cluster_status`**
:   Mgmt cluster status field.

type: keyword


**`nsxt.health.cluster_status.mgmt_cluster_node_count`**
:   Mgmt cluster node count field.

type: keyword


**`nsxt.health.cluster_status.group_id`**
:   Group id field.

type: keyword


**`nsxt.health.cluster_status.group_status`**
:   Group status field.

type: keyword


**`nsxt.health.cluster_status.group_type`**
:   Group type field.

type: keyword


**`nsxt.health.cluster_status.leaders`**
:   Leaders field.

type: keyword


## member_data [_member_data]

Member data field.

**`nsxt.health.cluster_status.member_data.fqdn`**
:   Member data.fqdn field.

type: keyword


**`nsxt.health.cluster_status.member_data.ip`**
:   Member data.ip field.

type: keyword


**`nsxt.health.cluster_status.member_data.status`**
:   Member data.status field.

type: keyword


**`nsxt.health.cluster_status.member_data.uuid`**
:   Member data.uuid field.

type: keyword


## edge_cluster [_edge_cluster]

Edge cluster field.

**`nsxt.health.edge_cluster.create_time`**
:   Create time field.

type: date


**`nsxt.health.edge_cluster.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.edge_cluster.last_modified_time`**
:   Last modified time field.

type: date


**`nsxt.health.edge_cluster.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.edge_cluster.protection`**
:   Protection field.

type: keyword


**`nsxt.health.edge_cluster.revision`**
:   Revision field.

type: keyword


**`nsxt.health.edge_cluster.system_owned`**
:   System owned field.

type: boolean


**`nsxt.health.edge_cluster.id`**
:   Id field.

type: keyword


**`nsxt.health.edge_cluster.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.edge_cluster.description`**
:   Description field.

type: keyword


**`nsxt.health.edge_cluster.deployment_type`**
:   Deployment type field.

type: keyword


**`nsxt.health.edge_cluster.enable_inter_site_forwarding`**
:   Enable inter site forwarding field.

type: keyword


**`nsxt.health.edge_cluster.member_node_type`**
:   Member node type field.

type: keyword


**`nsxt.health.edge_cluster.members`**
:   Members field.

type: keyword


**`nsxt.health.edge_cluster.cluster_profile_bindings`**
:   Cluster profile bindings field.

type: keyword


**`nsxt.health.edge_cluster.allocation_rules`**
:   Allocation rules field.

type: keyword


**`nsxt.health.edge_cluster.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.edge_cluster.tags`**
:   Tags field.

type: keyword


## firewall_section [_firewall_section]

Firewall section field.

**`nsxt.health.firewall_section.id`**
:   Id field.

type: keyword


**`nsxt.health.firewall_section.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.firewall_section.description`**
:   Description field.

type: keyword


**`nsxt.health.firewall_section.comments`**
:   Comments field.

type: keyword


**`nsxt.health.firewall_section.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.firewall_section.category`**
:   Category field.

type: keyword


**`nsxt.health.firewall_section.section_type`**
:   Section type field.

type: keyword


**`nsxt.health.firewall_section.enforced_on`**
:   Enforced on field.

type: keyword


**`nsxt.health.firewall_section.is_default`**
:   Is default field.

type: boolean


**`nsxt.health.firewall_section.locked`**
:   Locked field.

type: boolean


**`nsxt.health.firewall_section.lock_modified_by`**
:   Lock modified by field.

type: keyword


**`nsxt.health.firewall_section.lock_modified_time`**
:   Lock modified time field.

type: date


**`nsxt.health.firewall_section.stateful`**
:   Stateful field.

type: boolean


**`nsxt.health.firewall_section.tcp_strict`**
:   Tcp strict field.

type: boolean


**`nsxt.health.firewall_section.rule_count`**
:   Rule count field.

type: integer


**`nsxt.health.firewall_section.priority`**
:   Priority field.

type: keyword


**`nsxt.health.firewall_section.autoplumbed`**
:   Autoplumbed field.

type: keyword


**`nsxt.health.firewall_section.applied_tos`**
:   Applied tos field.

type: keyword


**`nsxt.health.firewall_section.tags`**
:   Tags field.

type: keyword


**`nsxt.health.firewall_section.create_time`**
:   Create time field.

type: date


**`nsxt.health.firewall_section.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.firewall_section.last_modified_time`**
:   Last modified time field.

type: date


**`nsxt.health.firewall_section.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.firewall_section.protection`**
:   Protection field.

type: keyword


**`nsxt.health.firewall_section.revision`**
:   Revision field.

type: integer


**`nsxt.health.firewall_section.system_owned`**
:   System owned field.

type: keyword


## ip_pool [_ip_pool]

Ip pool field.

**`nsxt.health.ip_pool.id`**
:   Id field.

type: keyword


**`nsxt.health.ip_pool.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.ip_pool.description`**
:   Description field.

type: keyword


## pool_usage [_pool_usage]

Pool usage field.

**`nsxt.health.ip_pool.pool_usage.allocated_ids`**
:   Pool usage.allocated ids field.

type: integer


**`nsxt.health.ip_pool.pool_usage.free_ids`**
:   Pool usage.free ids field.

type: integer


**`nsxt.health.ip_pool.pool_usage.total_ids`**
:   Pool usage.total ids field.

type: integer


**`nsxt.health.ip_pool.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.ip_pool.subnets`**
:   Subnets field.

type: keyword


**`nsxt.health.ip_pool.tags`**
:   Tags field.

type: keyword


**`nsxt.health.ip_pool.create_time`**
:   Create time field.

type: date


**`nsxt.health.ip_pool.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.ip_pool.last_modified_time`**
:   Last modified time field.

type: keyword


**`nsxt.health.ip_pool.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.ip_pool.protection`**
:   Protection field.

type: keyword


**`nsxt.health.ip_pool.revision`**
:   Revision field.

type: integer


**`nsxt.health.ip_pool.system_owned`**
:   System owned field.

type: boolean


## logical_router_port [_logical_router_port]

Logical router port field.

**`nsxt.health.logical_router_port.id`**
:   Id field.

type: keyword


**`nsxt.health.logical_router_port.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.logical_router_port.description`**
:   Description field.

type: keyword


**`nsxt.health.logical_router_port.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.logical_router_port.logical_router_id`**
:   Logical router id field.

type: keyword


**`nsxt.health.logical_router_port.mac_address`**
:   Mac address field.

type: keyword


**`nsxt.health.logical_router_port.subnets`**
:   Subnets field.

type: keyword


**`nsxt.health.logical_router_port.edge_cluster_member_index`**
:   Edge cluster member index field.

type: keyword


**`nsxt.health.logical_router_port.enable_multicast`**
:   Enable multicast field.

type: boolean


**`nsxt.health.logical_router_port.urpf_mode`**
:   Urpf mode field.

type: keyword


**`nsxt.health.logical_router_port.mode`**
:   Mode field.

type: keyword


**`nsxt.health.logical_router_port.mtu`**
:   Mtu field.

type: keyword


**`nsxt.health.logical_router_port.tags`**
:   Tags field.

type: keyword


**`nsxt.health.logical_router_port.service_bindings`**
:   Service bindings field.

type: keyword


**`nsxt.health.logical_router_port.linked_logical_router_port_id`**
:   Inked logical router port field.

type: keyword


**`nsxt.health.logical_router_port.inked_logical_switch_port_id`**
:   Inked logical router port id field.

type: keyword


## linked_logical_router_port [_linked_logical_router_port]

Linked logical router port field.

**`nsxt.health.logical_router_port.linked_logical_router_port.is_valid`**
:   Indicates if the linked logical router port is valid.

type: boolean


**`nsxt.health.logical_router_port.linked_logical_router_port.target_id`**
:   Target ID of the linked logical router port.

type: keyword


**`nsxt.health.logical_router_port.linked_logical_router_port.target_type`**
:   Target type of the linked logical router port.

type: keyword


**`nsxt.health.logical_router_port.linked_logical_router_port.target_display_name`**
:   Target display name of the linked logical router port.

type: keyword


## linked_logical_switch_port [_linked_logical_switch_port]

Linked logical router port field.

**`nsxt.health.logical_router_port.linked_logical_switch_port.is_valid`**
:   Indicates if the linked logical router port is valid.

type: boolean


**`nsxt.health.logical_router_port.linked_logical_switch_port.target_id`**
:   Target ID of the linked logical router port.

type: keyword


**`nsxt.health.logical_router_port.linked_logical_switch_port.target_type`**
:   Target type of the linked logical router port.

type: keyword


**`nsxt.health.logical_router_port.linked_logical_switch_port.target_display_name`**
:   Target display name of the linked logical router port.

type: keyword


## network_interface [_network_interface]

Network interface field.

**`nsxt.health.network_interface.schema`**
:   Schema field.

type: keyword


**`nsxt.health.network_interface.self`**
:   Self field.

type: keyword


**`nsxt.health.network_interface.admin_status`**
:   Admin status field.

type: keyword


**`nsxt.health.network_interface.broadcast_address`**
:   Broadcast address field.

type: keyword


**`nsxt.health.network_interface.default_gateway`**
:   Default gateway field.

type: keyword


**`nsxt.health.network_interface.interface_id`**
:   Interface id field.

type: keyword


**`nsxt.health.network_interface.ip_addresses`**
:   Ip addresses field.

type: keyword


**`nsxt.health.network_interface.ip_configuration`**
:   Ip configuration field.

type: keyword


**`nsxt.health.network_interface.link_status`**
:   Link status field.

type: keyword


**`nsxt.health.network_interface.mtu`**
:   Mtu field.

type: integer


**`nsxt.health.network_interface.physical_address`**
:   Physical address field.

type: keyword


## tier0 [_tier0]

Tier0 field.

**`nsxt.health.tier0.create_time`**
:   Create time field.

type: date


**`nsxt.health.tier0.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.tier0.last_modified_time`**
:   Last modified time field.

type: date


**`nsxt.health.tier0.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.tier0.protection`**
:   Protection field.

type: keyword


**`nsxt.health.tier0.revision`**
:   Revision field.

type: integer


**`nsxt.health.tier0.system_owned`**
:   System owned field.

type: boolean


**`nsxt.health.tier0.id`**
:   Id field.

type: keyword


**`nsxt.health.tier0.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.tier0.description`**
:   Description field.

type: keyword


**`nsxt.health.tier0.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.tier0.path`**
:   Path field.

type: keyword


**`nsxt.health.tier0.parent_path`**
:   Parent path field.

type: keyword


**`nsxt.health.tier0.relative_path`**
:   Relative path field.

type: keyword


**`nsxt.health.tier0.marked_for_delete`**
:   Marked for delete field.

type: boolean


**`nsxt.health.tier0.overridden`**
:   Overridden field.

type: boolean


**`nsxt.health.tier0.default_rule_logging`**
:   Default rule logging field.

type: boolean


**`nsxt.health.tier0.disable_firewall`**
:   Disable firewall field.

type: boolean


**`nsxt.health.tier0.force_whitelisting`**
:   Force whitelisting field.

type: boolean


**`nsxt.health.tier0.failover_mode`**
:   Failover mode field.

type: keyword


**`nsxt.health.tier0.ha_mode`**
:   Ha mode field.

type: keyword


**`nsxt.health.tier0.unique_id`**
:   Unique id field.

type: keyword


## advanced_config [_advanced_config]

Advanced config field.

**`nsxt.health.tier0.advanced_config.connectivity`**
:   Advanced config.connectivity field.

type: keyword


**`nsxt.health.tier0.advanced_config.forwarding_up_timer`**
:   Advanced config.forwarding up timer field.

type: integer


**`nsxt.health.tier0.internal_transit_subnets`**
:   Internal transit subnets field.

type: keyword


**`nsxt.health.tier0.transit_subnets`**
:   Transit subnets field.

type: keyword


**`nsxt.health.tier0.ipv6_profile_paths`**
:   Ipv6 profile paths field.

type: keyword


**`nsxt.health.tier0.tags`**
:   Tags field.

type: keyword


## tier1 [_tier1]

Tier1 field.

**`nsxt.health.tier1.create_time`**
:   Create time field.

type: date


**`nsxt.health.tier1.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.tier1.last_modified_time`**
:   Last modified time field.

type: date


**`nsxt.health.tier1.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.tier1.protection`**
:   Protection field.

type: keyword


**`nsxt.health.tier1.revision`**
:   Revision field.

type: integer


**`nsxt.health.tier1.system_owned`**
:   System owned field.

type: boolean


**`nsxt.health.tier1.id`**
:   Id field.

type: keyword


**`nsxt.health.tier1.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.tier1.description`**
:   Description field.

type: keyword


**`nsxt.health.tier1.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.tier1.path`**
:   Path field.

type: keyword


**`nsxt.health.tier1.parent_path`**
:   Parent path field.

type: keyword


**`nsxt.health.tier1.relative_path`**
:   Relative path field.

type: keyword


**`nsxt.health.tier1.marked_for_delete`**
:   Marked for delete field.

type: boolean


**`nsxt.health.tier1.overridden`**
:   Overridden field.

type: boolean


**`nsxt.health.tier1.default_rule_logging`**
:   Default rule logging field.

type: boolean


**`nsxt.health.tier1.disable_firewall`**
:   Disable firewall field.

type: boolean


**`nsxt.health.tier1.force_whitelisting`**
:   Force whitelisting field.

type: boolean


**`nsxt.health.tier1.failover_mode`**
:   Failover mode field.

type: keyword


**`nsxt.health.tier1.enable_standby_relocation`**
:   Enable standby relocation field.

type: boolean


**`nsxt.health.tier1.pool_allocation`**
:   Pool allocation field.

type: keyword


**`nsxt.health.tier1.ipv6_profile_paths`**
:   Ipv6 profile paths field.

type: keyword


**`nsxt.health.tier1.route_advertisement_types`**
:   Route advertisement types field.

type: keyword


**`nsxt.health.tier1.tier0_path`**
:   Tier0 path field.

type: keyword


**`nsxt.health.tier1.unique_id`**
:   Unique id field.

type: keyword


**`nsxt.health.tier1.tags`**
:   Tags field.

type: keyword


## transport_node [_transport_node]

Transport node field.

**`nsxt.health.transport_node.create_time`**
:   Create time field.

type: date


**`nsxt.health.transport_node.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.transport_node.last_modified_time`**
:   Last modified time field.

type: date


**`nsxt.health.transport_node.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.transport_node.protection`**
:   Protection field.

type: keyword


**`nsxt.health.transport_node.revision`**
:   Revision field.

type: integer


**`nsxt.health.transport_node.system_owned`**
:   System owned field.

type: boolean


**`nsxt.health.transport_node.id`**
:   Id field.

type: keyword


**`nsxt.health.transport_node.node_id`**
:   Node id field.

type: keyword


**`nsxt.health.transport_node.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.transport_node.description`**
:   Description field.

type: keyword


**`nsxt.health.transport_node.failure_domain_id`**
:   Failure domain id field.

type: keyword


**`nsxt.health.transport_node.is_overridden`**
:   Is overridden field.

type: boolean


**`nsxt.health.transport_node.maintenance_mode`**
:   Maintenance mode field.

type: keyword


**`nsxt.health.transport_node.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.transport_node.tags`**
:   Tags field.

type: keyword


**`nsxt.health.transport_node.transport_zone_endpoints`**
:   Transport zone endpoints field.

type: keyword


## node_deployment_info [_node_deployment_info]

Node deployment info field.

**`nsxt.health.transport_node.node_deployment_info.create_time`**
:   Node deployment info create time field.

type: date


**`nsxt.health.transport_node.node_deployment_info.create_user`**
:   Node deployment info create user field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.last_modified_time`**
:   Node deployment info last modified time field.

type: date


**`nsxt.health.transport_node.node_deployment_info.last_modified_user`**
:   Node deployment info last modified user field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.protection`**
:   Node deployment info protection field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.revision`**
:   Node deployment info revision field.

type: integer


**`nsxt.health.transport_node.node_deployment_info.system_owned`**
:   Node deployment info system owned field.

type: boolean


**`nsxt.health.transport_node.node_deployment_info.resource_type`**
:   Node deployment info resource type field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.deployment_type`**
:   Node deployment info deployment type field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.deployment_config`**
:   Node deployment info deployment config field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.display_name`**
:   Node deployment info display name field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.description`**
:   Node deployment info description field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.external_id`**
:   Node deployment info external id field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.id`**
:   Node deployment info id field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.ip_addresses`**
:   Node deployment info ip addresses field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.node_settings`**
:   Node deployment info node settings field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.discovered_node_id`**
:   Node deployment info discovered node id field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.fqdn`**
:   Node deployment info fqdn field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.managed_by_server`**
:   Node deployment info managed by server field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.discovered_ip_addresses`**
:   Node deployment info discovered ip addresses field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.os_type`**
:   Node deployment info os type field.

type: keyword


**`nsxt.health.transport_node.node_deployment_info.os_version`**
:   Node deployment info os version field.

type: keyword


## host_switch [_host_switch]

Host switch field.

**`nsxt.health.transport_node.host_switch.cpu_config`**
:   CPU configuration.

type: keyword


**`nsxt.health.transport_node.host_switch.host_switch_id`**
:   Host switch ID.

type: keyword


**`nsxt.health.transport_node.host_switch.host_switch_mode`**
:   Host switch mode.

type: keyword


**`nsxt.health.transport_node.host_switch.host_switch_name`**
:   Host switch name.

type: keyword


**`nsxt.health.transport_node.host_switch.host_switch_profile_ids`**
:   Host switch profile IDs.

type: keyword


**`nsxt.health.transport_node.host_switch.host_switch_type`**
:   Host switch type.

type: keyword


## ip_assignment_spec [_ip_assignment_spec]

IP assignment specification.

**`nsxt.health.transport_node.host_switch.ip_assignment_spec.resource_type`**
:   Resource type of the IP assignment specification.

type: keyword


**`nsxt.health.transport_node.host_switch.ip_assignment_spec.ip_pool_id`**
:   IP pool ID of the IP assignment specification.

type: keyword


**`nsxt.health.transport_node.host_switch.ip_assignment_spec.default_gateway`**
:   Default gateway of the IP assignment specification.

type: keyword


**`nsxt.health.transport_node.host_switch.ip_assignment_spec.ip_list`**
:   IP list of the IP assignment specification.

type: keyword


**`nsxt.health.transport_node.host_switch.ip_assignment_spec.subnet_mask`**
:   Subnet mask of the IP assignment specification.

type: keyword


**`nsxt.health.transport_node.host_switch.is_migrate_pnics`**
:   Indicates if PNICS are migrated.

type: boolean


**`nsxt.health.transport_node.host_switch.not_ready`**
:   Not ready flag.

type: boolean


**`nsxt.health.transport_node.host_switch.pnics`**
:   PNICS.

type: keyword


**`nsxt.health.transport_node.host_switch.pnics_uninstall_migration`**
:   PNICS uninstall migration.

type: keyword


**`nsxt.health.transport_node.host_switch.vmk_install_migration`**
:   VMK install migration.

type: keyword


**`nsxt.health.transport_node.host_switch.vmk_uninstall_migration`**
:   VMK uninstall migration.

type: keyword


**`nsxt.health.transport_node.host_switch.transport_zone_endpoints`**
:   Transport zone endpoints.

type: keyword


**`nsxt.health.transport_node.host_switch.uplinks`**
:   Uplinks.

type: keyword


## transport_zone [_transport_zone]

Transport zone field.

**`nsxt.health.transport_zone.create_time`**
:   Create time field.

type: date


**`nsxt.health.transport_zone.create_user`**
:   Create user field.

type: keyword


**`nsxt.health.transport_zone.last_modified_time`**
:   Last modified time field.

type: date


**`nsxt.health.transport_zone.last_modified_user`**
:   Last modified user field.

type: keyword


**`nsxt.health.transport_zone.protection`**
:   Protection field.

type: keyword


**`nsxt.health.transport_zone.revision`**
:   Revision field.

type: integer


**`nsxt.health.transport_zone.schema`**
:   Schema field.

type: keyword


**`nsxt.health.transport_zone.system_owned`**
:   System owned field.

type: boolean


**`nsxt.health.transport_zone.id`**
:   Id field.

type: keyword


**`nsxt.health.transport_zone.display_name`**
:   Display name field.

type: keyword


**`nsxt.health.transport_zone.host_switch_id`**
:   Host switch id field.

type: keyword


**`nsxt.health.transport_zone.host_switch_name`**
:   Host switch name field.

type: keyword


**`nsxt.health.transport_zone.host_switch_mode`**
:   Host switch mode field.

type: keyword


**`nsxt.health.transport_zone.is_default`**
:   Is default field.

type: boolean


**`nsxt.health.transport_zone.nested_nsx`**
:   Nested nsx field.

type: boolean


**`nsxt.health.transport_zone.resource_type`**
:   Resource type field.

type: keyword


**`nsxt.health.transport_zone.transport_type`**
:   Transport type field.

type: keyword


**`nsxt.health.transport_zone.tags`**
:   Tags field.

type: keyword


**`nsxt.health.transport_zone.transport_zone_profile_ids`**
:   Transport zone profile ids field.

type: keyword


**`nsxt.health.transport_zone.uplink_teaming_policy_names`**
:   Uplink teaming policy names field.

type: keyword


