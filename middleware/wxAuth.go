package middleware

import (
	"Doggggg/Init"
	"Doggggg/helping"
	"Doggggg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func AuthPhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		phone := c.Query("phone")

		user := &models.User{Phone: phone}
		// 查询是否注册过
		var count int64
		err := Init.DB.Model(new(models.User)).Where("phone = ? ", phone).Find(&user).Count(&count).Error
		if err != nil {
			log.Println("查询错误：", err)
			return
		}
		if count <= 0 {
			AuthPhoneRegisetr()
		}
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "登陆已过期,请重新登陆。",
				"data": nil,
			})
			c.Abort()
			return
		}
		userClaim, err := helping.AnalyseToken(auth)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code":  http.StatusUnauthorized,
				"error": err,
				"msg":   "Unauthorized Authorization",
			})
			c.Abort()
			return
		}
		c.Set("userClaim", userClaim)
	}

}

func AuthPhoneRegisetr() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 若未注册
		username := c.PostForm("username")
		password := c.PostForm("password")
		phone := c.PostForm("phone")
		if username == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "请提供用户名来保证程序运行",
			})
			return
		}
		if password == "" {
			password = "123456"
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "请及时修改初始密码",
			})
		}
		// 注册
		userUUID := helping.GetUUID()
		userData := &models.User{
			Uuid:      userUUID,
			Name:      username,
			Phone:     phone,
			IsStudent: "",
			Password:  password,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		}
		err := Init.DB.Create(&userData).Error
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "数据插入失败",
			})
			return
		}
		token, err := helping.GenerateToken(userUUID, username, phone)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "token生成错误",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功",
			"data": map[string]interface{}{
				"token": token,
			},
		})

	}
}
