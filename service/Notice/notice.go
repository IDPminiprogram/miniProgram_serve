package Notice

import "miniProgram_server/models"

func AddNotice(noticeDetial string) (res bool) {
	return models.AddNotice(noticeDetial)
}

func DelNotice(noticeid int) bool {
	return models.DelNotice(noticeid)
}

func GetNotice() []models.Notice {
	return models.GetNotice()
}
