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

	//分类组
	categoryGroup := app.Party("/api/admin/category")
	{
		//获取分类列表
		categoryGroup.Get("/", backend.GetCategoryList)
		//添加分类
		categoryGroup.Post("/", backend.AddCategory)
		//修改分类
		categoryGroup.Put("/", backend.UpdateCategory)
		//删除分类
		categoryGroup.Delete("/", backend.DeleteCategory)
	}

	adminNoAuth := app.Party("/api/admin")
	{
		//登录
		adminNoAuth.Post("/login", backend.Login)
	}
}
