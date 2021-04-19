package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Contact struct {
	ID           *int64  `json:"id" validate:"gt=0"`
	Content      *string `json:"content" validate:"gte=2,lte=255"`
	Name         *string `json:"name" validate:"gte=2,lte=20"`
	Email        *string `json:"email" validate:"email"`
	IsReply      *int64  `json:"is_reply" mapstructure:"is_reply"`
	CreatedUnix  int64   `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix  int64   `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt    string  `json:"created_at" gorm:"-"`
	UpdatedAt    string  `json:"updated_at" gorm:"-"`
	ReplyContent *string `json:"reply_content" mapstructure:"reply_content" validate:"gte=2,lte=255"`
}

func (Contact) TableName() string {
	return "contact"
}

func (this *Contact) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
