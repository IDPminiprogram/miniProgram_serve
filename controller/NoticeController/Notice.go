package NoticeController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniProgram_server/pkg/app"
	"miniProgram_server/pkg/e"
	"miniProgram_server/service/Notice"
	"net/http"
	"strconv"
)

// @Summary	ping example
// @Schemes
// @Description
// @Tags			添加
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/UserController/addrotograph [post]
func AddNotice(c *gin.Context) {
	app := app.Gin{C: c}
	noticeDetail := c.Query("noticeDetail")
	fmt.Println(noticeDetail)
	res := Notice.AddNotice(noticeDetail)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	}
}

// @Summary	ping example
// @Schemes
// @Description
// @Tags		删除
// @Accept		json
// @Produce		json
// @Success		200
// @Router			/UserController/delrotograph [delete]
func DelNotice(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	res := Notice.DelNotice(id)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	}
}

// @Summary	ping example
// @Schemes
// @Description
// @Tags			获取
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/UserController/addrotograph [post]
func GetNotice(c *gin.Context) {
	app := app.Gin{C: c}
	res := Notice.GetNotice()
	app.Response(http.StatusOK, e.Success, res)

}
