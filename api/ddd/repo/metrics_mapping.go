package repo

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
	"gorm.io/gorm"
)

func CreateMetricsMapping(ctx context.Context, db *gorm.DB, metricsMapping *model.MetricsMapping) error {
	return db.WithContext(ctx).Create(metricsMapping).Error
}

func BatchCreateMetricsMapping(ctx context.Context, db *gorm.DB, metricsMappingList []*model.MetricsMapping) error {
	return db.WithContext(ctx).Create(&metricsMappingList).Error
}

func ListMetricsMapping(ctx context.Context, db *gorm.DB) ([]model.MetricsMapping, error) {
	var metricsMappingList []model.MetricsMapping
	err := db.WithContext(ctx).Find(&metricsMappingList).Error
	return metricsMappingList, err
}
