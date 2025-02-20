package metrics

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r = r.Group("metrics")

	r.POST("/batch-query", BatchQuery)
}
