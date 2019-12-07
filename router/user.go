package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/controller"
	"github.com/hongjie104/NAS-server/middleware"
)

func initUserRouter(r *gin.Engine) {
	ctl := &controller.UserController{}
	r.POST("/api/admin/user/login", ctl.Login)
	api := r.Group("/api/admin")
	api.Use(middleware.JWT())
	{
		initController(api, "user", ctl)
		api.GET("/user/currentUser", ctl.CurrentUser)
	}
}
