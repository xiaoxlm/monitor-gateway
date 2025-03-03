package metrics

import (
	"fmt"
	"github.com/xiaoxlm/monitor-gateway/api/response"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lie-flat-planet/httputil"
	"github.com/xiaoxlm/monitor-gateway/api/controller"
	"github.com/xiaoxlm/monitor-gateway/api/request"
	"github.com/xiaoxlm/monitor-gateway/config"
)

// ListMetricsMapping
// @BasePath /
// PingExample godoc
// @Summary ListMetricsMapping
// @Schemes
// @Description 获取指标映射
// @Tags ListMetricsMapping
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Basic token"
// @Param category query enum.MetrcisMappingCategory false "类别"
// @Param metricsUniqueID query string false "metrics唯一id"
// @Success 200 {object} []internal_model.MetricsMapping 成功
// @Failure 500 {object} httputil.ErrorRESP 失败
// @Router /monitor-gateway/api/v1/metrics/mapping [GET]
// @ID ListMetricsMapping
func ListMetricsMapping(ctx *gin.Context) {
	var _ enum.MetrcisMappingCategory

	var query = request.ListMetricsMappingQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		(&httputil.RESP{
			Content:     "",
			ServiceCode: config.Config.Server.Code,
			Err:         fmt.Errorf("query parse failed. err:%v", err),
			HttpCode:    http.StatusBadRequest,
		}).Output(ctx)
		return
	}

	datas, err := controller.ListMetricsMapping(ctx, &query)

	(&httputil.RESP{
		Content:     datas,
		ServiceCode: config.Config.Server.Code,
		Err:         err,
	}).Output(ctx)

	return
}

// BatchQuery
// @BasePath /
// PingExample godoc
// @Summary BatchQuery
// @Schemes
// @Description 更具PromQL查询指标
// @Tags BatchQuery
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Basic token"
// @Param Body body request.MetricsBatchQueryBody true "body"
// @Success 200 {object} response.ListMetricsRESP 成功
// @Failure 400 {object} httputil.ErrorRESP 失败
// @Failure 500 {object} httputil.ErrorRESP 失败
// @Router /monitor-gateway/api/v1/metrics/batch-query [POST]
// @ID BatchQuery
func BatchQuery(ctx *gin.Context) {
	var _ response.ListMetricsRESP
	var body = request.MetricsBatchQueryBody{}
	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		(&httputil.RESP{
			Content:     "",
			ServiceCode: config.Config.Server.Code,
			Err:         fmt.Errorf("body parse failed. err:%v", err),
			HttpCode:    http.StatusBadRequest,
		}).Output(ctx)
		return
	}

	values, err := controller.ListMetrics(ctx, body.Queries)
	(&httputil.RESP{
		Content:     values,
		ServiceCode: config.Config.Server.Code,
		Err:         err,
	}).Output(ctx)

	return
}
