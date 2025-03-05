package factory

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/domain/entity"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"gorm.io/gorm"
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
	ret := make([]_interface.QueryFormItem, 0, len(queries))
	mm, err := FactoryMetricsMapping(ctx, db, queries)
	if err != nil {
		return nil, err
	}

	for _, query := range queries {
		q, err := mm.GetParsedExpression(query.MetricUniqueID)
		if err != nil {
			return nil, err
		}

		ret = append(ret, _interface.QueryFormItem{
			Start: query.Start,
			End:   query.End,
			Step:  query.Step,
			Query: q,
		})
	}

	return ret, nil
}
