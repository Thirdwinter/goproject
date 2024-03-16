package v1

import (
	"fmt"
	"goproject/middleware"
	"goproject/models"

	//_ "goproject/service"
	"goproject/utils/rspcode"

	validator "goproject/utils/vaildator"

	"github.com/gin-gonic/gin"
)

var code int

// 添加用户
func AddUser(c *gin.Context) {
	var msg string
	var data models.User
	//var headimg models.HeadImg
	data.Username = c.PostForm("username")
	data.Password = c.PostForm("password")
	data.Role = 1
	data.Email=c.PostForm("email")

	// file, fileHeader, _ := c.Request.FormFile("file")
	// filesize := fileHeader.Size
	// data.Image, code = service.UpLoadFile(file, filesize)
	// c.JSON(200, gin.H{
	// 	"code": code,
	// 	"msg":  rspcode.GetMsg(code),
	// 	"url":  data.Image,
	// })

	// fmt.Println("body:" ,dd)
	//fmt.Println("abc:",abc)
	//_ = c.ShouldBind(&data)
	fmt.Printf("这里%#v\n", data)
	msg, code = validator.ValidateUserRegistration(data)
	//code = 200
	if code != rspcode.SUCCESS {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}
	//fmt.Println("yonghum: ", data.Username)
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

func Login(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	//fmt.Println(err)
	//fmt.Println("json:",data)
	var atoken string
	var rtoken string
	code, user := models.CheckLogin(data.Username, data.Password)
	//fmt.Println("user:",user)
	if code == rspcode.SUCCESS {
		fmt.Println("login,r:", user.Role)
		atoken, rtoken, code = middleware.SetToken(data.Username, user.Role)
		//atoken, _ = mdw2.SetToken(data.Username)
		//c.SetCookie("token", atoken, 3600, "/", "", false, true)

		// c.JSON(200, gin.H{
		// 	"code": 200,
		// 	"msg":  "Login successful",
		// })

		c.JSON(200, gin.H{
			"code":   code,
			"msg":    rspcode.GetMsg(code),
			"atoken": atoken,
			"rtoken": rtoken,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  rspcode.GetMsg(code),
		})
		c.Abort()
		return
	}
}
