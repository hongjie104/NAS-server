package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/model"
	"github.com/hongjie104/NAS-server/pkg/response"
	"gopkg.in/mgo.v2/bson"
)

// SeriesController SeriesController
type SeriesController struct{}

// Index Index
func (ctl *SeriesController) Index(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	list, total := model.SeriesModelInstance.Index(page, pageSize, name)
	response := response.Gin{C: c}
	response.Success(gin.H{
		"list": list,
		"pagination": gin.H{
			"total": total,
		},
	})
}

// Show Show
func (ctl *SeriesController) Show(c *gin.Context) {
	id := c.Param("id")
	detail := model.SeriesModelInstance.Show(bson.ObjectIdHex(id))
	response := response.Gin{C: c}
	response.Success(detail)
}
