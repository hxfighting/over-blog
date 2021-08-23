package models

import (
	"net"

	"github.com/ohdata/blog/internal/pkg/geo"
	"github.com/ohdata/blog/tools/log"
	"github.com/ohdata/blog/tools/util"
)

type User struct {
	ID            *int64 `json:"id" validate:"gt=0"`
	Type          uint   `json:"type"`
	Name          string `json:"name"`
	OpenID        string `json:"openid" gorm:"column:openid"`
	AccessToken   string `json:"access_token"`
	LastLoginIp   string `json:"last_login_ip"`
	RememberToken string `json:"remember_token"`
	LoginTimes    *int64 `json:"login_times"`
	IsAdmin       *int8  `json:"is_admin" validate:"oneof=0 1" mapstructure:"is_admin"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar"`
	CreatedUnix   int64  `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix   int64  `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt     string `json:"created_at" gorm:"-"`
	UpdatedAt     string `json:"updated_at" gorm:"-"`
}

func (User) TableName() string {
	return "user"
}

func (this *User) AfterFind() {
	this.CreatedAt = util.GetDateTime(this.CreatedUnix, util.YMDHIS)
	this.UpdatedAt = util.GetDateTime(this.UpdatedUnix, util.YMDHIS)
	this.LastLoginIp = getAddress(this.LastLoginIp)
}

/**
获取IP对应的地址
*/
func getAddress(ip string) string {
	if geo.GeoDB == nil {
		return "未知地方"
	}
	right_ip := net.ParseIP(ip)
	record, err := geo.GeoDB.City(right_ip)
	if err != nil {
		log.Log.Error().Err(err).Send()
		return "未知地方"
	}
	var city, country, province string
	iso_code := record.Country.IsoCode
	if iso_code == "TW" {
		country = "中国"
		city = "台湾"
	} else if iso_code == "HK" {
		country = "中国"
		city = "香港"
	} else if iso_code == "MO" {
		country = "中国"
		city = "澳门"
	} else {
		city = record.City.Names["zh-CN"]
		country = record.Country.Names["zh-CN"]
	}
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}
	return country + province + city
}
