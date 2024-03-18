package v1

import (
	"goproject/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询一条记录下的所有自我推荐
// form提交"rid"
// 返回切片
func SelectThisJoins(c *gin.Context) {
	rrid := c.PostForm("rid")
	if rrid == "" {
		c.JSON(400, gin.H{
			"cdoe": 200,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	rid, err := strconv.Atoi(rrid)
	if err != nil {
		c.JSON(400, gin.H{
			"cdoe": 200,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	joins, code := models.SelectThisJoins(uint(rid))
	if code != 200 {
		c.JSON(500, gin.H{
			"cdoe": 500,
			"msg":  "查询错误",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": joins,
	})
}
