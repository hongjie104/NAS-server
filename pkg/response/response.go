package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin a
type Gin struct {
	C *gin.Context
}

// Success a
func (g *Gin) Success(data interface{}) {
	if data == nil {
		g.C.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    "ok",
		})
	} else {
		g.C.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    data,
		})
	}
}

// Fail a
func (g *Gin) Fail(err error) {
	g.C.JSON(http.StatusOK, gin.H{
		"success": false,
		"msg":     err.Error(),
	})
}
