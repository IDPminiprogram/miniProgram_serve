/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/17 16:00
 */

package main

import (
	"cea_api/middlewares"
	"cea_api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	router.UserRouters(r)
	router.UnitrsRouters(r)

	err := r.Run(":8889")
	if err != nil {
		return
	}
}
