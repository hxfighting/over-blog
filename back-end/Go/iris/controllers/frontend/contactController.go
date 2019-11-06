package frontend

import (
	"blog/controllers"
	"blog/models"
	"blog/service"
	template "blog/views"
	"bytes"
	"github.com/kataras/iris/v12"
)

/**
联系我页面
*/
func GetContactPage(ctx iris.Context) {
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	buffer := new(bytes.Buffer)
	user := GetUser(ctx)
	template.ContactPage(user, buffer)
	_, err := ctx.Write(buffer.Bytes())
	if err != nil {
		service.Log.Error(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("出错啦...")
	}
}

/**
留言
*/
func AddContact(ctx iris.Context) {
	contact := models.Contact{}
	fields := []string{"name", "email", "content"}
	validateFields := []string{"Name", "Email", "Content"}
	_, err := controllers.GetRightModel(ctx, &contact, fields, validateFields)
	if err != nil {
		Response.Code = -1
		Response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := contact.AddContact()
	if !res {
		Response.Code = -1
		Response.RenderError(ctx, "留言失败，请稍后再试！", nil)
		return
	}
	Response.RenderSuccess(ctx, "留言成功！", nil)
}
