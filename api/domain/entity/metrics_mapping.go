package entity

import (
	"fmt"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	"strings"
)

type MetricsMapping struct {
	labelValue map[enum.MetricUniqueID]map[string]string
	expression map[enum.MetricUniqueID]string // 数据库存的

	parsedExpression map[enum.MetricUniqueID]string
}

func NewMetricsMapping(labelValue map[enum.MetricUniqueID]map[string]string) (*MetricsMapping, error) {
	mm := &MetricsMapping{
		labelValue:       labelValue,
		parsedExpression: make(map[enum.MetricUniqueID]string),
	}

	err := mm.checkLabels()

	return mm, err
}

func (m *MetricsMapping) checkLabels() error {
	if len(m.labelValue) < 1 {
		return fmt.Errorf("MetricsMapping entity labels is empty")
	}

	return nil
}

func (m *MetricsMapping) checkExpressions() error {
	if len(m.expression) < 1 {
		return fmt.Errorf("MetricsMapping entity expression is empty")
	}

	return nil
}

func (m *MetricsMapping) ListMetricUniqueID() []enum.MetricUniqueID {
	var list []enum.MetricUniqueID
	for k := range m.labelValue {
		list = append(list, k)
	}
	return list
}

func (m *MetricsMapping) SetExpression(expression map[enum.MetricUniqueID]string) {
	m.expression = expression
}

func (m *MetricsMapping) parseExpression() error {
	if err := m.checkExpressions(); err != nil {
		return err
	}
	for uniqueID, expr := range m.expression {
		var replaceExpr = expr
		for k, v := range m.labelValue[uniqueID] {
			replaceExpr = strings.ReplaceAll(replaceExpr, "$"+k, v)
		}
		m.parsedExpression[uniqueID] = replaceExpr
	}
	return nil
}

func (m *MetricsMapping) GetParsedExpression(metricUniqueID enum.MetricUniqueID) (string, error) {
	if err := m.parseExpression(); err != nil {
		return "", err
	}

	return m.parsedExpression[metricUniqueID], nil
}
