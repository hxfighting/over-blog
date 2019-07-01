package routes

import (
	"blog/controllers/backend"
	"blog/service"
	"github.com/kataras/iris"
)

func RegisterRoutes(app *iris.Application) {
	adminNeedAuth := app.Party("/api/admin", service.GetJWTHandler().Serve)
	{
		adminNeedAuth.Post("/logout", backend.Logout)
	}

	adminNoAuth := app.Party("/api/admin")
	{
		adminNoAuth.Post("/login", backend.Login)
	}
}
