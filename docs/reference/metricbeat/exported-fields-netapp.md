---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/exported-fields-netapp.html
---

% This file is generated! See scripts/generate_fields_docs.py

# netapp fields [exported-fields-netapp]

netapp module

## netapp [_netapp]



## cluster [_cluster]

cluster

**`netapp.cluster.name`**
:   Cluster name

type: keyword


**`netapp.cluster.uuid`**
:   Cluster UUID

type: keyword


**`netapp.cluster.location`**
:   Cluster location

type: keyword


**`netapp.cluster.contact`**
:   Cluster contact

type: keyword


**`netapp.cluster.version`**
:   Cluster version

type: keyword


**`netapp.cluster.dns_domains`**
:   List of DNS domains

type: keyword


**`netapp.cluster.name_servers`**
:   List of name servers

type: keyword


**`netapp.cluster.ntp_servers`**
:   List of NTP servers

type: keyword


**`netapp.cluster.management_interfaces`**
:   List of management interfaces

type: keyword


**`netapp.cluster.metric`**
:   Cluster metric information

type: object


**`netapp.cluster.statistics`**
:   Cluster statistics

type: object


**`netapp.cluster.san_optimized`**
:   Whether the cluster is SAN optimized

type: boolean


**`netapp.cluster.disaggregated`**
:   Whether the cluster is disaggregated

type: boolean


**`netapp.cluster.auto_enable_analytics`**
:   Whether analytics are auto-enabled

type: boolean


**`netapp.cluster.auto_enable_activity_tracking`**
:   Whether activity tracking is auto-enabled

type: boolean


## timezone [_timezone]

Cluster timezone

**`netapp.cluster.timezone.name`**
:   Timezone name

type: keyword


## certificate [_certificate]

Cluster certificate

**`netapp.cluster.certificate.uuid`**
:   Certificate UUID

type: keyword


## peering_policy [_peering_policy]

Cluster peering policy

**`netapp.cluster.peering_policy.minimum_passphrase_length`**
:   Minimum passphrase length

type: integer


**`netapp.cluster.peering_policy.authentication_required`**
:   Whether authentication is required

type: boolean


**`netapp.cluster.peering_policy.encryption_required`**
:   Whether encryption is required

type: boolean


## peer [_peer]

Cluster peer information

## authentication [_authentication]

Authentication details for the peer

**`netapp.cluster.peer.authentication.expiry_time`**
:   Authentication expiry time

type: keyword


**`netapp.cluster.peer.authentication.in_use`**
:   Whether authentication is in use

type: boolean


**`netapp.cluster.peer.authentication.passphrase`**
:   Authentication passphrase

type: keyword


**`netapp.cluster.peer.authentication.state`**
:   Authentication state

type: keyword


## encryption [_encryption]

Encryption details for the peer

**`netapp.cluster.peer.encryption.proposed`**
:   Whether encryption is proposed

type: boolean


**`netapp.cluster.peer.encryption.state`**
:   Encryption state

type: keyword


**`netapp.cluster.peer.initial_allowed_svms`**
:   Initial allowed SVMs (as JSON string)

type: keyword


**`netapp.cluster.peer.ip_address`**
:   Peer IP address

type: keyword


**`netapp.cluster.peer.ipspace`**
:   IPspace information

type: object


**`netapp.cluster.peer.name`**
:   Peer name

type: keyword


**`netapp.cluster.peer.peer_applications`**
:   Peer applications (as JSON string)

type: keyword


## remote [_remote]

Remote peer information

**`netapp.cluster.peer.remote.ip_addresses`**
:   Remote peer IP addresses (as JSON string)

type: keyword


**`netapp.cluster.peer.remote.name`**
:   Remote peer name

type: keyword


**`netapp.cluster.peer.remote.serial_number`**
:   Remote peer serial number

type: keyword


## status [_status]

Peer status information

**`netapp.cluster.peer.status.state`**
:   Peer status state

type: keyword


**`netapp.cluster.peer.status.update_time`**
:   Peer status update time

type: keyword


**`netapp.cluster.peer.uuid`**
:   Peer UUID

type: keyword


## version [_version]

Peer version information

**`netapp.cluster.peer.version.full`**
:   Full version string

type: keyword


**`netapp.cluster.peer.version.generation`**
:   Version generation

type: integer


**`netapp.cluster.peer.version.major`**
:   Major version

type: integer


**`netapp.cluster.peer.version.minor`**
:   Minor version

type: integer


## node [_node]

Cluster node information

**`netapp.cluster.node.uuid`**
:   Node UUID

type: keyword


**`netapp.cluster.node.name`**
:   Node name

type: keyword


**`netapp.cluster.node.serial_number`**
:   Node serial number

type: keyword


**`netapp.cluster.node.location`**
:   Node location

type: keyword


**`netapp.cluster.node.owner`**
:   Node owner

type: keyword


**`netapp.cluster.node.model`**
:   Node model

type: keyword


**`netapp.cluster.node.system_id`**
:   Node system ID

type: keyword


## version [_version]

Node version information

**`netapp.cluster.node.version.full`**
:   Full version string

type: keyword


**`netapp.cluster.node.version.generation`**
:   Version generation

type: integer


**`netapp.cluster.node.version.major`**
:   Major version

type: integer


**`netapp.cluster.node.version.minor`**
:   Minor version

type: integer


**`netapp.cluster.node.date`**
:   Node date

type: keyword


**`netapp.cluster.node.uptime`**
:   Node uptime in seconds

type: long


**`netapp.cluster.node.state`**
:   Node state

type: keyword


**`netapp.cluster.node.membership`**
:   Node membership

type: keyword


**`netapp.cluster.node.management_interfaces`**
:   List of management interfaces (as JSON string)

type: keyword


**`netapp.cluster.node.cluster_interfaces`**
:   List of cluster interfaces (as JSON string)

type: keyword


**`netapp.cluster.node.storage_configuration`**
:   Node storage configuration

type: keyword


## system_aggregate [_system_aggregate]

System aggregate information

**`netapp.cluster.node.system_aggregate.uuid`**
:   System aggregate UUID

type: keyword


**`netapp.cluster.node.system_aggregate.name`**
:   System aggregate name

type: keyword


## controller [_controller]

Node controller information

**`netapp.cluster.node.controller.board`**
:   Controller board

type: keyword


**`netapp.cluster.node.controller.memory_size`**
:   Controller memory size

type: long


**`netapp.cluster.node.controller.over_temperature`**
:   Controller over-temperature status, e.g. "normal"

type: keyword


## failed_fan [_failed_fan]

Failed fan information

**`netapp.cluster.node.controller.failed_fan.count`**
:   Number of failed fans

type: integer


## message [_message]

Failed fan message

**`netapp.cluster.node.controller.failed_fan.message.message`**
:   Message text

type: keyword


**`netapp.cluster.node.controller.failed_fan.message.code`**
:   Message code

type: keyword


## failed_power_supply [_failed_power_supply]

Failed power supply information

**`netapp.cluster.node.controller.failed_power_supply.count`**
:   Number of failed power supplies

type: integer


## message [_message]

Failed power supply message

**`netapp.cluster.node.controller.failed_power_supply.message.message`**
:   Message text

type: keyword


**`netapp.cluster.node.controller.failed_power_supply.message.code`**
:   Message code

type: keyword


## cpu [_cpu]

CPU information

**`netapp.cluster.node.controller.cpu.firmware_release`**
:   CPU firmware release

type: keyword


**`netapp.cluster.node.controller.cpu.processor`**
:   CPU processor

type: keyword


**`netapp.cluster.node.controller.cpu.count`**
:   CPU count

type: integer


## service_processor [_service_processor]

Node service processor information

**`netapp.cluster.node.service_processor.dhcp_enabled`**
:   DHCP enabled

type: boolean


**`netapp.cluster.node.service_processor.state`**
:   Service processor state

type: keyword


**`netapp.cluster.node.service_processor.mac_address`**
:   MAC address

type: keyword


**`netapp.cluster.node.service_processor.firmware_version`**
:   Firmware version

type: keyword


**`netapp.cluster.node.service_processor.link_status`**
:   Link status

type: keyword


**`netapp.cluster.node.service_processor.type`**
:   Service processor type

type: keyword


**`netapp.cluster.node.service_processor.is_ip_configured`**
:   Is IP configured

type: boolean


**`netapp.cluster.node.service_processor.autoupdate_enabled`**
:   Autoupdate enabled

type: boolean


**`netapp.cluster.node.service_processor.last_update_state`**
:   Last update state

type: keyword


## ipv4_interface [_ipv4_interface]

IPv4 interface information

**`netapp.cluster.node.service_processor.ipv4_interface.address`**
:   IPv4 address

type: keyword


**`netapp.cluster.node.service_processor.ipv4_interface.netmask`**
:   IPv4 netmask

type: keyword


**`netapp.cluster.node.service_processor.ipv4_interface.gateway`**
:   IPv4 gateway

type: keyword


**`netapp.cluster.node.service_processor.ipv4_interface.enabled`**
:   IPv4 enabled

type: boolean


**`netapp.cluster.node.service_processor.ipv4_interface.setup_state`**
:   IPv4 setup state

type: keyword


## ipv6_interface [_ipv6_interface]

IPv6 interface information

**`netapp.cluster.node.service_processor.ipv6_interface.enabled`**
:   IPv6 enabled

type: boolean


## ssh_info [_ssh_info]

SSH information

**`netapp.cluster.node.service_processor.ssh_info.allowed_addresses`**
:   Allowed SSH addresses

type: keyword


## primary [_primary]

Primary service processor information

**`netapp.cluster.node.service_processor.primary.is_current`**
:   Is current

type: boolean


**`netapp.cluster.node.service_processor.primary.state`**
:   Primary state

type: keyword


**`netapp.cluster.node.service_processor.primary.version`**
:   Primary version

type: keyword


## backup [_backup]

Backup service processor information

**`netapp.cluster.node.service_processor.backup.is_current`**
:   Is current

type: boolean


**`netapp.cluster.node.service_processor.backup.state`**
:   Backup state

type: keyword


**`netapp.cluster.node.service_processor.backup.version`**
:   Backup version

type: keyword


## api_service [_api_service]

API service information

**`netapp.cluster.node.service_processor.api_service.enabled`**
:   API service enabled

type: boolean


**`netapp.cluster.node.service_processor.api_service.limit_access`**
:   API service limit access

type: boolean


**`netapp.cluster.node.service_processor.api_service.port`**
:   API service port

type: integer


## web_service [_web_service]

Web service information

**`netapp.cluster.node.service_processor.web_service.enabled`**
:   Web service enabled

type: boolean


**`netapp.cluster.node.service_processor.web_service.limit_access`**
:   Web service limit access

type: boolean


## nvram [_nvram]

NVRAM information

**`netapp.cluster.node.nvram.id`**
:   NVRAM ID

type: keyword


**`netapp.cluster.node.nvram.battery_state`**
:   NVRAM battery state

type: keyword


## external_cache [_external_cache]

External cache information

**`netapp.cluster.node.external_cache.is_enabled`**
:   External cache enabled

type: boolean


**`netapp.cluster.node.external_cache.is_hya_enabled`**
:   HYA enabled

type: boolean


**`netapp.cluster.node.external_cache.is_rewarm_enabled`**
:   Rewarm enabled

type: boolean


**`netapp.cluster.node.external_cache.pcs_size`**
:   PCS size

type: long


## hw_assist [_hw_assist]

Hardware assist information

## status [_status]

Hardware assist status

**`netapp.cluster.node.hw_assist.status.enabled`**
:   Hardware assist enabled

type: boolean


## local [_local]

Local hardware assist status

**`netapp.cluster.node.hw_assist.status.local.state`**
:   Local state

type: keyword


**`netapp.cluster.node.hw_assist.status.local.ip`**
:   Local IP

type: keyword


**`netapp.cluster.node.hw_assist.status.local.port`**
:   Local port

type: integer


## partner [_partner]

Partner hardware assist status

**`netapp.cluster.node.hw_assist.status.partner.state`**
:   Partner state

type: keyword


**`netapp.cluster.node.hw_assist.status.partner.ip`**
:   Partner IP

type: keyword


**`netapp.cluster.node.hw_assist.status.partner.port`**
:   Partner port

type: integer


**`netapp.cluster.node.anti_ransomware_version`**
:   Anti-ransomware version

type: keyword


**`netapp.cluster.node.metric`**
:   Node metric information

type: object


**`netapp.cluster.node.statistics`**
:   Node statistics

type: object


## sensor [_sensor]

Node sensor information

**`netapp.cluster.sensor.node`**
:   Node object associated with the sensor

type: object


**`netapp.cluster.sensor.index`**
:   Sensor index

type: integer


**`netapp.cluster.sensor.name`**
:   Sensor name

type: keyword


**`netapp.cluster.sensor.type`**
:   Sensor type

type: keyword


