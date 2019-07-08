package routes

import (
	"blog/controllers/backend"
	"blog/service"
	"github.com/kataras/iris"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/api/captcha", backend.GetCaptcha)

	adminNeedAuth := app.Party("/api/admin", service.GetJWTHandler().Serve)
	{
		//获取个人信息
		adminNeedAuth.Get("/", backend.GetUserInfo)
		//修改个人信息
		adminNeedAuth.Put("/", backend.UpdateInfo)
		//退出登录
		adminNeedAuth.Post("/logout", backend.Logout)
		//修改密码
		adminNeedAuth.Put("/password", backend.ResetPassword)
	}

	adminNoAuth := app.Party("/api/admin")
	{
		//登录
		adminNoAuth.Post("/login", backend.Login)
	}
}
