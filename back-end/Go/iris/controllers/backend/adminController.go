package backend

import (
	"blog/models"
	"blog/service"
	"github.com/kataras/iris"
)

var response = service.Response{}

func Login(ctx iris.Context) {
	admin := models.Admin{}
	err := ctx.ReadJSON(&admin)
	if err != nil {
		service.Log.Error(err.Error())
		response.RenderError(ctx, "数据错误！", nil)
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
