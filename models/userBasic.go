package models

import (
	"chat/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	HashPassword  string
	Salt          string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"` //valid校验
	Email         string `valid:"email"`                      //valid校验
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     string
	HeartbeatTime string
	LogOutTime    string
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	user := make([]*UserBasic, 10)
	utils.DB.Find(&user)
	return user
}

func UserCreate(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func UserDelete(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UserUpdataPhoneAndEmail(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Phone: user.Phone, Email: user.Email})
}

func FindUserByName(name string) UserBasic {
	var user UserBasic
	utils.DB.Model(&UserBasic{}).Where("name = ? ", name).First(&user)
	return user
}

func FindUserByPhone(phone string) UserBasic {
	var user UserBasic
	utils.DB.Model(&UserBasic{}).Where("phone = ? ", phone).First(&user)
	return user
}

func LoginUpdataUserInfo(user UserBasic) *gorm.DB {
	user.Identity = utils.MD5Encode(fmt.Sprintf("%v", time.Now().Unix()))
	newUser := UserBasic{
		Identity:  user.Identity,
		LoginTime: time.Now().Format("2006-01-02 15:03:04"),
	}
	return utils.DB.Model(&user).Updates(newUser)
}

func UpdataPasswordById(id uint, password string) *gorm.DB {
	var user UserBasic
	utils.DB.Where("id=?", id).First(&user)
	user.Salt = utils.MakeRandomSalt(16)
	user.HashPassword = utils.EncryptionPassword(user.Salt, password)
	return utils.DB.Model(user).Updates(user)

}
