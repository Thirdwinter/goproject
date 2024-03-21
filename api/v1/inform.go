package v1

import (
	"goproject/models"

	"github.com/gin-gonic/gin"
)

func NewInform(c *gin.Context) {
	var inf models.Inform
	err := c.ShouldBind(&inf)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
		})
		return
	}
	inf, code := models.CreateNewInf(inf.FId, inf.JId, inf.CId, inf.Ds)
	if code != 200 {
		c.JSON(500, gin.H{
			"code": 500,
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": inf,
	})
}
