package models

type Notice struct {
	NoticeID    int    `gorm:"column:notice_id"`
	NoticeDatil string `gorm:"column:notice_detail"`
}

func (r *Notice) TableName() string {
	return "mango_notice"
}

// AddRotoGraph 新增公告
func AddNotice(NoticeDatil string) (res bool) {
	result := DB.Debug().Create(&Notice{NoticeDatil: NoticeDatil})
	if result.Error != nil {
		return false
	}
	return true
}

// DelRotoGraph 删除公告
func DelNotice(Noticeid int) bool {
	result := DB.Where("notice_id = ?", Noticeid).Delete(&Notice{})
	if result.Error != nil {
		return false
	}
	return true
}

// GetRotoGraph 获取公告
func GetNotice() []Notice {
	var notices []Notice
	DB.Find(&notices)
	return notices
}
