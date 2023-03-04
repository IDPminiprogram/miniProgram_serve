/**
 * @Author: 戈亓
 * @Description:
 * @File:  units
 * @Version: 1.0.0
 * @Date: 2022/7/28 15:15
 */

package router

import (
	"cea_api/controller/UnitsController"
	"github.com/gin-gonic/gin"
)

func UnitrsRouters(r *gin.Engine) {
	UnitrsRouter := r.Group("/units")
	{
		UnitrsRouter.GET("/getUnits", UnitsController.GetUnit)             // 获取单位
		UnitrsRouter.GET("/getDepartments", UnitsController.GetDepart)     // 获取部门
		UnitrsRouter.POST("/addunit", UnitsController.AddUnit)             // 新增单位
		UnitrsRouter.POST("/valunit", UnitsController.ValUnitName)         // 验证单位是否存在
		UnitrsRouter.POST("/adddepartment", UnitsController.AddDeartment)  // 新增部门
		UnitrsRouter.POST("/valdepartment", UnitsController.ValDepartName) // 验证部门是否存在
	}
}
