package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/models"
	"github.com/hongjie104/NAS-server/app/pkg/e"
	"github.com/hongjie104/NAS-server/app/pkg/utils"
	"github.com/hongjie104/NAS-server/app/routers/api"
)

var userModel = &models.UserModel{}

// GetUser a
func GetUser(c *gin.Context) {
	response := api.Gin{C: c}
	id, ok := c.Get("id")
	if !ok {
		response.Fail(0)
		return
	}
	idStr := id.(string)
	u := userModel.Show(idStr)
	response.Success(gin.H{
		"name":   u.Name,
		"avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		"userid": u.ID.Hex(),
	})
}

// Login a
func Login(c *gin.Context) {
	// LoginData 绑定为json
	type LoginData struct {
		User     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	response := api.Gin{C: c}
	var loginData LoginData
	if err := c.BindJSON(&loginData); err == nil {
		u := userModel.Login(loginData.User)
		if u.Password == loginData.Password {
			token, _ := utils.GenerateToken(u.Name, u.ID.Hex())
			response.Success(gin.H{"token": token, "currentAuthority": "admin"})
		} else {
			response.Fail(e.LoginError)
		}
	} else {
		response.Fail(e.InvalidParams)
	}
}
