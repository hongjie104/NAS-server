package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/pkg/e"
	"github.com/hongjie104/NAS-server/app/pkg/utils"
	"github.com/hongjie104/NAS-server/app/routers/api"
)

// GetUser a
func GetUser(c *gin.Context) {
	response := api.Gin{C: c}
	response.Success(gin.H{
		"name":   "鸿杰",
		"avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		"userid": 3630,
	})
}

// LoginData 绑定为json
type LoginData struct {
	User     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login a
func Login(c *gin.Context) {
	response := api.Gin{C: c}
	var loginData LoginData
	if err := c.BindJSON(&loginData); err == nil {
		if loginData.User == "admin" && loginData.Password == "123" {
			token, _ := utils.GenerateToken(loginData.User, loginData.Password)
			response.Success(gin.H{"token": token})
		} else {
			response.Fail(e.LoginError)
		}
	} else {
		response.Fail(e.InvalidParams)
	}
}
