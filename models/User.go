package models

import (
	"encoding/base64"
	"goproject/global"
	"goproject/utils/rspcode"

	_ "github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" binding:"required" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required" validate:"required,min=0,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
	Image    string `gorm:"type:text;default:'http://s98w22032.hb-bkt.clouddn.com/default.jpg'" json:"img" label:"用户头像"`
}

type HeadImg struct {
	Imgstr string `json:"headimg"`
}

// 用户密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := []byte{12, 3, 4, 66, 234, 11, 42, 90}
	HashPwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		//!日后处理
		//log.Fatal("加盐加密错误：%s", err)
		return ""
	}
	fpw := base64.StdEncoding.EncodeToString(HashPwd)
	return fpw
}
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
	if u.Image == "" {
		u.Image = "http://s98w22032.hb-bkt.clouddn.com/default.jpg"
	}
} //?gorm 自带

// 注册时查询用户是否存在&&默认头像
func CheckUser(username string) (code int) {
	var users User
	global.Db.Select("id").Where("username=?", username).First(&users)
	if users.ID > 0 {
		return rspcode.ERROR_USERNAME_USED
	}
	return rspcode.SUCCESS
}

// 新增用户
func CreateUser(data *User) (code int) {
	//data.Password =ScryptPw(data.Password)
	err := global.Db.Create(&data).Error
	if err != nil {
		//log.Error("create user error: %s", err)
		return rspcode.ERROR
	}
	return rspcode.SUCCESS
}
