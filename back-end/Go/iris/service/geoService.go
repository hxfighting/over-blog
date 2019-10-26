package service

import "github.com/oschwald/geoip2-golang"

var GeoDB *geoip2.Reader

func NewGeoDb() {
	db, err := geoip2.Open("./geo.mmdb")
	if err != nil {
		Log.Error(err.Error())
	}
	GeoDB = db
}
