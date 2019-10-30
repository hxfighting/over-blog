package backend

import (
	"blog/config"
	"blog/controllers/frontend"
	"blog/helper"
	"blog/service"
	"blog/service/sms"
	"github.com/kataras/iris/v12"
	"math/rand"
	"strings"
	"time"
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

/**
发送短信
*/
func SendSms(ctx iris.Context) {
	token := strings.Trim(ctx.URLParamEscape("token"), "")
	right_token := config.GetConfig("sms.access_token").(string)
	if token == "" || right_token == "" || right_token != token {
		return
	}
	var smsService sms.SmsInterface
	ali := sms.Ali{
		AppID:    config.GetConfig("sms.sms_key").(string),
		Secret:   config.GetConfig("sms.sms_secret").(string),
		SignName: config.GetConfig("sms.sms_sign_name").(string),
	}
	for i := 0; i < 2; i++ {
		if i == 0 {
			rand.Seed(time.Now().Unix())
			num := rand.Intn(10)
			ali.Data = map[string]interface{}{"code": num}
			ali.TemplateCode = config.GetConfig("sms.phone_one_template").(string)
			ali.PhoneNumber = config.GetConfig("sms.phone_one").(string)
		} else {
			days, _ := helper.GetDateDiffDay("1994-03-10", time.Now().Format(helper.YMD))
			ali.Data = map[string]interface{}{"number": days}
			ali.TemplateCode = config.GetConfig("sms.phone_two_template").(string)
			ali.PhoneNumber = config.GetConfig("sms.phone_two").(string)
		}
		smsService = ali
		smsService.Send()
	}
}
