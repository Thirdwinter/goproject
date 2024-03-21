package models

import (
	"goproject/global"

	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	LeaderId uint `form:"lid"`  // 组长id
	MemberId uint `form:"mid"`  // 组员id
	Cid      uint `form:"cid"`  // 关联比赛id
}

func CreateGroup(lid, cid, mid uint) (Group, int) {
	var group Group
	group.LeaderId = lid
	group.MemberId = mid
	group.Cid = cid
	err := global.Db.Create(&group).Error
	if err != nil {
		return group, 500
	}
	return group, 200
}

func SelectGroupById(id uint) ([]Group, int, int) {
	var groups []Group
	var total int
	err := global.Db.Find(&groups, "id=?", id).Count(&total).Error
	if err != nil || len(groups) == 0 {
		return groups, 0, 500
	}
	return groups, total, 200
}

func SelectGroupByLid(lid uint) ([]Group, int, int) {
	var groups []Group
	var total int
	err := global.Db.Find(&groups, "lid=?", lid).Count(&total).Error
	if err != nil || len(groups) == 0 {
		return groups, 0, 500
	}
	return groups, total, 200
}

func SelectGroupMid(mid uint) ([]Group, int, int) {
	var groups []Group
	var total int
	err := global.Db.Find(&groups, "mid=?", mid).Count(&total).Error
	if err != nil || len(groups) == 0 {
		return groups, 0, 500
	}
	return groups, total, 200
}
