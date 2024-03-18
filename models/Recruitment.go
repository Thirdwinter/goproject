// 与比赛招募有关

package models

import (
	"goproject/global"
	"time"

	"github.com/jinzhu/gorm"
)

type Recruitment struct {
	gorm.Model
	UserId        uint
	CompetitionId uint
	Status        bool
	Vtime         int64
}

// 检查用户与比赛的合法性
func CheckIDsExist(uid, cid uint) bool {
	var userCount, comCount int

	// 查询 user 表中是否存在指定的 cid
	global.Db.Model(&User{}).Where("id = ?", uid).Count(&userCount)

	// 查询 com 表中是否存在指定的 cid
	global.Db.Model(&Competition{}).Where("id = ?", cid).Count(&comCount)

	return userCount > 0 && comCount > 0
}

// 查询一个用户是否存在一个已被批准且未过期的招募申请记录
func SelectOk(uid uint, cid uint) bool {
	var rec Recruitment
	global.Db.First(&rec, "userid = ? and competitionid = ?", uid, cid)
	return !(rec.Status && rec.Vtime > time.Now().Unix())
}

// 查询指定状态的记录
func SelectStatus(status bool) ([]Recruitment, int, int) {
	var recs []Recruitment
	var total int
	err := global.Db.Find(&recs, "status = ?", status).Count(&total).Error
	if err != nil {
		return recs, 0, 500
	}
	return recs, total, 200
}

// 新增一条记录
func CreateRecruitment(uid, cid uint) (Recruitment, int) {
	// 用户与比赛身份合法且未有相同记录
	if CheckIDsExist(uid, cid) && !SelectOk(uid, cid) {
		var rec Recruitment = Recruitment{UserId: uid, CompetitionId: cid, Status: false}
		err := global.Db.Create(&rec).Error
		if err == nil {
			return rec, 200
		}
	}
	return Recruitment{}, 500
}

// 管理员审核一条记录
func Reply(rid uint, ok bool) (Recruitment, int) {
	var rec Recruitment
	global.Db.First(&rec, "id=?", rid)
	if ok {
		rec.Status = ok
		rec.Vtime = time.Now().Add(7 * 24 * time.Hour).Unix()
		err := global.Db.Save(&rec).Error
		if err == nil {
			return rec, 200
		}
	}
	return rec, 500
}

// 用户撤销一条记录
func Undo(uid uint, rid uint) (int){
	recs, total, code := SelectStatus(true)
	if total == 1 && code == 200 && recs[0].ID == rid && recs[0].UserId == uid && recs[0].Status {
		recs[0].Status = false
		global.Db.Delete(&recs[0])
		return 200
	}
	return 500
}

