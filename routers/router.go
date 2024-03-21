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
	//r.Use(middleware.GetUser())

	user := r.Group("user")
	//user.Use(middleware.CheckUserRole())
	{
		//user.POST("login", v1.Login)
		user.POST("updatauimg",  v1.UpdateUserImage)
		user.GET("select/competiton/:pagesize/:pagenum", v1.SelectPageCompetiton)
		user.POST("select/info", v1.SelectUserDataById)
		user.POST("select/id", v1.SelectIdByPhone)
		user.POST("new/inf", v1.NewInform)
		user.POST("new/group", v1.CreateGroup)
		user.POST("select/group/id", v1.SelectGroupById)
		user.POST("select/group/lid", v1.SelectGroupByLid)
		user.POST("select/group/mid", v1.SelectGroupByMid)

	}
	admin := r.Group("admin")
	//admin.Use(middleware.CheckAdminRole())
	{
		admin.POST("login", v1.Login)
		admin.POST("newcom", v1.CreateCompetition)
		admin.POST("ucom", v1.UpdateCompetition)
		admin.POST("dcom", v1.DelCompetiton)
	}

	r.Run(global.Config.System.Addr())
}
