package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/ddd/repo"
	"github.com/xiaoxlm/monitor-gateway/internal/model"

	"github.com/xiaoxlm/monitor-gateway/api/ddd/entity"
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"

	common_model "github.com/prometheus/common/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

func ListMetricsMapping(ctx context.Context) ([]model.MetricsMapping, error) {
	return repo.ListMetricsMapping(ctx, config.Config.Mysql.GetDB())
}

func ListMetrics(ctx context.Context, queries []_interface.QueryFormItem) ([]common_model.Value, error) {
	prom, err := prometheus.NewPrometheus(config.Config.Prom.GetClient())
	if err != nil {
		return nil, err
	}

	m, err := entity.FactoryMetrics(ctx, prom, queries)
	if err != nil {
		return nil, err
	}
	return m.GetValues()
}
