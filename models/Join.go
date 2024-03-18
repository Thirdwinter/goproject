// 与用户对比赛招募的自我推荐有关

package models

import (
	"goproject/global"

	"github.com/jinzhu/gorm"
)

type Joins struct {
	gorm.Model
	UserId        uint
	RecruitmentId uint
	Status        bool
}

// 指定状态检查某个用户的所有申请
func SelectJoins(uid uint, status bool) ([]Joins, int) {
	var joins []Joins
	err := global.Db.Find(&joins, "userid = ? and stutus = ", uid, status).Error
	if err != nil {
		return joins, 500
	}
	return joins, 200
}

// 检查插入一个申请是否合法
func CheckJoin(uid uint, rid uint) bool {
	var joins Joins
	var total int
	err := global.Db.Find(&joins, "userid = ? and recruitmentid = ? and status = ture", uid, rid).Count(&total).Error
	if err != nil || total != 0 {
		return false
	}
	return true
}

// 新增一个加入申请
func CreateJoins(uid uint, rid uint) (Joins, int) {
	if CheckJoin(uid, rid) {
		var join Joins = Joins{UserId: uid, RecruitmentId: rid, Status: true}
		err := global.Db.Create(&join).Error
		if err == nil {
			return join, 200
		}
	}
	return Joins{}, 500
}

// 删除一个存在自荐申请
func ReturnJoins(uid uint, rid uint) int {
	var join Joins

	// 查询符合条件的申请记录
	err := global.Db.Where("userid = ? AND recruitmentid = ?", uid, rid).First(&join).Error
	if err != nil {
		// 未找到符合条件的申请记录
		return 500
	}

	// 删除申请记录
	err = global.Db.Delete(&join).Error
	if err != nil {
		// 删除失败
		return 500
	}

	return 200
}

// 查询一条记录下所有的自荐
func SelectThisJoins(rid uint)([]Joins,int){
	var joins []Joins
	err:=global.Db.Find(joins, "recruitmentid = ? and status = true",rid).Error
	if err!=nil{
		return joins,500
	}
	return joins,200
}