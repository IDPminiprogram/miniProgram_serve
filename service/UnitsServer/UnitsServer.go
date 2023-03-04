/**
 * @Author: 戈亓
 * @Description:
 * @File:  UnitsServer
 * @Version: 1.0.0
 * @Date: 2022/7/28 16:18
 */

package UnitsServer

import "cea_api/models"

type Unit struct {
	Unit string
	Mark string
}

func (u *Unit) ValUnitName() (res bool) {
	return models.ValUnitName(u.Unit)
}

func (u *Unit) AddUnit() (res bool) {
	return models.AddUnit(u.Unit, u.Mark)
}

func GetUnitMark(unit string) (data string) {
	return models.GetUnitMark(unit)
}

type AddDepart struct {
	Department string
	UnitMark   string
}

func GetDepartSer(unitmark, department string) (data int) {
	return models.GetDepartSer(unitmark, department)
}

// ValeDeaprtName 验证部门是否存在
func (d *AddDepart) ValeDeaprtName() (res bool) {
	return models.ValeDeaprtName(d.Department, d.UnitMark)
}

func (d *AddDepart) AddDepartment() (res bool) {
	return models.AddDepaerment(d.Department, d.UnitMark)
}
