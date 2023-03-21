package models

type User struct {
	UserID       int    `gorm:"column:user_id"`
	UserOpenId   string `gorm:"column:user_openid"`
	Usergender   int    `gorm:"column:user_gender"`
	userAvatar   string `gorm:"column:user_avatar"`
	userNickname int    `gorm:"column:user_nickname"`
	userIsAdmin  int    `gorm:"column:user_is_admin"`
	userAllow    string `gorm:"column:user_allow"`
}

func (u *User) TableName() string {
	return "mango_user"
}

// GetUser 获取用户列表
func GetUser(NumPage, NumSize int) []User {
	NumPage = NumPage*NumSize - NumSize
	var u []User
	DB.Limit(NumSize).Offset(NumPage).Find(&u)
	return u
}

func DelUser(id int) bool {
	result := DB.Where("user_id = ?", id).Delete(&User{})
	if result.Error != nil {
		return false
	}
	return true
}