**`netapp.cluster.sensor.value`**
:   Sensor value

type: double


**`netapp.cluster.sensor.value_units`**
:   Units for the sensor value

type: keyword


**`netapp.cluster.sensor.threshold_state`**
:   Sensor threshold state

type: keyword


**`netapp.cluster.sensor.critical_low_threshold`**
:   Critical low threshold for the sensor

type: double


**`netapp.cluster.sensor.warning_low_threshold`**
:   Warning low threshold for the sensor

type: double


**`netapp.cluster.sensor.warning_high_threshold`**
:   Warning high threshold for the sensor

type: double


**`netapp.cluster.sensor.critical_high_threshold`**
:   Critical high threshold for the sensor

type: double


## counters [_counters]

Counter table events and values

**`netapp.cluster.counters.counter_table`**
:   Name of the counter table

type: keyword


**`netapp.cluster.counters.node_name`**
:   Name of the node associated with the counter

type: keyword


**`netapp.cluster.counters.id`**
:   Unique identifier for the counter event

type: keyword


**`netapp.cluster.counters.properties`**
:   Additional properties for the counter event (as JSON string)

type: keyword


**`netapp.cluster.counters.counter_name`**
:   Name of the counter

type: keyword


**`netapp.cluster.counters.counter_value`**
:   Value of the counter

type: integer


## interfaces [_interfaces]

interfaces

**`netapp.interfaces.if_name`**
:   Interface name

type: keyword


**`netapp.interfaces.metric_name`**
:   Metric name

type: keyword


**`netapp.interfaces.value`**
:   Metric value

type: long


## protocols [_protocols]

protocols

## iscsi_service [_iscsi_service]

ISCSI service information

## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.iscsi_service.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.iscsi_service.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.iscsi_service.enabled`**
:   Whether the ISCSI service is enabled

type: boolean


## target [_target]

ISCSI target information

**`netapp.protocols.iscsi_service.target.iqn`**
:   ISCSI Qualified Name of the target

type: keyword


**`netapp.protocols.iscsi_service.target.alias`**
:   Alias of the ISCSI target

type: keyword


**`netapp.protocols.iscsi_service.metric`**
:   ISCSI service metrics

type: object


**`netapp.protocols.iscsi_service.statistics`**
:   ISCSI service statistics

type: object


## cifs_service [_cifs_service]

CIFS service information

## ad_domain [_ad_domain]

Active Directory domain information

**`netapp.protocols.cifs_service.ad_domain.default_site`**
:   Default site of the AD domain

type: keyword


**`netapp.protocols.cifs_service.ad_domain.fqdn`**
:   FQDN of the AD domain

type: keyword


**`netapp.protocols.cifs_service.ad_domain.organizational_unit`**
:   Organizational unit of the AD domain

type: keyword


**`netapp.protocols.cifs_service.auth_style`**
:   Authentication style

type: keyword


**`netapp.protocols.cifs_service.auth_user_type`**
:   Authentication user type

type: keyword


**`netapp.protocols.cifs_service.authentication_method`**
:   Authentication method

type: keyword


**`netapp.protocols.cifs_service.client_id`**
:   Client ID

type: keyword


**`netapp.protocols.cifs_service.comment`**
:   Comment for the CIFS service

type: keyword


**`netapp.protocols.cifs_service.default_unix_user`**
:   Default UNIX user

type: keyword


**`netapp.protocols.cifs_service.enabled`**
:   Whether the CIFS service is enabled

type: boolean


**`netapp.protocols.cifs_service.group_policy_object_enabled`**
:   Whether Group Policy Object is enabled

type: boolean


**`netapp.protocols.cifs_service.key_vault_uri`**
:   Key Vault URI

type: keyword


**`netapp.protocols.cifs_service.name`**
:   Name of the CIFS service

type: keyword


## netbios [_netbios]

NetBIOS information

**`netapp.protocols.cifs_service.netbios.aliases`**
:   List of NetBIOS aliases

type: keyword


**`netapp.protocols.cifs_service.netbios.enabled`**
:   Whether NetBIOS is enabled

type: boolean


**`netapp.protocols.cifs_service.netbios.wins_servers`**
:   WINS servers (JSON string)

type: keyword


**`netapp.protocols.cifs_service.oauth_host`**
:   OAuth host

type: keyword


**`netapp.protocols.cifs_service.options`**
:   CIFS service options

type: object


**`netapp.protocols.cifs_service.proxy_host`**
:   Proxy host

type: keyword


**`netapp.protocols.cifs_service.proxy_port`**
:   Proxy port

type: long


**`netapp.protocols.cifs_service.proxy_type`**
:   Proxy type

type: keyword


**`netapp.protocols.cifs_service.proxy_username`**
:   Proxy username

type: keyword


**`netapp.protocols.cifs_service.security`**
:   CIFS security settings

type: object


## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.cifs_service.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.cifs_service.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.cifs_service.tenant_id`**
:   Tenant ID

type: keyword


**`netapp.protocols.cifs_service.timeout`**
:   Timeout value

type: long


**`netapp.protocols.cifs_service.verify_host`**
:   Whether to verify the host

type: boolean


**`netapp.protocols.cifs_service.workgroup`**
:   Workgroup name

type: keyword


**`netapp.protocols.cifs_service.metric`**
:   CIFS service metrics

type: object


**`netapp.protocols.cifs_service.statistics`**
:   CIFS service statistics

type: object


## cifs_share [_cifs_share]

CIFS share information

**`netapp.protocols.cifs_share.access_based_enumeration`**
:   Whether access-based enumeration is enabled

type: boolean


**`netapp.protocols.cifs_share.acls`**
:   ACLs for the CIFS share (JSON string)

type: keyword


**`netapp.protocols.cifs_share.allow_unencrypted_access`**
:   Whether unencrypted access is allowed

type: boolean


**`netapp.protocols.cifs_share.attribute_cache`**
:   Whether attribute caching is enabled

type: boolean


**`netapp.protocols.cifs_share.browsable`**
:   Whether the share is browsable

type: boolean


**`netapp.protocols.cifs_share.change_notify`**
:   Whether change notify is enabled

type: boolean


**`netapp.protocols.cifs_share.comment`**
:   Comment for the CIFS share

type: keyword


**`netapp.protocols.cifs_share.continuously_available`**
:   Whether the share is continuously available

type: boolean


**`netapp.protocols.cifs_share.dir_umask`**
:   Directory umask for the share

type: keyword


**`netapp.protocols.cifs_share.encryption`**
:   Whether encryption is enabled for the share

type: boolean


**`netapp.protocols.cifs_share.file_umask`**
:   File umask for the share

type: keyword


**`netapp.protocols.cifs_share.force_group_for_create`**
:   Group to force for file creation

type: keyword


**`netapp.protocols.cifs_share.home_directory`**
:   Whether the share is a home directory

type: boolean


**`netapp.protocols.cifs_share.max_connections_per_share`**
:   Maximum connections allowed per share

type: long


**`netapp.protocols.cifs_share.name`**
:   Name of the CIFS share

type: keyword


**`netapp.protocols.cifs_share.namespace_caching`**
:   Whether namespace caching is enabled

type: boolean


**`netapp.protocols.cifs_share.no_strict_security`**
:   Whether strict security is disabled

type: boolean


**`netapp.protocols.cifs_share.offline_files`**
:   Offline files setting

type: keyword


**`netapp.protocols.cifs_share.oplocks`**
:   Whether opportunistic locking is enabled

type: boolean


**`netapp.protocols.cifs_share.path`**
:   Path of the CIFS share

type: keyword


**`netapp.protocols.cifs_share.show_previous_versions`**
:   Whether previous versions are shown

type: boolean


**`netapp.protocols.cifs_share.show_snapshot`**
:   Whether snapshots are shown

type: boolean


## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.cifs_share.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.cifs_share.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.cifs_share.unix_symlink`**
:   Whether UNIX symlinks are enabled

type: boolean


## volume [_volume]

Volume information

**`netapp.protocols.cifs_share.volume.name`**
:   Name of the volume

type: keyword


**`netapp.protocols.cifs_share.volume.uuid`**
:   UUID of the volume

type: keyword


**`netapp.protocols.cifs_share.vscan_profile`**
:   Vscan profile for the share

type: keyword


## fc_interface [_fc_interface]

FC interface information

**`netapp.protocols.fc_interface.comment`**
:   Comment for the FC interface

type: keyword


**`netapp.protocols.fc_interface.data_protocol`**
:   Data protocol used by the FC interface

type: keyword


**`netapp.protocols.fc_interface.enabled`**
:   Whether the FC interface is enabled

type: boolean


## location [_location]

Location information for the FC interface

## home_node [_home_node]

Home node information

**`netapp.protocols.fc_interface.location.home_node.name`**
:   Name of the home node

type: keyword


**`netapp.protocols.fc_interface.location.home_node.uuid`**
:   UUID of the home node

type: keyword


## home_port [_home_port]

Home port information

**`netapp.protocols.fc_interface.location.home_port.name`**
:   Name of the home port

type: keyword


**`netapp.protocols.fc_interface.location.home_port.uuid`**
:   UUID of the home port

type: keyword


**`netapp.protocols.fc_interface.location.is_home`**
:   Whether this is the home location

type: boolean


## node [_node]

Node information

**`netapp.protocols.fc_interface.location.node.name`**
:   Name of the node

type: keyword


**`netapp.protocols.fc_interface.location.node.uuid`**
:   UUID of the node

type: keyword


## port [_port]

Port information

**`netapp.protocols.fc_interface.location.port.name`**
:   Name of the port

type: keyword


**`netapp.protocols.fc_interface.location.port.uuid`**
:   UUID of the port

type: keyword


**`netapp.protocols.fc_interface.metric`**
:   FC interface metrics

type: object


**`netapp.protocols.fc_interface.name`**
:   Name of the FC interface

type: keyword


**`netapp.protocols.fc_interface.port_address`**
:   Port address of the FC interface

type: keyword


**`netapp.protocols.fc_interface.state`**
:   State of the FC interface

type: keyword


**`netapp.protocols.fc_interface.statistics`**
:   FC interface statistics

type: object


## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.fc_interface.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.fc_interface.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.fc_interface.uuid`**
:   UUID of the FC interface

type: keyword


**`netapp.protocols.fc_interface.wwnn`**
:   World Wide Node Name (WWNN) of the FC interface

type: keyword


**`netapp.protocols.fc_interface.wwpn`**
:   World Wide Port Name (WWPN) of the FC interface

type: keyword


## fc_port [_fc_port]

FC port information

## node [_node]

Node information

**`netapp.protocols.fc_port.node.name`**
:   Name of the node

type: keyword


**`netapp.protocols.fc_port.node.uuid`**
:   UUID of the node

type: keyword


**`netapp.protocols.fc_port.name`**
:   Name of the FC port

type: keyword


**`netapp.protocols.fc_port.uuid`**
:   UUID of the FC port

type: keyword


**`netapp.protocols.fc_port.description`**
:   Description of the FC port

type: keyword


**`netapp.protocols.fc_port.enabled`**
:   Whether the FC port is enabled

type: boolean


## fabric [_fabric]

Fabric information

**`netapp.protocols.fc_port.fabric.connected`**
:   Whether the port is connected to the fabric

type: boolean


**`netapp.protocols.fc_port.fabric.connected_speed`**
:   Connected speed of the fabric

type: keyword


**`netapp.protocols.fc_port.fabric.port_address`**
:   Port address in the fabric

type: keyword


**`netapp.protocols.fc_port.fabric.switch_port`**
:   Switch port in the fabric

type: keyword


**`netapp.protocols.fc_port.physical_protocol`**
:   Physical protocol of the FC port

type: keyword


## speed [_speed]

Speed information

**`netapp.protocols.fc_port.speed.maximum`**
:   Maximum speed of the port

type: keyword


**`netapp.protocols.fc_port.speed.configured`**
:   Configured speed of the port

type: keyword


**`netapp.protocols.fc_port.state`**
:   State of the FC port

type: keyword


**`netapp.protocols.fc_port.supported_protocols`**
:   Supported protocols (JSON string)

type: keyword


## transceiver [_transceiver]

Transceiver information

**`netapp.protocols.fc_port.transceiver.form_factor`**
:   Form factor of the transceiver

type: keyword


**`netapp.protocols.fc_port.transceiver.manufacturer`**
:   Manufacturer of the transceiver

type: keyword


**`netapp.protocols.fc_port.transceiver.capabilities`**
:   Capabilities of the transceiver (comma-separated string)

type: keyword


**`netapp.protocols.fc_port.transceiver.part_number`**
:   Part number of the transceiver

type: keyword


**`netapp.protocols.fc_port.wwnn`**
:   World Wide Node Name (WWNN) of the FC port

type: keyword


**`netapp.protocols.fc_port.wwpn`**
:   World Wide Port Name (WWPN) of the FC port

type: keyword


**`netapp.protocols.fc_port.metric`**
:   FC port metrics

