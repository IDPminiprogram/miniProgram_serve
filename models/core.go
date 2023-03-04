/**
 * @Author: 戈亓
 * @Description:
 * @File:  core.go
 * @Version: 1.0.0
 * @Date: 2022/7/22 1:55
 */

package models

import "fmt"
import "gorm.io/driver/mysql"
import "gorm.io/gorm"

var DB *gorm.DB
var err error

func init() {
	dsn := "root:fhtgis502@tcp(152.136.123.49:3306)/irms?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "用户名:密码@tcp(数据库地址:端口号 默认是3306)/数据库名称?charset=utf8mb4&parseTime=True&loc=Local'
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
