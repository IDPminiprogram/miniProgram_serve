package models

import "fmt"

type Business struct {
	BusinessID     int     `gorm:"column:business_id"`
	ShopId         int     `gorm:"column:shop_id"`
	ShopGoodsImage string  `gorm:"column:shop_goods_image"`
	ShopGoodsTitle string  `gorm:"column:shop_goods_title"`
	ShopGoodsPrice float64 `gorm:"column:shop_goods_price"`
}

func (b *Business) TableName() string {
	return "mango_shop_business"
}

func AddBusiness(id, shopid int, image, title string, piece float64) bool {
	if id > 0 {
		business := Business{
			ShopId:         shopid,
			ShopGoodsImage: image,
			ShopGoodsTitle: title,
			ShopGoodsPrice: piece,
		}
		DB.Model(&Business{}).Where("business_id = ?", id).Updates(&business)
	} else {
		business := Business{
			ShopId:         shopid,
			ShopGoodsImage: image,
			ShopGoodsTitle: title,
			ShopGoodsPrice: piece,
		}
		result := DB.Create(&business)
		if result.Error != nil {
			return false
		}
		return true
	}
	return true
}

// DelRotoGraph 删除商品
func DelBusiness(id int) bool {
	result := DB.Where("business_id = ?", id).Delete(&Business{})
	if result.Error != nil {
		return false
	}
	return true
}

func SearchBusiness(name string) []Business {
	var b []Business
	DB.Where("shop_goods_title LIKE ?", name).Find(&b)
	fmt.Println(b)
	return b
}

func GetBusiness(id, NumSize, NumPage int) (shop []Business, count int) {
	NumPage = NumPage*NumSize - NumSize
	var s []Business
	var b []Business
	DB.Limit(NumSize).Offset(NumPage).Where("shop_id = ?", id).Find(&s)
	result := DB.Where("shop_id = ?", id).Find(&b)
	count = int(result.RowsAffected)
	return s, count
}
