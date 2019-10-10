package models

import (
	"blog/database"
	"blog/helper"
	"time"
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
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

/**
获取照片列表
*/
func (this *Photo) GetPhotoList() []Image {
	image := []Image{}
	photos := []Photo{}
	database.Db.Where("image_type = ?", PHOTO_TYPE).Find(&image)
	if len(image) > 0 {
		ids := []int64{}
		for _, value := range image {
			ids = append(ids, *value.ID)
		}
		database.Db.Where("id in (?)", ids).Find(&photos)
		if len(photos) > 0 {
			for k, value := range image {
				for _, val := range photos {
					if *value.Image_id == *val.ID {
						value.Image = val
						image[k] = value
					}
				}
			}
		}
	}
	return image
}

/**
添加照片
*/
func (this *Photo) AddPhoto() bool {
	uninx_time := time.Now().Unix()
	tx := database.Db.Begin()
	this.CreatedUnix = uninx_time
	this.UpdatedUnix = uninx_time
	res := tx.Create(this)
	if res.Error != nil {
		tx.Rollback()
		return false
	}
	image := Image{}
	image_type := PHOTO_TYPE
	image.Image_type = &image_type
	image.Image_id = this.ID
	image.Image_url = &this.ImageUrl
	image.CreatedUnix = uninx_time
	image.UpdatedUnix = uninx_time
	result := tx.Create(&image)
	if result.Error != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

/**
修改照片
*/
func (this *Photo) UpdatePhoto(image Image) bool {
	now := time.Now().Unix()
	data := map[string]interface{}{
		"image_url":  image.Image_url,
		"updated_at": now,
	}
	result := database.Db.Table("image").
		Where("image_id = ? and image_type = ?", image.ID, PHOTO_TYPE).
		Updates(data)
	if result.Error != nil {
		return false
	}
	return true
}

/**
删除照片
*/
func (this *Photo) DeletePhoto() bool {
	tx := database.Db.Begin()
	res := tx.Table("image").Where("image_id = ? and image_type = ?", this.ID, PHOTO_TYPE).Delete(&Image{})
	if res.Error != nil {
		tx.Rollback()
		return false
	}
	res = tx.Delete(this)
	if res.Error != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
