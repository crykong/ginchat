package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// GetUserList
// @Summary 所有用户
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

// FindUserByNameAndPwd
// @Summary 登录
// @Tags 首页Api
// @Accept json
// @Param    name query string false "用户名称"
// @Param    password query string false "密码"
// @Success 200 {string} json {"code","message"}
// @Router /user/findUserbynameandpwd  [post]
func FindUserByNameAndPwd(r *gin.Context) {

	data := models.UserBasic{}
	name := r.Query("name")
	password := r.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		r.JSON(200, gin.H{
			"message": "该用户不存在",
		})
	}
	retsut := utils.ValidPassword(password, user.Salt, user.Password)
	if !retsut {
		r.JSON(200, gin.H{
			"message": "密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	r.JSON(200, gin.H{
		"message": data,
	})

}

// CreateUser
// @Summary 新增用户
// @Tags 首页Api
// @Param         name query string false "用户名称"
// @Param         password query string false "密码"
// @Param         repassword query string false "确定密码"
// @Accept json
// @Success 200 {string} json {"code","message"}
// @Failure 200 {object} errors.Error
// @Router /user/creatuser    [get]
func CreateUser(pa *gin.Context) {
	user := models.UserBasic{}
	name := pa.Query("name")
	user.Name = pa.Query("name")
	udata := models.FindUserByName(name)
	if udata.Name != "" {
		pa.JSON(-1, gin.H{
			"message": "用户已经注册",
		})
	}
	salt := fmt.Sprint("%06", rand.Int31())
	password := pa.Query("password")
	repassword := pa.Query("repassword")
	if password != repassword {

		pa.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = password
	user.Password = utils.MakePassword(password, salt)
	createdUser, err := models.CreateUser(user)
	if err != nil {
		pa.JSON(-1, gin.H{
			"message": err,
		})
	} else {
		pa.JSON(200, gin.H{
			"message": "用户创建成功",
			"user":    createdUser.Name,
		})
	}

}

// DeleteUser
// @Summary 删除用户
// @Tags 首页Api
// @Param id query string false "用户名称"
// @Accept json
// @Success 200 {string} json {"code","message"}
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
// @Summary 修改用户
// @Tags 首页Api
// @Param name query string false "用户名称"
// @Param password query string false "密码"
// @Param phone query string false "电话"
// @Param email query string false "邮箱"
// @Accept json
// @Success 200 {string} json {"code","message"}
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
