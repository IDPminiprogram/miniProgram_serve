/**
 * @Author: 戈亓
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/22 2:25
 */

package UserController

import (
	"cea_api/pkg/app"
	"cea_api/pkg/e"
	"cea_api/pkg/jwt"
	"cea_api/service"
	"cea_api/service/UserServer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetVerCode 获取验证码
func GetVerCode(c *gin.Context) {
	app := app.Gin{c}
	res := service.GetCaptcha()
	app.Response(200, e.Success, res)
}

// ValVerCode 验证验证码
func ValVerCode(c *gin.Context) {
	app := app.Gin{c}
	id := c.Query("codeid")
	ret_captcha := c.Query("vercode")
	if ret_captcha == "" {
		app.Response(http.StatusOK, e.Valerror, nil)
	} else {
		res := service.VerityCaptcha(id, ret_captcha)
		if res == false {
			app.Response(http.StatusOK, e.Valerror, nil)
		} else {
			app.Response(http.StatusOK, e.Valright, nil)
		}
	}
}

// Login 登录
func Login(c *gin.Context) {
	app := app.Gin{c}
	username := c.Query("username")
	password := c.Query("password")
	unit := c.Query("unit")
	department := c.Query("department")
	user := UserServer.UserLogin{username, password, unit, department}
	code := user.LoginCheck()
	switch code {
	case 200:
		result := UserServer.LoginRes{Name: username, Department: department, Unit: unit}
		res := result.CreatToken()
		app.Response(http.StatusOK, e.Success, res)
	case 20004:
		app.Response(http.StatusOK, e.UserPasswordError, nil)
	case 20002:
		app.Response(http.StatusOK, e.UserNoExits, nil)
	default:
		app.Response(http.StatusOK, e.Error, nil)
	}
}

// ValToken 验证token
func ValToken(c *gin.Context) {
	app := app.Gin{c}
	token := c.Query("token")
	tokenvaldata, _ := jwt.ParseToken(token)
	//fmt.Println(tokenvaldata)
	if tokenvaldata != nil {
		app.Response(http.StatusOK, e.Success, tokenvaldata)
	} else {
		app.Response(http.StatusOK, e.CheckTokenTimeout, nil)
	}
}

// ValOldPwd 旧密码验证
func ValOldPwd(c *gin.Context) {
	app := app.Gin{c}
	oldpwd := c.Query("oldpwd")
	token := c.Request.Header.Get("token")
	res := UserServer.ValOldPwd(oldpwd, token)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.OldPasswordError, nil)
	}
}

// UserUpdatePwd 修改密码
func UserUpdatePwd(c *gin.Context) {
	app := app.Gin{c}
	NewPwd := c.Query("newpwd")
	token := c.Request.Header.Get("token")
	res := UserServer.UpdateUserPwd(NewPwd, token)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.EditOldPasswordError, nil)
	}
}

// ValUserName 验证用户名是否已经存在
func ValUserName(c *gin.Context) {
	app := app.Gin{c}
	unit := c.Query("unit")
	depaerment := c.Query("department")
	username := c.Query("username")
	ValUserServer := &UserServer.User{Unit: unit, Department: depaerment, Name: username}
	res := ValUserServer.ValUserExists()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.UserExits, nil)
	}
}

// UserAdd  新增用户
func UserAdd(c *gin.Context) {
	app := app.Gin{c}
	unit := c.Query("unit")
	deaprtment := c.Query("department")
	realname := c.Query("realname")
	username := c.Query("username")
	password := c.Query("password")
	usertype := c.Query("userptype")
	addserver := &UserServer.User{Realname: realname, Name: username, Password: password, Unit: unit, Department: deaprtment, Usertype: usertype}
	res := addserver.AddUser()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.UserCreateError, nil)
	}

}

// GetMenu 获取菜单
func GetMenu(c *gin.Context) {
	app := app.Gin{c}
	token := c.Request.Header.Get("token")
	TokenData, _ := jwt.ParseToken(token)
	res := UserServer.GetMenu(TokenData.UserType)
	app.Response(http.StatusOK, e.Success, res)
}
