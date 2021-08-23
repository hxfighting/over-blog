package geo

import (
	_ "embed"

	"github.com/oschwald/geoip2-golang"

	"github.com/ohdata/blog/tools/log"
)

//go:embed geo.mmdb
var db []byte

var GeoDB *geoip2.Reader

func New() (err error) {
	GeoDB, err = geoip2.FromBytes(db)
	return
}

func Close() {
	if err := GeoDB.Close(); err != nil {
		log.Log.Err(err).Send()
	}
}
