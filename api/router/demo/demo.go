package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/lie-flat-planet/httputil"
	"github.com/xiaoxlm/monitor-gateway/api/controller"
	"github.com/xiaoxlm/monitor-gateway/config"
)

func FetchFirst(ctx *gin.Context) {
	ctl := controller.FactoryDemo()

	user, err := ctl.FetchFirst(ctx)

	(&httputil.RESP{
		Content:     user,
		ServiceCode: config.Config.Server.Code,
		Err:         err,
	}).Output(ctx)

	return
}