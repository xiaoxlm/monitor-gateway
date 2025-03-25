package entity

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/api/domain/model"
	"github.com/xiaoxlm/monitor-gateway/api/domain/repo"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	model2 "github.com/xiaoxlm/monitor-gateway/internal/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"gorm.io/gorm"
)

func FactoryAggr(ctx context.Context, db *gorm.DB, tsDB _interface.TimeSeriesDB, queries []model.MetricsQuery) (*Aggr, error) {
	var (
		metricUniqueID2MetricsMapping = make(map[enum.MetricUniqueID]model2.MetricsMapping)
		boardPayloadList              []model2.BoardPayload
	)
	{
		metricUniqueIDList := model.ListMetricsQueryMetricUniqueID(queries)

		mappingList, err := repo.ListMetricsMappingByUniqueIDs(ctx, db, metricUniqueIDList)
		if err != nil {
			return nil, err
		}
		metricUniqueID2MetricsMapping = model2.MetricUniqueID2MetricsMapping(mappingList)

		boardPayloadIDList := model2.ListBoardPayloadID(mappingList)
		boardPayloadList, err = repo.ListBoardPayload(ctx, db, boardPayloadIDList)
		if err != nil {
			return nil, err
		}
	}

	var entityMetricsMappingList []*MetricsMapping
	for _, query := range queries {
		mm, err := NewMetricsMapping(&query, metricUniqueID2MetricsMapping[query.MetricUniqueID], tsDB)
		if err != nil {
			return nil, err
		}

		entityMetricsMappingList = append(entityMetricsMappingList, mm)
	}

	return newAggr(entityMetricsMappingList, boardPayloadList, metricUniqueID2MetricsMapping), nil
}
