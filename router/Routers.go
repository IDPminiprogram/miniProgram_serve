/**
 * @Author: 戈亓
 * @Description: Routing branch
 * @File:  userRouters
 * @Version: 1.0.0
 * @Date: 2022/7/17 21:46
 */

package router

import (
	"github.com/gin-gonic/gin"
	"miniProgram_server/controller/NoticeController"
	"miniProgram_server/controller/RotoController"
	"miniProgram_server/controller/ShopController"
	"miniProgram_server/controller/UserController"
)

func ServerRoutes(r *gin.Engine) {
	ServerRouter := r.Group("/service")
	{
		ServerRouter.GET("/getPhotograph", RotoController.GetRotograph)
		ServerRouter.POST("/addPhotograph", RotoController.AddRotograph)
		ServerRouter.DELETE("/delPhotograph", RotoController.DelRotograph)

		ServerRouter.GET("/getUser", UserController.GetUser)
		ServerRouter.DELETE("/delUser", UserController.DelUser)
		ServerRouter.POST("/allowUser", UserController.AllowUser)
		ServerRouter.GET("/searchUser", UserController.SearchUser)

		ServerRouter.GET("/getNotice", NoticeController.GetNotice)
		ServerRouter.POST("/addNotice", NoticeController.AddNotice)
		ServerRouter.DELETE("/delNotice", NoticeController.DelNotice)

		ServerRouter.GET("/getShop", ShopController.GetShop)
		ServerRouter.POST("/addShop", ShopController.AddShop)
		ServerRouter.DELETE("/delShop", ShopController.DelShop)
		ServerRouter.GET("/searchShop", ShopController.SearchShop)

	}
}
