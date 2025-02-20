package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/service"
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"

	"github.com/prometheus/common/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

func ListMetrics(ctx context.Context, queries []_interface.QueryFormItem) ([]model.Value, error) {
	prom, err := prometheus.NewPrometheus(config.Config.Prom.GetClient())
	if err != nil {
		return nil, err
	}

	m := service.NewMetrics(prom)
	m.QueryMetrics(ctx, queries)
	return m.Output()

}
