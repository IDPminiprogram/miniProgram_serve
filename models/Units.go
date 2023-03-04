/**
 * @Author: 戈亓
 * @Description:
 * @File:  unit
 * @Version: 1.0.0
 * @Date: 2022/7/24 0:02
 */

package models

type Unit struct {
	Id   int    `json:"id"`
	Unit string `json:"unit"`
	Mark string `json:"mark"`
}

func (Unit) TableName() string {
	return "Unit"
}

func (Department) TableName() string {
	return "Department"
}

func GetUnitMark(unit string) (data string) {
	var u Unit
	DB.Where("Unit = ?", unit).First(&u)
	data = u.Mark
	return
}

func ValUnitName(unitname string) (res bool) {
	var u Unit
	DB.Where("Unit = ?", unitname).First(&u)
	if u.Id > 0 {
		return false
	} else {
		return true
	}
}

func AddUnit(unit, mark string) (res bool) {
	result := DB.Create(&Unit{Unit: unit, Mark: mark})
	if result.Error != nil {
		return false
	}
	return true
}

type Department struct {
	ID         int
	Department string `json:"department"`
	SerNum     int    `json:"ser_num"`
	UnitMark   string `json:"unit_mark"`
}

func GetDepartSer(unitmark, department string) (data int) {
	var d Department
	DB.Where("department = ? and unit_mark = ?", department, unitmark).Find(&d)
	data = d.SerNum
	return

}

func ValeDeaprtName(unitmark, deaprtname string) (res bool) {
	var d Department
	DB.Where("department = ? and unit_mark = ?", deaprtname, unitmark).First(&d)
	if d.ID > 0 {
		return false
	} else {
		return true
	}
}

func AddDepaerment(department, unitmark string) (res bool) {
	departmentList := []Department{}
	DB.Where("unit_mark = ?", unitmark).Find(&departmentList)
	sernum := len(departmentList) + 1
	result := DB.Create(&Department{Department: department, SerNum: sernum, UnitMark: unitmark})
	if result.Error != nil {
		return false
	}
	return true
}
