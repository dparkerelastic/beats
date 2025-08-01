---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/exported-fields-citrix_cms.html
---

% This file is generated! See scripts/generate_fields_docs.py

# citrix_cms fields [exported-fields-citrix_cms]

citrix_cms module

## citrix_cms [_citrix_cms]



## health [_health]

health

**`citrix_cms.health.api.odatacontext`**
:   The OData context of the API response.

type: keyword


**`citrix_cms.health.load.index.ID`**
:   The unique identifier of the load summary.

type: long


**`citrix_cms.health.load.index.EffectiveLoadIndex`**
:   The effective load index value.

type: float


**`citrix_cms.health.load.index.CPU`**
:   The CPU load percentage.

type: float


**`citrix_cms.health.load.index.Memory`**
:   The memory load percentage.

type: float


**`citrix_cms.health.load.index.Disk`**
:   The disk load percentage.

type: float


**`citrix_cms.health.load.index.Network`**
:   The network load percentage.

type: float


**`citrix_cms.health.load.index.SessionCount`**
:   The session count on the machine.

type: long


**`citrix_cms.health.load.index.MachineID`**
:   The ID of the machine associated with the load summary.

type: keyword


**`citrix_cms.health.load.index.CreatedDate`**
:   The date when the load summary was created.

type: date


**`citrix_cms.health.load.index.ModifiedDate`**
:   The date when the load summary was last modified.

type: date


**`citrix_cms.health.load.index.summary.ID`**
:   The unique identifier of the load summary.

type: long


**`citrix_cms.health.load.index.summary.SummaryDate`**
:   The date of the load summary.

type: date


**`citrix_cms.health.load.index.summary.MachineId`**
:   The ID of the machine associated with the load summary.

type: keyword


**`citrix_cms.health.load.index.summary.SumCount`**
:   The total count of load activities.

type: long


**`citrix_cms.health.load.index.summary.SumLoadIndex`**
:   The total load index value.

type: long


**`citrix_cms.health.load.index.summary.SumCpu`**
:   The total CPU load percentage.

type: long


**`citrix_cms.health.load.index.summary.SumNetwork`**
:   The total network load percentage.

type: long


**`citrix_cms.health.load.index.summary.SumDisk`**
:   The total disk load percentage.

type: long


**`citrix_cms.health.load.index.summary.SumMemory`**
:   The total memory load percentage.

type: long


**`citrix_cms.health.load.index.summary.SumSessionCount`**
:   The total session count on the machine.

type: long


**`citrix_cms.health.load.index.summary.Granularity`**
:   The granularity level of the load summary.

type: long


**`citrix_cms.health.load.index.summary.CreatedDate`**
:   The date when the load summary was created.

type: date


**`citrix_cms.health.load.index.summary.ModifiedDate`**
:   The date when the load summary was last modified.

type: date


**`citrix_cms.health.logon.summaries.ID`**
:   The unique identifier of the logon summary.

type: long


**`citrix_cms.health.logon.summaries.SummaryDate`**
:   The date of the logon summary.

type: date


**`citrix_cms.health.logon.summaries.DesktopGroupID`**
:   The ID of the desktop group associated with the logon summary.

type: keyword


**`citrix_cms.health.logon.summaries.TotalDuration`**
:   The total duration of logon activities.

type: float


**`citrix_cms.health.logon.summaries.TotalCount`**
:   The total count of logon activities.

type: long


**`citrix_cms.health.logon.summaries.BrokeringDuration`**
:   The duration of brokering activities.

type: float


**`citrix_cms.health.logon.summaries.BrokeringCount`**
:   The count of brokering activities.

type: long


**`citrix_cms.health.logon.summaries.VMStartDuration`**
:   The duration of virtual machine start activities.

type: float


**`citrix_cms.health.logon.summaries.VMPowerOnDuration`**
:   The duration of virtual machine power-on activities.

type: float


**`citrix_cms.health.logon.summaries.VMRegistrationDuration`**
:   The duration of virtual machine registration activities.

type: float


**`citrix_cms.health.logon.summaries.VMStartCount`**
:   The count of virtual machine start activities.

type: long


**`citrix_cms.health.logon.summaries.HdxDuration`**
:   The duration of HDX activities.

type: float


**`citrix_cms.health.logon.summaries.HdxCount`**
:   The count of HDX activities.

type: long


**`citrix_cms.health.logon.summaries.AuthenticationDuration`**
:   The duration of authentication activities.

type: float


**`citrix_cms.health.logon.summaries.AuthenticationCount`**
:   The count of authentication activities.

type: long


**`citrix_cms.health.logon.summaries.GpoDuration`**
:   The duration of Group Policy Object (GPO) activities.

type: float


**`citrix_cms.health.logon.summaries.GpoCount`**
:   The count of Group Policy Object (GPO) activities.

type: long


**`citrix_cms.health.logon.summaries.LogOnScriptsDuration`**
:   The duration of logon script activities.

type: float


**`citrix_cms.health.logon.summaries.LogOnScriptsCount`**
:   The count of logon script activities.

type: long


**`citrix_cms.health.logon.summaries.InteractiveDuration`**
:   The duration of interactive logon activities.

type: float


**`citrix_cms.health.logon.summaries.InteractiveCount`**
:   The count of interactive logon activities.

type: long


**`citrix_cms.health.logon.summaries.ProfileLoadDuration`**
:   The duration of profile load activities.

type: float


**`citrix_cms.health.logon.summaries.ProfileLoadCount`**
:   The count of profile load activities.

type: long


**`citrix_cms.health.logon.summaries.Granularity`**
:   The granularity level of the logon summary.

type: long


**`citrix_cms.health.logon.summaries.CreatedDate`**
:   The date when the logon summary was created.

type: date


**`citrix_cms.health.logon.summaries.ModifiedDate`**
:   The date when the logon summary was last modified.

type: date


**`citrix_cms.health.machine.details.ID`**
:   The unique identifier of the machine.

type: keyword


**`citrix_cms.health.machine.details.Sid`**
:   The security identifier of the machine.

type: keyword


**`citrix_cms.health.machine.details.Name`**
:   The name of the machine.

type: keyword


**`citrix_cms.health.machine.details.DNSName`**
:   The DNS name of the machine.

type: keyword


**`citrix_cms.health.machine.details.LifecycleState`**
:   The lifecycle state of the machine.

type: long


**`citrix_cms.health.machine.details.IPAddress`**
:   The IP address of the machine.

type: keyword


**`citrix_cms.health.machine.details.HostedMachineID`**
:   The ID of the hosted machine.

type: keyword


**`citrix_cms.health.machine.details.HostingServerName`**
:   The name of the hosting server.

type: keyword


**`citrix_cms.health.machine.details.HostedMachineName`**
:   The name of the hosted machine.

type: keyword


**`citrix_cms.health.machine.details.IsAssigned`**
:   Indicates if the machine is assigned.

type: boolean


**`citrix_cms.health.machine.details.IsInMaintenanceMode`**
:   Indicates if the machine is in maintenance mode.

type: boolean


