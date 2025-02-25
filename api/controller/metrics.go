package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/ddd/factory"
	"github.com/xiaoxlm/monitor-gateway/api/ddd/repo"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"github.com/xiaoxlm/monitor-gateway/internal/model"

	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"

	common_model "github.com/prometheus/common/model"
)

func ListMetricsMapping(ctx context.Context, query *request.ListMetricsMappingQuery) ([]model.MetricsMapping, error) {

	return repo.ListMetricsMapping(ctx, config.Config.Mysql.GetDB(), query.Category, query.MetricsUniqueID)
}

func ListMetrics(ctx context.Context, queries []request.MetricsQueryInfo) ([]common_model.Value, error) {
	prom, err := prometheus.NewPrometheus(config.Config.Prom.GetClient())
	if err != nil {
		return nil, err
	}

	db := config.Config.Mysql.GetDB()

	m, err := factory.FactoryMetrics(ctx, db, prom, queries)
	if err != nil {
		return nil, err
	}
	return m.GetValues()
}
