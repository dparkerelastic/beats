package health

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func isEmpty(value interface{}) bool {
	// we make use of the fact that all the dashboard API responses utilize
	// pointers for non-string types to filter out empty values from metric events.

	if value == nil {
		return true
	}

	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		return reflect.ValueOf(value).IsNil()
	}

	if t.Kind() == reflect.Slice || t.Kind() == reflect.String {
		return reflect.ValueOf(value).Len() == 0
	}

	return false
}

func reportMetricsForIntersight(reporter mb.ReporterV2, baseURL string, metrics ...[]mapstr.M) {
	for _, metricSlice := range metrics {
		for _, metric := range metricSlice {
			event := mb.Event{ModuleFields: mapstr.M{"base_url": baseURL}}
			if ts, ok := metric["@timestamp"]; ok {
				t, err := time.Parse(time.RFC3339, ts.(string))
				if err == nil {
					// if the timestamp parsing fails, we just fall back to the event time
					// (and leave the additional timestamp in the event for posterity)
					event.Timestamp = t
					delete(metric, "@timestamp")
				}
			}

			for k, v := range metric {
				if !isEmpty(v) {
					//fmt.Println("k =" + k + " v=" + string(v))
					event.ModuleFields.Put(k, v)
				}
			}

			reporter.Event(event)
		}
	}
}

