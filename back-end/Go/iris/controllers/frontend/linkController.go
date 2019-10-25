package frontend

import (
	"blog/controllers"
	"blog/models"
	"github.com/kataras/iris"
)

/**
申请友链
*/
func AddLink(ctx iris.Context) {
	link := models.Link{}
	fields := []string{"name", "url", "description"}
	validateFields := []string{"Name", "Description", "Url"}
	_, err := controllers.GetRightModel(ctx, &link, fields, validateFields)
	if err != nil {
		Response.Code = -1
		Response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := models.AddLink(&link)
	if !res {
		Response.Code = -1
		Response.RenderError(ctx, "申请友联失败,请稍后再试!", nil)
		return
	}
	Response.RenderSuccess(ctx, "申请友联成功！", nil)
}
