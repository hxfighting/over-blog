package models

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CategoryType int

func (c CategoryType) UInt() uint {
	return uint(c)
}

const (
	CategoryArticleType CategoryType = 1
	CategoryContactType CategoryType = 2
	CategoryChatType    CategoryType = 3
)

type Category struct {
	ID        *int64      `json:"id" validate:"gt=0"`
	Title     *string     `json:"title" validate:"gte=2,lte=20"`
	Pid       *int64      `json:"pid" validate:"gte=0"`
	Url       string      `json:"url"`
	Type      uint        `json:"type"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
	Children  []*Category `json:"children" gorm:"foreignkey:pid;PRELOAD:false"`
}

/**
获取分类列表
*/
func GetCategoryList(ctx *fiber.Ctx) (category []Category, err error) {
	err = db(ctx.Context(), nil).
		Where("pid = ?", 0).
		Preload("Children").Find(&category).Error
	return
}

func CountCategoryByTitle(ctx *fiber.Ctx, title string) (total int64, err error) {
	err = db(ctx.Context(), nil).Model(Category{}).
		Where("title = ?", title).Count(&total).Error
	return
}

func CreateCategory(ctx *fiber.Ctx, c *Category) error {
	if err := db(ctx.Context(), nil).Create(c).Error; err != nil {
		return err
	}
	return db(ctx.Context(), nil).Model(c).Update("url", "/category/"+fmt.Sprintf("%d", *c.ID)).Error
}

func UpdateCategory(ctx *fiber.Ctx, updateData map[string]interface{}) error {
	return db(ctx.Context(), nil).Model(Category{}).Updates(updateData).Error
}
