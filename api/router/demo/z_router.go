package demo

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r = r.Group("demo")

	r.POST("/hello-world", FetchFirst)
}
