---
mapped_pages:
  - https://www.elastic.co/guide/en/beats/metricbeat/current/exported-fields-emcunity.html
---

% This file is generated! See scripts/generate_fields_docs.py

# emcunity fields [exported-fields-emcunity]

emcunity module

## emcunity [_emcunity]



## health [_health]

health

**`emcunity.health.message`**
:   debug message of json message

type: text


**`emcunity.health.system.id`**
:   id of system

type: keyword


**`emcunity.health.system.name`**
:   name of system

type: keyword


**`emcunity.health.system.macaddress`**
:   macaddress of system

type: keyword


**`emcunity.health.system.model`**
:   model of system

type: keyword


**`emcunity.health.system.serial_number`**
:   searial number of the system

type: keyword


**`emcunity.health.system.internal_model`**
:   internal model of the system

type: keyword


**`emcunity.health.system.platform`**
:   platform of system

type: keyword


**`emcunity.health.system.health.value`**
:   value of health score

type: long


**`emcunity.health.system.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.system.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.pool.id`**
:   id of pool

type: keyword


**`emcunity.health.pool.name`**
:   name of pool

type: keyword


**`emcunity.health.pool.health.value`**
:   value of health score

type: long


**`emcunity.health.pool.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.pool.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.pool.description`**
:   description of pool

type: keyword


**`emcunity.health.pool.size.free`**
:   size free in the pool

type: long


**`emcunity.health.pool.size.used`**
:   size used in the pool

type: long


**`emcunity.health.pool.size.total`**
:   size total in the pool

type: long


**`emcunity.health.pool.size.subscribed`**
:   size subscribed in the pool

type: long


**`emcunity.health.pool.harvest.state`**
:   harvest state of the pool

type: long


**`emcunity.health.pool.metadata.size.subscribed`**
:   metadata size subscribed of the pool

type: long


**`emcunity.health.pool.metadata.size.used`**
:   metadata size used of the pool

type: long


**`emcunity.health.pool.snap.size.subscribed`**
:   snap size suscribed of the pool

type: long


**`emcunity.health.pool.snap.size.used`**
:   snap size used of the pool

type: long


**`emcunity.health.pool.rebalance.progress`**
:   rebalance progress of the pool

type: long


**`emcunity.health.pool.size.percent.used`**
:   pecent of pool used

type: long


**`emcunity.health.pool.unit.id`**
:   id of pool unit

type: keyword


**`emcunity.health.pool.unit.name`**
:   name of pool unit

type: keyword


**`emcunity.health.pool.unit.health.value`**
:   value of health score

type: long


**`emcunity.health.pool.unit.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.pool.unit.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.pool.unit.size.total`**
:   size total in the pool unit

type: long


**`emcunity.health.lun.id`**
:   id of lun

type: keyword


**`emcunity.health.lun.health.value`**
:   value of health score

type: long


**`emcunity.health.lun.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.lun.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.lun.size.total`**
:   size total in the lun

type: long


**`emcunity.health.lun.size.allocated`**
:   size allocated in the lun

type: long


**`emcunity.health.lun.metadata.size.total`**
:   metadata size total of the lun

type: long


**`emcunity.health.lun.metadata.size.allocated`**
:   metadata size allocated of the lun

type: long


**`emcunity.health.lun.snap.size.total`**
:   snap size total of the lun

type: long


**`emcunity.health.lun.snap.size.allocated`**
:   snap size allocated of the lun

type: long


**`emcunity.health.storage.processor.id`**
:   id of storage processor

type: keyword


**`emcunity.health.storage.processor.name`**
:   name of storage processor

type: keyword


**`emcunity.health.storage.processor.model`**
:   model of storage processor

type: keyword


**`emcunity.health.storage.processor.health.value`**
:   value of health score

type: long


**`emcunity.health.storage.processor.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.storage.processor.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.storage.resource.id`**
:   id of storage processor

type: keyword


**`emcunity.health.storage.resource.health.value`**
:   value of health score

type: long


**`emcunity.health.storage.resource.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.storage.resource.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.storage.resource.size.total`**
:   size total in the storage resource

type: long


**`emcunity.health.storage.resource.size.allocated`**
:   size allocated in the storage resource

type: long


**`emcunity.health.storage.resource.snap.count`**
:   snap count of the storage resource

type: long


**`emcunity.health.storage.resource.snaps.size.total`**
:   snaps size total of the storage resource

type: long


**`emcunity.health.storage.resource.snaps.size.allocated`**
:   snaps size allocated of the storage resource

type: long


**`emcunity.health.storage.resource.metadata.size.total`**
:   metadata size total of the storage resource

