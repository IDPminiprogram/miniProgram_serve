/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:21
 */

package models

import (
	"cea_api/pkg/jwt"
)

type User struct {
	ID         int
	Userid     string
	Realname   string `json:"realname"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Unit       string `json:"unit"`
	Department string `json:"department"`
	Usertype   string `json:"usertype"`
}

func (User) TableName() string {
	return "User"
}

type LoginUserRes struct {
	Realname   string
	Unit       string
	Department string
	Usertype   string
}

func GerUserCount(unit, department string) (data int) {
	userlist := []User{}
	DB.Where("unit = ? and department = ?", unit, department).Find(&userlist)
	data = len(userlist) + 1
	return data
}

func GetUserData(name, unit, department string) (userid, realname, usertype string) {
	var user User
	DB.Where("name = ? and unit = ? and department = ?", name, unit, department).Find(&user)
	if user.ID > 0 {
		userid := user.Userid
		realname := user.Realname
		usertype := user.Usertype
		return userid, realname, usertype
	} else {
		return
	}
}

// LoginCheck 登录检查
func LoginCheck(name, password, unit, department string) (code int) {
	var user User
	DB.Where("name = ? and unit = ? and department = ?", name, unit, department).Find(&user)
	if user.ID > 0 {
		if user.Password != password {
			return 20004
		} else {
			return 200
		}
	} else {
		return 20002
	}
}

// ValOldPwd  验证旧密码
func ValOldPwd(oldpwd, token string) (res bool) {
	var user User
	tokenvaldata, _ := jwt.ParseToken(token)
	UserId := tokenvaldata.UserId
	DB.Where("userid = ? ", UserId).Find(&user)
	if user.Password == oldpwd {
		return true
	} else {
		return false
	}
}

func UserUpdatePwd(newpwd, token string) (res bool) {
	userdata, _ := jwt.ParseToken(token)
	var user User
	DB.Where("userid = ?", userdata.UserId).First(&user)
	result := DB.Model(user).Where("userid = ?", userdata.UserId).Update("password", newpwd)
	if result.Error == nil {
		return true
	} else {
		return true
	}
}

// ValUserName 验证该用户名是否存在
func ValUserName(unit, department, username string) (res bool) {
	var u User
	DB.Where("unit = ? and department = ? and name = ?", unit, department, username).Find(&u)
	if u.ID > 0 {
		return false
	} else {
		return true
	}
}

// AddUser 新增用户
func AddUser(username, realname, password, usertype, unit, deaprtment, UserId string) (res bool) {
	result := DB.Create(&User{Userid: UserId, Name: username, Password: password, Realname: realname, Usertype: usertype, Unit: unit, Department: deaprtment})
	if result.Error != nil {
		return false
	}
	return true
}
