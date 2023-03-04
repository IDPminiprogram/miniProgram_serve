/**
 * @Author: 戈亓
 * @Description:
 * @File:  units
 * @Version: 1.0.0
 * @Date: 2022/7/28 15:18
 */

package UnitsController

import (
	"cea_api/models"
	"cea_api/pkg/app"
	"cea_api/pkg/e"
	"cea_api/service/UnitsServer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUnit 获取单位数据
func GetUnit(c *gin.Context) {
	app := app.Gin{c}
	unitList := []models.Unit{}
	result := models.DB.Find(&unitList)
	if result.Error == nil {
		app.Response(http.StatusOK, e.Success, unitList)
	} else {
		app.Response(http.StatusOK, e.UnitGetError, nil)
	}

}

// ValUnitName 验证单位名称是否存在
func ValUnitName(c *gin.Context) {
	app := app.Gin{c}
	unitname := c.Query("unitname")
	unitnameserver := UnitsServer.Unit{Unit: unitname}
	res := unitnameserver.ValUnitName()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.UnitExists, nil)
	}
}

// AddUnit 添加单位
func AddUnit(c *gin.Context) {
	app := app.Gin{c}
	unit := c.Query("unitname")
	mark := c.Query("mark")
	addserver := &UnitsServer.Unit{unit, mark}
	res := addserver.AddUnit()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.UnitAddError, nil)
	}
}

// ValDepartName 验证部门是否存在
func ValDepartName(c *gin.Context) {
	app := app.Gin{c}
	departname := c.Query("department")
	unitmark := c.Query("unitmarl")
	valdepartnameserver := UnitsServer.AddDepart{departname, unitmark}
	res := valdepartnameserver.ValeDeaprtName()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.DepartExists, nil)
	}
}

// GetDepart 获取部门数据
func GetDepart(c *gin.Context) {
	app := app.Gin{c}
	unitmark := c.Query("unit_mark")
	departlist := []models.Department{}
	result := models.DB.Where("unit_mark = ? ", unitmark).Find(&departlist)
	if result.Error == nil {
		app.Response(http.StatusOK, e.Success, departlist)
	} else {
		app.Response(http.StatusOK, e.DepartGetError, nil)
	}
}

// AddDeartment 添加部门
func AddDeartment(c *gin.Context) {
	app := app.Gin{c}
	unitmark := c.Query("unitmark")
	department := c.Query("deapartment")
	addserver := &UnitsServer.AddDepart{UnitMark: unitmark, Department: department}
	res := addserver.AddDepartment()
	if res == true {
		app.Response(http.StatusOK, e.Success, nil)
	} else {
		app.Response(http.StatusOK, e.DepartAddError, nil)
	}
}
