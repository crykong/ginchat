package utils

import (
	"context"
	"fmt"
	//"github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

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

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       2,
		PoolSize: 30,
	})
	//pong, err := Red.Ping().Result()
	//if err != nil {
	//	fmt.Println("init redis", err)
	//} else {
	//	fmt.Println("redis inited ...", pong)
	//}
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish 。。。。", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe 。。。。", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe 。。。。", msg.Payload)
	return msg.Payload, err
}
