package entity

import (
	"context"
	"github.com/lie-flat-planet/httputil"
	"github.com/xiaoxlm/monitor-gateway/api/response"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	model2 "github.com/xiaoxlm/monitor-gateway/internal/model"
)

type Aggr struct {
	//queries            []domain_model.MetricsQuery
	metricsMappingList []*MetricsMapping

	boardPayloadList              []model2.BoardPayload
	metricUniqueID2MetricsMapping map[enum.MetricUniqueID]model2.MetricsMapping
}

func newAggr(metricsMappingList []*MetricsMapping, boardPayloadList []model2.BoardPayload, metricUniqueID2MetricsMapping map[enum.MetricUniqueID]model2.MetricsMapping) *Aggr {
	return &Aggr{
		metricsMappingList:            metricsMappingList,
		boardPayloadList:              boardPayloadList,
		metricUniqueID2MetricsMapping: metricUniqueID2MetricsMapping,
	}
}

func (agg *Aggr) ListMetricsValue(ctx context.Context) ([]response.MetricsData, error) {
	var respData = make([]response.MetricsData, 0)

	for _, mapping := range agg.metricsMappingList {

		metricUniqueID := mapping.getMapping().MetricUniqueID
		hostIP := mapping.getQuery().LabelValue["host_ip"]

		exprValue, err := mapping.getMetricsFromExpr(ctx)
		if err != nil {
			return nil, err
		}

		if err = agg.setColor(metricUniqueID, exprValue); err != nil {
			return nil, err
		}

		if exprValue == nil {
			exprValue = make(httputil.MetricsFromExpr, 0)
		}

		respData = append(respData, response.MetricsData{
			MetricUniqueID:   metricUniqueID,
			Label:            mapping.getQuery().LabelValue,
			HostIP:           hostIP,
			MultiMetricsData: metricMultiDataMapping(metricUniqueID),
			Values:           exprValue,
		})

	}

	return respData, nil
}

func (agg *Aggr) setColor(metricUniqueID enum.MetricUniqueID, data httputil.MetricsFromExpr) error {
	if data == nil {
		return nil
	}

	boardPayloadID := agg.metricUniqueID2MetricsMapping[metricUniqueID].BoardPayloadID

	panelID := agg.metricUniqueID2MetricsMapping[metricUniqueID].PanelID
	panel, err := model2.GetPanelByBoardIDAndPanelID(agg.boardPayloadList, boardPayloadID, panelID)
	if err != nil {
		return err
	}

	colorH := colorHandlerVO{
		panel:       panel,
		metricsData: data,
	}
	return colorH.setMetricsDataColor()
}

func metricMultiDataMapping(uniqueID enum.MetricUniqueID) bool {
	switch uniqueID {
	case enum.MetricUniqueID_Gpu_All_Util:
		return true
	default:
		return false
	}
}
