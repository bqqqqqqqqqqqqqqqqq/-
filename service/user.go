package service

import (
	"dogking_shop/models"
	"dogking_shop/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "-1",
		"test": "true",
	})
}

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	email := c.Query("email")
	if email == "" || !util.IsEmail(email) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不正确",
		})
		return
	}
	code := util.GetRand()
	models.RDB.Set(c, email, code, time.Second*300)
	err := util.SendCodeHelp(c, email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send Code Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// Register 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userCode := c.PostForm("code")
	phone := c.PostForm("phone")
	//email := c.PostForm("email")
	if username == "" || password == "<empty>" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "信息错误",
		})
		return
	}
	//if !util.IsEmail(email) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": -1,
	//		"msg":  "邮箱格式错误",
	//	})
	//	return
	//}

	password = util.GetMd5(password)
	sysCode, err := models.RDB.Get(c, phone).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		log.Printf("获取验证码错误:%v\n", err)
		return
	}
	if sysCode != userCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}
	var cnt int64
	err = models.DB.Model(new(models.User)).Where("phone = ?", phone).Count(&cnt).Error
	if err != nil {
		log.Printf("错误:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证手机错误",
		})
		return
	}

	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱已被注册",
		})
		return
	}

	userUUID := util.GetUUID()
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
	err = models.DB.Create(&userData).Error
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "数据插入错误",
		})
		return
	}
	token, err := util.GenerateToken(userUUID, username, phone)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "token生成错误",
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

// Login 登陆
func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	if phone == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}
	password = util.GetMd5(password)

	data := new(models.User)
	err := models.DB.Where("phone= ? AND password = ? ", phone, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}
	token, err := util.GenerateToken(data.Uuid, data.Name, phone)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "token生成错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登陆成功",
		"data": token,
	})
}

// GetUserDetail 获取用户信息
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识不能为空",
		})
		return
	}
	data := new(models.User)
	err := models.DB.Omit("password").Where("identity = ? ", identity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Detail By Identity:" + identity + " Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
