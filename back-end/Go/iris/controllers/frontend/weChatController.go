package frontend

import (
	"blog/config"
	"blog/controllers"
	"blog/database"
	"blog/models"
	"blog/service"
	"blog/service/wechat"
	template "blog/views"
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

type weChatLoginStruct struct {
	Code   string `json:"code" form:"code" validate:"gte=2"`
	Avatar string `json:"avatar" form:"avatar" validate:"url"`
	Name   string `json:"name" form:"name" validate:"gte=2"`
	Sence  string `json:"sence" form:"sence" validate:"gte=2"`
}

/**
获取小程序scene
*/
func GetScene(ctx iris.Context) {
	randStr := "sdasqrqweqw1232342354+123.*-+" + fmt.Sprintf("%d", time.Now().UnixNano())
	scene := fmt.Sprintf("%x", md5.Sum([]byte(randStr)))
	Response.RenderSuccess(ctx, "获取scene成功！", scene)
}

/**
获取微信小程序二维码
*/
func GetQrCodeForWeChat(ctx iris.Context) {
	scene := strings.Trim(ctx.Params().GetEscape("scene"), "")
	if scene == "" {
		return
	}
	coderConfig := wechat.QRCoder{Width: 200, Scene: scene}
	wechatConfig := wechat.Config{
		ClientID: config.GetConfig("miniProgram.client_id").(string),
		Secret:   config.GetConfig("miniProgram.secret").(string),
	}
	if wechatConfig.Secret == "" || wechatConfig.ClientID == "" {
		return
	}
	qrcode, err := wechatConfig.GetQrCode(coderConfig)
	if err != nil {
		service.Log.Error(err.Error())
		return
	}
	ctx.ContentType("image/jpeg")
	_, err = ctx.Write(qrcode)
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦")
	}
}

/**
获取登录结果
*/
func GetLoginResult(ctx iris.Context) {
	if ctx.IsAjax() {
		scene := strings.Trim(ctx.URLParamEscape("scene"), "")
		if scene == "" {
			Response.RenderError(ctx, "参数缺失，请重试", nil)
			return
		}
		res := service.Redis.Get(scene)
		if res.Err() == nil {
			Response.RenderSuccess(ctx, "登录成功", nil)
			return
		}
	} else {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("404")
		return
	}
}

/**
微信小程序登录
*/
func WeChatLogin(ctx iris.Context) {
	ip := ctx.Request().Header.Get("X-Real-Ip")
	needField := weChatLoginStruct{}
	fields := []string{"code", "avatar", "name", "sence"}
	validateFields := []string{"Name", "Code", "Avatar", "Sence"}
	_, err := controllers.GetRightModel(ctx, &needField, fields, validateFields)
	if err != nil {
		service.Log.Error(err.Error())
		Response.Code = 500
		Response.RenderError(ctx, err.Error(), nil)
		return
	}
	wechatConfig := wechat.Config{
		ClientID: config.GetConfig("miniProgram.client_id").(string),
		Secret:   config.GetConfig("miniProgram.secret").(string),
	}
	if wechatConfig.Secret == "" || wechatConfig.ClientID == "" {
		service.Log.Error("缺少微信小程序配置")
		Response.Code = 500
		Response.RenderError(ctx, "缺少微信小程序配置！", nil)
		return
	}
	openid, sessionKey, err := wechatConfig.GetOpenID(needField.Code)
	if err != nil {
		service.Log.Error(err.Error())
		Response.Code = 500
		Response.RenderError(ctx, "登录失败！", nil)
		return
	}
	var loginTimes int64 = 1
	avatar := strings.Replace(needField.Avatar, "http", "https", -1)
	user := make(map[string]string)
	user["name"] = needField.Name
	user["avatar"] = avatar
	user["access_token"] = sessionKey
	user["openid"] = openid
	user_model := models.User{}
	var res *gorm.DB
	database.Db.Where("openid = ? and type = ?", openid, 2).First(&user_model)
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
		user_model.Type = 2
		user_model.LastLoginIp = ip
		user_model.AccessToken = sessionKey
		user_model.LoginTimes = &loginTimes
		res = database.Db.Create(&user_model)
	}
	if res.Error != nil {
		service.Log.Error(res.Error.Error())
		Response.Code = 500
		Response.RenderError(ctx, "登录失败！", nil)
		return
	}
	user["is_admin"] = fmt.Sprintf("%d", *user_model.IsAdmin)
	user["id"] = fmt.Sprintf("%d", *user_model.ID)
	user["email"] = user_model.Email
	Sess.Start(ctx).Set("is_login", true)
	template.AuthInfo = user
	resRedis := service.Redis.Set(needField.Sence, 1, 60*time.Second)
	if resRedis.Err() != nil {
		service.Log.Error(resRedis.Err().Error())
		Response.Code = 500
		Response.RenderError(ctx, "登录失败！", nil)
		return
	}
	Response.RenderSuccess(ctx, "登录成功！", nil)
}
