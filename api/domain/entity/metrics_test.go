package entity

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/xiaoxlm/monitor-gateway/config"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"
	"testing"
	"time"
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
			//Start: 1740471662,
			End: time.Now().Unix(),
			//End:  1740475262,
			Step: 15,
		},
	}

	metrics := NewMetrics(queries, prom)
	metrics.FetchMetrics(context.Background())

	output, err := metrics.GetValues()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, len(output) > 0)
}
