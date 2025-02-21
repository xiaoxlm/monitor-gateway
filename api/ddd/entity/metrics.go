package entity

import (
	"context"

	"github.com/prometheus/common/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

func FactoryMetrics(ctx context.Context, tsDB _interface.TimeSeriesDB, queries []_interface.QueryFormItem) (*Metrics, error) {
	m := &Metrics{timeSeriesDB: tsDB, queries: queries}

	if err := m.query(ctx); err != nil {
		return nil, err
	}

	return m, nil
}

type Metrics struct {
	timeSeriesDB _interface.TimeSeriesDB

	queries []_interface.QueryFormItem

	output []model.Value
}

func (m *Metrics) Output() ([]model.Value, error) {
	return m.output, nil
}

func (m *Metrics) query(ctx context.Context) error {
	values, err := m.timeSeriesDB.BatchQueryRange(ctx, m.queries)
	if err != nil {
		return err
	}

	m.output = values

	return nil
}
