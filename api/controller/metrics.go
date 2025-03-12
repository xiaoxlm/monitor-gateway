package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/domain/factory"
	"github.com/xiaoxlm/monitor-gateway/api/domain/repo"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"github.com/xiaoxlm/monitor-gateway/api/response"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	"github.com/xiaoxlm/monitor-gateway/internal/model"

	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"
)

func ListMetricsMapping(ctx context.Context, query *request.ListMetricsMappingQuery) ([]model.MetricsMapping, error) {

	return repo.ListMetricsMapping(ctx, config.Config.Mysql.GetDB(), query.Category, query.MetricsUniqueID)
}

func ListMetrics(ctx context.Context, queries []request.MetricsQueryInfo) (*response.ListMetricsRESP, error) {
	prom, err := prometheus.NewPrometheus(config.Config.Prom.GetClient())
	if err != nil {
		return nil, err
	}

	db := config.Config.Mysql.GetDB()

	m, err := factory.FactoryMetrics(ctx, db, prom, queries)
	if err != nil {
		return nil, err
	}

	multiExprValueList, err := m.ListValues()
	if err != nil {
		return nil, err
	}

	var respData []response.MetricsData
	for idx, v := range multiExprValueList {
		respData = append(respData, response.MetricsData{
			MetricUniqueID:   queries[idx].MetricUniqueID,
			HostIP:           v[0].Metric["host_ip"],
			MultiMetricsData: MetricMultiDataMapping(queries[idx].MetricUniqueID),
			Values:           v,
		})
	}

	return &response.ListMetricsRESP{
		Data: respData,
	}, nil
}

func MetricMultiDataMapping(uniqueID enum.MetricUniqueID) bool {
	switch uniqueID {
	case enum.MetricUniqueID_Gpu_All_Util:
		return true
	default:
		return false
	}
}
