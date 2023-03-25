package Message

import "miniProgram_server/models"

func GetMessage(pageSize, pageNum int, name string) (shop []models.Message, cout int) {
	return models.GetMessage(pageSize, pageNum, name)
}

func DelMessage(id int) bool {
	return models.DelMessage(id)
}
