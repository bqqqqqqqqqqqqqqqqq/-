package router

import (
	"dogking_shop/middleware"
	"dogking_shop/models"
	"dogking_shop/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.Ping)
	r.POST("/code", service.SendCode)
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
	r.POST("/AllProblem", models.GetAllProblem)
	r.POST("/AllStudent", models.GetAllUser)
	r.POST("/wxAuth", middleware.AuthPhone())

	return r
}
