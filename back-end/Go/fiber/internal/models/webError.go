package models

import (
	"github.com/ohdata/blog/tools/util"
)

type WebError struct {
	ID          *int64  `json:"id" validate:"gt=0"`
	Type        string  `json:"type" validate:"gte=2,lte=255"`
	Code        string  `json:"code" validate:"numeric"`
	Mes         string  `json:"mes" validate:"gte=2,lte=255"`
	Url         string  `json:"url" validate:"url,lte=255"`
	CreatedUnix int64   `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64   `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string  `json:"created_at" gorm:"-"`
	UpdatedAt   string  `json:"updated_at" gorm:"-"`
	IDs         []int64 `json:"ids" gorm:"-" validate:"gt=0,dive,gt=0,numeric"`
}

func (WebError) TableName() string {
	return "web_error"
}

func (this *WebError) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
