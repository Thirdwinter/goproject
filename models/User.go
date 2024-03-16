package models

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"goproject/global"
	"goproject/utils/rspcode"

	//"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

// type User struct {
// 	gorm.Model
// 	Username string `gorm:"type:varchar(20);not null" json:"username" binding:"required" validate:"required,min=4,max=12" label:"用户名"`
// 	Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required" validate:"required,min=6,max=20" label:"密码"`
// 	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
// 	Image    string `gorm:"type:text" label:"用户头像"`
// 	Salt     []byte `gorm:"type:varchar(20)" label:"m"`
// }

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" binding:"required" label:"用户名"`
	Password string `gorm:"type:varchar(50);not null" json:"password" binding:"required" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" label:"角色码"`
	Image    string `gorm:"type:text" label:"用户头像"`
	Email    string `gorm:"type:varchar(30);not null" json:"email" binding:"required" label:"用户邮箱"`
	Salt     string `gorm:"type:varchar(20)" label:"m"`
	Q        string `gorm:"type:varchar(30)"`
	A        string `gorm:"type:varchar(30)"`
}

type HeadImg struct {
	Imgstr string `json:"headimg"`
}

func generateRandomSalt() string {
	// 生成随机的8字节数据作为盐值
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		randomBytes = []byte{12, 3, 4, 66, 234, 11, 42, 90}
	}
	// 将随机字节数据转换为十六进制字符串
	salt := hex.EncodeToString(randomBytes)
	return salt
}

// 用户密码加密
func ScryptPw(password string, salt string) string {
	// 对密码和盐值分别进行单向哈希
	passwordHash := hashString(password)
	saltHash := hashString(salt)

	// 将密码哈希和盐值哈希拼接后再进行一次哈希
	combinedHash := hashString(passwordHash + saltHash)

	return combinedHash
}
func (u *User) BeforeSave() {
	u.Salt = generateRandomSalt()
	u.Password = ScryptPw(u.Password, u.Salt)
	if u.Image == "" {
		u.Image = "http://s98w22032.hb-bkt.clouddn.com/default.jpg"
	}
} //?gorm 自带

// 注册时查询用户是否存在
func CheckUser(username string) (code int) {
	var count int
	global.Db.Model(&User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
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

// 登录验证
func CheckLogin(username string, password string) (int, User) {
	var user User
	global.Db.Where("username=?", username).First(&user)
	//fmt.Println(username,password)
	//fmt.Println(user)
	if user.ID == 0 {
		return rspcode.ERROR_USER_NOT_EXIST, user
	}
	if password == "" {
		return rspcode.ERROR_PASSWORD_NO_EXIST, user
	}
	encryptedPassword := ScryptPw(password, user.Salt) // 对输入的密码进行加密

	if encryptedPassword != user.Password {
		return rspcode.ERROR_PASSWORD_WRONG, user
	}
	if user.Role != 1 && user.Role != 2 {
		return rspcode.ERROR_USER_NO_RIGHT, user
	}
	return rspcode.SUCCESS, user
}

func hashString(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
