package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"github.com/xiaoxlm/monitor-gateway/pkg/util"
	"testing"
)

func TestMetrics_ListMetrics(t *testing.T) {

	ctx := context.Background()

	queries := []request.MetricsQueryInfo{
		{
			MetricUniqueID: `gpu_all_util`,
			LabelValue: map[string]string{
				"IBN":     "算网A",
				"host_ip": "10.10.1.84",
				"start":   "1742194554",
				"end":     "1742194554",
				"step":    "10",
			},
		},
		//{
		//	MetricUniqueID: `mem_util`,
		//	IBN:            "算网A",
		//	HostIP:         "10.10.1.84",
		//	Start:          1740980024, //time.Now().Add(-time.Hour).Unix(),
		//	End:            1740980624, //time.Now().Unix(),
		//	Step:           15,
		//},
	}

	values, err := ListMetrics(ctx, queries)
	if err != nil {
		t.Fatal(err)
	}

	util.LogJSON(values)
}
