// 鉴定用户权限
package middleware

import (
	"goproject/utils/rspcode"

	"github.com/gin-gonic/gin"
)

func CheckUserRole() gin.HandlerFunc{
	return func (c *gin.Context)  {
		role, exists :=c.Get("role")
		if !exists {
			c.JSON(200, gin.H{
				"code":rspcode.ERROR_USER_NO_RIGHT,
				"msg":"用户不正确",
			})
			c.Abort()
			return
		}
		if role == 1{
			c.Next()
			return
		} else{
			c.JSON(200, gin.H{
				"code":rspcode.ERROR_USER_NO_RIGHT,
				"msg":"用户权限不足",
			})
			c.Abort()
			return
		}
	}
}

func CheckAdminRole() gin.HandlerFunc{
	return func(c *gin.Context) {
		role, exists :=c.Get("role")
		if !exists {
			c.JSON(200, gin.H{
				"code":rspcode.ERROR_USER_NO_RIGHT,
				"msg":"用户不正确",
			})
			c.Abort()
			return
		}
		if role == 1 || role == 2{
			c.Next()
			return
		} else {
			c.JSON(200, gin.H{
				"code":rspcode.ERROR_USER_NO_RIGHT,
				"msg":"用户不正确",
			})
		}
	}
}