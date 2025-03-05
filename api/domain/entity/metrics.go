package entity

import (
	"context"
	"fmt"
	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
)

// MetricsQuery represents a domain object for metrics queries
type metricsQuery struct {
	Query string
	Start int64
	End   int64
	Step  uint
}

// Metrics is the aggregate root entity for metrics data
type Metrics struct {
	queries      []metricsQuery
	timeSeriesDB _interface.TimeSeriesDB
	values       []model.Value
}

// NewMetrics creates a new Metrics aggregate
func NewMetrics(queries []_interface.QueryFormItem, tsDB _interface.TimeSeriesDB) *Metrics {
	metricQueries := make([]metricsQuery, len(queries))
	for i, q := range queries {
		metricQueries[i] = metricsQuery{
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
func (m *Metrics) GetValues() ([]MetricsFormedData, error) {
	var ret []MetricsFormedData
	for _, result := range m.values {
		switch result.Type() {
		case model.ValScalar:
			logrus.Warnf("need to parse 'Scalar' type value")
		case model.ValVector:
			logrus.Warnf("need to parse 'Vector' type value")
		case model.ValMatrix:
			var values []Values
			matrix := result.(model.Matrix)
			for _, stream := range matrix {
				for _, value := range stream.Values {
					values = append(values, Values{
						Value:     value.Value.String(),
						Timestamp: value.Timestamp.Unix(),
					})
				}
			}

			ret = append(ret, MetricsFormedData{
				Values: values,
			})
		case model.ValString:
			logrus.Warnf("need to parse 'String' type value")
		default:
			return nil, fmt.Errorf("unknown metric type: %s", result.Type())
		}
	}

	return ret, nil
}

type MetricsFormedData struct {
	Values []Values
}

type Values struct {
	Value     string
	Timestamp int64
}