**`citrix_cms.health.machine.details.IsPendingUpdate`**
:   Indicates if the machine has a pending update.

type: boolean


**`citrix_cms.health.machine.details.AgentVersion`**
:   The version of the agent installed on the machine.

type: keyword


**`citrix_cms.health.machine.details.AssociatedUserFullNames`**
:   The full names of associated users.

type: keyword


**`citrix_cms.health.machine.details.AssociatedUserNames`**
:   The usernames of associated users.

type: keyword


**`citrix_cms.health.machine.details.AssociatedUserUPNs`**
:   The UPNs of associated users.

type: keyword


**`citrix_cms.health.machine.details.CurrentRegistrationState`**
:   The current registration state of the machine.

type: long


**`citrix_cms.health.machine.details.RegistrationStateChangeDate`**
:   The date when the registration state last changed.

type: date


**`citrix_cms.health.machine.details.LastDeregisteredCode`**
:   The code of the last deregistration event.

type: long


**`citrix_cms.health.machine.details.LastDeregisteredDate`**
:   The date of the last deregistration event.

type: date


**`citrix_cms.health.machine.details.CurrentPowerState`**
:   The current power state of the machine.

type: long


**`citrix_cms.health.machine.details.CurrentSessionCount`**
:   The current session count on the machine.

type: long


**`citrix_cms.health.machine.details.ControllerDNSName`**
:   The DNS name of the controller.

type: keyword


**`citrix_cms.health.machine.details.PoweredOnDate`**
:   The date when the machine was powered on.

type: date


**`citrix_cms.health.machine.details.PowerStateChangeDate`**
:   The date when the power state last changed.

type: date


**`citrix_cms.health.machine.details.FunctionalLevel`**
:   The functional level of the machine.

type: long


**`citrix_cms.health.machine.details.FailureDate`**
:   The date of the last failure event.

type: date


**`citrix_cms.health.machine.details.WindowsConnectionSetting`**
:   The Windows connection setting of the machine.

type: long


**`citrix_cms.health.machine.details.IsPreparing`**
:   Indicates if the machine is preparing.

type: boolean


**`citrix_cms.health.machine.details.FaultState`**
:   The fault state of the machine.

type: long


**`citrix_cms.health.machine.details.OSType`**
:   The operating system type of the machine.

type: keyword


**`citrix_cms.health.machine.details.CurrentLoadIndexID`**
:   The ID of the current load index.

type: long


**`citrix_cms.health.machine.details.CatalogID`**
:   The ID of the catalog.

type: keyword


**`citrix_cms.health.machine.details.DesktopGroupID`**
:   The ID of the desktop group.

type: keyword


**`citrix_cms.health.machine.details.HypervisorID`**
:   The ID of the hypervisor.

type: keyword


**`citrix_cms.health.machine.details.LastPowerActionCompletedDate`**
:   The date when the last power action was completed.

type: date


**`citrix_cms.health.machine.details.LastUpgradeState`**
:   The last upgrade state of the machine.

type: long


**`citrix_cms.health.machine.details.LastUpgradeStateChangeDate`**
:   The date when the last upgrade state changed.

type: date


**`citrix_cms.health.machine.details.Hash`**
:   The hash of the machine.

type: keyword


**`citrix_cms.health.machine.details.MachineRole`**
:   The role of the machine.

type: long


**`citrix_cms.health.machine.details.CreatedDate`**
:   The date when the machine was created.

type: date


**`citrix_cms.health.machine.details.ModifiedDate`**
:   The date when the machine was last modified.

type: date


## CurrentLoadIndex [_CurrentLoadIndex]

The current load index details of the machine.

**`citrix_cms.health.machine.details.CurrentLoadIndex.ID`**
:   The unique identifier of the load index.

type: long


**`citrix_cms.health.machine.details.CurrentLoadIndex.EffectiveLoadIndex`**
:   The effective load index value.

type: long


**`citrix_cms.health.machine.details.CurrentLoadIndex.CPU`**
:   The CPU load percentage.

type: float


**`citrix_cms.health.machine.details.CurrentLoadIndex.Memory`**
:   The memory load percentage.

type: float


**`citrix_cms.health.machine.details.CurrentLoadIndex.Disk`**
:   The disk load percentage.

type: float


**`citrix_cms.health.machine.details.CurrentLoadIndex.Network`**
:   The network load percentage.

type: float


**`citrix_cms.health.machine.details.CurrentLoadIndex.SessionCount`**
:   The session count on the machine.

type: long


**`citrix_cms.health.machine.details.CurrentLoadIndex.MachineID`**
:   The ID of the machine associated with the load index.

type: keyword


**`citrix_cms.health.machine.details.CurrentLoadIndex.CreatedDate`**
:   The date when the load index was created.

type: date


**`citrix_cms.health.machine.details.CurrentLoadIndex.ModifiedDate`**
:   The date when the load index was last modified.

type: date


## Catalog [_Catalog]

The catalog details of the machine.

**`citrix_cms.health.machine.details.Catalog.ID`**
:   The unique identifier of the catalog.

type: keyword


**`citrix_cms.health.machine.details.Catalog.Name`**
:   The name of the catalog.

type: keyword


**`citrix_cms.health.machine.details.Catalog.LifecycleState`**
:   The lifecycle state of the catalog.

type: long


**`citrix_cms.health.machine.details.Catalog.ProvisioningType`**
:   The provisioning type of the catalog.

type: long


**`citrix_cms.health.machine.details.Catalog.PersistentUserChanges`**
:   Indicates if persistent user changes are allowed.

type: long


**`citrix_cms.health.machine.details.Catalog.IsMachinePhysical`**
:   Indicates if the machine is physical.

type: boolean


**`citrix_cms.health.machine.details.Catalog.AllocationType`**
:   The allocation type of the catalog.

type: long


**`citrix_cms.health.machine.details.Catalog.SessionSupport`**
:   The session support type of the catalog.

type: long


**`citrix_cms.health.machine.details.Catalog.ProvisioningSchemeID`**
:   The ID of the provisioning scheme.

type: keyword


**`citrix_cms.health.machine.details.Catalog.ZoneUID`**
:   The unique identifier of the zone.

type: keyword


**`citrix_cms.health.machine.details.Catalog.ZoneName`**
:   The name of the zone.

type: keyword


**`citrix_cms.health.machine.details.Catalog.CreatedDate`**
:   The date when the catalog was created.

type: date


**`citrix_cms.health.machine.details.Catalog.ModifiedDate`**
:   The date when the catalog was last modified.

type: date


## DesktopGroup [_DesktopGroup]

The desktop group details of the machine.

**`citrix_cms.health.machine.details.DesktopGroup.ID`**
:   The unique identifier of the desktop group.

type: keyword


**`citrix_cms.health.machine.details.DesktopGroup.Name`**
:   The name of the desktop group.

type: keyword


**`citrix_cms.health.machine.details.DesktopGroup.IsRemotePC`**
:   Indicates if the desktop group is a remote PC.

type: boolean


**`citrix_cms.health.machine.details.DesktopGroup.DesktopKind`**
:   The kind of desktop in the group.

type: long


