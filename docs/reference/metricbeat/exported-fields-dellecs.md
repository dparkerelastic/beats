---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/exported-fields-dellecs.html
---

% This file is generated! See scripts/generate_fields_docs.py

# dellecs fields [exported-fields-dellecs]

dellecs module

## dellecs [_dellecs]

Dell EMC ECS metrics

## health [_health]

This module collects health metrics from Dell EMC ECS.

## disks [_disks]

Information about disks

**`dellecs.health.disks.disk_id`**
:   Unique identifier for the disk

type: keyword


**`dellecs.health.disks.node_id`**
:   Unique identifier for the node

type: keyword


**`dellecs.health.disks.storage_pool_name`**
:   Name of the storage pool

type: keyword


**`dellecs.health.disks.display_name`**
:   Display name of the disk

type: keyword


**`dellecs.health.disks.node_display_name`**
:   Display name of the node

type: keyword


**`dellecs.health.disks.slot_id`**
:   Slot identifier for the disk

type: keyword


**`dellecs.health.disks.storage_pool_id`**
:   Unique identifier for the storage pool

type: keyword


**`dellecs.health.disks.ssm_l2_status`**
:   SSM level 2 status

type: keyword


**`dellecs.health.disks.health_status`**
:   Health status of the disk

type: keyword


**`dellecs.health.disks.ssm_l1_status`**
:   SSM level 1 status

type: keyword


**`dellecs.health.disks.disk_space_free_current_l1`**
:   Current free disk space at level 1

type: long


**`dellecs.health.disks.disk_space_free_current_l2`**
:   Current free disk space at level 2

type: long


**`dellecs.health.disks.disk_space_free_l2`**
:   Free disk space at level 2

type: long


**`dellecs.health.disks.disk_space_free_l1`**
:   Free disk space at level 1

type: long


**`dellecs.health.disks.disk_space_total`**
:   Total disk space

type: long


**`dellecs.health.disks.disk_space_total_current_l1`**
:   Current total disk space at level 1

type: long


**`dellecs.health.disks.disk_space_total_current_l2`**
:   Current total disk space at level 2

type: long


**`dellecs.health.disks.disk_space_allocated_l2`**
:   Allocated disk space at level 2

type: long


**`dellecs.health.disks.disk_space_allocated_l1`**
:   Allocated disk space at level 1

type: long


**`dellecs.health.disks.disk_space_free_current`**
:   Current free disk space

type: long


**`dellecs.health.disks.disk_space_total_current`**
:   Current total disk space

type: long


**`dellecs.health.disks.disk_space_allocated_current`**
:   Current allocated disk space

type: long


**`dellecs.health.disks.disk_space_allocated_current_l1`**
:   Current allocated disk space at level 1

type: long


**`dellecs.health.disks.disk_space_allocated`**
:   Allocated disk space

type: long


**`dellecs.health.disks.disk_space_allocated_current_l2`**
:   Current allocated disk space at level 2

type: long


**`dellecs.health.disks.disk_space_free`**
:   Free disk space

type: long


**`dellecs.health.disks.disk_space_allocated_percent`**
:   Percentage of allocated disk space

type: float


**`dellecs.health.disks.disk_space_allocated_percentage_current`**
:   Current percentage of allocated disk space

type: float


## capacity [_capacity]

Capacity details

**`dellecs.health.capacity.total_provisioned_gb`**
:   Total provisioned capacity in GB

type: long


**`dellecs.health.capacity.total_free_gb`**
:   Total free capacity in GB

type: long


## storage_pool [_storage_pool]

Information about storage pools

**`dellecs.health.storage_pool.pool_id`**
:   Unique identifier for the storage pool

type: keyword


**`dellecs.health.storage_pool.status`**
:   Status of the storage pool

type: keyword


**`dellecs.health.storage_pool.chunks_l1_journal_total_size`**
:   Total size of L1 journal chunks

type: long


**`dellecs.health.storage_pool.chunks_l1_btree_number`**
:   Number of L1 B-tree chunks

type: long


**`dellecs.health.storage_pool.chunks_l1_btree_total_size`**
:   Total size of L1 B-tree chunks

