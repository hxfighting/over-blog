package models

import (
	"blog/database"
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"strings"
	"time"
)

type Admin struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" validate:"required"`
	Password  string `gorm:"-" json:"password" validate:"required"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Captcha   string `json:"captcha" gorm:"-"`
	Key       string `json:"key" gorm:"-"`
}

/**
登录
*/
func (this *Admin) Login() (token string, err error) {
	existAdmin := Admin{}
	database.Db.Where("name = ?", this.Name).First(&existAdmin)
	if existAdmin.ID == 0 {
		return "", errors.New("登录失败,用户名或密码错误")
	}
	if !service.PasswordVerify(this.Password, existAdmin.Password) {
		return "", errors.New("登录失败,用户名或密码错误")
	}
	token, _, err = service.GenerateToken(existAdmin.ID, 0)
	if err != nil {
		return "", errors.New("登录失败,用户名或密码错误")
	}
	return token, nil
}

/**
获取用户信息
*/
func GetUserInfo(ctx iris.Context) (user map[string]interface{}, err error) {
	id := service.GetUserId(ctx)
	if id == 0 {
		return nil, errors.New("用户不存在！")
	}
	admin := Admin{}
	database.Db.Where("id = ?", id).First(&admin)
	if admin.ID == 0 {
		return nil, errors.New("用户不存在！")
	}
	avatarName := admin.Avatar[strings.LastIndex(admin.Avatar, "/")+1:]
	user = make(map[string]interface{})
	user["avatarName"] = avatarName
	user["user_id"] = admin.ID
	user["name"] = admin.Name
	user["access"] = []string{"super_admin"}
	user["avatar"] = admin.Avatar
	user["email"] = admin.Email
	user["mobile"] = admin.Mobile
	return
}

/**
修改个人信息
*/
func UpdateInfo(ctx iris.Context, admin *Admin) error {
	id := service.GetUserId(ctx)
	if id == 0 {
		return errors.New("用户不存在！")
	}
	admin.ID = id
	updates := database.Db.Model(admin).Updates(map[string]interface{}{"name": admin.Name,
		"avatar": admin.Avatar, "email": admin.Email, "mobile": admin.Mobile, "updated_at": time.Now().Unix()})
	if updates.Error != nil {
		return errors.New("修改个人信息失败，请稍后再试！")
	}
	return nil
}

/**
修改密码
*/
func ResetPassword(ctx iris.Context, admin *Admin) error {
	id := service.GetUserId(ctx)
	if id == 0 {
		return errors.New("用户不存在！")
	}
	admin.ID = id
	pass, err := service.PasswordHash(admin.Password)
	if err != nil {
		return errors.New("修改密码失败，请稍后再试！")
	}
	res := database.Db.Model(admin).Updates(map[string]interface{}{"password": pass, "updated_at": time.Now().Unix()})
	if res.Error != nil {
		return errors.New("修改密码失败，请稍后再试！")
	}
	return nil
}
