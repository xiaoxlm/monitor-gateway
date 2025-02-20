package request

import _interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"

type MetricsBatchQueryBody struct {
	Queries []_interface.QueryFormItem `json:"queries" binding:"required"`
}
