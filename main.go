package main

import (
	"ginchat/router" // router   "ginchat/router"  // 请确保这里的路径正确，指向你的路由包

	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Router()
	r.Run()

}
