package service

import "github.com/gin-gonic/gin"

//GetIndex
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} welcom
// @Router /index  [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcom !!",
	})
}