type: object


**`netapp.protocols.fc_port.statistics`**
:   FC port statistics

type: object


## fcp_service [_fcp_service]

FCP service information

**`netapp.protocols.fcp_service.enabled`**
:   Whether the FCP service is enabled

type: boolean


**`netapp.protocols.fcp_service.metric`**
:   FCP service metrics

type: object


**`netapp.protocols.fcp_service.statistics`**
:   FCP service statistics

type: object


## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.fcp_service.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.fcp_service.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.fcp_service.target_name`**
:   Name of the FCP target

type: keyword


## nfs_service [_nfs_service]

NFS service information

## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.nfs_service.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.nfs_service.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.nfs_service.enabled`**
:   Whether the NFS service is enabled

type: boolean


**`netapp.protocols.nfs_service.state`**
:   State of the NFS service

type: keyword


## transport [_transport]

NFS transport protocol settings

**`netapp.protocols.nfs_service.transport.udp_enabled`**
:   Whether UDP transport is enabled

type: boolean


**`netapp.protocols.nfs_service.transport.tcp_enabled`**
:   Whether TCP transport is enabled

type: boolean


**`netapp.protocols.nfs_service.transport.rdma_enabled`**
:   Whether RDMA transport is enabled

type: boolean


## protocol [_protocol]

NFS protocol settings

**`netapp.protocols.nfs_service.protocol.v3_enabled`**
:   Whether NFSv3 is enabled

type: boolean


**`netapp.protocols.nfs_service.protocol.v3_64bit_identifiers_enabled`**
:   Whether 64-bit identifiers are enabled for NFSv3

type: boolean


**`netapp.protocols.nfs_service.protocol.v4_id_domain`**
:   ID domain for NFSv4

type: keyword


**`netapp.protocols.nfs_service.protocol.v4_64bit_identifiers_enabled`**
:   Whether 64-bit identifiers are enabled for NFSv4

type: boolean


**`netapp.protocols.nfs_service.protocol.v40_enabled`**
:   Whether NFSv4.0 is enabled

type: boolean


**`netapp.protocols.nfs_service.protocol.v41_enabled`**
:   Whether NFSv4.1 is enabled

type: boolean


**`netapp.protocols.nfs_service.protocol.v4_grace_seconds`**
:   Grace period in seconds for NFSv4

type: long


## v40_features [_v40_features]

NFSv4.0 feature settings

**`netapp.protocols.nfs_service.protocol.v40_features.acl_enabled`**
:   Whether ACL is enabled for NFSv4.0

type: boolean


**`netapp.protocols.nfs_service.protocol.v40_features.read_delegation_enabled`**
:   Whether read delegation is enabled for NFSv4.0

type: boolean


**`netapp.protocols.nfs_service.protocol.v40_features.write_delegation_enabled`**
:   Whether write delegation is enabled for NFSv4.0

type: boolean


**`netapp.protocols.nfs_service.protocol.v40_features.acl_preserve`**
:   Whether ACL preservation is enabled for NFSv4.0

type: boolean


## v41_features [_v41_features]

NFSv4.1 feature settings

**`netapp.protocols.nfs_service.protocol.v41_features.acl_enabled`**
:   Whether ACL is enabled for NFSv4.1

type: boolean


**`netapp.protocols.nfs_service.protocol.v41_features.read_delegation_enabled`**
:   Whether read delegation is enabled for NFSv4.1

type: boolean


**`netapp.protocols.nfs_service.protocol.v41_features.write_delegation_enabled`**
:   Whether write delegation is enabled for NFSv4.1

type: boolean


**`netapp.protocols.nfs_service.protocol.v41_features.pnfs_enabled`**
:   Whether pNFS is enabled for NFSv4.1

type: boolean


## v3_features [_v3_features]

NFSv3 feature settings

**`netapp.protocols.nfs_service.protocol.v3_features.mount_root_only`**
:   Whether only root can mount NFSv3 exports

type: boolean


**`netapp.protocols.nfs_service.vstorage_enabled`**
:   Whether vStorage is enabled

type: boolean


**`netapp.protocols.nfs_service.rquota_enabled`**
:   Whether rquota is enabled

type: boolean


**`netapp.protocols.nfs_service.showmount_enabled`**
:   Whether showmount is enabled

type: boolean


**`netapp.protocols.nfs_service.auth_sys_extended_groups_enabled`**
:   Whether AUTH_SYS extended groups are enabled

type: boolean


**`netapp.protocols.nfs_service.extended_groups_limit`**
:   Limit for extended groups

type: long


## credential_cache [_credential_cache]

Credential cache settings

**`netapp.protocols.nfs_service.credential_cache.positive_ttl`**
:   Positive TTL for credential cache

type: long


## qtree [_qtree]

Qtree export settings

**`netapp.protocols.nfs_service.qtree.export_enabled`**
:   Whether qtree export is enabled

type: boolean


**`netapp.protocols.nfs_service.qtree.validate_export`**
:   Whether qtree export validation is enabled

type: boolean


## access_cache_config [_access_cache_config]

Access cache configuration

**`netapp.protocols.nfs_service.access_cache_config.ttl_positive`**
:   Positive TTL for access cache

type: long


**`netapp.protocols.nfs_service.access_cache_config.ttl_negative`**
:   Negative TTL for access cache

type: long


**`netapp.protocols.nfs_service.access_cache_config.harvest_timeout`**
:   Harvest timeout for access cache

type: long


**`netapp.protocols.nfs_service.access_cache_config.is_dns_ttl_enabled`**
:   Whether DNS TTL is enabled for access cache

type: boolean


**`netapp.protocols.nfs_service.file_session_io_grouping_count`**
:   File session IO grouping count

type: long


**`netapp.protocols.nfs_service.file_session_io_grouping_duration`**
:   File session IO grouping duration

type: long


## exports [_exports]

NFS export settings

**`netapp.protocols.nfs_service.exports.name_service_lookup_protocol`**
:   Name service lookup protocol for exports

type: keyword


## security [_security]

NFS security settings

**`netapp.protocols.nfs_service.security.permitted_encryption_types`**
:   Permitted encryption types (comma-separated string)

type: keyword


## windows [_windows]

Windows-specific NFS settings

**`netapp.protocols.nfs_service.windows.v3_ms_dos_client_enabled`**
:   Whether NFSv3 MS-DOS client support is enabled

type: boolean


## metric [_metric]

NFS service metrics

**`netapp.protocols.nfs_service.metric.v3`**
:   NFSv3 metrics

type: object


**`netapp.protocols.nfs_service.metric.v4`**
:   NFSv4 metrics

type: object


**`netapp.protocols.nfs_service.metric.v41`**
:   NFSv4.1 metrics

type: object


## statistics [_statistics]

NFS service statistics

**`netapp.protocols.nfs_service.statistics.v3`**
:   NFSv3 statistics

type: object


**`netapp.protocols.nfs_service.statistics.v4`**
:   NFSv4 statistics

type: object


**`netapp.protocols.nfs_service.statistics.v41`**
:   NFSv4.1 statistics

type: object


## nfs_export_policy [_nfs_export_policy]

NFS export policy information

## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.nfs_export_policy.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.nfs_export_policy.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.nfs_export_policy.id`**
:   ID of the NFS export policy

type: long


**`netapp.protocols.nfs_export_policy.name`**
:   Name of the NFS export policy

type: keyword


## ip_interface [_ip_interface]

IP interface information

**`netapp.protocols.ip_interface.ddns_enabled`**
:   Whether DDNS is enabled

type: boolean


**`netapp.protocols.ip_interface.dns_zone`**
:   DNS zone of the interface

type: keyword


**`netapp.protocols.ip_interface.enabled`**
:   Whether the IP interface is enabled

type: boolean


## ip [_ip]

IP address information

**`netapp.protocols.ip_interface.ip.address`**
:   IP address of the interface

type: ip


**`netapp.protocols.ip_interface.ip.family`**
:   IP address family (e.g., IPv4, IPv6)

type: keyword


**`netapp.protocols.ip_interface.ip.netmask`**
:   Netmask of the IP address

type: keyword


## ipspace [_ipspace]

IPspace information

**`netapp.protocols.ip_interface.ipspace.name`**
:   Name of the IPspace

type: keyword


**`netapp.protocols.ip_interface.ipspace.uuid`**
:   UUID of the IPspace

type: keyword


## location [_location]

Location information for the IP interface

**`netapp.protocols.ip_interface.location.auto_revert`**
:   Whether auto-revert is enabled

type: boolean


**`netapp.protocols.ip_interface.location.failover`**
:   failover policy for a logical interface (LIF) — specifically, how it behaves when a failover event occurs  (e.g., node or port failure). Sample values include "disabled", "local_only", "home_port_only", and "home_node_only".

type: keyword


## home_node [_home_node]

Home node information

**`netapp.protocols.ip_interface.location.home_node.name`**
:   Name of the home node

type: keyword


**`netapp.protocols.ip_interface.location.home_node.uuid`**
:   UUID of the home node

type: keyword


## home_port [_home_port]

Home port information

**`netapp.protocols.ip_interface.location.home_port.name`**
:   Name of the home port

type: keyword


**`netapp.protocols.ip_interface.location.home_port.uuid`**
:   UUID of the home port

type: keyword


**`netapp.protocols.ip_interface.location.is_home`**
:   Whether this is the home location

type: boolean


## node [_node]

Node information

**`netapp.protocols.ip_interface.location.node.name`**
:   Name of the node

type: keyword


**`netapp.protocols.ip_interface.location.node.uuid`**
:   UUID of the node

type: keyword


## port [_port]

Port information

**`netapp.protocols.ip_interface.location.port.name`**
:   Name of the port

type: keyword


**`netapp.protocols.ip_interface.location.port.uuid`**
:   UUID of the port

type: keyword


**`netapp.protocols.ip_interface.metric`**
:   IP interface metrics

type: object


**`netapp.protocols.ip_interface.name`**
:   Name of the IP interface

type: keyword


**`netapp.protocols.ip_interface.probe_port`**
:   Probe port of the IP interface

type: long


**`netapp.protocols.ip_interface.rdma_protocols`**
:   RDMA protocols (JSON string)

type: keyword


**`netapp.protocols.ip_interface.scope`**
:   Scope of the IP interface

type: keyword


## service_policy [_service_policy]

Service policy information

**`netapp.protocols.ip_interface.service_policy.name`**
:   Name of the service policy

type: keyword


**`netapp.protocols.ip_interface.service_policy.uuid`**
:   UUID of the service policy

type: keyword


**`netapp.protocols.ip_interface.services`**
:   Services (JSON string)

type: keyword


**`netapp.protocols.ip_interface.state`**
:   State of the IP interface

type: keyword


**`netapp.protocols.ip_interface.statistics`**
:   IP interface statistics

type: object


## subnet [_subnet]

Subnet information

**`netapp.protocols.ip_interface.subnet.name`**
:   Name of the subnet

type: keyword


**`netapp.protocols.ip_interface.subnet.uuid`**
:   UUID of the subnet

type: keyword


## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.ip_interface.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.ip_interface.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.ip_interface.uuid`**
:   UUID of the IP interface

type: keyword


**`netapp.protocols.ip_interface.vip`**
:   Whether the interface is a virtual IP (VIP)

type: boolean


## iscsi_session [_iscsi_session]

ISCSI session information

**`netapp.protocols.iscsi_session.igroups`**
:   ISCSI initiator groups (JSON string)

type: keyword


## initiator [_initiator]

ISCSI initiator information

**`netapp.protocols.iscsi_session.initiator.alias`**
:   Alias of the ISCSI initiator

type: keyword


**`netapp.protocols.iscsi_session.initiator.comment`**
:   Comment for the ISCSI initiator

type: keyword


**`netapp.protocols.iscsi_session.initiator.name`**
:   Name of the ISCSI initiator

type: keyword


**`netapp.protocols.iscsi_session.isid`**
:   ISCSI session identifier (ISID)

type: keyword


## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.iscsi_session.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.iscsi_session.svm.uuid`**
:   UUID of the SVM

type: keyword


**`netapp.protocols.iscsi_session.target_portal_group`**
:   Target portal group

type: keyword


**`netapp.protocols.iscsi_session.target_portal_group_tag`**
:   Target portal group tag

type: long


**`netapp.protocols.iscsi_session.tsih`**
:   Target session identifier handle (TSIH)

type: long


## connection [_connection]

ISCSI connection in the session

**`netapp.protocols.iscsi_session.connection.authentication_type`**
:   Authentication type used for the connection

type: keyword


**`netapp.protocols.iscsi_session.connection.cid`**
:   Connection identifier (CID)

type: long


## initiator_address [_initiator_address]

Initiator address information

**`netapp.protocols.iscsi_session.connection.initiator_address.address`**
:   IP address of the initiator

type: ip


**`netapp.protocols.iscsi_session.connection.initiator_address.port`**
:   Port of the initiator

