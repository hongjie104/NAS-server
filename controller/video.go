package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/model"
	"github.com/hongjie104/NAS-server/pkg/response"
	"github.com/hongjie104/NAS-server/pkg/utils"
	"gopkg.in/mgo.v2/bson"
)

// VideoController VideoController
type VideoController struct{}

// Index Index
func (ctl *VideoController) Index(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	sortBy := c.DefaultQuery("sortBy", "")
	actressID := c.DefaultQuery("actressId", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	seriesID := c.DefaultQuery("seriesId", "")
	indexOption := &model.VideoIndexOption{
		Page:      page,
		PageSize:  pageSize,
		Code:      code,
		SortBy:    sortBy,
		ActressID: actressID,
		SeriesID:  seriesID,
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

// Show Show
func (ctl *VideoController) Show(c *gin.Context) {
	id := c.Param("id")
	detail := model.VideoModelInstance.Show(id)
	series := model.SeriesModelInstance.Show(detail.Series)
	categoryArr := model.CategoryModelInstance.ShowMany(detail.Category)
	actress := model.ActressModelInstance.ShowByIDList(detail.Actress)
	response := response.Gin{C: c}
	response.Success(gin.H{
		"video":       detail,
		"series":      series,
		"categoryArr": categoryArr,
		"actress":     actress,
	})
}

// Update update
func (ctl *VideoController) Update(c *gin.Context) {
	response := response.Gin{C: c}
	var newData model.VideoModel
	err := c.BindJSON(&newData)
	if err == nil {
		id := c.Param("id")
		newDataBson, _ := utils.Struct2Bson(newData)
		model.VideoModelInstance.Update(id, bson.M{"$set": newDataBson})
		response.Success(nil)
	} else {
		response.Fail(err)
	}
}
