package utils

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("配置文件 app:", viper.Get("app"))
	fmt.Println("配置文件 mysql:", viper.Get("mysql"))
}

func InitMysql() {
	newlog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 修正这里的log.new -> log.New
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	dbname := viper.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newlog}) // 修正这里的logger -> Logger
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	fmt.Println("Connected to MySQL database!")
}
