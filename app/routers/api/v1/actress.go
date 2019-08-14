package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/models"
	"github.com/hongjie104/NAS-server/app/pkg/config"
	"github.com/hongjie104/NAS-server/app/pkg/e"
	"github.com/hongjie104/NAS-server/app/routers/api"
)

// ActressModel ActressModel
var ActressModel = &models.ActressModel{}

// ActressController ActressController
type ActressController struct{}

// Index 获取女演员列表
func (ctl *ActressController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", strconv.Itoa(config.Config.APP.PageSize)))
	name := c.DefaultQuery("name", "")
	sortBy := c.DefaultQuery("sortBy", "")
	actress, total := ActressModel.Index(models.ActressIndexOption{
		Page:     page,
		PageSize: pageSize,
		Name:     name,
		SortBy:   sortBy,
	})
	response := api.Gin{C: c}
	response.Success(gin.H{
		"list":  actress,
		"total": total,
	})
}

// Show a
func (ctl *ActressController) Show(c *gin.Context) {
	id := c.Param("id")
	actress := ActressModel.Show(id)
	response := api.Gin{C: c}
	response.Success(gin.H{"detail": actress})
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
