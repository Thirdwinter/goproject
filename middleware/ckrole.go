// 鉴定用户权限
package middleware

import (
	"fmt"
	"goproject/utils/rspcode"

	"github.com/gin-gonic/gin"
)

func CheckUserRole() gin.HandlerFunc{
	return func (c *gin.Context)  {
		role, exists :=c.Get("role")
		fmt.Println("r:",role)
		if !exists {
			c.JSON(200, gin.H{
				"code":rspcode.ERROR_USER_NO_RIGHT,
				"msg":"用户不正确",
			})
			c.Abort()
			return
		}
		userRole := role.(int)
		if userRole ==1||userRole ==2{
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
		fmt.Println("r:",role)
		if !exists {
			c.JSON(200, gin.H{
				"code":rspcode.ERROR_USER_NO_RIGHT,
				"msg":"用户不正确",
			})
			c.Abort()
			return
		}
		adminRole:=role.(int)
		if adminRole ==2{
			c.Next()
			return
		} else {
			c.JSON(200, gin.H{
				"code":rspcode.ERROR_USER_NO_RIGHT,
				"msg":"用户不正确",
			})
			c.Abort()
			return
		}
	}
}