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
	Phone         string `valid:"matcher(^1[3-9]{1}\\d{9$})"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClinetPort    string
	Salt          string
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64 `gorm:"column:login_out_time" json:"login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string
}

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

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word=?", name, password).First(&user)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

func CreateUser(user UserBasic) (*gorm.DB, error) {
	fmt.Println(user.Name, user.Password)
	result := utils.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func DeleteUesr(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone})
}
