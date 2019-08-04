package utils

import (
	// "github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/pkg/setting"
)

// GetPage 获取分页页码
func GetPage(c *gin.Context) int {
	result := 0
	// page, _ := com.StrTo(c.Query("page")).Int()
	page := 10
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
