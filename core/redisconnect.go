package core

import (
	"fmt"
	"goproject/global"

	"github.com/gomodule/redigo/redis"
)


func InitRedis() {
	// 创建Redis连接
	c, err := redis.Dial("tcp", global.Config.Redis.Conn(), redis.DialPassword(global.Config.Redis.Password))
	if err != nil {
		panic(err)
	}
	fmt.Println("redis ok!")
	// 将连接赋值给全局变量
	global.Redis = c
}
