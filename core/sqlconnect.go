package core

import (
	"time"

	"goproject/global"
	//"github.com/ThirdWinter/Go/gvb_servehr/models"
	"goproject/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func InitGorm() {
	global.Db, err = gorm.Open("mysql", global.Config.Mysql.Dsn())
	if err != nil {
		panic(err)
	}
	global.Db.DB().SetMaxIdleConns(10)
	global.Db.DB().SetMaxOpenConns(100)
	global.Db.DB().SetConnMaxLifetime(10 * time.Second)
	global.Db.SingularTable(true)                                                // 禁止表名复数化
	global.Db.AutoMigrate(&models.User{},&models.Competition{},models.Inform{},models.Group{}) // 数据模型迁移
	//global.Db.Close()
}
