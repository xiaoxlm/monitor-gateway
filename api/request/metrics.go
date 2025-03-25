package request

type MetricsBatchQueryBody struct {
	Queries []MetricsQueryInfo `json:"queries" binding:"required,dive"`
}

type MetricsQueryInfo struct {
	MetricUniqueID string            `json:"metricUniqueID" binding:"required"`
	LabelValue     map[string]string `json:"labelValue" binding:"required"`
}
