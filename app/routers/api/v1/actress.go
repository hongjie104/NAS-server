package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/models"
	"github.com/hongjie104/NAS-server/app/routers/api"
)

var model = &models.ActressModel{}

// ActressController ActressController
type ActressController struct{}

// Index 获取女演员列表
func (controller *ActressController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	actress := model.Index(page, pageSize)
	response := api.Gin{C: c}
	response.Success(actress)
}

// Show a
func (controller *ActressController) Show(c *gin.Context) {
	id := c.Param("id")
	actress := model.Show(id)
	response := api.Gin{C: c}
	response.Success(actress)
}
