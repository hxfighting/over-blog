package backend

import (
	"blog/controllers"
	"blog/models"
	"github.com/kataras/iris"
)

/**
获取错误日志
*/
func GetWebErrorList(ctx iris.Context) {
	pageNum := ctx.URLParamInt64Default("pageNum", 1)
	pageSize := ctx.URLParamInt64Default("pageSize", 10)
	list := models.WebError{}.GetWebErrorList(pageNum, pageSize)
	if list["total"].(int64) > 0 {
		response.RenderSuccess(ctx, "获取错误日志列表成功", list)
		return
	}
	response.RenderError(ctx, "暂无错误日志列表数据", nil)
}

/**
添加错误日志
*/
func AddWebError(ctx iris.Context) {
	webError := models.WebError{}
	fields := []string{"code", "mes", "url", "type"}
	validateFields := []string{"Code", "Mes", "Url", "Type"}
	_, err := controllers.GetRightModel(ctx, &webError, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := webError.AddWebError()
	if !res {
		response.RenderError(ctx, "添加错误日志失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "添加错误日志成功！", nil)
}

/**
删除错误日志
*/
func DeleteWebError(ctx iris.Context) {
	webError := models.WebError{}
	fields := []string{"ids"}
	validateFields := []string{"IDs"}
	_, err := controllers.GetRightModel(ctx, &webError, fields, validateFields)
	if err != nil {
		response.RenderError(ctx, err.Error(), nil)
		return
	}
	res := webError.DeleteWebError()
	if !res {
		response.RenderError(ctx, "删除错误日志失败，请稍后再试！", nil)
		return
	}
	response.RenderSuccess(ctx, "删除错误日志成功！", nil)
}
