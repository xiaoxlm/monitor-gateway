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
		MetricUniqueID: enum.MetricUniqueID_Gpu_Avg_Util,
		Labels: map[string]interface{}{
			"IBN":     "算网名",
			"host_ip": "节点ip",
		},
		Expression:     `avg by(host_ip) (DCGM_FI_DEV_GPU_UTIL{IBN="$IBN", host_ip="$host_ip"})`,
		Desc:           "gpu利用率",
		Category:       enum.MetrcisMappingCategory_Gpu,
		BoardPayloadID: 2,
		PanelID:        "7384ef77-1d5c-4133-a4ce-a496a4fa5114",
	})

	assert.Equal(t, true, err == nil)
}

func TestMetricsMapping_BatchCreate(t *testing.T) {
	list := []*model.MetricsMapping{
		&model.MetricsMapping{
			MetricUniqueID: enum.MetricUniqueID_Cpu_Avg_Util,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `100 * (1 - sum by (host_ip)(increase(node_cpu_seconds_total{mode="idle",IBN="$IBN",host_ip="$host_ip"}[30s])) / sum by (host_ip)(increase(node_cpu_seconds_total{IBN="$IBN",host_ip="$host_ip"}[30s])))`,
			Desc:           "cpu利用率",
			Category:       enum.MetrcisMappingCategory_Cpu,
			BoardPayloadID: 2,
			PanelID:        "b93d912c-b3bf-4c6e-a77b-77ea6123ffe9",
		},
		&model.MetricsMapping{
			MetricUniqueID: enum.MetricUniqueID_Mem_Util,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `(1 - (node_memory_MemAvailable_bytes{IBN="$IBN", host_ip="$host_ip"} / node_memory_MemTotal_bytes{IBN="$IBN", host_ip="$host_ip"})) * 100`,
			Desc:           "内存利用率",
			Category:       enum.MetrcisMappingCategory_Memory,
			BoardPayloadID: 2,
			PanelID:        "34ef9aa4-0328-41a2-aa32-265604e2abfb",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_Gpu_Avg_Util,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `avg by(host_ip) (DCGM_FI_DEV_GPU_UTIL{IBN="$IBN", host_ip="$host_ip"})`,
			Desc:           "gpu平均利用率",
			Category:       enum.MetrcisMappingCategory_Gpu,
			BoardPayloadID: 2,
			PanelID:        "7384ef77-1d5c-4133-a4ce-a496a4fa5114",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_Gpu_All_Util,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `avg by(host_ip,gpu) (DCGM_FI_DEV_GPU_UTIL{IBN="$IBN", host_ip="$host_ip"})`,
			Desc:           "所有gpu利用率",
			Category:       enum.MetrcisMappingCategory_Gpu,
			BoardPayloadID: 2,
			PanelID:        "7384ef77-1d5c-4133-a4ce-a496a4fa5114",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_Gpu_Mem_Avg_Util,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `avg by (host_ip)(last_over_time(DCGM_FI_DEV_FB_USED{IBN="$IBN", host_ip="$host_ip"}[1m])/(last_over_time(DCGM_FI_DEV_FB_FREE{IBN="$IBN", host_ip="$host_ip"}[1m]) + last_over_time(DCGM_FI_DEV_FB_USED{IBN="$IBN", host_ip="$host_ip"}[1m])))`,
			Desc:           "gpu平均内存利用率",
			Category:       enum.MetrcisMappingCategory_Gpu,
			BoardPayloadID: 2,
			PanelID:        "4f046feb-4caf-4174-9c10-9b7c6a4ee795",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_Gpu_Avg_Temp,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `avg by (host_ip)DCGM_FI_DEV_GPU_TEMP{IBN="$IBN", host_ip="$host_ip"})`,
			Desc:           "gpu平均温度",
			Category:       enum.MetrcisMappingCategory_Gpu,
			BoardPayloadID: 2,
			PanelID:        "2b51fdb8-6a56-481c-91f0-8d4c09cccc14",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_Disk_Util,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `100 - ((node_filesystem_avail_bytes{IBN="$IBN", host_ip="$host_ip",mountpoint="/",fstype!="rootfs"} * 100) / node_filesystem_size_bytes{IBN="$IBN", host_ip="$host_ip",mountpoint="/",fstype!="rootfs"})`,
			Desc:           "磁盘利用率",
			Category:       enum.MetrcisMappingCategory_Disk,
			BoardPayloadID: 2,
			PanelID:        "4c2a360d-ef1c-4268-bc18-c8cb5aab2744",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_Eth_Recv,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `sum(rate(node_network_receive_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:           "以太网卡端口接收数据总和的变化率",
			Category:       enum.MetrcisMappingCategory_Network,
			BoardPayloadID: 2,
			PanelID:        "7f075e4e-c4c2-43a6-a099-d76cc93fd8f4",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_Eth_Trans,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `sum(rate(node_network_transmit_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:           "以太网卡端口发送数据总和的变化率",
			Category:       enum.MetrcisMappingCategory_Network,
			BoardPayloadID: 2,
			PanelID:        "7f075e4e-c4c2-43a6-a099-d76cc93fd8f4",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_IB_Recv,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `sum(rate(node_infiniband_port_data_received_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:           "IB网卡端口接收数据总和的变化率",
			Category:       enum.MetrcisMappingCategory_Network,
			BoardPayloadID: 2,
			PanelID:        "10762210-8b41-447b-be50-52673f1091f5",
		},
		{
			MetricUniqueID: enum.MetricUniqueID_IB_Trans,
			Labels: map[string]interface{}{
				"IBN":     "算网名",
				"host_ip": "节点ip",
			},
			Expression:     `sum(rate(node_infiniband_port_data_transmitted_bytes_total{IBN="$IBN", host_ip="$host_ip"}[1m]))`,
			Desc:           "IB网卡端口发送数据总和的变化率",
			Category:       enum.MetrcisMappingCategory_Network,
			BoardPayloadID: 2,
			PanelID:        "10762210-8b41-447b-be50-52673f1091f5",
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
