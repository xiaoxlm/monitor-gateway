package entity

import (
	"context"
	"sort"

	"github.com/lie-flat-planet/httputil"
	"github.com/prometheus/common/model"
	domain_model "github.com/xiaoxlm/monitor-gateway/api/domain/model"
	"github.com/xiaoxlm/monitor-gateway/api/response"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"

	model2 "github.com/xiaoxlm/monitor-gateway/internal/model"
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

	boardPayloadList              []model2.BoardPayload
	metricUniqueID2MetricsMapping map[enum.MetricUniqueID]model2.MetricsMapping
}

// NewMetrics creates a new Metrics aggregate
func NewMetrics(queries []_interface.QueryFormItem, tsDB _interface.TimeSeriesDB, metricUniqueID2MetricsMapping map[enum.MetricUniqueID]model2.MetricsMapping, boardPayloadList []model2.BoardPayload) *Metrics {
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
		queries:                       metricQueries,
		timeSeriesDB:                  tsDB,
		metricUniqueID2MetricsMapping: metricUniqueID2MetricsMapping,
		boardPayloadList:              boardPayloadList,
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
func (m *Metrics) ListValues(ctx context.Context, queries []domain_model.MetricsQuery) ([]response.MetricsData, error) {
	if err := m.FetchMetrics(ctx); err != nil {
		return nil, err
	}

	multiExprValueList, err := httputil.PromCommonModelValue(m.values)
	if err != nil {
		return nil, err
	}

	return m.metricsFromExpr2RESPMetricsData(queries, multiExprValueList)

}

func (m *Metrics) metricsFromExpr2RESPMetricsData(queries []domain_model.MetricsQuery, multiExprValueList []httputil.MetricsFromExpr) ([]response.MetricsData, error) {

	var respData = make([]response.MetricsData, 0)
	for idx, v := range multiExprValueList {
		metricUniqueID := queries[idx].MetricUniqueID
		if err := m.setColor(metricUniqueID, v); err != nil {
			return nil, err
		}

		respData = append(respData, response.MetricsData{
			MetricUniqueID:   metricUniqueID,
			HostIP:           v[0].Metric["host_ip"],
			MultiMetricsData: MetricMultiDataMapping(queries[idx].MetricUniqueID),
			Values:           v,
		})
	}

	return respData, nil
}

func (m *Metrics) setColor(metricUniqueID enum.MetricUniqueID, data httputil.MetricsFromExpr) error {
	boardPayloadID := m.metricUniqueID2MetricsMapping[metricUniqueID].BoardPayloadID
	panelID := m.metricUniqueID2MetricsMapping[metricUniqueID].PanelID
	panel, err := model2.GetPanelByBoardIDAndPanelID(m.boardPayloadList, boardPayloadID, panelID)
	if err != nil {
		return err
	}

	colorH := colorHandlerVO{
		panel:       panel,
		metricsData: data,
	}
	return colorH.setMetricsDataColor()
}

func MetricMultiDataMapping(uniqueID enum.MetricUniqueID) bool {
	switch uniqueID {
	case enum.MetricUniqueID_Gpu_All_Util:
		return true
	default:
		return false
	}
}

func sortThresholdsStepsByValue(p model2.Panel) []model2.Step {
	thresholdsSteps := p.Options.Thresholds.Steps

	sort.Slice(thresholdsSteps, func(i, j int) bool {
		if thresholdsSteps[i].Value == nil || thresholdsSteps[j].Value == nil {
			return false
		}
		return *thresholdsSteps[i].Value > *thresholdsSteps[j].Value
	})

	return thresholdsSteps
}
