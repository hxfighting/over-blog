package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Rotation struct {
	ID          *int64 `json:"id" validate:"gt=0"`
	Words       string `json:"words" validate:"gte=2,lte=60"`
	CreatedUnix int64  `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64  `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string `json:"created_at" gorm:"-"`
	UpdatedAt   string `json:"updated_at" gorm:"-"`
	ImageUrl    string `json:"image_url" gorm:"-" mapstructure:"image_url" validate:"url"`
}

const ROTATION_TYPE = "App\\Http\\Models\\RotationImage"

func (this *Rotation) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}

func (Rotation) TableName() string {
	return "rotation_image"
}
