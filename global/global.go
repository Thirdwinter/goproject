package global

import (
	"goproject/config"
	"github.com/jinzhu/gorm"
)

var (
	Config *config.Config
	Db *gorm.DB
)
