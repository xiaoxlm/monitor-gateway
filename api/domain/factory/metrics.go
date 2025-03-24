package factory

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/domain/entity"
	"github.com/xiaoxlm/monitor-gateway/api/domain/model"
	"github.com/xiaoxlm/monitor-gateway/api/domain/repo"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	model2 "github.com/xiaoxlm/monitor-gateway/internal/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"gorm.io/gorm"
)

// FactoryMetrics creates and initializes a new Metrics aggregate
func FactoryMetrics(ctx context.Context, db *gorm.DB, tsDB _interface.TimeSeriesDB, queries []model.MetricsQuery) (*entity.Metrics, error) {
	items, mm, err := convertMetricsQueryInfoToItem(ctx, db, queries)
	if err != nil {
		return nil, err
	}

	var metricUniqueID2MetricsMapping = make(map[enum.MetricUniqueID]model2.MetricsMapping)
	{
		mappingList := mm.GetMappingList()
		metricUniqueID2MetricsMapping = model2.MetricUniqueID2MetricsMapping(mappingList)
	}

	var panel = &model.Panel{}
	{
		panel, err = repo.GetPanelContent(ctx, metricsMapping.BoardPayloadID, metricsMapping.PanelID)
		if err != nil {
			return nil, err
		}
	}

	metrics := entity.NewMetrics(items, tsDB, metricUniqueID2MetricsMapping, panel)

	return metrics, nil
}

// convertMetricsQueryInfoToItem 获取完整表达式并转换数据结构
func convertMetricsQueryInfoToItem(ctx context.Context, db *gorm.DB, queries []model.MetricsQuery) ([]_interface.QueryFormItem, *entity.MetricsMapping, error) {
	ret := make([]_interface.QueryFormItem, 0, len(queries))
	mm, err := FactoryMetricsMapping(ctx, db, queries)
	if err != nil {
		return nil, nil, err
	}

	for _, query := range queries {
		q, err := mm.GetParsedExpression(query.MetricUniqueID)
		if err != nil {
			return nil, nil, err
		}

		ret = append(ret, _interface.QueryFormItem{
			Start: query.Start,
			End:   query.End,
			Step:  query.Step,
			Query: q,
		})
	}

	return ret, mm, nil
}