type: long


**`emcunity.health.storage.resource.metadata.size.allocated`**
:   metadata size allocated of the storage resource

type: long


**`emcunity.health.storage.tier.id`**
:   id of storage tier

type: keyword


**`emcunity.health.storage.tier.size.total`**
:   size total in the storage tier

type: long


**`emcunity.health.storage.tier.size.free`**
:   size free in the storage tier

type: long


**`emcunity.health.storage.tier.disk.total`**
:   disk total in the storage tier

type: long


**`emcunity.health.storage.tier.disk.unused`**
:   disk unused in the storage tier

type: long


**`emcunity.health.storage.tier.virtual.disk.total`**
:   virtual disk total in the storage tier

type: long


**`emcunity.health.storage.tier.virtual.disk.unused`**
:   virtual disk unused in the storage tier

type: long


**`emcunity.health.storage.tier.virtual.disk.used`**
:   virtual disk used in the storage tier

type: long


**`emcunity.health.storage.tier.disk.used`**
:   disk used in the storage tier

type: long


**`emcunity.health.license.id`**
:   id of license

type: keyword


**`emcunity.health.license.name`**
:   name of license

type: keyword


**`emcunity.health.license.isvalid`**
:   boolean is license valid

type: boolean


**`emcunity.health.license.ispermanent`**
:   boolean is license pemanent

type: boolean


**`emcunity.health.license.isinstalled`**
:   boolean is license installed

type: boolean


**`emcunity.health.license.expires`**
:   date license expires

type: date


**`emcunity.health.license.feature.id`**
:   license feature id

type: keyword


**`emcunity.health.ethernet.port.id`**
:   id of ethernet port

type: keyword


**`emcunity.health.ethernet.port.health.value`**
:   value of health score

type: long


**`emcunity.health.ethernet.port.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.ethernet.port.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.file.interface.id`**
:   id of file interface

type: keyword


**`emcunity.health.file.interface.health.value`**
:   value of health score

type: long


**`emcunity.health.file.interface.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.file.interface.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.remote.system.id`**
:   id of remote system

type: keyword


**`emcunity.health.remote.system.health.value`**
:   value of health score

type: long


**`emcunity.health.remote.system.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.remote.system.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.disk.id`**
:   id of disk

type: keyword


**`emcunity.health.disk.name`**
:   name of disk

type: keyword


**`emcunity.health.disk.size`**
:   size of disk

type: long


**`emcunity.health.disk.vendor.size`**
:   vendor size of disk

type: long


**`emcunity.health.disk.raw.size`**
:   raw size of disk

type: long


**`emcunity.health.disk.health.value`**
:   value of health score

type: long


**`emcunity.health.disk.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.disk.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.datastore.id`**
:   id of datastore

type: keyword


**`emcunity.health.datastore.name`**
:   name of datastore

type: keyword


**`emcunity.health.datastore.size.total`**
:   size total of datastore

type: long


**`emcunity.health.datastore.size.used`**
:   size used of datastore

type: long


**`emcunity.health.datastore.size.percent.used`**
:   size percent used of datastore

type: long


**`emcunity.health.filesystem.id`**
:   id of filesystem

type: keyword


**`emcunity.health.filesystem.name`**
:   name of filesystem

type: keyword


**`emcunity.health.filesystem.health.value`**
:   value of health score

type: long


**`emcunity.health.filesystem.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.filesystem.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.filesystem.size.total`**
:   size total in filesystem

type: long


**`emcunity.health.filesystem.size.used`**
:   size used in filesystem

type: long


**`emcunity.health.filesystem.size.allocated`**
:   size allocated in filesystem

type: long


**`emcunity.health.filesystem.metadata.size.total`**
:   size total in filesystem

type: long


**`emcunity.health.filesystem.metadata.size.allocated`**
:   size allocated in filesystem

type: long


**`emcunity.health.filesystem.snap.count`**
:   snap count in filesystem

type: long


**`emcunity.health.filesystem.snaps.size.total`**
:   snaps size total in filesystem

type: long


**`emcunity.health.filesystem.snaps.size.allocated`**
:   snaps size allocated in filesystem

type: long


**`emcunity.health.filesystem.size.percent.used`**
:   size percent used in filesystem

type: long


**`emcunity.health.snap.id`**
:   id of snap

type: keyword


**`emcunity.health.snap.name`**
:   name of snap

type: keyword


**`emcunity.health.snap.size.total`**
:   size total of snap

type: long


**`emcunity.health.snap.state`**
:   state of snap

type: long


**`emcunity.health.snap.creation.time`**
:   creation time of snap

type: date


