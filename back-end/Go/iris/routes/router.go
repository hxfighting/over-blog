package routes

import (
	"blog/controllers/backend"
	"blog/controllers/frontend"
	"blog/service"
	template "blog/views"
	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"strings"
)

func RegisterRoutes(app *iris.Application) {
	registerApiRoutes(app)
	registerWebRoutes(app)
}

/**
注册api路由
*/
func registerApiRoutes(app *iris.Application) {
	app.Get("/api/captcha", backend.GetCaptcha)

	adminNeedAuth := app.Party("/api/admin", service.GetJWTHandler().Serve)

	//图片上传
	adminNeedAuth.Post("/upload", backend.Upload)
	//后台首页统计
	adminNeedAuth.Get("/count", backend.GetTotalCount)
	//刷新token
	adminNeedAuth.Get("/token", backend.RefreshToken)

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

	//标签组
	tagGroup := adminNeedAuth.Party("/tag")
	{
		//获取标签列表
		tagGroup.Get("/", backend.GetTagList)
		//添加标签
		tagGroup.Post("/", backend.AddTag)
		//修改标签
		tagGroup.Put("/", backend.UpdateTag)
		//删除标签
		tagGroup.Delete("/", backend.DeleteTag)
	}

	//用户组
	userGroup := adminNeedAuth.Party("/user")
	{
		//获取用户列表
		userGroup.Get("/", backend.GetUserList)
		//修改会员信息
		userGroup.Put("/", backend.UpdateUser)
		//删除会员
		userGroup.Delete("/", backend.DeleteUser)
	}

	//轮播图组
	rotationGroup := adminNeedAuth.Party("/rotation")
	{
		//获取轮播图列表
		rotationGroup.Get("/", backend.GetRotationList)
		//添加轮播图
		rotationGroup.Post("/", backend.AddRotation)
		//修改轮播图
		rotationGroup.Put("/", backend.UpdateRotation)
		//删除轮播图
		rotationGroup.Delete("/", backend.DeleteRotation)
	}

	//错误日志组
	webErrorGroup := adminNeedAuth.Party("/error")
	{
		//获取错误日志列表
		webErrorGroup.Get("/", backend.GetWebErrorList)
		//删除错误日志
		webErrorGroup.Delete("/", backend.DeleteWebError)
	}

	//文章组
	articleGroup := adminNeedAuth.Party("/article")
	{
		//获取文章列表
		articleGroup.Get("/", backend.GetArticleList)
		//添加文章
		articleGroup.Post("/", backend.AddArticle)
		//修改文章
		articleGroup.Put("/", backend.UpdateArticle)
		//删除文章
		articleGroup.Delete("/", backend.DeleteArticle)
	}

	adminNoAuth := app.Party("/api/admin")
	{
		//登录
		adminNoAuth.Post("/login", backend.Login)
		//添加错误日志
		adminNoAuth.Post("/error", backend.AddWebError)
	}
}

/**
注册web路由
*/
func registerWebRoutes(app *iris.Application) {
	csrfMiddleware := csrf.Protect([]byte("893263524e68086c9ac536e1c638d7da"))
	frontendRoute := app.Party("/", csrfMiddleware, initTemplateCsrfToken, favicon)
	{
		//首页
		frontendRoute.Get("/", frontend.Index)
		//文章搜索
		frontendRoute.Get("/search", frontend.SearchArticle)
		//获取博客统计
		frontendRoute.Get("/getBlogCount", frontend.GetBlogCount)
		//联系我
		frontendRoute.Get("/contact", frontend.GetContactPage)
		//说说页面
		frontendRoute.Get("/chat", frontend.GetChatPage)
		//申请友链
		frontendRoute.Post("/link", frontend.AddLink)
		//添加留言
		frontendRoute.Post("/contact", frontend.AddContact)
		//添加评论
		frontendRoute.Post("/comment", frontend.AddComment)
		//文章页
		frontendRoute.Get("/article/{id:int min(1) else 404}", frontend.GetArticleDetail)
		//根据标签获取文章列表
		frontendRoute.Get("/tag/{id:int min(1) else 404}", frontend.GetArticleByTag)
		//根据文章分类获取文章
		frontendRoute.Get("/category/{id:int min(1) else 404}", frontend.GetArticleByCategory)
	}

	//退出登录
	app.Get("/logout", frontend.Logout)

	wechatRoute := app.Party("/wechat", favicon)
	{
		//获取小程序scene
		wechatRoute.Get("/scene", frontend.GetScene)
		//获取登录结果
		wechatRoute.Get("/status", frontend.GetLoginResult)
		//获取小程序二维码
		wechatRoute.Get("/qrcode/{scene}", frontend.GetQrCodeForWeChat)
	}

	oauthRoute := app.Party("/oauth", favicon)
	{
		//三方授权
		oauthRoute.Get("/redirectToProvider/{service:string}", frontend.Oauth)
		//三方授权回调
		oauthRoute.Get("/handleOauth/{service:string}", frontend.OauthCallback)
	}

	//微信小程序登录
	app.Post("/home/wechat", frontend.WeChatLogin)
}

/**
csfrtoken赋值到模板
*/
func initTemplateCsrfToken(ctx context.Context) {
	token := csrf.Token(ctx)
	template.CsrfToken = token
	ctx.Next()
}

/**
检查是否登录
*/
func checkLogin(ctx context.Context) {
	if auth, _ := frontend.Sess.Start(ctx).GetBoolean("is_login"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.Next()
}

func favicon(ctx context.Context) {
	if strings.Contains(ctx.FullRequestURI(), "/favicon.ico") {
		return
	}
	ctx.Next()
}
