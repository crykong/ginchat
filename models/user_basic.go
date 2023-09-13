package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
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

// GetUserList 获取用户列表
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// FindUserByNameAndPwd 登录
func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and password = ?", name, password).First(&user)
	//token 加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5Encode(str)
	utils.DB.Model(&user).Where("id =?", user.ID).Update("identity", temp)
	return user
}

// FindUserByName 根据用户名称获取对象
func FindUserByName(name string) (UserBasic, error) {
	user := UserBasic{}
	result := utils.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

// CreateUser  创建用户
func CreateUser(user UserBasic) (*gorm.DB, error) {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

// DeleteUesr 删除用户
func DeleteUesr(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

// UpdateUser 更新用户对象
//
//	func UpdateUser(user UserBasic) *gorm.DB {
//		return utils.DB.Model(&user).Where("id = ?", user.ID).Updates(UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone})
//		//return utils.DB.Model(&UserBasic{}).Where("id = ?", user.ID).Updates(&user)
//	}
//
// UpdateUser 更新用户对象
func UpdateUser(user UserBasic) (*gorm.DB, error) {
	// 构建要更新的字段映射
	updates := map[string]interface{}{
		"name":     user.Name,
		"password": user.Password,
		"phone":    user.Phone,
		"email":    user.Email,
	}
	// 执行更新并获取结果
	result := utils.DB.Model(&UserBasic{}).Where("id = ?", user.ID).Updates(updates)

	if result.Error != nil {
		return nil, result.Error
	}

	return result, nil
}
