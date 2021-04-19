package init

import (
	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/internal/database"
)

func Setup() {
	var err error
	f := func(ff func() error) {
		if err != nil {
			return
		}
		err = ff()
	}
	f(configs.New)
	f(database.New)
	if err != nil {
		panic(err)
	}
}

func Close() {
	database.Close()
}
