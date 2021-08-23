package init

import (
	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/internal/pkg/cache"
	"github.com/ohdata/blog/internal/pkg/database"
	"github.com/ohdata/blog/internal/pkg/event"
	"github.com/ohdata/blog/internal/pkg/geo"
	"github.com/ohdata/blog/internal/subscribes"
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
	f(geo.New)
	f(event.New)
	f(registerEvents)
	f(cache.New)
	f(database.New)
	if err != nil {
		panic(err)
	}
}

func Close() {
	event.Close()
	geo.Close()
	cache.Close()
	database.Close()
}

func registerEvents() (err error) {
	f := func(topic string, fn interface{}) {
		if err != nil {
			return
		}
		err = event.EventBus.Subscribe(topic, fn)
	}
	f(event.CategoryCreateEvent.String(), subscribes.CategoryCreateSubscribe)
	return
}
