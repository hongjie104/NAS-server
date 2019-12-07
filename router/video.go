package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/controller"
	"github.com/hongjie104/NAS-server/middleware"
)

func initVideoRouter(r *gin.Engine) {
	ctl := &controller.VideoController{}
	api := r.Group("/api/admin")
	api.Use(middleware.JWT())
	{
		initController(api, "video", ctl)
	}
}