**`emcunity.health.snap.expiration.time`**
:   expiration time of snap

type: date


**`emcunity.health.sas.port.id`**
:   id of sas port

type: keyword


**`emcunity.health.sas.port.name`**
:   name of sas port

type: keyword


**`emcunity.health.sas.port.needs_replacement`**
:   boolean needs_replacement of sas port

type: boolean


**`emcunity.health.sas.port.health.value`**
:   value of health score

type: long


**`emcunity.health.sas.port.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.sas.port.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.power.supply.id`**
:   id of power supply

type: keyword


**`emcunity.health.power.supply.name`**
:   name of power supply

type: keyword


**`emcunity.health.power.supply.needs_replacement`**
:   boolean needs_replacement of power supply

type: boolean


**`emcunity.health.power.supply.health.value`**
:   value of health score

type: long


**`emcunity.health.power.supply.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.power.supply.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.fan.id`**
:   id of fan

type: keyword


**`emcunity.health.fan.name`**
:   name of fan

type: keyword


**`emcunity.health.fan.needs_replacement`**
:   boolean needs_replacement of fan

type: boolean


**`emcunity.health.fan.health.value`**
:   value of health score

type: long


**`emcunity.health.fan.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.fan.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.dae.id`**
:   id of disk array enclosure (dae)

type: keyword


**`emcunity.health.dae.name`**
:   name of disk array enclosure (dae)

type: keyword


**`emcunity.health.dae.needs_replacement`**
:   boolean needs_replacement of dae

type: boolean


**`emcunity.health.dae.health.value`**
:   value of health score

type: long


**`emcunity.health.dae.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.dae.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.dae.current.power`**
:   current power of dae

type: long


**`emcunity.health.dae.avg.power`**
:   avf power of dae

type: long


**`emcunity.health.dae.max.power`**
:   max power of dae

type: long


**`emcunity.health.dae.current.temperature`**
:   current temperature of dae

type: long


**`emcunity.health.dae.avg.temperature`**
:   avf temperature of dae

type: long


**`emcunity.health.dae.max.temperature`**
:   max temperature of dae

type: long


**`emcunity.health.memory.module.id`**
:   id of memory module

type: keyword


**`emcunity.health.memory.module.name`**
:   name of memory module

type: keyword


**`emcunity.health.memory.module.size`**
:   size of memory module

type: long


**`emcunity.health.memory.module.needs_replacement`**
:   boolean needs_replacement of memory module

type: boolean


**`emcunity.health.memory.module.health.value`**
:   value of health score

type: long


**`emcunity.health.memory.module.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.memory.module.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.battery.id`**
:   id of battery

type: keyword


**`emcunity.health.battery.name`**
:   name of battery

type: keyword


**`emcunity.health.battery.needs_replacement`**
:   boolean needs_replacement of battery

type: boolean


**`emcunity.health.battery.health.value`**
:   value of health score

type: long


**`emcunity.health.battery.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.battery.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.ssd.id`**
:   id of ssd

type: keyword


**`emcunity.health.ssd.name`**
:   name of fan

type: keyword


**`emcunity.health.ssd.needs_replacement`**
:   boolean needs_replacement of ssd

type: boolean


**`emcunity.health.ssd.health.value`**
:   value of health score

type: long


**`emcunity.health.ssd.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.ssd.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.raid.group.id`**
:   id of raid group

type: keyword


**`emcunity.health.raid.group.name`**
:   name of raid group

type: keyword


**`emcunity.health.raid.group.size.total`**
:   total size of raid group

type: long


**`emcunity.health.raid.group.needs_replacement`**
:   boolean needs_replacement of raid group

type: boolean


**`emcunity.health.raid.group.health.value`**
:   value of health score

type: long


**`emcunity.health.raid.group.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.raid.group.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.tree.quota.id`**
:   id of tree quota

type: keyword


**`emcunity.health.tree.quota.soft.limit`**
:   soft limit of tree quota

type: long


**`emcunity.health.tree.quota.hard.limit`**
:   hard limit of tree quota

type: long


**`emcunity.health.tree.quota.size.used`**
:   size used of tree quota

type: long


**`emcunity.health.tree.quota.size.state`**
:   size state of tree quota

type: long


**`emcunity.health.disk.group.id`**
:   id of disk group

type: keyword


**`emcunity.health.disk.group.name`**
:   name of disk group

type: keyword


**`emcunity.health.disk.group.advertised.size`**
:   advertised size of disk group

type: long


**`emcunity.health.disk.group.disk.size`**
:   disk size of disk group

type: long


**`emcunity.health.disk.group.hot.spare.policy.status`**
:   hot spare policy status of disk group

