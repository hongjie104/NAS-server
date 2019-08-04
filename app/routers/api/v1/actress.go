package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/models"
	response "github.com/hongjie104/NAS-server/app/pkg/app"
)

// IndexAcctress 获取女演员列表
func IndexAcctress(c *gin.Context) {
	response := response.Gin{C: c}
	acctress := models.Index(1, 10)
	response.Success(acctress)
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