func reportMetrics(reporter mb.ReporterV2, baseURL string, data IntersightData, debug bool) {

	metrics := []mapstr.M{}

	for _, summary := range data.computePhysicalSummaryList.Results {
		metric := mapstr.M{}
		metric["health.compute_physical_summary.class_id"] = summary.ClassId
		metric["health.compute_physical_summary.object_type"] = summary.ObjectType
		if summary.AdminPowerState != nil {
			metric["health.compute_physical_summary.admin_power_state"] = *summary.AdminPowerState
		}
		if summary.AssetTag != nil {
			metric["health.compute_physical_summary.asset_tag"] = *summary.AssetTag
		}
		if summary.AvailableMemory != nil {
			metric["health.compute_physical_summary.available_memory"] = *summary.AvailableMemory
		}
		if summary.BiosPostComplete != nil {
			metric["health.compute_physical_summary.bios_post_complete"] = *summary.BiosPostComplete
		}
		if summary.ChassisId != nil {
			metric["health.compute_physical_summary.chassis_id"] = *summary.ChassisId
		}
		if summary.ConnectionStatus != nil {
			metric["health.compute_physical_summary.connection_status"] = *summary.ConnectionStatus
		}
		if summary.CoolingMode != nil {
			metric["health.compute_physical_summary.cooling_mode"] = *summary.CoolingMode
		}
		if summary.CpuCapacity != nil {
			metric["health.compute_physical_summary.cpu_capacity"] = *summary.CpuCapacity
		}
		if summary.DeviceMoId != nil {
			metric["health.compute_physical_summary.device_mo_id"] = *summary.DeviceMoId
		}
		if summary.Dn != nil {
			metric["health.compute_physical_summary.dn"] = *summary.Dn
		}
		if summary.FaultSummary != nil {
			metric["health.compute_physical_summary.fault_summary"] = *summary.FaultSummary
		}
		if summary.Firmware != nil {
			metric["health.compute_physical_summary.firmware"] = *summary.Firmware
		}
		if summary.FrontPanelLockState != nil {
			metric["health.compute_physical_summary.front_panel_lock_state"] = *summary.FrontPanelLockState
		}
		if summary.HardwareUuid != nil {
			metric["health.compute_physical_summary.hardware_uuid"] = *summary.HardwareUuid
		}
		if summary.Ipv4Address != nil {
			metric["health.compute_physical_summary.ipv4_address"] = *summary.Ipv4Address
		}
		if summary.IsUpgraded != nil {
			metric["health.compute_physical_summary.is_upgraded"] = *summary.IsUpgraded
		}
		if summary.KvmServerStateEnabled != nil {
			metric["health.compute_physical_summary.kvm_server_state_enabled"] = *summary.KvmServerStateEnabled
		}
		if summary.KvmVendor != nil {
			metric["health.compute_physical_summary.kvm_vendor"] = *summary.KvmVendor
		}
		if summary.Lifecycle != nil {
			metric["health.compute_physical_summary.lifecycle"] = *summary.Lifecycle
		}
		if summary.ManagementMode != nil {
			metric["health.compute_physical_summary.management_mode"] = *summary.ManagementMode
		}
		if summary.MemorySpeed != nil {
			metric["health.compute_physical_summary.memory_speed"] = *summary.MemorySpeed
		}
		if summary.MgmtIpAddress != nil {
			metric["health.compute_physical_summary.mgmt_ip_address"] = *summary.MgmtIpAddress
		}
		if summary.Model != nil {
			metric["health.compute_physical_summary.model"] = *summary.Model
		}
		if summary.Name != nil {
			metric["health.compute_physical_summary.name"] = *summary.Name
		}
		if summary.NumAdaptors != nil {
			metric["health.compute_physical_summary.num_adaptors"] = *summary.NumAdaptors
		}
		if summary.NumCpuCores != nil {
			metric["health.compute_physical_summary.num_cpu_cores"] = *summary.NumCpuCores
		}
		if summary.NumCpuCoresEnabled != nil {
			metric["health.compute_physical_summary.num_cpu_cores_enabled"] = *summary.NumCpuCoresEnabled
		}
		if summary.NumCpus != nil {
			metric["health.compute_physical_summary.num_cpus"] = *summary.NumCpus
		}
		if summary.NumEthHostInterfaces != nil {
			metric["health.compute_physical_summary.num_eth_host_interfaces"] = *summary.NumEthHostInterfaces
		}
		if summary.NumFcHostInterfaces != nil {
			metric["health.compute_physical_summary.num_fc_host_interfaces"] = *summary.NumFcHostInterfaces
		}
		if summary.NumThreads != nil {
			metric["health.compute_physical_summary.num_threads"] = *summary.NumThreads
		}
		if summary.OperPowerState != nil {
			metric["health.compute_physical_summary.oper_power_state"] = *summary.OperPowerState
		}
		if summary.OperState != nil {
			metric["health.compute_physical_summary.oper_state"] = *summary.OperState
		}
		if summary.Operability != nil {
			metric["health.compute_physical_summary.operability"] = *summary.Operability
		}
		if summary.PackageVersion != nil {
			metric["health.compute_physical_summary.package_version"] = *summary.PackageVersion
		}
		if summary.Personality != nil {
			metric["health.compute_physical_summary.personality"] = *summary.Personality
		}
		if summary.PlatformType != nil {
			metric["health.compute_physical_summary.platform_type"] = *summary.PlatformType
		}
		if summary.Presence != nil {
			metric["health.compute_physical_summary.presence"] = *summary.Presence
		}
		if summary.Revision != nil {
			metric["health.compute_physical_summary.revision"] = *summary.Revision
		}
		if summary.Rn != nil {
			metric["health.compute_physical_summary.rn"] = *summary.Rn
		}
		if summary.ScaledMode != nil {
			metric["health.compute_physical_summary.scaled_mode"] = *summary.ScaledMode
		}
		if summary.Serial != nil {
			metric["health.compute_physical_summary.serial"] = *summary.Serial
		}
		if summary.ServerId != nil {
			metric["health.compute_physical_summary.server_id"] = *summary.ServerId
		}
		if summary.ServiceProfile != nil {
			metric["health.compute_physical_summary.service_profile"] = *summary.ServiceProfile
		}
		if summary.SlotId != nil {
			metric["health.compute_physical_summary.slot_id"] = *summary.SlotId
		}
		if summary.SourceObjectType != nil {
			metric["health.compute_physical_summary.source_object_type"] = *summary.SourceObjectType
		}
		if summary.TopologyScanStatus != nil {
			metric["health.compute_physical_summary.topology_scan_status"] = *summary.TopologyScanStatus
		}
		if summary.TotalMemory != nil {
			metric["health.compute_physical_summary.total_memory"] = *summary.TotalMemory
		}
		if summary.TunneledKvm != nil {
			metric["health.compute_physical_summary.tunneled_kvm"] = *summary.TunneledKvm
		}
		if summary.UserLabel != nil {
			metric["health.compute_physical_summary.user_label"] = *summary.UserLabel
		}
		if summary.Uuid != nil {
			metric["health.compute_physical_summary.uuid"] = *summary.Uuid
		}
		if summary.Vendor != nil {
			metric["health.compute_physical_summary.vendor"] = *summary.Vendor
		}
		if len(summary.OperReason) > 0 {
			metric["health.compute_physical_summary.oper_reason"] = summary.OperReason
		}
		if debug {
			if summaryJson, err := json.Marshal(summary); err == nil {
				metric["health.message"] = string(summaryJson)
			} else {
				metric["health.message"] = "error marshaling computePhysicalSummary"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, chassis := range data.chassisList.Results {
		metric := mapstr.M{}
		metric["health.chassis.class_id"] = chassis.ClassId
		metric["health.chassis.object_type"] = chassis.ObjectType
		if chassis.ChassisId != nil {
			metric["health.chassis.chassis_id"] = *chassis.ChassisId
		}
		if chassis.ConnectionPath != nil {
			metric["health.chassis.connection_path"] = *chassis.ConnectionPath
		}
		if chassis.ConnectionStatus != nil {
			metric["health.chassis.connection_status"] = *chassis.ConnectionStatus
		}
		if chassis.Description != nil {
			metric["health.chassis.description"] = *chassis.Description
		}
		if chassis.FaultSummary != nil {
			metric["health.chassis.fault_summary"] = *chassis.FaultSummary
		}
		if chassis.ManagementMode != nil {
			metric["health.chassis.management_mode"] = *chassis.ManagementMode
		}
		if chassis.Name != nil {
			metric["health.chassis.name"] = *chassis.Name
		}
		if len(chassis.OperReason) > 0 {
			metric["health.chassis.oper_reason"] = chassis.OperReason
		}
		if chassis.OperState != nil {
			metric["health.chassis.oper_state"] = *chassis.OperState
		}
		if chassis.PartNumber != nil {
			metric["health.chassis.part_number"] = *chassis.PartNumber
		}
		if chassis.Pid != nil {
			metric["health.chassis.pid"] = *chassis.Pid
		}
		if chassis.PlatformType != nil {
			metric["health.chassis.platform_type"] = *chassis.PlatformType
		}
		if chassis.ProductName != nil {
			metric["health.chassis.product_name"] = *chassis.ProductName
		}
		if chassis.Sku != nil {
			metric["health.chassis.sku"] = *chassis.Sku
		}
		if chassis.UserLabel != nil {
			metric["health.chassis.user_label"] = *chassis.UserLabel
		}
		if chassis.Vid != nil {
			metric["health.chassis.vid"] = *chassis.Vid
		}
		if chassis.Presence != nil {
			metric["health.chassis.presence"] = *chassis.Presence
		}
		if chassis.Serial != nil {
			metric["health.chassis.serial"] = *chassis.Serial
		}
		if debug {
			if chassisJson, err := json.Marshal(chassis); err == nil {
				metric["health.message"] = string(chassisJson)
			} else {
				metric["health.message"] = "error marshaling chassis"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, networkElement := range data.networkElementList.Results {
		metric := mapstr.M{}
		metric["health.network_element.class_id"] = networkElement.ClassId
		metric["health.network_element.object_type"] = networkElement.ObjectType

		if networkElement.AdminEvacState != nil {
			metric["health.network_element.admin_evac_state"] = *networkElement.AdminEvacState
		}
		if networkElement.AdminInbandInterfaceState != nil {
			metric["health.network_element.admin_inband_interface_state"] = *networkElement.AdminInbandInterfaceState
		}
		if networkElement.AvailableMemory != nil {
			metric["health.network_element.available_memory"] = *networkElement.AvailableMemory
		}
		if networkElement.Chassis != nil {
			metric["health.network_element.chassis"] = *networkElement.Chassis
		}
		if networkElement.ConfModTs != nil {
			metric["health.network_element.conf_mod_ts"] = *networkElement.ConfModTs
		}
		if networkElement.ConfModTsBackup != nil {
			metric["health.network_element.conf_mod_ts_backup"] = *networkElement.ConfModTsBackup
		}
		if networkElement.ConnectionStatus != nil {
			metric["health.network_element.connection_status"] = *networkElement.ConnectionStatus
		}
		if networkElement.DefaultDomain != nil {
			metric["health.network_element.default_domain"] = *networkElement.DefaultDomain
		}
		if networkElement.EthernetMode != nil {
			metric["health.network_element.ethernet_mode"] = *networkElement.EthernetMode
		}
		if networkElement.EthernetSwitchingMode != nil {
			metric["health.network_element.ethernet_switching_mode"] = *networkElement.EthernetSwitchingMode
		}
		if networkElement.FaultSummary != nil {
			metric["health.network_element.fault_summary"] = *networkElement.FaultSummary
		}
		if networkElement.FcMode != nil {
			metric["health.network_element.fc_mode"] = *networkElement.FcMode
		}
		if networkElement.FcSwitchingMode != nil {
			metric["health.network_element.fc_switching_mode"] = *networkElement.FcSwitchingMode
		}
		if networkElement.FpgaUpgradeNeeded != nil {
			metric["health.network_element.fpga_upgrade_needed"] = *networkElement.FpgaUpgradeNeeded
		}
		if networkElement.InbandIpAddress != nil {
			metric["health.network_element.inband_ip_address"] = *networkElement.InbandIpAddress
		}
		if networkElement.InbandIpGateway != nil {
			metric["health.network_element.inband_ip_gateway"] = *networkElement.InbandIpGateway
		}
		if networkElement.InbandIpMask != nil {
			metric["health.network_element.inband_ip_mask"] = *networkElement.InbandIpMask
		}
		if networkElement.InbandVlan != nil {
			metric["health.network_element.inband_vlan"] = *networkElement.InbandVlan
		}
		if networkElement.InterClusterLinkState != nil {
			metric["health.network_element.inter_cluster_link_state"] = *networkElement.InterClusterLinkState
		}
		if networkElement.ManagementMode != nil {
			metric["health.network_element.management_mode"] = *networkElement.ManagementMode
		}
		if networkElement.Model != nil {
			metric["health.network_element.model"] = *networkElement.Model
		}
		// if networkElement.Name != nil {
		// 	metric["health.network_element.name"] = *networkElement.Name
		// }
		if networkElement.OperEvacState != nil {
			metric["health.network_element.oper_evac_state"] = *networkElement.OperEvacState
		}
		if networkElement.Operability != nil {
			metric["health.network_element.operability"] = *networkElement.Operability
		}
		if networkElement.OutOfBandIpAddress != nil {
			metric["health.network_element.out_of_band_ip_address"] = *networkElement.OutOfBandIpAddress
		}
		if networkElement.OutOfBandIpGateway != nil {
			metric["health.network_element.out_of_band_ip_gateway"] = *networkElement.OutOfBandIpGateway
		}
		if networkElement.OutOfBandIpMask != nil {
			metric["health.network_element.out_of_band_ip_mask"] = *networkElement.OutOfBandIpMask
		}
		if networkElement.OutOfBandIpv4Address != nil {
			metric["health.network_element.out_of_band_ipv4_address"] = *networkElement.OutOfBandIpv4Address
		}
		if networkElement.OutOfBandIpv4Gateway != nil {
			metric["health.network_element.out_of_band_ipv4_gateway"] = *networkElement.OutOfBandIpv4Gateway
		}
		if networkElement.OutOfBandIpv4Mask != nil {
			metric["health.network_element.out_of_band_ipv4_mask"] = *networkElement.OutOfBandIpv4Mask
		}
		if networkElement.OutOfBandIpv6Address != nil {
			metric["health.network_element.out_of_band_ipv6_address"] = *networkElement.OutOfBandIpv6Address
		}
		if networkElement.OutOfBandIpv6Gateway != nil {
			metric["health.network_element.out_of_band_ipv6_gateway"] = *networkElement.OutOfBandIpv6Gateway
		}
		if networkElement.OutOfBandIpv6Prefix != nil {
			metric["health.network_element.out_of_band_ipv6_prefix"] = *networkElement.OutOfBandIpv6Prefix
		}
		if networkElement.OutOfBandMac != nil {
			metric["health.network_element.out_of_band_mac"] = *networkElement.OutOfBandMac
		}
		if networkElement.PartNumber != nil {
			metric["health.network_element.part_number"] = *networkElement.PartNumber
		}
		if networkElement.PeerFirmwareOutOfSync != nil {
			metric["health.network_element.peer_firmware_out_of_sync"] = *networkElement.PeerFirmwareOutOfSync
		}
		if networkElement.Presence != nil {
			metric["health.network_element.presence"] = *networkElement.Presence
		}
		if networkElement.ReservedVlanStartId != nil {
			metric["health.network_element.reserved_vlan_start_id"] = *networkElement.ReservedVlanStartId
		}
		if networkElement.Serial != nil {
			metric["health.network_element.serial"] = *networkElement.Serial
		}
		if networkElement.Status != nil {
			metric["health.network_element.status"] = *networkElement.Status
		}
		if networkElement.SwitchId != nil {
			metric["health.network_element.switch_id"] = *networkElement.SwitchId
		}
		if networkElement.SwitchProfileName != nil {
			metric["health.network_element.switch_profile_name"] = *networkElement.SwitchProfileName
		}
		if networkElement.SwitchType != nil {
			metric["health.network_element.switch_type"] = *networkElement.SwitchType
		}
		if networkElement.SystemUpTime != nil {
			metric["health.network_element.system_up_time"] = *networkElement.SystemUpTime
		}
		if networkElement.Thermal != nil {
			metric["health.network_element.thermal"] = *networkElement.Thermal
		}
		if networkElement.TotalMemory != nil {
			metric["health.network_element.total_memory"] = *networkElement.TotalMemory
		}
		if networkElement.UserLabel != nil {
			metric["health.network_element.user_label"] = *networkElement.UserLabel
		}
		if networkElement.Version != nil {
			metric["health.network_element.version"] = *networkElement.Version
		}
		if debug {
			if networkElementJson, err := json.Marshal(networkElement); err == nil {
				metric["health.message"] = string(networkElementJson)
			} else {
				metric["health.message"] = "error marshaling networkElement"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, asset := range data.assetDeviceRegistrationList.Results {
		metric := mapstr.M{}
		metric["health.asset_device_registration.class_id"] = asset.ClassId
		metric["health.asset_device_registration.object_type"] = asset.ObjectType

		// if asset.AccessKeyId != nil {
		// 	metric["health.asset_device_registration.access_key_id"] = *asset.AccessKeyId
		// }
		if asset.ClaimedByUserName != nil {
			metric["health.asset_device_registration.claimed_by_user_name"] = *asset.ClaimedByUserName
		}
		if asset.ClaimedTime != nil {
			metric["health.asset_device_registration.claimed_time"] = asset.ClaimedTime.Format(time.RFC3339)
		}
		if len(asset.DeviceHostname) > 0 {
			metric["health.asset_device_registration.device_hostname"] = asset.DeviceHostname
		}
		if len(asset.DeviceIpAddress) > 0 {
			metric["health.asset_device_registration.device_ip_address"] = asset.DeviceIpAddress
		}
		if asset.ExecutionMode != nil {
			metric["health.asset_device_registration.execution_mode"] = *asset.ExecutionMode
		}
		if len(asset.Pid) > 0 {
			metric["health.asset_device_registration.pid"] = asset.Pid
		}
		if asset.PlatformType != nil {
			metric["health.asset_device_registration.platform_type"] = *asset.PlatformType
		}
		// if asset.PublicAccessKey != nil {
		// 	metric["health.asset_device_registration.public_access_key"] = *asset.PublicAccessKey
		// }
		// if asset.PublicAccessKeyRotated != nil {
		// 	metric["health.asset_device_registration.public_access_key_rotated"] = *asset.PublicAccessKeyRotated
		// }
		// if asset.PublicEncryptionKey != nil {
		// 	metric["health.asset_device_registration.public_encryption_key"] = *asset.PublicEncryptionKey
		// }
		if asset.ReadOnly != nil {
			metric["health.asset_device_registration.read_only"] = *asset.ReadOnly
		}
		// if asset.RotateAccessKey != nil {
		// 	metric["health.asset_device_registration.rotate_access_key"] = *asset.RotateAccessKey
		// }
		if len(asset.Serial) > 0 {
			metric["health.asset_device_registration.serial"] = asset.Serial
		}
		if asset.Vendor != nil {
			metric["health.asset_device_registration.vendor"] = *asset.Vendor
		}
		// Relationships and additional properties can be marshaled for debug or extended as needed

		if debug {
			if assetJson, err := json.Marshal(asset); err == nil {
				metric["health.message"] = string(assetJson)
			} else {
				metric["health.message"] = "error marshaling assetDeviceRegistration"
			}
		}
		metrics = append(metrics, metric)
	}

	// Process fanEvent
	for _, telemetry := range data.fanEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if fanEventJson, err := json.Marshal(data.fanEvent); err == nil {
				metric["health.message"] = string(fanEventJson)
			} else {
				metric["health.message"] = "error marshaling fanEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.memoryEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if memoryEventJson, err := json.Marshal(data.memoryEvent); err == nil {
				metric["health.message"] = string(memoryEventJson)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.physicalProcessorEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if physicalProcessorEvent, err := json.Marshal(data.physicalProcessorEvent); err == nil {
				metric["health.message"] = string(physicalProcessorEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.powerSupplyEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if powerSupplyEvent, err := json.Marshal(data.powerSupplyEvent); err == nil {
				metric["health.message"] = string(powerSupplyEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.temperatureEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if temperatureEvent, err := json.Marshal(data.temperatureEvent); err == nil {
				metric["health.message"] = string(temperatureEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.systemCPUEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if systemCPUEvent, err := json.Marshal(data.systemCPUEvent); err == nil {
				metric["health.message"] = string(systemCPUEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.systemMemoryEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if systemMemoryEvent, err := json.Marshal(data.systemMemoryEvent); err == nil {
				metric["health.message"] = string(systemMemoryEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.hostPowerAndStatusEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if hostPowerAndStatusEvent, err := json.Marshal(data.hostPowerAndStatusEvent); err == nil {
				metric["health.message"] = string(hostPowerAndStatusEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.signalPowerEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if signalPowerEvent, err := json.Marshal(data.signalPowerEvent); err == nil {
				metric["health.message"] = string(signalPowerEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.graphicalProcessingUnitEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if graphicalProcessingUnitEvent, err := json.Marshal(data.graphicalProcessingUnitEvent); err == nil {
				metric["health.message"] = string(graphicalProcessingUnitEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.electricCurrentEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if electricCurrentEvent, err := json.Marshal(data.electricCurrentEvent); err == nil {
				metric["health.message"] = string(electricCurrentEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	for _, telemetry := range data.voltageEvent {
		metric := mapstr.M{}
		for k, v := range telemetry.Event {
			metric["health."+k] = v
		}
		if telemetry.Timestamp != nil {
			metric["health.time_stamp"] = telemetry.Timestamp.Format(time.RFC3339)
		}
		if telemetry.Version != nil && *telemetry.Version != "" {
			metric["health.version"] = *telemetry.Version
		}

		if debug {
			if voltageEvent, err := json.Marshal(data.voltageEvent); err == nil {
				metric["health.message"] = string(voltageEvent)
			} else {
				metric["health.message"] = "error marshaling memoryEvent"
			}
		}
		metrics = append(metrics, metric)
	}

	reportMetricsForIntersight(reporter, baseURL, metrics)
}
