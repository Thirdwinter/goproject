package main


import(
	"goproject/core"
	"goproject/routers"
	
)

func main() {
	core.InitConf()
	//core.Initlog()
	core.InitGorm()
	core.InitRedis()
	routers.InitRouter()
}