package factory

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/ddd/repo"
	"gorm.io/gorm"

	"github.com/xiaoxlm/monitor-gateway/api/ddd/entity"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
)

func FactoryMetricsMapping(ctx context.Context, db *gorm.DB, queries []request.MetricsQueryInfo) (*entity.MetricsMapping, error) {
	labels := make(map[enum.MetricUniqueID]*entity.Labels)
	for _, query := range queries {
		labels[query.MetricUniqueID] = &entity.Labels{
			IBN:    query.IBN,
			HostIP: query.HostIP,
		}
	}

	mm := entity.NewMetricsMapping(labels)

	metricsMappingList, err := repo.ListMetricsMappingByUniqueID(ctx, db, mm)
	if err != nil {
		return nil, err
	}

	var expressionMap = make(map[enum.MetricUniqueID]string)
	for _, metricsMapping := range metricsMappingList {
		expressionMap[metricsMapping.MetricUniqueID] = metricsMapping.Expression
	}

	mm.SetExpression(expressionMap)

	return mm, nil
}
