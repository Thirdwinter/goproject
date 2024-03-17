package v1

import (
	"goproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompetition(c *gin.Context) {
	var data models.Competition
	data.Title = c.PostForm("title")
	data.Info = c.PostForm("info")
	if data.Title == ""|| data.Info == ""{
		c.JSON(200, gin.H{
			"code":500,
			"msg":"存在空值,重新提交",
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
	ntitle := c.PostForm("ntitle")
	ninfo := c.PostForm("ninfo")

	// 参数校验
	if ntitle == "" || ninfo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "标题和信息不能为空",
		})
		return
	}

	code := models.UpdateCompetition(ntitle, ninfo)
	if code != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg":  "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "更新成功",
	})
}
