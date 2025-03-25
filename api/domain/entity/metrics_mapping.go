package entity

import (
	"context"
	"fmt"
	"github.com/lie-flat-planet/httputil"
	prom_model "github.com/prometheus/common/model"
	domain_model "github.com/xiaoxlm/monitor-gateway/api/domain/model"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"strings"
)

type MetricsMapping struct {
	mapping      *model.MetricsMapping
	query        *domain_model.MetricsQuery
	timeSeriesDB _interface.TimeSeriesDB

	parsedExpression string

	promValue prom_model.Value
}

func NewMetricsMapping(query *domain_model.MetricsQuery, mapping model.MetricsMapping, timeSeriesDB _interface.TimeSeriesDB) (*MetricsMapping, error) {
	mm := &MetricsMapping{
		query:        query,
		mapping:      &mapping,
		timeSeriesDB: timeSeriesDB,
	}

	err := mm.checkLabels()

	return mm, err
}

func (m *MetricsMapping) checkLabels() error {
	if len(m.query.LabelValue) < 1 {
		return fmt.Errorf("MetricsMapping entity labels is empty")
	}

	return nil
}

func (m *MetricsMapping) checkExpression() error {
	return nil
}

func (m *MetricsMapping) getMapping() *model.MetricsMapping {
	return m.mapping
}

func (m *MetricsMapping) getQuery() *domain_model.MetricsQuery {
	return m.query
}

func (m *MetricsMapping) fetchMetrics(ctx context.Context) error {
	if err := m.parseExpression(); err != nil {
		return err
	}

	value, err := m.timeSeriesDB.QueryRange(ctx, _interface.QueryFormItem{
		Start: m.query.Start,
		End:   m.query.End,
		Step:  m.query.Step,
		Query: m.parsedExpression,
	})
	if err != nil {
		return err
	}

	m.promValue = value
	return nil
}

func (m *MetricsMapping) getMetricsFromExpr(ctx context.Context) (httputil.MetricsFromExpr, error) {
	if err := m.fetchMetrics(ctx); err != nil {
		return nil, err
	}

	return httputil.ParseModelValue2MetricsData(m.promValue)
}

func (m *MetricsMapping) parseExpression() error {
	if err := m.checkExpression(); err != nil {
		return err
	}

	var replaceExpr = m.mapping.Expression

	for k, v := range m.query.LabelValue {
		k = m.replaceLabelKey(k)
		replaceExpr = strings.ReplaceAll(replaceExpr, "$"+k, v)
	}

	m.parsedExpression = replaceExpr

	return nil
}

func (m *MetricsMapping) replaceLabelKey(key string) string {
	tmpKey := strings.ToLower(key)

	if tmpKey == "ibn" {
		return "IBN"
	}

	if tmpKey == "hostip" {
		return "host_ip"
	}

	return key
}
