package enum

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
