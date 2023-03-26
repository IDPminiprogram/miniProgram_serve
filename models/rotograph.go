package models

type RotoGrph struct {
	SwiperID       int    `gorm:"column:swiper_id"`
	SwiperName     string `gorm:"column:swiper_name"`
	SwiperImageUrl string `gorm:"column:swiper_image_url"`
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

// DelRotoGraph 删除轮播图
func DelRotoGraph(swiperid string) bool {
	result := DB.Where("swiper_id = ?", swiperid).Delete(&RotoGrph{})
	if result.Error != nil {
		return false
	}
	return true
}

// GetRotoGraph 获取轮播图数据
func GetRotoGraph() []RotoGrph {
	var rotos []RotoGrph
	DB.Find(&rotos)
	return rotos
}
