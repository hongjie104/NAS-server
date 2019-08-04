package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/pkg/e"
	"github.com/hongjie104/NAS-server/app/pkg/utils"
)

// JWT a
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		// var data interface{}
		code = e.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"errCode": code,
				"errMsg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
