package models

import (
	"dogking_shop/define"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type Problem struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
	Goods   int    `gorm:"column:goods" json:"goods"`
	Comment string `gorm:"column:Comment" json:"comment"`
}

func (table *Problem) TableName() string {
	return "problem"
}

// GetAllProblem 分页查询问题
func GetAllProblem(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("分页strconv错误:", err)
		return
	}
	page = (page - 1) * size
	problems := make([]*Problem, 0)
	tx := DB.Model(&Problem{}).Offset(page).Limit(10).Omit("password").Find(&problems)
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
		"data": problems,
	})
}
