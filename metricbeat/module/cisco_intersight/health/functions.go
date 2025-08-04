package health

import (
	"time"

	intersight "github.com/CiscoDevNet/intersight-go"
)

// Build the TelemetryDruidGroupByRequest struct
// To get extra attributes, add them as dimensions in the request.
// For example, to get "hw.fan.serial" and "hw.fan.model", add them to Dimensions.

func FanTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		//Granularity: "all",
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		// },
		// Granularity: intersight.TelemetryDruidGranularity{
		// 	TelemetryDruidDurationGranularity: &intersight.TelemetryDruidDurationGranularity{
		// 		Type:     "duration",
		// 		Duration: int64(0),
		// 	},
		// },
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			// {
			// 	TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
			// 		Type:       "default",
			// 		Dimension:  "id",
			// 		OutputName: "unit.id",
			// 	},
			// },
			// {
			// 	TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
			// 		Type:       "default",
			// 		Dimension:  "id",
			// 		OutputName: "unit.id",
			// 	},
			// },
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			// {
			// 	TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
			// 		Type:       "default",
			// 		Dimension:  "intersight.asset.cluster_member.moid",
			// 		OutputName: "intersight.asset.cluster_member.moid",
			// 	},
			// },
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.fan",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.fan.speed_max",
					FieldName: "hw.fan.speed_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.fan.speed_ratio_max",
					FieldName: "hw.fan.speed_ratio_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.fan.speed.limit_low_critical_max",
					FieldName: "hw.fan.speed.limit_low_critical_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.fan.speed.limit_low_degraded_max",
					FieldName: "hw.fan.speed.limit_low_degraded_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.fan.speed.limit_max_max",
					FieldName: "hw.fan.speed.limit_max_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.status_min",
					FieldName: "hw.status_min",
				},
			},
		},
	}
}

func MemoryTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.model",
					OutputName: "parent.model",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.serial_number",
					OutputName: "parent.serial_number",
				},
			},
			// {
			// 	TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
			// 		Type:       "default",
			// 		Dimension:  "serial_number",
			// 		OutputName: "serial_number",
			// 	},
			// },
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.memory",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.memory.size_max",
					FieldName: "hw.memory.size_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.memory.size_min",
					FieldName: "hw.memory.size_min",
				},
			},
			{
				TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
					Type:      "longSum",
					Name:      "hw.memory.size_count",
					FieldName: "hw.memory.size_count",
				},
			},
			{
				TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
					Type:      "longSum",
					Name:      "hw.memory.size",
					FieldName: "hw.memory.size",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.status_min",
					FieldName: "hw.status_min",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.errors_uncorrectable_ecc_errors_max",
					FieldName: "hw.errors_uncorrectable_ecc_errors_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.errors_correctable_ecc_errors_max",
					FieldName: "hw.errors_correctable_ecc_errors_max",
				},
			},
			// {
			// 	TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
			// 		Type:      "longMax",
			// 		Name:      "hw.memory.size_max",
			// 		FieldName: "hw.memory.size_max",
			// 	},
			// },
		},
		PostAggregations: []intersight.TelemetryDruidPostAggregator{
			{
				TelemetryDruidExpressionPostAggregator: &intersight.TelemetryDruidExpressionPostAggregator{
					Type:       "expression",
					Name:       func() *string { s := "hw.memory.size_avg"; return &s }(),
					Expression: func() *string { s := "(\"hw.memory.size\" / \"hw.memory.size_count\")"; return &s }(),
				},
			},
		},
	}

}

func PhysicalProcessorTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
			// {
			// 	TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
			// 		Type:       "default",
			// 		Dimension:  "parent.model",
			// 		OutputName: "parent.model",
			// 	},
			// },
			// {
			// 	TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
			// 		Type:       "default",
			// 		Dimension:  "parent.serial_number",
			// 		OutputName: "parent.serial_number",
			// 	},
			// },
			// {
			// 	TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
			// 		Type:       "default",
			// 		Dimension:  "serial_number",
			// 		OutputName: "serial_number",
			// 	},
			// },
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.cpu",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.cpu.utilization_c0_max",
					FieldName: "hw.cpu.utilization_c0_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.cpu.utilization_c0_min",
					FieldName: "hw.cpu.utilization_c0_min",
				},
			},
			{
				TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
					Type:      "longSum",
					Name:      "hw.cpu.utilization_c0_count",
					FieldName: "hw.cpu.utilization_c0_count",
				},
			},
			{
				TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
					Type:      "longSum",
					Name:      "hw.cpu.utilization_c0",
					FieldName: "hw.cpu.utilization_c0",
				},
			},
		},
		PostAggregations: []intersight.TelemetryDruidPostAggregator{
			{
				TelemetryDruidExpressionPostAggregator: &intersight.TelemetryDruidExpressionPostAggregator{
					Type:       "expression",
					Name:       func() *string { s := "hw.cpu.utilization_c0_avg"; return &s }(),
					Expression: func() *string { s := "(\"hw.cpu.utilization_c0\" / \"hw.cpu.utilization_c0_count\")"; return &s }(),
				},
			},
		},
	}

}

func PowerSupplyTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.power_supply",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.power_supply.utilization_max",
					FieldName: "hw.power_supply.utilization_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.power_max",
					FieldName: "hw.power_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.power_out_max",
					FieldName: "hw.power_out_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.power_supply.limit_max_max",
					FieldName: "hw.power_supply.limit_max_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.status_min",
					FieldName: "hw.status_min",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.power_supply.limit_throttled_max",
					FieldName: "hw.power_supply.limit_throttled_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.power_supply.limit_high_critical_max",
					FieldName: "hw.power_supply.limit_high_critical_max",
				},
			},
			// {
			// 	TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
			// 		Type:      "longMin",
			// 		Name:      "hw.cpu.utilization_c0_min",
			// 		FieldName: "hw.cpu.utilization_c0_min",
			// 	},
			// },
			// {
			// 	TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
			// 		Type:      "longSum",
			// 		Name:      "hw.cpu.utilization_c0_count",
			// 		FieldName: "hw.cpu.utilization_c0_count",
			// 	},
			// },
			// {
			// 	TelemetryDruidSumAggregator: &intersight.TelemetryDruidSumAggregator{
			// 		Type:      "longSum",
			// 		Name:      "hw.cpu.utilization_c0",
			// 		FieldName: "hw.cpu.utilization_c0",
			// 	},
			// },
		},
		// PostAggregations: []intersight.TelemetryDruidPostAggregator{
		// 	{
		// 		TelemetryDruidExpressionPostAggregator: &intersight.TelemetryDruidExpressionPostAggregator{
		// 			Type:       "expression",
		// 			Name:       func() *string { s := "hw.cpu.utilization_c0_avg"; return &s }(),
		// 			Expression: func() *string { s := "(\"hw.cpu.utilization_c0\" / \"hw.cpu.utilization_c0_count\")"; return &s }(),
		// 		},
		// 	},
		// },
	}

}

func TemperatureTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.temperature",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.temperature_max",
					FieldName: "hw.temperature_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.temperature_min",
					FieldName: "hw.temperature_min",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.status_min",
					FieldName: "hw.status_min",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.temperature.limit_low_critical_min",
					FieldName: "hw.temperature.limit_low_critical_min",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.temperature.limit_high_critical_max",
					FieldName: "hw.temperature.limit_high_critical_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.temperature.limit_high_degraded_max",
					FieldName: "hw.temperature.limit_high_degraded_max",
				},
			},
		},
	}
}

func SystemCPUTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "system.cpu",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "system.cpu.utilization_system_max",
					FieldName: "system.cpu.utilization_system_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "system.cpu.utilization_idle_max",
					FieldName: "system.cpu.utilization_idle_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "system.cpu.utilization_user_max",
					FieldName: "system.cpu.utilization_user_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMin",
					Name:      "system.cpu.utilization_system_min",
					FieldName: "system.cpu.utilization_system_min",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMin",
					Name:      "system.cpu.utilization_idle_min",
					FieldName: "system.cpu.utilization_idle_min",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMin",
					Name:      "system.cpu.utilization_user_min",
					FieldName: "system.cpu.utilization_user_min",
				},
			},
		},
	}
}

func SystemMemoryTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "system.memory",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "system.memory.usage_free_max",
					FieldName: "system.memory.usage_free_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "system.memory.usage_cached_max",
					FieldName: "system.memory.usage_cached_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "system.memory.utilization_max",
					FieldName: "system.memory.utilization_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "system.memory.usage_used_max",
					FieldName: "system.memory.usage_used_max",
				},
			},
		},
	}
}

func HostPowerAndStatusTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.host",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "hw.host.power_max",
					FieldName: "hw.host.power_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMin",
					Name:      "hw.host.power_state_min",
					FieldName: "hw.host.power_state_min",
				},
			},
		},
	}
}

func GraphicalProcessingUnitTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.gpu",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.gpu.clockspeed_max",
					FieldName: "hw.gpu.clockspeed_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "longMax",
					Name:      "hw.gpu.memory.clockspeed_max",
					FieldName: "hw.gpu.memory.clockspeed_max",
				},
			},
		},
	}
}

func SignalPowerTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.signal_power",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "hw.signal_power_receive_max",
					FieldName: "hw.signal_power_receive_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "hw.signal_power_transmit_max",
					FieldName: "hw.signal_power_transmit_max",
				},
			},
		},
	}
}

func ElectricCurrentTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.current",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "hw.current_max",
					FieldName: "hw.current_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMin",
					Name:      "hw.current_min",
					FieldName: "hw.current_min",
				},
			},
		},
	}
}

func VoltageTelemetryDruidGroupByRequestStruct(currentTime time.Time, previousTime time.Time) intersight.TelemetryDruidGroupByRequest {

	return intersight.TelemetryDruidGroupByRequest{
		QueryType: "groupBy",
		DataSource: intersight.TelemetryDruidDataSource{
			TelemetryDruidTableDataSource: &intersight.TelemetryDruidTableDataSource{
				Type: "table",
				Name: "PhysicalEntities",
			},
		},
		Granularity: intersight.TelemetryDruidGranularity{
			TelemetryDruidPeriodGranularity: &intersight.TelemetryDruidPeriodGranularity{
				Type:   "all",
				Period: "all",
			},
		},
		Intervals: []string{
			previousTime.UTC().Format("2006-01-02T15:04:05.000Z") + "/" +
				currentTime.UTC().Format("2006-01-02T15:04:05.000Z"),
		},
		Dimensions: []intersight.TelemetryDruidDimensionSpec{
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.name",
					OutputName: "host.name",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.type",
					OutputName: "host.type",
				},
			},
			//
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "host.id",
					OutputName: "host.id",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "intersight.asset.device_registration.moid",
					OutputName: "intersight.asset.device_registration.moid",
				},
			},
			{
				TelemetryDruidDefaultDimensionSpec: &intersight.TelemetryDruidDefaultDimensionSpec{
					Type:       "default",
					Dimension:  "parent.id",
					OutputName: "parent.id",
				},
			},
		},
		Filter: &intersight.TelemetryDruidFilter{
			TelemetryDruidAndFilter: &intersight.TelemetryDruidAndFilter{
				Type: "and",
				Fields: []intersight.TelemetryDruidFilter{
					{
						TelemetryDruidSelectorFilter: &intersight.TelemetryDruidSelectorFilter{
							Type:      "selector",
							Dimension: "instrument.name",
							Value:     "hw.voltage",
						},
					},
				},
			},
		},
		Aggregations: []intersight.TelemetryDruidAggregator{
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMax",
					Name:      "hw.voltage_max",
					FieldName: "hw.voltage_max",
				},
			},
			{
				TelemetryDruidMinMaxAggregator: &intersight.TelemetryDruidMinMaxAggregator{
					Type:      "floatMin",
					Name:      "hw.voltage_min",
					FieldName: "hw.voltage_min",
				},
			},
		},
	}
}
