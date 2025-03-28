package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"github.com/xiaoxlm/monitor-gateway/pkg/util"
	"testing"
)

func TestAllMapping(t *testing.T) {
	ctx := context.Background()

	mappingList, err := ListMetricsMapping(ctx, &request.ListMetricsMappingQuery{})
	if err != nil {
		t.Fatal(err)
	}

	for _, mapping := range mappingList {
		metricUniqueID := mapping.MetricUniqueID

		q := []request.MetricsQueryInfo{
			{
				MetricUniqueID: string(metricUniqueID),
				LabelValue: map[string]string{
					"IBN":     "算网A",
					"host_ip": "10.10.1.85",
					"end":     "1742803857",
					"start":   "1742803857",
					"step":    "10",
				},
			},
		}

		values, err := ListMetrics(ctx, q)
		if err != nil {
			t.Fatal("metricUniqueID:", metricUniqueID, err)
		}

		util.LogJSON(values)
	}
}

func TestMetrics_ListMetrics(t *testing.T) {
	ctx := context.Background()

	queries := []request.MetricsQueryInfo{
		{
			MetricUniqueID: `cpu_avg_util`,
			LabelValue: map[string]string{
				"IBN":     "算网A",
				"host_ip": "10.10.1.85",
				"end":     "1742803857",
				"start":   "1742803857",
				"step":    "10",
			},
		},
		{
			MetricUniqueID: `ib_trans_bytes_rate`,
			LabelValue: map[string]string{
				"IBN":     "算网A",
				"host_ip": "10.10.1.84",
				"start":   "1642803857",
				"end":     "1642803857",
				"step":    "10",
			},
		},
	}

	values, err := ListMetrics(ctx, queries)
	if err != nil {
		t.Fatal(err)
	}

	util.LogJSON(values)
}
