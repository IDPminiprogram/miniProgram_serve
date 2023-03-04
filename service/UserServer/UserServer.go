/**
 * @Author: 戈亓
 * @Description:
 * @File:  UserServer
 * @Version: 1.0.0
 * @Date: 2022/7/24 20:33
 */

package UserServer

import (
	"cea_api/models"
	"cea_api/pkg/jwt"
	"cea_api/service/UnitsServer"
	"fmt"
	"strconv"
)

type User struct {
	Userid     string
	Realname   string
	Name       string
	Password   string
	Unit       string
	Department string
	Usertype   string
}

func (u *User) AddUser() (res bool) {
	UnitMark := UnitsServer.GetUnitMark(u.Unit)                   // 单位缩写
	DepartSer := UnitsServer.GetDepartSer(UnitMark, u.Department) // 部门序号
	UserCount := GetUserCount(u.Unit, u.Department)
	UserId := UnitMark + SwapIntFormat(DepartSer) + SwapIntFormat(UserCount)
	return models.AddUser(u.Name, u.Realname, u.Password, u.Usertype, u.Unit, u.Department, UserId)
}

func SwapIntFormat(data int) (res string) {
	if data < 10 {
		res = "0" + strconv.Itoa(data) // int类型 转 字符串
		return
	} else {
		res = strconv.Itoa(data)
		return
	}
}

func GetUserCount(unit, department string) (data int) {
	return models.GerUserCount(unit, department)
}
func (u *User) ValUserExists() (res bool) {
	return models.ValUserName(u.Unit, u.Department, u.Name)
}

func ValOldPwd(oldpwd, token string) (res bool) {
	return models.ValOldPwd(oldpwd, token)
}

func UpdateUserPwd(newpwd, token string) (res bool) {
	return models.UserUpdatePwd(newpwd, token)
}

type UserLogin struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Unit       string `json:"unit"`
	Department string `json:"department"`
}

type LoginRes struct {
	Name       string
	Realname   string `json:"realname"`
	Usertype   string `json:"usertype"`
	Unit       string `json:"unit"`
	Department string `json:"department"`
	Token      string `json:"token"`
}

func (u *UserLogin) LoginCheck() (code int) {
	return models.LoginCheck(u.Name, u.Password, u.Unit, u.Department)
}

func (l *LoginRes) CreatToken() (res interface{}) {

	userid, realname, usertype := models.GetUserData(l.Name, l.Unit, l.Department)
	token_ := jwt.TokenData{Realname: realname, UserType: usertype, Unit: l.Unit, Department: l.Department, UserId: userid}
	token, _ := token_.GenToken(userid, realname, usertype, l.Unit, l.Department)
	result := LoginRes{Name: l.Name, Unit: l.Unit, Department: l.Department, Realname: realname, Usertype: usertype, Token: token}
	res = result
	return res
}

func GetMenu(userType string) (res []models.Menu) {
	fmt.Println(userType)
	model := models.UserType{UserType: userType}
	res = model.GetUsertypeID()
	fmt.Println(res)
	return res
}