**`citrix_cms.health.machine.details.DesktopGroup.LifecycleState`**
:   The lifecycle state of the desktop group.

type: long


**`citrix_cms.health.machine.details.DesktopGroup.SessionSupport`**
:   The session support type of the desktop group.

type: long


**`citrix_cms.health.machine.details.DesktopGroup.DeliveryType`**
:   The delivery type of the desktop group.

type: long


**`citrix_cms.health.machine.details.DesktopGroup.IsInMaintenanceMode`**
:   Indicates if the desktop group is in maintenance mode.

type: boolean


**`citrix_cms.health.machine.details.DesktopGroup.MachineCost`**
:   The cost of the machine in the desktop group.

type: float


**`citrix_cms.health.machine.details.DesktopGroup.AutoscaleTagID`**
:   The ID of the autoscale tag.

type: long


**`citrix_cms.health.machine.details.DesktopGroup.CreatedDate`**
:   The date when the desktop group was created.

type: date


**`citrix_cms.health.machine.details.DesktopGroup.ModifiedDate`**
:   The date when the desktop group was last modified.

type: date


## Hypervisor [_Hypervisor]

The hypervisor details of the machine.

**`citrix_cms.health.machine.details.Hypervisor.ID`**
:   The unique identifier of the hypervisor.

type: keyword


**`citrix_cms.health.machine.details.Hypervisor.Name`**
:   The name of the hypervisor.

type: keyword


**`citrix_cms.health.machine.details.Hypervisor.Type`**
:   The type of the hypervisor.

type: keyword


**`citrix_cms.health.machine.details.Hypervisor.LifecycleState`**
:   The lifecycle state of the hypervisor.

type: long


**`citrix_cms.health.machine.details.Hypervisor.CreatedDate`**
:   The date when the hypervisor was created.

type: date


**`citrix_cms.health.machine.details.Hypervisor.ModifiedDate`**
:   The date when the hypervisor was last modified.

type: date


## MachineCost [_MachineCost]

The cost details of the machine.

**`citrix_cms.health.machine.details.MachineCost.MachineID`**
:   The unique identifier of the machine.

type: keyword


**`citrix_cms.health.machine.details.MachineCost.SpecID`**
:   The specification ID of the machine.

type: long


**`citrix_cms.health.machine.details.MachineCost.CostPerHour`**
:   The cost per hour of the machine.

type: float


**`citrix_cms.health.machine.details.MachineCost.PowerOnComputeCostPerHour`**
:   The compute cost per hour when powered on.

type: float


**`citrix_cms.health.machine.details.MachineCost.PowerOnStorageCostPerHour`**
:   The storage cost per hour when powered on.

type: float


**`citrix_cms.health.machine.details.MachineCost.PowerOffStorageCostPerHour`**
:   The storage cost per hour when powered off.

type: float


**`citrix_cms.health.machine.details.MachineCost.CreatedDate`**
:   The date when the cost details were created.

type: date


**`citrix_cms.health.machine.details.MachineCost.ModifiedDate`**
:   The date when the cost details were last modified.

type: date


**`citrix_cms.health.machine.metric.details.MachineID`**
:   The unique identifier of the machine.

type: keyword


**`citrix_cms.health.machine.metric.details.CollectedDate`**
:   The date when the metrics were collected.

type: date


**`citrix_cms.health.machine.metric.details.Iops`**
:   The input/output operations per second (IOPS) of the machine.

type: float


**`citrix_cms.health.machine.metric.details.Latency`**
:   The latency of the machine in milliseconds.

type: float


**`citrix_cms.health.machine.summaries.ID`**
:   The unique identifier of the machine summary.

type: long


**`citrix_cms.health.machine.summaries.SummaryDate`**
:   The date of the machine summary.

type: date


**`citrix_cms.health.machine.summaries.DesktopGroupID`**
:   The ID of the desktop group associated with the machine summary.

type: keyword


**`citrix_cms.health.machine.summaries.PoweredOnMachinesCount`**
:   The count of powered-on machines.

type: long


**`citrix_cms.health.machine.summaries.RegisteredMachinesCount`**
:   The count of registered machines.

type: long


**`citrix_cms.health.machine.summaries.MachinesInMaintenanceModeCount`**
:   The count of machines in maintenance mode.

type: long


**`citrix_cms.health.machine.summaries.MachinesCount`**
:   The total count of machines.

type: long


**`citrix_cms.health.machine.summaries.Granularity`**
:   The granularity level of the machine summary.

type: long


**`citrix_cms.health.machine.summaries.CreatedDate`**
:   The date when the machine summary was created.

type: date


**`citrix_cms.health.machine.summaries.ModifiedDate`**
:   The date when the machine summary was last modified.

type: date


**`citrix_cms.health.resource.utilization.MachineID`**
:   The unique identifier of the machine.

type: keyword


**`citrix_cms.health.resource.utilization.CollectedDate`**
:   The date when the resource utilization metrics were collected.

type: date


**`citrix_cms.health.resource.utilization.PercentCPU`**
:   The percentage of CPU utilization.

type: float


**`citrix_cms.health.resource.utilization.UsedMemory`**
:   The amount of memory used in megabytes.

type: float


**`citrix_cms.health.resource.utilization.TotalMemory`**
:   The total memory available in megabytes.

type: float


**`citrix_cms.health.resource.utilization.CreatedDate`**
:   The date when the resource utilization record was created.

type: date


**`citrix_cms.health.resource.utilization.ModifiedDate`**
:   The date when the resource utilization record was last modified.

type: date


**`citrix_cms.health.resource.utilization.DesktopGroupID`**
:   The ID of the desktop group associated with the resource utilization.

type: keyword


**`citrix_cms.health.resource.utilization.summary.ID`**
:   The unique identifier of the resource utilization summary.

type: long


**`citrix_cms.health.resource.utilization.summary.SummaryDate`**
:   The date of the resource utilization summary.

type: date


**`citrix_cms.health.resource.utilization.summary.DesktopGroupID`**
:   The ID of the desktop group associated with the resource utilization summary.

type: keyword


**`citrix_cms.health.resource.utilization.summary.ConnectedSessionCount`**
:   The count of connected sessions.

type: long


**`citrix_cms.health.resource.utilization.summary.DisconnectedSessionCount`**
:   The count of disconnected sessions.

type: long


**`citrix_cms.health.resource.utilization.summary.ConcurrentSessionCount`**
:   The count of concurrent sessions.

type: long


**`citrix_cms.health.resource.utilization.summary.TotalLogOnDuration`**
:   The total duration of logon activities.

type: float


**`citrix_cms.health.resource.utilization.summary.TotalLogOnCount`**
:   The total count of logon activities.

type: long


**`citrix_cms.health.resource.utilization.summary.Granularity`**
:   The granularity level of the resource utilization summary.

type: long


**`citrix_cms.health.resource.utilization.summary.CreatedDate`**
:   The date when the resource utilization summary was created.

type: date


**`citrix_cms.health.resource.utilization.summary.ModifiedDate`**
:   The date when the resource utilization summary was last modified.

type: date


