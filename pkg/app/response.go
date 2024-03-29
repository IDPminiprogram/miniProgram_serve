/**
 * @Author: 戈亓
 * @Description: 封装返回值
 * @File:  responde
 * @Version: 1.0.0
 * @Date: 2022/7/30 15:00
 */

package app

import (
	"github.com/gin-gonic/gin"
	"miniProgram_server/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(HttpCode, ErrCode int, data interface{}) {
	g.C.JSON(HttpCode, gin.H{
		"code": ErrCode,
		"msg":  e.GetMsg(ErrCode),
		"data": data,
	})

}
func (g *Gin) CountResponse(HttpCode, ErrCode int, data interface{}, count int) {
	g.C.JSON(HttpCode, gin.H{
		"code":  ErrCode,
		"msg":   e.GetMsg(ErrCode),
		"data":  data,
		"count": count,
	})

}
