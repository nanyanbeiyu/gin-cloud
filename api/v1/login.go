package v1

import (
	"github.com/gin-gonic/gin"
	"go_cloud/middleware"
	"go_cloud/model"
	"go_cloud/util/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	// 获取参数
	var data model.User
	// 获取用户信息
	c.ShouldBind(&data)

	// 校验用户
	user, code := model.CheckLogin(data.UserName, data.PassWord)
	if code != errmsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetMsg(code),
			"data": nil,
		})
	}

	// 获取jwt
	token := middleware.SetJWTToken(data.UserName)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errmsg.GetMsg(code),
		"data":  *user,
		"token": token,
	})
}