**`citrix_cms.health.session.metric.details.ID`**
:   The unique identifier of the session metric.

type: long


**`citrix_cms.health.session.metric.details.CollectedDate`**
:   The date when the metrics were collected.

type: date


**`citrix_cms.health.session.metric.details.IcaRttMS`**
:   The ICA round-trip time in milliseconds.

type: float


**`citrix_cms.health.session.metric.details.IcaLatency`**
:   The ICA latency in milliseconds.

type: float


**`citrix_cms.health.session.metric.details.ClientL7Latency`**
:   The client layer 7 latency in milliseconds.

type: float


**`citrix_cms.health.session.metric.details.ServerL7Latency`**
:   The server layer 7 latency in milliseconds.

type: float


**`citrix_cms.health.session.metric.details.SessionID`**
:   The unique identifier of the session.

type: keyword


**`citrix_cms.health.session.metric.details.CreatedDate`**
:   The date when the session metric was created.

type: date


**`citrix_cms.health.session.metric.details.ModifiedDate`**
:   The date when the session metric was last modified.

type: date


## Session [_Session]

Details of the session.

**`citrix_cms.health.session.metric.details.Session.SessionKey`**
:   The session key.

type: keyword


**`citrix_cms.health.session.metric.details.Session.StartDate`**
:   The start date of the session.

type: date


**`citrix_cms.health.session.metric.details.Session.LogOnDuration`**
:   The duration of the logon in seconds.

type: long


**`citrix_cms.health.session.metric.details.Session.EndDate`**
:   The end date of the session.

type: date


**`citrix_cms.health.session.metric.details.Session.ExitCode`**
:   The exit code of the session.

type: long


**`citrix_cms.health.session.metric.details.Session.FailureID`**
:   The failure ID associated with the session.

type: long


**`citrix_cms.health.session.metric.details.Session.FailureDate`**
:   The date of the failure.

type: date


**`citrix_cms.health.session.metric.details.Session.ConnectionState`**
:   The connection state of the session.

type: long


**`citrix_cms.health.session.metric.details.Session.SessionIdleTime`**
:   The idle time of the session.

type: date


**`citrix_cms.health.session.metric.details.Session.ConnectionStateChangeDate`**
:   The date when the connection state last changed.

type: date


**`citrix_cms.health.session.metric.details.Session.LifecycleState`**
:   The lifecycle state of the session.

type: long


**`citrix_cms.health.session.metric.details.Session.CurrentConnectionID`**
:   The ID of the current connection.

type: long


**`citrix_cms.health.session.metric.details.Session.UserID`**
:   The ID of the user associated with the session.

type: long


**`citrix_cms.health.session.metric.details.Session.MachineID`**
:   The ID of the machine associated with the session.

type: keyword


**`citrix_cms.health.session.metric.details.Session.SessionType`**
:   The type of the session.

type: long


**`citrix_cms.health.session.metric.details.Session.IsAnonymous`**
:   Indicates if the session is anonymous.

type: boolean


**`citrix_cms.health.session.metric.details.Session.PublishedDesktopID`**
:   The ID of the published desktop.

type: long


**`citrix_cms.health.session.metric.details.Session.CreatedDate`**
:   The date when the session was created.

type: date


**`citrix_cms.health.session.metric.details.Session.ModifiedDate`**
:   The date when the session was last modified.

type: date


**`citrix_cms.health.session.activity.summaries.ID`**
:   The unique identifier of the session activity summary.

type: long


**`citrix_cms.health.session.activity.summaries.SummaryDate`**
:   The date of the session activity summary.

type: date


**`citrix_cms.health.session.activity.summaries.DesktopGroupID`**
:   The ID of the desktop group associated with the session activity summary.

type: keyword


**`citrix_cms.health.session.activity.summaries.ConnectedSessionCount`**
:   The count of connected sessions.

type: long


**`citrix_cms.health.session.activity.summaries.DisconnectedSessionCount`**
:   The count of disconnected sessions.

type: long


**`citrix_cms.health.session.activity.summaries.ConcurrentSessionCount`**
:   The count of concurrent sessions.

type: long


**`citrix_cms.health.session.activity.summaries.TotalLogOnDuration`**
:   The total duration of logon activities.

type: float


**`citrix_cms.health.session.activity.summaries.TotalLogOnCount`**
:   The total count of logon activities.

type: long


**`citrix_cms.health.session.activity.summaries.Granularity`**
:   The granularity level of the session activity summary.

type: long


**`citrix_cms.health.session.activity.summaries.CreatedDate`**
:   The date when the session activity summary was created.

type: date


**`citrix_cms.health.session.activity.summaries.ModifiedDate`**
:   The date when the session activity summary was last modified.

type: date


**`citrix_cms.health.session.details.SessionKey`**
:   The session key.

type: keyword


**`citrix_cms.health.session.details.StartDate`**
:   The start date of the session.

type: date


**`citrix_cms.health.session.details.LogOnDuration`**
:   The duration of the logon in seconds.

type: long


**`citrix_cms.health.session.details.EndDate`**
:   The end date of the session.

type: date


**`citrix_cms.health.session.details.ExitCode`**
:   The exit code of the session.

type: long


**`citrix_cms.health.session.details.FailureID`**
:   The failure ID associated with the session.

type: long


**`citrix_cms.health.session.details.FailureDate`**
:   The date of the failure.

type: date


**`citrix_cms.health.session.details.ConnectionState`**
:   The connection state of the session.

type: long


**`citrix_cms.health.session.details.SessionIdleTime`**
:   The idle time of the session.

type: date


**`citrix_cms.health.session.details.ConnectionStateChangeDate`**
:   The date when the connection state last changed.

type: date


**`citrix_cms.health.session.details.LifecycleState`**
:   The lifecycle state of the session.

type: long


**`citrix_cms.health.session.details.CurrentConnectionID`**
:   The ID of the current connection.

type: long


**`citrix_cms.health.session.details.UserID`**
:   The ID of the user associated with the session.

type: long


**`citrix_cms.health.session.details.MachineID`**
:   The ID of the machine associated with the session.

type: keyword


**`citrix_cms.health.session.details.SessionType`**
:   The type of the session.

type: long


**`citrix_cms.health.session.details.IsAnonymous`**
:   Indicates if the session is anonymous.

type: boolean


**`citrix_cms.health.session.details.PublishedDesktopID`**
:   The ID of the published desktop.

type: long


**`citrix_cms.health.session.details.CreatedDate`**
:   The date when the session was created.

type: date


**`citrix_cms.health.session.details.ModifiedDate`**
:   The date when the session was last modified.

type: date


## Failure [_Failure]

Details of the failure associated with the session.

**`citrix_cms.health.session.details.Failure.ID`**
:   The unique identifier of the failure.

type: long


**`citrix_cms.health.session.details.Failure.ConnectionFailureEnumValue`**
:   The connection failure enum value.

type: long


**`citrix_cms.health.session.details.Failure.Category`**
:   The category of the failure.

type: long


**`citrix_cms.health.session.details.Failure.CreatedDate`**
:   The date when the failure was created.

type: date


