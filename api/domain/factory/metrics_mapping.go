package factory

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/domain/model"
	"github.com/xiaoxlm/monitor-gateway/api/domain/repo"
	"gorm.io/gorm"

	"github.com/xiaoxlm/monitor-gateway/api/domain/entity"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
)

func FactoryMetricsMapping(ctx context.Context, db *gorm.DB, queries []model.MetricsQuery) (*entity.MetricsMapping, error) {
	labelValue := make(map[enum.MetricUniqueID]map[string]string)
	for _, query := range queries {
		labelValue[query.MetricUniqueID] = query.LabelValue
	}

	mm, err := entity.NewMetricsMapping(labelValue)
	if err != nil {
		return nil, err
	}

	metricsMappingList, err := repo.ListMetricsMappingByUniqueID(ctx, db, mm)
	if err != nil {
		return nil, err
	}

	mm.SetMappingList(metricsMappingList)

	return mm, nil
}
