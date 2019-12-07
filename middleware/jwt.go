package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/pkg/utils"
)

// JWT a
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		// var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		// var claims *utils.Claims
		var claims *utils.Claims
		var err error
		if token == "" {
			code = 401
		} else {
			claims, err = utils.ParseToken(token)
			if err != nil {
				code = 401
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401
			}
		}
		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"code":    code,
				"msg":     "jwt error",
			})
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
		c.Next()
	}
}