**`citrix_cms.health.session.details.Failure.ModifiedDate`**
:   The date when the failure was last modified.

type: date


## CurrentConnection [_CurrentConnection]

Details of the current connection associated with the session.

**`citrix_cms.health.session.details.CurrentConnection.ID`**
:   The unique identifier of the current connection.

type: long


**`citrix_cms.health.session.details.CurrentConnection.ClientName`**
:   The name of the client.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ClientAddress`**
:   The address of the client.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ClientPublicIP`**
:   The public IP address of the client.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ClientVersion`**
:   The version of the client.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ClientPlatform`**
:   The platform of the client.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ClientISP`**
:   The ISP of the client.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ClientLocationCountry`**
:   The country of the client's location.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ClientLocationCity`**
:   The city of the client's location.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ConnectedViaHostName`**
:   The hostname through which the client is connected.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ConnectedViaIPAddress`**
:   The IP address through which the client is connected.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.LaunchedViaHostName`**
:   The hostname through which the session was launched.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.LaunchedViaIPAddress`**
:   The IP address through which the session was launched.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.IsReconnect`**
:   Indicates if the session is a reconnect.

type: boolean


**`citrix_cms.health.session.details.CurrentConnection.IsSecureIca`**
:   Indicates if the session uses Secure ICA.

type: boolean


**`citrix_cms.health.session.details.CurrentConnection.Protocol`**
:   The protocol used for the session.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.LogOnStartDate`**
:   The start date of the logon process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.LogOnEndDate`**
:   The end date of the logon process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.BrokeringDuration`**
:   The duration of the brokering process.

type: long


**`citrix_cms.health.session.details.CurrentConnection.BrokeringDate`**
:   The date of the brokering process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.DisconnectCode`**
:   The disconnect code of the session.

type: long


**`citrix_cms.health.session.details.CurrentConnection.DisconnectDate`**
:   The date of the disconnection.

type: date


**`citrix_cms.health.session.details.CurrentConnection.VMStartStartDate`**
:   The start date of the virtual machine start process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.VMPoweredOnDate`**
:   The date when the virtual machine was powered on.

type: date


**`citrix_cms.health.session.details.CurrentConnection.VMStartEndDate`**
:   The end date of the virtual machine start process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.ClientSessionValidateDate`**
:   The date when the client session was validated.

type: date


**`citrix_cms.health.session.details.CurrentConnection.ServerSessionValidateDate`**
:   The date when the server session was validated.

type: date


**`citrix_cms.health.session.details.CurrentConnection.EstablishmentDate`**
:   The date when the session was established.

type: date


**`citrix_cms.health.session.details.CurrentConnection.HdxStartDate`**
:   The start date of the HDX process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.HdxEndDate`**
:   The end date of the HDX process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.AuthenticationDuration`**
:   The duration of the authentication process.

type: long


**`citrix_cms.health.session.details.CurrentConnection.GpoStartDate`**
:   The start date of the GPO process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.GpoEndDate`**
:   The end date of the GPO process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.LogOnScriptsStartDate`**
:   The start date of the logon scripts process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.LogOnScriptsEndDate`**
:   The end date of the logon scripts process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.ProfileLoadStartDate`**
:   The start date of the profile load process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.ProfileLoadEndDate`**
:   The end date of the profile load process.

type: date


**`citrix_cms.health.session.details.CurrentConnection.InteractiveStartDate`**
:   The start date of the interactive session.

type: date


**`citrix_cms.health.session.details.CurrentConnection.InteractiveEndDate`**
:   The end date of the interactive session.

type: date


**`citrix_cms.health.session.details.CurrentConnection.SessionKey`**
:   The session key.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.CreatedDate`**
:   The date when the connection was created.

type: date


**`citrix_cms.health.session.details.CurrentConnection.ModifiedDate`**
:   The date when the connection was last modified.

type: date


## ConnectionFailureLog [_ConnectionFailureLog]

Details of the connection failure log.

**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.ID`**
:   The unique identifier of the connection failure log.

type: long


**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.SessionKey`**
:   The session key associated with the failure log.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.FailureDate`**
:   The date of the failure.

type: date


**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.UserID`**
:   The ID of the user associated with the failure log.

type: long


**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.MachineID`**
:   The ID of the machine associated with the failure log.

type: keyword


**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.ConnectionFailureEnumValue`**
:   The connection failure enum value.

type: long


**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.CreatedDate`**
:   The date when the failure log was created.

type: date


**`citrix_cms.health.session.details.CurrentConnection.ConnectionFailureLog.ModifiedDate`**
:   The date when the failure log was last modified.

type: date


## User [_User]

Details of the user associated with the session.

**`citrix_cms.health.session.details.User.ID`**
:   The unique identifier of the user.

type: long


**`citrix_cms.health.session.details.User.Sid`**
:   The security identifier of the user.

type: keyword


**`citrix_cms.health.session.details.User.Upn`**
:   The UPN of the user.

type: keyword


**`citrix_cms.health.session.details.User.UserName`**
:   The username of the user.

type: keyword


**`citrix_cms.health.session.details.User.FullName`**
:   The full name of the user.

type: keyword


**`citrix_cms.health.session.details.User.Domain`**
:   The domain of the user.

type: keyword


**`citrix_cms.health.session.details.User.CreatedDate`**
:   The date when the user was created.

type: date


**`citrix_cms.health.session.details.User.ModifiedDate`**
:   The date when the user was last modified.

type: date


## Machine [_Machine]

Details of the machine associated with the session.

**`citrix_cms.health.session.details.Machine.ID`**
:   The unique identifier of the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.Sid`**
:   The security identifier of the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.Name`**
:   The name of the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.DNSName`**
:   The DNS name of the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.LifecycleState`**
:   The lifecycle state of the machine.

type: long


**`citrix_cms.health.session.details.Machine.IPAddress`**
:   The IP address of the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.HostedMachineID`**
:   The ID of the hosted machine.

type: keyword


**`citrix_cms.health.session.details.Machine.HostingServerName`**
:   The name of the hosting server.

type: keyword


**`citrix_cms.health.session.details.Machine.HostedMachineName`**
:   The name of the hosted machine.

type: keyword


**`citrix_cms.health.session.details.Machine.IsAssigned`**
:   Indicates if the machine is assigned.

type: boolean


**`citrix_cms.health.session.details.Machine.IsInMaintenanceMode`**
:   Indicates if the machine is in maintenance mode.

type: boolean


**`citrix_cms.health.session.details.Machine.IsPendingUpdate`**
:   Indicates if the machine has a pending update.

type: boolean


