package enum

type MetricUniqueID string

const (
	MetricUniqueID_Avg_Cpu_Util MetricUniqueID = "avg_cpu_util"
	MetricUniqueID_Mem_Util     MetricUniqueID = "mem_util"
	MetricUniqueID_Avg_Gpu_Util MetricUniqueID = "avg_gpu_util"
	MetricUniqueID_All_Gpu_Util MetricUniqueID = "all_gpu_util"
	MetricUniqueID_Disk_Util    MetricUniqueID = "disk_util"
	MetricUniqueID_Eth_Recv     MetricUniqueID = "eth_recv_bytes_rate"
	MetricUniqueID_Eth_Trans    MetricUniqueID = "eth_trans_bytes_rate"
	MetricUniqueID_IB_Recv      MetricUniqueID = "ib_recv_bytes_rate"
	MetricUniqueID_IB_Trans     MetricUniqueID = "ib_trans_bytes_rate"
)
