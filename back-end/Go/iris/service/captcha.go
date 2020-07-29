package service

import (
	"blog/helper"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

/**
生成验证码
*/
func GenerateCaptcha() (string, string) {
	var digit = base64Captcha.NewDriverDigit(80,240,6,0.7,80)
	var driver base64Captcha.Driver
	driver = digit
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		Log.Error(err.Error())
		return "", ""
	}
	return id, b64s
}

/**
验证验证码
*/
func ValidateCaptcha(key, captcha string) bool {
	if !helper.CheckDebug() {
		return store.Verify(key, captcha, true)
	}
	return true
}
