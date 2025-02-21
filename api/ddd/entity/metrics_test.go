package entity

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/xiaoxlm/monitor-gateway/config"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"
)

func TestMetrics_Output(t *testing.T) {
	prom, err := prometheus.NewPrometheus(config.Config.Prom.GetClient())
	if err != nil {
		t.Fatal(err)
	}

	queries := []_interface.QueryFormItem{
		{
			Query: `DCGM_FI_DEV_POWER_USAGE{IBN="算网A", host_ip="10.10.1.84"}`,
			Start: time.Now().Add(-time.Minute * 1).Unix(),
			End:   time.Now().Unix(),
			Step:  15,
		},
	}

	metrics, err := FactoryMetrics(context.Background(), prom, queries)
	if err != nil {
		t.Fatal(err)
	}

	output, err := metrics.Output()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, len(output) > 0)
}