**`citrix_cms.health.session.details.Machine.AgentVersion`**
:   The version of the agent installed on the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.AssociatedUserFullNames`**
:   The full names of associated users.

type: keyword


**`citrix_cms.health.session.details.Machine.AssociatedUserNames`**
:   The usernames of associated users.

type: keyword


**`citrix_cms.health.session.details.Machine.AssociatedUserUPNs`**
:   The UPNs of associated users.

type: keyword


**`citrix_cms.health.session.details.Machine.CurrentRegistrationState`**
:   The current registration state of the machine.

type: long


**`citrix_cms.health.session.details.Machine.RegistrationStateChangeDate`**
:   The date when the registration state last changed.

type: date


**`citrix_cms.health.session.details.Machine.LastDeregisteredCode`**
:   The code of the last deregistration event.

type: long


**`citrix_cms.health.session.details.Machine.LastDeregisteredDate`**
:   The date of the last deregistration event.

type: date


**`citrix_cms.health.session.details.Machine.CurrentPowerState`**
:   The current power state of the machine.

type: long


**`citrix_cms.health.session.details.Machine.CurrentSessionCount`**
:   The current session count on the machine.

type: long


**`citrix_cms.health.session.details.Machine.ControllerDNSName`**
:   The DNS name of the controller.

type: keyword


**`citrix_cms.health.session.details.Machine.PoweredOnDate`**
:   The date when the machine was powered on.

type: date


**`citrix_cms.health.session.details.Machine.PowerStateChangeDate`**
:   The date when the power state last changed.

type: date


**`citrix_cms.health.session.details.Machine.FunctionalLevel`**
:   The functional level of the machine.

type: long


**`citrix_cms.health.session.details.Machine.FailureDate`**
:   The date of the last failure event.

type: date


**`citrix_cms.health.session.details.Machine.WindowsConnectionSetting`**
:   The Windows connection setting of the machine.

type: long


**`citrix_cms.health.session.details.Machine.IsPreparing`**
:   Indicates if the machine is preparing.

type: boolean


**`citrix_cms.health.session.details.Machine.FaultState`**
:   The fault state of the machine.

type: long


**`citrix_cms.health.session.details.Machine.OSType`**
:   The operating system type of the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.CurrentLoadIndexID`**
:   The ID of the current load index.

type: long


**`citrix_cms.health.session.details.Machine.CatalogID`**
:   The ID of the catalog.

type: keyword


**`citrix_cms.health.session.details.Machine.DesktopGroupID`**
:   The ID of the desktop group.

type: keyword


**`citrix_cms.health.session.details.Machine.HypervisorID`**
:   The ID of the hypervisor.

type: keyword


**`citrix_cms.health.session.details.Machine.LastPowerActionCompletedDate`**
:   The date when the last power action was completed.

type: date


**`citrix_cms.health.session.details.Machine.LastUpgradeState`**
:   The last upgrade state of the machine.

type: long


**`citrix_cms.health.session.details.Machine.LastUpgradeStateChangeDate`**
:   The date when the last upgrade state changed.

type: date


**`citrix_cms.health.session.details.Machine.Hash`**
:   The hash of the machine.

type: keyword


**`citrix_cms.health.session.details.Machine.MachineRole`**
:   The role of the machine.

type: long


**`citrix_cms.health.session.details.Machine.CreatedDate`**
:   The date when the machine was created.

type: date


**`citrix_cms.health.session.details.Machine.ModifiedDate`**
:   The date when the machine was last modified.

type: date


## SessionMetrics [_SessionMetrics]

Metrics related to the session.

**`citrix_cms.health.session.details.SessionMetrics.ID`**
:   The unique identifier of the session metric.

type: long


**`citrix_cms.health.session.details.SessionMetrics.CollectedDate`**
:   The date when the metrics were collected.

type: date


**`citrix_cms.health.session.details.SessionMetrics.IcaRttMS`**
:   The ICA round-trip time in milliseconds.

type: long


**`citrix_cms.health.session.details.SessionMetrics.IcaLatency`**
:   The ICA latency in milliseconds.

type: long


**`citrix_cms.health.session.details.SessionMetrics.ClientL7Latency`**
:   The client layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.session.details.SessionMetrics.ServerL7Latency`**
:   The server layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.session.details.SessionMetrics.SessionID`**
:   The unique identifier of the session.

type: keyword


**`citrix_cms.health.session.details.SessionMetrics.CreatedDate`**
:   The date when the session metric was created.

type: date


**`citrix_cms.health.session.details.SessionMetrics.ModifiedDate`**
:   The date when the session metric was last modified.

type: date


## SessionRecordingServer [_SessionRecordingServer]

Details of the session recording server.

**`citrix_cms.health.session.details.SessionRecordingServer.SessionKey`**
:   The session key associated with the recording server.

type: keyword


**`citrix_cms.health.session.details.SessionRecordingServer.SessionRecordingServerName`**
:   The name of the session recording server.

type: keyword


**`citrix_cms.health.session.details.SessionRecordingServer.CreatedDate`**
:   The date when the session recording server was created.

type: date


## PublishedDesktopName [_PublishedDesktopName]

Details of the published desktop name.

**`citrix_cms.health.session.details.PublishedDesktopName.ID`**
:   The unique identifier of the published desktop.

type: long


**`citrix_cms.health.session.details.PublishedDesktopName.PublishedName`**
:   The name of the published desktop.

type: keyword


## SessionMetricsLatest [_SessionMetricsLatest]

Latest metrics related to the session.

**`citrix_cms.health.session.details.SessionMetricsLatest.SessionKey`**
:   The session key associated with the latest metrics.

type: keyword


**`citrix_cms.health.session.details.SessionMetricsLatest.ClientL7Latency`**
:   The client layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.session.details.SessionMetricsLatest.CloudConnectorName`**
:   The name of the cloud connector.

type: keyword


**`citrix_cms.health.session.details.SessionMetricsLatest.CreatedDate`**
:   The date when the latest session metrics were created.

type: date


**`citrix_cms.health.session.details.SessionMetricsLatest.EdtMtu`**
:   The EDT MTU value.

type: keyword


**`citrix_cms.health.session.details.SessionMetricsLatest.GatewayPopName`**
:   The name of the gateway point of presence.

type: keyword


**`citrix_cms.health.session.details.SessionMetricsLatest.HdxConnectionType`**
:   The type of HDX connection.

type: long


**`citrix_cms.health.session.details.SessionMetricsLatest.HdxProtocolName`**
:   The name of the HDX protocol.

type: keyword


**`citrix_cms.health.session.details.SessionMetricsLatest.ModifiedDate`**
:   The date when the latest session metrics were last modified.

type: date


**`citrix_cms.health.session.details.SessionMetricsLatest.ServerL7Latency`**
:   The server layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.session.failure.details.SessionKey`**
:   The session key.

type: keyword


**`citrix_cms.health.session.failure.details.StartDate`**
:   The start date of the session.

type: date


**`citrix_cms.health.session.failure.details.LogOnDuration`**
:   The duration of the logon in seconds.

type: long


**`citrix_cms.health.session.failure.details.EndDate`**
:   The end date of the session.

type: date


**`citrix_cms.health.session.failure.details.ExitCode`**
:   The exit code of the session.

type: long


**`citrix_cms.health.session.failure.details.FailureID`**
:   The failure ID associated with the session.

type: long


**`citrix_cms.health.session.failure.details.FailureDate`**
:   The date of the failure.

type: date


**`citrix_cms.health.session.failure.details.ConnectionState`**
:   The connection state of the session.

type: long


**`citrix_cms.health.session.failure.details.SessionIdleTime`**
:   The idle time of the session.

