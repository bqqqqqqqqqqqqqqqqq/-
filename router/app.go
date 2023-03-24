package router

import (
	"dogking_shop/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.Ping)

	return r
}
