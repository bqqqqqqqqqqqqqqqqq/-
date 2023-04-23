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

type Problem struct {
	ID        uint           `gorm:"primarykey;" json:"id"`
	Title     string         `gorm:"column:title" json:"title"`
	Goods     int            `gorm:"column:goods" json:"goods"`
	Content   string         `gorm:"column:content" json:"content"`
	AuthorId  int            `gorm:"column:author_id" json:"authorId"`
	CreatedAt time.Time      `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
}

//type Content struct {
//	ID     string `gorm:"column:id" json:"id"`
//	UserId *User  `gorm:"foreignKey:id:reference:id" json:"UserId"`
//}

func (table *Problem) TableName() string {
	return "problem"
}

// 查询问题
func AllProblem(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("分页strconv错误:", err)
		return
	}
	page = (page - 1) * size
	problems := make([]*Problem, 0)
	tx := Init.DB.Model(&Problem{}).Offset(page).Limit(10).Find(&problems)
	err = tx.Error
	if err != nil {
		log.Println("查询错误", err)
		return
	}
	tx.Order("problems.id DESC")
	//for _, v := range users {
	//	fmt.Printf("user ===> %v\n", v)
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": problems,
	})
}

// SearchProblem 分页查询问题
func SearchProblem(c *gin.Context) {
	_, page := helping.Paging(c)
	problems := make([]*Problem, 0)
	author := c.PostForm("author")
	content := c.PostForm("content")
	title := c.PostForm("title")
	tx := Init.DB.Model(&Problem{}).Offset(page).Limit(10).Omit("password").Where("author like ? or content like ? or title like ?", author+"%", content+"%", title+"%").Find(&problems)
	err := tx.Error
	if err != nil {
		log.Println("查询错误", err)
		return
	}
	//tx.Order("User.id DESC")
	//for _, v := range users {
	//	fmt.Printf("user ===> %v\n", v)
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": problems,
	})
}
