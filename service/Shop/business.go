package Shop

import "miniProgram_server/models"

func AddShop(shop_name, shop_intro, shop_avatar, shop_latitude, shop_longitude string, shopPhone int) (res bool) {
	return models.AddShop(shop_name, shop_intro, shop_avatar, shop_latitude, shop_longitude, shopPhone)
}

func DelShop(noticeid int) bool {
	return models.DelShop(noticeid)
}

func GetShop(pageSize, pageNum int) (shop []models.Shop, cout int) {
	return models.GetShop(pageSize, pageNum)
}
func SeachShop(name string) []models.Shop {
	return models.SearchShop(name)
}
