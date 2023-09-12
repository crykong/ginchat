package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:aNXvf&qYQ4NJTt7DPzPdGoGg@tcp(127.0.0.1:3306)/ginchat?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("数据库链接失败")
	}
	db.AutoMigrate(&models.UserBasic{})

	//Create
	//user := &models.UserBasic{}
	//user.Name = "王五"
	////	insertProduct := &Product{Code: "D42", Price: 100}
	//
	//db.Create(user)
	//fmt.Printf("insert ID: %d, Name: %s \n", user.ID, user.Name)
	//
	////red
	//db.First(user, user.ID)
	//
	//db.Model(user).Update("Password", "aNXvf&qYQ4NJTt7DPzPdGoGg")
	//
	//fmt.Printf("Update ID: %d,  Password: %d \n", user.ID, user.Password)

}
