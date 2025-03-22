package controller

import (
	"context"
	"fmt"

	"github.com/spf13/cast"

	"github.com/xiaoxlm/monitor-gateway/api/domain/factory"
	"github.com/xiaoxlm/monitor-gateway/api/domain/repo"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"github.com/xiaoxlm/monitor-gateway/api/response"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	"github.com/xiaoxlm/monitor-gateway/internal/model"

	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/pkg/metrics/prometheus"

	domain_model "github.com/xiaoxlm/monitor-gateway/api/domain/model"
)

func ListMetricsMapping(ctx context.Context, query *request.ListMetricsMappingQuery) ([]model.MetricsMapping, error) {

	return repo.ListMetricsMapping(ctx, config.Config.Mysql.GetDB(), query.Category, query.MetricsUniqueID)
}

func ListMetrics(ctx context.Context, queryInfos []request.MetricsQueryInfo) (*response.ListMetricsRESP, error) {
	prom, err := prometheus.NewPrometheus(config.Config.Prom.GetClient())
	if err != nil {
		return nil, err
	}

	db := config.Config.Mysql.GetDB()

	queries, err := ConvertMetricsQueryInfo(queryInfos)
	if err != nil {
		return nil, err
	}

	m, err := factory.FactoryMetrics(ctx, db, prom, queries)
	if err != nil {
		return nil, err
	}

	multiExprValueList, err := m.ListValues()
	if err != nil {
		return nil, err
	}

	var respData = make([]response.MetricsData, 0)
	for idx, v := range multiExprValueList {
		respData = append(respData, response.MetricsData{
			MetricUniqueID:   queries[idx].MetricUniqueID,
			HostIP:           v[0].Metric["host_ip"],
			MultiMetricsData: MetricMultiDataMapping(queries[idx].MetricUniqueID),
			Values:           v,
		})
	}

	return &response.ListMetricsRESP{
		Data: respData,
	}, nil
}

func MetricMultiDataMapping(uniqueID enum.MetricUniqueID) bool {
	switch uniqueID {
	case enum.MetricUniqueID_Gpu_All_Util:
		return true
	default:
		return false
	}
}

const (
	FixedQueryKeyStart = "start"
	FixedQueryKeyEnd   = "end"
	FixedQueryKeyStep  = "step"
)

func ConvertMetricsQueryInfo(queries []request.MetricsQueryInfo) (ret []domain_model.MetricsQuery, err error) {

	for _, query := range queries {

		var (
			start, end int64
			step       uint
		)
		{
			startSTR, ok := query.LabelValue[FixedQueryKeyStart]
			if !ok {
				return nil, fmt.Errorf("start label is required")
			}
			start, err = cast.ToInt64E(startSTR)
			if err != nil {
				return nil, err
			}

			endSTR, ok := query.LabelValue[FixedQueryKeyEnd]
			if !ok {
				return nil, fmt.Errorf("end label is required")
			}
			end, err = cast.ToInt64E(endSTR)
			if err != nil {
				return nil, err
			}

			stepSTR, ok := query.LabelValue[FixedQueryKeyStep]
			if !ok {
				return nil, fmt.Errorf("step label is required")
			}
			step, err = cast.ToUintE(stepSTR)
			if err != nil {
				return nil, err
			}

		}

		ret = append(ret, domain_model.MetricsQuery{
			MetricUniqueID: query.MetricUniqueID,
			LabelValue:     query.LabelValue,
			Start:          start,
			End:            end,
			Step:           step,
		})
	}

	return ret, nil
}
