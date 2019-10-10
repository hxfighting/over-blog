package models

import (
	"blog/database"
	"blog/helper"
	"time"
)

type Chat struct {
	ID          int64  `json:"id" validate:"gt=0"`
	Content     string `json:"content" validate:"gte=2,lte=255"`
	IsShow      *uint  `json:"is_show" mapstructure:"is_show" validate:"oneof=0 1"`
	CreatedUnix int64  `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64  `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string `json:"created_at" gorm:"-"`
	UpdatedAt   string `json:"updated_at" gorm:"-"`
}

func (this *Chat) AfterFind() {
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取说说列表
*/
func GetChatList(pageNum, pageSize int64) []Chat {
	chat := []Chat{}
	database.Db.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&chat)
	return chat
}

/**
添加说说
*/
func AddChat(chat *Chat) bool {
	chat.CreatedUnix = time.Now().Unix()
	chat.UpdatedUnix = time.Now().Unix()
	res := database.Db.Create(chat)
	if res.Error != nil {
		return false
	}
	return true
}

/**
修改说说
*/
func UpdateChat(chat *Chat) bool {
	now := time.Now().Unix()
	data := map[string]interface{}{
		"is_show":    chat.IsShow,
		"content":    chat.Content,
		"updated_at": now,
	}
	result := database.Db.Model(&chat).Updates(data)
	if result.Error != nil {
		return false
	}
	return true
}

/**
删除说说
*/
func DeleteChat(chat *Chat) bool {
	res := database.Db.Delete(chat)
	if res.Error != nil {
		return false
	}
	return true
}