type: date


**`citrix_cms.health.session.failure.details.ConnectionStateChangeDate`**
:   The date when the connection state last changed.

type: date


**`citrix_cms.health.session.failure.details.LifecycleState`**
:   The lifecycle state of the session.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnectionID`**
:   The ID of the current connection.

type: long


**`citrix_cms.health.session.failure.details.UserID`**
:   The ID of the user associated with the session.

type: long


**`citrix_cms.health.session.failure.details.MachineID`**
:   The ID of the machine associated with the session.

type: keyword


**`citrix_cms.health.session.failure.details.SessionType`**
:   The type of the session.

type: long


**`citrix_cms.health.session.failure.details.IsAnonymous`**
:   Indicates if the session is anonymous.

type: boolean


**`citrix_cms.health.session.failure.details.PublishedDesktopID`**
:   The ID of the published desktop.

type: long


**`citrix_cms.health.session.failure.details.CreatedDate`**
:   The date when the session was created.

type: date


**`citrix_cms.health.session.failure.details.ModifiedDate`**
:   The date when the session was last modified.

type: date


## Failure [_Failure]

Details of the failure associated with the session.

**`citrix_cms.health.session.failure.details.Failure.ID`**
:   The unique identifier of the failure.

type: long


**`citrix_cms.health.session.failure.details.Failure.ConnectionFailureEnumValue`**
:   The connection failure enum value.

type: long


**`citrix_cms.health.session.failure.details.Failure.Category`**
:   The category of the failure.

type: long


**`citrix_cms.health.session.failure.details.Failure.CreatedDate`**
:   The date when the failure was created.

type: date


**`citrix_cms.health.session.failure.details.Failure.ModifiedDate`**
:   The date when the failure was last modified.

type: date


## CurrentConnection [_CurrentConnection]

Details of the current connection associated with the session.

**`citrix_cms.health.session.failure.details.CurrentConnection.ID`**
:   The unique identifier of the current connection.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientName`**
:   The name of the client.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientAddress`**
:   The address of the client.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientPublicIP`**
:   The public IP address of the client.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientVersion`**
:   The version of the client.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientPlatform`**
:   The platform of the client.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientISP`**
:   The ISP of the client.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientLocationCountry`**
:   The country of the client's location.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientLocationCity`**
:   The city of the client's location.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectedViaHostName`**
:   The hostname through which the client is connected.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectedViaIPAddress`**
:   The IP address through which the client is connected.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.LaunchedViaHostName`**
:   The hostname through which the session was launched.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.LaunchedViaIPAddress`**
:   The IP address through which the session was launched.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.IsReconnect`**
:   Indicates if the session is a reconnect.

type: boolean


**`citrix_cms.health.session.failure.details.CurrentConnection.IsSecureIca`**
:   Indicates if the session uses Secure ICA.

type: boolean


**`citrix_cms.health.session.failure.details.CurrentConnection.Protocol`**
:   The protocol used for the session.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.LogOnStartDate`**
:   The start date of the logon process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.LogOnEndDate`**
:   The end date of the logon process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.BrokeringDuration`**
:   The duration of the brokering process.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnection.BrokeringDate`**
:   The date of the brokering process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.DisconnectCode`**
:   The disconnect code of the session.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnection.DisconnectDate`**
:   The date of the disconnection.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.VMStartStartDate`**
:   The start date of the virtual machine start process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.VMPoweredOnDate`**
:   The date when the virtual machine was powered on.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.VMStartEndDate`**
:   The end date of the virtual machine start process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.ClientSessionValidateDate`**
:   The date when the client session was validated.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.ServerSessionValidateDate`**
:   The date when the server session was validated.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.EstablishmentDate`**
:   The date when the session was established.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.HdxStartDate`**
:   The start date of the HDX process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.HdxEndDate`**
:   The end date of the HDX process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.AuthenticationDuration`**
:   The duration of the authentication process.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnection.GpoStartDate`**
:   The start date of the GPO process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.GpoEndDate`**
:   The end date of the GPO process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.LogOnScriptsStartDate`**
:   The start date of the logon scripts process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.LogOnScriptsEndDate`**
:   The end date of the logon scripts process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.ProfileLoadStartDate`**
:   The start date of the profile load process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.ProfileLoadEndDate`**
:   The end date of the profile load process.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.InteractiveStartDate`**
:   The start date of the interactive session.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.InteractiveEndDate`**
:   The end date of the interactive session.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.SessionKey`**
:   The session key.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.CreatedDate`**
:   The date when the connection was created.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.ModifiedDate`**
:   The date when the connection was last modified.

type: date


## ConnectionFailureLog [_ConnectionFailureLog]

Details of the connection failure log.

**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.ID`**
:   The unique identifier of the connection failure log.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.SessionKey`**
:   The session key associated with the failure log.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.FailureDate`**
:   The date of the failure.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.UserID`**
:   The ID of the user associated with the failure log.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.MachineID`**
:   The ID of the machine associated with the failure log.

type: keyword


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.ConnectionFailureEnumValue`**
:   The connection failure enum value.

type: long


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.CreatedDate`**
:   The date when the failure log was created.

type: date


**`citrix_cms.health.session.failure.details.CurrentConnection.ConnectionFailureLog.ModifiedDate`**
:   The date when the failure log was last modified.

type: date


## User [_User]

Details of the user associated with the session.

**`citrix_cms.health.session.failure.details.User.ID`**
:   The unique identifier of the user.

type: long


**`citrix_cms.health.session.failure.details.User.Sid`**
:   The security identifier of the user.

type: keyword


**`citrix_cms.health.session.failure.details.User.Upn`**
:   The UPN of the user.

type: keyword


**`citrix_cms.health.session.failure.details.User.UserName`**
:   The username of the user.

type: keyword


**`citrix_cms.health.session.failure.details.User.FullName`**
:   The full name of the user.

type: keyword


**`citrix_cms.health.session.failure.details.User.Domain`**
:   The domain of the user.

type: keyword


**`citrix_cms.health.session.failure.details.User.CreatedDate`**
:   The date when the user was created.

type: date


**`citrix_cms.health.session.failure.details.User.ModifiedDate`**
:   The date when the user was last modified.

type: date


## Machine [_Machine]

Details of the machine associated with the session.

**`citrix_cms.health.session.failure.details.Machine.ID`**
:   The unique identifier of the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.Sid`**
:   The security identifier of the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.Name`**
:   The name of the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.DNSName`**
:   The DNS name of the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.LifecycleState`**
:   The lifecycle state of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.IPAddress`**
:   The IP address of the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.HostedMachineID`**
:   The ID of the hosted machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.HostingServerName`**
:   The name of the hosting server.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.HostedMachineName`**
:   The name of the hosted machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.IsAssigned`**
:   Indicates if the machine is assigned.

type: boolean


**`citrix_cms.health.session.failure.details.Machine.IsInMaintenanceMode`**
:   Indicates if the machine is in maintenance mode.

type: boolean


**`citrix_cms.health.session.failure.details.Machine.IsPendingUpdate`**
:   Indicates if the machine has a pending update.

