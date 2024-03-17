package v1

import (
	"fmt"
	"goproject/middleware"
	"goproject/models"
	"goproject/service"
	"goproject/utils/rspcode"
	validator "goproject/utils/vaildator"

	"github.com/gin-gonic/gin"
)

//var code int

// 添加用户
func AddUser(c *gin.Context) {
	var msg string
	var data models.User
	//var headimg models.HeadImg
	data.Username = c.PostForm("username")
	data.Password = c.PostForm("password")
	data.Role = 1
	data.Email = c.PostForm("email")

	//fmt.Printf("这里%#v\n", data)
	msg, code := validator.ValidateUserRegistration(data)
	//code = 200
	if code != rspcode.SUCCESS {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}
	code = models.CheckUser(data.Username)
	if code == rspcode.SUCCESS {
		file, fileHeader, _ := c.Request.FormFile("file")
		filesize := fileHeader.Size
		data.Image, code = service.UpLoadFile(file, filesize)
		if code != 200 {
			data.Image = "http://s98w22032.hb-bkt.clouddn.com/default.jpg"
			return
		}
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

// 用户登录
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

// 修改头像
func UpdateUserImage(c *gin.Context) {
	newfile, fileHeader, _ := c.Request.FormFile("newfile")
	filesize := fileHeader.Size
	url, code := service.UpLoadFile(newfile, filesize)
	if code != 200 {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "更换头像失败,上传头像失败",
		})
		c.Abort()
		return
	}
	who, exists := c.Get("username")
	name := who.(string)
	if !exists {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "更换头像失败,无法获取用户权限",
		})
		c.Abort()
		return
	}
	code, user := models.UpdateUserImage(name, url)
	if code == 200 {
		c.JSON(200, gin.H{
			"code": 200,
			"data": user.Image,
			"msg":  "更新头像成功",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "更新头像失败",
		})
		c.Abort()
		return
	}
}

// 修改个人信息
// !先扔着
func UpdateUserInfo(c *gin.Context) {
}
