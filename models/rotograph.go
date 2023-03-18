package models

type RotoGrph struct {
	SwiperID       int    `gorm:"column:id"`
	SwiperName     string `gorm:"column:swiper_name"`
	SwiperImageUrl string `gorm:"column:swiper_image"`
	SwiperjumpUrl  string `gorm:"column:swiper_jump"`
	Swiperdesc     string `gorm:"column:swiper_desc"`
}

func (r *RotoGrph) TableName() string {
	return "mango_swiper"
}

// AddRotoGraph 新增轮播图
func AddRotoGraph(swiperName, swiperImnageUrl, swiperJumpurl, swiperDesc string) (res bool) {
	result := DB.Create(&RotoGrph{SwiperName: swiperName, SwiperImageUrl: swiperImnageUrl, SwiperjumpUrl: swiperJumpurl, Swiperdesc: swiperDesc})
	if result.Error != nil {
		return false
	}
	return true
}
func DelRotoGraph(swiperid string) bool {
	result := DB.Where("id = ?", swiperid).Delete(&RotoGrph{})
	if result.Error != nil {
		return false
	}
	return true
}
func GetRotoGraph() []RotoGrph {
	var rotos []RotoGrph
	DB.Find(&rotos)
	return rotos
}
