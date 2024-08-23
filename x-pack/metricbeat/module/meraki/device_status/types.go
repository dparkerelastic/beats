// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package device_status

// device unique identifier
type Serial string

// Device contains static device attributes (i.e. dimensions)
type Device struct {
	Address     string
	Details     map[string]string
	Firmware    string
	Imei        *float64
	LanIP       string
	Location    []*float64
	Mac         string
	Model       string
	Name        string
	NetworkID   string
	Notes       string
	ProductType string // one of ["appliance", "camera", "cellularGateway", "secureConnect", "sensor", "switch", "systemsManager", "wireless", "wirelessController"]
	Serial      string
	Tags        []string
}

// DeviceStatus contains dynamic device attributes
type DeviceStatus struct {
	Gateway        string
	IPType         string
	LastReportedAt string
	PrimaryDNS     string
	PublicIP       string
	SecondaryDNS   string
	Status         string // one of ["online", "alerting", "offline", "dormant"]
}
