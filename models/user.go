package models

import (
	"fmt"
	"time"
)

type User struct {
	UserID        int       `gorm:"column:user_id"`
	UserOpenId    string    `gorm:"column:user_openid"`
	Usergender    int       `gorm:"column:user_gender"`
	UserAvatar    string    `gorm:"column:user_avatar"`
	UserNickname  string    `gorm:"column:user_nickname"`
	UserIsAdmin   int       `gorm:"column:user_is_admin"`
	UserAllow     string    `gorm:"column:user_allow"`
	UserCreatTime time.Time `gorm:"column:user_create_time"`
}

func (u *User) TableName() string {
	return "mango_user"
}

// GetUser 获取用户列表
func GetUser(NumPage, NumSize int) (user []User, count int) {
	var u []User
	NumPage = NumPage*NumSize - NumSize
	DB.Limit(NumSize).Offset(NumPage).Find(&u)
	result := DB.Find(&[]User{})
	count = int(result.RowsAffected)
	fmt.Println(count)
	//fmt.Println(result.RowsAffected)
	return u, count
}

func DelUser(id int) bool {
	result := DB.Where("user_id = ?", id).Delete(&User{})
	if result.Error != nil {
		return false
	}
	return true
}

func AllowUser(id, allow int) bool {
	result := DB.Model(&User{}).Where("user_id = ?", id).Update("user_is_admin", allow)
	if result.Error != nil {
		return false
	}
	return true
}
func AllowRelease(id int, allow string) bool {
	result := DB.Model(&User{}).Where("user_id = ?", id).Update("user_allow", allow)
	if result.Error != nil {
		return false
	}
	return true
}

func SearchUser(name string) []User {
	var u []User
	DB.Where("user_nickname LIKE ?", name).Find(&u)
	return u
}
