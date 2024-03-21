package v1

import (
	"fmt"
	"goproject/middleware"
	"goproject/models"
	"goproject/service"
	"goproject/utils/rspcode"
	validator "goproject/utils/vaildator"
	"strconv"

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
	data.Phonenumber = c.PostForm("phonenumber")
	data.Email = c.PostForm("email")
	data.Role = 1
	if data.Username==""||data.Password==""||data.Phonenumber==""||data.Email==""{
		c.JSON(400, gin.H{
			"code":400,
			"msg":"用户信息上传错误",
		})
		return
	}
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
		file, fileHeader, err := c.Request.FormFile("file")
		if err!=nil{
			c.JSON(400, gin.H{
				"code":400,
				"msg":"图片上传错误",
			})
			return
		}
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
	//_ = c.ShouldBindJSON(&data)
	data.Username = c.PostForm("username")
	data.Password = c.PostForm("password")
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
			"username":data.Username,
			"uid":user.ID,
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
	username:=c.PostForm("username")
	if username == ""{
		c.JSON(400, gin.H{
			"code":400,
		})
		return
	}
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
	// who, exists := c.Get("username")
	// name := who.(string)
	// if !exists {
	// 	c.JSON(200, gin.H{
	// 		"code": 500,
	// 		"msg":  "更换头像失败,无法获取用户权限",
	// 	})
	// 	c.Abort()
	// 	return
	// }
	code, user := models.UpdateUserImage(username, url)
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

// form提交"phone"
func SelectIdByPhone(c *gin.Context){
	phone:=c.PostForm("phone")
	if phone == ""{
		c.JSON(400, gin.H{
			"code":400,
			"msg":"参数错误",
		})
		c.Abort()
		return
	}
	uid,code:=models.SelectIdByPhone(phone)
	if code==500{
		c.JSON(500, gin.H{
			"code":500,
			"msg":"查询错误",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code":200,
		"msg":"ok",
		"data":uid,
	})
}

// form提交"uid"
func SelectUserDataById(c *gin.Context){
	uid:=c.PostForm("uid")
	if uid == ""{
		c.JSON(400, gin.H{
			"code":400,
			"msg":"参数错误",
		})
		c.Abort()
		return		
	}
	id,err:=strconv.Atoi(uid)
	if err!= nil{
		c.JSON(400, gin.H{
			"code":400,
			"msg":"参数错误",
		})
		c.Abort()
		return
	}
	userinfo,code:=models.SelectUserDataById(uint(id))
	if code != 200{
		c.JSON(500, gin.H{
			"code":500,
			"msg":"查询失败",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code":200,
		"msg":"ok",
		"data":userinfo,
	})
}