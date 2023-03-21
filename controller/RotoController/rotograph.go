package RotoController

import (
	"github.com/gin-gonic/gin"
	"miniProgram_server/pkg/app"
	"miniProgram_server/pkg/e"
	"miniProgram_server/service/Rotogaph"
	"net/http"
)

// @Summary	ping example
// @Schemes
// @Description
// @Tags			添加轮播图
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/UserController/addrotograph [post]
func AddRotograph(c *gin.Context) {
	app := app.Gin{C: c}
	swiperName := c.Query("name")
	swiperImageUrl := c.Query("imageUrl")
	swiperJumpurl := c.Query("jumpUrl")
	swiperDesc := c.Query("desc")
	addData := &Rotogaph.Rotograph{SwiperName: swiperName, SwiperImnageUrl: swiperImageUrl, SwiperjumpUrl: swiperJumpurl, Swiperdesc: swiperDesc}
	res := addData.AddRoto()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.UserCreateError, nil)
	}
}

// @Summary	ping example
// @Schemes
// @Description
// @Tags		删除轮播图
// @Accept		json
// @Produce		json
// @Success		200
// @Router			/UserController/delrotograph [delete]
func DelRotograph(c *gin.Context) {
	app := app.Gin{C: c}
	id := c.Query("id")
	addData := &Rotogaph.Rotograph{SwiperID: id}
	res := addData.DelRoto()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.UserCreateError, nil)
	}
}

// @Summary	ping example
// @Schemes
// @Description
// @Tags			获取轮播图
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/UserController/addrotograph [post]
func GetRotograph(c *gin.Context) {
	app := app.Gin{C: c}
	Rotostruct := &Rotogaph.Rotograph{}
	res := Rotostruct.GetRoto()
	app.Response(http.StatusOK, e.Success, res)

}
