package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/pkg/e"
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
		})
	} else {
		g.C.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    data,
		})
	}
}

// Fail a
func (g *Gin) Fail(errCode int) {
	g.C.JSON(http.StatusOK, gin.H{
		"success": false,
		"errCode": errCode,
		"errMsg":  e.GetMsg(errCode),
	})
}
