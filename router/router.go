package router

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/config"
)

// InitRouter a
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	gin.SetMode(config.Config.Server.RunMode)

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong 1.0.0", "success": true})
	})

	initTestRouter(r)
	initUserRouter(r)
	initActressRouter(r)
	initVideoRouter(r)
	initSeriesRouter(r)

	return r
}

type indexController interface {
	Index(c *gin.Context)
}

type showController interface {
	Show(c *gin.Context)
}

type updateController interface {
	Update(c *gin.Context)
}

type createController interface {
	Create(c *gin.Context)
}

type destoryController interface {
	Destroy(c *gin.Context)
}

func initController(routerGroup *gin.RouterGroup, routerName string, ctl interface{}) {
	t := reflect.TypeOf(ctl)
	if _, existing := t.MethodByName("Index"); existing {
		routerGroup.GET("/"+routerName, ctl.(indexController).Index)
	}
	if _, existing := t.MethodByName("Show"); existing {
		routerGroup.GET("/"+routerName+"/show/:id", ctl.(showController).Show)
	}
	if _, existing := t.MethodByName("Update"); existing {
		routerGroup.PUT("/"+routerName+"/update/:id", ctl.(updateController).Update)
	}
	if _, existing := t.MethodByName("Create"); existing {
		routerGroup.POST("/"+routerName, ctl.(createController).Create)
	}
	if _, existing := t.MethodByName("Destroy"); existing {
		routerGroup.DELETE("/"+routerName+"/:id", ctl.(destoryController).Destroy)
	}
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
