package models

import (
	"blog/database"
	"blog/service"
	"errors"
	"log"
)

type Admin struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Password  string `gorm:"-" json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

/**
登录
*/
func (this *Admin) Login() (token string, err error) {
	s, _ := service.PasswordHash(this.Password)
	log.Println(s, this)
	existAdmin := Admin{}
	database.Db.Where("name = ?", this.Name).First(&existAdmin)
	if existAdmin.ID == 0 {
		return "", errors.New("登录失败,用户名或密码错误")
	}
	if !service.PasswordVerify(this.Password, existAdmin.Password) {
		return "", errors.New("登录失败,用户名或密码错误")
	}
	token = service.GenerateToken(existAdmin.ID)
	return token, nil
}
