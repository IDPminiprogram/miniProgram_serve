package User

import "miniProgram_server/models"

type Page struct {
	PageNum  int
	PageSize int
}

func (p *Page) GetUser() []models.User {
	return models.GetUser(p.PageSize, p.PageNum)
}

func DelUser(id int) bool {
	return models.DelUser(id)
}
