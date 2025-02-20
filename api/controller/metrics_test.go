package controller

import (
	"context"
	"fmt"
	"testing"
	"time"

	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

func TestMetrics_ListMetrics(t *testing.T) {

	ctx := context.Background()
	queries := []_interface.QueryFormItem{
		{
			Query: `DCGM_FI_DEV_POWER_USAGE{IBN="算网A", host_ip="10.10.1.84"}`,
			Start: time.Now().Add(-time.Hour).Unix(),
			End:   time.Now().Unix(),
			Step:  15,
		},
	}

	values, err := ListMetrics(ctx, queries)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(values)
}
