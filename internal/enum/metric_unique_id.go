package enum

type MetricUniqueID string

const (
	MetricUniqueID_Cpu_Util  MetricUniqueID = "cpu_util"
	MetricUniqueID_Mem_Util  MetricUniqueID = "mem_util"
	MetricUniqueID_Disk_Util MetricUniqueID = "disk_util"
	MetricUniqueID_Eth_Recv  MetricUniqueID = "eth_recv_bytes_rate"
	MetricUniqueID_Eth_Trans MetricUniqueID = "eth_trans_bytes_rate"
	MetricUniqueID_IB_Recv   MetricUniqueID = "ib_recv_bytes_rate"
	MetricUniqueID_IB_Trans  MetricUniqueID = "ib_trans_bytes_rate"
)
