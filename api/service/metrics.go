package service

import (
	"context"
	"github.com/prometheus/common/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

type Metrics struct {
	timeSeriesDB _interface.TimeSeriesDB

	output []model.Value
}

func NewMetrics(tsDB _interface.TimeSeriesDB) *Metrics {
	return &Metrics{timeSeriesDB: tsDB}
}

func (m *Metrics) Output() ([]model.Value, error) {
	return m.output, nil
}

func (m *Metrics) QueryMetrics(ctx context.Context, queries []_interface.QueryFormItem) error {
	values, err := m.timeSeriesDB.BatchQueryRange(ctx, queries)
	if err != nil {
		return err
	}

	m.output = values

	return nil
}