type: long


**`emcunity.health.disk.group.min.spare.policy.status`**
:   min spare policy status of disk group

type: long


**`emcunity.health.disk.group.rpm`**
:   rpm of disk group

type: long


**`emcunity.health.disk.group.speed`**
:   speed of disk group

type: long


**`emcunity.health.disk.group.total.disks`**
:   total disks of disk group

type: long


**`emcunity.health.disk.group.unconfigured.disks`**
:   unconfigured disks of disk group

type: long


**`emcunity.health.cifs.server.id`**
:   id of cifs server

type: keyword


**`emcunity.health.cifs.server.name`**
:   name of cifs server

type: keyword


**`emcunity.health.cifs.server.health.value`**
:   value of health score

type: long


**`emcunity.health.cifs.server.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.cifs.server.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.fast.cache.id`**
:   id of fast cache

type: keyword


**`emcunity.health.fast.cache.name`**
:   name of fast cache

type: keyword


**`emcunity.health.fast.cache.number_of_disks`**
:   number of disks of fast cache

type: long


**`emcunity.health.fast.cache.size.free`**
:   size free of fast cache

type: long


**`emcunity.health.fast.cache.size.total`**
:   size total of fast cache

type: long


**`emcunity.health.fast.cache.health.value`**
:   value of health score

type: long


**`emcunity.health.fast.cache.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.fast.cache.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.fast.vp.id`**
:   id of fast vp

type: keyword


**`emcunity.health.fast.vp.status`**
:   name of fast vp

type: long


**`emcunity.health.fast.vp.is_schedule_enabled`**
:   is schedule enabled of fast vp

type: boolean


**`emcunity.health.fast.vp.relocation.duration.estimate`**
:   duration estimate of fast vp

type: keyword


**`emcunity.health.fast.vp.relocation.rate`**
:   relocation rate of fast vp

type: long


**`emcunity.health.fast.vp.size.moving.down`**
:   size moving down of fast vp

type: long


**`emcunity.health.fast.vp.size.moving.up`**
:   size moving up of fast vp

type: long


**`emcunity.health.fast.vp.size.moving.up`**
:   size moving up of fast vp

type: long


**`emcunity.health.fast.vp.size.moving.within`**
:   size moving within of fast vp

type: long


**`emcunity.health.fc.port.id`**
:   id of fc port

type: keyword


**`emcunity.health.fc.port.name`**
:   name of fc port

type: keyword


**`emcunity.health.fc.port.health.value`**
:   value of health score

type: long


**`emcunity.health.fc.port.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.fc.port.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.host.container.id`**
:   id of host container

type: keyword


**`emcunity.health.host.container.name`**
:   name of host container

type: keyword


**`emcunity.health.host.container.health.value`**
:   value of health score

type: long


**`emcunity.health.host.container.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.host.container.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.host.initiator.id`**
:   id of host container

type: keyword


**`emcunity.health.host.initiator.health.value`**
:   value of health score

type: long


**`emcunity.health.host.initiator.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.host.initiator.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.host.id`**
:   id of host

type: keyword


**`emcunity.health.host.name`**
:   name of host

type: keyword


**`emcunity.health.host.health.value`**
:   value of health score

type: long


**`emcunity.health.host.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.host.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.io.module.id`**
:   id of io module

type: keyword


**`emcunity.health.io.module.name`**
:   name of io module

type: keyword


**`emcunity.health.io.module.needs_replacement`**
:   boolean needs_replacement of io module

type: boolean


**`emcunity.health.io.module.health.value`**
:   value of health score

type: long


**`emcunity.health.io.module.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.io.module.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.lcc.id`**
:   id of lcc

type: keyword


**`emcunity.health.lcc.name`**
:   name of lcc

type: keyword


**`emcunity.health.lcc.needs_replacement`**
:   boolean needs_replacement of lcc

type: boolean


**`emcunity.health.lcc.health.value`**
:   value of health score

type: long


**`emcunity.health.lcc.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.lcc.health.description.ids`**
:   descriptions of the health score ids

type: keyword


**`emcunity.health.nas.server.id`**
:   id of nas server

type: keyword


**`emcunity.health.nas.server.name`**
:   name of nas server

type: keyword


**`emcunity.health.nas.server.needs_replacement`**
:   boolean needs_replacement of io module

type: boolean


**`emcunity.health.nas.server.health.value`**
:   value of health score

type: long


**`emcunity.health.nas.server.health.descriptions`**
:   descriptions of the value of health score

type: keyword


**`emcunity.health.nas.server.health.description.ids`**
:   descriptions of the health score ids

type: keyword


