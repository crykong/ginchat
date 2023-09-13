package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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
	user, _ := models.FindUserByName(name)
	if user.Name == "" {
		r.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户不存在",
			"data":    data,
		})
	}
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		r.JSON(200, gin.H{
			"code":    -1,
			"message": "密码不正确",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	r.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    data,
	})

}

// CreateUser
// @Summary 新增用户
// @Tags 首页Api
// @Param         name query string false "用户名称"
// @Param         password query string false "密码"
// @Param         repassword query string false "确定密码"
// @Accept json
// @Success 200 {string} json{"code","message"}
// @Router /user/creatuser   [get]
func CreateUser(pa *gin.Context) {
	user := models.UserBasic{}
	name := pa.Query("name")
	udata, _ := models.FindUserByName(name) // 使用下划线来忽略第二个返回值

	if udata.Name != "" {
		pa.JSON(-1, gin.H{

			"code":    -1,
			"message": "用户已经注册",
			"data":    udata,
		})
		return
	}
	user.Name = name
	salt := fmt.Sprintf("%06d", rand.Int31()) // 修正 salt 的格式化
	password := pa.Query("password")
	repassword := pa.Query("repassword")
	if password != repassword {

		pa.JSON(-1, gin.H{
			"code":    -1,
			"message": "两次密码不一致",
			"data":    udata,
		})
		return
	}
	user.Password = password
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	dauser, _ := models.CreateUser(user)
	if dauser != nil {
		pa.JSON(200, gin.H{
			"code":    0,
			"message": "创建用户成功",
			"data":    dauser.Name,
		})
	} else {
		pa.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户创建失败",
			"data":    dauser.Name,
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
		"code":    0,
		"message": "删除用户成功",
		"data":    user,
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 首页Api
// @Param id query string false "用户名称"
// @Param name query string false "用户名称"
// @Param password query string false "密码"
// @Param phone query string false "电话"
// @Param email query string false "邮箱"
// @Accept json
// @Success 200 {string} json {"code","message"}
// @Router /user/updateteuser [post]
func UpdateUser(pa *gin.Context) {
	ubasic := models.UserBasic{}

	id, _ := strconv.Atoi(pa.PostForm("id"))
	fmt.Println("ID:", pa.PostForm("id"))
	ubasic.ID = uint(id)
	ubasic.Name = pa.PostForm("name")
	fmt.Println("Name:", pa.PostForm("name"))
	ubasic.Password = pa.PostForm("password")
	ubasic.Phone = pa.PostForm("phone")
	ubasic.Email = pa.PostForm("email")
	fmt.Println("Request Form Data:", pa.Request.Form)
	// 打印 ubasic 变量内容
	fmt.Printf("Updating user: %+v\n", ubasic)
	if err := pa.ShouldBind(&ubasic); err != nil {
		fmt.Println(err)
		pa.JSON(200, gin.H{
			"code":    -1,
			"message": "错误修改参数不正确",
			"data":    ubasic,
		})
		return
	}

	if _, err := govalidator.ValidateStruct(ubasic); err != nil {
		fmt.Println(err)
		pa.JSON(200, gin.H{
			"code":    -1,
			"message": "修改参数不正确",
			"data":    ubasic,
		})
		return
	}
	if _, err := models.UpdateUser(ubasic); err != nil {
		pa.JSON(200, gin.H{
			//"message": err.Error(),
			"code":    -1,
			"message": "修改失败异常",
			"data":    err.Error(),
		})
		return
	}

	pa.JSON(200, gin.H{
		"code":    0,
		"message": "修改成功",
		"data":    ubasic,
	})
}

//防止跨域占点 伪造请求

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Sendmsg  发送消息
func Sendmsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println("deferfunc", err)
		}
	}(ws)
	MsgHandler(ws, c)
}

// MsgHandler  处理消息
func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.Publishkey)
		if err != nil {
			fmt.Println("MsgHandLer 发送失败", err)
		}
		fmt.Println("发送消息", msg)
		tm := time.Now().Format("2006-01-12 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg) //ws//localhost:8080/user/sendmsg
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
