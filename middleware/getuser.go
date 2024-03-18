package middleware

import (
	"goproject/global"
	"goproject/models"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userlogin := c.GetHeader("userlogin")
		var u models.User
		if userlogin == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "未授权",
			})
			c.Abort()
			return
		}
		err := global.Db.First(&u, "username = ?", userlogin).Error
		if u.ID <=0||err!=nil{
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "未授权",
			})
			c.Abort()
			return
		}

		c.Set("username", userlogin)
		c.Set("uid", u.ID)
		c.Next()
	}
}
