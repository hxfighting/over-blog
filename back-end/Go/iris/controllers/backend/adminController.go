package backend

import (
	"blog/controllers"
	"blog/helper"
	"blog/models"
	"blog/service"
	"crypto/md5"
	"fmt"
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
)

/**
登录
*/
func Login(ctx iris.Context) {
	admin := models.Admin{}
	fields := []string{"password", "name", "captcha", "key"}
	validateFields := []string{"Password", "Name", "Captcha"}
	_, err := controllers.GetRightModel(ctx, &admin, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	if service.ValidateCaptcha(admin.Key, admin.Captcha) {
		token, err := admin.Login()
		if err != nil {
			response.RenderError(ctx, err.Error(), nil)
			return
		}
		response.RenderSuccess(ctx, "登录成功！", token)
		return
	} else {
		response.RenderError(ctx, "验证码错误！", nil)
		return
	}
}

/**
退出登录
*/
func Logout(ctx iris.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		response.RenderError(ctx, "退出登录失败，缺少认证token！", nil)
		return
	}
	key := md5.Sum([]byte(authHeader))
	res := service.Redis.Del(fmt.Sprintf("%x", key))
	_, e := res.Result()
	if e != nil {
		response.RenderError(ctx, "退出登录失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "退出登录成功！", nil)
	return
}

/**
获取个人信息
*/
func GetUserInfo(ctx iris.Context) {
	user, err := models.GetUserInfo(ctx)
	if err != nil {
		response.RenderError(ctx, "暂无该用户数据！", nil)
		return
	}
	response.RenderSuccess(ctx, "获取用户信息成功！", user)
}

/**
修改个人信息
*/
func UpdateInfo(ctx iris.Context) {
	admin := models.Admin{}
	fields := []string{"avatar", "name", "email"}
	validateFields := []string{"Avatar", "Name", "Email"}
	data, err := controllers.GetRightModel(ctx, &admin, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	phone := gjson.Get(data, "phone").String()
	if phone == "" {
		response.RenderError(ctx, "请输入正确的电话号码", nil)
		return
	}
	res := helper.VerifyMobileFormat(phone)
	if !res {
		response.RenderError(ctx, "请输入正确的电话号码", nil)
		return
	}
	admin.Mobile = phone
	err = models.UpdateInfo(ctx, &admin)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "修改个人信息成功！", nil)
}

/**
修改密码
*/
func ResetPassword(ctx iris.Context) {
	admin := models.Admin{}
	fields := []string{"password"}
	validateFields := []string{"Password"}
	_, err := controllers.GetRightModel(ctx, &admin, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	err = models.ResetPassword(ctx, &admin)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "修改密码成功！", nil)
}
