package database

import (
	"blog/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MyDb struct {
	User     string
	Password string
	Port     int
	Host     string
	Database string
}

var (
	Db = getDb()
)

func getDb() *gorm.DB {
	my_db := MyDb{
		config.GetConfig("database.username").(string),
		config.GetConfig("database.password").(string),
		config.GetConfig("database.port").(int),
		config.GetConfig("database.host").(string),
		config.GetConfig("database.db").(string),
	}
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		my_db.User, my_db.Password, my_db.Host, my_db.Port, my_db.Database)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {

	}
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(5)
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	return db, nil
}
