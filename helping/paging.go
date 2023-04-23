package helping

import (
	"Doggggg/define"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Paging(c *gin.Context) (size, page int) {
	size, _ = strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("分页strconv错误:", err)
		return 0, 0
	}
	page = (page - 1) * size
	return size, page
}