type: long


## interface [_interface]

Interface information

## ip [_ip]

IP information of the interface

**`netapp.protocols.iscsi_session.connection.interface.ip.address`**
:   IP address of the interface

type: ip


**`netapp.protocols.iscsi_session.connection.interface.ip.port`**
:   Port of the interface

type: long


**`netapp.protocols.iscsi_session.connection.interface.name`**
:   Name of the interface

type: keyword


**`netapp.protocols.iscsi_session.connection.interface.uuid`**
:   UUID of the interface

type: keyword


## igroup [_igroup]

ISCSI initiator group (igroup) information

**`netapp.protocols.igroup.comment`**
:   Comment for the igroup

type: keyword


**`netapp.protocols.igroup.connectivity_tracking`**
:   Connectivity tracking information for the igroup

type: object


**`netapp.protocols.igroup.delete_on_unmap`**
:   Whether to delete the igroup on unmap

type: boolean


**`netapp.protocols.igroup.lun_maps`**
:   LUN maps for the igroup (JSON string)

type: keyword


**`netapp.protocols.igroup.name`**
:   Name of the igroup

type: keyword


**`netapp.protocols.igroup.os_type`**
:   Operating system type for the igroup

type: keyword


**`netapp.protocols.igroup.portset`**
:   Portset information for the igroup (JSON string)

type: keyword


**`netapp.protocols.igroup.protocol`**
:   Protocol used by the igroup

type: keyword


**`netapp.protocols.igroup.replication`**
:   Replication information for the igroup

type: object


**`netapp.protocols.igroup.supports_igroups`**
:   Whether igroups are supported

type: boolean


## svm [_svm]

Storage Virtual Machine (SVM) information

**`netapp.protocols.igroup.svm.name`**
:   Name of the SVM

type: keyword


**`netapp.protocols.igroup.svm.uuid`**
:   UUID of the SVM

type: keyword


## target [_target]

Target information for the igroup

**`netapp.protocols.igroup.target.firmware_revision`**
:   Firmware revision of the target

type: keyword


**`netapp.protocols.igroup.target.product_id`**
:   Product ID of the target

type: keyword


**`netapp.protocols.igroup.target.vendor_id`**
:   Vendor ID of the target

type: keyword


**`netapp.protocols.igroup.uuid`**
:   UUID of the igroup

type: keyword


## initiator [_initiator]

ISCSI initiator information

**`netapp.protocols.igroup.initiator.comment`**
:   Comment for the ISCSI initiator

type: keyword


**`netapp.protocols.igroup.initiator.connection_state`**
:   Connection state of the initiator

type: keyword


## igroup [_igroup]

Igroup information for the initiator

**`netapp.protocols.igroup.initiator.igroup.comment`**
:   Comment for the igroup

type: keyword


**`netapp.protocols.igroup.initiator.igroup.igroups`**
:   Placeholder for igroups (always null to avoid recursion)

type: object


**`netapp.protocols.igroup.initiator.igroup.name`**
:   Name of the igroup

type: keyword


**`netapp.protocols.igroup.initiator.igroup.uuid`**
:   UUID of the igroup

type: keyword


**`netapp.protocols.igroup.initiator.name`**
:   Name of the ISCSI initiator

type: keyword


**`netapp.protocols.igroup.initiator.proximity`**
:   Proximity information for the initiator

type: object


## storage [_storage]

storage

## snapmirror [_snapmirror]

Example group

**`netapp.storage.snapmirror.backoff_level`**
:   The backoff level of the SnapMirror relationship.

type: keyword


## consistency_group_failover [_consistency_group_failover]

Consistency group failover details.

## error [_error]

Error details for consistency group failover.

**`netapp.storage.snapmirror.consistency_group_failover.error.code`**
:   Error code for the failover.

type: keyword


**`netapp.storage.snapmirror.consistency_group_failover.error.message`**
:   Error message for the failover.

type: keyword


**`netapp.storage.snapmirror.consistency_group_failover.error.state`**
:   Error state for the failover.

type: keyword


**`netapp.storage.snapmirror.consistency_group_failover.state`**
:   State of the consistency group failover.

type: keyword


## status [_status]

Status details for consistency group failover.

**`netapp.storage.snapmirror.consistency_group_failover.status.code`**
:   Status code for the failover.

type: keyword


**`netapp.storage.snapmirror.consistency_group_failover.status.message`**
:   Status message for the failover.

type: keyword


**`netapp.storage.snapmirror.consistency_group_failover.status.state`**
:   Status state for the failover.

type: keyword


**`netapp.storage.snapmirror.consistency_group_failover.type`**
:   Type of the consistency group failover.

type: keyword


## destination [_destination]

Destination endpoint details.

## cluster [_cluster]

Cluster information for the destination.

**`netapp.storage.snapmirror.destination.cluster.name`**
:   Name of the cluster.

type: keyword


**`netapp.storage.snapmirror.destination.cluster.uuid`**
:   UUID of the cluster.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the destination.

**`netapp.storage.snapmirror.destination.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.snapmirror.destination.svm.uuid`**
:   UUID of the SVM.

type: keyword


## luns [_luns]

LUNs information for the destination.

**`netapp.storage.snapmirror.destination.luns.name`**
:   Name of the LUN.

type: keyword


**`netapp.storage.snapmirror.destination.luns.uuid`**
:   UUID of the LUN.

type: keyword


**`netapp.storage.snapmirror.destination.path`**
:   Path for the destination endpoint.

type: keyword


**`netapp.storage.snapmirror.destination.consistency_group_volumes`**
:   Consistency group volumes for the destination.

type: keyword


**`netapp.storage.snapmirror.exported_snapshot`**
:   The exported snapshot of the SnapMirror relationship.

type: keyword


**`netapp.storage.snapmirror.group_type`**
:   The group type of the SnapMirror relationship.

type: keyword


**`netapp.storage.snapmirror.healthy`**
:   Indicates if the SnapMirror relationship is healthy.

type: boolean


**`netapp.storage.snapmirror.identity_preservation`**
:   Indicates if identity preservation is enabled.

type: boolean


**`netapp.storage.snapmirror.io_serving_copy`**
:   Indicates if this is the IO serving copy.

type: boolean


**`netapp.storage.snapmirror.lag_time`**
:   Lag time in ISO 8601 format (e.g. "PT25M5S").

type: keyword


**`netapp.storage.snapmirror.last_transfer_network_compression_ratio`**
:   The network compression ratio of the last transfer, e.g "1:1".

type: keyword


**`netapp.storage.snapmirror.last_transfer_type`**
:   The type of the last transfer.

type: keyword


**`netapp.storage.snapmirror.master_bias_activated_site`**
:   The master bias activated site.

type: keyword


## policy [_policy]

The SnapMirror policy.

**`netapp.storage.snapmirror.policy.name`**
:   Name of the SnapMirror policy.

type: keyword


**`netapp.storage.snapmirror.policy.type`**
:   Type of the SnapMirror policy.

type: keyword


**`netapp.storage.snapmirror.policy.uuid`**
:   UUID of the SnapMirror policy.

type: keyword


**`netapp.storage.snapmirror.preferred_site`**
:   The preferred site.

type: keyword


**`netapp.storage.snapmirror.restore`**
:   Indicates if restore is in progress.

type: boolean


## source [_source]

Source endpoint details.

## cluster [_cluster]

Cluster information for the source.

**`netapp.storage.snapmirror.source.cluster.name`**
:   Name of the cluster.

type: keyword


**`netapp.storage.snapmirror.source.cluster.uuid`**
:   UUID of the cluster.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the source.

**`netapp.storage.snapmirror.source.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.snapmirror.source.svm.uuid`**
:   UUID of the SVM.

type: keyword


## luns [_luns]

LUNs information for the source.

**`netapp.storage.snapmirror.source.luns.name`**
:   Name of the LUN.

type: keyword


**`netapp.storage.snapmirror.source.luns.uuid`**
:   UUID of the LUN.

type: keyword


**`netapp.storage.snapmirror.source.path`**
:   Path for the source endpoint.

type: keyword


**`netapp.storage.snapmirror.source.consistency_group_volumes`**
:   Consistency group volumes for the source.

type: keyword


**`netapp.storage.snapmirror.state`**
:   The state of the SnapMirror relationship.

type: keyword


**`netapp.storage.snapmirror.svmdr_volumes`**
:   SVMDR volumes.

type: keyword


**`netapp.storage.snapmirror.throttle`**
:   Throttle information.

type: keyword


**`netapp.storage.snapmirror.total_transfer_bytes`**
:   Total bytes transferred.

type: long


**`netapp.storage.snapmirror.total_transfer_duration`**
:   Duration in ISO 8601 format (e.g. "PT25M5S").

type: keyword


## transfer [_transfer]

Transfer information.

**`netapp.storage.snapmirror.transfer.bytes_transferred`**
:   Number of bytes transferred.

type: long


**`netapp.storage.snapmirror.transfer.end_time`**
:   End time of the transfer.

type: date


**`netapp.storage.snapmirror.transfer.last_updated_time`**
:   Last updated time of the transfer.

type: date


**`netapp.storage.snapmirror.transfer.state`**
:   State of the transfer.

type: keyword


**`netapp.storage.snapmirror.transfer.total_duration`**
:   Total duration of the transfer.

type: keyword


**`netapp.storage.snapmirror.transfer.type`**
:   Type of the transfer.

type: keyword


**`netapp.storage.snapmirror.transfer.uuid`**
:   UUID of the transfer.

type: keyword


## transfer_schedule [_transfer_schedule]

Transfer schedule information.

**`netapp.storage.snapmirror.transfer_schedule.name`**
:   Name of the transfer schedule.

type: keyword


**`netapp.storage.snapmirror.transfer_schedule.uuid`**
:   UUID of the transfer schedule.

type: keyword


**`netapp.storage.snapmirror.unhealthy_reason`**
:   Reason for unhealthy state.

type: keyword


**`netapp.storage.snapmirror.uuid`**
:   UUID of the SnapMirror relationship.

type: keyword


## aggregate [_aggregate]

Example group

**`netapp.storage.aggregate.uuid`**
:   UUID of the aggregate.

type: keyword


**`netapp.storage.aggregate.name`**
:   Name of the aggregate.

type: keyword


## node [_node]

Node information for the aggregate.

**`netapp.storage.aggregate.node.name`**
:   Name of the node.

type: keyword


**`netapp.storage.aggregate.node.uuid`**
:   UUID of the node.

type: keyword


## home_node [_home_node]

Home node information for the aggregate.

**`netapp.storage.aggregate.home_node.name`**
:   Name of the home node.

type: keyword


**`netapp.storage.aggregate.home_node.uuid`**
:   UUID of the home node.

type: keyword


## snapshot [_snapshot]

Snapshot information for the aggregate.

**`netapp.storage.aggregate.snapshot.files_total`**
:   Total number of files in the snapshot.

type: long


**`netapp.storage.aggregate.snapshot.files_used`**
:   Number of files used in the snapshot.

type: long


**`netapp.storage.aggregate.snapshot.max_files_available`**
:   Maximum number of files available in the snapshot.

type: long


**`netapp.storage.aggregate.snapshot.max_files_used`**
:   Maximum number of files used in the snapshot.

type: long


## space [_space]

Space information for the aggregate.

## block_storage [_block_storage]

Block storage space information for the aggregate.

**`netapp.storage.aggregate.space.block_storage.size`**
:   Total size of block storage in bytes.

type: long


**`netapp.storage.aggregate.space.block_storage.available`**
:   Available block storage space in bytes.

type: long


**`netapp.storage.aggregate.space.block_storage.used`**
:   Used block storage space in bytes.

type: long


**`netapp.storage.aggregate.space.block_storage.used_percent`**
:   Percentage of used block storage space.

type: float


**`netapp.storage.aggregate.space.block_storage.full_threshold_percent`**
:   Full threshold percentage for block storage.

type: float


**`netapp.storage.aggregate.space.block_storage.physical_used`**
:   Physical used space in block storage in bytes.

type: long


**`netapp.storage.aggregate.space.block_storage.physical_used_percent`**
:   Percentage of physical used space in block storage.

type: float


**`netapp.storage.aggregate.space.block_storage.data_compacted_count`**
:   Number of data compaction operations performed.

type: long


**`netapp.storage.aggregate.space.block_storage.data_compaction_space_saved`**
:   Space saved by data compaction in bytes.

type: long


**`netapp.storage.aggregate.space.block_storage.data_compaction_space_saved_percent`**
:   Percentage of space saved by data compaction.

type: float


**`netapp.storage.aggregate.space.block_storage.volume_deduplication_shared_count`**
:   Number of deduplication shared blocks in volumes.

type: long


**`netapp.storage.aggregate.space.block_storage.volume_deduplication_space_saved`**
:   Space saved by volume deduplication in bytes.

type: long


**`netapp.storage.aggregate.space.block_storage.volume_deduplication_space_saved_percent`**
:   Percentage of space saved by volume deduplication.

