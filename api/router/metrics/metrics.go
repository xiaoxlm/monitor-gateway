package metrics

import (
	"fmt"
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
// @Param Authorization header string true "Authorization bearer token"
// @Success 200 {object} []model.MetricsMapping 成功
// @Failure 500 {object} httputil.ErrorRESP 失败
// @Router /monitor-gateway/api/v1/metrics/mapping [GET]
// @ID ListMetricsMapping
func ListMetricsMapping(ctx *gin.Context) {
	datas, err := controller.ListMetricsMapping(ctx)

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
// @Param Authorization header string true "Authorization bearer token"
// @Param Body body request.MetricsBatchQueryBody true "body"
// @Success 200 {object} []model.Value 成功
// @Failure 400 {object} httputil.ErrorRESP 失败
// @Failure 500 {object} httputil.ErrorRESP 失败
// @Router /monitor-gateway/api/v1/metrics/batch-query [POST]
// @ID BatchQuery
func BatchQuery(ctx *gin.Context) {
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
