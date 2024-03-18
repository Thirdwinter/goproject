package v1

import (
	"goproject/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 查询指定状态(true/false)的小组申请记录(已通过或者被驳回)
// -->返回切片
// form提交一个 "rec_status"(true/false)
func SelectStatus(c *gin.Context) {
	ok := c.PostForm("rec_status")
	ok = strings.ToLower(ok)
	var stutas bool
	if ok == "true" {
		stutas = true
	}
	if ok == "false" {
		stutas = false
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	rec, total, code := models.SelectStatus(stutas)
	if code == 500 {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "查询失败",
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"code":  code,
		"data":  rec,
		"total": total,
	})
}

// 新增一个记录
// form提交 "uid"(发起比赛招募申请的用户id);
// "cid"(对那种比赛发布组队申请)
// 返回创建的组队申请记录
func CreateRec(c *gin.Context) {
	uuid := c.PostForm("uid")
	ccid := c.PostForm("cid")
	if uuid == "" || ccid == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	uid, err := strconv.Atoi(uuid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	creuid := uint(uid)
	cid, err := strconv.Atoi(ccid)
	if err != nil {

		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	crecid := uint(cid)
	rec, code := models.CreateRecruitment(crecid, creuid)
	if code == 500 {
		c.JSON(500, gin.H{
			"code": code,
			"msg":  "创建错误",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": rec,
	})
}

// 管理员审核一条待审核记录
// form提交"rid"(申请记录的id)
func Reply(c *gin.Context){
	rrid:=c.PostForm("rid")
	if rrid == "" {
		c.JSON(400, gin.H{
			"code":400,
			"msg":"参数错误",
		})
		c.Abort()
		return
	}
	irid,err:=strconv.Atoi("rrid")
	if err != nil{
		c.JSON(400, gin.H{
			"code":400,
			"msg":"参数错误",
		})
		c.Abort()
		return
	}
	rec,code:=models.Reply(uint(irid), false)
	if code == 500 {
		c.JSON(500, gin.H{
			"code":500,
			"msg":"错误",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code":200,
		"msg":"ok",
		"data":rec,
	})
}

// 用户删除记录
// form提交"uid","rid"
func Undo(c *gin.Context){
	uuid := c.PostForm("uid")
	ccid := c.PostForm("cid")
	if uuid == "" || ccid == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	uid, err := strconv.Atoi(uuid)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	creuid := uint(uid)
	cid, err := strconv.Atoi(ccid)
	if err != nil {

		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		c.Abort()
		return
	}
	crecid := uint(cid)
	code:=models.Undo(crecid, creuid)
	if code == 500{
		c.JSON(500, gin.H{
			"code":500,
			"msg":"删除错误",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code":200,
		"msg":"ok",
	})
}

