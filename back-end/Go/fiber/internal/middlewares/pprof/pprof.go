package pprof

import (
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/ohdata/blog/configs"
)

// Set pprof adaptors
var (
	pprofIndex        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
	pprofCmdline      = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	pprofProfile      = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	pprofSymbol       = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	pprofTrace        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	pprofAllocs       = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("allocs").ServeHTTP)
	pprofBlock        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("block").ServeHTTP)
	pprofGoroutine    = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("goroutine").ServeHTTP)
	pprofHeap         = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("heap").ServeHTTP)
	pprofMutex        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("mutex").ServeHTTP)
	pprofThreadcreate = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("threadcreate").ServeHTTP)
)

func New() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		// We are only interested in /debug/pprof routes
		if len(path) < 12 || !strings.HasPrefix(path, "/debug/pprof") {
			return ctx.Next()
		}
		if !configs.Config.Server.Debug {
			token := ctx.Query("token")
			if token != configs.Config.Server.PProfToken {
				return ctx.JSON(fiber.Map{"msg": "invalid route"})
			}
		}
		// Switch to original path without stripped slashes
		switch path {
		case "/debug/pprof/":
			pprofIndex(ctx.Context())
		case "/debug/pprof/cmdline":
			pprofCmdline(ctx.Context())
		case "/debug/pprof/profile":
			pprofProfile(ctx.Context())
		case "/debug/pprof/symbol":
			pprofSymbol(ctx.Context())
		case "/debug/pprof/trace":
			pprofTrace(ctx.Context())
		case "/debug/pprof/allocs":
			pprofAllocs(ctx.Context())
		case "/debug/pprof/block":
			pprofBlock(ctx.Context())
		case "/debug/pprof/goroutine":
			pprofGoroutine(ctx.Context())
		case "/debug/pprof/heap":
			pprofHeap(ctx.Context())
		case "/debug/pprof/mutex":
			pprofMutex(ctx.Context())
		case "/debug/pprof/threadcreate":
			pprofThreadcreate(ctx.Context())
		default:
			// pprof index only works with trailing slash
			return ctx.Redirect("/debug/pprof/", http.StatusFound)
		}
		return nil
	}
}