type: float


## snapshot [_snapshot]

Snapshot space information for the aggregate.

**`netapp.storage.aggregate.space.snapshot.size`**
:   Total size of snapshots in bytes.

type: long


**`netapp.storage.aggregate.space.snapshot.used`**
:   Used snapshot space in bytes.

type: long


**`netapp.storage.aggregate.space.snapshot.available`**
:   Available snapshot space in bytes.

type: long


**`netapp.storage.aggregate.space.snapshot.used_percent`**
:   Percentage of used snapshot space.

type: float


**`netapp.storage.aggregate.space.snapshot.reserve_percent`**
:   Percentage of reserved snapshot space.

type: float


**`netapp.storage.aggregate.space.snapshot.total`**
:   Total snapshot space in bytes.

type: long


## cloud_storage [_cloud_storage]

Cloud storage space information for the aggregate.

**`netapp.storage.aggregate.space.cloud_storage.used`**
:   Used cloud storage space in bytes.

type: long


## efficiency [_efficiency]

Efficiency information for the aggregate.

**`netapp.storage.aggregate.space.efficiency.savings`**
:   Space savings in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency.ratio`**
:   Efficiency ratio.

type: float


**`netapp.storage.aggregate.space.efficiency.logical_used`**
:   Logical used space in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency.cross_volume_background_dedupe`**
:   Cross-volume background deduplication value.

type: boolean


**`netapp.storage.aggregate.space.efficiency.cross_volume_inline_dedupe`**
:   Indicates whether cross-volume inline deduplication is enabled on the aggregate.

type: boolean


**`netapp.storage.aggregate.space.efficiency.cross_volume_dedupe_savings`**
:   Cross-volume deduplication savings in bytes.

type: boolean


**`netapp.storage.aggregate.space.efficiency.auto_adaptive_compression_savings`**
:   Auto-adaptive compression savings in bytes.

type: boolean


**`netapp.storage.aggregate.space.efficiency.enable_workload_informed_tsse`**
:   Indicates if workload-informed TSSE is enabled.

type: boolean


**`netapp.storage.aggregate.space.efficiency.wise_tsse_min_used_capacity_pct`**
:   Minimum used capacity percent for WISE TSSE.

type: float


**`netapp.storage.aggregate.space.efficiency.wise_tsse_space_savings_pct`**
:   Percentage of space savings.

type: float


## efficiency_without_snapshots [_efficiency_without_snapshots]

Efficiency information without snapshots for the aggregate.