type: boolean


**`citrix_cms.health.session.failure.details.Machine.AgentVersion`**
:   The version of the agent installed on the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.AssociatedUserFullNames`**
:   The full names of associated users.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.AssociatedUserNames`**
:   The usernames of associated users.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.AssociatedUserUPNs`**
:   The UPNs of associated users.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.CurrentRegistrationState`**
:   The current registration state of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.RegistrationStateChangeDate`**
:   The date when the registration state last changed.

type: date


**`citrix_cms.health.session.failure.details.Machine.LastDeregisteredCode`**
:   The code of the last deregistration event.

type: long


**`citrix_cms.health.session.failure.details.Machine.LastDeregisteredDate`**
:   The date of the last deregistration event.

type: date


**`citrix_cms.health.session.failure.details.Machine.CurrentPowerState`**
:   The current power state of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.CurrentSessionCount`**
:   The current session count on the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.ControllerDNSName`**
:   The DNS name of the controller.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.PoweredOnDate`**
:   The date when the machine was powered on.

type: date


**`citrix_cms.health.session.failure.details.Machine.PowerStateChangeDate`**
:   The date when the power state last changed.

type: date


**`citrix_cms.health.session.failure.details.Machine.FunctionalLevel`**
:   The functional level of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.FailureDate`**
:   The date of the last failure event.

type: date


**`citrix_cms.health.session.failure.details.Machine.WindowsConnectionSetting`**
:   The Windows connection setting of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.IsPreparing`**
:   Indicates if the machine is preparing.

type: boolean


**`citrix_cms.health.session.failure.details.Machine.FaultState`**
:   The fault state of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.OSType`**
:   The operating system type of the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.CurrentLoadIndexID`**
:   The ID of the current load index.

type: long


**`citrix_cms.health.session.failure.details.Machine.CatalogID`**
:   The ID of the catalog.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.DesktopGroupID`**
:   The ID of the desktop group.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.HypervisorID`**
:   The ID of the hypervisor.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.LastPowerActionCompletedDate`**
:   The date when the last power action was completed.

type: date


**`citrix_cms.health.session.failure.details.Machine.LastUpgradeState`**
:   The last upgrade state of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.LastUpgradeStateChangeDate`**
:   The date when the last upgrade state changed.

type: date


**`citrix_cms.health.session.failure.details.Machine.Hash`**
:   The hash of the machine.

type: keyword


**`citrix_cms.health.session.failure.details.Machine.MachineRole`**
:   The role of the machine.

type: long


**`citrix_cms.health.session.failure.details.Machine.CreatedDate`**
:   The date when the machine was created.

type: date


**`citrix_cms.health.session.failure.details.Machine.ModifiedDate`**
:   The date when the machine was last modified.

type: date


## SessionMetrics [_SessionMetrics]

Metrics related to the session.

**`citrix_cms.health.session.failure.details.SessionMetrics.ID`**
:   The unique identifier of the session metric.

type: long


**`citrix_cms.health.session.failure.details.SessionMetrics.CollectedDate`**
:   The date when the metrics were collected.

type: date


**`citrix_cms.health.session.failure.details.SessionMetrics.IcaRttMS`**
:   The ICA round-trip time in milliseconds.

type: long


**`citrix_cms.health.session.failure.details.SessionMetrics.IcaLatency`**
:   The ICA latency in milliseconds.

type: long


**`citrix_cms.health.session.failure.details.SessionMetrics.ClientL7Latency`**
:   The client layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.session.failure.details.SessionMetrics.ServerL7Latency`**
:   The server layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.session.failure.details.SessionMetrics.SessionID`**
:   The unique identifier of the session.

type: keyword


**`citrix_cms.health.session.failure.details.SessionMetrics.CreatedDate`**
:   The date when the session metric was created.

type: date


**`citrix_cms.health.session.failure.details.SessionMetrics.ModifiedDate`**
:   The date when the session metric was last modified.

type: date


## SessionRecordingServer [_SessionRecordingServer]

Details of the session recording server.

**`citrix_cms.health.session.failure.details.SessionRecordingServer.SessionKey`**
:   The session key associated with the recording server.

type: keyword


**`citrix_cms.health.session.failure.details.SessionRecordingServer.SessionRecordingServerName`**
:   The name of the session recording server.

type: keyword


**`citrix_cms.health.session.failure.details.SessionRecordingServer.CreatedDate`**
:   The date when the session recording server was created.

type: date


## PublishedDesktopName [_PublishedDesktopName]

Details of the published desktop name.

**`citrix_cms.health.session.failure.details.PublishedDesktopName.ID`**
:   The unique identifier of the published desktop.

type: long


**`citrix_cms.health.session.failure.details.PublishedDesktopName.PublishedName`**
:   The name of the published desktop.

type: keyword


## SessionMetricsLatest [_SessionMetricsLatest]

Latest metrics related to the session.

**`citrix_cms.health.session.failure.details.SessionMetricsLatest.SessionKey`**
:   The session key associated with the latest metrics.

type: keyword


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.ClientL7Latency`**
:   The client layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.CloudConnectorName`**
:   The name of the cloud connector.

type: keyword


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.CreatedDate`**
:   The date when the latest session metrics were created.

type: date


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.EdtMtu`**
:   The EDT MTU value.

type: keyword


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.GatewayPopName`**
:   The name of the gateway point of presence.

type: keyword


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.HdxConnectionType`**
:   The type of HDX connection.

type: long


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.HdxProtocolName`**
:   The name of the HDX protocol.

type: keyword


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.ModifiedDate`**
:   The date when the latest session metrics were last modified.

type: date


**`citrix_cms.health.session.failure.details.SessionMetricsLatest.ServerL7Latency`**
:   The server layer 7 latency in milliseconds.

type: long


**`citrix_cms.health.server.os.desktop.summary.ID`**
:   The unique identifier of the desktop OS summary.

type: long


**`citrix_cms.health.server.os.desktop.summary.SummaryDate`**
:   The date of the desktop OS summary.

type: date


**`citrix_cms.health.server.os.desktop.summary.DesktopGroupId`**
:   The ID of the desktop group associated with the desktop OS summary.

type: keyword


**`citrix_cms.health.server.os.desktop.summary.PeakConcurrentInstanceCount`**
:   The peak count of concurrent instances.

type: long


**`citrix_cms.health.server.os.desktop.summary.TotalUsageDuration`**
:   The total usage duration in hours.

type: float


**`citrix_cms.health.server.os.desktop.summary.TotalLaunchesCount`**
:   The total count of launches.

type: long


**`citrix_cms.health.server.os.desktop.summary.StartingInstanceCount`**
:   The count of starting instances.

type: long


**`citrix_cms.health.server.os.desktop.summary.Granularity`**
:   The granularity level of the desktop OS summary.

type: long


**`citrix_cms.health.server.os.desktop.summary.CreatedDate`**
:   The date when the desktop OS summary was created.

type: date


**`citrix_cms.health.server.os.desktop.summary.ModifiedDate`**
:   The date when the desktop OS summary was last modified.

type: date


