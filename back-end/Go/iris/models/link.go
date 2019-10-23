package models

import (
	"blog/database"
	"blog/helper"
	"time"
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
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取友联列表
*/
func GetLinkList(pageNum, pageSize int64, name string) map[string]interface{} {
	link := []Link{}
	var db = database.Db.Table("link")
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	var total int64 = 0
	db.Count(&total)
	db.Order("`order` desc").Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&link)
	data := make(map[string]interface{})
	data["total"] = total
	data["list"] = link
	return data
}

/**
删除友联
*/
func DeleteLink(link *Link) bool {
	res := database.Db.Delete(link)
	if res.Error != nil {
		return false
	}
	return true
}

/**
添加友联
*/
func AddLink(link *Link) bool {
	link.CreatedUnix = time.Now().Unix()
	link.UpdatedUnix = time.Now().Unix()
	res := database.Db.Create(link)
	if res.Error != nil {
		return false
	}
	return true
}

/**
修养友联
*/
func UpdateLink(link *Link) bool {
	now := time.Now().Unix()
	data := map[string]interface{}{
		"is_show":     link.IsShow,
		"name":        link.Name,
		"description": link.Description,
		"order":       link.Order,
		"url":         link.Url,
		"updated_at":  now,
	}
	result := database.Db.Model(&link).Updates(data)
	if result.Error != nil {
		return false
	}
	return true
}
