package model

import "github.com/xiaoxlm/monitor-gateway/internal/enum"

type MetricsQuery struct {
	MetricUniqueID enum.MetricUniqueID
	LabelValue     map[string]string
	Start          int64 // 开始时间
	End            int64 // 结束时间
	Step           uint  // 步长
}

func ListMetricsQueryMetricUniqueID(queries []MetricsQuery) []enum.MetricUniqueID {
	var ret []enum.MetricUniqueID
	for _, q := range queries {
		ret = append(ret, q.MetricUniqueID)
	}

	return ret
}
