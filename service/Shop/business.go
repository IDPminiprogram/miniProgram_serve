package Shop

import "miniProgram_server/models"

func AddBusiness(id, shopid int, image, title string, piece float64) (res bool) {
	return models.AddBusiness(id, shopid, image, title, piece)
}

func DelBusiness(business int) bool {
	return models.DelBusiness(business)
}

func GetBusiness(id, pageSize, pageNum int) (shop []models.Business, cout int) {
	return models.GetBusiness(id, pageSize, pageNum)
}
func SeachBusiness(name string) []models.Business {
	return models.SearchBusiness(name)
}
