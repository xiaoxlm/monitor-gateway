package entity

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/xiaoxlm/monitor-gateway/config"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"
	"github.com/xiaoxlm/monitor-gateway/pkg/util"
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
			Start: time.Now().Unix(),
			//Start: 1740471662,
			End: time.Now().Unix(),
			//End:  1740475262,
			Step: 15,
		},
		{
			Query: `100 * (1 - sum by (host_ip,cpu)(increase(node_cpu_seconds_total{mode="idle",IBN="算网A",host_ip="10.10.1.84"}[15s])) / sum by (host_ip,cpu)(increase(node_cpu_seconds_total{IBN="算网A",host_ip="10.10.1.84"}[15s])))`,
			Start: time.Now().Unix(),
			//Start: 1740471662,
			End: time.Now().Unix(),
			//End:  1740475262,
			Step: 15,
		},
	}

	metrics := NewMetrics(queries, prom)
	metrics.FetchMetrics(context.Background())

	output, err := metrics.ListValues()
	if err != nil {
		t.Fatal(err)
	}

	util.LogJSON(output)

	assert.Equal(t, true, len(output) > 0)
}
