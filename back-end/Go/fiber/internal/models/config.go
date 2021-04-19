package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Config struct {
	ID          *int64  `json:"id" validate:"gt=0"`
	Title       *string `json:"title" validate:"gte=2,lte=200"`
	Name        *string `json:"name" validate:"gte=2,lte=200"`
	Val         *string `json:"val" validate:"gte=2,lte=65535"`
	Type        *int64  `json:"type" validate:"oneof=1 2 3"`
	CreatedUnix int64   `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64   `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string  `json:"created_at" gorm:"-"`
	UpdatedAt   string  `json:"updated_at" gorm:"-"`
}

type SimpleConfig struct {
	Title string
	Val   string
	Name  string
}

func (Config) TableName() string {
	return "web_config"
}

func (this *Config) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}

const FOOTER_TYPE int64 = 2
