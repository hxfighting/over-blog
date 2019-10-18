package backend

import (
	"blog/service"
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"strings"
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
获取需要的数据
*/
func getRequestData(ctx iris.Context, fields ...string) (map[string]interface{}, string, error) {
	jsonBytes, _ := ctx.GetBody()
	jsonData := string(jsonBytes)
	if !gjson.Valid(jsonData) {
		return nil, "", errors.New("json格式错误!")
	}
	mapData := make(map[string]interface{})
	for _, field := range fields {
		data := gjson.Get(jsonData, field).Value()
		if data == nil {
			return nil, "", errors.New("缺少参数：" + field)
		}
		mapData[field] = data
	}
	return mapData, jsonData, nil
}

/**
验证数据，获取模型
*/
func getRightModel(ctx iris.Context, model interface{}, fields, validateField []string) (string, error) {
	requestData, jsonData, err := getRequestData(ctx, fields...)
	if err != nil {
		return jsonData, err
	}
	decode := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           model,
	}
	decoder, err := mapstructure.NewDecoder(decode)
	if err != nil {
		service.Log.Error(err.Error())
		return jsonData, errors.New("参数错误！")
	}
	err = decoder.Decode(requestData)
	if err != nil {
		service.Log.Error(err.Error())
		return jsonData, errors.New("参数错误！")
	}
	err = service.Validate.StructPartial(model, validateField...)
	if err != nil {
		zhs := zh.New()
		uni := ut.New(zhs, zhs)
		trans, _ := uni.GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(service.Validate, trans)
		errs := err.(validator.ValidationErrors)
		error_msg := ""
		for _, e := range errs {
			msg := e.Translate(trans) + ","
			error_msg += msg
		}
		error_msg = strings.TrimRight(error_msg, ",")
		return jsonData, errors.New(error_msg)
	}
	return jsonData, nil
}

/**
刷新token
 */
func RefreshToken(ctx iris.Context)  {
	token, e := service.RefreshToken(ctx)
	if e!=nil{
		response.RenderError(ctx, e.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "获取token成功", token)
}
