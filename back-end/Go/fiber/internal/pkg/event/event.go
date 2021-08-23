package event

import (
	"sync"

	"github.com/ohdata/blog/configs"
	messagebus "github.com/ohdata/blog/pkg/event"
)

var (
	EventBus messagebus.MessageBus
	once     sync.Once
)

type Event string

func (e Event) String() string {
	return string(e)
}

const (
	CategoryCreateEvent Event = "category_create"
)

func New() (err error) {
	once.Do(func() {
		EventBus = messagebus.New(
			messagebus.WithQueueSize(uint(configs.Config.Event.Size)),
			messagebus.WithConcurrent(byte(configs.Config.Event.Concurrent)))
	})
	return
}

func Close() {
	EventBus.Close()
}
