package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/internal/model"

	"github.com/xiaoxlm/monitor-gateway/api/service"
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"

	common_model "github.com/prometheus/common/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

func ListMetricsMapping(ctx context.Context) ([]model.MetricsMapping, error) {
	mapping := service.NewMetricsMapping(config.Config.Mysql.GetDB())
	return mapping.List(ctx)
}

func ListMetrics(ctx context.Context, queries []_interface.QueryFormItem) ([]common_model.Value, error) {
	prom, err := prometheus.NewPrometheus(config.Config.Prom.GetClient())
	if err != nil {
		return nil, err
	}

	m := service.NewMetrics(prom)
	if err = m.Query(ctx, queries); err != nil {
		return nil, err
	}
	return m.Output()
}
