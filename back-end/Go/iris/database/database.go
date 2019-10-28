package database

import (
	"blog/config"
	"blog/helper"
	"blog/service"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type MyDb struct {
	User     string
	Password string
	Port     int64
	Host     string
	Database string
}

var (
	Db *gorm.DB
)

func NewDB() {
	my_db := MyDb{
		config.GetConfig("database.username").(string),
		config.GetConfig("database.password").(string),
		config.GetConfig("database.port").(int64),
		config.GetConfig("database.host").(string),
		config.GetConfig("database.db").(string),
	}
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		my_db.User, my_db.Password, my_db.Host, my_db.Port, my_db.Database)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)

	if helper.CheckDebug() {
		db.LogMode(true)
	}
	err = db.DB().Ping()
	if err != nil {
		service.Log.Error(err.Error())
	}
	db.Callback().Create().Remove("gorm:update_time_stamp")
	db.Callback().Update().Remove("gorm:update_time_stamp")
	db.SingularTable(true)
	Db = db
}
