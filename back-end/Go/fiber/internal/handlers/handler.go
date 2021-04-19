package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/internal/handlers/backend/admin"
	"github.com/ohdata/blog/internal/middlewares/jwt"
	"github.com/ohdata/blog/tools"
)

func RegisterRoute(app *fiber.App) {

	app.Get("/favicon.ico", func(ctx *fiber.Ctx) error {
		return nil
	})

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"msg": "ok"})
	})
	handlers := []fiber.Handler{jwt.New(jwt.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return tools.Error(ctx, "非法token")
		},
		SigningKey:    configs.Config.JWT.Secret,
		SigningMethod: configs.Config.JWT.SigningAlgorithm,
		Expire:        configs.Config.JWT.Expire,
	})}

	admin.Route(app, handlers...)
}
