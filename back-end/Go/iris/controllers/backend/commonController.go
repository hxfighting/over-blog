package backend

import (
	"blog/controllers/frontend"
	"blog/service"
	"github.com/kataras/iris"
)

var response = service.Response{}

/**
生成验证码
*/
func GetCaptcha(ctx iris.Context) {
	key, captchaBase64 := service.GenerateCaptcha()
	data := make(map[string]string)
	data["img"] = captchaBase64
	data["key"] = key
	response.RenderSuccess(ctx, "登录成功！", data)
}

/**
刷新token
*/
func RefreshToken(ctx iris.Context) {
	token, e := service.RefreshToken(ctx)
	if e != nil {
		response.RenderError(ctx, e.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "获取token成功", token)
}

/**
删除缓存的key
*/
func removeFrontendCache(keys ...string) {
	del := service.Redis.Del(keys...)
	if del.Err() != nil {
		service.Log.Error(del.Err().Error())
	} else {
		frontend.InitData()
	}
}
