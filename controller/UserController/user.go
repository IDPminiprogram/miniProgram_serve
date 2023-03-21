package UserController

import (
	"github.com/gin-gonic/gin"
	"miniProgram_server/pkg/app"
	"miniProgram_server/pkg/e"
	"miniProgram_server/service/User"
	"net/http"
	"strconv"
)

// GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			获取用户
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/getUser/GetUser [GET]
func GetUser(c *gin.Context) {
	app := app.Gin{C: c}
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	getData := &User.Page{PageNum: pageNum, PageSize: pageSize}
	res := getData.GetUser()
	if res != nil {
		app.Response(http.StatusOK, e.Success, res)
	}
}

// DelUser GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			删除用户
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/delUser/DelUser [GET]
func DelUser(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	res := User.DelUser(id)
	if res == true {
		app.Response(http.StatusOK, e.Success, res)
	}
}

// DelUser GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			修改用户权限用户
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/delUser/DelUser [GET]
func AllowUser(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	allow, _ := strconv.Atoi(c.Query("allow"))
	res := User.AllowUser(id, allow)
	if res == true {
		app.Response(http.StatusOK, e.Success, res)
	}
}

// DelUser GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			搜索用户
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/searchUser/SearchUser [GET]
func SearchUser(c *gin.Context) {
	app := app.Gin{C: c}
	name := c.Query("name")
	res := User.SearchUser(name)
	app.Response(http.StatusOK, e.Success, res)
}
