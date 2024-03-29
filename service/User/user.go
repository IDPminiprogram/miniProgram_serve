package User

import "miniProgram_server/models"

type Page struct {
	PageNum  int
	PageSize int
}

func (p *Page) GetUser() (user []models.User, count int) {
	return models.GetUser(p.PageSize, p.PageNum)
}

func DelUser(id int) bool {
	return models.DelUser(id)
}

func AllowUser(id, allow int) bool {
	return models.AllowUser(id, allow)
}

func SearchUser(name string) []models.User {
	return models.SearchUser(name)
}
func AllowRelease(id int, allow string) bool {
	return models.AllowRelease(id, allow)
}
