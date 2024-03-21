package v1

import (
	"fmt"
	"goproject/models"

	"github.com/gin-gonic/gin"
)

// 队长id(lid),队员id(mid),比赛id(cid)==>新增信息
func CreateGroup(c *gin.Context) {
	var group models.Group
	err := c.ShouldBind(&group)
	if err != nil {
		c.JSON(400, gin.H{
			"code": "400",
		})
		return
	}
	data, code := models.CreateGroup(group.LeaderId, group.Cid, group.MemberId)
	if code != 200 {
		c.JSON(500, gin.H{
			"code": 500,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

func SelectGroupById(c *gin.Context) {
	var group models.Group
	err := c.ShouldBind(&group)
	if err != nil || group.ID == 0 {
		c.JSON(400, gin.H{
			"code": 400,
		})
		return
	}
	fmt.Println(group.ID)
	groups, total, code := models.SelectGroupById(group.ID)
	if code != 200 {
		c.JSON(500, gin.H{
			"code": 500,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  groups,
		"total": total,
	})
}
func SelectGroupByLid(c *gin.Context) {
	var group models.Group
	err := c.ShouldBind(&group)
	if err != nil || group.LeaderId == 0 {
		c.JSON(400, gin.H{
			"code": 400,
		})
		return
	}
	fmt.Println(group.ID)
	groups, total, code := models.SelectGroupById(group.LeaderId)
	if code != 200 {
		c.JSON(500, gin.H{
			"code": 500,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  groups,
		"total": total,
	})
}
func SelectGroupByMid(c *gin.Context) {
	var group models.Group
	err := c.ShouldBind(&group)
	if err != nil || group.MemberId == 0 {
		c.JSON(400, gin.H{
			"code": 400,
		})
		return
	}
	fmt.Println(group.ID)
	groups, total, code := models.SelectGroupById(group.MemberId)
	if code != 200 {
		c.JSON(500, gin.H{
			"code": 500,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  groups,
		"total": total,
	})
}
