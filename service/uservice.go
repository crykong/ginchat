package service

import (
	"fmt"
	"ginchat/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserList
// Summary 所有用户
// @Tags 首页
// @Accept json
// @Success 200 {string} json{"code","message"}
// @Router /user/getuserlist  [get]
func GetUserList(c *gin.Context) {

	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// Summary 新增用户
// @Tags 首页
// @param name query string false "用户名称"
// @param password query string false "密码"
// @param repassword query string false "确定密码"
// @Accept json
// @Success 200 {string} json{"code","message"}
// @Router /user/creatuser [get]
func CreateUser(pa *gin.Context) {
	user := models.UserBasic{}
	name := pa.Query("name")
	user.Name = name
	udata := models.FindUserByName(name)
	if udata != nil {
		pa.JSON(-1, gin.H{
			"message": "用户已经注册",
		})
	}
	password := pa.Query("password")
	repassword := pa.Query("repassword")
	if password != repassword {

		pa.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = password
	models.CreateUesr(user)
	pa.JSON(200, gin.H{
		"message": "用户创建成功",
	})

}

// DeleteUser
// Summary 删除用户
// @Tags 首页
// @param name query string false "用户名称"
// @param password query string false "密码"
// @param repassword query string false "确定密码"
// @Accept json
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteuser [get]
func DeleteUser(pa *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(pa.Query("id"))
	user.ID = uint(id)
	models.DeleteUesr(user)
	pa.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// Summary 修改用户
// @Tags 首页
// @param name query string false "用户名称"
// @param password query string false "密码"
// @param phone query string false "电话"
// @param email query string false "邮箱"
// @Accept json
// @Success 200 {string} json{"code","message"}
// @Router /user/updateteuser [post]
func UpdateUser(pa *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(pa.PostForm("id"))
	user.ID = uint(id)
	user.Name = pa.PostForm("name")
	user.Password = pa.PostForm("password")
	user.Phone = pa.PostForm("phone")
	user.Email = pa.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		pa.JSON(200, gin.H{
			"message": "修改参数不正确",
		})
		return
	} else {
		models.UpdateUser(user)
		pa.JSON(200, gin.H{
			"message": "修改用户和曾给",
		})
	}
}
