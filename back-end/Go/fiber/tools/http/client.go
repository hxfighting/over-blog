package http

import (
	"sync"

	"github.com/valyala/fasthttp"

	"github.com/ohdata/blog/configs"
)

var (
	HTTPClient fasthttp.Client
	once       sync.Once
)

func New() {
	cfg := configs.Config.HTTPClient
	once.Do(func() {
		HTTPClient = fasthttp.Client{
			NoDefaultUserAgentHeader:  true,
			MaxConnsPerHost:           cfg.MaxConnsPerHost,
			MaxIdleConnDuration:       cfg.MaxIdleConnDuration,
			MaxConnDuration:           cfg.MaxConnDuration,
			ReadTimeout:               cfg.ReadTimeout,
			WriteTimeout:              cfg.WriteTimeout,
			MaxConnWaitTimeout:        cfg.MaxConnWaitTimeout,
			MaxIdemponentCallAttempts: cfg.MaxIdemponentCallAttempts,
		}
	})
}
