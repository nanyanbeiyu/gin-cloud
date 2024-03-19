package v1

import (
	"github.com/gin-gonic/gin"
	"go_cloud/model"
	"go_cloud/util/errmsg"
)

// 注册
func Register(c *gin.Context) {
	// 获取数据
	var data model.User
	// 绑定数据
	c.ShouldBindJSON(&data)
	// 注册逻辑
	code, user := model.Register(data.UserName, data.PassWord)
	if code != errmsg.Success {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  errmsg.GetMsg(code),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  errmsg.GetMsg(code),
		"data": user.OpenID,
	})
}
