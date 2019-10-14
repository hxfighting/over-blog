package models

import (
	"blog/database"
	"blog/helper"
	"time"
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
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取错误日志列表
*/
func (this WebError) GetWebErrorList(pageNum, pageSize int64) map[string]interface{} {
	webError := []WebError{}
	db := database.Db.Table("web_error")
	var total int64 = 0
	db.Count(&total)
	db.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&webError)
	data := make(map[string]interface{})
	data["total"] = total
	data["list"] = webError
	return data
}

/**
添加错误日志
*/
func (this WebError) AddWebError() bool {
	this.CreatedUnix = time.Now().Unix()
	this.UpdatedUnix = time.Now().Unix()
	res := database.Db.Create(&this)
	if res.Error != nil {
		return false
	}
	return true
}

/**
删除错误日志
*/
func (this WebError) DeleteWebError() bool {
	res := database.Db.Where("id in (?)", this.IDs).Delete(WebError{})
	if res.Error != nil {
		return false
	}
	return true
}
