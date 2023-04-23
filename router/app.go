package router

import (
	"Doggggg/middleware"
	"Doggggg/models"
	"Doggggg/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	//r.MaxMultipartMemory = 8 << 20
	r.GET("/ping", service.Ping)

	//  公共方法
	r.POST("/code", service.SendCode)
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
	r.POST("/AllStudent", models.GetAllUser)
	// 操作问题
	r.POST("/AllProblem", models.AllProblem)
	r.POST("/uProblem", service.PublishProblem)
	r.POST("/delProblem", service.DeleteProblem)
	r.POST("/serProblem", models.SearchProblem)

	// 私有方法
	// 管理者
	// 操作学生
	r.POST("/wxAuth", middleware.AuthPhone())

	return r
}
