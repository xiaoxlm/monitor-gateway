package entity

import (
	"context"
	"github.com/prometheus/common/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

// MetricsQuery represents a domain object for metrics queries
type MetricsQuery struct {
	Query string
	Start int64
	End   int64
	Step  uint
}

// Metrics is the aggregate root entity for metrics data
type Metrics struct {
	queries      []MetricsQuery
	timeSeriesDB _interface.TimeSeriesDB
	values       []model.Value
}

// NewMetrics creates a new Metrics aggregate
func NewMetrics(queries []_interface.QueryFormItem, tsDB _interface.TimeSeriesDB) *Metrics {
	metricQueries := make([]MetricsQuery, len(queries))
	for i, q := range queries {
		metricQueries[i] = MetricsQuery{
			Query: q.Query,
			Start: q.Start,
			End:   q.End,
			Step:  q.Step,
		}
	}

	return &Metrics{
		queries:      metricQueries,
		timeSeriesDB: tsDB,
	}
}

// FetchMetrics fetches metrics data from the time series database
func (m *Metrics) FetchMetrics(ctx context.Context) error {
	// Convert domain queries back to interface format
	interfaceQueries := make([]_interface.QueryFormItem, len(m.queries))
	for i, q := range m.queries {
		interfaceQueries[i] = _interface.QueryFormItem{
			Query: q.Query,
			Start: q.Start,
			End:   q.End,
			Step:  q.Step,
		}
	}

	values, err := m.timeSeriesDB.BatchQueryRange(ctx, interfaceQueries)
	if err != nil {
		return err
	}

	m.values = values
	return nil
}

// GetValues returns the fetched metrics values
func (m *Metrics) GetValues() ([]model.Value, error) {
	return m.values, nil
}
