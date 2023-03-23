/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/1/17 16:00
 */

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "miniProgram_server/docs" // 导入swagger
	"miniProgram_server/middlewares"
	"miniProgram_server/router"
)

// @title			小程序后端和部分管理系统接囗文档
// @version		1.0
// @contact.name	戈亓
// @contact.email	1994975650@qq.com
// @host			127.0.0.1:8889 Ceshi
func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	router.ServerRoutes(r)
	r.GET("/swagger/*any", func(c *gin.Context) {
		ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c)
	})
	err := r.Run(":8889")
	if err != nil {
		return
	}
}
