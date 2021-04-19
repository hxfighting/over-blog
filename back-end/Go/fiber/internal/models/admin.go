package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Admin struct {
	Model
	Name     string `json:"name" validate:"gte=2,lte=15"`
	Password string `json:"-" validate:"gte=6,lte=16"`
	Email    string `json:"email" validate:"email"`
	Mobile   string `json:"mobile"`
	Avatar   string `json:"avatar" validate:"url"`
	Captcha  string `json:"captcha" gorm:"-" validate:"len=6"`
	Key      string `json:"key" gorm:"-"`
}

func (a Admin) TableName() string {
	return "admin"
}

// 获取用户信息
func GetUserInfo(ctx *fiber.Ctx, uid int64) (admin Admin, err error) {
	err = db(ctx.Context()).Where("id = ?", uid).First(&admin).Error
	return
}

// 修改个人信息
func UpdateInfo(ctx *fiber.Ctx, uid int64, name, avatar, email, phone string) error {
	return db(ctx.Context()).Table(Admin{}.TableName()).Where("id = ?", uid).Updates(map[string]interface{}{
		"name":       name,
		"avatar":     avatar,
		"email":      email,
		"mobile":     phone,
		"updated_at": time.Now().Unix()}).Error
}

// 修改密码
func ResetPassword(ctx *fiber.Ctx, uid int64, pass string) error {
	return db(ctx.Context()).Table(Admin{}.TableName()).
		Where("id = ?", uid).Update("password", pass).Error
}
