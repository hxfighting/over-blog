package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Photo struct {
	ID          *int64 `json:"id" validate:"gt=0"`
	CreatedUnix int64  `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64  `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string `json:"created_at" gorm:"-"`
	UpdatedAt   string `json:"updated_at" gorm:"-"`
	ImageUrl    string `json:"image_url" gorm:"-" mapstructure:"image_url" validate:"url"`
}

const PHOTO_TYPE = "App\\Http\\Models\\Photo"

func (this *Photo) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
