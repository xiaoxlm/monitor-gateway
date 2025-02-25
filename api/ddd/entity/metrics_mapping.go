package entity

import (
	"fmt"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
)

type MetricsMapping struct {
	labels     map[enum.MetricUniqueID]*Labels
	expression map[enum.MetricUniqueID]string
}

func (m *MetricsMapping) CheckLabels() error {
	if len(m.labels) < 1 {
		return fmt.Errorf("MetricsMapping entity labels is empty")
	}

	return nil
}

func (m *MetricsMapping) CheckExpressions() error {
	if len(m.expression) < 1 {
		return fmt.Errorf("MetricsMapping entity expression is empty")
	}

	return nil
}

func NewMetricsMapping(labels map[enum.MetricUniqueID]*Labels) *MetricsMapping {
	return &MetricsMapping{
		labels: labels,
	}
}

func (m *MetricsMapping) ListLabel() []enum.MetricUniqueID {
	var list []enum.MetricUniqueID
	for k := range m.labels {
		list = append(list, k)
	}
	return list
}

func (m *MetricsMapping) SetExpression(expression map[enum.MetricUniqueID]string) {
	m.expression = expression
}

func (m *MetricsMapping) GetExpression(metricUniqueID enum.MetricUniqueID) string {
	return m.expression[metricUniqueID]
}
