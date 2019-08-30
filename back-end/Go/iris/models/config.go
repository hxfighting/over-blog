package models

import (
	"blog/database"
	"blog/helper"
)

type Config struct {
	ID          *int64  `json:"id"`
	Title       *string `json:"title"`
	Name        *string `json:"name"`
	Val         *string `json:"val"`
	Type        *int64  `json:"type"`
	CreatedUnix int64   `json:"created_unix" gorm:"column:created_at"`
	UpdatedUnix int64   `json:"updated_unix" gorm:"column:updated_at"`
	CreatedAt   string  `json:"created_at" gorm:"-"`
	UpdatedAt   string  `json:"updated_at" gorm:"-"`
}

func (Config) TableName() string {
	return "web_config"
}

func (this *Config) AfterFind() {
	this.CreatedAt = helper.GetDateTime(this.CreatedUnix, helper.YMDHIS)
	this.UpdatedAt = helper.GetDateTime(this.UpdatedUnix, helper.YMDHIS)
}

const FOOTER_TYPE int64 = 2

/**
获取配置列表
*/
func GetConfigList(pageNum, pageSize, type_id int64) map[string]interface{} {
	var data = make(map[string]interface{})
	configs := []Config{}
	var db = database.Db
	var total int64
	if type_id != 0 {
		db = db.Where("type = ?", type_id)
	}
	db.Table("web_config").Count(&total)
	if total > 0 {
		db.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&configs)
	}
	data["list"] = configs
	data["total"] = total
	return data
}

/**
添加配置
 */
func AddConfig(config *Config) map[string]interface{} {
	data := make(map[string]interface{})
	data["flag"] = false
	checkRes := checkConfig(*config.Name, *config.Type)
	if !checkRes["flag"].(bool) {
		data["msg"] = checkRes["msg"]
		return data
	}
	database.Db.Where("name = ? and type = ?",*config.Name,*config.Type).FirstOrCreate(&config)
	if config.ID != nil {
		data["flag"] = true
		data["msg"] = "添加配置成功！"
	} else {
		data["flag"] = false
		data["msg"] = "添加配置失败，请稍后再试！"
	}
	return data
}

/**
检查配置项
*/
func checkConfig(name string, type_id int64) map[string]interface{} {
	res := make(map[string]interface{})
	res["flag"] = false
	var total int64
	database.Db.Table("web_config").Where("name = ? and type = ?", name, type_id).Count(&total)
	if type_id == FOOTER_TYPE {
		if name == "copyright" && total >= 1 {
			res["msg"] = "版权信息只能有一条"
			return res
		}
		if total >= 4 {
			res["msg"] = "footer内容每一项最多添加4个!"
			return res
		}
	}
	res["flag"] = true
	return res
}
