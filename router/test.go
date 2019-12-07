package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/controller"
)

func initTestRouter(r *gin.Engine) {
	r.GET("/api/test", controller.Test)
}
