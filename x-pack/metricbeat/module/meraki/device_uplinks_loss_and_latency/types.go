// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package device_uplinks_loss_and_latency

import "time"

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

// Uplink contains static device uplink attributes; uplinks are always associated with a device
type Uplink struct {
	DeviceSerial Serial
	IP           string
	Interface    string
	Metrics      []*UplinkMetric
}

// UplinkMetric contains timestamped device uplink metric data points
type UplinkMetric struct {
	Timestamp   time.Time
	LossPercent *float64
	LatencyMs   *float64
}
