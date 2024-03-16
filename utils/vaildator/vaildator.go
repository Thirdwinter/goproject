package validator

import (
	"goproject/models"
	"goproject/utils/rspcode"
	"regexp"
)

// 使用正则表达式匹配邮箱格式
func validateEmail(email string) (bool,error) {
	// 使用正则表达式匹配邮箱格式
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regex, email)
	return match, err
}


func ValidateUserRegistration(user models.User) (string, int) {

	if len(user.Username) < 4 || len(user.Username) > 12 {
		return "用户名长度限制4-12位", rspcode.ERROR
	}

	if len(user.Password) < 4 || len(user.Password) > 12 {
		return "用户密码限制4-12位", rspcode.ERROR
	}

	if user.Role != 1 {
		return "用户无法申请更高权限", rspcode.ERROR
	}

	ok,err:=validateEmail(user.Email)
	if !ok || err != nil{
		return "用户邮箱错误", rspcode.ERROR
	}


	return "", rspcode.SUCCESS
}
