package models

import (
	"goproject/global"

	"github.com/jinzhu/gorm"
)

type Inform struct {
	gorm.Model
	FId    uint `form:"fid" binding:"required"` // 发起者id
	JId    uint `form:"jid" binding:"required"` // 接收者id
	CId    uint `form:"cid" binding:"required"` // 比赛id
	Ds     uint `form:"ds" binding:"required"`  // 1/2==>两种状态
	Stutas uint `form:"stutas"`                 // 1/2/3==>三种状态
}

// 发起/接收/比赛/动作==>新增信息
func CreateNewInf(fid, jid, cid, ds uint) (Inform, int) {
	var inf Inform
	inf.JId = jid
	inf.FId = fid
	inf.CId = cid
	inf.Ds = ds
	inf.Stutas = 1
	err := global.Db.Create(&inf).Error
	if err != nil {
		return Inform{}, 500
	}
	return inf, 200
}

// 更改状态(id,状态)
func ChangeStatus(iid, status uint) (Inform, int) {
	var inf Inform
	err := global.Db.First(&inf, "iid = ?", iid).Error
	if err != nil {
		return Inform{}, 500
	}
	inf.Stutas = status
	global.Db.Save(inf)
	return inf, 200
}

// 删除通知
func DeleteInfo(iid uint) int {
	var inf Inform
	err := global.Db.Delete(&inf, "id = ?", iid).Error
	if err != nil {
		return 500
	}
	return 200
}

// 查询通知(id)
func SelectInfoById(iid uint) ([]Inform, int, int) {
	var infs []Inform
	var total int
	global.Db.Find(&infs, "id = ?", iid).Count(&total)
	return infs, total, 200
}

// 查询通知(fid)
func SelectInfoByFid(fid uint) ([]Inform, int, int) {
	var infs []Inform
	var total int
	global.Db.Find(&infs, "fid = ?", fid).Count(&total)
	return infs, total, 200
}

// 查询通知(status)
func SelectInfoByStatus(status uint) ([]Inform, int, int) {
	var infs []Inform
	var total int
	global.Db.Find(&infs, "status = ?", status).Count(&total)
	return infs, total, 200
}

// 查询通知(jid)
func SelectInfoByJid(jid uint) ([]Inform, int, int) {
	var infs []Inform
	var total int
	global.Db.Find(&infs, "jid = ?", jid).Count(&total)
	return infs, total, 200
}
