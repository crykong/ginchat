package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义一个DB
var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig() // 修复了错误变量名error -> err
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("配置文件 app:", viper.Get("app"))
	fmt.Println("配置文件 mysql:", viper.Get("mysql"))
}
func InitMysql() {
	//DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})

	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	dbname := viper.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	fmt.Println("Connected to MySQL database!")
}
