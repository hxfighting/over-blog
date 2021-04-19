package middlewares

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"

	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/internal/middlewares/pprof"
	"github.com/ohdata/blog/tools"
	"github.com/ohdata/blog/tools/util"
)

func RegisterMiddleware(app *fiber.App) {
	// recover
	app.Use(recover.New())
	// request id
	app.Use(requestid.New())
	// pprof
	app.Use(pprof.New())
	// limiter
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1" || configs.Config.Server.Debug
		},
		Max: configs.Config.RateLimit.Max,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		Expiration: configs.Config.RateLimit.Expire,
		LimitReached: func(ctx *fiber.Ctx) error {
			return tools.Error(ctx, http.StatusText(http.StatusTooManyRequests))
		},
	}))
	// etag
	app.Use(etag.New(etag.Config{
		Next: func(ctx *fiber.Ctx) bool {
			return ctx.IP() == "127.0.0.1" || configs.Config.Server.Debug
		},
	}))
	// monitor
	app.Get("/blog/monitor", func(ctx *fiber.Ctx) error {
		if !configs.Config.Server.Debug {
			token := ctx.Query("token")
			if token != configs.Config.Server.PProfToken {
				return tools.Error(ctx, http.StatusText(http.StatusNotFound))
			}
		}
		return ctx.Next()
	}, monitor.New())
	// request log
	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return !configs.Config.Server.Debug
		},
		Format:     "[${time}] ${status} - ${latency} ${method} ${path} ${query}\n",
		TimeFormat: util.YMDHIS,
		TimeZone:   "Asia/Chongqing",
		Output:     os.Stdout,
	}))
	// secure
	app.Use(helmet.New())
	// cors
	corsConfig := configs.Config.CORS
	app.Use(cors.New(cors.Config{
		Next: func(c *fiber.Ctx) bool {
			return !corsConfig.Enable
		},
		AllowOrigins:     corsConfig.AllowOrigins,
		AllowMethods:     corsConfig.AllowMethods,
		AllowHeaders:     corsConfig.AllowHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
		ExposeHeaders:    corsConfig.ExposeHeaders,
		MaxAge:           corsConfig.MaxAge,
	}))
}
