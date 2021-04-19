package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Link struct {
	ID          *int64  `json:"id" validate:"gt=0"`
	Url         *string `json:"url" validate:"url"`
	Name        *string `json:"name" validate:"gte=2,lte=30"`
	Description *string `json:"description" validate:"gte=2,lte=50"`
	Order       *int    `json:"order" validate:"gte=0,lte=9999999"`
	IsShow      *int    `json:"is_show" validate:"oneof=0 1" mapstructure:"is_show"`
	CreatedUnix int64   `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64   `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string  `json:"created_at" gorm:"-"`
	UpdatedAt   string  `json:"updated_at" gorm:"-"`
}

type SimpleLink struct {
	Url         string `json:"url"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

func (this *Link) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
