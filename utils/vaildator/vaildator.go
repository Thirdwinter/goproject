package validator

import (
	"goproject/models"
	"goproject/utils/rspcode"
)

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
	// if err := validate.Var(user.Password, "required,min=6,max=20"); err != nil {
	// 	return "密码长度限制6-20位",rspcode.ERROR
	// }

	// if err := validate.Var(user.Role, "required,gte=2"); err != nil {
	// 	return "权限错误",rspcode.ERROR
	// }

	return "", rspcode.SUCCESS
}
