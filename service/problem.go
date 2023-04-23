package service

import (
	"Doggggg/Init"
	"Doggggg/define"
	"Doggggg/helping"
	"Doggggg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// 查询自己的问题
func UserProblem(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("分页strconv错误:", err)
		return
	}
	page = (page - 1) * size
	problems := make([]*models.Problem, 0)
	tx := Init.DB.Model(&models.Problem{}).Offset(page).Limit(10).Find(&problems)
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

// 发表问题 (内有上传图片)
func PublishProblem(c *gin.Context) {
	userID := c.PostForm("user_id")
	username := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")
	form, err := c.MultipartForm()
	//userID := form.Value["user_id"]
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "上传错误",
		})
		return
	}
	files := form.File["file"]
	for k, file := range files {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			})
			return
		}
		nameFileName := helping.GetNowTime() + ":" + username + strconv.Itoa(k)
		err := c.SaveUploadedFile(file, "./image/blog/"+nameFileName+fileExt)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "上传错误",
			})
			return
		}
		newPicture := models.Picture{
			Path:   "/image/blog/" + nameFileName,
			UserId: userID,
		}
		fmt.Println(newPicture)
		err = Init.DB.Create(&newPicture).Error
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "用户创建错误",
			})
			return
		}
	}

	newProblem := models.Problem{
		Title:   title,
		Goods:   0,
		Content: content,
	}
	err = Init.DB.Create(&newProblem).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "创建错误",
		})
		return
	}
}

// 管理员删除问题
func AdminDeleteProblem(c *gin.Context) {
	problemID := c.Param("problemID")
	problem := new(models.Problem)
	err := Init.DB.Model(&models.Problem{}).Where("id = ?", problemID).First(problem).Error
	if err != nil || err != gorm.ErrRecordNotFound {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除错误或问题不存在",
		})
	}
	Init.DB.Delete(&problem)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})

}

// 删除自己的问题
func DeleteProblem(c *gin.Context) {
	problemID := c.Param("problemID")
	userID := c.Param("userID")
	fproblemUser := new(models.User)
	err := Init.DB.Model(&models.Problem{}).Where("problem_id = ?", problemID).First(fproblemUser).Error
	if err != nil {
		log.Println(err)
		return
	}
	if fproblemUser.Uuid != userID {
		log.Println(err)
		return
	}
	problem := new(models.Problem)
	err = Init.DB.Model(&models.Problem{}).Where("id = ?", problemID).First(problem).Error
	if err != nil || err != gorm.ErrRecordNotFound {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除错误或问题不存在",
		})
	}
	Init.DB.Delete(&problem)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
