package response

import (
	"github.com/lie-flat-planet/httputil"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
)

type ListMetricsRESP struct {
	Data []MetricsData `json:"data"`
}

type MetricsData struct {
	MetricUniqueID   enum.MetricUniqueID      `json:"metricUniqueID"` // 指标唯一标识
	HostIP           string                   `json:"hostIP"`
	Label            map[string]string        `json:"label"`            // 暂存字段。将请求参数的label透传给 ibn 用
	MultiMetricsData bool                     `json:"multiMetricsData"` // values 是否有多个值
	Values           httputil.MetricsFromExpr `json:"values"`           // 时序数值
}

type MetricsValues struct {
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
}
