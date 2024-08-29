package device_uplinks_cellular_gateway_status

import (
	"fmt"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/x-pack/metricbeat/module/meraki"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"

	meraki_api "github.com/meraki/dashboard-api-go/v3/sdk"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet(meraki.ModuleName, "device_uplinks_cellular_gateway_status", New)
}

type config struct {
	BaseURL       string   `config:"apiBaseURL"`
	ApiKey        string   `config:"apiKey"`
	DebugMode     string   `config:"apiDebugMode"`
	Organizations []string `config:"organizations"`
	// todo: device filtering?
}

func defaultConfig() *config {
	return &config{
		BaseURL:   "https://api.meraki.com",
		DebugMode: "false",
	}
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	logger        *logp.Logger
	client        *meraki_api.Client
	organizations []string
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The meraki device_uplinks_cellular_gateway_status metricset is beta.")

	logger := logp.NewLogger(base.FullyQualifiedName())

	config := defaultConfig()
	if err := base.Module().UnpackConfig(config); err != nil {
		return nil, err
	}

	logger.Debugf("loaded config: %v", config)
	client, err := meraki_api.NewClientWithOptions(config.BaseURL, config.ApiKey, config.DebugMode, "Metricbeat Elastic")
	if err != nil {
		logger.Error("creating Meraki dashboard API client failed: %w", err)
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		logger:        logger,
		client:        client,
		organizations: config.Organizations,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	for _, org := range m.organizations {

		devices, err := meraki.GetDevices(m.client, org)
		if err != nil {
			return err
		}

		val, res, err := m.client.CellularGateway.GetOrganizationCellularGatewayUplinkStatuses(org, &meraki_api.GetOrganizationCellularGatewayUplinkStatusesQueryParams{})
		if err != nil {
			return fmt.Errorf("CellularGateway.GetOrganizationCellularGatewayUplinkStatuses failed; [%d] %s. %w", res.StatusCode(), res.Body(), err)
		}

		fmt.Printf("CellularGateway.GetOrganizationCellularGatewayUplinkStatuses debug; [%d] %s.", res.StatusCode(), res.Body())

		reportApplianceUplinkStatuses(reporter, org, devices, val)
	}

	return nil
}

func reportApplianceUplinkStatuses(reporter mb.ReporterV2, organizationID string, devices map[meraki.Serial]*meraki.Device, responseCellularGatewayUplinkStatuses *meraki_api.ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses) {

	metrics := []mapstr.M{}

	for _, uplink := range *responseCellularGatewayUplinkStatuses {

		if device, ok := devices[meraki.Serial(uplink.Serial)]; ok {
			metric := mapstr.M{
				//this one should be deleted, I just want to see if it matches the device.network_id
				"device.uplink.cellular.gateway.networkd_id":      uplink.NetworkID,
				"device.uplink.cellular.gateway.last_reported_at": uplink.LastReportedAt,
				"device.address":      device.Address,
				"device.firmware":     device.Firmware,
				"device.imei":         device.Imei,
				"device.lan_ip":       device.LanIP,
				"device.location":     device.Location,
				"device.mac":          device.Mac,
				"device.model":        device.Model,
				"device.name":         device.Name,
				"device.network_id":   device.NetworkID,
				"device.notes":        device.Notes,
				"device.product_type": device.ProductType,
				"device.serial":       device.Serial,
				"device.tags":         device.Tags,
			}

			//Not sure if this is really needed on uplink status
			// for k, v := range device.Details {
			// 	metric[fmt.Sprintf("device.details.%s", k)] = v
			// }

			for _, item := range *uplink.Uplinks {
				metrics = append(metrics, mapstr.Union(metric, mapstr.M{
					"device.uplink.cellular.gateway.apn":              item.Apn,
					"device.uplink.cellular.gateway.connection_type":  item.ConnectionType,
					"device.uplink.cellular.gateway.dns1":             item.DNS1,
					"device.uplink.cellular.gateway.dns2":             item.DNS2,
					"device.uplink.cellular.gateway.gateway":          item.Gateway,
					"device.uplink.cellular.gateway.iccid":            item.Iccid,
					"device.uplink.cellular.gateway.interface":        item.Interface,
					"device.uplink.cellular.gateway.ip":               item.IP,
					"device.uplink.cellular.gateway.model":            item.Model,
					"device.uplink.cellular.gateway.provider":         item.Provider,
					"device.uplink.cellular.gateway.public_ip":        item.PublicIP,
					"device.uplink.cellular.gateway.signal_stat.rsrp": item.SignalStat.Rsrp,
					"device.uplink.cellular.gateway.signal_stat.rsrq": item.SignalStat.Rsrq,
					"device.uplink.cellular.gateway.signal_type":      item.SignalType,
					"device.uplink.cellular.gateway.status":           item.Status,
				}))

			}
		}
	}

	// Apn            string                                                                                    `json:"apn,omitempty"`            // Access Point Name
	// ConnectionType string                                                                                    `json:"connectionType,omitempty"` // Connection Type
	// DNS1           string                                                                                    `json:"dns1,omitempty"`           // Primary DNS IP
	// DNS2           string                                                                                    `json:"dns2,omitempty"`           // Secondary DNS IP
	// Gateway        string                                                                                    `json:"gateway,omitempty"`        // Gateway IP
	// Iccid          string                                                                                    `json:"iccid,omitempty"`          // Integrated Circuit Card Identification Number
	// Interface      string                                                                                    `json:"interface,omitempty"`      // Uplink interface
	// IP             string                                                                                    `json:"ip,omitempty"`             // Uplink IP
	// Model          string                                                                                    `json:"model,omitempty"`          // Uplink model
	// Provider       string                                                                                    `json:"provider,omitempty"`       // Network Provider
	// PublicIP       string                                                                                    `json:"publicIp,omitempty"`       // Public IP
	// SignalStat     *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat `json:"signalStat,omitempty"`     // Tower Signal Status
	// SignalType     string                                                                                    `json:"signalType,omitempty"`     // Signal Type
	// Status         string                                                                                    `json:"status,omitempty"`         // Uplink status

	meraki.ReportMetricsForOrganization(reporter, organizationID, metrics)
}
