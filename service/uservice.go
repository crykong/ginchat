package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} welcom
// @Router /user/getuserList  [get]
func GetUserList(c *gin.Context) {

	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}
