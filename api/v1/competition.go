package v1

import (
	"goproject/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCompetition(c *gin.Context) {
	var data models.Competition
	data.Title = c.PostForm("title")
	data.Info = c.PostForm("info")
	if data.Title == "" || data.Info == "" {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "存在空值,重新提交",
		})
		c.Abort()
		return
	}
	code := models.CheckCompetition(&data)
	if code != 200 {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  "赛事已存在",
		})
		c.Abort()
		return
	}
	code = models.CreateCompetition(&data)
	if code != 200 {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  "添加赛事失败",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加赛事成功",
	})
}

func UpdateCompetition(c *gin.Context) {
	title := c.PostForm("title")
	ntitle := c.PostForm("ntitle")
	ninfo := c.PostForm("ninfo")

	// 参数校验
	if title == "" || ntitle == "" || ninfo == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "标题和信息不能为空",
		})
		return
	}

	code := models.UpdateCompetition(title, ntitle, ninfo)
	if code != 200 {
		c.JSON(200, gin.H{
			"code": code,
			"msg":  "更新失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg":  "更新成功",
	})
}

func DelCompetiton(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "请求错误,title值为空",
		})
		return
	}
	code := models.DelCompetiton(title)
	if code != 200 {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func SelectPageCompetiton(c *gin.Context) {
	pagesize, err := strconv.Atoi(c.Param("pagesize"))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "分页参数错误",
		})
		return
	}
	pagenum, err := strconv.Atoi(c.Param("pagenum"))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "分页参数错误",
		})
		return
	}
	com, total := models.SelectPage(pagesize, pagenum)
	if total == 0 {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "查询错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "查询成功",
		"data":  com,
		"total": total,
	})
}
