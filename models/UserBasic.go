package models

import (
	"fmt"
	"heyChat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Salt          string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	HeartbeatTime time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	LogOutTime    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	fmt.Print(user)
	return user
}

//	func FindUserByNameAndPassword(name, password string) UserBasic {
//		user := UserBasic{}
//		utils.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
//		fmt.Println(user)
//		return user
//	}
func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Phone = ?", phone).First(&user)
}
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Email = ?", email).First(&user)
}
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email})
}
