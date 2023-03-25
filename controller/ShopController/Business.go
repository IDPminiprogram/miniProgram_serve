package ShopController

import (
	"github.com/gin-gonic/gin"
	"miniProgram_server/pkg/app"
	"miniProgram_server/pkg/e"
	"miniProgram_server/service/Shop"
	"net/http"
	"strconv"
)

// GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			获取商户
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/getShop/GetShop [GET]
func GetShop(c *gin.Context) {
	app := app.Gin{C: c}
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	res, count := Shop.GetShop(pageNum, pageSize)
	if res != nil {
		app.CountResponse(http.StatusOK, e.Success, res, count)
	}
}

// @Summary	ping example
// @Schemes
// @Description
// @Tags			添加
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/ShopController/AddShop [post]
func AddShop(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	shopName := c.Query("shop_name")
	shopIntro := c.Query("shop_intro")
	shopPhone := c.Query("shop_phone")
	shopAvatar := c.Query("shop_avatar")
	shopLatitude := c.Query("shop_latitude")
	shopLongitude := c.Query("shop_longitude")
	res := Shop.AddShop(id, shopName, shopIntro, shopAvatar, shopLatitude, shopLongitude, shopPhone)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
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
func DelShop(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	res := Shop.DelShop(id)
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
func SearchShop(c *gin.Context) {
	app := app.Gin{C: c}
	name := c.Query("name")
	res := Shop.SeachShop(name)
	app.Response(http.StatusOK, e.Success, res)
}
