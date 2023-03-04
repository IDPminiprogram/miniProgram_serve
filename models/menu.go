/**
 * @Author: 戈亓
 * @Description:
 * @File:  menu
 * @Version: 1.0.0
 * @Date: 2022/8/4 21:15
 */

package models

import "fmt"

type Menu struct {
	Id        int
	MenuName  string
	OrderNum  int
	Menuicon  string
	ChildItem []MenuItem
}

func (Menu) TableName() string {
	return "menu"
}

type MenuItem struct {
	Id       int
	MenuId   int
	ItemName string
	OrderNum int
	Url      string
	Itemicon string
}

func (MenuItem) TableName() string {
	return "menu_item"
}

type MenuPermission struct {
	Id       int
	Usertype int
	Item     int
}

func (MenuPermission) TableName() string {
	return "MenuPermission"
}

type UserType struct {
	Id       int
	UserType string
}

func (UserType) TableName() string {
	return "UserType"
}

func (u *UserType) GetUsertypeID() (menudata []Menu) {
	DB.Where("Usertype = ?", u.UserType).First(&u)
	//fmt.Println(u.Id)
	MenuPermissionlist := []MenuPermission{}
	DB.Where("usertype = ?", u.Id).Find(&MenuPermissionlist)
	//fmt.Println(MenuPermissionlist)
	menulist := []Menu{}
	DB.Find(&menulist)
	//fmt.Println(menulist, "1111111")
	menuitemlist := []MenuItem{}
	itemmenu := []Menu{}
	for i := 0; i < len(MenuPermissionlist); i++ {
		DB.Where("id = ?", MenuPermissionlist[i].Item).Find(&menuitemlist)
		//fmt.Println(menuitemlist)
		// 查询一级菜单
		DB.Where("id = ?", menuitemlist[0].MenuId).Find(&itemmenu)
		//fmt.Println(itemmenu)
		for j := 0; j < len(menulist); j++ {
			//fmt.Println(menulist[j])
			if menulist[j].Id == itemmenu[0].Id {
				menulist[j].ChildItem = append(menulist[j].ChildItem, menuitemlist[0])
			}
		}
	}
	//menudata = menulist
	menudata = MenuSort(menulist)
	return menudata
}
func MenuSort(menudata []Menu) (sortdata []Menu) {
	newMenu := []Menu{}
	for i := 0; i < len(menudata); i++ {
		fmt.Println(menudata[i].OrderNum)
		//newMenu[0] = menudata[i]
	}
	fmt.Println(newMenu)
	return sortdata
}
