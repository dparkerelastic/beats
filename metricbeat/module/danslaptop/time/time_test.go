//go:build darwin || freebsd || linux || openbsd || windows || aix

package time

import (
	"testing"

	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
	_ "github.com/elastic/beats/v7/metricbeat/module/system"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	f := mbtest.NewReportingMetricSetV2Error(t, getConfig())
	events, errs := mbtest.ReportingFetchV2Error(f)

	assert.Empty(t, errs)
	// if !assert.NotEmpty(t, events) {
	// 	t.FailNow()
	// }

	//var counter int= events[0].BeatEvent("danslaptop", "time").
	//var counter int = events[0].BeatEvent("danslaptop", "time")

	if !assert.NotEmpty(t, events) {
		t.FailNow()
	}

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(),
		events[0].BeatEvent("danslaptop", "time").Fields.StringToPrint())
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "danslaptop",
		"metricsets": []string{"time"},
	}
}