**`netapp.storage.aggregate.space.efficiency_without_snapshots.logical_used`**
:   Logical used space in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency_without_snapshots.physical_used`**
:   Physical used space in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency_without_snapshots.savings`**
:   Space savings in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency_without_snapshots.savings_percent`**
:   Percentage of space savings.

type: float


## efficiency_without_snapshots_flexclones [_efficiency_without_snapshots_flexclones]

Efficiency information without snapshots and FlexClones for the aggregate.

**`netapp.storage.aggregate.space.efficiency_without_snapshots_flexclones.logical_used`**
:   Logical used space in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency_without_snapshots_flexclones.physical_used`**
:   Physical used space in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency_without_snapshots_flexclones.savings`**
:   Space savings in bytes.

type: long


**`netapp.storage.aggregate.space.efficiency_without_snapshots_flexclones.savings_percent`**
:   Percentage of space savings.

type: float


**`netapp.storage.aggregate.state`**
:   State of the aggregate.

type: keyword


**`netapp.storage.aggregate.snaplock_type`**
:   SnapLock type of the aggregate.

type: keyword


**`netapp.storage.aggregate.create_time`**
:   Creation time of the aggregate.

type: date


## data_encryption [_data_encryption]

Data encryption information for the aggregate.

**`netapp.storage.aggregate.data_encryption.savings`**
:   Space savings in bytes.

type: long


**`netapp.storage.aggregate.data_encryption.ratio`**
:   Efficiency ratio.

type: float


**`netapp.storage.aggregate.data_encryption.logical_used`**
:   Logical used space in bytes. Type of encryption.

type: long


## block_storage [_block_storage]

Block storage information for the aggregate.

## primary [_primary]

Primary block storage information.

**`netapp.storage.aggregate.block_storage.primary.disk_count`**
:   Number of disks in the primary block storage.

type: long


**`netapp.storage.aggregate.block_storage.primary.disk_class`**
:   Disk class of the primary block storage.

type: keyword


**`netapp.storage.aggregate.block_storage.primary.raid_type`**
:   RAID type of the primary block storage.

type: keyword


**`netapp.storage.aggregate.block_storage.primary.raid_size`**
:   RAID size of the primary block storage.

type: long


**`netapp.storage.aggregate.block_storage.primary.checksum_style`**
:   Checksum style of the primary block storage.

type: keyword


**`netapp.storage.aggregate.block_storage.primary.disk_type`**
:   Disk type of the primary block storage.

type: keyword


**`netapp.storage.aggregate.block_storage.type`**
:   Type of block storage.

type: keyword


## cloud_storage [_cloud_storage]

Cloud storage information for the aggregate.

**`netapp.storage.aggregate.cloud_storage.attach_eligible`**
:   Indicates if the aggregate is eligible for cloud storage attachment.

type: boolean


**`netapp.storage.aggregate.cloud_storage.stores`**
:   JSON string representing the cloud storage stores.

type: keyword


## inactive_data_reporting [_inactive_data_reporting]

Inactive data reporting information for the aggregate.

**`netapp.storage.aggregate.inactive_data_reporting.enabled`**
:   Indicates if inactive data reporting is enabled.

type: boolean


**`netapp.storage.aggregate.inactive_data_reporting.start_time`**
:   Start time of inactive data reporting (present only if enabled).

type: date


## inode_attributes [_inode_attributes]

Inode attributes for the aggregate.

**`netapp.storage.aggregate.inode_attributes.files_total`**
:   Total number of files.

type: long


**`netapp.storage.aggregate.inode_attributes.files_used`**
:   Number of files used.

type: long


**`netapp.storage.aggregate.inode_attributes.max_files_available`**
:   Maximum number of files available.

type: long


**`netapp.storage.aggregate.inode_attributes.max_files_possible`**
:   Maximum possible number of files.

type: long


**`netapp.storage.aggregate.inode_attributes.max_files_used`**
:   Maximum number of files used.

type: long


**`netapp.storage.aggregate.inode_attributes.used_percent`**
:   Percentage of used files. Total number of files.

type: float


**`netapp.storage.aggregate.volume_count`**
:   Number of volumes in the aggregate.

type: long


## metrics [_metrics]

Metrics for the aggregate.

**`netapp.storage.aggregate.metrics.timestamp`**
:   Timestamp of the metric sample.

type: date


**`netapp.storage.aggregate.metrics.duration`**
:   Duration in ISO 8601 format (e.g. "PT25M5S").

type: keyword


**`netapp.storage.aggregate.metrics.status`**
:   Status of the metric sample.

type: keyword


## latency [_latency]

IO latency metrics.

**`netapp.storage.aggregate.metrics.latency.read`**
:   Read IO latency in milliseconds.

type: float


**`netapp.storage.aggregate.metrics.latency.write`**
:   Write IO latency in milliseconds.

type: float


**`netapp.storage.aggregate.metrics.latency.other`**
:   Other IO latency in milliseconds.

type: float


**`netapp.storage.aggregate.metrics.latency.total`**
:   Total IO latency in milliseconds.

type: float


## iops [_iops]

Input/output operations per second metrics.

**`netapp.storage.aggregate.metrics.iops.read`**
:   Read IOPS.

type: float


**`netapp.storage.aggregate.metrics.iops.write`**
:   Write IOPS.

type: float


**`netapp.storage.aggregate.metrics.iops.other`**
:   Other IOPS.

type: float


**`netapp.storage.aggregate.metrics.iops.total`**
:   Total IOPS.

type: float


## throughput [_throughput]

Throughput metrics in bytes per second.

**`netapp.storage.aggregate.metrics.throughput.read`**
:   Read throughput in bytes per second.

type: float


**`netapp.storage.aggregate.metrics.throughput.write`**
:   Write throughput in bytes per second.

type: float


**`netapp.storage.aggregate.metrics.throughput.other`**
:   Other throughput in bytes per second.

type: float


**`netapp.storage.aggregate.metrics.throughput.total`**
:   Total throughput in bytes per second.

type: float


## statistics [_statistics]

Statistics for the aggregate.

**`netapp.storage.aggregate.statistics.timestamp`**
:   Timestamp of the storage statistics sample.

type: date


**`netapp.storage.aggregate.statistics.status`**
:   Status of the storage statistics sample.

type: keyword


## throughput_raw [_throughput_raw]

Raw throughput metrics.

**`netapp.storage.aggregate.statistics.throughput_raw.read`**
:   Raw read throughput in bytes per second.

type: float


**`netapp.storage.aggregate.statistics.throughput_raw.write`**
:   Raw write throughput in bytes per second.

type: float


**`netapp.storage.aggregate.statistics.throughput_raw.other`**
:   Raw other throughput in bytes per second.

type: float


**`netapp.storage.aggregate.statistics.throughput_raw.total`**
:   Raw total throughput in bytes per second.

type: float


## iops_raw [_iops_raw]

Raw IOPS metrics.

**`netapp.storage.aggregate.statistics.iops_raw.read`**
:   Raw read IO latency.

type: float


**`netapp.storage.aggregate.statistics.iops_raw.write`**
:   Raw write IO latency.

type: float


**`netapp.storage.aggregate.statistics.iops_raw.other`**
:   Raw other IO latency.

type: float


**`netapp.storage.aggregate.statistics.iops_raw.total`**
:   Raw total IO latency.

type: float


## latency_raw [_latency_raw]

Raw latency metrics.

**`netapp.storage.aggregate.statistics.latency_raw.read`**
:   Raw read latency.

type: float


**`netapp.storage.aggregate.statistics.latency_raw.write`**
:   Raw write latency.

type: float


**`netapp.storage.aggregate.statistics.latency_raw.other`**
:   Raw other latency.

type: float


**`netapp.storage.aggregate.statistics.latency_raw.total`**
:   Raw total latency.

type: float


## lun [_lun]

Example group

**`netapp.storage.lun.uuid`**
:   UUID of the LUN.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the LUN.

**`netapp.storage.lun.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.lun.svm.uuid`**
:   UUID of the SVM.

type: keyword


**`netapp.storage.lun.name`**
:   Name of the LUN.

type: keyword


## location [_location]

Location information for the LUN.

**`netapp.storage.lun.location.logical_unit`**
:   Logical unit of the LUN location.

type: keyword


## node [_node]

Node information for the LUN location.

**`netapp.storage.lun.location.node.name`**
:   Name of the node.

type: keyword


**`netapp.storage.lun.location.node.uuid`**
:   UUID of the node.

type: keyword


## volume [_volume]

Volume information for the LUN location.

**`netapp.storage.lun.location.volume.name`**
:   Name of the volume.

type: keyword


**`netapp.storage.lun.location.volume.uuid`**
:   UUID of the volume.

type: keyword


**`netapp.storage.lun.class`**
:   Class of the LUN.

type: keyword


**`netapp.storage.lun.create_time`**
:   Creation time of the LUN.

type: date


**`netapp.storage.lun.enabled`**
:   Indicates if the LUN is enabled.

type: boolean


**`netapp.storage.lun.os_type`**
:   OS type of the LUN.

type: keyword


**`netapp.storage.lun.serial_number`**
:   Serial number of the LUN.

type: keyword


## space [_space]

Space information for the LUN.

**`netapp.storage.lun.space.scsi_thin_provisioning_support_enabled`**
:   Indicates if SCSI thin provisioning support is enabled for the LUN.

type: boolean


**`netapp.storage.lun.space.size`**
:   Total size of the LUN in bytes.

type: long


**`netapp.storage.lun.space.used`**
:   Used space of the LUN in bytes.

type: long


## guarantee [_guarantee]

Guarantee information for the LUN space.

**`netapp.storage.lun.space.guarantee.requested`**
:   Requested space for the LUN guarantee in bytes.

type: boolean


**`netapp.storage.lun.space.guarantee.reserved`**
:   Reserved space for the LUN guarantee in bytes.

type: boolean


## status [_status]

Status information for the LUN.

**`netapp.storage.lun.status.container_state`**
:   State of the LUN container.

type: keyword


**`netapp.storage.lun.status.mapped`**
:   Indicates if the LUN is mapped.

type: boolean


**`netapp.storage.lun.status.read_only`**
:   Indicates if the LUN is read-only.

type: boolean


**`netapp.storage.lun.status.state`**
:   State of the LUN. State of the LUN.

type: keyword


## vvol [_vvol]

VVol information for the LUN.

**`netapp.storage.lun.vvol.is_bound`**
:   Indicates if the VVol is bound.

type: boolean


**`netapp.storage.lun.vvol.bindings`**
:   JSON string representing the VVol bindings (present only if is_bound is true).

type: keyword


## qos_policy [_qos_policy]

Example group

**`netapp.storage.qos_policy.name`**
:   Name of the QoS policy.

type: keyword


**`netapp.storage.qos_policy.object_count`**
:   Number of objects associated with the QoS policy.

type: long


**`netapp.storage.qos_policy.pgid`**
:   PGID of the QoS policy.

type: keyword


**`netapp.storage.qos_policy.policy_class`**
:   Policy class of the QoS policy.

type: keyword


**`netapp.storage.qos_policy.scope`**
:   Scope of the QoS policy.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the QoS policy.

**`netapp.storage.qos_policy.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.qos_policy.svm.uuid`**
:   UUID of the QoS policy.

type: keyword


## adaptive [_adaptive]

Adaptive QoS policy settings.

**`netapp.storage.qos_policy.adaptive.absolute_min_iops`**
:   Absolute minimum IOPS for the adaptive QoS policy.

type: long


**`netapp.storage.qos_policy.adaptive.block_size`**
:   Block size for the adaptive QoS policy.

type: long


**`netapp.storage.qos_policy.adaptive.expected_iops`**
:   Expected IOPS for the adaptive QoS policy.

type: long


**`netapp.storage.qos_policy.adaptive.expected_iops_allocation`**
:   Expected IOPS allocation for the adaptive QoS policy.

type: long


**`netapp.storage.qos_policy.adaptive.peak_iops`**
:   Peak IOPS for the adaptive QoS policy.

type: long


**`netapp.storage.qos_policy.adaptive.peak_iops_allocation`**
:   Peak IOPS allocation for the adaptive QoS policy.

type: long


## fixed [_fixed]

Fixed QoS policy settings.

**`netapp.storage.qos_policy.fixed.capacity_shared`**
:   Indicates if the fixed QoS policy is capacity shared.

type: boolean


**`netapp.storage.qos_policy.fixed.max_throughput_iops`**
:   Maximum throughput in IOPS for the fixed QoS policy.

type: long


**`netapp.storage.qos_policy.fixed.max_throughput_mbps`**
:   Maximum throughput in MBps for the fixed QoS policy.

type: long


**`netapp.storage.qos_policy.fixed.min_throughput_iops`**
:   Minimum throughput in IOPS for the fixed QoS policy.

type: long


**`netapp.storage.qos_policy.fixed.min_throughput_mbps`**
:   Minimum throughput in MBps for the fixed QoS policy.

type: long


## qtree [_qtree]

Qtree information.

## volume [_volume]

Volume information for the qtree.

**`netapp.storage.qtree.volume.name`**
:   Name of the volume.

type: keyword


**`netapp.storage.qtree.volume.uuid`**
:   UUID of the volume.

type: keyword


**`netapp.storage.qtree.id`**
:   ID of the qtree.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the qtree.

**`netapp.storage.qtree.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.qtree.svm.uuid`**
:   UUID of the SVM.

type: keyword


**`netapp.storage.qtree.name`**
:   Name of the qtree.

type: keyword


**`netapp.storage.qtree.security_style`**
:   Security style of the qtree.

type: keyword


**`netapp.storage.qtree.unix_permissions`**
:   UNIX permissions of the qtree.

type: integer


## export_policy [_export_policy]

Export policy information for the qtree.

**`netapp.storage.qtree.export_policy.name`**
:   Name of the export policy.

type: keyword


**`netapp.storage.qtree.export_policy.id`**
:   ID of the export policy.

type: keyword


**`netapp.storage.qtree.path`**
:   Path of the qtree.

type: keyword


**`netapp.storage.qtree.nas_path`**
:   NAS path of the qtree.

type: keyword


**`netapp.storage.qtree.user_id`**
:   User ID for the qtree.

type: keyword


**`netapp.storage.qtree.group_id`**
:   Group ID for the qtree.

type: keyword


## metric [_metric]

Metrics for the qtree.

**`netapp.storage.qtree.metric.duration`**
:   Duration in ISO 8601 format (e.g. "PT25M5S").

type: keyword


## iops [_iops]

IOPS metrics for the qtree.

**`netapp.storage.qtree.metric.iops.read`**
:   Read IOPS.

type: float


**`netapp.storage.qtree.metric.iops.write`**
:   Write IOPS.

type: float


**`netapp.storage.qtree.metric.iops.other`**
:   Other IOPS.

type: float


**`netapp.storage.qtree.metric.iops.total`**
:   Total IOPS.

type: float


## latency [_latency]

Latency metrics for the qtree.

**`netapp.storage.qtree.metric.latency.read`**
:   Read latency in milliseconds.

type: float


**`netapp.storage.qtree.metric.latency.write`**
:   Write latency in milliseconds.

type: float


**`netapp.storage.qtree.metric.latency.other`**
:   Other latency in milliseconds.

type: float


**`netapp.storage.qtree.metric.latency.total`**
:   Total latency in milliseconds.

type: float


## throughput [_throughput]

Throughput metrics for the qtree in bytes per second.

**`netapp.storage.qtree.metric.throughput.read`**
:   Read throughput in bytes per second.

type: float


**`netapp.storage.qtree.metric.throughput.write`**
:   Write throughput in bytes per second.

type: float


**`netapp.storage.qtree.metric.throughput.other`**
:   Other throughput in bytes per second.

type: float


**`netapp.storage.qtree.metric.throughput.total`**
:   Total throughput in bytes per second.

type: float


## qtree [_qtree]

Brief qtree information for the metric.

**`netapp.storage.qtree.metric.qtree.name`**
:   Name of the qtree.

type: keyword


**`netapp.storage.qtree.metric.qtree.id`**
:   ID of the qtree.

type: keyword


**`netapp.storage.qtree.metric.status`**
:   Status of the metric sample.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the metric.

**`netapp.storage.qtree.metric.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.qtree.metric.svm.uuid`**
:   UUID of the SVM.

type: keyword


## volume [_volume]

Volume information for the metric.

**`netapp.storage.qtree.metric.volume.name`**
:   Name of the volume.

type: keyword


**`netapp.storage.qtree.metric.volume.uuid`**
:   UUID of the volume.

type: keyword


## quota_report [_quota_report]

Quota report information.

## svm [_svm]

SVM (Storage Virtual Machine) information for the quota report.

**`netapp.storage.quota_report.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.quota_report.svm.uuid`**
:   UUID of the SVM.

type: keyword


## volume [_volume]

Volume information for the quota report.

**`netapp.storage.quota_report.volume.name`**
:   Name of the volume.

type: keyword


**`netapp.storage.quota_report.volume.uuid`**
:   UUID of the volume.

type: keyword


## qtree [_qtree]

Qtree information for the quota report.

**`netapp.storage.quota_report.qtree.name`**
:   Name of the qtree.

type: keyword


**`netapp.storage.quota_report.qtree.id`**
:   ID of the qtree.

type: keyword


**`netapp.storage.quota_report.type`**
:   Type of the quota report.

type: keyword


**`netapp.storage.quota_report.index`**
:   Index of the quota report.

type: long


## group [_group]

Group information for the quota report.

**`netapp.storage.quota_report.group.name`**
:   Name of the group.

type: keyword


**`netapp.storage.quota_report.group.uuid`**
:   UUID of the group.

type: keyword


**`netapp.storage.quota_report.users`**
:   JSON string representing users for the quota report.

type: keyword


## files [_files]

File limits and usage for the quota report.

**`netapp.storage.quota_report.files.hard_limit`**
:   Hard file limit.

type: long


**`netapp.storage.quota_report.files.soft_limit`**
:   Soft file limit.

type: long


**`netapp.storage.quota_report.files.used`**
:   Number of files used.

type: long


## space [_space]

Space limits and usage for the quota report.

**`netapp.storage.quota_report.space.hard_limit`**
:   Hard space limit in bytes.

type: long


**`netapp.storage.quota_report.space.soft_limit`**
:   Soft space limit in bytes.

type: long


**`netapp.storage.quota_report.space.used`**
:   Space used in bytes.

type: long


## quota_rule [_quota_rule]

Quota rule information.

## files [_files]

File limits for the quota rule.

**`netapp.storage.quota_rule.files.hard_limit`**
:   Hard file limit.

type: long


**`netapp.storage.quota_rule.files.soft_limit`**
:   Soft file limit.

type: long


## space [_space]

Space limits for the quota rule.

**`netapp.storage.quota_rule.space.hard_limit`**
:   Hard space limit in bytes.

type: long


**`netapp.storage.quota_rule.space.soft_limit`**
:   Soft space limit in bytes.

type: long


## qtree [_qtree]

Brief qtree information for the quota rule.

**`netapp.storage.quota_rule.qtree.name`**
:   Name of the qtree.

type: keyword


**`netapp.storage.quota_rule.qtree.id`**
:   ID of the qtree.

type: keyword


**`netapp.storage.quota_rule.svm_name`**
:   Name of the SVM for the quota rule.

type: keyword


**`netapp.storage.quota_rule.type`**
:   Type of the quota rule.

type: keyword


**`netapp.storage.quota_rule.user_mapping`**
:   If true, mapping applies after user mapping resolved.

type: boolean


**`netapp.storage.quota_rule.users`**
:   JSON string representing users for the quota rule.

type: keyword


**`netapp.storage.quota_rule.uuid`**
:   UUID of the quota rule.

type: keyword


**`netapp.storage.quota_rule.volume`**
:   Name of the volume for the quota rule.

type: keyword


## volume [_volume]

Example group

**`netapp.storage.volume.uuid`**
:   UUID of the volume.

type: keyword


**`netapp.storage.volume.comment`**
:   Comment for the volume.

type: keyword


**`netapp.storage.volume.create_time`**
:   Creation time of the volume.

type: date


**`netapp.storage.volume.language`**
:   Language of the volume.

type: keyword


**`netapp.storage.volume.name`**
:   Name of the volume.

type: keyword


**`netapp.storage.volume.size`**
:   Size of the volume in bytes.

type: long


**`netapp.storage.volume.state`**
:   State of the volume.

type: keyword


**`netapp.storage.volume.style`**
:   Style of the volume.

type: keyword


**`netapp.storage.volume.tiering_policy`**
:   Tiering policy of the volume.

type: keyword


**`netapp.storage.volume.cloud_retrieval_policy`**
:   Cloud retrieval policy of the volume.

type: keyword


**`netapp.storage.volume.type`**
:   Type of the volume.

type: keyword


**`netapp.storage.volume.aggregates`**
:   JSON string representing aggregates for the volume.

type: keyword


**`netapp.storage.volume.snapshot_count`**
:   Number of snapshots for the volume.

type: long


**`netapp.storage.volume.msid`**
:   MSID of the volume.

type: keyword


**`netapp.storage.volume.scheduled_snapshot_naming_scheme`**
:   Scheduled snapshot naming scheme for the volume.

type: keyword


## clone [_clone]

Clone information for the volume.

**`netapp.storage.volume.clone.is_flexclone`**
:   Indicates if the volume is a FlexClone.

type: boolean


**`netapp.storage.volume.clone.has_flexclone`**
:   Indicates if the volume has FlexClones.

type: boolean


## nas [_nas]

NAS information for the volume.

**`netapp.storage.volume.nas.gid`**
:   GID for the NAS volume.

type: keyword


**`netapp.storage.volume.nas.security_style`**
:   Security style for the NAS volume.

type: keyword


**`netapp.storage.volume.nas.uid`**
:   UID for the NAS volume.

type: keyword


**`netapp.storage.volume.nas.unix_permissions`**
:   UNIX permissions for the NAS volume.

type: integer


## export_policy [_export_policy]

Export policy information for the NAS volume.

**`netapp.storage.volume.nas.export_policy.name`**
:   Name of the export policy.

type: keyword


**`netapp.storage.volume.nas.export_policy.id`**
:   ID of the export policy.

type: keyword


**`netapp.storage.volume.snapshot_locking_enabled`**
:   Indicates if snapshot locking is enabled for the volume.

type: boolean


## snapshot_policy [_snapshot_policy]

Snapshot policy information for the volume.

**`netapp.storage.volume.snapshot_policy.name`**
:   Name of the snapshot policy.

type: keyword


**`netapp.storage.volume.snapshot_policy.uuid`**
:   UUID of the snapshot policy.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the volume.

**`netapp.storage.volume.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.volume.svm.uuid`**
:   UUID of the SVM.

type: keyword


## space [_space]

Space information for the volume.

**`netapp.storage.volume.space.size`**
:   Total size of the volume in bytes.

type: long


**`netapp.storage.volume.space.available`**
:   Available space in bytes.

type: long


**`netapp.storage.volume.space.used`**
:   Used space in bytes.

type: long


**`netapp.storage.volume.space.is_used_stale`**
:   Indicates if the used space is stale.

type: boolean


**`netapp.storage.volume.space.block_storage_inactive_user_data`**
:   Block storage inactive user data in bytes.

type: long


**`netapp.storage.volume.space.local_tier_footprint`**
:   Local tier footprint in bytes.

type: long


**`netapp.storage.volume.space.footprint`**
:   Footprint in bytes.

type: long


**`netapp.storage.volume.space.over_provisioned`**
:   Indicates if the volume is over-provisioned.

type: integer


**`netapp.storage.volume.space.metadata`**
:   Metadata size in bytes.

type: long


**`netapp.storage.volume.space.total_footprint`**
:   Total footprint in bytes.

type: long


**`netapp.storage.volume.space.delayed_free_footprint`**
:   Delayed free footprint in bytes.

type: long


**`netapp.storage.volume.space.file_operation_metadata`**
:   File operation metadata in bytes.

type: long


**`netapp.storage.volume.space.volume_guarantee_footprint`**
:   Volume guarantee footprint in bytes.

type: long


**`netapp.storage.volume.space.effective_total_footprint`**
:   Effective total footprint in bytes.

type: long


**`netapp.storage.volume.space.user_data`**
:   User data size in bytes.

type: long


**`netapp.storage.volume.space.used_by_afs`**
:   Space used by AFS in bytes.

type: long


**`netapp.storage.volume.space.available_percent`**
:   Available space as a percentage.

type: float


**`netapp.storage.volume.space.afs_total`**
:   Total AFS space in bytes.

type: long


**`netapp.storage.volume.space.full_threshold_percent`**
:   Full threshold percentage.

type: float


**`netapp.storage.volume.space.nearly_full_threshold_percent`**
:   Nearly full threshold percentage.

type: float


**`netapp.storage.volume.space.overwrite_reserve`**
:   Overwrite reserve in bytes.

type: long


**`netapp.storage.volume.space.overwrite_reserve_used`**
:   Overwrite reserve used in bytes.

type: long


**`netapp.storage.volume.space.size_available_for_snapshots`**
:   Size available for snapshots in bytes.

type: long


**`netapp.storage.volume.space.percent_used`**
:   Percent of space used.

type: float


**`netapp.storage.volume.space.fractional_reserve`**
:   Fractional reserve percentage.

type: float


**`netapp.storage.volume.space.block_storage_inactive_user_data_pct`**
:   Block storage inactive user data percentage.

type: float


**`netapp.storage.volume.space.physical_used_percent`**
:   Physical used percentage.

type: float


**`netapp.storage.volume.space.physical_used`**
:   Physical used space in bytes.

type: long


**`netapp.storage.volume.space.expected_available`**
:   Expected available space in bytes.

type: long


**`netapp.storage.volume.space.filesystem_size`**
:   Filesystem size in bytes.

type: long


**`netapp.storage.volume.space.filesystem_size_fixed`**
:   Indicates if the filesystem size is fixed.

type: boolean


**`netapp.storage.volume.space.large_size_enabled`**
:   Indicates if large size is enabled.

type: boolean


**`netapp.storage.volume.space.total_metadata`**
:   Total metadata in bytes.

type: long


**`netapp.storage.volume.space.total_metadata_footprint`**
:   Total metadata footprint in bytes.

type: long


## logical_space [_logical_space]

Logical space information for the volume.

**`netapp.storage.volume.space.logical_space.reporting`**
:   Indicates if logical space reporting is enabled.

type: boolean


**`netapp.storage.volume.space.logical_space.enforcement`**
:   Indicates if logical space enforcement is enabled.

type: boolean


**`netapp.storage.volume.space.logical_space.used_by_afs`**
:   Logical space used by AFS in bytes.

type: long


**`netapp.storage.volume.space.logical_space.used_percent`**
:   Logical space used percent.

type: float


**`netapp.storage.volume.space.logical_space.used`**
:   Logical space used in bytes.

type: long


**`netapp.storage.volume.space.logical_space.used_by_snapshots`**
:   Logical space used by snapshots in bytes.

type: long


## snapshot [_snapshot]

Snapshot space information for the volume.

**`netapp.storage.volume.space.snapshot.used`**
:   Snapshot space used in bytes.

type: long


**`netapp.storage.volume.space.snapshot.reserve_percent`**
:   Snapshot reserve percent.

type: float


**`netapp.storage.volume.space.snapshot.autodelete_enabled`**
:   Indicates if snapshot autodelete is enabled.

type: boolean


**`netapp.storage.volume.space.snapshot.reserve_size`**
:   Snapshot reserve size in bytes.

type: long


**`netapp.storage.volume.space.snapshot.space_used_percent`**
:   Snapshot space used percent.

type: float


**`netapp.storage.volume.space.snapshot.reserve_available`**
:   Snapshot reserve available in bytes.

type: long


**`netapp.storage.volume.space.snapshot.autodelete_trigger`**
:   Snapshot autodelete trigger.

type: keyword


## autodelete [_autodelete]

Autodelete information for the snapshot.

**`netapp.storage.volume.space.snapshot.autodelete.enabled`**
:   Indicates if autodelete is enabled.

type: boolean


**`netapp.storage.volume.space.snapshot.autodelete.trigger`**
:   Trigger for autodelete.

type: keyword


**`netapp.storage.volume.space.snapshot.autodelete.delete_order`**
:   Delete order for autodelete.

type: keyword


**`netapp.storage.volume.space.snapshot.autodelete.defer_delete`**
:   Type of defer delete for autodelete, e.g. "user_created"

type: keyword


**`netapp.storage.volume.space.snapshot.autodelete.commitment`**
:   Commitment for autodelete.

type: keyword


**`netapp.storage.volume.space.snapshot.autodelete.target_free_space`**
:   Target free space in bytes for autodelete.

type: long


**`netapp.storage.volume.space.snapshot.autodelete.prefix`**
:   Prefix for autodelete.

type: keyword


## metrics [_metrics]

Metrics for the volume.

**`netapp.storage.volume.metrics.timestamp`**
:   Timestamp of the metric sample.

type: date


**`netapp.storage.volume.metrics.duration`**
:   Duration in ISO 8601 format (e.g. "PT25M5S").

type: keyword


**`netapp.storage.volume.metrics.status`**
:   Status of the metric sample.

type: keyword


## latency [_latency]

IO latency metrics.

**`netapp.storage.volume.metrics.latency.read`**
:   Read IO latency in milliseconds.

type: float


**`netapp.storage.volume.metrics.latency.write`**
:   Write IO latency in milliseconds.

type: float


**`netapp.storage.volume.metrics.latency.other`**
:   Other IO latency in milliseconds.

type: float


**`netapp.storage.volume.metrics.latency.total`**
:   Total IO latency in milliseconds.

type: float


## iops [_iops]

Input/output operations per second metrics.

**`netapp.storage.volume.metrics.iops.read`**
:   Read IOPS.

type: float


**`netapp.storage.volume.metrics.iops.write`**
:   Write IOPS.

type: float


**`netapp.storage.volume.metrics.iops.other`**
:   Other IOPS.

type: float


**`netapp.storage.volume.metrics.iops.total`**
:   Total IOPS.

type: float


## throughput [_throughput]

Throughput metrics in bytes per second.

**`netapp.storage.volume.metrics.throughput.read`**
:   Read throughput in bytes per second.

type: float


**`netapp.storage.volume.metrics.throughput.write`**
:   Write throughput in bytes per second.

type: float


**`netapp.storage.volume.metrics.throughput.other`**
:   Other throughput in bytes per second.

type: float


**`netapp.storage.volume.metrics.throughput.total`**
:   Total throughput in bytes per second.

type: float


## snapmirror [_snapmirror]

SnapMirror information for the volume.

**`netapp.storage.volume.snapmirror.is_protected`**
:   Indicates if the volume is protected by SnapMirror.

type: boolean


## destinations [_destinations]

SnapMirror destination information.

**`netapp.storage.volume.snapmirror.destinations.is_ontap`**
:   Indicates if the destination is ONTAP.

type: boolean


**`netapp.storage.volume.snapmirror.destinations.is_cloud`**
:   Indicates if the destination is cloud.

type: boolean


## activity_tracking [_activity_tracking]

Activity tracking information for the volume.

**`netapp.storage.volume.activity_tracking.supported`**
:   Indicates if activity tracking is supported for the volume.

type: boolean


**`netapp.storage.volume.activity_tracking.state`**
:   State of activity tracking for the volume.

type: keyword


**`netapp.storage.volume.granular_data`**
:   Indicates if granular data is enabled for the volume.

type: boolean


**`netapp.storage.volume.granular_data_mode`**
:   Granular data mode for the volume.

type: keyword


## svm_peer [_svm_peer]

SVM peer information.

**`netapp.storage.svm_peer.applications`**
:   JSON string representing the applications associated with the SVM peer.

type: keyword


**`netapp.storage.svm_peer.name`**
:   Name of the SVM peer.

type: keyword


## peer [_peer]

Peer information for the SVM peer.

## cluster [_cluster]

Cluster information for the peer.

**`netapp.storage.svm_peer.peer.cluster.name`**
:   Name of the cluster.

type: keyword


**`netapp.storage.svm_peer.peer.cluster.uuid`**
:   UUID of the cluster.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the peer.

**`netapp.storage.svm_peer.peer.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.svm_peer.peer.svm.uuid`**
:   UUID of the SVM.

type: keyword


**`netapp.storage.svm_peer.state`**
:   State of the SVM peer relationship.

type: keyword


## svm [_svm]

SVM (Storage Virtual Machine) information for the SVM peer.

**`netapp.storage.svm_peer.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.svm_peer.svm.uuid`**
:   UUID of the SVM.

type: keyword


**`netapp.storage.svm_peer.uuid`**
:   UUID of the SVM peer.

type: keyword


## svm [_svm]

Example group

**`netapp.storage.svm.uuid`**
:   UUID of the SVM.

type: keyword


**`netapp.storage.svm.name`**
:   Name of the SVM.

type: keyword


**`netapp.storage.svm.subtype`**
:   Subtype of the SVM.

type: keyword


**`netapp.storage.svm.language`**
:   Language of the SVM.

type: keyword


**`netapp.storage.svm.aggregates`**
:   JSON string representing aggregates for the SVM.

type: keyword


**`netapp.storage.svm.state`**
:   State of the SVM.

type: keyword


**`netapp.storage.svm.comment`**
:   Comment for the SVM.

type: keyword


## ipspace [_ipspace]

IPSpace information for the SVM.

**`netapp.storage.svm.ipspace.name`**
:   Name of the IPSpace.

type: keyword


**`netapp.storage.svm.ipspace.uuid`**
:   UUID of the IPSpace.

type: keyword


**`netapp.storage.svm.ip_interfaces`**
:   JSON string representing IP interfaces for the SVM.

type: keyword


## snapshot_policy [_snapshot_policy]

Snapshot policy information for the SVM.

**`netapp.storage.svm.snapshot_policy.name`**
:   Name of the snapshot policy.

type: keyword


**`netapp.storage.svm.snapshot_policy.uuid`**
:   UUID of the snapshot policy.

type: keyword


**`netapp.storage.svm.nis_enabled`**
:   Indicates if NIS is enabled for the SVM.

type: boolean


**`netapp.storage.svm.ldap_enabled`**
:   Indicates if LDAP is enabled for the SVM.

type: boolean


## nfs [_nfs]

NFS protocol status for the SVM.

**`netapp.storage.svm.nfs.allowed`**
:   Indicates if NFS is allowed.

type: boolean


**`netapp.storage.svm.nfs.enabled`**
:   Indicates if NFS is enabled.

type: boolean


## cifs [_cifs]

CIFS protocol status for the SVM.

**`netapp.storage.svm.cifs.allowed`**
:   Indicates if CIFS is allowed.

type: boolean


**`netapp.storage.svm.cifs.enabled`**
:   Indicates if CIFS is enabled.

type: boolean


## iscsi [_iscsi]

iSCSI protocol status for the SVM.

**`netapp.storage.svm.iscsi.allowed`**
:   Indicates if iSCSI is allowed.

type: boolean


**`netapp.storage.svm.iscsi.enabled`**
:   Indicates if iSCSI is enabled.

type: boolean


## fcp [_fcp]

FCP protocol status for the SVM.

**`netapp.storage.svm.fcp.allowed`**
:   Indicates if FCP is allowed.

type: boolean


**`netapp.storage.svm.fcp.enabled`**
:   Indicates if FCP is enabled.

type: boolean


## nvme [_nvme]

NVMe protocol status for the SVM.

**`netapp.storage.svm.nvme.allowed`**
:   Indicates if NVMe is allowed.

type: boolean


**`netapp.storage.svm.nvme.enabled`**
:   Indicates if NVMe is enabled.

type: boolean


**`netapp.storage.svm.ndmp_allowed`**
:   Indicates if NDMP is allowed for the SVM.

type: boolean


## s3 [_s3]

S3 protocol status for the SVM.

**`netapp.storage.svm.s3.allowed`**
:   Indicates if S3 is allowed.

type: boolean


**`netapp.storage.svm.s3.enabled`**
:   Indicates if S3 is enabled.

type: boolean


**`netapp.storage.svm.certificate`**
:   UUID of the certificate for the SVM.

type: keyword


**`netapp.storage.svm.aggregates_delegated`**
:   Indicates if aggregates are delegated for the SVM.

type: boolean


**`netapp.storage.svm.retention_period`**
:   Retention period for the SVM.

type: long


**`netapp.storage.svm.max_volumes`**
:   Maximum number of volumes for the SVM: can be set to a specific value or unlimited.

type: keyword


**`netapp.storage.svm.anti_ransomware_default_volume_state`**
:   Default anti-ransomware state for volumes in the SVM.

type: keyword


**`netapp.storage.svm.is_space_reporting_logical`**
:   Indicates if space reporting is logical for the SVM.

type: boolean


**`netapp.storage.svm.is_space_enforcement_logical`**
:   Indicates if space enforcement is logical for the SVM.

type: boolean


**`netapp.storage.svm.auto_enable_analytics`**
:   Indicates if analytics are auto-enabled for the SVM.

type: boolean


**`netapp.storage.svm.auto_enable_activity_tracking`**
:   Indicates if activity tracking is auto-enabled for the SVM.

type: boolean


**`netapp.storage.svm.anti_ransomware_auto_switch_enabled`**
:   Indicates if anti-ransomware auto-switch is enabled for the SVM.

type: boolean


**`netapp.storage.svm.anti_ransomware_auto_switch_data_percent`**
:   Data percent threshold for anti-ransomware auto-switch.

type: float


**`netapp.storage.svm.anti_ransomware_auto_switch_no_ext_days`**
:   Number of days with no extensions for anti-ransomware auto-switch.

type: long


**`netapp.storage.svm.anti_ransomware_auto_switch_min_period`**
:   Minimum period for anti-ransomware auto-switch.

type: long


**`netapp.storage.svm.anti_ransomware_auto_switch_min_files`**
:   Minimum files for anti-ransomware auto-switch.

type: long


**`netapp.storage.svm.anti_ransomware_auto_switch_min_exts`**
:   Minimum extensions for anti-ransomware auto-switch. Example field

type: long


## disk [_disk]

Disk related information.

**`netapp.storage.disk.name`**
:   Name of the disk.

type: keyword


**`netapp.storage.disk.uid`**
:   UID of the disk.

type: keyword


**`netapp.storage.disk.serial_number`**
:   Serial number of the disk.

type: keyword


**`netapp.storage.disk.model`**
:   Model of the disk.

type: keyword


**`netapp.storage.disk.vendor`**
:   Vendor of the disk.

type: keyword


**`netapp.storage.disk.firmware_version`**
:   Firmware version of the disk.

type: keyword


**`netapp.storage.disk.usable_size`**
:   Usable size of the disk in bytes.

type: long


**`netapp.storage.disk.rated_life_used_percent`**
:   Rated life used percent of the disk.

type: float


**`netapp.storage.disk.type`**
:   Type of the disk.

type: keyword


**`netapp.storage.disk.effective_type`**
:   Effective type of the disk.

type: keyword


**`netapp.storage.disk.class`**
:   Class of the disk.

type: keyword


**`netapp.storage.disk.container_type`**
:   Container type of the disk.

type: keyword


**`netapp.storage.disk.pool`**
:   Pool of the disk.

type: keyword


**`netapp.storage.disk.state`**
:   State of the disk.

type: keyword


## node [_node]

Node information for the disk.

**`netapp.storage.disk.node.name`**
:   Name of the node.

type: keyword


**`netapp.storage.disk.node.uuid`**
:   UUID of the node.

type: keyword


## home_node [_home_node]

Home node information for the disk.

**`netapp.storage.disk.home_node.name`**
:   Name of the home node.

type: keyword


**`netapp.storage.disk.home_node.uuid`**
:   UUID of the home node.

type: keyword


**`netapp.storage.disk.aggregates`**
:   JSON string containing Aggregates associated with the disk.

type: keyword


**`netapp.storage.disk.shelf_uuid`**
:   UUID of the shelf containing the disk.

type: keyword


**`netapp.storage.disk.local`**
:   Indicates if the disk is local.

type: boolean


**`netapp.storage.disk.bay`**
:   Bay number of the disk.

type: long


**`netapp.storage.disk.self_encrypting`**
:   Indicates if the disk is self-encrypting.

type: boolean


**`netapp.storage.disk.fips_certified`**
:   Indicates if the disk is FIPS certified.

type: boolean


**`netapp.storage.disk.bytes_per_sector`**
:   Number of bytes per sector on the disk.

type: long


**`netapp.storage.disk.sector_count`**
:   Total sector count of the disk.

type: long


**`netapp.storage.disk.right_size_sector_count`**
:   Right size sector count of the disk.

type: long


**`netapp.storage.disk.physical_size`**
:   Physical size of the disk in bytes.

type: long


## stats [_stats]

Statistics for the disk.

**`netapp.storage.disk.stats.average_latency`**
:   Average latency of the disk in milliseconds.

type: float


**`netapp.storage.disk.stats.throughput`**
:   Throughput of the disk in bytes per second.

type: float


**`netapp.storage.disk.stats.iops_total`**
:   Total input/output operations per second for the disk.

type: float


**`netapp.storage.disk.stats.path_error_count`**
:   Number of path errors for the disk.

type: long


**`netapp.storage.disk.stats.power_on_hours`**
:   Number of hours the disk has been powered on.

type: long


## shelf [_shelf]

Shelf related information.

**`netapp.storage.shelf.uid`**
:   UID of the shelf.

type: keyword


**`netapp.storage.shelf.name`**
:   Name of the shelf.

type: keyword


**`netapp.storage.shelf.id`**
:   ID of the shelf.

type: keyword


**`netapp.storage.shelf.serial_number`**
:   Serial number of the shelf.

type: keyword


**`netapp.storage.shelf.model`**
:   Model of the shelf.

type: keyword


**`netapp.storage.shelf.module_type`**
:   Module type of the shelf.

type: keyword


**`netapp.storage.shelf.internal`**
:   Indicates if the shelf is internal.

type: boolean


**`netapp.storage.shelf.local`**
:   Indicates if the shelf is local.

type: boolean


**`netapp.storage.shelf.manufacturer`**
:   Manufacturer name of the shelf.

type: keyword


**`netapp.storage.shelf.state`**
:   State of the shelf.

type: keyword


**`netapp.storage.shelf.connection_type`**
:   Connection type of the shelf.

type: keyword


**`netapp.storage.shelf.disk_count`**
:   Number of disks in the shelf.

type: long


**`netapp.storage.shelf.location_led`**
:   Location LED status of the shelf.

type: keyword


**`netapp.storage.shelf.paths`**
:   JSON string representing the paths of the shelf.

type: keyword


**`netapp.storage.shelf.bays`**
:   JSON string representing the bays of the shelf.

type: keyword


## acp [_acp]

ACP (Automatic Configuration Protocol) information for the shelf.

**`netapp.storage.shelf.acp.enabled`**
:   Indicates if ACP is enabled for the shelf.

type: boolean


**`netapp.storage.shelf.acp.channel`**
:   Channel of the ACP for the shelf.

type: keyword


**`netapp.storage.shelf.acp.connection_state`**
:   Connection state of the ACP for the shelf.

type: keyword


## node [_node]

Node information for the ACP.

**`netapp.storage.shelf.acp.node.name`**
:   Name of the node.

type: keyword


**`netapp.storage.shelf.acp.node.uuid`**
:   UUID of the node.

type: keyword


## port [_port]

Port information for the shelf.

**`netapp.storage.shelf.port.id`**
:   ID of the port.

type: keyword


**`netapp.storage.shelf.port.module_id`**
:   Module ID of the port.

type: keyword


**`netapp.storage.shelf.port.designator`**
:   Designator of the port.

type: keyword


**`netapp.storage.shelf.port.state`**
:   State of the port.

type: keyword


**`netapp.storage.shelf.port.internal`**
:   Indicates if the port is internal.

type: boolean


**`netapp.storage.shelf.port.wwn`**
:   WWN of the port.

type: keyword


**`netapp.storage.shelf.port.cable_id`**
:   Identifier of the cable attached to the port.

type: keyword


**`netapp.storage.shelf.port.remote_wwn`**
:   Remote WWN connected to the port.

type: keyword


## current_sensor [_current_sensor]

Current sensor information for the shelf.

**`netapp.storage.shelf.current_sensor.id`**
:   ID of the current sensor.

type: keyword


**`netapp.storage.shelf.current_sensor.location`**
:   Location of the current sensor.

type: keyword


**`netapp.storage.shelf.current_sensor.current`**
:   Current value measured by the sensor (in Amperes).

type: float


**`netapp.storage.shelf.current_sensor.state`**
:   State of the current sensor.

type: keyword


**`netapp.storage.shelf.current_sensor.installed`**
:   Indicates if the current sensor is installed.

type: boolean


## voltage_sensor [_voltage_sensor]

Voltage sensor information for the shelf.

**`netapp.storage.shelf.voltage_sensor.id`**
:   ID of the voltage sensor.

type: keyword


**`netapp.storage.shelf.voltage_sensor.location`**
:   Location of the voltage sensor.

type: keyword


**`netapp.storage.shelf.voltage_sensor.voltage`**
:   Voltage value measured by the sensor (in Volts).

type: float


**`netapp.storage.shelf.voltage_sensor.state`**
:   State of the voltage sensor.

type: keyword


**`netapp.storage.shelf.voltage_sensor.installed`**
:   Indicates if the voltage sensor is installed.

type: boolean


## temperature_sensor [_temperature_sensor]

Temperature sensor information for the shelf.

**`netapp.storage.shelf.temperature_sensor.id`**
:   ID of the temperature sensor.

type: keyword


**`netapp.storage.shelf.temperature_sensor.location`**
:   Location of the temperature sensor.

type: keyword


**`netapp.storage.shelf.temperature_sensor.temperature`**
:   Temperature value measured by the sensor (in Celsius).

type: float


**`netapp.storage.shelf.temperature_sensor.is_ambient`**
:   Indicates if the temperature sensor is measuring ambient temperature.

type: boolean


**`netapp.storage.shelf.temperature_sensor.state`**
:   State of the temperature sensor.

type: keyword


**`netapp.storage.shelf.temperature_sensor.installed`**
:   Indicates if the temperature sensor is installed.

type: boolean


## threshold [_threshold]

Threshold information for the temperature sensor.

## high [_high]

High threshold values.

**`netapp.storage.shelf.temperature_sensor.threshold.high.critical`**
:   Critical high temperature threshold.

type: float


**`netapp.storage.shelf.temperature_sensor.threshold.high.warning`**
:   Warning high temperature threshold.

type: float


## low [_low]

Low threshold values.

**`netapp.storage.shelf.temperature_sensor.threshold.low.critical`**
:   Critical low temperature threshold.

type: float


**`netapp.storage.shelf.temperature_sensor.threshold.low.warning`**
:   Warning low temperature threshold.

type: float


## fan [_fan]

Fan information for the shelf.

**`netapp.storage.shelf.fan.id`**
:   ID of the fan.

type: keyword


**`netapp.storage.shelf.fan.location`**
:   Location of the fan.

type: keyword


**`netapp.storage.shelf.fan.rpm`**
:   Rotations per minute of the fan.

type: long


**`netapp.storage.shelf.fan.state`**
:   State of the fan.

type: keyword


**`netapp.storage.shelf.fan.installed`**
:   Indicates if the fan is installed.

type: boolean


## psu [_psu]

Power Supply Unit (PSU) information for the shelf.

**`netapp.storage.shelf.psu.type`**
:   Type of the PSU.

type: keyword


**`netapp.storage.shelf.psu.id`**
:   ID of the PSU.

type: keyword


**`netapp.storage.shelf.psu.state`**
:   State of the PSU.

type: keyword


**`netapp.storage.shelf.psu.part_number`**
:   Part number of the PSU.

type: keyword


**`netapp.storage.shelf.psu.serial_number`**
:   Serial number of the PSU.

type: keyword


**`netapp.storage.shelf.psu.firmware_version`**
:   Firmware version of the PSU.

type: keyword


**`netapp.storage.shelf.psu.installed`**
:   Indicates if the PSU is installed.

type: boolean


**`netapp.storage.shelf.psu.model`**
:   Model of the PSU.

type: keyword


**`netapp.storage.shelf.psu.power_drawn`**
:   Power drawn by the PSU (in Watts).

type: float


**`netapp.storage.shelf.psu.power_rating`**
:   Power rating of the PSU (in Watts).

type: float


**`netapp.storage.shelf.psu.crest_factor`**
:   Crest factor of the PSU.

type: float