type: long


**`dellecs.health.storage_pool.chunks_l1_journal_avg_size`**
:   Average size of L1 journal chunks

type: long


**`dellecs.health.storage_pool.chunks_l1_journal_number`**
:   Number of L1 journal chunks

type: long


**`dellecs.health.storage_pool.chunks_l0_journal_total_size`**
:   Total size of L0 journal chunks

type: long


**`dellecs.health.storage_pool.chunks_l0_btree_number`**
:   Number of L0 B-tree chunks

type: long


**`dellecs.health.storage_pool.chunks_l0_btree_total_size`**
:   Total size of L0 B-tree chunks

type: long


**`dellecs.health.storage_pool.chunks_l0_btree_avg_size`**
:   Average size of L0 B-tree chunks

type: long


**`dellecs.health.storage_pool.chunks_l0_journal_avg_size`**
:   Average size of L0 journal chunks

type: long


**`dellecs.health.storage_pool.chunks_l0_journal_number`**
:   Number of L0 journal chunks

type: long


**`dellecs.health.storage_pool.chunks_repo_total_seal_size`**
:   Total seal size of repository chunks

type: long


**`dellecs.health.storage_pool.chunks_repo_number`**
:   Number of repository chunks

type: long


**`dellecs.health.storage_pool.chunks_xor_number`**
:   Number of XOR chunks

type: long


**`dellecs.health.storage_pool.chunks_xor_total_size`**
:   Total size of XOR chunks

type: long


**`dellecs.health.storage_pool.chunks_geo_cache_total_size`**
:   Total size of geo cache chunks

type: long


**`dellecs.health.storage_pool.chunks_geo_cache_count`**
:   Count of geo cache chunks

type: long


**`dellecs.health.storage_pool.chunks_geo_copy_number`**
:   Number of geo copy chunks

type: long


**`dellecs.health.storage_pool.chunks_geo_copy_total_size`**
:   Total size of geo copy chunks

type: long


**`dellecs.health.storage_pool.num_nodes`**
:   Number of nodes

type: long


**`dellecs.health.storage_pool.num_disks`**
:   Number of disks

type: long


**`dellecs.health.storage_pool.num_bad_nodes`**
:   Number of bad nodes

type: long


**`dellecs.health.storage_pool.num_good_nodes`**
:   Number of good nodes

type: long


**`dellecs.health.storage_pool.num_bad_disks`**
:   Number of bad disks

type: long


**`dellecs.health.storage_pool.num_good_disks`**
:   Number of good disks

type: long


**`dellecs.health.storage_pool.num_maintenance_nodes`**
:   Number of nodes under maintenance

type: long


**`dellecs.health.storage_pool.num_maintenance_disks`**
:   Number of disks under maintenance

type: long


**`dellecs.health.storage_pool.num_ready_to_replace_disks`**
:   Number of disks ready to be replaced

type: long


**`dellecs.health.storage_pool.num_nodes_with_sufficient_disk_space`**
:   Number of nodes with sufficient disk space

type: long


**`dellecs.health.storage_pool.gc_user_data_is_enabled`**
:   Indicates if user data garbage collection is enabled

type: boolean


**`dellecs.health.storage_pool.gc_system_metadata_is_enabled`**
:   Indicates if system metadata garbage collection is enabled

type: boolean


**`dellecs.health.storage_pool.recovery_complete_time_estimate`**
:   Estimated time to complete recovery

type: long


**`dellecs.health.storage_pool.chunks_ec_complete_time_estimate`**
:   Estimated time to complete EC chunks

type: long


**`dellecs.health.storage_pool.name`**
:   Name of the pool

type: keyword


**`dellecs.health.storage_pool.gc_user_unreclaimable_current`**
:   Current unreclaimable user data

type: long


**`dellecs.health.storage_pool.gc_user_total_detected_current`**
:   Current total detected user data

type: long


**`dellecs.health.storage_pool.gc_system_reclaimed_current`**
:   Current reclaimed system data

type: long


