package request

type ListMetricsMappingQuery struct {
	Category        string `form:"category" binding:"omitempty"`
	MetricsUniqueID string `form:"metricsUniqueID" binding:"omitempty"`
}
