package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/model"
	"github.com/hongjie104/NAS-server/pkg/response"
)

// VideoController VideoController
type VideoController struct{}

// Index Index
func (ctl *VideoController) Index(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	actressID := c.DefaultQuery("actressId", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	indexOption := &model.VideoIndexOption{
		Page:      page,
		PageSize:  pageSize,
		Code:      code,
		ActressID: actressID,
	}
	list, total := model.VideoModelInstance.Index(*indexOption)
	response := response.Gin{C: c}
	response.Success(gin.H{
		"list": list,
		"pagination": gin.H{
			"total": total,
		},
	})
}
