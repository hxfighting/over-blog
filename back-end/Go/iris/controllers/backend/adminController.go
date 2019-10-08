package backend

import (
	"blog/helper"
	"blog/models"
	"blog/service"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
)

/**
获取admin model
*/
func getModel(ctx iris.Context, validates []service.BlogValidate) (models.Admin, error) {
	admin := models.Admin{}
	requestData, err := getRequestData(ctx, validates)
	if err != nil {
		return admin, err
	}
	err = mapstructure.Decode(requestData, &admin)
	if err != nil {
		return admin, errors.New("参数错误！")
	}
	return admin, nil
}

/**
登录
*/
func Login(ctx iris.Context) {
	validates := []service.BlogValidate{
		{"name", "string", "required,myString,gte=2,lte=15", "请输入合法的用户名！"},
		{"password", "string", "required,myString,gte=6,lte=16", "请输入合法的密码！"},
		{"captcha", "string", "required,myString,len=6", "请输入合法的验证码！"},
		{"key", "string", "required,myString", "请输入合法的验证码key！"},
	}
	admin, err := getModel(ctx, validates)
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
	validates := []service.BlogValidate{
		{"avatar", "string", "required,myString,url", "请输入正确的头像地址"},
		{"email", "string", "required,myString,email", "请输入正确的邮箱地址"},
		{"name", "string", "required,myString,gte=2,lte=30", "请输入正确的姓名"},
	}
	admin, err := getModel(ctx, validates)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := helper.VerifyMobileFormat(admin.Mobile)
	if !res {
		response.RenderError(ctx, "请输入正确的电话号码", nil)
		return
	}
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
	validates := []service.BlogValidate{
		{"password", "string", "required,myString,gte=6,lte=20", "请输入正确的密码，6到20个字符！"},
	}
	admin, err := getModel(ctx, validates)
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
