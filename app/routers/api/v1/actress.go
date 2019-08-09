package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/models"
	"github.com/hongjie104/NAS-server/app/pkg/e"
	"github.com/hongjie104/NAS-server/app/routers/api"
)

var model = &models.ActressModel{}

// ActressController ActressController
type ActressController struct{}

// Index 获取女演员列表
func (ctl *ActressController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	actress := model.Index(page, pageSize)
	response := api.Gin{C: c}
	response.Success(actress)
}

// Show a
func (ctl *ActressController) Show(c *gin.Context) {
	id := c.Param("id")
	actress := model.Show(id)
	response := api.Gin{C: c}
	response.Success(actress)
}

// Update Update
func (ctl *ActressController) Update(c *gin.Context) {
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
