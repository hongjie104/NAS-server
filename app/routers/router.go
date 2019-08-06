package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/pkg/setting"
	v1 "github.com/hongjie104/NAS-server/app/routers/api/v1"
)

// InitRouter a
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "success": true})
	})

	r.POST("/api/v1/login", v1.Login)

	apiV1 := r.Group("/api/v1")
	// apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/actress", v1.IndexActress)
		apiV1.GET("/actress/:id", v1.ShowActress)
		// apiV1.GET("/test/:id", v1.Test)

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