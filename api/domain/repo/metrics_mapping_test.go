package repo

import (
	"context"
	"testing"

	"github.com/xiaoxlm/monitor-gateway/internal/enum"

	"github.com/xiaoxlm/monitor-gateway/pkg/util"

	"github.com/stretchr/testify/assert"
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
)

var db = config.Config.Mysql.GetDB()

func TestMetricsMapping_Create(t *testing.T) {

	ctx := context.Background()
	err := CreateMetricsMapping(ctx, db, &model.MetricsMapping{
		MetricUniqueID: "gpu_util",
		Labels: map[string]interface{}{
			"IBN":     "算网名",
			"host_ip": "节点ip",
		},
		Expression: `DCGM_FI_DEV_GPU_UTIL{IBN="$IBN", host_ip="$host_ip"}`,
		Desc:       "gpu利用率",
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
			Category:   enum.MetrcisMappingCategory_Cpu,
		},
		&model.MetricsMapping{
			MetricUniqueID: "mem_util",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `(1 - (node_memory_MemAvailable_bytes{IBN="$IBN", host_ip="$host_ip"} / node_memory_MemTotal_bytes{IBN="$IBN", host_ip="$host_ip"})) * 100`,
			Desc:       "内存利用率",
			Category:   enum.MetrcisMappingCategory_Memory,
		},
		{
			MetricUniqueID: "gpu_util",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `DCGM_FI_DEV_GPU_UTIL{IBN="$IBN", host_ip="$host_ip"}`,
			Desc:       "gpu利用率",
			Category:   enum.MetrcisMappingCategory_Gpu,
		},
		{
			MetricUniqueID: "disk_util",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `100 - ((node_filesystem_avail_bytes{IBN="$IBN", host_ip="$host_ip",mountpoint="/",fstype!="rootfs"} * 100) / node_filesystem_size_bytes{IBN="$IBN", host_ip="$host_ip",mountpoint="/",fstype!="rootfs"})`,
			Desc:       "磁盘利用率",
			Category:   enum.MetrcisMappingCategory_Disk,
		},
		{
			MetricUniqueID: "eth_recv_bytes_rate",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `sum(rate(node_network_receive_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:       "以太网卡端口接收数据总和的变化率",
			Category:   enum.MetrcisMappingCategory_Network,
		},
		{
			MetricUniqueID: "eth_trans_bytes_rate",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `sum(rate(node_network_transmit_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:       "以太网卡端口发送数据总和的变化率",
			Category:   enum.MetrcisMappingCategory_Network,
		},
		{
			MetricUniqueID: "ib_recv_bytes_rate",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `sum(rate(node_infiniband_port_data_received_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:       "IB网卡端口接收数据总和的变化率",
			Category:   enum.MetrcisMappingCategory_Network,
		},
		{
			MetricUniqueID: "ib_trans_bytes_rate",
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression: `sum(rate(node_infiniband_port_data_transmitted_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:       "IB网卡端口发送数据总和的变化率",
			Category:   enum.MetrcisMappingCategory_Network,
		},
	}

	ctx := context.Background()
	err := BatchCreateMetricsMapping(ctx, db, list)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMetricsMapping_List(t *testing.T) {

	ctx := context.Background()
	datas, err := ListMetricsMapping(ctx, db, "", "")
	if err != nil {
		t.Fatal(err)
	}

	util.LogJSON(datas)
}
