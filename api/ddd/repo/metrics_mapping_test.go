package repo

import (
	"context"
	"testing"

	"github.com/xiaoxlm/monitor-gateway/pkg/util"

	"github.com/stretchr/testify/assert"
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
)

var db = config.Config.Mysql.GetDB()

func TestMetricsMapping_Create(t *testing.T) {

	ctx := context.Background()
	err := CreateMetricsMapping(ctx, db, &model.MetricsMapping{
		MetricUniqueID: "mem_util",
		Labels: map[string]interface{}{
			"IBN":     "算网名",
			"host_ip": "节点ip",
		},
		Expression: `(1 - (node_memory_MemAvailable_bytes{IBN="$IBN", host_ip="$host_ip"} / node_memory_MemTotal_bytes{IBN="$IBN", host_ip="$host_ip"})) * 100`,
		Desc:       "内存利用率",
	})

	assert.Equal(t, true, err == nil)
}

func TestMetricsMapping_BatchCreate(t *testing.T) {
	list := []*model.MetricsMapping{
		&model.MetricsMapping{
			MetricUniqueID: "cpu_util",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `(1 - avg(rate(node_cpu_seconds_total{IBN="$IBN", mode="idle", host_ip="$host_ip"}[1m]))) * 100`,
			Desc:       "cpu利用率",
		},
		&model.MetricsMapping{
			MetricUniqueID: "mem_util",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `(1 - (node_memory_MemAvailable_bytes{IBN="$IBN", host_ip="$host_ip"} / node_memory_MemTotal_bytes{IBN="$IBN", host_ip="$host_ip"})) * 100`,
			Desc:       "内存利用率",
		},
	}

	ctx := context.Background()
	err := BatchCreateMetricsMapping(ctx, db, list)

	assert.Equal(t, true, err == nil)
}

func TestMetricsMapping_List(t *testing.T) {

	ctx := context.Background()
	datas, err := ListMetricsMapping(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	util.LogJSON(datas)
}
