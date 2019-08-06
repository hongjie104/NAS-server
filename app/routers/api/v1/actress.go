package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/models"
	response "github.com/hongjie104/NAS-server/app/routers/response"
)

// IndexActress 获取女演员列表
func IndexActress(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	actress := models.IndexActress(page, pageSize)
	response := response.Gin{C: c}
	response.Success(actress)
}

// ShowActress a
func ShowActress(c *gin.Context) {
	id := c.Param("id")
	actress := models.ShowActress(id)
	response := response.Gin{C: c}
	response.Success(actress)
}

// // AddTag 新增文章标签
// func AddTag(c *gin.Context) {
// }

// // EditTag 修改文章标签
// func EditTag(c *gin.Context) {
// }

// // DeleteTag 删除文章标签
// func DeleteTag(c *gin.Context) {
// }
