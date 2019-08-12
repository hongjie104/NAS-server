package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/models"
	"github.com/hongjie104/NAS-server/app/pkg/config"
	"github.com/hongjie104/NAS-server/app/pkg/e"
	"github.com/hongjie104/NAS-server/app/routers/api"
)

var videoModel = &models.VideoModel{}

// VideoController VideoController
type VideoController struct{}

// Index 获取女演员列表
func (ctl *VideoController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", strconv.Itoa(config.Config.APP.PageSize)))
	video, total := videoModel.Index(page, pageSize)
	response := api.Gin{C: c}
	response.Success(gin.H{
		"list":  video,
		"total": total,
	})
}

// Show a
func (ctl *VideoController) Show(c *gin.Context) {
	id := c.Param("id")
	video := videoModel.Show(id)
	response := api.Gin{C: c}
	response.Success(video)
}

// Update Update
func (ctl *VideoController) Update(c *gin.Context) {
	// 田嶋涼子
	type UpdateData struct {
		Name string `json:"name" binding:"required"`
	}
	response := api.Gin{C: c}
	// id := c.Param("id")
	var updateData UpdateData
	if err := c.BindJSON(&updateData); err == nil {
		// model.Update(id, bson.M{"$set"})
		response.Success(nil)
	} else {
		response.Fail(e.InvalidParams)
	}
}
