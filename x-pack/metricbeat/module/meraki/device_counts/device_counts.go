package device_counts

import (
	"encoding/json"
	"log"

	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/helper"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/elastic/beats/v7/metricbeat/mb/parse"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

var (
	hostParser = parse.URLHostParserBuilder{
		DefaultScheme: "http",
		DefaultPath:   "/api/v1/organizations",
	}.Build()
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host is defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("meraki", "device_counts", New,
		mb.WithHostParser(hostParser),
		mb.DefaultMetricSet())
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	http *helper.HTTP
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The meraki device_counts metricset is beta.")

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	http, err := helper.NewHTTP(base)
	if err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		http:          http,
		//counter:       1,
	}, nil
}

// Fetch method implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {

	//m.http.FetchContent()
	orgIDs := GetOrganizationIDList(m.http)
	for _, orgIdValue := range orgIDs {

		if orgIdValue.API_status.Enabled {
			orgIDDeviceCounts := GetDeviceCountsByOrgId(orgIdValue.Id, m.http)
			report.Event(mb.Event{
				MetricSetFields: mapstr.M{
					//"counter":  m.counter,
					"id":       orgIdValue.Id,
					"online":   orgIDDeviceCounts.OrgCounts.OrgByStatus.Online,
					"alerting": orgIDDeviceCounts.OrgCounts.OrgByStatus.Alerting,
					"offline":  orgIDDeviceCounts.OrgCounts.OrgByStatus.Offline,
					"dormant":  orgIDDeviceCounts.OrgCounts.OrgByStatus.Dormant,
				},
			})
		} else {
			log.Println("Meracki Organization ID " + orgIdValue.Id + " api are disabled")
		}

	}

	return nil
}

func GetDeviceCountsByOrgId(id string, httpClient *helper.HTTP) ResponseDeviceCounts {

	initialURL := httpClient.GetURI()
	httpClient.SetURI(httpClient.GetURI() + "/" + id + "/devices/statuses/overview")

	responseData, err := httpClient.FetchContent()
	if err != nil {
		httpClient.SetURI(initialURL)
		log.Fatal(err)
	}

	httpClient.SetURI(initialURL)
	//fmt.Println(string(responseData))

	var responseObject ResponseDeviceCounts
	json.Unmarshal(responseData, &responseObject)

	return responseObject

}

type ResponseDeviceCounts struct {
	OrgCounts Counts `json:"counts"`
}

type Counts struct {
	OrgByStatus ByStatus `json:"byStatus"`
}

type ByStatus struct {
	Online   int `json:"online"`
	Alerting int `json:"alerting"`
	Offline  int `json:"offline"`
	Dormant  int `json:"dormant"`
}

func GetOrganizationIDList(httpClient *helper.HTTP) []Org_Response {

	responseData, err := httpClient.FetchContent()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))

	var responseObject []Org_Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject

}

// A Response struct to map the Entire Response
type Org_Response struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Url        string      `json:"url"`
	API_status API_enabled `json:"api"`
}

type API_enabled struct {
	Enabled bool `json:"enabled"`
}
