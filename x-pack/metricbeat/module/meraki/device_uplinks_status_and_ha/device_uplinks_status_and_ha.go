package device_uplinks_status_and_ha

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
	mb.Registry.MustAddMetricSet(meraki.ModuleName, "device_uplinks_status_and_ha", New)
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
	cfgwarn.Beta("The meraki device_uplinks_status_and_ha metricset is beta.")

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

		val, res, err := m.client.Appliance.GetOrganizationApplianceUplinkStatuses(org, &meraki_api.GetOrganizationApplianceUplinkStatusesQueryParams{})
		if err != nil {
			return fmt.Errorf("Appliance.GetOrganizationApplianceUplinkStatuses failed; [%d] %s. %w", res.StatusCode(), res.Body(), err)
		}

		reportApplianceUplinkStatuses(reporter, org, devices, val)
	}

	return nil
}

func reportApplianceUplinkStatuses(reporter mb.ReporterV2, organizationID string, devices map[meraki.Serial]*meraki.Device, responseApplianceUplinkStatuses *meraki_api.ResponseApplianceGetOrganizationApplianceUplinkStatuses) {

	metrics := []mapstr.M{}

	for _, uplink := range *responseApplianceUplinkStatuses {

		if device, ok := devices[meraki.Serial(uplink.Serial)]; ok {
			metric := mapstr.M{
				"device.uplink.high_availablity.enabled": uplink.HighAvailability.Enabled,
				"device.uplink.high_availablity.role":    uplink.HighAvailability.Role,
				"device.uplink.last_reported_at":         uplink.LastReportedAt,
				"device.address":                         device.Address,
				"device.firmware":                        device.Firmware,
				"device.imei":                            device.Imei,
				"device.lan_ip":                          device.LanIP,
				"device.location":                        device.Location,
				"device.mac":                             device.Mac,
				"device.model":                           device.Model,
				"device.name":                            device.Name,
				"device.network_id":                      device.NetworkID,
				"device.notes":                           device.Notes,
				"device.product_type":                    device.ProductType,
				"device.serial":                          device.Serial,
				"device.tags":                            device.Tags,
			}

			//Not sure if this is really needed on uplink status
			// for k, v := range device.Details {
			// 	metric[fmt.Sprintf("device.details.%s", k)] = v
			// }

			for _, item := range *uplink.Uplinks {
				metrics = append(metrics, mapstr.Union(metric, mapstr.M{
					"device.uplink.interface":      item.Interface,
					"device.uplink.status":         item.Status,
					"device.uplink.ip":             item.IP,
					"device.uplink.gateway":        item.Gateway,
					"device.uplink.public_ip":      item.PublicIP,
					"device.uplink.primary_dns":    item.PrimaryDNS,
					"device.uplink.secondary_dns":  item.SecondaryDNS,
					"device.uplink.ip_assigned_by": item.IPAssignedBy,
				}))

			}
		}
	}

	meraki.ReportMetricsForOrganization(reporter, organizationID, metrics)
}
