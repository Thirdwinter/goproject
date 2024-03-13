package routers

import (
	"goproject/global"
	"goproject/api/v1"

	"github.com/gin-gonic/gin"
)


func InitRouter() {
	gin.SetMode(global.Config.System.Env)
	r:=gin.Default()
	//r := gin.New()
	//r.Use(middleware.Logrus())
	user:=r.Group("api")
	{
		user.POST("user/add", v1.AddUser)
	}

	r.Run(global.Config.System.Addr())
}