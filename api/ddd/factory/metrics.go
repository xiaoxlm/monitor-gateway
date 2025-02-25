package factory

import (
	"context"
	"gorm.io/gorm"

	"github.com/xiaoxlm/monitor-gateway/api/ddd/entity"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

// FactoryMetrics creates and initializes a new Metrics aggregate
func FactoryMetrics(ctx context.Context, db *gorm.DB, tsDB _interface.TimeSeriesDB, queries []request.MetricsQueryInfo) (*entity.Metrics, error) {
	items, err := convertMetricsQueryInfoToItem(ctx, db, queries)
	if err != nil {
		return nil, err
	}

	metrics := entity.NewMetrics(items, tsDB)
	if err = metrics.FetchMetrics(ctx); err != nil {
		return nil, err
	}
	return metrics, nil
}

func convertMetricsQueryInfoToItem(ctx context.Context, db *gorm.DB, queries []request.MetricsQueryInfo) ([]_interface.QueryFormItem, error) {
	ret := make([]_interface.QueryFormItem, len(queries))
	mm, err := FactoryMetricsMapping(ctx, db, queries)
	if err != nil {
		return nil, err
	}

	if err = mm.CheckExpressions(); err != nil {
		return nil, err
	}

	for _, query := range queries {
		ret = append(ret, _interface.QueryFormItem{
			Start: query.Start,
			End:   query.End,
			Step:  query.Step,
			Query: mm.GetExpression(query.MetricUniqueID),
		})
	}

	return ret, nil
}
