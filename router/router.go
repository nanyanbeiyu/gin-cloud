package router

import (
	"github.com/gin-gonic/gin"
	v1 "go_cloud/api/v1"
	"go_cloud/middleware"
	"go_cloud/util/conf"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.SetTrustedProxies([]string{"127.0.0.1"})
	//登录
	r.POST("/login", v1.Login)
	//注册
	r.POST("/register", v1.Register)
	r.Run(conf.HTTP_PORT)
}
