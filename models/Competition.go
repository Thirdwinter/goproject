// 与比赛列表相关

package models

import (
	"goproject/global"
	"goproject/utils/rspcode"

	"github.com/jinzhu/gorm"
)

type Competition struct {
	gorm.Model
	Title string `gorm:"type:varchar(30)"`
	Info  string `gorm:"type:longtext"`
}

// 注册赛事前检测是否存在
func CheckCompetition(com *Competition) (code int) {
	var count int
	global.Db.Model(&Competition{}).Where("title = ?", com.Title).Count(&count)
	if count > 0 {
		return rspcode.ERROR_USERNAME_USED
	}
	return rspcode.SUCCESS
}

// 注册赛事
func CreateCompetition(com *Competition) (code int) {
	err := global.Db.Create(&com).Error
	if err != nil {
		return 500
	}
	return 200
}

// 更新赛事
func UpdateCompetition(ntitle,ninfo string) (code int) {
	var com Competition
	if ntitle != "" {
		global.Db.Find(&com, "title = ?",com.Title)
		if com.ID != 0 && ninfo != ""{
			com.Title = ntitle
			com.Info = ninfo
			global.Db.Save(&com)
			return 200
		}
	}
	return 500
}

// 删除赛事
func DelCompetiton() {

}
