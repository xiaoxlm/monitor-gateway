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

func ListMetricsMapping(ctx context.Context, db *gorm.DB, category, uniqueID string) ([]model.MetricsMapping, error) {
	var metricsMappingList []model.MetricsMapping
	sql := db.WithContext(ctx)
	if category != "" {
		sql = sql.Where("category = ?", category)
	}
	if uniqueID != "" {
		sql = sql.Where("metric_unique_id LIKE ?", "%"+uniqueID+"%")
	}

	err := sql.Find(&metricsMappingList).Error
	return metricsMappingList, err
}