**`dellecs.health.storage_pool.gc_system_reclaimed_per_interval`**
:   System data reclaimed per interval

type: long


**`dellecs.health.storage_pool.gc_system_reclaimed_over_time_range`**
:   System data reclaimed over time range

type: long


**`dellecs.health.storage_pool.gc_combined_reclaimed_current`**
:   Current combined reclaimed data

type: long


**`dellecs.health.storage_pool.gc_user_reclaimed_over_time_range`**
:   User data reclaimed over time range

type: long


**`dellecs.health.storage_pool.gc_system_pending_current`**
:   Current pending system data

type: long


**`dellecs.health.storage_pool.gc_user_reclaimed_per_interval`**
:   User data reclaimed per interval

type: long


**`dellecs.health.storage_pool.gc_combined_total_detected_current`**
:   Current combined total detected data

type: long


**`dellecs.health.storage_pool.gc_system_total_detected_current`**
:   Current total detected system data

type: long


**`dellecs.health.storage_pool.gc_user_reclaimed_current`**
:   Current reclaimed user data

type: long


**`dellecs.health.storage_pool.gc_user_pending_current`**
:   Current pending user data

type: long


**`dellecs.health.storage_pool.gc_system_unreclaimable_current`**
:   Current unreclaimable system data

type: long


**`dellecs.health.storage_pool.gc_combined_reclaimed_over_time_range`**
:   Combined data reclaimed over time range

type: long


**`dellecs.health.storage_pool.gc_combined_pending_current`**
:   Current pending combined data

type: long


**`dellecs.health.storage_pool.gc_combined_unreclaimable_current`**
:   Current unreclaimable combined data

type: long


**`dellecs.health.storage_pool.disk_space_allocated_geo_cache_current`**
:   Current allocated disk space for geo cache

type: long


**`dellecs.health.storage_pool.disk_space_allocated_local_protection_current`**
:   Current allocated disk space for local protection

type: long


**`dellecs.health.storage_pool.disk_space_allocated_system_metadata_current`**
:   Current allocated disk space for system metadata

type: long


**`dellecs.health.storage_pool.disk_space_allocated_user_data_current`**
:   Current allocated disk space for user data

type: long


**`dellecs.health.storage_pool.allocated_capacity_forecast`**
:   Forecasted allocated capacity

type: long


**`dellecs.health.storage_pool.disk_space_allocated_percentage`**
:   Percentage of allocated disk space

type: float


**`dellecs.health.storage_pool.chunks_ec_coded_ratio_current`**
:   Current EC coded ratio

type: float


**`dellecs.health.storage_pool.chunks_ec_coded_ratio`**
:   EC coded ratio

type: float


**`dellecs.health.storage_pool.disk_space_allocated_percentage_current`**
:   Current percentage of allocated disk space

type: float


## rep_group [_rep_group]

Information about replication groups

**`dellecs.health.rep_group.name`**
:   Name of the replication group

type: keyword


**`dellecs.health.rep_group.group_id`**
:   Unique identifier for the replication group

type: keyword


**`dellecs.health.rep_group.num_zones`**
:   Number of zones in the replication group

type: long


**`dellecs.health.rep_group.chunks_pending_xor_total_size`**
:   Total size of pending XOR chunks

type: long


**`dellecs.health.rep_group.chunks_repo_pending_replication_total_size`**
:   Total size of repository chunks pending replication

type: long


**`dellecs.health.rep_group.chunks_journal_pending_replication_total_size`**
:   Total size of journal chunks pending replication

type: long


**`dellecs.health.rep_group.replication_egress_traffic`**
:   Replication egress traffic

type: float


**`dellecs.health.rep_group.replication_ingress_traffic`**
:   Replication ingress traffic

type: float


**`dellecs.health.rep_group.replication_ingress_traffic_current`**
:   Current replication ingress traffic

type: float


**`dellecs.health.rep_group.replication_egress_traffic_current`**
:   Current replication egress traffic

type: float


## node_processes [_node_processes]

Information about node processes

**`dellecs.health.node_processes.process_id`**
:   Unique identifier for the process

type: keyword


**`dellecs.health.node_processes.process_name`**
:   Name of the process

