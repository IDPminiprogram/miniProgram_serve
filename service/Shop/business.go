package Shop

import "miniProgram_server/models"

func AddShop(id int, shop_name, shop_intro, shop_avatar, shop_latitude, shop_longitude, shopPhone string) (res bool) {
	return models.AddShop(id, shop_name, shop_intro, shop_avatar, shop_latitude, shop_longitude, shopPhone)
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
