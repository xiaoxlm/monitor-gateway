package service

import (
	"context"

	"github.com/xiaoxlm/monitor-gateway/internal/model"

	"gorm.io/gorm"
)

type MetricsMapping struct {
	db *gorm.DB
}

func NewMetricsMapping(db *gorm.DB) *MetricsMapping {
	return &MetricsMapping{db: db}
}

func (m *MetricsMapping) Create(ctx context.Context, metricsMapping *model.MetricsMapping) error {
	return m.db.WithContext(ctx).Create(metricsMapping).Error
}

func (m *MetricsMapping) BatchCreate(ctx context.Context, metricsMappingList []*model.MetricsMapping) error {
	return m.db.WithContext(ctx).Create(&metricsMappingList).Error
}

func (m *MetricsMapping) List(ctx context.Context) ([]model.MetricsMapping, error) {

	var metricsMappingList []model.MetricsMapping
	err := m.db.WithContext(ctx).Find(&metricsMappingList).Error
	return metricsMappingList, err
}
