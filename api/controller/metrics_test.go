package controller

import (
	"context"
	"fmt"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"testing"
)

func TestMetrics_ListMetrics(t *testing.T) {

	ctx := context.Background()

	queries := []request.MetricsQueryInfo{
		{
			MetricUniqueID: `cpu_util`,
			IBN:            "算网A",
			HostIP:         "10.10.1.84",
			Start:          1740980024, //time.Now().Add(-time.Hour).Unix(),
			End:            1740980624, //time.Now().Unix(),
			Step:           15,
		},
		{
			MetricUniqueID: `mem_util`,
			IBN:            "算网A",
			HostIP:         "10.10.1.84",
			Start:          1740980024, //time.Now().Add(-time.Hour).Unix(),
			End:            1740980624, //time.Now().Unix(),
			Step:           15,
		},
	}

	values, err := ListMetrics(ctx, queries)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(values)
}
