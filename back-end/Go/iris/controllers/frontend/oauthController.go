package frontend

import (
	"blog/config"
	"blog/database"
	"blog/models"
	"blog/service"
	"blog/service/oauth"
	template "blog/views"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

/**
qq、微博授权
*/
func Oauth(ctx iris.Context) {
	refer := ctx.GetReferrer().URL
	oauthService := strings.Trim(ctx.Params().GetEscape("service"), "")
	if oauthService == "" {
		return
	}
	oauthConfig := oauth.OauthConfig{}
	oauthUrl := ""
	if oauthService == "qq" {
		oauthConfig.ClientID = config.GetConfig("qq.client_id").(string)
		oauthConfig.Secret = config.GetConfig("qq.secret").(string)
		oauthConfig.Callback = config.GetConfig("qq.callback").(string)
		oauthUrl = oauthConfig.GetOauthUrlFromQQ(fmt.Sprintf("%d", time.Now().UnixNano()))
	} else if oauthService == "weibo" {
		oauthConfig.ClientID = config.GetConfig("weibo.client_id").(string)
		oauthConfig.Secret = config.GetConfig("weibo.secret").(string)
		oauthConfig.Callback = config.GetConfig("weibo.callback").(string)
		oauthUrl = oauthConfig.GetOauthUrlFromWeiBo()
	}
	if oauthConfig.ClientID == "" || oauthConfig.Secret == "" || oauthConfig.Callback == "" {
		service.Log.Error("缺少qq授权配置")
		return
	}
	Sess.Start(ctx).Set("target_url", refer)
	ctx.Redirect(oauthUrl)
}

/**
授权回调
*/
func OauthCallback(ctx iris.Context) {
	ip := ctx.Request().Header.Get("X-Real-Ip")
	code := ctx.URLParamTrim("code")
	oauthService := strings.Trim(ctx.Params().GetEscape("service"), "")
	if oauthService == "" || code == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	oauth_type := map[string]uint{
		"weibo": 3,
		"qq":    1,
	}
	oauthConfig := oauth.OauthConfig{}
	if oauthService == "qq" {
		oauthConfig.ClientID = config.GetConfig("qq.client_id").(string)
		oauthConfig.Secret = config.GetConfig("qq.secret").(string)
		oauthConfig.Callback = config.GetConfig("qq.callback").(string)
	} else if oauthService == "weibo" {
		oauthConfig.ClientID = config.GetConfig("weibo.client_id").(string)
		oauthConfig.Secret = config.GetConfig("weibo.secret").(string)
		oauthConfig.Callback = config.GetConfig("weibo.callback").(string)
	}
	var user map[string]string
	var e error
	if oauthService == "qq" {
		user, e = oauthConfig.GetUserFromQQ(code)
	} else {
		user, e = oauthConfig.GetUserFromWeiBo(code)
	}
	if e != nil {
		service.Log.Error(e.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
		return
	}
	var loginTimes int64 = 1
	var IsAdmin int8 = 0
	user_model := models.User{}
	var res *gorm.DB
	database.Db.Where("openid = ? and type = ?", user["openid"], oauth_type[oauthService]).First(&user_model)
	if user_model.ID != nil {
		user["login_times"] = fmt.Sprintf("%d", *user_model.LoginTimes+1)
		user["last_login_ip"] = ip
		res = database.Db.Model(&user_model).Updates(user)
	} else {
		user_model.Name = user["name"]
		user_model.Avatar = user["avatar"]
		user_model.CreatedUnix = time.Now().Unix()
		user_model.UpdatedUnix = time.Now().Unix()
		user_model.OpenID = user["openid"]
		user_model.AccessToken = user["access_token"]
		user_model.Type = oauth_type[oauthService]
		user_model.LastLoginIp = ip
		user_model.LoginTimes = &loginTimes
		user_model.IsAdmin = &IsAdmin
		res = database.Db.Create(&user_model)
	}
	if res.Error != nil {
		service.Log.Error(res.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
		return
	}
	user["is_admin"] = fmt.Sprintf("%d", *user_model.IsAdmin)
	user["id"] = fmt.Sprintf("%d", *user_model.ID)
	user["email"] = user_model.Email
	Sess.Start(ctx).Set("is_login", true)
	template.AuthInfo = user
	refer := Sess.Start(ctx).GetString("target_url")
	if refer == "" {
		app_url := config.GetConfig("app.url").(string)
		ctx.Redirect(app_url)
		return
	}
	ctx.Redirect(refer)
}

/**
退出登录
*/
func Logout(ctx iris.Context) {
	Sess.Start(ctx).Set("is_login", false)
	template.AuthInfo = make(map[string]string)
	Response.RenderSuccess(ctx, "退出登录成功！", nil)
}
