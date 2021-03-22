package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/model"
	"github.com/hongjie104/NAS-server/pkg/response"
	"github.com/hongjie104/NAS-server/pkg/utils"
)

// UserController UserController
type UserController struct{}

// CurrentUser CurrentUser
func (ctl *UserController) CurrentUser(c *gin.Context) {
	id, _ := c.Get("id")
	response := response.Gin{C: c}
	u := model.UserModelInstance.Show(id.(string))
	if u.Name == "" {
		response.Fail(errors.New("用户名错误"))
		return
	}
	response.Success(gin.H{
		"name":   u.Name,
		"avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		"userid": u.ID.Hex(),
	})
}

// Login 登录
func (ctl *UserController) Login(c *gin.Context) {
	// LoginData 绑定为json
	type LoginData struct {
		User     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	response := response.Gin{C: c}
	var loginData LoginData
	if err := c.BindJSON(&loginData); err == nil {
		uModel := model.UserModelInstance.Login(loginData.User)
		if uModel.Password == "" {
			response.Fail(errors.New("用户名错误"))
			return
		}
		if uModel.Password != loginData.Password {
			response.Fail(errors.New("密码错误"))
			return
		}
		token, _ := utils.GenerateToken(loginData.User, uModel.ID.Hex())
		response.Success(gin.H{"token": token, "currentAuthority": "admin"})
	} else {
		response.Fail(errors.New("参数错误"))
	}
}

/*
// Index Index
func (ctl *UserController) Index(c *gin.Context) {
	uList := model.UserModelInstance.Index()
	response := response.Gin{C: c}
	response.Success(uList)
}

// Show Show
func (ctl *UserController) Show(c *gin.Context) {
	id := c.Param("id")
	u := model.UserModelInstance.Show(id)
	response := response.Gin{C: c}
	response.Success(u)
}
*/
