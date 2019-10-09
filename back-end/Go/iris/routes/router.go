package routes

import (
	"blog/controllers/backend"
	"blog/service"
	"github.com/kataras/iris"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/api/captcha", backend.GetCaptcha)

	adminNeedAuth := app.Party("/api/admin", service.GetJWTHandler().Serve)

	//管理员组
	adminGroup := adminNeedAuth.Party("/")
	{
		//获取个人信息
		adminGroup.Get("/", backend.GetUserInfo)
		//修改个人信息
		adminGroup.Put("/", backend.UpdateInfo)
		//退出登录
		adminGroup.Post("/logout", backend.Logout)
		//修改密码
		adminGroup.Put("/password", backend.ResetPassword)
	}

	//分类组
	categoryGroup := adminNeedAuth.Party("/category")
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

	//说说组
	chatGroup := adminNeedAuth.Party("/chat")
	{
		//获取说说列表
		chatGroup.Get("/", backend.GetChatList)
		//添加说说
		chatGroup.Post("/", backend.AddChat)
		//修改说说
		chatGroup.Put("/", backend.UpdateChat)
		//删除说说
		chatGroup.Delete("/", backend.DeleteChat)
	}

	//友联组
	linkGroup := adminNeedAuth.Party("/link")
	{
		//获取友联列表
		linkGroup.Get("/", backend.GetLinkList)
		//添加友联
		linkGroup.Post("/", backend.AddLink)
		//修改友联
		linkGroup.Put("/", backend.UpdateLink)
		//删除友联
		linkGroup.Delete("/", backend.DeleteLink)
	}

	//照片组
	photoGroup := adminNeedAuth.Party("/photo")
	{
		//获取照片列表
		photoGroup.Get("/", backend.GetPhotoList)
		//添加照片
		photoGroup.Post("/", backend.AddPhoto)
		//修改照片
		photoGroup.Put("/", backend.UpdatePhoto)
		//删除照片
		photoGroup.Delete("/", backend.DeletePhoto)
	}

	//评论组
	commentGroup := adminNeedAuth.Party("/comment")
	{
		//获取评论列表
		commentGroup.Get("/", backend.GetCommentList)
		//回复评论
		commentGroup.Post("/", backend.ReplyComment)
		//删除评论
		commentGroup.Delete("/", backend.DeleteComment)
	}

	//配置组
	configGroup := adminNeedAuth.Party("/config")
	{
		//获取配置列表
		configGroup.Get("/", backend.GetConfigList)
		//添加配置
		configGroup.Post("/", backend.AddConfig)
		//修改配置
		configGroup.Put("/", backend.UpdateConfig)
		//删除配置
		configGroup.Delete("/", backend.DeleteConfig)
	}

	//留言组
	contactGroup := adminNeedAuth.Party("/contact")
	{
		//获取留言列表
		contactGroup.Get("/", backend.GetContactList)
		//删除留言
		contactGroup.Delete("/", backend.DeleteContact)
		//回复留言
		contactGroup.Post("/reply", backend.ReplyContact)
	}

	adminNoAuth := app.Party("/api/admin")
	{
		//登录
		adminNoAuth.Post("/login", backend.Login)
	}
}
