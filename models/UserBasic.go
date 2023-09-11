package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `gorm:"名字"`
	Password      string `gorm:"密码"`
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClinetPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64 `gorm:"column:login_out_time" json:"login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string
}

// /创建表
func (table *UserBasic) TableName() string {
	return "User_basic"

}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
func CreateUesr(user UserBasic) *gorm.DB {
	return utils.DB.Create(user)
}

func DeleteUesr(user UserBasic) *gorm.DB {
	return utils.DB.Delete(user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(user)
}
