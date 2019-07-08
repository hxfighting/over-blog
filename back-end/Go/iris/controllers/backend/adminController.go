package backend

import (
	"blog/helper"
	"blog/models"
	"blog/service"
	"crypto/md5"
	"fmt"
	"github.com/kataras/iris"
)

/**
登录
*/
func Login(ctx iris.Context) {
	admin := models.Admin{}
	err := ctx.ReadJSON(&admin)
	if err != nil {
		response.RenderError(ctx, "数据错误！", nil)
		return
	}
	err = service.Validate.Var(admin.Name, "required,gte=2,lte=15")
	if err != nil {
		response.RenderError(ctx, "请输入合法的用户名！", nil)
		return
	}
	err = service.Validate.Var(admin.Password, "required,gte=6,lte=16")
	if err != nil {
		response.RenderError(ctx, "请输入合法的密码！", nil)
		return
	}
	err = service.Validate.Var(admin.Captcha, "required,len=6")
	if err != nil {
		response.RenderError(ctx, "请输入合法的验证码！", nil)
		return
	}
	err = service.Validate.Var(admin.Key, "required")
	if err != nil {
		response.RenderError(ctx, "请输入合法的验证码key！", nil)
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
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "获取用户信息成功！", user)
}

/**
修改个人信息
*/
func UpdateInfo(ctx iris.Context) {
	admin := models.Admin{}
	err := ctx.ReadJSON(&admin)
	if err != nil {
		response.RenderError(ctx, "数据错误！", nil)
		return
	}
	err = service.Validate.Var(admin.Avatar, "required,url")
	if err != nil {
		response.RenderError(ctx, "请输入正确的头像地址", nil)
		return
	}
	res := helper.VerifyMobileFormat(admin.Mobile)
	if !res {
		response.RenderError(ctx, "请输入正确的电话号码", nil)
		return
	}
	err = service.Validate.Var(admin.Email, "required,email")
	if err != nil {
		response.RenderError(ctx, "请输入正确的邮箱地址", nil)
		return
	}
	err = service.Validate.Var(admin.Name, "required,required,gte=2,lte=30")
	if err != nil {
		response.RenderError(ctx, "请输入正确的邮箱地址", nil)
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
	admin := models.Admin{}
	err := ctx.ReadJSON(&admin)
	if err != nil {
		response.RenderError(ctx, "数据错误！", nil)
		return
	}
	err = service.Validate.Var(admin.Password, "required,gte=6,lte=20")
	if err != nil {
		response.RenderError(ctx, "请输入正确的密码，6到20个字符！", nil)
		return
	}
	err = models.ResetPassword(ctx, &admin)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "修改密码成功！", nil)
}
