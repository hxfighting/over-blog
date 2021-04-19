package models

import (
	"blog/database"
	"blog/helper"
	"time"
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
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取标签列表
*/
func (this Tag) GetTagList(pageNum, pageSize int64) map[string]interface{} {
	tag := []Tag{}
	db := database.Db.Table("tag")
	var total int64 = 0
	db.Count(&total)
	db.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&tag)
	data := make(map[string]interface{})
	data["total"] = total
	data["list"] = tag
	return data
}

/**
删除tag
*/
func (this Tag) DeleteTag() bool {
	tx := database.Db.Begin()
	re := database.Db.Table("article_tag").Where("tag_id = ?", *this.ID).Delete(Tag{})
	if re.Error != nil {
		tx.Rollback()
		return false
	}
	res := tx.Delete(&this)
	if res.Error != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

/**
检查tag是否存在
*/
func checkExistTagName(name string, id int64) bool {
	exist_tag := Tag{}
	database.Db.Table("tag").Where("name = ?", name).First(&exist_tag)
	if exist_tag.ID == nil {
		return true
	}
	if *exist_tag.ID == id {
		return true
	}
	return false
}

/**
添加tag
*/
func (this Tag) AddTag() bool {
	if !checkExistTagName(this.Name, 0) {
		return false
	}
	this.CreatedUnix = time.Now().Unix()
	this.UpdatedUnix = time.Now().Unix()
	res := database.Db.Create(&this)
	if res.Error != nil {
		return false
	}
	return true
}

/**
修改tag信息
*/
func (this Tag) UpdateTag() bool {
	if !checkExistTagName(this.Name, *this.ID) {
		return false
	}
	now := time.Now().Unix()
	data := map[string]interface{}{
		"name":       this.Name,
		"updated_at": now,
	}
	result := database.Db.Model(&this).Updates(data)
	if result.Error != nil {
		return false
	}
	return true
}
