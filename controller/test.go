package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/pkg/response"
)

// Test Test
func Test(c *gin.Context) {
	response := response.Gin{C: c}
	response.Success("this is a test")
}
