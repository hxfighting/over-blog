package models

import (
	"blog/database"
	"fmt"
	"time"
)

type Category struct {
	ID        *int64     `json:"id"`
	Title     *string    `json:"title"`
	Pid       *int64     `json:"pid"`
	Url       string     `json:"url"`
	Type      uint       `json:"type"`
	CreatedAt int64      `json:"created_at"`
	UpdatedAt int64      `json:"updated_at"`
	Children  []Category `json:"children" gorm:"foreignkey:pid;PRELOAD:false;-"`
}

/**
获取分类列表
*/
func GetCategoryList() (category []Category) {
	database.Db.Where("pid = ?", 0).Preload("Children").Find(&category)
	return
}

/**
添加分类
*/
func AddCategory(category *Category) bool {
	now := time.Now().Unix()
	category.CreatedAt = now
	category.UpdatedAt = now
	category.Type = 1
	res := database.Db.Create(category)
	if res.Error != nil {
		return false
	}
	result := database.Db.Model(&category).Update("url", "/category/"+fmt.Sprintf("%d", *category.ID))
	if result.Error != nil {
		return false
	}
	return true
}

/**
修改分类名称
*/
func UpdateCategory(category *Category) bool {
	now := time.Now().Unix()
	category.UpdatedAt = now
	result := database.Db.Model(&category).Update("title", category.Title)
	if result.Error != nil {
		return false
	}
	return true
}

/**
删除分类
*/
func DeleteCategory(category *Category) bool {
	exist_cate := Category{}
	database.Db.Where("id = ?", category.ID).Find(&exist_cate)
	if exist_cate.Url == "" || exist_cate.Type != 1 {
		return false
	}
	res := database.Db.Delete(category)
	if res.Error != nil {
		return false
	}
	return true
}
