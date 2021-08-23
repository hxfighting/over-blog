package event

import (
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/rs/zerolog"
)

// MessageBus implements publish/subscribe messaging paradigm
type MessageBus interface {
	// Publish publishes arguments to the given topic subscribers
	// Publish block only when the buffer of one of the subscribers is full.
	Publish(topic string, args ...interface{})
	// Close unsubscribe all handlers
	Close()
	// CloseTopic unsubscribe all handlers from given topic
	CloseTopic(topic string)
	// Subscribe subscribes to the given topic
	Subscribe(topic string, fn interface{}) error
	// Unsubscribe unsubscribe handler from the given topic
	Unsubscribe(topic string, fn interface{}) error
}

type Logger interface {
	Error(msg string)
}

type Option func(*option)

type option struct {
	logger           Logger
	handlerQueueSize uint
	concurrent       byte
}

type handlersMap map[string][]*handler

type handler struct {
	callback reflect.Value
	queue    chan []reflect.Value
}

type messageBus struct {
	handlerQueueSize uint
	handlers         handlersMap
	mtx              sync.RWMutex
	logger           Logger
	g                sync.WaitGroup
	concurrent       byte
}

func (b *messageBus) Publish(topic string, args ...interface{}) {
	defer logErr(b, false)
	rArgs := buildHandlerArgs(args)

	b.mtx.RLock()
	defer b.mtx.RUnlock()

	if hs, ok := b.handlers[topic]; ok {
		for _, h := range hs {
			h.queue <- rArgs
		}
	}
}

func (b *messageBus) Subscribe(topic string, fn interface{}) error {
	if err := isValidHandler(fn); err != nil {
		return err
	}

	h := &handler{
		callback: reflect.ValueOf(fn),
		queue:    make(chan []reflect.Value, b.handlerQueueSize),
	}

	for i := 0; i < int(b.concurrent); i++ {
		b.g.Add(1)
		go func() {
			defer logErr(b, true)
			for args := range h.queue {
				h.callback.Call(args)
			}
		}()
	}

	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.handlers[topic] = append(b.handlers[topic], h)

	return nil
}

func logErr(b *messageBus, decreaseGoroutine bool) {
	if decreaseGoroutine {
		b.g.Done()
	}
	if err := recover(); err != nil {
		b.logger.Error(fmt.Sprintf("event err: %v", err))
	}
}

func (b *messageBus) Unsubscribe(topic string, fn interface{}) error {
	defer logErr(b, false)
	if err := isValidHandler(fn); err != nil {
		return err
	}

	rv := reflect.ValueOf(fn)

	b.mtx.Lock()
	defer b.mtx.Unlock()

	if _, ok := b.handlers[topic]; ok {
		for i, h := range b.handlers[topic] {
			if h.callback == rv {
				close(h.queue)

				if len(b.handlers[topic]) == 1 {
					delete(b.handlers, topic)
				} else {
					b.handlers[topic] = append(b.handlers[topic][:i], b.handlers[topic][i+1:]...)
				}
			}
		}

		return nil
	}

	return fmt.Errorf("topic %s doesn't exist", topic)
}

func (b *messageBus) Close() {
	defer logErr(b, false)
	b.mtx.Lock()
	for topic, handlers := range b.handlers {
		for _, h := range handlers {
			close(h.queue)
		}
		delete(b.handlers, topic)
	}
	b.mtx.Unlock()
	b.g.Wait()
}

func (b *messageBus) CloseTopic(topic string) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if _, ok := b.handlers[topic]; ok {
		for _, h := range b.handlers[topic] {
			close(h.queue)
		}

		delete(b.handlers, topic)

		return
	}
}

func isValidHandler(fn interface{}) error {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		return fmt.Errorf("%s is not a reflect.Func", reflect.TypeOf(fn))
	}

	return nil
}

func buildHandlerArgs(args []interface{}) []reflect.Value {
	reflectedArgs := make([]reflect.Value, 0)

	for _, arg := range args {
		reflectedArgs = append(reflectedArgs, reflect.ValueOf(arg))
	}

	return reflectedArgs
}

// New creates new MessageBus
// handlerQueueSize sets buffered channel length per subscriber
func New(options ...Option) MessageBus {
	opt := new(option)
	for _, f := range options {
		f(opt)
	}
	if opt.handlerQueueSize == 0 {
		panic("handlerQueueSize has to be greater then 0")
	}
	if opt.logger == nil {
		opt.logger = getLogger()
	}

	return &messageBus{
		concurrent:       opt.concurrent,
		logger:           opt.logger,
		handlerQueueSize: opt.handlerQueueSize,
		handlers:         make(handlersMap),
	}
}

func WithConcurrent(concurrent byte) Option {
	return func(o *option) {
		o.concurrent = concurrent
	}
}

func WithQueueSize(size uint) Option {
	return func(o *option) {
		o.handlerQueueSize = size
	}
}

func WithLogger(logger Logger) Option {
	return func(o *option) {
		o.logger = logger
	}
}

type log struct {
	zerolog.Logger
}

func (l *log) Error(msg string) {
	l.Logger.Error().Str("msg", msg).Send()
}

func getLogger() Logger {
	return &log{Logger: zerolog.New(os.Stderr).With().Caller().Timestamp().Logger()}
}
