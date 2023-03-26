package models

import (
	"fmt"
	"time"
)

type Shop struct {
	ShopID          int       `gorm:"column:shop_id"`
	ShopName        string    `gorm:"column:shop_name"`
	ShopIntro       string    `gorm:"column:shop_intro"`
	ShopPhoneNumber string    `gorm:"column:shop_phone"`
	ShopAvatar      string    `gorm:"column:shop_avatar"`
	ShopLatitude    string    `gorm:"column:shop_latitude"`
	ShopLongitude   string    `gorm:"column:shop_longitude"`
	ShopCreatTime   time.Time `gorm:"column:shop_creat_time"`
}

func (s *Shop) TableName() string {
	return "mango_shop"
}

func AddShop(id int, shopName, shopIntro, shopAvatar, shopLatitude, shopLongitude, shopPhone string) bool {
	createTime := time.Now()
	if id > 0 {
		shops := Shop{
			ShopName:        shopName,
			ShopIntro:       shopIntro,
			ShopCreatTime:   createTime,
			ShopPhoneNumber: shopPhone,
			ShopAvatar:      shopAvatar,
			ShopLatitude:    shopLatitude,
			ShopLongitude:   shopLongitude,
		}
		DB.Model(&Shop{}).Where("shop_id = ?", id).Updates(&shops)
	} else {
		fmt.Println(shopName, shopIntro, shopPhone, shopAvatar, shopLatitude, shopLongitude)
		shops := Shop{
			ShopName:        shopName,
			ShopIntro:       shopIntro,
			ShopCreatTime:   createTime,
			ShopPhoneNumber: shopPhone,
			ShopAvatar:      shopAvatar,
			ShopLatitude:    shopLatitude,
			ShopLongitude:   shopLongitude,
		}
		fmt.Println(shops)
		result := DB.Create(&shops)
		if result.Error != nil {
			return false
		}
		return true
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
	result := DB.Find(&[]Shop{})
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
