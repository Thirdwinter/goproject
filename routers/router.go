package routers

import (
	v1 "goproject/api/v1"
	"goproject/global"

	"github.com/gin-gonic/gin"
	"goproject/middleware"
)

func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	r.Use(middleware.Logrus())
	r.Use(middleware.Next())
	//r := gin.New()

	r.POST("api/user/add", v1.AddUser) // 注册
	r.POST("api/user/login", v1.Login)  // 登录
	//r.Use(middleware.JwtToken())

	user := r.Group("user")
	user.Use(middleware.CheckUserRole())
	{
		user.POST("login", v1.Login)
	}
	admin := r.Group("admin")
	admin.Use(middleware.CheckAdminRole())
	{
		admin.POST("login", v1.Login)
	}

	r.Run(global.Config.System.Addr())
}
