package service

import (
	"blog/helper"
	"github.com/mojocn/base64Captcha"
)

/**
生成验证码
*/
func GenerateCaptcha() (string, string) {
	var configs = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	key, captchaInstance := base64Captcha.GenerateCaptcha("", configs)
	base64string := base64Captcha.CaptchaWriteToBase64Encoding(captchaInstance)
	return key, base64string
}

/**
验证验证码
*/
func ValidateCaptcha(key, captcha string) bool {
	if !helper.CheckDebug() {
		return base64Captcha.VerifyCaptcha(key, captcha)
	}
	return true
}
