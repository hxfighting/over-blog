package models

import (
	"blog/database"
	"blog/helper"
	"time"
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
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

func (Rotation) TableName() string {
	return "rotation_image"
}

/**
获取轮播图列表
*/
func (this *Rotation) GetRotationList() []Image {
	image := []Image{}
	rotations := []Rotation{}
	database.Db.Where("image_type = ?", ROTATION_TYPE).Find(&image)
	if len(image) > 0 {
		ids := []int64{}
		for _, value := range image {
			ids = append(ids, *value.Image_id)
		}
		database.Db.Where("id in (?)", ids).Find(&rotations)
		if len(rotations) > 0 {
			for k, value := range image {
				for _, val := range rotations {
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
删除轮播图
*/
func (this *Rotation) DeleteRotation() bool {
	tx := database.Db.Begin()
	res := tx.Table("image").Where("image_id = ? and image_type = ?", this.ID, ROTATION_TYPE).Delete(&Image{})
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

/**
添加轮播图
*/
func (this *Rotation) AddRotation() bool {
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
	image_type := ROTATION_TYPE
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
修改轮播图
*/
func (this *Rotation) UpdateRotation() bool {
	tx := database.Db.Begin()
	now := time.Now().Unix()
	data := map[string]interface{}{
		"image_url":  this.ImageUrl,
		"updated_at": now,
	}
	result := tx.Table("image").
		Where("image_id = ? and image_type = ?", this.ID, ROTATION_TYPE).
		Updates(data)
	if result.Error != nil {
		tx.Rollback()
		return false
	}
	res := tx.Model(this).Updates(map[string]interface{}{
		"words":      this.Words,
		"updated_at": now,
	})
	if res.Error != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
