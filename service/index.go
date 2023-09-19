package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"text/template"
)

func GetIndex(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "welcom !!",
	//})
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
}

// ToRegister 用户注册
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	// 这一行代码使用一个模板引擎对象（假设是ind）的Execute方法，将user对象渲染到响应中。这通常用于生成动态的HTML或其他文本内容，
	//其中c.Writer表示响应的写入流，user对象是模板的数据上下文，模板引擎将根据模板文件和数据上下文生成最终的输出。
	ind.Execute(c.Writer, user)
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}
func Chat(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
