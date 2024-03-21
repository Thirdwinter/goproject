// 与比赛列表相关

package models

import (
	"errors"
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
func UpdateCompetition(title, ntitle, ninfo string) (Competition ,int) {
	var com Competition
	if ntitle != "" {
		global.Db.Find(&com, "title = ?", title)
		if com.ID != 0 && ninfo != "" {
			com.Title = ntitle
			com.Info = ninfo
			global.Db.Save(&com)
			return com,200
		}
	}
	return com, 500
}

// 删除赛事
func DelCompetiton(title string) int {
	var com Competition

	global.Db.First(&com, "title=?", title)
	if com.ID != 0 {
		result := global.Db.Delete(com)
		if result.Error != nil {
			return 500
		}
		return 200
	}
	return 500
}


// 分页查询
func SelectPage(pageSize, pageNum int) ([]Competition, int) {
	var competitions []Competition
	var total int

	var err error
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize
	err = global.Db.Limit(pageSize).Offset(offset).Find(&competitions).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return competitions, 0
		}
		return nil, 0
	}
	err = global.Db.Model(&User{}).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return competitions, total
}
