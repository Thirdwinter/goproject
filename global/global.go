package global

import (
	"goproject/config"
	"github.com/jinzhu/gorm"
	"github.com/gomodule/redigo/redis"
)

var (
	Config *config.Config
	Db *gorm.DB
	Redis redis.Conn
)
