package models

import (
	"time"
)

type Message struct {
	MessageId        int       `gorm:"column:message_id"`
	UserIdAnonymity  int       `gorm:"column:user_id_anonymity"`
	UserId           int       `gorm:"column:user_id"`
	CategoryId       int       `gorm:"column:category_id"`
	UserPhone        string    `gorm:"column:user_phone"`
	UserMajor        string    `gorm:"column:user_major"`
	UserLevel        string    `gorm:"column:user_level"`
	MessageDetail    string    `gorm:"column:message_detail"`
	MessageShare     int       `gorm:"column:message_share"`
	MessageComment   int       `gorm:"column:message_comment"`
	MessageWatch     int       `gorm:"column:message_watch"`
	MessageCreatTime time.Time `gorm:"column:message_creat_time"`
}

type Messageback struct {
	MessageId        int
	UserIdAnonymity  string
	UserId           string
	CategoryId       string
	UserPhone        string
	UserMajor        string
	UserLevel        string
	MessageDetail    string
	MessageShare     int
	MessageComment   int
	MessageWatch     int
	MessageCreatTime time.Time
}

func (m *Message) TableName() string {
	return "mango_message"
}

func GetMessage(NumSize, NumPage int, name string) (message []Message, count int) {
	if name != "" {
		NumPage = NumPage*NumSize - NumSize
		var m []Message
		DB.Limit(NumSize).Offset(NumPage).Where("message_detail LIKE ?", name).Find(&m)
		result := DB.Where("message_detail LIKE ?", name).Find(&[]Message{})
		count = int(result.RowsAffected)
		return m, count
	} else {
		NumPage = NumPage*NumSize - NumSize
		var m []Message
		DB.Limit(NumSize).Offset(NumPage).Find(&m)
		result := DB.Find(&[]Message{})
		count = int(result.RowsAffected)
		return m, count
	}
}

// DelRotoGraph 删除
func DelMessage(id int) bool {
	result := DB.Where("message_id = ?", id).Delete(&Message{})
	if result.Error != nil {
		return false
	}
	return true
}
