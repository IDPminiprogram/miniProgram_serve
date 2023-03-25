package Message

import (
	"github.com/gin-gonic/gin"
	"miniProgram_server/pkg/app"
	"miniProgram_server/pkg/e"
	"miniProgram_server/service/Message"
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
func GetMessage(c *gin.Context) {
	app := app.Gin{C: c}
	name := c.Query("name")
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	res, count := Message.GetMessage(pageNum, pageSize, name)
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
func DelMessage(c *gin.Context) {
	app := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Query("id"))
	res := Message.DelMessage(id)
	if res == true {
		app.Response(http.StatusOK, e.Success, res)
	}
}

// DelUser GetUser @Summary	ping example
// @Schemes
// @Description
// @Tags			搜索信息
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/searchMessage/SearchShop [GET]
func SearchMessage(c *gin.Context) {
	app := app.Gin{C: c}
	name := c.Query("name")
	res := Shop.SeachShop(name)
	app.Response(http.StatusOK, e.Success, res)
}
