package Rotogaph

import (
	"miniProgram_server/models"
)

type Rotograph struct {
	SwiperID        string
	SwiperName      string
	SwiperImnageUrl string
	SwiperjumpUrl   string
	Swiperdesc      string
}

func (r *Rotograph) AddRoto() (res bool) {
	return models.AddRotoGraph(r.SwiperName, r.SwiperImnageUrl, r.SwiperjumpUrl, r.Swiperdesc)
}

func (r *Rotograph) DelRoto() bool {
	return models.DelRotoGraph(r.SwiperID)
}

func (r *Rotograph) GetRoto() []models.RotoGrph {
	return models.GetRotoGraph()
}
