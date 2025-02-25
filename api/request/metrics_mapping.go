package request

import "github.com/xiaoxlm/monitor-gateway/internal/enum"

type ListMetricsMappingQuery struct {
	Category        enum.MetrcisMappingCategory `form:"category" binding:"omitempty"`
	MetricsUniqueID string                      `form:"metricsUniqueID" binding:"omitempty"`
}
