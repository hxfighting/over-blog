package models

import (
	"github.com/ohdata/blog/tools/util"
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
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