type: keyword


**`dellecs.health.node_processes.cpu_utilization`**
:   CPU utilization percentage

type: float


**`dellecs.health.node_processes.java_heap_utilization`**
:   Java heap utilization percentage

type: float


**`dellecs.health.node_processes.max_java_heap_size`**
:   Maximum Java heap size in bytes

type: long


**`dellecs.health.node_processes.memory_utilization_bytes`**
:   Memory utilization in bytes

type: long


**`dellecs.health.node_processes.memory_utilization`**
:   Memory utilization percentage

type: float


**`dellecs.health.node_processes.thread_count`**
:   Number of threads

type: long


**`dellecs.health.node_processes.restart_time`**
:   Time of the last restart

type: date


## node_details [_node_details]

Details about nodes

**`dellecs.health.node_details.node_id`**
:   Unique identifier for the node

type: keyword


**`dellecs.health.node_details.storage_pool_id`**
:   Unique identifier for the storage pool

type: keyword


**`dellecs.health.node_details.name`**
:   Display name of the node

type: keyword


**`dellecs.health.node_details.api_change`**
:   Indicates if there was an API change

type: integer


**`dellecs.health.node_details.num_bad_disks`**
:   Number of bad disks

type: long


**`dellecs.health.node_details.storage_pool_name`**
:   Name of the storage pool

type: keyword


**`dellecs.health.node_details.display_name`**
:   Display name of the node

type: keyword


**`dellecs.health.node_details.num_ready_to_replace_disks`**
:   Number of disks ready to be replaced

type: long


**`dellecs.health.node_details.num_maintenance_disks`**
:   Number of disks under maintenance

type: long


**`dellecs.health.node_details.health_status`**
:   Health status of the node

type: keyword


**`dellecs.health.node_details.num_good_disks`**
:   Number of good disks

type: long


**`dellecs.health.node_details.num_disks`**
:   Total number of disks

type: long


**`dellecs.health.node_details.disk_space_free_current_l1`**
:   Current free disk space at level 1

type: long


**`dellecs.health.node_details.allocated_capacity_forecast`**
:   Forecasted allocated capacity

type: long


**`dellecs.health.node_details.disk_space_free_current_l2`**
:   Current free disk space at level 2

type: long


**`dellecs.health.node_details.disk_space_reserved_current`**
:   Current reserved disk space

type: long


**`dellecs.health.node_details.disk_space_free_l2`**
:   Free disk space at level 2

type: long


**`dellecs.health.node_details.disk_space_free_l1`**
:   Free disk space at level 1

type: long


**`dellecs.health.node_details.disk_space_total`**
:   Total disk space

type: long


**`dellecs.health.node_details.disk_space_total_current_l1`**
:   Current total disk space at level 1

type: long


**`dellecs.health.node_details.disk_space_total_current_l2`**
:   Current total disk space at level 2

type: long


**`dellecs.health.node_details.disk_space_offline_total_current`**
:   Current total offline disk space

type: long


**`dellecs.health.node_details.disk_space_allocated_l2`**
:   Allocated disk space at level 2

type: long


**`dellecs.health.node_details.disk_space_allocated_l1`**
:   Allocated disk space at level 1

type: long


**`dellecs.health.node_details.disk_space_free_current`**
:   Current free disk space

type: long


**`dellecs.health.node_details.disk_space_total_current`**
:   Current total disk space

type: long


**`dellecs.health.node_details.disk_space_allocated_current`**
:   Current allocated disk space

type: long


**`dellecs.health.node_details.disk_space_allocated_current_l1`**
:   Current allocated disk space at level 1

type: long


**`dellecs.health.node_details.disk_space_allocated_current_l2`**
:   Current allocated disk space at level 2

type: long


**`dellecs.health.node_details.disk_space_allocated`**
:   Allocated disk space

type: long


**`dellecs.health.node_details.disk_space_allocated_percent`**
:   Percentage of allocated disk space

type: float


**`dellecs.health.node_details.disk_space_allocated_percentage_current`**
:   Current percentage of allocated disk space

type: float


