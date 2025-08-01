---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/exported-fields-purestorage.html
---

% This file is generated! See scripts/generate_fields_docs.py

# purestorage fields [exported-fields-purestorage]

purestorage module

## purestorage [_purestorage]



## arrays [_arrays]

arrays

**`purestorage.arrays.performance.read_bandwidth`**
:   Read bandwidth in bytes per second.

type: long


**`purestorage.arrays.performance.read_iops`**
:   Read IOPS (Input/Output Operations Per Second).

type: long


**`purestorage.arrays.performance.read_latency`**
:   Read latency in milliseconds.

type: long


**`purestorage.arrays.performance.write_bandwidth`**
:   Write bandwidth in bytes per second.

type: long


**`purestorage.arrays.performance.write_iops`**
:   Write IOPS (Input/Output Operations Per Second).

type: long


**`purestorage.arrays.performance.write_latency`**
:   Write latency in milliseconds.

type: long


## health [_health]

health

**`purestorage.health.array_controller.status`**
:   Status of the health check

type: keyword


**`purestorage.health.array_controller.name`**
:   Name of the entity

type: keyword


**`purestorage.health.array_controller.version`**
:   Version information

type: keyword


**`purestorage.health.array_controller.mode`**
:   Mode of operation

type: keyword


**`purestorage.health.array_controller.model`**
:   Model identifier

type: keyword


**`purestorage.health.array_monitor.writes_per_sec`**
:   Number of writes per second

type: long


**`purestorage.health.array_monitor.usec_per_write_op`**
:   Microseconds per write operation

type: long


**`purestorage.health.array_monitor.output_per_sec`**
:   Output per second

type: long


**`purestorage.health.array_monitor.reads_per_sec`**
:   Number of reads per second

type: long


**`purestorage.health.array_monitor.input_per_sec`**
:   Input per second

type: long


**`purestorage.health.array_monitor.time`**
:   Timestamp of the health check

type: date


**`purestorage.health.array_monitor.usec_per_read_op`**
:   Microseconds per read operation

type: long


**`purestorage.health.array_monitor.queue_depth`**
:   Queue depth

type: long


**`purestorage.health.array_space.capacity`**
:   Capacity information

type: long


**`purestorage.health.array_space.hostname`**
:   Hostname of the system

type: keyword


**`purestorage.health.array_space.system`**
:   System identifier

type: keyword


**`purestorage.health.array_space.snapshots`**
:   Number of snapshots

type: long


**`purestorage.health.array_space.volumes`**
:   Number of volumes

type: long


**`purestorage.health.array_space.data_reduction`**
:   Data reduction ratio

type: float


**`purestorage.health.array_space.total`**
:   Total capacity

type: long


**`purestorage.health.array_space.shared_space`**
:   Shared space capacity

type: long


**`purestorage.health.array_space.thin_provisioning`**
:   Thin provisioning ratio

type: float


**`purestorage.health.array_space.total_reduction`**
:   Total reduction ratio

type: float


**`purestorage.health.hardware.status`**
:   Status of the health check

type: keyword


**`purestorage.health.hardware.slot`**
:   Slot identifier

type: keyword


**`purestorage.health.hardware.name`**
:   Name of the entity

type: keyword


**`purestorage.health.hardware.temperature`**
:   Temperature reading

type: float


**`purestorage.health.hardware.index`**
:   Index number

type: long


**`purestorage.health.hardware.identify`**
:   Identify flag

type: keyword


**`purestorage.health.hardware.speed`**
:   Speed measurement

type: long


**`purestorage.health.hardware.details`**
:   Additional details

type: text


**`purestorage.health.drive.status`**
:   Status of the health check

type: keyword


**`purestorage.health.drive.capacity`**
:   Capacity information

type: long


**`purestorage.health.drive.type`**
:   Type of entity

type: keyword


**`purestorage.health.drive.name`**
:   Name of the entity

type: keyword


**`purestorage.health.pgroup.name`**
:   Name of prgroup

type: keyword


**`purestorage.health.pgroup.physical_bytes_written`**
:   Physical bytes written

type: long


**`purestorage.health.pgroup.started`**
:   Start time

type: date


**`purestorage.health.pgroup.completed`**
:   Completion time

type: date


**`purestorage.health.pgroup.created`**
:   Creation time

type: date


**`purestorage.health.pgroup.source`**
:   Source identifier

type: keyword


**`purestorage.health.pgroup.time_remaining`**
:   Time remaining

type: long


**`purestorage.health.pgroup.progress`**
:   Progress percentage

type: float


**`purestorage.health.pgroup.data_transferred`**
:   Data transferred

type: long


**`purestorage.health.volume_monitor.writes_per_sec`**
:   Number of writes per second

type: long


**`purestorage.health.volume_monitor.usec_per_write_op`**
:   Microseconds per write operation

type: long


**`purestorage.health.volume_monitor.output_per_sec`**
:   Output per second

type: long


**`purestorage.health.volume_monitor.reads_per_sec`**
:   Number of reads per second

type: long


**`purestorage.health.volume_monitor.input_per_sec`**
:   Input per second

type: long


**`purestorage.health.volume_monitor.time`**
:   Timestamp of the health check

type: date


**`purestorage.health.volume_monitor.usec_per_read_op`**
:   Microseconds per read operation

type: long


**`purestorage.health.volume_monitor.name`**
:   Volume name

type: keyword


**`purestorage.health.volume_space.size`**
:   Capacity information

type: long


**`purestorage.health.volume_space.name`**
:   Hostname of the system

type: keyword


**`purestorage.health.volume_space.system`**
:   System identifier

type: keyword


**`purestorage.health.volume_space.snapshots`**
:   Number of snapshots

type: long


**`purestorage.health.volume_space.volumes`**
:   Number of volumes

type: long


**`purestorage.health.volume_space.data_reduction`**
:   Data reduction ratio

type: float


**`purestorage.health.volume_space.total`**
:   Total capacity

type: long


**`purestorage.health.volume_space.shared_space`**
:   Shared space capacity

type: long


**`purestorage.health.volume_space.thin_provisioning`**
:   Thin provisioning ratio

type: float


**`purestorage.health.volume_space.total_reduction`**
:   Total reduction ratio

type: float


**`purestorage.health.array.revision`**
:   Revision identifier

type: keyword


**`purestorage.health.array.version`**
:   Version information

type: keyword


**`purestorage.health.array.array_name`**
:   Array name

type: keyword


**`purestorage.health.array.id`**
:   Unique identifier

type: keyword


**`purestorage.health.volume.name`**
:   Volume name

type: keyword


**`purestorage.health.volume.created`**
:   Creation time

type: date


**`purestorage.health.volume.source`**
:   Source identifier

type: keyword


**`purestorage.health.volume.serial`**
:   Serial number

type: keyword


**`purestorage.health.volume.size`**
:   Volume size

type: long


