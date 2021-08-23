package models

import (
	"github.com/ohdata/blog/tools/util"
)

type Tag struct {
	ID          *int64 `json:"id" validate:"gt=0"`
	Name        string `json:"name" validate:"gte=2,lte=20"`
	CreatedUnix int64  `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64  `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string `json:"created_at" gorm:"-"`
	UpdatedAt   string `json:"updated_at" gorm:"-"`
}

func (Tag) TableName() string {
	return "tag"
}

func (this *Tag) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
}
