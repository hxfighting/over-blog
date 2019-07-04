package backend

import (
	"blog/models"
	"blog/service"
	"github.com/kataras/iris"
)

var response = service.Response{}

/**
登录
*/
func Login(ctx iris.Context) {
	admin := models.Admin{}
	err := ctx.ReadJSON(&admin)
	if err != nil {
		service.Log.Error(err.Error())
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
	token, err := admin.Login()
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	response.RenderSuccess(ctx, "登录成功！", token)
}

func Logout(ctx iris.Context) {

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
