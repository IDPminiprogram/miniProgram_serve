package models

import (
	"strconv"
	"time"
)

type Shop struct {
	ShopID        int       `gorm:"column:shop_id"`
	ShopName      string    `gorm:"column:shop_name"`
	ShopIntro     string    `gorm:"column:shop_intro"`
	shopPhone     string    `gorm:"column:shop_phone"`
	shopAvatar    string    `gorm:"column:shop_avatar"`
	shopLatitude  string    `gorm:"column:shop_latitude"`
	shopLongitude string    `gorm:"column:shop_longitude"`
	ShopCreatTime time.Time `gorm:"column:shop_creat_time"`
}

func (s *Shop) TableName() string {
	return "mango_shop"
}

func AddShop(shopName, shopIntro, shopAvatar, shopLatitude, shopLongitude string, shopPhone int) bool {
	createTime := time.Now()
	//fmt.Println(shopName, shopIntro, shopPhone, shopAvatar, shopLatitude, shopLongitude)
	shopPhones := strconv.Itoa(shopPhone)
	result := DB.Create(&Shop{ShopName: shopName, ShopIntro: shopIntro, shopPhone: shopPhones, shopAvatar: shopAvatar, shopLatitude: shopLatitude, shopLongitude: shopLongitude, ShopCreatTime: createTime})
	if result.Error != nil {
		return false
	}
	return true

}

func SearchShop(name string) []Shop {
	var s []Shop
	DB.Where("shop_name LIKE ?", name).Find(&s)
	return s
}

func GetShop(NumSize, NumPage int) (shop []Shop, count int) {
	NumPage = NumPage*NumSize - NumSize
	var s []Shop
	DB.Limit(NumSize).Offset(NumPage).Find(&s)
	result := DB.Find(&Shop{})
	count = int(result.RowsAffected)
	return s, count
}

// DelRotoGraph 删除商户
func DelShop(id int) bool {
	result := DB.Where("shop_id = ?", id).Delete(&Shop{})
	if result.Error != nil {
		return false
	}
	return true
}
