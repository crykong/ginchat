package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) //localhost:8080/swagger/index.html
	r.GET("/index", service.GetIndex)
	r.GET("/user/getuserlist", service.GetUserList)
	r.GET("/user/creatuser", service.CreateUser)

	r.GET("/user/deleteuser", service.DeleteUser)
	r.POST("/user/updateteuser", service.UpdateUser)

	r.POST("/user/findUserbynameandpwd", service.FindUserByNameAndPwd)

	return r
}
