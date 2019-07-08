package backend

import (
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
