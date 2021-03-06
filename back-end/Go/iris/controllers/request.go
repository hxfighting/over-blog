package controllers

import (
	"blog/service"
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"strings"
)

/**
获取需要的数据
*/
func GetRequestField(ctx iris.Context, fields ...string) (map[string]interface{}, string, error) {
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
func GetRightModel(ctx iris.Context, model interface{}, fields, validateField []string) (string, error) {
	requestData, jsonData, err := GetRequestField(ctx, fields...)
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
