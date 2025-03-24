package entity

import (
	"fmt"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
	"strings"
)

type MetricsMapping struct {
	labelValue map[enum.MetricUniqueID]map[string]string
	//expression map[enum.MetricUniqueID]string // 数据库存的

	mappingList []model.MetricsMapping

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
	if len(m.mappingList) < 1 {
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

func (m *MetricsMapping) SetMappingList(mappingList []model.MetricsMapping) {
	m.mappingList = mappingList
}

func (m *MetricsMapping) GetParsedExpression(metricUniqueID enum.MetricUniqueID) (string, error) {
	if err := m.parseExpression(); err != nil {
		return "", err
	}

	return m.parsedExpression[metricUniqueID], nil
}

func (m *MetricsMapping) parseExpression() error {
	if err := m.checkExpressions(); err != nil {
		return err
	}

	expressionMap := m.metricUniqueID2Expression()

	for uniqueID, expr := range expressionMap {
		var replaceExpr = expr
		for k, v := range m.labelValue[uniqueID] {
			replaceExpr = strings.ReplaceAll(replaceExpr, "$"+k, v)
		}
		m.parsedExpression[uniqueID] = replaceExpr
	}
	return nil
}

func (m *MetricsMapping) metricUniqueID2Expression() map[enum.MetricUniqueID]string {
	var expressionMap = make(map[enum.MetricUniqueID]string)
	for _, metricsMapping := range m.mappingList {
		expressionMap[metricsMapping.MetricUniqueID] = metricsMapping.Expression
	}

	return expressionMap
}
