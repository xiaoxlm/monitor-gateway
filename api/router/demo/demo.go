package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lie-flat-planet/httputil"
	"github.com/xiaoxlm/monitor-gateway/api/controller"
	resp "github.com/xiaoxlm/monitor-gateway/api/response"
	"github.com/xiaoxlm/monitor-gateway/config"
	"net/http"
)

type UserBody struct {
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

// FetchFirst
// @BasePath /
// PingExample godoc
// @Summary FetchFirst
// @Schemes
// @Description 查询用户第一条数据
// @Tags FetchFirst
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization bearer token"
// @Param Body body UserBody true "body"
// @Param userID query string true "用户id"
// @Success 200 {object} resp.UserRESP 成功
// @Failure 400 {object} httputil.ErrorRESP 失败
// @Failure 500 {object} httputil.ErrorRESP 失败
// @Router /monitor-gateway/api/v1/demo/hello-world [POST]
// @ID FetchFirst
func FetchFirst(ctx *gin.Context) {
	userID := ctx.Query("userID")
	fmt.Println("query userID:", userID)

	var body = UserBody{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		(&httputil.RESP{
			Content:     "",
			ServiceCode: config.Config.Server.Code,
			Err:         fmt.Errorf("body parse failed. err:%v", err),
			HttpCode:    http.StatusBadRequest,
		}).Output(ctx)
		return
	}
	fmt.Println("body:", body)

	ctl := controller.FactoryDemo()

	user, err := ctl.FetchFirst(ctx)
	(&httputil.RESP{
		Content: resp.UserRESP{
			ID:       user.ID.ID,
			Username: user.Username,
			Nickname: user.Nickname,
		},
		ServiceCode: config.Config.Server.Code,
		Err:         err,
	}).Output(ctx)

	return
}
