package health

import (
	"time"

	intersight "github.com/CiscoDevNet/intersight-go"
)

type Serial string

type ErrorResponse struct {
	Error struct {
		ErrorCode      int `json:"errorCode"`
		HTTPStatusCode int `json:"httpStatusCode"`
		Messages       []struct {
			EnUS string `json:"en-US"`
		} `json:"messages"`
		Created time.Time `json:"created"`
	} `json:"error"`
}

type IntersightData struct {
	computePhysicalSummaryList   intersight.ComputePhysicalSummaryList
	chassisList                  intersight.EquipmentChassisList
	networkElementList           intersight.NetworkElementList
	assetDeviceRegistrationList  intersight.AssetDeviceRegistrationList
	fanEvent                     []intersight.TelemetryDruidGroupByResult
	memoryEvent                  []intersight.TelemetryDruidGroupByResult
	physicalProcessorEvent       []intersight.TelemetryDruidGroupByResult
	powerSupplyEvent             []intersight.TelemetryDruidGroupByResult
	temperatureEvent             []intersight.TelemetryDruidGroupByResult
	systemCPUEvent               []intersight.TelemetryDruidGroupByResult
	systemMemoryEvent            []intersight.TelemetryDruidGroupByResult
	hostPowerAndStatusEvent      []intersight.TelemetryDruidGroupByResult
	graphicalProcessingUnitEvent []intersight.TelemetryDruidGroupByResult
	signalPowerEvent             []intersight.TelemetryDruidGroupByResult
	electricCurrentEvent         []intersight.TelemetryDruidGroupByResult
	voltageEvent                 []intersight.TelemetryDruidGroupByResult
}
