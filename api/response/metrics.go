package response

import "github.com/xiaoxlm/monitor-gateway/internal/enum"

type ListMetricsRESP struct {
	Data []MetricsData `json:"data"`
}

type MetricsData struct {
	MetricUniqueID enum.MetricUniqueID `json:"metricUniqueID"` // 指标唯一标识
	Values         []MetricsValues     `json:"values"`         // 时序数值
}

type MetricsValues struct {
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
}
