package request

import "github.com/xiaoxlm/monitor-gateway/internal/enum"

type MetricsBatchQueryBody struct {
	Queries []MetricsQueryInfo `json:"queries" binding:"required,dive"`
}

type MetricsQueryInfo struct {
	MetricUniqueID enum.MetricUniqueID `json:"metricUniqueID" binding:"required"`
	IBN            string              `json:"ibn" binding:"required"`    // 算网名
	HostIP         string              `json:"hostIP" binding:"required"` // 节点ip
	Start          int64               `json:"start" binding:"required"`  // 开始时间
	End            int64               `json:"end" binding:"required"`    // 结束时间
	Step           uint                `json:"step" binding:"required"`   // 步长
}
