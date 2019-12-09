package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/model"
	"github.com/hongjie104/NAS-server/pkg/response"
	"github.com/hongjie104/NAS-server/pkg/utils"
	"gopkg.in/mgo.v2/bson"
)

// ActressController ActressController
type ActressController struct{}

// Index Index
func (ctl *ActressController) Index(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	sortBy := c.DefaultQuery("sortBy", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	actressIndexOption := &model.ActressIndexOption{
		Name:     name,
		Page:     page,
		PageSize: pageSize,
		SortBy:   sortBy,
	}
	list, total := model.ActressModelInstance.Index(*actressIndexOption)
	response := response.Gin{C: c}
	response.Success(gin.H{
		"list": list,
		"pagination": gin.H{
			"total": total,
		},
	})
}

// Show Show
func (ctl *ActressController) Show(c *gin.Context) {
	id := c.Param("id")
	detail := model.ActressModelInstance.Show(id)
	response := response.Gin{C: c}
	response.Success(detail)
}

// Update update
func (ctl *ActressController) Update(c *gin.Context) {
	response := response.Gin{C: c}
	var newData model.ActressModel
	err := c.BindJSON(&newData)
	if err == nil {
		id := c.Param("id")
		newDataBson, _ := utils.Struct2Bson(newData)
		model.ActressModelInstance.Update(id, bson.M{"$set": newDataBson})
		response.Success(nil)
	} else {
		response.Fail(err)
	}
}
