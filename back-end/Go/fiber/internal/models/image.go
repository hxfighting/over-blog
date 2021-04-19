package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Image struct {
	ID          *int64      `json:"id"`
	Image_type  *string     `json:"image_type"`
	Image_id    *int64      `json:"image_id"`
	Image_url   *string     `json:"image_url"`
	CreatedUnix int64       `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64       `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string      `json:"created_at" gorm:"-"`
	UpdatedAt   string      `json:"updated_at" gorm:"-"`
	Image       interface{} `json:"image" gorm:"-"`
}

func (this *Image) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
