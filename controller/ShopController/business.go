package ShopController

import (
	"github.com/gin-gonic/gin"
	"miniProgram_server/pkg/app"
	"miniProgram_server/pkg/e"
	"miniProgram_server/service/Shop"
	"net/http"
	"strconv"
)

// AddBusiness @Summary	ping example
// @Schemes
// @Description
// @Tags			添加商品
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/ShopController/AddBusiness [post]
func AddBusiness(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	shopid, _ := strconv.Atoi(c.Query("shop_id"))
	image := c.Query("image")
	title := c.Query("title")
	piece, _ := strconv.ParseFloat(c.Query("piece"), 64)
	res := Shop.AddBusiness(id, shopid, image, title, piece)
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	}
}

// GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			获取商品
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/getShop/GetShop [GET]
func GetBusiness(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("shop_id"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	res, count := Shop.GetBusiness(id, pageNum, pageSize)
	if res != nil {
		app.CountResponse(http.StatusOK, e.Success, res, count)
	}
}

// DelUser GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			删除商品
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/delUser/DelUser [GET]
func DelBusiness(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	res := Shop.DelBusiness(id)
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
func SearchBusiness(c *gin.Context) {
	app := app.Gin{C: c}
	name := c.Query("name")
	res := Shop.SeachBusiness(name)
	app.Response(http.StatusOK, e.Success, res)
}
