package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/middleware"
	"github.com/hongjie104/NAS-server/app/pkg/config"
	v1 "github.com/hongjie104/NAS-server/app/routers/api/v1"
)

// InitRouter a
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.Config.Server.RunMode)

	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong", "success": true})
	})

	r.POST("/api/v1/login", v1.Login)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		actressController := &v1.ActressController{}
		apiV1.GET("/actress", actressController.Index)
		apiV1.GET("/actress/:id", actressController.Show)
		apiV1.PUT("/actress/:id", actressController.Update)
		// apiV1.GET("/test/:id", v1.Test)

		videoController := &v1.VideoController{}
		apiV1.GET("/video", videoController.Index)
		apiV1.GET("/video/:id", videoController.Show)
		apiV1.PUT("/video/:id", videoController.Update)

		apiV1.GET("/user/currentUser", v1.GetUser)
	}

	return r
}

/*
记录一下，总是忘记，还要去查，烦得要死
Method        Path                Route Name      Controller.Action
GET           /posts              posts           app.controllers.posts.index
GET           /posts/new          new_post        app.controllers.posts.new
GET           /posts/:id          post            app.controllers.posts.show
GET           /posts/:id/edit     edit_post       app.controllers.posts.edit
POST          /posts              posts           app.controllers.posts.create
PUT           /posts/:id          post            app.controllers.posts.update
DELETE        /posts/:id          post            app.controllers.posts.destroy
*/
