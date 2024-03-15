package v1

import (
	"goproject/service"
	"goproject/utils/rspcode"
	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	filesize := fileHeader.Size
	url, code := service.UpLoadFile(file, filesize)
	c.JSON(200, gin.H{
		"code": code,
		"msg":  rspcode.GetMsg(code),
		"url":  url,
	})
}
