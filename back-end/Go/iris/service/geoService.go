package service

import (
	"blog/config"
	"github.com/oschwald/geoip2-golang"
	"log"
)

var GeoDB *geoip2.Reader

func NewGeoDb() {
	db, err := geoip2.Open(config.ConfigPath + "/geo.mmdb")
	if err != nil {
		log.Fatalln(err.Error())
	}
	GeoDB = db
}
