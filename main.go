package main

import (
	"ginchat/models"
	"ginchat/router" // router   "ginchat/router"  // 请确保这里的路径正确，指向你的路由包
	"github.com/spf13/viper"
	"time"

	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	InitTimer()
	r := router.Router()                  // router.Router()
	r.Run(viper.GetString("port.server")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

// 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
