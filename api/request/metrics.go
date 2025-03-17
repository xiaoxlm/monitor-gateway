package request

import "github.com/xiaoxlm/monitor-gateway/internal/enum"

type MetricsBatchQueryBody struct {
	Queries []MetricsQueryInfo `json:"queries" binding:"required,dive"`
}

type MetricsQueryInfo struct {
	MetricUniqueID enum.MetricUniqueID `json:"metricUniqueID" binding:"required"`
	LabelValue     map[string]string   `json:"labelValue" binding:"required"`
}
