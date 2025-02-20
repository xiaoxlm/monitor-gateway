package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xiaoxlm/monitor-gateway/api/router/demo"
	_ "github.com/xiaoxlm/monitor-gateway/cmd/monitor-gateway/docs"
	"github.com/xiaoxlm/monitor-gateway/config"
)

func NewRoot(r *gin.Engine) {
	basePath := r.Group("/monitor-gateway/api")
	v1 := basePath.Group("/v1")

	demo.Router(v1)

}

func Start() {
	r := &router{}
	r.init().registerHandler().swagger()

	config.Config.Server.GinServe(r.getEngin())
}

type router struct {
	ginEngine *gin.Engine
}

func (r *router) init() *router {
	gin.SetMode(config.Config.Server.RunMode)
	r.ginEngine = gin.Default()

	return r
}

func (r *router) registerHandler() *router {
	NewRoot(r.ginEngine)

	return r
}

func (r *router) swagger() *router {
	//docs.SwaggerInfo.BasePath = "/monitor-gateway/api/v1"
	r.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func (r *router) getEngin() *gin.Engine {
	return r.ginEngine
}
