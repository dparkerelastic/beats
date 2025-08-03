package health

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/logp"

	intersight "github.com/CiscoDevNet/intersight-go"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("cisco_intersight", "health", New)
}

type config struct {
	// Hosts     []string      `config:"hosts"`
	IntersightHost       string        `config:"intersight_host"`
	IntersightAPIKeyID   string        `config:"intersight_api_key_id"`
	IntersightPrivateKey string        `config:"intersight_private_key"`
	Period               time.Duration `config:"period"`
	DebugMode            bool          `config:"debug"`
	UserName             string        `config:"username"`
	Password             string        `config:"password"`
	PerPage              int           `config:"per_page"`
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	//config  *config
	logger               *logp.Logger
	counter              int
	debug                bool
	intersightHost       string
	intersightAPIKeyID   string
	intersightPrivateKey string
	perPage              int
	period               time.Duration
	previousTime         time.Time // Timestamp of the previous fetch.
}

// type MetricSet struct {
// 	mb.BaseMetricSet
// 	counter int
// }

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The cisco_intersight health metricset is beta.")

	//config := struct{}{}
	//config := defaultConfig()
	config := &config{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}
	logger := logp.NewLogger(base.FullyQualifiedName())

	return &MetricSet{
		BaseMetricSet:        base,
		counter:              1,
		debug:                config.DebugMode,
		logger:               logger,
		perPage:              config.PerPage,
		intersightHost:       config.IntersightHost,
		intersightAPIKeyID:   config.IntersightAPIKeyID,
		intersightPrivateKey: config.IntersightPrivateKey,
		period:               config.Period,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error(to).
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {

	currentTime := time.Now().Truncate(time.Minute)
	// fmt.Println("Current time:", currentTime.Format(time.RFC3339))
	// fmt.Println("Code is Running")

	if m.previousTime.IsZero() {
		m.previousTime = time.Now().Truncate(time.Minute).Add(-m.period)
	}

	//Setup Connection Info for this Fetch
	hostInfo := ConnectionInfo{}
	hostInfo.Host = m.intersightHost
	hostInfo.ApiKeyID = m.intersightAPIKeyID
	hostInfo.PrivateKey = m.intersightPrivateKey

	// Configuration for Intersight API
	config := intersight.NewConfiguration()
	config.Host = "intersight.com" // Default Intersight SaaS endpoint

	authConfig := intersight.HttpSignatureAuth{
		KeyId:            m.intersightAPIKeyID,
		PrivateKeyReader: strings.NewReader(m.intersightPrivateKey),
		SigningScheme:    intersight.HttpSigningSchemeHs2019, //WORKS!!!
		HashAlgorithm:    intersight.HttpHashAlgorithmSha256, // was HttpSigningAlgorithmRsaPSS

		SignedHeaders: []string{
			intersight.HttpSignatureParameterRequestTarget, // The special (request-target) parameter expresses the HTTP request target.
			intersight.HttpSignatureParameterCreated,       // Time when request was signed, formatted as a Unix timestamp integer value.
			"Host",                                         // The Host request header specifies the domain name of the server, and optionally the TCP port number.
			"Date",                                         // The date and time at which the message was originated.
			//"Content-Type",                                 // The Media type of the body of the request.
			"Digest", // A cryptographic digest of the request body.
		},
		SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPKCS1v15, // Other value
	}

	authCtx, err := authConfig.ContextWithValue(context.Background())
	if err != nil {
		m.logger.Warnf("GetSystemMetrics failed; %v", err)
	}
	client := intersight.NewAPIClient(config)

	var metricData IntersightData

	// Get the list of physical summaries by MOID
	physicalSummariesList, resp, err := client.ComputeApi.GetComputePhysicalSummaryList(authCtx).Execute()
	if err != nil {
		m.logger.Warnf("GetOrganizationOrganizationList failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)
	} else {
		metricData.computePhysicalSummaryList = *physicalSummariesList.ComputePhysicalSummaryList

	}

	// Get the list of equipment chassis by MOID
	chassisList, resp, err := client.EquipmentApi.GetEquipmentChassisList(authCtx).Execute()
	if err != nil {
		m.logger.Warnf("GetOrganizationOrganizationList failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)
	} else {
		metricData.chassisList = *chassisList.EquipmentChassisList
	}

	// Get the list of elements by MOID
	elementsList, resp, err := client.NetworkApi.GetNetworkElementList(authCtx).Execute()
	if err != nil {
		m.logger.Warnf("GetOrganizationOrganizationList failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)
	} else {
		metricData.networkElementList = *elementsList.NetworkElementList
	}

	// Get the list of devices registertion by MOID
	deviceRegistrationList, resp, err := client.AssetApi.GetAssetDeviceRegistrationList(authCtx).Execute()
	if err != nil {
		m.logger.Warnf("GetOrganizationOrganizationList failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)
	} else {
		metricData.assetDeviceRegistrationList = *deviceRegistrationList.AssetDeviceRegistrationList
	}

	// Execute the telemetry query using the API client
	responseData, resp, err := client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(FanTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("FanTelemetryDruidGroupByRequest failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.fanEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(MemoryTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("VoltgageTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {

			metricData.memoryEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(PhysicalProcessorTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("VoltgageTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.physicalProcessorEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(PowerSupplyTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("PowerSupplyTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.powerSupplyEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(TemperatureTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("TemperatureTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.temperatureEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(SystemCPUTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("TemperatureTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.systemCPUEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(SystemMemoryTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("TemperatureTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {

		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.systemMemoryEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(HostPowerAndStatusTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("HostPowerAndStatusTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.hostPowerAndStatusEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}

	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(GraphicalProcessingUnitTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("HostPowerAndStatusTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.graphicalProcessingUnitEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(SignalPowerTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("HostPowerAndStatusTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {

		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.signalPowerEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(ElectricCurrentTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("HostPowerAndStatusTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.electricCurrentEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	// Execute the telemetry query using the API client
	responseData, resp, err = client.TelemetryApi.QueryTelemetryGroupBy(authCtx).TelemetryDruidGroupByRequest(VoltageTelemetryDruidGroupByRequestStruct(currentTime, m.previousTime)).Execute()
	if err != nil {
		m.logger.Warnf("HostPowerAndStatusTelemetryDruidGroupByRequestStruct failed; %v", err)
	} else if resp.StatusCode != 200 {
		m.logger.Warnf("Unexpected response status: %d", resp.StatusCode)

	} else {
		if _, err := json.MarshalIndent(responseData, "", "  "); err == nil {
			metricData.voltageEvent = responseData
		} else {
			m.logger.Warnf("Failed to marshal responseData to JSON: %v", err)
		}
	}

	reportMetrics(reporter, hostInfo.Host, metricData, m.debug)

	// fmt.Println("Previous time:", m.previousTime.Format(time.RFC3339))
	// fmt.Println("Period:", m.period)
	// fmt.Println("Current time:", currentTime.Format(time.RFC3339))

	m.previousTime = currentTime // Update the previous time to the current time for the next fetch.

	return nil
}

type ConnectionInfo struct {
	Host       string
	ApiKeyID   string
	PrivateKey string
}
