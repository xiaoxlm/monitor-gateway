package enum

import (
	golang_set "github.com/deckarep/golang-set/v2"
)

type MetricUniqueID string

const (
	MetricUniqueID_Cpu_Avg_Util     MetricUniqueID = "cpu_avg_util"
	MetricUniqueID_Mem_Util         MetricUniqueID = "mem_util"
	MetricUniqueID_Gpu_Mem_Avg_Util MetricUniqueID = "gpu_mem_avg_util"
	MetricUniqueID_Gpu_Avg_Util     MetricUniqueID = "gpu_avg_util"
	MetricUniqueID_Gpu_All_Util     MetricUniqueID = "gpu_all_util"
	MetricUniqueID_Gpu_Avg_Temp     MetricUniqueID = "gpu_avg_temp"
	MetricUniqueID_Disk_Util        MetricUniqueID = "disk_util"
	MetricUniqueID_Eth_Recv         MetricUniqueID = "eth_recv_bytes_rate"
	MetricUniqueID_Eth_Trans        MetricUniqueID = "eth_trans_bytes_rate"
	MetricUniqueID_IB_Recv          MetricUniqueID = "ib_recv_bytes_rate"
	MetricUniqueID_IB_Trans         MetricUniqueID = "ib_trans_bytes_rate"
)

var metricUniqueIDList []MetricUniqueID = []MetricUniqueID{
	MetricUniqueID_Cpu_Avg_Util,
	MetricUniqueID_Mem_Util,
	MetricUniqueID_Gpu_Mem_Avg_Util,
	MetricUniqueID_Gpu_Avg_Util,
	MetricUniqueID_Gpu_All_Util,
	MetricUniqueID_Gpu_Avg_Temp,
	MetricUniqueID_Disk_Util,
	MetricUniqueID_Eth_Recv,
	MetricUniqueID_Eth_Trans,
	MetricUniqueID_IB_Recv,
}

func CheckMetricUniqueIDExist(uniqueID MetricUniqueID) bool {
	list := golang_set.NewSet[MetricUniqueID](metricUniqueIDList...)

	return list.Contains(uniqueID)
}
