package v1

import (
	"goproject/models"
	_"goproject/service"
	"goproject/utils/rspcode"
	_"goproject/utils/vaildator"

	"github.com/gin-gonic/gin"
)

var code int
// 添加用户
func AddUser(c *gin.Context) {
	// var msg string
	var data models.User
	var headimg models.HeadImg
	_ = c.ShouldBindJSON(&data)
	_ = c.ShouldBindJSON(&headimg)
	// msg, code = validator.Validate(&data)
	// if code != rspcode.SUCCESS {
	// 	c.JSON(200, gin.H{
	// 		"code": code,
	// 		"msg":  msg,
	// 	})
	// 	return
	// }

	code = models.CheckUser(data.Username)
	if code == rspcode.SUCCESS {
		//!有问题,先不用
		//data.Image, _= service.UpLoadBase64Image(headimg.Imgstr)
		models.CreateUser(&data)
	}
	if code == rspcode.ERROR_USERNAME_USED {
		code = rspcode.ERROR_USERNAME_USED
	}
	c.JSON(200, gin.H{
		"code": code,
		//"data":   data,
		"msg": rspcode.GetMsg(code),
	})

}
