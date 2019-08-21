package models

type User struct {
	ID            *int64 `json:"id"`
	Type          uint   `json:"type"`
	Name          string `json:"name"`
	OpenID        string `json:"openid" gorm:"column:openid"`
	AccessToken   string `json:"access_token"`
	LastLoginIp   string `json:"last_login_ip"`
	RememberToken string `json:"remember_token"`
	LoginTimes    *int64 `json:"login_times"`
	IsAdmin       *int8  `json:"is_admin"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar"`
	CreatedUnix   int64  `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix   int64  `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt     string `json:"created_at" gorm:"-"`
	UpdatedAt     string `json:"updated_at" gorm:"-"`
}
