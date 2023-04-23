package models

import (
	"Doggggg/Init"
	"Doggggg/define"
	"Doggggg/helping"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	//gorm.Model
	ID        uint           `gorm:"primarykey;" json:"id"`
	Uuid      string         `gorm:"column:uuid" json:"uuid"`
	Name      string         `gorm:"column:name" json:"name"`
	Phone     string         `gorm:"column:phone" json:"phone"`
	IsStudent string         `gorm:"column:isStudent" json:"isStudent"`
	Password  string         `gorm:"column:password;"`
	CreatedAt time.Time      `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
}

func (table *User) TableName() string {
	return "users"
}

// 暴露给前端的struct
//type UserAPI struct {
//	ID        uint   `gorm:"primarykey;" json:"id"`
//	Uuid      string `gorm:"column:uuid" json:"uuid"`
//	IsStudent string `gorm:"column:isAdmin;type:varchar(2)" json:"isAdmination"`
//}

// AddStudent 添加学生
func AddStudent(c *gin.Context) {
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	password := ""
	userUUID := helping.GetUUID()
	userData := &User{
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
			"msg": "数据插入错误",
		})
		return
	}

}

// GetAllUser 分页查询
func GetAllUser(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("分页strconv错误:", err)
		return
	}
	page = (page - 1) * size
	users := make([]*User, 0)
	tx := Init.DB.Model(&User{}).Offset(page).Limit(10).Omit("password").Find(&users)
	err = tx.Error
	if err != nil {
		log.Println("查询错误", err)
		return
	}
	tx.Order("User.id DESC")
	//for _, v := range users {
	//	fmt.Printf("user ===> %v\n", v)
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": users,
	})
}
