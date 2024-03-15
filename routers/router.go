package routers

import (
	v1 "goproject/api/v1"
	"goproject/global"

	"goproject/middleware"
	"github.com/gin-gonic/gin"
)


func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	r:=gin.Default()
	r.Use(middleware.Logrus())
	r.Use(middleware.Next())
	r.Use(middleware.JwtToken())
	//r := gin.New()
	
	r.POST("user/login", v1.Login)	// 登录
	r.POST("user/add", v1.AddUser)	// 注册
	
	user:=r.Group("api")
	user.Use(middleware.CheckUserRole())
	{
	}

	r.Run(global.Config.System.Addr())
}